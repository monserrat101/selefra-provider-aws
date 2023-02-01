package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFrauddetectorBatchImportsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFrauddetectorBatchImportsGenerator{}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetTableName() string {
	return "aws_frauddetector_batch_imports"
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := frauddetector.NewGetBatchImportJobsPaginator(client.(*aws_client.Client).AwsServices().Frauddetector, nil)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.BatchImports
			}
			return nil
		},
	}
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("frauddetector")
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("iam_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IamRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("JobId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processed_records_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProcessedRecordsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("completion_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompletionTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_type_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventTypeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutputPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InputPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_records_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TotalRecordsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failed_records_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("FailedRecordsCount")).Build(),
	}
}

func (x *TableAwsFrauddetectorBatchImportsGenerator) GetSubTables() []*schema.Table {
	return nil
}
