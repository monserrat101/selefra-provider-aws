package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbEngineVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbEngineVersionsGenerator{}

func (x *TableAwsDocdbEngineVersionsGenerator) GetTableName() string {
	return "aws_docdb_engine_versions"
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := &docdb.DescribeDBEngineVersionsInput{
				Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
			}

			p := docdb.NewDescribeDBEngineVersionsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.DBEngineVersions
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("supports_log_exports_to_cloudwatch_logs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SupportsLogExportsToCloudwatchLogs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_upgrade_target").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ValidUpgradeTarget")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_engine_version_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBEngineVersionDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exportable_log_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ExportableLogTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_engine_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBEngineDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_parameter_group_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupFamily")).Build(),
	}
}

func (x *TableAwsDocdbEngineVersionsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsDocdbClusterParametersGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsDocdbOrderableDbInstanceOptionsGenerator{}),
	}
}
