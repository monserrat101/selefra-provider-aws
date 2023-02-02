package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApprunnerAutoScalingConfigurationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApprunnerAutoScalingConfigurationsGenerator{}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetTableName() string {
	return "aws_apprunner_auto_scaling_configurations"
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apprunner.ListAutoScalingConfigurationsInput
			svc := client.(*aws_client.Client).AwsServices().Apprunner
			paginator := apprunner.NewListAutoScalingConfigurationsPaginator(svc, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.AutoScalingConfigurationSummaryList, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Apprunner
					asConfig := result.(types.AutoScalingConfigurationSummary)

					describeTaskDefinitionOutput, err := svc.DescribeAutoScalingConfiguration(ctx, &apprunner.DescribeAutoScalingConfigurationInput{AutoScalingConfigurationArn: asConfig.AutoScalingConfigurationArn})
					if err != nil {
						return nil, err
					}
					return describeTaskDefinitionOutput.AutoScalingConfiguration, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apprunner")
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("max_concurrency").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxConcurrency")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MinSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_configuration_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoScalingConfigurationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DeletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_configuration_revision").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AutoScalingConfigurationRevision")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoScalingConfigurationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Latest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_configuration_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoScalingConfigurationName")).Build(),
	}
}

func (x *TableAwsApprunnerAutoScalingConfigurationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
