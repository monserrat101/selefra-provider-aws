package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheReplicationGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheReplicationGroupsGenerator{}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetTableName() string {
	return "aws_elasticache_replication_groups"
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := elasticache.NewDescribeReplicationGroupsPaginator(client.(*aws_client.Client).AwsServices().Elasticache, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.ReplicationGroups
			}
			return nil
		},
	}
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cache_node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheNodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TransitEncryptionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_delivery_configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogDeliveryConfigurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_az").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MultiAZ")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_token_last_modified_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("AuthTokenLastModifiedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_minor_version_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoMinorVersionUpgrade")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("automatic_failover").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutomaticFailover")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ClusterEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_endpoint").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConfigurationEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshotting_cluster_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshottingClusterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("at_rest_encryption_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AtRestEncryptionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_tiering").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DataTiering")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_replication_group_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GlobalReplicationGroupInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_clusters").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("MemberClusters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_modified_values").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingModifiedValues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_discovery").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpDiscovery")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TransitEncryptionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_token_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AuthTokenEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_clusters_outpost_arns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("MemberClustersOutpostArns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_retention_limit").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SnapshotRetentionLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_group_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("UserGroupIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodeGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_group_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ReplicationGroupCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationGroupId")).Build(),
	}
}

func (x *TableAwsElasticacheReplicationGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
