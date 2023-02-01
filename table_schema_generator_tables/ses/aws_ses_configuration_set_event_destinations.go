package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSesConfigurationSetEventDestinationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSesConfigurationSetEventDestinationsGenerator{}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetTableName() string {
	return "aws_ses_configuration_set_event_destinations"
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"configuration_set_name",
			"name",
		},
	}
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sesv2

			s := task.ParentRawResult.(*sesv2.GetConfigurationSetOutput)

			output, err := svc.GetConfigurationSetEventDestinations(ctx,
				&sesv2.GetConfigurationSetEventDestinationsInput{
					ConfigurationSetName: s.ConfigurationSetName,
				},
			)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- output.EventDestinations

			return nil
		},
	}
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("email")
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kinesis_firehose_destination").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("KinesisFirehoseDestination")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pinpoint_destination").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PinpointDestination")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sns_destination").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SnsDestination")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("matching_event_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("MatchingEventTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_watch_destination").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CloudWatchDestination")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_set_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ses_configuration_sets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ses_configuration_sets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
	}
}

func (x *TableAwsSesConfigurationSetEventDestinationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
