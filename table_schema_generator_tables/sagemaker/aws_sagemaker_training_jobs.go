package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSagemakerTrainingJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSagemakerTrainingJobsGenerator{}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetTableName() string {
	return "aws_sagemaker_training_jobs"
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sagemaker
			config := sagemaker.ListTrainingJobsInput{}

			for {
				response, err := svc.ListTrainingJobs(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.TrainingJobSummaries, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sagemaker
					n := result.(types.TrainingJobSummary)
					config := sagemaker.DescribeTrainingJobInput{
						TrainingJobName: n.TrainingJobName,
					}
					response, err := svc.DescribeTrainingJob(ctx, &config)
					if err != nil {
						return nil, err
					}
					return response, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("api.sagemaker")
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("algorithm_specification").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AlgorithmSpecification")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stopping_condition").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StoppingCondition")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("checkpoint_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CheckpointConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profiler_rule_configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProfilerRuleConfigurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_status_transitions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecondaryStatusTransitions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecondaryStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_job_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrainingJobName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hyper_parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HyperParameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retry_strategy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RetryStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("warm_pool_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("WarmPoolStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profiling_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProfilingStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_artifacts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModelArtifacts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_job_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrainingJobStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_inter_container_traffic_encryption").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableInterContainerTrafficEncryption")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_network_isolation").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableNetworkIsolation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("final_metric_data_list").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FinalMetricDataList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_end_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("TrainingEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_time_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TrainingTimeInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("debug_rule_evaluation_statuses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DebugRuleEvaluationStatuses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_managed_spot_training").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableManagedSpotTraining")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_data_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InputDataConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_data_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutputDataConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The tags associated with the model.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_ml_job_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoMLJobArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labeling_job_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelingJobArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profiler_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProfilerConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("TrainingStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("debug_hook_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DebugHookConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profiler_rule_evaluation_statuses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProfilerRuleEvaluationStatuses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("experiment_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExperimentConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tensor_board_output_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TensorBoardOutputConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tuning_job_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TuningJobArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrainingJobArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_job_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrainingJobArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billable_time_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("BillableTimeInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("debug_rule_configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DebugRuleConfigurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Environment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleArn")).Build(),
	}
}

func (x *TableAwsSagemakerTrainingJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
