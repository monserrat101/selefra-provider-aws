package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsmInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsmInstancesGenerator{}

func (x *TableAwsSsmInstancesGenerator) GetTableName() string {
	return "aws_ssm_instances"
}

func (x *TableAwsSsmInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsmInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsmInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSsmInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Ssm

			var input ssm.DescribeInstanceInformationInput
			for {
				output, err := svc.DescribeInstanceInformation(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.InstanceInformationList
				if aws.ToString(output.NextToken) == "" {
					break
				}
				input.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsSsmInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ssm")
}

func (x *TableAwsSsmInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_association_execution_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastAssociationExecutionDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_ping_date_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastPingDateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ping_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PingStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AgentVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("computer_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ComputerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_latest_version").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsLatestVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activation_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActivationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_successful_association_execution_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastSuccessfulAssociationExecutionDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registration_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("RegistrationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IamRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					instance := result.(types.InstanceInformation)
					cl := client.(*aws_client.Client)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	"ssm",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	fmt.Sprintf("managed-instance/%s", aws.ToString(instance.InstanceId)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_overview").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AssociationOverview")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssociationStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformType")).Build(),
	}
}

func (x *TableAwsSsmInstancesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsSsmInstanceComplianceItemsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsSsmInstancePatchesGenerator{}),
	}
}
