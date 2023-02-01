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

type TableAwsEc2NetworkInterfacesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2NetworkInterfacesGenerator{}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetTableName() string {
	return "aws_ec2_network_interfaces"
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ec2
			input := ec2.DescribeNetworkInterfacesInput{}
			for {
				output, err := svc.DescribeNetworkInterfaces(ctx, &input, func(o *ec2.Options) {
					o.Region = c.Region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.NetworkInterfaces
				if aws.ToString(output.NextToken) == "" {
					break
				}
				input.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("network_interface_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkInterfaceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outpost_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutpostArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateDnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Groups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ipv6Address")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_native").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Ipv6Native")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_dest_check").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SourceDestCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Association")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mac_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MacAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requester_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequesterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_addresses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateIpAddresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.NetworkInterface)
					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"ec2",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"network-interface/" + aws.ToString(item.NetworkInterfaceId),
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
		table_schema_generator.NewColumnBuilder().ColumnName("ipv4_prefixes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Ipv4Prefixes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_addresses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Ipv6Addresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_prefixes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Ipv6Prefixes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deny_all_igw_traffic").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DenyAllIgwTraffic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interface_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InterfaceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag_set").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TagSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attachment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requester_managed").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RequesterManaged")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
	}
}

func (x *TableAwsEc2NetworkInterfacesGenerator) GetSubTables() []*schema.Table {
	return nil
}
