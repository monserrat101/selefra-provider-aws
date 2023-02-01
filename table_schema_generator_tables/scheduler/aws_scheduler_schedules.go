package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSchedulerSchedulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSchedulerSchedulesGenerator{}

func (x *TableAwsSchedulerSchedulesGenerator) GetTableName() string {
	return "aws_scheduler_schedules"
}

func (x *TableAwsSchedulerSchedulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSchedulerSchedulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSchedulerSchedulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSchedulerSchedulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := scheduler.ListSchedulesInput{
				MaxResults: aws.Int32(100),
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Scheduler
			paginator := scheduler.NewListSchedulesPaginator(svc, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.Schedules, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Scheduler
					scheduleSummary := result.(types.ScheduleSummary)

					describeTaskDefinitionOutput, err := svc.GetSchedule(ctx, &scheduler.GetScheduleInput{
						Name:		scheduleSummary.Name,
						GroupName:	scheduleSummary.GroupName,
					})
					if err != nil {
						return nil, err
					}
					return describeTaskDefinitionOutput, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsSchedulerSchedulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("scheduler")
}

func (x *TableAwsSchedulerSchedulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modification_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModificationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("flexible_time_window").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FlexibleTimeWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StartDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Target")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EndDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScheduleExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_expression_timezone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScheduleExpressionTimezone")).Build(),
	}
}

func (x *TableAwsSchedulerSchedulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
