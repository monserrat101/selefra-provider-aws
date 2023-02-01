package ram

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRamResourceSharePermissionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRamResourceSharePermissionsGenerator{}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetTableName() string {
	return "aws_ram_resource_share_permissions"
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := &ram.ListResourceSharePermissionsInput{
				MaxResults:		aws.Int32(500),
				ResourceShareArn:	task.ParentRawResult.(types.ResourceShare).ResourceShareArn,
			}
			paginator := ram.NewListResourceSharePermissionsPaginator(client.(*aws_client.Client).AwsServices().RAM, input)
			for paginator.HasMorePages() {
				response, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Permissions
			}
			return nil
		},
	}
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_version").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DefaultVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_resource_type_default").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsResourceTypeDefault")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ram_resource_shares_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ram_resource_shares.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					permission := result.(types.ResourceSharePermissionSummary)
					version, err := strconv.ParseInt(aws.ToString(permission.Version), 10, 32)
					if err != nil {
						return nil, err
					}
					input := &ram.GetPermissionInput{
						PermissionArn:		permission.Arn,
						PermissionVersion:	aws.Int32(int32(version)),
					}
					response, err := client.(*aws_client.Client).AwsServices().RAM.GetPermission(ctx, input)
					if err != nil {
						return nil, err
					}
					return response.Permission.Permission, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
	}
}

func (x *TableAwsRamResourceSharePermissionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
