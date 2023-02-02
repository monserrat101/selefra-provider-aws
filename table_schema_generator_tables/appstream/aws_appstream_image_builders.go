package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppstreamImageBuildersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamImageBuildersGenerator{}

func (x *TableAwsAppstreamImageBuildersGenerator) GetTableName() string {
	return "aws_appstream_image_builders"
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input appstream.DescribeImageBuildersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeImageBuilders(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ImageBuilders

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_endpoints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AccessEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_join_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DomainJoinInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_builder_errors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageBuilderErrors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_access_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkAccessConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_change_reason").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateChangeReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_default_internet_access").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableDefaultInternetAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Platform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("appstream_agent_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AppstreamAgentVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IamRoleArn")).Build(),
	}
}

func (x *TableAwsAppstreamImageBuildersGenerator) GetSubTables() []*schema.Table {
	return nil
}
