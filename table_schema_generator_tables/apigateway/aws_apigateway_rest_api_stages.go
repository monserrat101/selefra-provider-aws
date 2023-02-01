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

type TableAwsApigatewayRestApiStagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayRestApiStagesGenerator{}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetTableName() string {
	return "aws_apigateway_rest_api_stages"
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.RestApi)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			config := apigateway.GetStagesInput{RestApiId: r.Id}

			response, err := svc.GetStages(ctx, &config)
			if err != nil {
				if c.IsNotFoundError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.Item

			return nil
		},
	}
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					s := result.(types.Stage)
					rapi := task.ParentRawResult.(types.RestApi)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/restapis/%s/stages/%s", aws.ToString(rapi.Id), aws.ToString(s.StageName)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rest_api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_size").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheClusterSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("canary_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CanarySettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("documentation_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DocumentationVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Variables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stage_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StageName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_acl_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebAclArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_log_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AccessLogSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheClusterStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_certificate_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientCertificateId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigateway_rest_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigateway_rest_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("CacheClusterEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("method_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MethodSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tracing_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TracingEnabled")).Build(),
	}
}

func (x *TableAwsApigatewayRestApiStagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
