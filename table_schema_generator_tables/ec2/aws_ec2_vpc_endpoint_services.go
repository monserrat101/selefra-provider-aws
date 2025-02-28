package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2VpcEndpointServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2VpcEndpointServicesGenerator{}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetTableName() string {
	return "aws_ec2_vpc_endpoint_services"
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ec2.DescribeVpcEndpointServicesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ec2
			for {
				output, err := svc.DescribeVpcEndpointServices(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ServiceDetails
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acceptance_required").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AcceptanceRequired")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AvailabilityZones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("base_endpoint_dns_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("BaseEndpointDnsNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_ip_address_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SupportedIpAddressTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_endpoint_policy_supported").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("VpcEndpointPolicySupported")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manages_vpc_endpoints").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ManagesVpcEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Owner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_type").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateDnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_name_verification_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateDnsNameVerificationState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.ServiceDetail)
					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"ec2",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"vpc-endpoint-service/" + aws.ToString(item.ServiceId),
					}
					return a.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payer_responsibility").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PayerResponsibility")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_names").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateDnsNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceId")).Build(),
	}
}

func (x *TableAwsEc2VpcEndpointServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
