package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApplicationautoscalingPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApplicationautoscalingPoliciesGenerator{}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetTableName() string {
	return "aws_applicationautoscaling_policies"
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Applicationautoscaling

			config := applicationautoscaling.DescribeScalingPoliciesInput{
				ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
			}
			for {
				output, err := svc.DescribeScalingPolicies(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- output.ScalingPolicies

				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("target_tracking_scaling_policy_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetTrackingScalingPolicyConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarms").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Alarms")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceNamespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scalable_dimension").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScalableDimension")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("step_scaling_policy_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StepScalingPolicyConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyARN")).Build(),
	}
}

func (x *TableAwsApplicationautoscalingPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
