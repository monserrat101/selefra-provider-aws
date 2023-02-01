package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsElastictranscoderPipelineJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElastictranscoderPipelineJobsGenerator{}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetTableName() string {
	return "aws_elastictranscoder_pipeline_jobs"
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Elastictranscoder

			p := elastictranscoder.NewListJobsByPipelinePaginator(
				svc,
				&elastictranscoder.ListJobsByPipelineInput{
					PipelineId: aws.String(task.ParentRow.GetStringOrDefault("id", "")),
				},
			)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.Jobs
			}

			return nil
		},
	}
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elastictranscoder_pipelines_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elastictranscoder_pipelines.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inputs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Inputs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Output")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("playlists").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Playlists")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timing").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Timing")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pipeline_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PipelineId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Input")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_key_prefix").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutputKeyPrefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outputs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Outputs")).Build(),
	}
}

func (x *TableAwsElastictranscoderPipelineJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
