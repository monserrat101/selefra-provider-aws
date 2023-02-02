package mwaa

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsMwaaEnvironmentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsMwaaEnvironmentsGenerator{}

func (x *TableAwsMwaaEnvironmentsGenerator) GetTableName() string {
	return "aws_mwaa_environments"
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := mwaa.ListEnvironmentsInput{}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Mwaa
			p := mwaa.NewListEnvironmentsPaginator(svc, &config)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Environments, func(result any) (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Mwaa
					name := result.(string)

					output, err := svc.GetEnvironment(ctx, &mwaa.GetEnvironmentInput{Name: &name})
					if err != nil {
						return nil, err
					}
					return output.Environment, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("airflow")
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("logging_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LoggingConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_workers").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MinWorkers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("webserver_access_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebserverAccessMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("weekly_maintenance_window_start").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WeeklyMaintenanceWindowStart")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastUpdate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requirements_s3_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequirementsS3Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requirements_s3_object_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequirementsS3ObjectVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedulers").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Schedulers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dag_s3_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DagS3Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("airflow_configuration_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AirflowConfigurationOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("airflow_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AirflowVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExecutionRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plugins_s3_object_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PluginsS3ObjectVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("webserver_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebserverUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_workers").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxWorkers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plugins_s3_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PluginsS3Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_bucket_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceBucketArn")).Build(),
	}
}

func (x *TableAwsMwaaEnvironmentsGenerator) GetSubTables() []*schema.Table {
	return nil
}
