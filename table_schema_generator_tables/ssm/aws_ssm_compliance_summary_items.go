package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsmComplianceSummaryItemsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsmComplianceSummaryItemsGenerator{}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetTableName() string {
	return "aws_ssm_compliance_summary_items"
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"compliance_type",
		},
	}
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ssm

			params := ssm.ListComplianceSummariesInput{
				MaxResults: aws.Int32(50),
			}
			for {
				output, err := svc.ListComplianceSummaries(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ComplianceSummaryItems

				if aws.ToString(output.NextToken) == "" {
					break
				}
				params.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ssm")
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("compliant_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CompliantSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("non_compliant_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NonCompliantSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compliance_type").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsSsmComplianceSummaryItemsGenerator) GetSubTables() []*schema.Table {
	return nil
}
