package redshift

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRedshiftSnapshotsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRedshiftSnapshotsGenerator{}

func (x *TableAwsRedshiftSnapshotsGenerator) GetTableName() string {
	return "aws_redshift_snapshots"
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Redshift
			cluster := task.ParentRawResult.(types.Cluster)
			params := redshift.DescribeClusterSnapshotsInput{
				ClusterExists:     aws.Bool(true),
				ClusterIdentifier: cluster.ClusterIdentifier,
				MaxRecords:        aws.Int32(100),
			}
			for {
				result, err := svc.DescribeClusterSnapshots(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.Snapshots
				if aws.ToString(result.Marker) == "" {
					break
				}
				params.Marker = result.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("redshift")
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restorable_node_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("RestorableNodeTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elapsed_time_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ElapsedTimeInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enhanced_vpc_routing").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnhancedVpcRouting")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manual_snapshot_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ManualSnapshotRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_nodes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumberOfNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_retention_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("SnapshotRetentionStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceRegion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_backup_rate_in_mega_bytes_per_second").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("CurrentBackupRateInMegaBytesPerSecond")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encrypted_with_hsm").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EncryptedWithHSM")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_full_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineFullVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("estimated_seconds_to_completion").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EstimatedSecondsToCompletion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("SnapshotCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encrypted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Encrypted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manual_snapshot_remaining_days").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ManualSnapshotRemainingDays")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterUsername")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_progress_in_mega_bytes").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("BackupProgressInMegaBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_backup_size_in_mega_bytes").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("TotalBackupSizeInMegaBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_redshift_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_redshift_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_track_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MaintenanceTrackName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("`ARN of the snapshot.`").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					snapshot := result.(types.Snapshot)
					return snapshotARN(cl, *snapshot.ClusterIdentifier, *snapshot.SnapshotIdentifier), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("accounts_with_restore_access").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AccountsWithRestoreAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actual_incremental_backup_size_in_mega_bytes").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("ActualIncrementalBackupSizeInMegaBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`Tags consisting of a name/value pair for a resource.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func snapshotARN(cl *aws_client.Client, clusterName, snapshotName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   "redshift",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("snapshot:%s/%s", clusterName, snapshotName),
	}.String()
}

func (x *TableAwsRedshiftSnapshotsGenerator) GetSubTables() []*schema.Table {
	return nil
}
