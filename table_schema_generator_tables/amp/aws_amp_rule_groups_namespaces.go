package amp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAmpRuleGroupsNamespacesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAmpRuleGroupsNamespacesGenerator{}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetTableName() string {
	return "aws_amp_rule_groups_namespaces"
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Amp

			p := amp.NewListRuleGroupsNamespacesPaginator(svc,
				&amp.ListRuleGroupsNamespacesInput{
					WorkspaceId:	task.ParentRawResult.(*types.WorkspaceDescription).WorkspaceId,
					MaxResults:	aws.Int32(int32(1000)),
				},
			)
			for p.HasMorePages() {
				out, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, out.RuleGroupsNamespaces, func(result any) (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Amp

					out, err := svc.DescribeRuleGroupsNamespace(ctx,
						&amp.DescribeRuleGroupsNamespaceInput{WorkspaceId: task.ParentRawResult.(*types.WorkspaceDescription).WorkspaceId},
					)
					if err != nil {
						return nil, err
					}
					return out.RuleGroupsNamespace, nil

				})
			}

			return nil
		},
	}
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data").ColumnType(schema.ColumnTypeByteArray).
			Extractor(column_value_extractor.StructSelector("Data")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workspace_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_amp_workspaces_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_amp_workspaces.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsAmpRuleGroupsNamespacesGenerator) GetSubTables() []*schema.Table {
	return nil
}
