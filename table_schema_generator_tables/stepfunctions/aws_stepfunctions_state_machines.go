package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsStepfunctionsStateMachinesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsStepfunctionsStateMachinesGenerator{}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetTableName() string {
	return "aws_stepfunctions_state_machines"
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Sfn
			config := sfn.ListStateMachinesInput{}
			paginator := sfn.NewListStateMachinesPaginator(svc, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.StateMachines, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sfn
					sm := result.(types.StateMachineListItem)

					stateMachineDetails, err := svc.DescribeStateMachine(ctx, &sfn.DescribeStateMachineInput{StateMachineArn: sm.StateMachineArn})
					if err != nil {
						return nil, err
					}
					return stateMachineDetails, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("states")
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Label")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateMachineArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("definition").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Definition")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tracing_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TracingConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_machine_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateMachineArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LoggingConfiguration")).Build(),
	}
}

func (x *TableAwsStepfunctionsStateMachinesGenerator) GetSubTables() []*schema.Table {
	return nil
}
