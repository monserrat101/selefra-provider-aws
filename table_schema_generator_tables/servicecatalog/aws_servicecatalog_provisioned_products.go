package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsServicecatalogProvisionedProductsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsServicecatalogProvisionedProductsGenerator{}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetTableName() string {
	return "aws_servicecatalog_provisioned_products"
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Servicecatalog

			listInput := new(servicecatalog.SearchProvisionedProductsInput)
			for {
				output, err := svc.SearchProvisionedProducts(ctx, listInput)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- output.ProvisionedProducts

				if aws.ToString(output.NextPageToken) == "" {
					break
				}
				listInput.PageToken = output.NextPageToken
			}

			return nil
		},
	}
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("servicecatalog")
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("physical_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PhysicalId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_arn_session").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserArnSession")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_artifact_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProvisioningArtifactId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("idempotency_token").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdempotencyToken")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_record_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastRecordId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_artifact_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProvisioningArtifactName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_successful_provisioning_record_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastSuccessfulProvisioningRecordId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_provisioning_record_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastProvisioningRecordId")).Build(),
	}
}

func (x *TableAwsServicecatalogProvisionedProductsGenerator) GetSubTables() []*schema.Table {
	return nil
}
