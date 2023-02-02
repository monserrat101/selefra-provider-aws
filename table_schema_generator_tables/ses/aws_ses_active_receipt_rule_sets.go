package ses

import (
	"context"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSesActiveReceiptRuleSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSesActiveReceiptRuleSetsGenerator{}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetTableName() string {
	return "aws_ses_active_receipt_rule_sets"
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"name",
		},
	}
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ses

			set, err := svc.DescribeActiveReceiptRuleSet(ctx, nil)
			if err != nil {
				if !isRegionSupported(c.Region) && aws_client.IgnoreWithInvalidAction(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			if set.Metadata != nil && set.Metadata.Name != nil {
				resultChannel <- set
			}

			return nil
		},
	}
}

var supportedRegions = []string{"us-east-1", "us-west-2", "eu-west-1"}

func isRegionSupported(region string) bool {
	for _, r := range supportedRegions {
		if r == region {
			return true
		}
	}
	return false
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("email")
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Metadata.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Metadata.CreatedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Rules")).Build(),
	}
}

func (x *TableAwsSesActiveReceiptRuleSetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
