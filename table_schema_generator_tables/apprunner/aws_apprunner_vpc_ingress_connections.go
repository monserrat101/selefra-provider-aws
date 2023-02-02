package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApprunnerVpcIngressConnectionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApprunnerVpcIngressConnectionsGenerator{}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetTableName() string {
	return "aws_apprunner_vpc_ingress_connections"
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apprunner.ListVpcIngressConnectionsInput
			svc := client.(*aws_client.Client).AwsServices().Apprunner
			paginator := apprunner.NewListVpcIngressConnectionsPaginator(svc, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.VpcIngressConnectionSummaryList, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Apprunner
					asConfig := result.(types.VpcIngressConnectionSummary)

					describeTaskDefinitionOutput, err := svc.DescribeVpcIngressConnection(ctx, &apprunner.DescribeVpcIngressConnectionInput{VpcIngressConnectionArn: asConfig.VpcIngressConnectionArn})
					if err != nil {
						return nil, err
					}
					return describeTaskDefinitionOutput.VpcIngressConnection, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apprunner")
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("ingress_vpc_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IngressVpcConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_ingress_connection_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcIngressConnectionName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DeletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcIngressConnectionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_ingress_connection_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcIngressConnectionArn")).Build(),
	}
}

func (x *TableAwsApprunnerVpcIngressConnectionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
