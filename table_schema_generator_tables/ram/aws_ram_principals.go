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

type TableAwsRamPrincipalsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRamPrincipalsGenerator{}

func (x *TableAwsRamPrincipalsGenerator) GetTableName() string {
	return "aws_ram_principals"
}

func (x *TableAwsRamPrincipalsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRamPrincipalsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRamPrincipalsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"id",
		},
	}
}

func (x *TableAwsRamPrincipalsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			err := fetchRamPrincipalsByOwner(ctx, client, types.ResourceOwnerSelf, resultChannel)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			err = fetchRamPrincipalsByOwner(ctx, client, types.ResourceOwnerOtherAccounts, resultChannel)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			return nil
		},
	}
}

func fetchRamPrincipalsByOwner(ctx context.Context, client any, shareType types.ResourceOwner, res chan<- any) error {
	input := &ram.ListPrincipalsInput{
		MaxResults:	aws.Int32(500),
		ResourceOwner:	shareType,
	}
	paginator := ram.NewListPrincipalsPaginator(client.(*aws_client.Client).AwsServices().RAM, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Principals
	}
	return nil
}

func (x *TableAwsRamPrincipalsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ram")
}

func (x *TableAwsRamPrincipalsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("External")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsRamPrincipalsGenerator) GetSubTables() []*schema.Table {
	return nil
}
