package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightGroupMembersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightGroupMembersGenerator{}

func (x *TableAwsQuicksightGroupMembersGenerator) GetTableName() string {
	return "aws_quicksight_group_members"
}

func (x *TableAwsQuicksightGroupMembersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightGroupMembersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightGroupMembersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"user_arn",
			"group_arn",
		},
	}
}

func (x *TableAwsQuicksightGroupMembersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			item := task.ParentRawResult.(types.Group)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight

			input := quicksight.ListGroupMembershipsInput{
				AwsAccountId:	aws.String(cl.AccountID),
				Namespace:	aws.String(defaultNamespace),
				GroupName:	item.GroupName,
			}
			for {
				out, err := svc.ListGroupMemberships(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- out.GroupMemberList

				if aws.ToString(out.NextToken) == "" {
					break
				}
				input.NextToken = out.NextToken
			}
			return nil
		},
	}
}

var defaultNamespace = "default"

func (x *TableAwsQuicksightGroupMembersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightGroupMembersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MemberName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_quicksight_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_quicksight_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsQuicksightGroupMembersGenerator) GetSubTables() []*schema.Table {
	return nil
}
