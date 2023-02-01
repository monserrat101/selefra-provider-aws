package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIdentitystoreUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIdentitystoreUsersGenerator{}

func (x *TableAwsIdentitystoreUsersGenerator) GetTableName() string {
	return "aws_identitystore_users"
}

func (x *TableAwsIdentitystoreUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIdentitystoreUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIdentitystoreUsersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIdentitystoreUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			instance, err := getIamInstance(ctx, client)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			svc := client.(*aws_client.Client).AwsServices().Identitystore
			config := identitystore.ListUsersInput{}
			config.IdentityStoreId = instance.IdentityStoreId
			for {
				response, err := svc.ListUsers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Users

				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsIdentitystoreUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsIdentitystoreUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_store_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityStoreId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addresses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Addresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Timezone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("emails").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Emails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locale").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Locale")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nick_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NickName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profile_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProfileUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_language").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredLanguage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_ids").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExternalIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone_numbers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PhoneNumbers")).Build(),
	}
}

func (x *TableAwsIdentitystoreUsersGenerator) GetSubTables() []*schema.Table {
	return nil
}
