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

type TableAwsApigatewayRestApiAuthorizersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayRestApiAuthorizersGenerator{}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetTableName() string {
	return "aws_apigateway_rest_api_authorizers"
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.RestApi)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			config := apigateway.GetAuthorizersInput{RestApiId: r.Id}
			for {
				response, err := svc.GetAuthorizers(ctx, &config)
				if err != nil {
					if c.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
				if aws.ToString(response.Position) == "" {
					break
				}
				config.Position = response.Position
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("rest_api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigateway_rest_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigateway_rest_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					auth := result.(types.Authorizer)
					rapi := task.ParentRawResult.(types.RestApi)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/restapis/%s/authorizers/%s", aws.ToString(rapi.Id), aws.ToString(auth.Id)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_result_ttl_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AuthorizerResultTtlInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_validation_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityValidationExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_ar_ns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ProviderARNs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_credentials").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerCredentials")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentitySource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsApigatewayRestApiAuthorizersGenerator) GetSubTables() []*schema.Table {
	return nil
}
