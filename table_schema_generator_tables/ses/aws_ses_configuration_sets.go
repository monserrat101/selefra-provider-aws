package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsSesConfigurationSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSesConfigurationSetsGenerator{}

func (x *TableAwsSesConfigurationSetsGenerator) GetTableName() string {
	return "aws_ses_configuration_sets"
}

func (x *TableAwsSesConfigurationSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSesConfigurationSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSesConfigurationSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSesConfigurationSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sesv2

			p := sesv2.NewListConfigurationSetsPaginator(svc, nil)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.ConfigurationSets, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sesv2
					csName := result.(string)

					getOutput, err := svc.GetConfigurationSet(ctx, &sesv2.GetConfigurationSetInput{ConfigurationSetName: &csName})
					if err != nil {
						return nil, err
					}
					return getOutput, nil

				})
			}

			return nil
		},
	}
}

func (x *TableAwsSesConfigurationSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("email")
}

func (x *TableAwsSesConfigurationSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("sending_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SendingOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suppression_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SuppressionOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vdm_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VdmOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					return client.(*aws_client.Client).ARN("configuration-set", *result.(*sesv2.GetConfigurationSetOutput).ConfigurationSetName), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_set_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ConfigurationSetName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeliveryOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reputation_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReputationOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tracking_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TrackingOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableAwsSesConfigurationSetsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsSesConfigurationSetEventDestinationsGenerator{}),
	}
}
