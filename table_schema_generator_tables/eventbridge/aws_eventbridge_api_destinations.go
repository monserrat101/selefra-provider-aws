package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEventbridgeApiDestinationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEventbridgeApiDestinationsGenerator{}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetTableName() string {
	return "aws_eventbridge_api_destinations"
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input eventbridge.ListApiDestinationsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Eventbridge
			for {
				response, err := svc.ListApiDestinations(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ApiDestinations
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("events")
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("invocation_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InvocationEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invocation_rate_limit_per_second").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("InvocationRateLimitPerSecond")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiDestinationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_destination_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiDestinationState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ConnectionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_method").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HttpMethod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_destination_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiDestinationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsEventbridgeApiDestinationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
