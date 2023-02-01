package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlueDatabaseTablesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueDatabaseTablesGenerator{}

func (x *TableAwsGlueDatabaseTablesGenerator) GetTableName() string {
	return "aws_glue_database_tables"
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"database_arn",
			"name",
		},
	}
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Database)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Glue
			input := glue.GetTablesInput{
				DatabaseName: r.Name,
			}
			for {
				result, err := svc.GetTables(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.TableList
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_glue_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_glue_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_registered_with_lake_formation").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsRegisteredWithLakeFormation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_access_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastAccessTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("partition_keys").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PartitionKeys")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_descriptor").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StorageDescriptor")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("view_original_text").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ViewOriginalText")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_analyzed_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastAnalyzedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Owner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Parameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TableType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VersionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("catalog_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CatalogId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Retention")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_table").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetTable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("view_expanded_text").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ViewExpandedText")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsGlueDatabaseTablesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsGlueDatabaseTableIndexesGenerator{}),
	}
}
