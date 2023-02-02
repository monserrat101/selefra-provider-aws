package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailAlarmsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailAlarmsGenerator{}

func (x *TableAwsLightsailAlarmsGenerator) GetTableName() string {
	return "aws_lightsail_alarms"
}

func (x *TableAwsLightsailAlarmsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailAlarmsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailAlarmsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLightsailAlarmsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetAlarmsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetAlarms(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Alarms
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailAlarmsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailAlarmsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("comparison_operator").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ComparisonOperator")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Period")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("Threshold")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("NotificationEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_triggers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("NotificationTriggers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("contact_protocols").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ContactProtocols")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MetricName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitored_resource_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MonitoredResourceInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statistic").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Statistic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("treat_missing_data").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TreatMissingData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("datapoints_to_alarm").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DatapointsToAlarm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evaluation_periods").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EvaluationPeriods")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SupportCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Unit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsLightsailAlarmsGenerator) GetSubTables() []*schema.Table {
	return nil
}
