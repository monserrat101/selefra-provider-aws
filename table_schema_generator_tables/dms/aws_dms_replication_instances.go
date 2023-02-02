package dms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDmsReplicationInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDmsReplicationInstancesGenerator{}

func (x *TableAwsDmsReplicationInstancesGenerator) GetTableName() string {
	return "aws_dms_replication_instances"
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Databasemigrationservice

			var describeReplicationInstancesInput *databasemigrationservice.DescribeReplicationInstancesInput
			describeReplicationInstancesOutput, err := svc.DescribeReplicationInstances(ctx, describeReplicationInstancesInput)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if len(describeReplicationInstancesOutput.ReplicationInstances) == 0 {
				return nil
			}

			listTagsForResourceInput := databasemigrationservice.ListTagsForResourceInput{}
			for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
				listTagsForResourceInput.ResourceArnList = append(listTagsForResourceInput.ResourceArnList, *replicationInstance.ReplicationInstanceArn)
			}
			var listTagsForResourceOutput *databasemigrationservice.ListTagsForResourceOutput
			listTagsForResourceOutput, err = svc.ListTagsForResource(ctx, &listTagsForResourceInput)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			replicationInstanceTags := make(map[string]map[string]any)
			for _, tag := range listTagsForResourceOutput.TagList {
				if replicationInstanceTags[*tag.ResourceArn] == nil {
					replicationInstanceTags[*tag.ResourceArn] = make(map[string]any)
				}
				replicationInstanceTags[*tag.ResourceArn][*tag.Key] = *tag.Value
			}

			for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
				wrapper := ReplicationInstanceWrapper{
					ReplicationInstance:	replicationInstance,
					Tags:			replicationInstanceTags[*replicationInstance.ReplicationInstanceArn],
				}
				resultChannel <- wrapper
			}
			return nil
		},
	}
}

type ReplicationInstanceWrapper struct {
	types.ReplicationInstance
	Tags	map[string]any
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("dms")
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReplicationInstance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationInstanceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
