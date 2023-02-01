package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcrRepositoryImageScanFindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcrRepositoryImageScanFindingsGenerator{}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetTableName() string {
	return "aws_ecr_repository_image_scan_findings"
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			image := task.ParentRawResult.(types.ImageDetail)
			repo := task.ParentTask.ParentRawResult.(types.Repository)
			for _, tag := range image.ImageTags {
				config := ecr.DescribeImageScanFindingsInput{
					RepositoryName:	repo.RepositoryName,
					ImageId: &types.ImageIdentifier{
						ImageDigest:	image.ImageDigest,
						ImageTag:	aws.String(tag),
					},
					MaxResults:	aws.Int32(1000),
				}

				paginator := ecr.NewDescribeImageScanFindingsPaginator(client.(*aws_client.Client).AwsServices().Ecr, &config)
				for paginator.HasMorePages() {
					output, err := paginator.NextPage(ctx)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- ImageScanWrapper{
						ImageScanFindings:	output.ImageScanFindings,
						ImageTag:		aws.String(tag),
						ImageDigest:		image.ImageDigest,
						ImageScanStatus:	output.ImageScanStatus,
						RegistryId:		repo.RegistryId,
						RepositoryName:		repo.RepositoryName,
					}
				}
			}

			return nil
		},
	}
}

type ImageScanWrapper struct {
	ImageTag	*string
	ImageDigest	*string

	*types.ImageScanFindings

	*types.ImageScanStatus

	RegistryId	*string

	RepositoryName	*string
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ecr_repository_images_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ecr_repository_images.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_digest").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageDigest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_scan_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageScanStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_tag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageTag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_scan_findings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageScanFindings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registry_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistryId")).Build(),
	}
}

func (x *TableAwsEcrRepositoryImageScanFindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
