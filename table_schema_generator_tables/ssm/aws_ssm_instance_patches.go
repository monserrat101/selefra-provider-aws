package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsmInstancePatchesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsmInstancePatchesGenerator{}

func (x *TableAwsSsmInstancePatchesGenerator) GetTableName() string {
	return "aws_ssm_instance_patches"
}

func (x *TableAwsSsmInstancePatchesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsmInstancePatchesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsmInstancePatchesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"kb_id",
		},
	}
}

func (x *TableAwsSsmInstancePatchesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ssm
			item := task.ParentRawResult.(types.InstanceInformation)

			paginator := ssm.NewDescribeInstancePatchesPaginator(svc, &ssm.DescribeInstancePatchesInput{
				InstanceId: item.InstanceId,
			})
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.Patches
			}
			return nil
		},
	}
}

func (x *TableAwsSsmInstancePatchesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ssm")
}

func (x *TableAwsSsmInstancePatchesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kb_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KBId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cve_ids").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CVEIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ssm_instances_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ssm_instances.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("classification").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Classification")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("installed_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("InstalledTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsSsmInstancePatchesGenerator) GetSubTables() []*schema.Table {
	return nil
}
