package timestream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsTimestreamTablesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsTimestreamTablesGenerator{}

func (x *TableAwsTimestreamTablesGenerator) GetTableName() string {
	return "aws_timestream_tables"
}

func (x *TableAwsTimestreamTablesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsTimestreamTablesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsTimestreamTablesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsTimestreamTablesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := &timestreamwrite.ListTablesInput{
				DatabaseName:	task.ParentRawResult.(types.Database).DatabaseName,
				MaxResults:	aws.Int32(20),
			}
			paginator := timestreamwrite.NewListTablesPaginator(client.(*aws_client.Client).AwsServices().Timestreamwrite, input)
			for paginator.HasMorePages() {
				response, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Tables
			}
			return nil
		},
	}
}

func (x *TableAwsTimestreamTablesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsTimestreamTablesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TableName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TableStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_timestream_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_timestream_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("magnetic_store_write_properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MagneticStoreWriteProperties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RetentionProperties")).Build(),
	}
}

func (x *TableAwsTimestreamTablesGenerator) GetSubTables() []*schema.Table {
	return nil
}
