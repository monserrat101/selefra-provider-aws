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

type TableAwsApigatewayv2ApiRoutesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApiRoutesGenerator{}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetTableName() string {
	return "aws_apigatewayv2_api_routes"
}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Api)
			config := apigatewayv2.GetRoutesInput{
				ApiId: r.ApiId,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetRoutes(ctx, &config)

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

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					r := result.(types.Route)
					p := task.ParentRawResult.(types.Api)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/apis/%s/routes/%s", aws.ToString(p.ApiId), aws.ToString(r.RouteId)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigatewayv2_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigatewayv2_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorization_scopes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AuthorizationScopes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_models").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RequestModels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RequestParameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RouteId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_response_selection_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RouteResponseSelectionExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_key_required").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ApiKeyRequired")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RouteKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorization_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorizer_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_gateway_managed").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ApiGatewayManaged")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_selection_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelSelectionExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Target")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsApigatewayv2ApiRoutesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiRouteResponsesGenerator{}),
	}
}
