package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheClustersGenerator{}

func (x *TableAwsElasticacheClustersGenerator) GetTableName() string {
	return "aws_elasticache_clusters"
}

func (x *TableAwsElasticacheClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElasticacheClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input elasticache.DescribeCacheClustersInput
			input.ShowCacheNodeInfo = aws.Bool(true)

			paginator := elasticache.NewDescribeCacheClustersPaginator(client.(*aws_client.Client).AwsServices().Elasticache, &input)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.CacheClusters
			}
			return nil
		},
	}
}

func (x *TableAwsElasticacheClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("client_download_landing_page").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientDownloadLandingPage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_discovery").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpDiscovery")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredMaintenanceWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CacheClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheClusterStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheNodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredAvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NotificationConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheClusterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CacheSecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_group_log_delivery_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ReplicationGroupLogDeliveryEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_modified_values").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingModifiedValues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TransitEncryptionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("at_rest_encryption_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AtRestEncryptionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_token_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AuthTokenEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_minor_version_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoMinorVersionUpgrade")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_parameter_group").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CacheParameterGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_endpoint").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConfigurationEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_nodes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CacheNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_outpost_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredOutpostArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TransitEncryptionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_token_last_modified_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("AuthTokenLastModifiedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_subnet_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheSubnetGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_delivery_configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogDeliveryConfigurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_cache_nodes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumCacheNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_retention_limit").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SnapshotRetentionLimit")).Build(),
	}
}

func (x *TableAwsElasticacheClustersGenerator) GetSubTables() []*schema.Table {
	return nil
}
