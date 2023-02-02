package docdb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsDocdbEventCategoriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbEventCategoriesGenerator{}

func (x *TableAwsDocdbEventCategoriesGenerator) GetTableName() string {
	return "aws_docdb_event_categories"
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := &docdb.DescribeEventCategoriesInput{
				Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
			}

			response, err := svc.DescribeEventCategories(ctx, input)
			if err != nil {
				if c.IsNotFoundError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.EventCategoriesMapList
			return nil
		},
	}
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("event_categories").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsDocdbEventCategoriesGenerator) GetSubTables() []*schema.Table {
	return nil
}
