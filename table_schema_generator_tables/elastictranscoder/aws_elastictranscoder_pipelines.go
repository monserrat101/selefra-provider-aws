package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElastictranscoderPipelinesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElastictranscoderPipelinesGenerator{}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetTableName() string {
	return "aws_elastictranscoder_pipelines"
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Elastictranscoder

			p := elastictranscoder.NewListPipelinesPaginator(svc, nil)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.Pipelines
			}

			return nil
		},
	}
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elastictranscoder")
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_kms_key_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AwsKmsKeyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ContentConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_bucket").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InputBucket")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notifications").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Notifications")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thumbnail_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ThumbnailConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_bucket").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutputBucket")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsElastictranscoderPipelinesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsElastictranscoderPipelineJobsGenerator{}),
	}
}
