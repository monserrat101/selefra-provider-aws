package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCodebuildProjectsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCodebuildProjectsGenerator{}

func (x *TableAwsCodebuildProjectsGenerator) GetTableName() string {
	return "aws_codebuild_projects"
}

func (x *TableAwsCodebuildProjectsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCodebuildProjectsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCodebuildProjectsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsCodebuildProjectsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Codebuild
			config := codebuild.ListProjectsInput{}
			for {
				response, err := svc.ListProjects(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if len(response.Projects) == 0 {
					break
				}
				projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: response.Projects})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- projectsOutput.Projects
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCodebuildProjectsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("codebuild")
}

func (x *TableAwsCodebuildProjectsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_locations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FileSystemLocations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("badge").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Badge")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_source_versions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecondarySourceVersions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queued_timeout_in_minutes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("QueuedTimeoutInMinutes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_project_alias").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicProjectAlias")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Cache")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logs_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogsConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_sources").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecondarySources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_visibility").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProjectVisibility")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Source")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("artifacts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Artifacts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Environment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("webhook").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Webhook")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("concurrent_build_limit").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ConcurrentBuildLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_artifacts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecondaryArtifacts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout_in_minutes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TimeoutInMinutes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_batch_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BuildBatchConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_access_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceAccessRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModified")).Build(),
	}
}

func (x *TableAwsCodebuildProjectsGenerator) GetSubTables() []*schema.Table {
	return nil
}
