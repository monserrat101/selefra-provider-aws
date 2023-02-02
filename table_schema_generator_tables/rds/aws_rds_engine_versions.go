package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRdsEngineVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRdsEngineVersionsGenerator{}

func (x *TableAwsRdsEngineVersionsGenerator) GetTableName() string {
	return "aws_rds_engine_versions"
}

func (x *TableAwsRdsEngineVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRdsEngineVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRdsEngineVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"_engine_version_hash",
		},
	}
}

func (x *TableAwsRdsEngineVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Rds

			input := &rds.DescribeDBEngineVersionsInput{}

			p := rds.NewDescribeDBEngineVersionsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.DBEngineVersions
			}
			return nil
		},
	}
}

func (x *TableAwsRdsEngineVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("rds")
}

func (x *TableAwsRdsEngineVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("supported_engine_modes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SupportedEngineModes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_nchar_character_sets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SupportedNcharCharacterSets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_global_databases").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsGlobalDatabases")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_installation_files_s3_bucket_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseInstallationFilesS3BucketName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_parallel_query").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsParallelQuery")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_read_replica").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsReadReplica")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("_engine_version_hash").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					hash, err := hashstructure.Hash(result, hashstructure.FormatV2, nil)
					if err != nil {
						return nil, err
					}
					return fmt.Sprint(hash), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_character_set").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultCharacterSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_character_sets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SupportedCharacterSets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Image")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KMSKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_babelfish").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsBabelfish")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_engine_version_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBEngineVersionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_ca_certificate_identifiers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SupportedCACertificateIdentifiers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_feature_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SupportedFeatureNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_certificate_rotation_without_restart").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsCertificateRotationWithoutRestart")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag_list").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.TagsExtractor("TagList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_engine_media_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBEngineMediaType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_engine_version_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBEngineVersionDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_log_exports_to_cloudwatch_logs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsLogExportsToCloudwatchLogs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_installation_files_s3_prefix").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseInstallationFilesS3Prefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_timezones").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SupportedTimezones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_upgrade_target").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ValidUpgradeTarget")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_db_engine_version_manifest").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomDBEngineVersionManifest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_engine_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBEngineDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("major_engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MajorEngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_parameter_group_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exportable_log_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ExportableLogTypes")).Build(),
	}
}

func (x *TableAwsRdsEngineVersionsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsRdsClusterParametersGenerator{}),
	}
}
