package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticsearchVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticsearchVersionsGenerator{}

func (x *TableAwsElasticsearchVersionsGenerator) GetTableName() string {
	return "aws_elasticsearch_versions"
}

func (x *TableAwsElasticsearchVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticsearchVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticsearchVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"version",
		},
	}
}

func (x *TableAwsElasticsearchVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

			p := elasticsearchservice.NewListElasticsearchVersionsPaginator(svc,
				&elasticsearchservice.ListElasticsearchVersionsInput{MaxResults: 100},
			)
			for p.HasMorePages() {
				out, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- out.ElasticsearchVersions
			}

			return nil
		},
	}
}

func (x *TableAwsElasticsearchVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("es")
}

func (x *TableAwsElasticsearchVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					return result.(string), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_types").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

					var instanceTypes []types.ESPartitionInstanceType
					p := elasticsearchservice.NewListElasticsearchInstanceTypesPaginator(svc,
						&elasticsearchservice.ListElasticsearchInstanceTypesInput{
							ElasticsearchVersion: aws.String(result.(string)),
						},
					)
					for p.HasMorePages() {
						out, err := p.NextPage(ctx)
						if err != nil {
							return nil, err
						}

						instanceTypes = append(instanceTypes, out.ElasticsearchInstanceTypes...)
					}

					return instanceTypes, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsElasticsearchVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
