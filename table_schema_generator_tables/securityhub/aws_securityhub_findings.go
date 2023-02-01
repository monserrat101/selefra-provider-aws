package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSecurityhubFindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSecurityhubFindingsGenerator{}

func (x *TableAwsSecurityhubFindingsGenerator) GetTableName() string {
	return "aws_securityhub_findings"
}

func (x *TableAwsSecurityhubFindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSecurityhubFindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSecurityhubFindingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsSecurityhubFindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Securityhub
			config := securityhub.GetFindingsInput{
				MaxResults: 100,
			}
			p := securityhub.NewGetFindingsPaginator(svc, &config)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Findings
			}
			return nil
		},
	}
}

func (x *TableAwsSecurityhubFindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("securityhub")
}

func (x *TableAwsSecurityhubFindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vulnerabilities").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Vulnerabilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finding_provider_fields").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FindingProviderFields")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("malware").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Malware")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("confidence").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Confidence")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("note").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Note")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("patch_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PatchSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("record_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecordState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_defined_fields").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserDefinedFields")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threat_intel_indicators").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ThreatIntelIndicators")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompanyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compliance").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Compliance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verification_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VerificationState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_path").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("action").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Action")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_observed_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FirstObservedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("process").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Process")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AwsAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generator_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GeneratorId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remediation").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Remediation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Types")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workflow").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Workflow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SchemaVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_fields").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProductFields")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sample").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Sample")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_observed_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastObservedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("related_findings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RelatedFindings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threats").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Threats")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workflow_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkflowState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("criticality").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Criticality")).Build(),
	}
}

func (x *TableAwsSecurityhubFindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
