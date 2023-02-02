package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElastictranscoderPresetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElastictranscoderPresetsGenerator{}

func (x *TableAwsElastictranscoderPresetsGenerator) GetTableName() string {
	return "aws_elastictranscoder_presets"
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Elastictranscoder

			p := elastictranscoder.NewListPresetsPaginator(svc, nil)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.Presets
			}

			return nil
		},
	}
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elastictranscoder")
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("video").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Video")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("audio").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Audio")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Container")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thumbnails").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Thumbnails")).Build(),
	}
}

func (x *TableAwsElastictranscoderPresetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
