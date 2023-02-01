package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsOrganizationResourcePoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsOrganizationResourcePoliciesGenerator{}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetTableName() string {
	return "aws_organization_resource_policies"
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			o, err := c.AwsServices().Organizations.DescribeResourcePolicy(ctx, &organizations.DescribeResourcePolicyInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- o.ResourcePolicy
			return nil
		},
	}
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("content").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Content")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_policy_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourcePolicySummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsOrganizationResourcePoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
