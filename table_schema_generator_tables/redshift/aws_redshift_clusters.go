package redshift

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRedshiftClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRedshiftClustersGenerator{}

func (x *TableAwsRedshiftClustersGenerator) GetTableName() string {
	return "aws_redshift_clusters"
}

func (x *TableAwsRedshiftClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRedshiftClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRedshiftClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRedshiftClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config redshift.DescribeClustersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Redshift
			for {
				response, err := svc.DescribeClusters(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Clusters
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRedshiftClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("redshift")
}

func (x *TableAwsRedshiftClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_parameter_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ClusterParameterGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enhanced_vpc_routing").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnhancedVpcRouting")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_roles").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IamRoles")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredMaintenanceWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encrypted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Encrypted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_nodes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumberOfNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("`The Amazon Resource Name (ARN) for the resource.`").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{fmt.Sprintf("cluster:%s", *result.(types.Cluster).ClusterIdentifier)}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"redshift",
					Region:		cl.Region,
					AccountID:	cl.AccountID,
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expected_next_snapshot_schedule_time_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExpectedNextSnapshotScheduleTimeStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deferred_maintenance_windows").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeferredMaintenanceWindows")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elastic_resize_number_of_node_options").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ElasticResizeNumberOfNodeOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_schedule_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotScheduleIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_revision_number").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterRevisionNumber")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_iam_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultIamRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_actions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("PendingActions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_version_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AllowVersionUpgrade")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_availability_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterAvailabilityStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_snapshot_copy_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ClusterSnapshotCopyStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elastic_ip_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ElasticIpStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expected_next_snapshot_schedule_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ExpectedNextSnapshotScheduleTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("automated_snapshot_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AutomatedSnapshotRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterUsername")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aqua_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AquaConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_storage_capacity_in_mega_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TotalStorageCapacityInMegaBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcSecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_modified_values").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingModifiedValues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_public_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterPublicKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_transfer_progress").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DataTransferProgress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hsm_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HsmStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_subnet_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterSubnetGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_track_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MaintenanceTrackName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_nodes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ClusterNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone_relocation_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZoneRelocationStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_namespace_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterNamespaceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ClusterSecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restore_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RestoreStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manual_snapshot_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ManualSnapshotRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_maintenance_window_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("NextMaintenanceWindowStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reserved_node_exchange_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReservedNodeExchangeStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_schedule_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotScheduleState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("publicly_accessible").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PubliclyAccessible")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modify_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModifyStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resize_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResizeInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_status").ColumnType(schema.ColumnTypeJSON).Description("`Describes the status of logging for a cluster.`").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(types.Cluster)

					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Redshift
					cfg := redshift.DescribeLoggingStatusInput{
						ClusterIdentifier: r.ClusterIdentifier,
					}
					response, err := svc.DescribeLoggingStatus(ctx, &cfg)
					if err != nil {
						return nil, err
					}

					return response, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
	}
}

func (x *TableAwsRedshiftClustersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsRedshiftSnapshotsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsRedshiftClusterParameterGroupsGenerator{}),
	}
}
