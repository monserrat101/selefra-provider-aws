package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppstreamStackUserAssociationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamStackUserAssociationsGenerator{}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetTableName() string {
	return "aws_appstream_stack_user_associations"
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"stack_name",
			"user_name",
			"authentication_type",
		},
	}
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input appstream.DescribeUserStackAssociationsInput
			input.StackName = task.ParentRawResult.(types.Stack).Name
			input.MaxResults = aws.Int32(25)

			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeUserStackAssociations(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.UserStackAssociations
				if response.NextToken == nil {
					break
				}
				input.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("send_email_notification").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SendEmailNotification")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_appstream_stacks_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_appstream_stacks.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authentication_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsAppstreamStackUserAssociationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
