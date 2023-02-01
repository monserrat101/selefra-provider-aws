package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaLayerVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaLayerVersionsGenerator{}

func (x *TableAwsLambdaLayerVersionsGenerator) GetTableName() string {
	return "aws_lambda_layer_versions"
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(types.LayersListItem)
			svc := client.(*aws_client.Client).AwsServices().Lambda
			config := lambda.ListLayerVersionsInput{
				LayerName: p.LayerName,
			}

			for {
				output, err := svc.ListLayerVersions(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.LayerVersions
				if output.NextMarker == nil {
					break
				}
				config.Marker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LayerVersionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compatible_runtimes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CompatibleRuntimes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layer_version_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LayerVersionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_info").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LicenseInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layer_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lambda_layers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lambda_layers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compatible_architectures").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CompatibleArchitectures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
	}
}

func (x *TableAwsLambdaLayerVersionsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLambdaLayerVersionPoliciesGenerator{}),
	}
}
