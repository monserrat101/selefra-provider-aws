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

type TableAwsApigatewayv2ApiStagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApiStagesGenerator{}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetTableName() string {
	return "aws_apigatewayv2_api_stages"
}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Api)
			config := apigatewayv2.GetStagesInput{
				ApiId: r.ApiId,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetStages(ctx, &config)

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

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					r := result.(types.Stage)
					p := task.ParentRawResult.(types.Api)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/apis/%s/stages/%s", aws.ToString(p.ApiId), aws.ToString(r.StageName)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_gateway_managed").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ApiGatewayManaged")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_deploy").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoDeploy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_certificate_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientCertificateId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stage_variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StageVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_log_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AccessLogSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigatewayv2_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigatewayv2_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_deployment_status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastDeploymentStatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RouteSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stage_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StageName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_route_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultRouteSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsApigatewayv2ApiStagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
