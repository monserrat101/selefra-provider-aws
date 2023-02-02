package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2SubnetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2SubnetsGenerator{}

func (x *TableAwsEc2SubnetsGenerator) GetTableName() string {
	return "aws_ec2_subnets"
}

func (x *TableAwsEc2SubnetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2SubnetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2SubnetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2SubnetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ec2.DescribeSubnetsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ec2
			for {
				output, err := svc.DescribeSubnets(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Subnets
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2SubnetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2SubnetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assign_ipv6_address_on_creation").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AssignIpv6AddressOnCreation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_ip_address_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AvailableIpAddressCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_owned_ipv4_pool").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomerOwnedIpv4Pool")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_dns64").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableDns64")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_block").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CidrBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("map_customer_owned_ip_on_launch").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MapCustomerOwnedIpOnLaunch")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outpost_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutpostArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_for_az").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DefaultForAz")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_lni_at_device_index").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EnableLniAtDeviceIndex")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_block_association_set").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Ipv6CidrBlockAssociationSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_native").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Ipv6Native")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("map_public_ip_on_launch").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MapPublicIpOnLaunch")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_name_options_on_launch").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateDnsNameOptionsOnLaunch")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
	}
}

func (x *TableAwsEc2SubnetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
