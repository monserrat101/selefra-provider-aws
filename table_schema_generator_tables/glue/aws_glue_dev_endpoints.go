package glue

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsGlueDevEndpointsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueDevEndpointsGenerator{}

func (x *TableAwsGlueDevEndpointsGenerator) GetTableName() string {
	return "aws_glue_dev_endpoints"
}

func (x *TableAwsGlueDevEndpointsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueDevEndpointsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueDevEndpointsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsGlueDevEndpointsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Glue
			input := glue.GetDevEndpointsInput{}
			for {
				result, err := svc.GetDevEndpoints(ctx, &input)
				if err != nil {
					if cl.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.DevEndpoints
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueDevEndpointsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueDevEndpointsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("security_configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extra_jars_s3_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExtraJarsS3Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdateStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_group_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SecurityGroupIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("glue_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GlueVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_nodes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumberOfNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_keys").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("PublicKeys")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("worker_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkerType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zeppelin_remote_spark_interpreter_port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ZeppelinRemoteSparkInterpreterPort")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arguments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Arguments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EndpointName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("yarn_endpoint_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("YarnEndpointAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					return devEndpointARN(cl, aws.ToString(result.(types.DevEndpoint).EndpointName)), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extra_python_libs_s3_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExtraPythonLibsS3Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_workers").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumberOfWorkers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicKey")).Build(),
	}
}

func devEndpointARN(cl *aws_client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   "glue",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("devEndpoint/%s", name),
	}.String()
}

func (x *TableAwsGlueDevEndpointsGenerator) GetSubTables() []*schema.Table {
	return nil
}
