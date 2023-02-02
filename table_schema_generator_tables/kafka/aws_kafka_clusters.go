package kafka

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsKafkaClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKafkaClustersGenerator{}

func (x *TableAwsKafkaClustersGenerator) GetTableName() string {
	return "aws_kafka_clusters"
}

func (x *TableAwsKafkaClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKafkaClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKafkaClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsKafkaClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input kafka.ListClustersV2Input
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kafka
			for {
				response, err := svc.ListClustersV2(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.ClusterInfoList, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Kafka
					var input kafka.DescribeClusterV2Input = describeClustersInput(result)
					output, err := svc.DescribeClusterV2(ctx, &input)
					if err != nil {
						return nil, err
					}
					return output.ClusterInfo, nil

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

func describeClustersInput(result any) kafka.DescribeClusterV2Input {
	return kafka.DescribeClusterV2Input{
		ClusterArn: result.(types.Cluster).ClusterArn,
	}
}

func (x *TableAwsKafkaClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kafka")
}

func (x *TableAwsKafkaClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serverless").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Serverless")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_operation_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActiveOperationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CurrentVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioned").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Provisioned")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
	}
}

func (x *TableAwsKafkaClustersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsKafkaNodesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsKafkaClusterOperationsGenerator{}),
	}
}
