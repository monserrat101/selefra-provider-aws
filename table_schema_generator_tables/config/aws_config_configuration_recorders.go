package config

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsConfigConfigurationRecordersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsConfigConfigurationRecordersGenerator{}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetTableName() string {
	return "aws_config_configuration_recorders"
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)

			resp, err := c.AwsServices().Configservice.DescribeConfigurationRecorders(ctx, &configservice.DescribeConfigurationRecordersInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if len(resp.ConfigurationRecorders) == 0 {
				return nil
			}
			names := make([]string, len(resp.ConfigurationRecorders))
			for i, configurationRecorder := range resp.ConfigurationRecorders {
				names[i] = *configurationRecorder.Name
			}
			status, err := c.AwsServices().Configservice.DescribeConfigurationRecorderStatus(ctx, &configservice.DescribeConfigurationRecorderStatusInput{
				ConfigurationRecorderNames: names,
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for _, configurationRecorder := range resp.ConfigurationRecorders {
				if configurationRecorder.Name == nil {
					continue
				}
				var configurationRecorderStatus types.ConfigurationRecorderStatus
				for _, s := range status.ConfigurationRecordersStatus {
					if s.Name == nil {
						continue
					}
					if *s.Name == *configurationRecorder.Name {
						configurationRecorderStatus = s
						resultChannel <- ConfigurationRecorderWrapper{
							ConfigurationRecorder:		configurationRecorder,
							StatusLastErrorCode:		configurationRecorderStatus.LastErrorCode,
							StatusLastErrorMessage:		configurationRecorderStatus.LastErrorMessage,
							StatusLastStartTime:		configurationRecorderStatus.LastStartTime,
							StatusLastStatus:		configurationRecorderStatus.LastStatus,
							StatusLastStatusChangeTime:	configurationRecorderStatus.LastStatusChangeTime,
							StatusLastStopTime:		configurationRecorderStatus.LastStopTime,
							StatusRecording:		configurationRecorderStatus.Recording,
						}

						break
					}
				}
			}
			return nil
		},
	}
}

type ConfigurationRecorderWrapper struct {
	types.ConfigurationRecorder
	StatusLastErrorCode		*string
	StatusLastErrorMessage		*string
	StatusLastStartTime		*time.Time
	StatusLastStatus		types.RecorderStatus
	StatusLastStatusChangeTime	*time.Time
	StatusLastStopTime		*time.Time
	StatusRecording			bool
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("config")
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_recorder").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConfigurationRecorder")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_last_error_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusLastErrorCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_last_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusLastStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_last_stop_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StatusLastStopTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					cfg := result.(ConfigurationRecorderWrapper)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	"config",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	fmt.Sprintf("config-recorder/%s", aws.ToString(cfg.Name)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_last_error_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusLastErrorMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_last_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StatusLastStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_last_status_change_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StatusLastStatusChangeTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_recording").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("StatusRecording")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsConfigConfigurationRecordersGenerator) GetSubTables() []*schema.Table {
	return nil
}
