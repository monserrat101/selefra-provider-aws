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

type TableAwsRamResourceShareAssociationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRamResourceShareAssociationsGenerator{}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetTableName() string {
	return "aws_ram_resource_share_associations"
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"associated_entity",
			"resource_share_arn",
		},
	}
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			err := fetchRamResourceShareAssociationsByType(ctx, client, &ram.GetResourceShareAssociationsInput{
				AssociationType:	types.ResourceShareAssociationTypeResource,
				MaxResults:		aws.Int32(500),
			}, resultChannel)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			err = fetchRamResourceShareAssociationsByType(ctx, client, &ram.GetResourceShareAssociationsInput{
				AssociationType:	types.ResourceShareAssociationTypePrincipal,
				MaxResults:		aws.Int32(500),
			}, resultChannel)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			return nil
		},
	}
}

func fetchRamResourceShareAssociationsByType(ctx context.Context, client any, resourceShareInput *ram.GetResourceShareAssociationsInput, res chan<- any) error {
	paginator := ram.NewGetResourceShareAssociationsPaginator(client.(*aws_client.Client).AwsServices().RAM, resourceShareInput)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ResourceShareAssociations
	}
	return nil
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ram")
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("association_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssociationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("External")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_entity").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsRamResourceShareAssociationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
