package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsIamRolesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamRolesGenerator{}

func (x *TableAwsIamRolesGenerator) GetTableName() string {
	return "aws_iam_roles"
}

func (x *TableAwsIamRolesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamRolesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamRolesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"id",
		},
	}
}

func (x *TableAwsIamRolesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config iam.ListRolesInput
			svc := client.(*aws_client.Client).AwsServices().IAM
			paginator := iam.NewListRolesPaginator(svc, &config)
			for paginator.HasMorePages() {
				response, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Roles, func(result any) (any, error) {
					role := result.(types.Role)
					svc := client.(*aws_client.Client).AwsServices().IAM
					roleDetails, err := svc.GetRole(ctx, &iam.GetRoleInput{
						RoleName: role.RoleName,
					})
					if err != nil {
						return nil, err
					}
					return roleDetails.Role, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsIamRolesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamRolesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assume_role_policy_document").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(*types.Role)
					if r.AssumeRolePolicyDocument == nil {
						return nil, nil
					}
					decodedDocument, err := url.QueryUnescape(*r.AssumeRolePolicyDocument)
					if err != nil {
						return nil, err
					}
					var d map[string]any
					err = json.Unmarshal([]byte(decodedDocument), &d)
					if err != nil {
						return nil, err
					}
					return d, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_last_used").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RoleLastUsed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(*types.Role)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().IAM
					input := iam.ListAttachedRolePoliciesInput{
						RoleName: r.RoleName,
					}
					policies := map[string]*string{}
					paginator := iam.NewListAttachedRolePoliciesPaginator(svc, &input)
					for paginator.HasMorePages() {
						response, err := paginator.NextPage(ctx)
						if err != nil {
							if cl.IsNotFoundError(err) {
								return nil, nil
							}
							return nil, err
						}
						for _, p := range response.AttachedPolicies {
							policies[*p.PolicyArn] = p.PolicyName
						}
					}
					return policies, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreateDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_session_duration").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxSessionDuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions_boundary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PermissionsBoundary")).Build(),
	}
}

func (x *TableAwsIamRolesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsIamRolePoliciesGenerator{}),
	}
}
