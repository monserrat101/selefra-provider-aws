package redshift

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRedshiftEventSubscriptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRedshiftEventSubscriptionsGenerator{}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetTableName() string {
	return "aws_redshift_event_subscriptions"
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Redshift
			var params redshift.DescribeEventSubscriptionsInput
			params.MaxRecords = aws.Int32(100)
			for {
				result, err := svc.DescribeEventSubscriptions(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.EventSubscriptionsList
				if aws.ToString(result.Marker) == "" {
					break
				}
				params.Marker = result.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("redshift")
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cust_subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustSubscriptionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sns_topic_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnsTopicArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("`ARN of the event subscription.`").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					sub := result.(types.EventSubscription)
					return eventSubscriptionARN(client.(*aws_client.Client), *sub.CustSubscriptionId), nil

				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_categories_list").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EventCategoriesList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_ids_list").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SourceIdsList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("SubscriptionCreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_aws_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomerAwsId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func eventSubscriptionARN(cl *aws_client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   "redshift",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("eventsubscription:%s", name),
	}.String()
}

func (x *TableAwsRedshiftEventSubscriptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
