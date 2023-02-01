package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcsClusterTasksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcsClusterTasksGenerator{}

func (x *TableAwsEcsClusterTasksGenerator) GetTableName() string {
	return "aws_ecs_cluster_tasks"
}

func (x *TableAwsEcsClusterTasksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcsClusterTasksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcsClusterTasksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEcsClusterTasksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cluster, ok := task.ParentRawResult.(types.Cluster)
			if !ok {
				maybeError := fmt.Errorf("expected to have types.Cluster but got %T", task.ParentRawResult)
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			}

			svc := client.(*aws_client.Client).AwsServices().Ecs
			config := ecs.ListTasksInput{
				Cluster: cluster.ClusterArn,
			}
			for {
				listTasks, err := svc.ListTasks(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if len(listTasks.TaskArns) == 0 {
					return nil
				}
				describeServicesInput := ecs.DescribeTasksInput{
					Cluster:	cluster.ClusterArn,
					Tasks:		listTasks.TaskArns,
					Include:	[]types.TaskField{types.TaskFieldTags},
				}
				describeTasks, err := svc.DescribeTasks(ctx, &describeServicesInput)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- describeTasks.Tasks

				if listTasks.NextToken == nil {
					break
				}
				config.NextToken = listTasks.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEcsClusterTasksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ecs")
}

func (x *TableAwsEcsClusterTasksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("containers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Containers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Memory")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pull_started_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("PullStartedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("started_by").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StartedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stopping_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StoppingAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_provider_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CapacityProviderName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container_instance_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ContainerInstanceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inference_accelerators").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InferenceAccelerators")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("overrides").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Overrides")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pull_stopped_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("PullStoppedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("started_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StartedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("desired_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DesiredStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stopped_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StoppedReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ecs_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ecs_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connectivity_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ConnectivityAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ephemeral_storage").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EphemeralStorage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_stopped_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ExecutionStoppedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HealthStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_protection").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Ecs
					task := result.(types.Task)
					resp, err := svc.GetTaskProtection(ctx, &ecs.GetTaskProtectionInput{
						Cluster:	task.ClusterArn,
						Tasks:		[]string{aws.ToString(task.TaskArn)},
					})
					if err != nil {
						return nil, err
					}
					if len(resp.Failures) > 0 {

						return nil, nil
					}
					return resp.ProtectedTasks, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Group")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attachments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connectivity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Connectivity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_execute_command").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableExecuteCommand")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stopped_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StoppedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LaunchType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stop_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StopCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Cpu")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_definition_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskDefinitionArn")).Build(),
	}
}

func (x *TableAwsEcsClusterTasksGenerator) GetSubTables() []*schema.Table {
	return nil
}
