package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEventbridgeReplaysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEventbridgeReplaysGenerator{}

func (x *TableAwsEventbridgeReplaysGenerator) GetTableName() string {
	return "aws_eventbridge_replays"
}

func (x *TableAwsEventbridgeReplaysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEventbridgeReplaysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEventbridgeReplaysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEventbridgeReplaysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input eventbridge.ListReplaysInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Eventbridge
			for {
				response, err := svc.ListReplays(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Replays
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEventbridgeReplaysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("events")
}

func (x *TableAwsEventbridgeReplaysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_last_replayed_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EventLastReplayedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replay_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EventStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replay_end_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ReplayEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replay_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ReplayStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_end_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EventEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_source_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventSourceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)

					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"events",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"replay/" + aws.ToString(result.(types.Replay).ReplayName),
					}

					return a.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
	}
}

func (x *TableAwsEventbridgeReplaysGenerator) GetSubTables() []*schema.Table {
	return nil
}
