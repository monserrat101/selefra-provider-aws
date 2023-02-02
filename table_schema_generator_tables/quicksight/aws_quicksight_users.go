package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightUsersGenerator{}

func (x *TableAwsQuicksightUsersGenerator) GetTableName() string {
	return "aws_quicksight_users"
}

func (x *TableAwsQuicksightUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightUsersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsQuicksightUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListUsersInput{
				AwsAccountId:	aws.String(cl.AccountID),
				Namespace:	aws.String(defaultNamespace),
			}
			var ae smithy.APIError

			for {
				out, err := svc.ListUsers(ctx, &input)
				if err != nil {
					if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- out.UserList

				if aws.ToString(out.NextToken) == "" {
					break
				}
				input.NextToken = out.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("principal_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrincipalId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_login_federation_provider_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExternalLoginFederationProviderType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_permissions_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomPermissionsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Email")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Active")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_login_federation_provider_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExternalLoginFederationProviderUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_login_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExternalLoginId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsQuicksightUsersGenerator) GetSubTables() []*schema.Table {
	return nil
}
