package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/mitchellh/mapstructure"
	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSnsTopicsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSnsTopicsGenerator{}

func (x *TableAwsSnsTopicsGenerator) GetTableName() string {
	return "aws_sns_topics"
}

func (x *TableAwsSnsTopicsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSnsTopicsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSnsTopicsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSnsTopicsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sns
			config := sns.ListTopicsInput{}
			for {
				output, err := svc.ListTopics(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.Topics, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sns
					topic := result.(types.Topic)

					attrs, err := svc.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
					if err != nil {
						return nil, err
					}

					t := &Topic{Arn: topic.TopicArn}
					dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: t})
					if err != nil {
						return nil, err
					}
					if err := dec.Decode(attrs.Attributes); err != nil {
						return nil, err
					}
					return t, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

type Topic struct {
	DeliveryPolicy	*string

	DisplayName	*string

	Owner	*string

	Policy	*string

	SubscriptionsConfirmed	*int

	SubscriptionsDeleted	*int

	SubscriptionsPending	*int

	Arn	*string	`mapstructure:"TopicArn"`

	EffectiveDeliveryPolicy	*string

	KmsMasterKeyId	*string

	FifoTopic	*bool

	ContentBasedDeduplication	*bool

	UnknownFields	map[string]any	`mapstructure:",remain"`
}

func (x *TableAwsSnsTopicsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("sns")
}

func (x *TableAwsSnsTopicsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kms_master_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsMasterKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fifo_topic").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("FifoTopic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_based_deduplication").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ContentBasedDeduplication")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unknown_fields").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UnknownFields")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscriptions_confirmed").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SubscriptionsConfirmed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("effective_delivery_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscriptions_pending").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SubscriptionsPending")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Owner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscriptions_deleted").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SubscriptionsDeleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsSnsTopicsGenerator) GetSubTables() []*schema.Table {
	return nil
}
