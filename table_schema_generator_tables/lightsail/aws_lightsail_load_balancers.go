package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailLoadBalancersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailLoadBalancersGenerator{}

func (x *TableAwsLightsailLoadBalancersGenerator) GetTableName() string {
	return "aws_lightsail_load_balancers"
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetLoadBalancersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetLoadBalancers(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.LoadBalancers
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_health_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InstanceHealthSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tls_certificate_summaries").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TlsCertificateSummaries")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HealthCheckPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("https_redirection_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("HttpsRedirectionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SupportCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConfigurationOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("InstancePort")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpAddressType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocol").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Protocol")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ports").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("PublicPorts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tls_policy_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TlsPolicyName")).Build(),
	}
}

func (x *TableAwsLightsailLoadBalancersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLightsailLoadBalancerTlsCertificatesGenerator{}),
	}
}
