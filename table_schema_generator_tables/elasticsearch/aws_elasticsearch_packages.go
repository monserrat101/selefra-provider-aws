package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticsearchPackagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticsearchPackagesGenerator{}

func (x *TableAwsElasticsearchPackagesGenerator) GetTableName() string {
	return "aws_elasticsearch_packages"
}

func (x *TableAwsElasticsearchPackagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticsearchPackagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticsearchPackagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"id",
		},
	}
}

func (x *TableAwsElasticsearchPackagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

			p := elasticsearchservice.NewDescribePackagesPaginator(svc, nil)
			for p.HasMorePages() {
				out, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- out.PackageDetailsList
			}

			return nil
		},
	}
}

func (x *TableAwsElasticsearchPackagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("es")
}

func (x *TableAwsElasticsearchPackagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_package_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailablePackageVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ErrorDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageID")).Build(),
	}
}

func (x *TableAwsElasticsearchPackagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
