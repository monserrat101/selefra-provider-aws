package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsServicecatalogProductsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsServicecatalogProductsGenerator{}

func (x *TableAwsServicecatalogProductsGenerator) GetTableName() string {
	return "aws_servicecatalog_products"
}

func (x *TableAwsServicecatalogProductsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsServicecatalogProductsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsServicecatalogProductsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsServicecatalogProductsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Servicecatalog

			listInput := new(servicecatalog.SearchProductsAsAdminInput)
			for {
				output, err := svc.SearchProductsAsAdmin(ctx, listInput)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- output.ProductViewDetails

				if aws.ToString(output.NextPageToken) == "" {
					break
				}
				listInput.PageToken = output.NextPageToken
			}

			return nil
		},
	}
}

func (x *TableAwsServicecatalogProductsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("servicecatalog")
}

func (x *TableAwsServicecatalogProductsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("product_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_view_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProductViewSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_connection").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceConnection")).Build(),
	}
}

func (x *TableAwsServicecatalogProductsGenerator) GetSubTables() []*schema.Table {
	return nil
}
