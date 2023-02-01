package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbPendingMaintenanceActionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbPendingMaintenanceActionsGenerator{}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetTableName() string {
	return "aws_docdb_pending_maintenance_actions"
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := &docdb.DescribePendingMaintenanceActionsInput{
				Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
			}

			p := docdb.NewDescribePendingMaintenanceActionsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.PendingMaintenanceActions
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_maintenance_action_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingMaintenanceActionDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsDocdbPendingMaintenanceActionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
