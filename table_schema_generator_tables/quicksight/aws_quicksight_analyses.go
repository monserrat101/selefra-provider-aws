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

type TableAwsQuicksightAnalysesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightAnalysesGenerator{}

func (x *TableAwsQuicksightAnalysesGenerator) GetTableName() string {
	return "aws_quicksight_analyses"
}

func (x *TableAwsQuicksightAnalysesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightAnalysesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightAnalysesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsQuicksightAnalysesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListAnalysesInput{
				AwsAccountId: aws.String(cl.AccountID),
			}
			var ae smithy.APIError

			paginator := quicksight.NewListAnalysesPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, result.AnalysisSummaryList, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Quicksight
					item := result.(types.AnalysisSummary)

					out, err := svc.DescribeAnalysis(ctx, &quicksight.DescribeAnalysisInput{
						AwsAccountId:	aws.String(cl.AccountID),
						AnalysisId:	item.AnalysisId,
					})
					if err != nil {
						return nil, err
					}
					return out.Analysis, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightAnalysesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightAnalysesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_set_arns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("DataSetArns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("errors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Errors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("theme_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ThemeArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("analysis_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AnalysisId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sheets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Sheets")).Build(),
	}
}

func (x *TableAwsQuicksightAnalysesGenerator) GetSubTables() []*schema.Table {
	return nil
}
