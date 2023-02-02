package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudwatchlogsResourcePoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudwatchlogsResourcePoliciesGenerator{}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetTableName() string {
	return "aws_cloudwatchlogs_resource_policies"
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"policy_name",
		},
	}
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config cloudwatchlogs.DescribeResourcePoliciesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudwatchlogs
			for {
				response, err := svc.DescribeResourcePolicies(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ResourcePolicies
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("logs")
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
	}
}

func (x *TableAwsCloudwatchlogsResourcePoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
