package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsRamResourceShareInvitationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRamResourceShareInvitationsGenerator{}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetTableName() string {
	return "aws_ram_resource_share_invitations"
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input ram.GetResourceShareInvitationsInput = getResourceShareInvitationsInput()
			c := client.(*aws_client.Client)
			svc := c.AwsServices().RAM
			for {
				response, err := svc.GetResourceShareInvitations(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ResourceShareInvitations

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func getResourceShareInvitationsInput() ram.GetResourceShareInvitationsInput {
	return ram.GetResourceShareInvitationsInput{
		MaxResults: aws.Int32(500),
	}
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ram")
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_associations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceShareAssociations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_invitation_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareInvitationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareInvitationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("receiver_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReceiverArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("receiver_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReceiverAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_share_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceShareArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sender_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SenderAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invitation_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("InvitationTimestamp")).Build(),
	}
}

func (x *TableAwsRamResourceShareInvitationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
