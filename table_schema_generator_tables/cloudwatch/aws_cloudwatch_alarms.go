package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudwatchAlarmsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudwatchAlarmsGenerator{}

func (x *TableAwsCloudwatchAlarmsGenerator) GetTableName() string {
	return "aws_cloudwatch_alarms"
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config cloudwatch.DescribeAlarmsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudwatch
			for {
				response, err := svc.DescribeAlarms(ctx, &config)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.MetricAlarms
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("logs")
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason_data").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateReasonData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("treat_missing_data").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TreatMissingData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actions_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ActionsEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evaluate_low_sample_count_percentile").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EvaluateLowSampleCountPercentile")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evaluation_periods").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EvaluationPeriods")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evaluation_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EvaluationState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("Threshold")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarm_actions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AlarmActions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_updated_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StateUpdatedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold_metric_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ThresholdMetricId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AlarmArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarm_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AlarmDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dimensions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					alarm := result.(types.MetricAlarm)
					dimensions := make(map[string]*string)
					for _, d := range alarm.Dimensions {
						dimensions[*d.Name] = d.Value
					}
					return dimensions, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_value").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateValue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Unit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarm_configuration_updated_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("AlarmConfigurationUpdatedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MetricName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarm_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AlarmArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_statistic").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExtendedStatistic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("insufficient_data_actions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("InsufficientDataActions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Period")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statistic").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Statistic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarm_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AlarmName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comparison_operator").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ComparisonOperator")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("datapoints_to_alarm").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DatapointsToAlarm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ok_actions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("OKActions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_transitioned_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StateTransitionedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metrics").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metrics")).Build(),
	}
}

func (x *TableAwsCloudwatchAlarmsGenerator) GetSubTables() []*schema.Table {
	return nil
}
