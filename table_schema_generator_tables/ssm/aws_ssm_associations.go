package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsmAssociationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsmAssociationsGenerator{}

func (x *TableAwsSsmAssociationsGenerator) GetTableName() string {
	return "aws_ssm_associations"
}

func (x *TableAwsSsmAssociationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsmAssociationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsmAssociationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"association_id",
		},
	}
}

func (x *TableAwsSsmAssociationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ssm

			paginator := ssm.NewListAssociationsPaginator(svc, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.Associations
			}
			return nil
		},
	}
}

func (x *TableAwsSsmAssociationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ssm")
}

func (x *TableAwsSsmAssociationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("association_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssociationVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_offset").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ScheduleOffset")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("targets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Targets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssociationName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_execution_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastExecutionDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScheduleExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_maps").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetMaps")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("document_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DocumentVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("overview").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Overview")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsSsmAssociationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
