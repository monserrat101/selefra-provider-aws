package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsServicequotasServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsServicequotasServicesGenerator{}

func (x *TableAwsServicequotasServicesGenerator) GetTableName() string {
	return "aws_servicequotas_services"
}

func (x *TableAwsServicequotasServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsServicequotasServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsServicequotasServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"service_code",
			"service_name",
		},
	}
}

func (x *TableAwsServicequotasServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := servicequotas.ListServicesInput{
				MaxResults: defaultMaxResults,
			}

			svc := client.(*aws_client.Client).AwsServices().Servicequotas
			servicePaginator := servicequotas.NewListServicesPaginator(svc, &config)
			for servicePaginator.HasMorePages() {
				output, err := servicePaginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Services
			}
			return nil
		},
	}
}

func (x *TableAwsServicequotasServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("servicequotas")
}

func (x *TableAwsServicequotasServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_code").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsServicequotasServicesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsServicequotasQuotasGenerator{}),
	}
}
