package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaFunctionVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaFunctionVersionsGenerator{}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetTableName() string {
	return "aws_lambda_function_versions"
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(*lambda.GetFunctionOutput)
			if p.Configuration == nil {
				return nil
			}

			svc := client.(*aws_client.Client).AwsServices().Lambda
			config := lambda.ListVersionsByFunctionInput{
				FunctionName: p.Configuration.FunctionName,
			}

			for {
				output, err := svc.ListVersionsByFunction(ctx, &config)
				if err != nil {
					if client.(*aws_client.Client).IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Versions
				if output.NextMarker == nil {
					break
				}
				config.Marker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("master_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime_version_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RuntimeVersionConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Timeout")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architectures").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Architectures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FunctionName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Environment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_config_response").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageConfigResponse")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_status_reason_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdateStatusReasonCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MemorySize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snap_start").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SnapStart")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lambda_functions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lambda_functions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KMSKeyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("signing_job_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SigningJobArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateReasonCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tracing_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TracingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dead_letter_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeadLetterConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ephemeral_storage").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EphemeralStorage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("handler").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Handler")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdateStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Layers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastModified")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revision_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RevisionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("signing_profile_version_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SigningProfileVersionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_sha256").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CodeSha256")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CodeSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_configs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FileSystemConfigs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_status_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdateStatusReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Runtime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsLambdaFunctionVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
