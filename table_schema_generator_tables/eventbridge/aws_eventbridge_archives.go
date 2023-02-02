package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEventbridgeArchivesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEventbridgeArchivesGenerator{}

func (x *TableAwsEventbridgeArchivesGenerator) GetTableName() string {
	return "aws_eventbridge_archives"
}

func (x *TableAwsEventbridgeArchivesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEventbridgeArchivesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEventbridgeArchivesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEventbridgeArchivesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input eventbridge.ListArchivesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Eventbridge
			for {
				response, err := svc.ListArchives(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Archives
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEventbridgeArchivesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("events")
}

func (x *TableAwsEventbridgeArchivesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("archive_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ArchiveName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_days").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RetentionDays")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SizeBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EventCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_source_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventSourceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateReason")).Build(),
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
						Resource:	"archive/" + aws.ToString(result.(types.Archive).ArchiveName),
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
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsEventbridgeArchivesGenerator) GetSubTables() []*schema.Table {
	return nil
}
