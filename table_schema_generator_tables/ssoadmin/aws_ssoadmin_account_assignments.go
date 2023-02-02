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

type TableAwsSsoadminAccountAssignmentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsoadminAccountAssignmentsGenerator{}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetTableName() string {
	return "aws_ssoadmin_account_assignments"
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Ssoadmin
			permission_set_arn := task.ParentRawResult.(*types.PermissionSet).PermissionSetArn
			instance_arn := task.ParentTask.ParentRawResult.(types.InstanceMetadata).InstanceArn
			config := ssoadmin.ListAccountAssignmentsInput{
				AccountId:		&cl.AccountID,
				InstanceArn:		instance_arn,
				PermissionSetArn:	permission_set_arn,
			}

			for {
				response, err := svc.ListAccountAssignments(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.AccountAssignments
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ssoadmin_permission_sets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ssoadmin_permission_sets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission_set_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PermissionSetArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("principal_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrincipalId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("principal_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrincipalType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsSsoadminAccountAssignmentsGenerator) GetSubTables() []*schema.Table {
	return nil
}
