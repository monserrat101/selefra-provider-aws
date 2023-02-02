package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFrauddetectorModelVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFrauddetectorModelVersionsGenerator{}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetTableName() string {
	return "aws_frauddetector_model_versions"
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := frauddetector.NewDescribeModelVersionsPaginator(client.(*aws_client.Client).AwsServices().Frauddetector,
				&frauddetector.DescribeModelVersionsInput{ModelId: task.ParentRawResult.(types.Model).ModelId})
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ModelVersionDetails
			}
			return nil
		},
	}
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("training_result").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TrainingResult")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_events_detail").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExternalEventsDetail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingested_events_detail").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IngestedEventsDetail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_frauddetector_models_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_frauddetector_selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_data_schema").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TrainingDataSchema")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_data_source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrainingDataSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_result_v2").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TrainingResultV2")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_version_number").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelVersionNumber")).Build(),
	}
}

func (x *TableAwsFrauddetectorModelVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
