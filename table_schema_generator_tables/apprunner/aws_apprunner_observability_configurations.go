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

type TableAwsApprunnerObservabilityConfigurationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApprunnerObservabilityConfigurationsGenerator{}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetTableName() string {
	return "aws_apprunner_observability_configurations"
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apprunner.ListObservabilityConfigurationsInput
			svc := client.(*aws_client.Client).AwsServices().Apprunner
			paginator := apprunner.NewListObservabilityConfigurationsPaginator(svc, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.ObservabilityConfigurationSummaryList, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Apprunner
					service := result.(types.ObservabilityConfigurationSummary)

					describeTaskDefinitionOutput, err := svc.DescribeObservabilityConfiguration(ctx, &apprunner.DescribeObservabilityConfigurationInput{ObservabilityConfigurationArn: service.ObservabilityConfigurationArn})
					if err != nil {
						return nil, err
					}
					return describeTaskDefinitionOutput.ObservabilityConfiguration, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apprunner")
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ObservabilityConfigurationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DeletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Latest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("observability_configuration_revision").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ObservabilityConfigurationRevision")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trace_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TraceConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("observability_configuration_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ObservabilityConfigurationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("observability_configuration_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ObservabilityConfigurationName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsApprunnerObservabilityConfigurationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
