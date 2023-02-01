package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRdsClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRdsClustersGenerator{}

func (x *TableAwsRdsClustersGenerator) GetTableName() string {
	return "aws_rds_clusters"
}

func (x *TableAwsRdsClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRdsClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRdsClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRdsClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config rds.DescribeDBClustersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Rds
			for {
				response, err := svc.DescribeDBClusters(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.DBClusters
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRdsClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("rds")
}

func (x *TableAwsRdsClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("backup_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("BackupRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_az").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MultiAZ")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("percent_progress").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PercentProgress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_endpoints").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CustomEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("publicly_accessible").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PubliclyAccessible")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_encrypted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("StorageEncrypted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backtrack_consumed_change_records").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("BacktrackConsumedChangeRecords")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_memberships").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DomainMemberships")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_restorable_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LatestRestorableTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("performance_insights_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PerformanceInsightsEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cross_account_clone").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("CrossAccountClone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_instance_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterInstanceClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_subnet_group").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSubnetGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serverless_v2_scaling_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServerlessV2ScalingConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSystemId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("earliest_backtrack_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EarliestBacktrackTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterUsername")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("automatic_restart_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("AutomaticRestartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_write_forwarding_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GlobalWriteForwardingStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reader_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReaderEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AvailabilityZones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backtrack_window").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("BacktrackWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredMaintenanceWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag_list").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TagList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcSecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activity_stream_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActivityStreamMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("copy_tags_to_snapshot").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("CopyTagsToSnapshot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("earliest_restorable_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EarliestRestorableTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_write_forwarding_requested").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("GlobalWriteForwardingRequested")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HostedZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_modified_values").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingModifiedValues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activity_stream_kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActivityStreamKmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_option_group_memberships").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DBClusterOptionGroupMemberships")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("performance_insights_kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PerformanceInsightsKMSKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_replica_identifiers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ReadReplicaIdentifiers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activity_stream_kinesis_stream_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActivityStreamKinesisStreamName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("performance_insights_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PerformanceInsightsRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_roles").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AssociatedRoles")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_endpoint_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("HttpEndpointEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_database_authentication_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IAMDatabaseAuthenticationEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iops").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Iops")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_backup_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredBackupWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocated_storage").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AllocatedStorage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_members").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DBClusterMembers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_user_secret").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MasterUserSecret")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitoring_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MonitoringRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_minor_version_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoMinorVersionUpgrade")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_parameter_group").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterParameterGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_source_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationSourceIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activity_stream_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActivityStreamStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitoring_interval").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MonitoringInterval")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Capacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("character_set_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CharacterSetName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DbClusterResourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_configuration_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ScalingConfigurationInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("clone_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CloneGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DeletionProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled_cloudwatch_logs_exports").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EnabledCloudwatchLogsExports")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
	}
}

func (x *TableAwsRdsClustersGenerator) GetSubTables() []*schema.Table {
	return nil
}
