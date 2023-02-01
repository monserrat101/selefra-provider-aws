package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRamResourceSharesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRamResourceSharesGenerator{}

func (x *TableAwsRamResourceSharesGenerator) GetTableName() string {
	return "aws_ram_resource_shares"
}

func (x *TableAwsRamResourceSharesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRamResourceSharesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRamResourceSharesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRamResourceSharesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			err := fetchRamResourceSharesByType(ctx, client, types.ResourceOwnerOtherAccounts, resultChannel)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			err = fetchRamResourceSharesByType(ctx, client, types.ResourceOwnerSelf, resultChannel)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			return nil
		},
	}
}

func fetchRamResourceSharesByType(ctx context.Context, client any, shareType types.ResourceOwner, res chan<- any) error {
	input := &ram.GetResourceSharesInput{
		MaxResults:	aws.Int32(500),
		ResourceOwner:	shareType,
	}
	paginator := ram.NewGetResourceSharesPaginator(client.(*aws_client.Client).AwsServices().RAM, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ResourceShares
	}
	return nil
}

func (x *TableAwsRamResourceSharesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ram")
}

func (x *TableAwsRamResourceSharesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("feature_set").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FeatureSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owning_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwningAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_external_principals").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AllowExternalPrincipals")).Build(),
	}
}

func (x *TableAwsRamResourceSharesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsRamResourceSharePermissionsGenerator{}),
	}
}
