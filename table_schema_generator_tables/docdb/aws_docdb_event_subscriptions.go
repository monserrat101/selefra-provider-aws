package docdb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsDocdbEventSubscriptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbEventSubscriptionsGenerator{}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetTableName() string {
	return "aws_docdb_event_subscriptions"
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := &docdb.DescribeEventSubscriptionsInput{
				Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
			}

			p := docdb.NewDescribeEventSubscriptionsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.EventSubscriptionsList
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("event_categories_list").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EventCategoriesList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_subscription_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventSubscriptionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cust_subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustSubscriptionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sns_topic_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnsTopicArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_ids_list").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SourceIdsList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_creation_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscriptionCreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_aws_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomerAwsId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsDocdbEventSubscriptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
