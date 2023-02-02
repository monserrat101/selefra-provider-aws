package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFrauddetectorEventTypesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFrauddetectorEventTypesGenerator{}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetTableName() string {
	return "aws_frauddetector_event_types"
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := frauddetector.NewGetEventTypesPaginator(client.(*aws_client.Client).AwsServices().Frauddetector, nil)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.EventTypes
			}
			return nil
		},
	}
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("frauddetector")
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entity_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EntityTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_ingestion").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventIngestion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_variables").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EventVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingested_event_statistics").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IngestedEventStatistics")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsFrauddetectorEventTypesGenerator) GetSubTables() []*schema.Table {
	return nil
}
