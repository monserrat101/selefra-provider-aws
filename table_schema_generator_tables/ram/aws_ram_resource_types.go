package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRamResourceTypesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRamResourceTypesGenerator{}

func (x *TableAwsRamResourceTypesGenerator) GetTableName() string {
	return "aws_ram_resource_types"
}

func (x *TableAwsRamResourceTypesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRamResourceTypesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRamResourceTypesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsRamResourceTypesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := &ram.ListResourceTypesInput{MaxResults: aws.Int32(500)}
			paginator := ram.NewListResourceTypesPaginator(client.(*aws_client.Client).AwsServices().RAM, input)
			for paginator.HasMorePages() {
				response, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ResourceTypes
			}
			return nil
		},
	}
}

func (x *TableAwsRamResourceTypesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ram")
}

func (x *TableAwsRamResourceTypesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_region_scope").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceRegionScope")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsRamResourceTypesGenerator) GetSubTables() []*schema.Table {
	return nil
}
