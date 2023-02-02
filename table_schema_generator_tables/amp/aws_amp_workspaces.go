package amp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAmpWorkspacesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAmpWorkspacesGenerator{}

func (x *TableAwsAmpWorkspacesGenerator) GetTableName() string {
	return "aws_amp_workspaces"
}

func (x *TableAwsAmpWorkspacesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAmpWorkspacesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAmpWorkspacesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAmpWorkspacesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Amp

			p := amp.NewListWorkspacesPaginator(svc, &amp.ListWorkspacesInput{MaxResults: aws.Int32(int32(1000))})
			for p.HasMorePages() {
				out, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, out.Workspaces, func(result any) (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Amp

					out, err := svc.DescribeWorkspace(ctx,
						&amp.DescribeWorkspaceInput{WorkspaceId: result.(types.WorkspaceSummary).WorkspaceId},
					)
					if err != nil {
						return nil, err
					}
					return out.Workspace, nil

				})
			}

			return nil
		},
	}
}

func (x *TableAwsAmpWorkspacesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("amp")
}

func (x *TableAwsAmpWorkspacesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("logging_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Amp

					out, err := svc.DescribeLoggingConfiguration(ctx,
						&amp.DescribeLoggingConfigurationInput{WorkspaceId: result.(*types.WorkspaceDescription).WorkspaceId},
					)
					if err != nil {
						return nil, err
					}

					return out.LoggingConfiguration, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_manager_definition").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Amp

					out, err := svc.DescribeAlertManagerDefinition(ctx,
						&amp.DescribeAlertManagerDefinitionInput{WorkspaceId: result.(*types.WorkspaceDescription).WorkspaceId},
					)
					if err != nil {
						return nil, err
					}

					return out.AlertManagerDefinition, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workspace_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkspaceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Alias")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("prometheus_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrometheusEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsAmpWorkspacesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAmpRuleGroupsNamespacesGenerator{}),
	}
}
