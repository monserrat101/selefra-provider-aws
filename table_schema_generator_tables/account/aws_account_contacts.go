package account

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAccountContactsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAccountContactsGenerator{}

func (x *TableAwsAccountContactsGenerator) GetTableName() string {
	return "aws_account_contacts"
}

func (x *TableAwsAccountContactsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAccountContactsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAccountContactsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsAccountContactsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Account
			var input account.GetContactInformationInput
			output, err := svc.GetContactInformation(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.ContactInformation
			return nil
		},
	}
}

func (x *TableAwsAccountContactsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsAccountContactsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("country_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CountryCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address_line2").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AddressLine2")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address_line3").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AddressLine3")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompanyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("website_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebsiteUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("city").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("City")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("postal_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PostalCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_or_region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateOrRegion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("district_or_county").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DistrictOrCounty")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address_line1").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AddressLine1")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("full_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FullName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone_number").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PhoneNumber")).Build(),
	}
}

func (x *TableAwsAccountContactsGenerator) GetSubTables() []*schema.Table {
	return nil
}
