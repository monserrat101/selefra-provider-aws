package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsConfigConfigRuleCompliancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsConfigConfigRuleCompliancesGenerator{}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetTableName() string {
	return "aws_config_config_rule_compliances"
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			ruleDetail := task.ParentRawResult.(types.ConfigRule)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Configservice

			input := &configservice.DescribeComplianceByConfigRuleInput{
				ConfigRuleNames: []string{aws.ToString(ruleDetail.ConfigRuleName)},
			}
			p := configservice.NewDescribeComplianceByConfigRulePaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ComplianceByConfigRules
			}
			return nil
		},
	}
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("config")
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("compliance").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Compliance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config_rule_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ConfigRuleName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_config_config_rules_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_config_config_rules.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsConfigConfigRuleCompliancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
