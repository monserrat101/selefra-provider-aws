package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFrauddetectorRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFrauddetectorRulesGenerator{}

func (x *TableAwsFrauddetectorRulesGenerator) GetTableName() string {
	return "aws_frauddetector_rules"
}

func (x *TableAwsFrauddetectorRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFrauddetectorRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFrauddetectorRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFrauddetectorRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := frauddetector.NewGetRulesPaginator(client.(*aws_client.Client).AwsServices().Frauddetector,
				&frauddetector.GetRulesInput{DetectorId: task.ParentRawResult.(types.Detector).DetectorId})
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.RuleDetails
			}
			return nil
		},
	}
}

func (x *TableAwsFrauddetectorRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsFrauddetectorRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("language").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Language")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outcomes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Outcomes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Expression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DetectorId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RuleId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RuleVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_frauddetector_detectors_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_frauddetector_detectors.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsFrauddetectorRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
