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

type TableAwsKafkaClusterOperationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKafkaClusterOperationsGenerator{}

func (x *TableAwsKafkaClusterOperationsGenerator) GetTableName() string {
	return "aws_kafka_cluster_operations"
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			if task.ParentRawResult.(*types.Cluster).ClusterType == types.ClusterTypeServerless {

				return nil
			}

			var input = getListClusterOperationsInput(task)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kafka
			for {
				response, err := svc.ListClusterOperations(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ClusterOperationInfoList
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func getListClusterOperationsInput(task *schema.DataSourcePullTask) kafka.ListClusterOperationsInput {
	return kafka.ListClusterOperationsInput{
		ClusterArn:	task.ParentRawResult.(*types.Cluster).ClusterArn,
		MaxResults:	100,
	}
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kafka")
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("operation_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_request_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientRequestId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ErrorInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_kafka_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_kafka_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_cluster_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetClusterInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_steps").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OperationSteps")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_cluster_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceClusterInfo")).Build(),
	}
}

func (x *TableAwsKafkaClusterOperationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
