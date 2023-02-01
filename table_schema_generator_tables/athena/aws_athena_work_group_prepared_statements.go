package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAthenaWorkGroupPreparedStatementsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAthenaWorkGroupPreparedStatementsGenerator{}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetTableName() string {
	return "aws_athena_work_group_prepared_statements"
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Athena
			wg := task.ParentRawResult.(types.WorkGroup)
			input := athena.ListPreparedStatementsInput{WorkGroup: wg.Name}
			for {
				response, err := svc.ListPreparedStatements(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.PreparedStatements, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Athena
					wg := task.ParentRawResult.(types.WorkGroup)

					d := result.(types.PreparedStatementSummary)
					dc, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
						WorkGroup:	wg.Name,
						StatementName:	d.StatementName,
					})
					if err != nil {
						return nil, err
					}
					return *dc.PreparedStatement, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("athena")
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("statement_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatementName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("work_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_statement").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("QueryStatement")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("work_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_athena_work_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_athena_work_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsAthenaWorkGroupPreparedStatementsGenerator) GetSubTables() []*schema.Table {
	return nil
}
