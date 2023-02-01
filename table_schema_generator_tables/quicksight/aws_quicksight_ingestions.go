package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightIngestionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightIngestionsGenerator{}

func (x *TableAwsQuicksightIngestionsGenerator) GetTableName() string {
	return "aws_quicksight_ingestions"
}

func (x *TableAwsQuicksightIngestionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightIngestionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightIngestionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
			"data_set_arn",
		},
	}
}

func (x *TableAwsQuicksightIngestionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			item := task.ParentRawResult.(types.DataSetSummary)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListIngestionsInput{
				AwsAccountId:	aws.String(cl.AccountID),
				DataSetId:	item.DataSetId,
			}
			paginator := quicksight.NewListIngestionsPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.Ingestions
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightIngestionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightIngestionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingestion_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IngestionStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingestion_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IngestionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queue_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("QueueInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ErrorInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingestion_size_in_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("IngestionSizeInBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingestion_time_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("IngestionTimeInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequestType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_set_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_quicksight_data_sets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_quicksight_data_sets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequestSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("row_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RowInfo")).Build(),
	}
}

func (x *TableAwsQuicksightIngestionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
