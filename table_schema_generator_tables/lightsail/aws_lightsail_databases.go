package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailDatabasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailDatabasesGenerator{}

func (x *TableAwsLightsailDatabasesGenerator) GetTableName() string {
	return "aws_lightsail_databases"
}

func (x *TableAwsLightsailDatabasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailDatabasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailDatabasesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailDatabasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetRelationalDatabasesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetRelationalDatabases(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.RelationalDatabases
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailDatabasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailDatabasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_restorable_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LatestRestorableTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_modified_values").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingModifiedValues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("relational_database_blueprint_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RelationalDatabaseBlueprintId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_retention_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("BackupRetentionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hardware").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Hardware")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredMaintenanceWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecondaryAvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_backup_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredBackupWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ca_certificate_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CaCertificateIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_endpoint").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MasterEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SupportCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("publicly_accessible").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PubliclyAccessible")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterUsername")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_database_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterDatabaseName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_apply_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ParameterApplyStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_maintenance_actions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingMaintenanceActions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("relational_database_bundle_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RelationalDatabaseBundleId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
	}
}

func (x *TableAwsLightsailDatabasesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLightsailDatabaseParametersGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLightsailDatabaseEventsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLightsailDatabaseLogEventsGenerator{}),
	}
}
