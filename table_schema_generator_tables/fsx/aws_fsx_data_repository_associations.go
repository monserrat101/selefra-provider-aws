package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFsxDataRepositoryAssociationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFsxDataRepositoryAssociationsGenerator{}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetTableName() string {
	return "aws_fsx_data_repository_associations"
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Fsx
			input := fsx.DescribeDataRepositoryAssociationsInput{MaxResults: aws.Int32(25)}
			paginator := fsx.NewDescribeDataRepositoryAssociationsPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.Associations
			}
			return nil
		},
	}
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("fsx")
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileSystemId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssociationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("batch_import_meta_data_on_create").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("BatchImportMetaDataOnCreate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_cache_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileCachePath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nfs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NFS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("s3").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("S3")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_cache_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileCacheId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("imported_file_chunk_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ImportedFileChunkSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Lifecycle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_repository_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DataRepositoryPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_repository_subdirectories").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("DataRepositorySubdirectories")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FailureDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileSystemPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsFsxDataRepositoryAssociationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
