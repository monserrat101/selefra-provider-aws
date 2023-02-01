package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsoadminPermissionSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsoadminPermissionSetsGenerator{}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetTableName() string {
	return "aws_ssoadmin_permission_sets"
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Ssoadmin
			instance_arn := task.ParentRawResult.(types.InstanceMetadata).InstanceArn
			config := ssoadmin.ListPermissionSetsInput{
				InstanceArn: instance_arn,
			}

			for {
				response, err := svc.ListPermissionSets(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.PermissionSets, func(result any) (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Ssoadmin
					permission_set_arn := result.(string)
					instance_arn := task.ParentRawResult.(types.InstanceMetadata).InstanceArn
					config := ssoadmin.DescribePermissionSetInput{
						InstanceArn:		instance_arn,
						PermissionSetArn:	&permission_set_arn,
					}

					response, err := svc.DescribePermissionSet(ctx, &config)
					if err != nil {
						return nil, err
					}
					return response.PermissionSet, nil

				})

				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("session_duration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SessionDuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("relay_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RelayState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission_set_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PermissionSetArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inline_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Ssoadmin
					permissionSetARN := result.(*types.PermissionSet).PermissionSetArn
					instanceARN := task.ParentRawResult.(types.InstanceMetadata).InstanceArn
					config := ssoadmin.GetInlinePolicyForPermissionSetInput{
						InstanceArn:		instanceARN,
						PermissionSetArn:	permissionSetARN,
					}

					response, err := svc.GetInlinePolicyForPermissionSet(ctx, &config)
					if err != nil {
						return nil, err
					}

					return response.InlinePolicy, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ssoadmin_instances_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ssoadmin_instances.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
	}
}

func (x *TableAwsSsoadminPermissionSetsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsSsoadminAccountAssignmentsGenerator{}),
	}
}
