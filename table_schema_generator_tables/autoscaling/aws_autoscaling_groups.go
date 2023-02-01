package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAutoscalingGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAutoscalingGroupsGenerator{}

func (x *TableAwsAutoscalingGroupsGenerator) GetTableName() string {
	return "aws_autoscaling_groups"
}

func (x *TableAwsAutoscalingGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAutoscalingGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAutoscalingGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAutoscalingGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Autoscaling
			processGroupsBundle := func(groups []types.AutoScalingGroup) error {
				input := autoscaling.DescribeNotificationConfigurationsInput{
					MaxRecords: aws.Int32(100),
				}
				for _, h := range groups {
					input.AutoScalingGroupNames = append(input.AutoScalingGroupNames, *h.AutoScalingGroupName)
				}
				var configurations []types.NotificationConfiguration
				for {
					output, err := svc.DescribeNotificationConfigurations(ctx, &input, func(o *autoscaling.Options) {
						o.Region = c.Region
					})
					if err != nil {
						return err
					}
					configurations = append(configurations, output.NotificationConfigurations...)
					if aws.ToString(output.NextToken) == "" {
						break
					}
					input.NextToken = output.NextToken
				}
				for _, gr := range groups {
					wrapper := AutoScalingGroupWrapper{
						AutoScalingGroup:		gr,
						NotificationConfigurations:	getNotificationConfigurationByGroupName(*gr.AutoScalingGroupName, configurations),
					}
					resultChannel <- wrapper
				}
				return nil
			}

			config := autoscaling.DescribeAutoScalingGroupsInput{}
			for {
				output, err := svc.DescribeAutoScalingGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				groups := output.AutoScalingGroups
				for i := 0; i < len(groups); i += 255 {
					end := i + 255

					if end > len(groups) {
						end = len(groups)
					}
					t := groups[i:end]
					err := processGroupsBundle(t)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
				}

				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func getNotificationConfigurationByGroupName(name string, set []types.NotificationConfiguration) []types.NotificationConfiguration {
	var response []types.NotificationConfiguration
	for _, s := range set {
		if *s.AutoScalingGroupName == name {
			response = append(response, s)
		}
	}
	return response
}

func (x *TableAwsAutoscalingGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("autoscaling")
}

func (x *TableAwsAutoscalingGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_group").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AutoScalingGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NotificationConfigurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(AutoScalingGroupWrapper)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Autoscaling
					config := autoscaling.DescribeLoadBalancersInput{AutoScalingGroupName: p.AutoScalingGroupName}
					j := map[string]any{}
					for {
						output, err := svc.DescribeLoadBalancers(ctx, &config)
						if err != nil {
							if isAutoScalingGroupNotExistsError(err) {
								return nil, nil
							}
							return nil, err
						}
						for _, lb := range output.LoadBalancers {
							j[*lb.LoadBalancerName] = *lb.State
						}

						if aws.ToString(output.NextToken) == "" {
							break
						}
						config.NextToken = output.NextToken
					}
					return j, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_target_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(AutoScalingGroupWrapper)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Autoscaling
					config := autoscaling.DescribeLoadBalancerTargetGroupsInput{AutoScalingGroupName: p.AutoScalingGroupName}
					j := map[string]any{}
					for {
						output, err := svc.DescribeLoadBalancerTargetGroups(ctx, &config)
						if err != nil {
							if isAutoScalingGroupNotExistsError(err) {
								return nil, nil
							}
							return nil, err
						}
						for _, lb := range output.LoadBalancerTargetGroups {
							j[*lb.LoadBalancerTargetGroupARN] = *lb.State
						}

						if aws.ToString(output.NextToken) == "" {
							break
						}
						config.NextToken = output.NextToken
					}
					return j, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoScalingGroupARN")).Build(),
	}
}

func (x *TableAwsAutoscalingGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAutoscalingGroupLifecycleHooksGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsAutoscalingGroupScalingPoliciesGenerator{}),
	}
}
