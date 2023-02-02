package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudformationStackResourcesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudformationStackResourcesGenerator{}

func (x *TableAwsCloudformationStackResourcesGenerator) GetTableName() string {
	return "aws_cloudformation_stack_resources"
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			stack := task.ParentRawResult.(types.Stack)
			config := cloudformation.ListStackResourcesInput{
				StackName: stack.StackName,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudformation
			for {
				output, err := svc.ListStackResources(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.StackResourceSummaries
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cloudformation")
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("logical_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LogicalResourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("physical_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PhysicalResourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("drift_information").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DriftInformation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("module_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModuleInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_status_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceStatusReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_cloudformation_stacks_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_cloudformation_stacks.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsCloudformationStackResourcesGenerator) GetSubTables() []*schema.Table {
	return nil
}
