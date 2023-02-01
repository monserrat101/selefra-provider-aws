package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppstreamApplicationFleetAssociationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamApplicationFleetAssociationsGenerator{}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetTableName() string {
	return "aws_appstream_application_fleet_associations"
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"application_arn",
			"fleet_name",
		},
	}
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			parentApplication := task.ParentRawResult.(types.Application)

			var input appstream.DescribeApplicationFleetAssociationsInput
			input.ApplicationArn = parentApplication.Arn

			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeApplicationFleetAssociations(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ApplicationFleetAssociations
				if response.NextToken == nil {
					break
				}
				input.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fleet_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_appstream_applications_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_appstream_applications.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsAppstreamApplicationFleetAssociationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
