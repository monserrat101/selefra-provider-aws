package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsmDocumentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsmDocumentsGenerator{}

func (x *TableAwsSsmDocumentsGenerator) GetTableName() string {
	return "aws_ssm_documents"
}

func (x *TableAwsSsmDocumentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsmDocumentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsmDocumentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSsmDocumentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ssm

			params := ssm.ListDocumentsInput{
				Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}},
			}
			for {
				output, err := svc.ListDocuments(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.DocumentIdentifiers, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Ssm
					d := result.(types.DocumentIdentifier)

					dd, err := svc.DescribeDocument(ctx, &ssm.DescribeDocumentInput{Name: d.Name})
					if err != nil {
						return nil, err
					}
					return dd.Document, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				params.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsSsmDocumentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ssm")
}

func (x *TableAwsSsmDocumentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("latest_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LatestVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_review_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PendingReviewVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SchemaVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_information").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusInformation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("document_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DocumentVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Owner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TargetType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approved_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApprovedVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("document_format").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DocumentFormat")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					d := result.(*types.DocumentDescription)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Ssm

					input := ssm.DescribeDocumentPermissionInput{
						Name:		d.Name,
						PermissionType:	types.DocumentPermissionTypeShare,
					}
					var permissions []*ssm.DescribeDocumentPermissionOutput
					for {
						output, err := svc.DescribeDocumentPermission(ctx, &input)
						if err != nil {
							return nil, err
						}
						permissions = append(permissions, output)
						if aws.ToString(output.NextToken) == "" {
							break
						}
						input.NextToken = output.NextToken
					}
					return permissions, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments_information").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AttachmentsInformation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("author").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Author")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VersionName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Category")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("document_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DocumentType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Parameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requires").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Requires")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("review_information").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReviewInformation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("PlatformTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("review_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReviewStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sha1").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Sha1")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category_enum").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CategoryEnum")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hash").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Hash")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hash_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HashType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					d := result.(*types.DocumentDescription)
					cl := client.(*aws_client.Client)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	"ssm",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	fmt.Sprintf("document/%s", aws.ToString(d.Name)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
	}
}

func (x *TableAwsSsmDocumentsGenerator) GetSubTables() []*schema.Table {
	return nil
}
