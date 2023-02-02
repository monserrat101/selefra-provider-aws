package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIdentitystoreGroupMembershipsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIdentitystoreGroupMembershipsGenerator{}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetTableName() string {
	return "aws_identitystore_group_memberships"
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Identitystore
			group := task.ParentRawResult.(types.Group)
			config := identitystore.ListGroupMembershipsInput{
				GroupId:		group.GroupId,
				IdentityStoreId:	group.IdentityStoreId,
			}
			for {
				response, err := svc.ListGroupMemberships(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.GroupMemberships

				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("identity_store_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityStoreId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("membership_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MembershipId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_identitystore_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_identitystore_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsIdentitystoreGroupMembershipsGenerator) GetSubTables() []*schema.Table {
	return nil
}
