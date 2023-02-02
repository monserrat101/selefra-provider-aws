package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightDataSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightDataSetsGenerator{}

func (x *TableAwsQuicksightDataSetsGenerator) GetTableName() string {
	return "aws_quicksight_data_sets"
}

func (x *TableAwsQuicksightDataSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightDataSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightDataSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsQuicksightDataSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListDataSetsInput{
				AwsAccountId: aws.String(cl.AccountID),
			}
			var ae smithy.APIError

			paginator := quicksight.NewListDataSetsPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.DataSetSummaries
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightDataSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightDataSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("row_level_permission_data_set").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RowLevelPermissionDataSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("column_level_permission_rules_applied").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ColumnLevelPermissionRulesApplied")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_set_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DataSetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImportMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("row_level_permission_tag_configuration_applied").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RowLevelPermissionTagConfigurationApplied")).Build(),
	}
}

func (x *TableAwsQuicksightDataSetsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsQuicksightIngestionsGenerator{}),
	}
}
