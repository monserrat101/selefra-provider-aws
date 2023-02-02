package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbEventsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbEventsGenerator{}

func (x *TableAwsDocdbEventsGenerator) GetTableName() string {
	return "aws_docdb_events"
}

func (x *TableAwsDocdbEventsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbEventsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbEventsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDocdbEventsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := &docdb.DescribeEventsInput{}

			p := docdb.NewDescribeEventsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Events
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbEventsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbEventsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Date")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_categories").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EventCategories")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceIdentifier")).Build(),
	}
}

func (x *TableAwsDocdbEventsGenerator) GetSubTables() []*schema.Table {
	return nil
}
