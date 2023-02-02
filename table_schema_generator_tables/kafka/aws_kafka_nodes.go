package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsKafkaNodesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKafkaNodesGenerator{}

func (x *TableAwsKafkaNodesGenerator) GetTableName() string {
	return "aws_kafka_nodes"
}

func (x *TableAwsKafkaNodesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKafkaNodesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKafkaNodesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsKafkaNodesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			if task.ParentRawResult.(*types.Cluster).ClusterType == types.ClusterTypeServerless {

				return nil
			}
			if task.ParentRawResult.(*types.Cluster).State == types.ClusterStateCreating {

				return nil
			}

			var input = getListNodesInput(task)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kafka
			for {
				response, err := svc.ListNodes(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.NodeInfoList
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func getListNodesInput(task *schema.DataSourcePullTask) kafka.ListNodesInput {
	return kafka.ListNodesInput{
		ClusterArn:	task.ParentRawResult.(*types.Cluster).ClusterArn,
		MaxResults:	100,
	}
}

func (x *TableAwsKafkaNodesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kafka")
}

func (x *TableAwsKafkaNodesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("broker_node_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BrokerNodeInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zookeeper_node_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ZookeeperNodeInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("added_to_cluster_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AddedToClusterTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_kafka_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_kafka_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsKafkaNodesGenerator) GetSubTables() []*schema.Table {
	return nil
}
