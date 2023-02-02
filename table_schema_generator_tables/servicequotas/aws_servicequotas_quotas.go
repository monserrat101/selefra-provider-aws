package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsServicequotasQuotasGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsServicequotasQuotasGenerator{}

func (x *TableAwsServicequotasQuotasGenerator) GetTableName() string {
	return "aws_servicequotas_quotas"
}

func (x *TableAwsServicequotasQuotasGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsServicequotasQuotasGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsServicequotasQuotasGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsServicequotasQuotasGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Servicequotas
			service := task.ParentRawResult.(types.ServiceInfo)
			config := servicequotas.ListServiceQuotasInput{
				ServiceCode:	service.ServiceCode,
				MaxResults:	defaultMaxResults,
			}
			quotasPaginator := servicequotas.NewListServiceQuotasPaginator(svc, &config)
			for quotasPaginator.HasMorePages() {
				output, err := quotasPaginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Quotas
			}
			return nil
		},
	}
}

var defaultMaxResults = aws.Int32(100)

func (x *TableAwsServicequotasQuotasGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsServicequotasQuotasGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("global_quota").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("GlobalQuota")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quota_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("QuotaArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("QuotaArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_metric").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UsageMetric")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_reason").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ErrorReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quota_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("QuotaCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quota_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("QuotaName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Unit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("adjustable").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Adjustable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("period").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Period")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_servicequotas_services_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_servicequotas_services.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsServicequotasQuotasGenerator) GetSubTables() []*schema.Table {
	return nil
}
