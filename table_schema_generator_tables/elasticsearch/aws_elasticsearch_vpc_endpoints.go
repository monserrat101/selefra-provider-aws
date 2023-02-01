package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticsearchVpcEndpointsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticsearchVpcEndpointsGenerator{}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetTableName() string {
	return "aws_elasticsearch_vpc_endpoints"
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

			listInput := new(elasticsearchservice.ListVpcEndpointsInput)
			var vpcEndpointIDs []string

			for {
				out, err := svc.ListVpcEndpoints(ctx, listInput)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for _, summary := range out.VpcEndpointSummaryList {
					vpcEndpointIDs = append(vpcEndpointIDs, *summary.VpcEndpointId)
				}

				if out.NextToken == nil {
					break
				}

				listInput.NextToken = out.NextToken
			}

			var maxLen = 100
			for len(vpcEndpointIDs) > 0 {
				var part []string
				if len(vpcEndpointIDs) > maxLen {
					part, vpcEndpointIDs = vpcEndpointIDs[:maxLen], vpcEndpointIDs[maxLen:]
				} else {
					part, vpcEndpointIDs = vpcEndpointIDs, nil
				}

				out, err := svc.DescribeVpcEndpoints(ctx,
					&elasticsearchservice.DescribeVpcEndpointsInput{VpcEndpointIds: part},
				)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- out.VpcEndpoints
			}

			return nil
		},
	}
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("es")
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_endpoint_owner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcEndpointOwner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcEndpointId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_endpoint_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcEndpointId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsElasticsearchVpcEndpointsGenerator) GetSubTables() []*schema.Table {
	return nil
}
