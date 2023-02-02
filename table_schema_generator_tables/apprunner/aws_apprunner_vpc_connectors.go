package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApprunnerVpcConnectorsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApprunnerVpcConnectorsGenerator{}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetTableName() string {
	return "aws_apprunner_vpc_connectors"
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apprunner.ListVpcConnectorsInput
			svc := client.(*aws_client.Client).AwsServices().Apprunner
			paginator := apprunner.NewListVpcConnectorsPaginator(svc, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.VpcConnectors
			}
			return nil
		},
	}
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apprunner")
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subnets").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Subnets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_connector_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcConnectorArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_connector_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcConnectorName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcConnectorArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DeletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_groups").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_connector_revision").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("VpcConnectorRevision")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsApprunnerVpcConnectorsGenerator) GetSubTables() []*schema.Table {
	return nil
}
