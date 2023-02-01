package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApprunnerCustomDomainsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApprunnerCustomDomainsGenerator{}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetTableName() string {
	return "aws_apprunner_custom_domains"
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := apprunner.NewDescribeCustomDomainsPaginator(client.(*aws_client.Client).AwsServices().Apprunner,
				&apprunner.DescribeCustomDomainsInput{ServiceArn: task.ParentRawResult.(*types.Service).ServiceArn})
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.CustomDomains
			}
			return nil
		},
	}
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_validation_records").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CertificateValidationRecords")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apprunner_services_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apprunner_services.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_www_subdomain").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableWWWSubdomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
	}
}

func (x *TableAwsApprunnerCustomDomainsGenerator) GetSubTables() []*schema.Table {
	return nil
}
