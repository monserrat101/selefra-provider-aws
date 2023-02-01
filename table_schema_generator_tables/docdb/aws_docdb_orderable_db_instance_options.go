package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbOrderableDbInstanceOptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbOrderableDbInstanceOptionsGenerator{}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetTableName() string {
	return "aws_docdb_orderable_db_instance_options"
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			item := task.ParentRawResult.(types.DBEngineVersion)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := &docdb.DescribeOrderableDBInstanceOptionsInput{Engine: item.Engine}

			p := docdb.NewDescribeOrderableDBInstanceOptionsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.OrderableDBInstanceOptions
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("license_model").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LicenseModel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Vpc")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_docdb_engine_versions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_docdb_engine_versions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AvailabilityZones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBInstanceClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
	}
}

func (x *TableAwsDocdbOrderableDbInstanceOptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
