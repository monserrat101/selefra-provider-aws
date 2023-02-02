package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayDomainNameBasePathMappingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayDomainNameBasePathMappingsGenerator{}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetTableName() string {
	return "aws_apigateway_domain_name_base_path_mappings"
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.DomainName)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			config := apigateway.GetBasePathMappingsInput{DomainName: r.DomainName}
			for p := apigateway.NewGetBasePathMappingsPaginator(svc, &config); p.HasMorePages(); {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("rest_api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RestApiId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigateway_domain_names_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigateway_domain_names.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("base_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BasePath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stage").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Stage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					domain := task.ParentRawResult.(types.DomainName)
					mapping := result.(types.BasePathMapping)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/domainnames/%s/basepathmappings/%s", aws.ToString(domain.DomainName), aws.ToString(mapping.BasePath)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
	}
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
