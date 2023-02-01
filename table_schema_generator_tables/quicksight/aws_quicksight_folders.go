package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightFoldersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightFoldersGenerator{}

func (x *TableAwsQuicksightFoldersGenerator) GetTableName() string {
	return "aws_quicksight_folders"
}

func (x *TableAwsQuicksightFoldersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightFoldersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightFoldersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsQuicksightFoldersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListFoldersInput{
				AwsAccountId: aws.String(cl.AccountID),
			}
			var ae smithy.APIError

			for {
				out, err := svc.ListFolders(ctx, &input)
				if err != nil {
					if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, out.FolderSummaryList, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Quicksight
					item := result.(types.FolderSummary)

					out, err := svc.DescribeFolder(ctx, &quicksight.DescribeFolderInput{
						AwsAccountId:	aws.String(cl.AccountID),
						FolderId:	item.FolderId,
					})
					if err != nil {
						return nil, err
					}
					return out.Folder, nil

				})
				if aws.ToString(out.NextToken) == "" {
					break
				}
				input.NextToken = out.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightFoldersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightFoldersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("folder_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FolderId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("folder_path").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("FolderPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("folder_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FolderType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsQuicksightFoldersGenerator) GetSubTables() []*schema.Table {
	return nil
}
