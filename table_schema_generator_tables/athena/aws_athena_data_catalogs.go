package athena

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAthenaDataCatalogsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAthenaDataCatalogsGenerator{}

func (x *TableAwsAthenaDataCatalogsGenerator) GetTableName() string {
	return "aws_athena_data_catalogs"
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Athena
			input := athena.ListDataCatalogsInput{}
			for {
				response, err := svc.ListDataCatalogs(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.DataCatalogsSummary, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Athena
					catalogSummary := result.(types.DataCatalogSummary)
					dc, err := svc.GetDataCatalog(ctx, &athena.GetDataCatalogInput{
						Name: catalogSummary.CatalogName,
					})
					if err != nil {

						if aws_client.IsAWSError(err, "InvalidRequestException") && *catalogSummary.CatalogName == "AwsDataCatalog" {
							return types.DataCatalog{Name: catalogSummary.CatalogName, Type: catalogSummary.Type}, nil

						}
						return nil, err
					}
					return *dc.DataCatalog, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("athena")
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					dc := result.(types.DataCatalog)
					return createDataCatalogArn(cl, *dc.Name), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Parameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func createDataCatalogArn(cl *aws_client.Client, catalogName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   "athena",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("datacatalog/%s", catalogName),
	}.String()
}

func (x *TableAwsAthenaDataCatalogsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAthenaDataCatalogDatabasesGenerator{}),
	}
}
