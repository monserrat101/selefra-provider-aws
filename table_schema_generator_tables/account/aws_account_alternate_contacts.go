package account

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAccountAlternateContactsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAccountAlternateContactsGenerator{}

func (x *TableAwsAccountAlternateContactsGenerator) GetTableName() string {
	return "aws_account_alternate_contacts"
}

func (x *TableAwsAccountAlternateContactsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAccountAlternateContactsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAccountAlternateContactsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"alternate_contact_type",
		},
	}
}

func (x *TableAwsAccountAlternateContactsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Account
			var contactTypes types.AlternateContactType
			for _, acType := range contactTypes.Values() {
				var input account.GetAlternateContactInput
				input.AlternateContactType = acType
				output, err := svc.GetAlternateContact(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.AlternateContact
			}
			return nil
		},
	}
}

func (x *TableAwsAccountAlternateContactsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsAccountAlternateContactsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("alternate_contact_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EmailAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone_number").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PhoneNumber")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsAccountAlternateContactsGenerator) GetSubTables() []*schema.Table {
	return nil
}
