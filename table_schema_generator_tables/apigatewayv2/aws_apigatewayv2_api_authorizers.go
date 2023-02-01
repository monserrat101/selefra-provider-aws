package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayv2ApiAuthorizersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApiAuthorizersGenerator{}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetTableName() string {
	return "aws_apigatewayv2_api_authorizers"
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Api)
			config := apigatewayv2.GetAuthorizersInput{
				ApiId: r.ApiId,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetAuthorizers(ctx, &config)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_validation_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityValidationExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_credentials_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerCredentialsArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_payload_format_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerPayloadFormatVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigatewayv2_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigatewayv2_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("jwt_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("JwtConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_simple_responses").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableSimpleResponses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					r := result.(types.Authorizer)
					p := task.ParentRawResult.(types.Api)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/apis/%s/authorizers/%s", aws.ToString(p.ApiId), aws.ToString(r.AuthorizerId)),
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
		table_schema_generator.NewColumnBuilder().ColumnName("identity_source").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("IdentitySource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
	}
}

func (x *TableAwsApigatewayv2ApiAuthorizersGenerator) GetSubTables() []*schema.Table {
	return nil
}
