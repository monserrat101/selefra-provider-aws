package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightGroupsGenerator{}

func (x *TableAwsQuicksightGroupsGenerator) GetTableName() string {
	return "aws_quicksight_groups"
}

func (x *TableAwsQuicksightGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsQuicksightGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListGroupsInput{
				AwsAccountId:	aws.String(cl.AccountID),
				Namespace:	aws.String(defaultNamespace),
			}
			var ae smithy.APIError

			for {
				out, err := svc.ListGroups(ctx, &input)
				if err != nil {
					if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- out.GroupList

				if aws.ToString(out.NextToken) == "" {
					break
				}
				input.NextToken = out.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("principal_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrincipalId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsQuicksightGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsQuicksightGroupMembersGenerator{}),
	}
}
