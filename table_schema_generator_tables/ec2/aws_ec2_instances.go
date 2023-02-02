package ec2

import (
	"context"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2InstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2InstancesGenerator{}

func (x *TableAwsEc2InstancesGenerator) GetTableName() string {
	return "aws_ec2_instances"
}

func (x *TableAwsEc2InstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2InstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2InstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2InstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ec2.DescribeInstancesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ec2
			for {
				output, err := svc.DescribeInstances(ctx, &config, func(options *ec2.Options) {
					options.Region = c.Region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, reservation := range output.Reservations {
					resultChannel <- reservation.Instances
				}

				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

var stateTransitionReasonTimeRegex = regexp.MustCompile(`\((.*)\)`)

func (x *TableAwsEc2InstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2InstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("launch_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LaunchTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MetadataOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_interfaces").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkInterfaces")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ramdisk_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RamdiskId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_transition_reason_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					instance := result.(types.Instance)
					if instance.StateTransitionReason == nil {
						return nil, nil
					}
					match := regexp.MustCompile(`\((.*)\)`).
						FindStringSubmatch(*instance.StateTransitionReason)
					if len(match) < 2 {

						return nil, nil
					}
					var layout = "2006-01-02 15:04:05 MST"
					tm, err := time.Parse(layout, match[1])
					if err != nil {
						return nil, err
					}
					return tm, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_token").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientToken")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("root_device_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RootDeviceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_reservation_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CapacityReservationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ipv6Address")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_dest_check").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SourceDestCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_transition_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateTransitionReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hibernation_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HibernationOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outpost_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutpostArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_details").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicDnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enclave_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EnclaveOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_lifecycle").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceLifecycle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sriov_net_support").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SriovNetSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_instance_profile").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IamInstanceProfile")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitoring").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Monitoring")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_operation").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UsageOperation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_device_mappings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BlockDeviceMappings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("placement").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Placement")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_name_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateDnsNameOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hypervisor").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Hypervisor")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaintenanceOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_codes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProductCodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tpm_support").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TpmSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Platform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateDnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("root_device_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RootDeviceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ami_launch_index").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AmiLaunchIndex")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architecture").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Architecture")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CpuOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elastic_inference_accelerator_associations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ElasticInferenceAcceleratorAssociations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtualization_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VirtualizationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("boot_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BootMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ebs_optimized").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EbsOptimized")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kernel_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KernelId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_instance_request_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SpotInstanceRequestId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("licenses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Licenses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elastic_gpu_associations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ElasticGpuAssociations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_reservation_specification").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CapacityReservationSpecification")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ena_support").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnaSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_operation_update_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UsageOperationUpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.Instance)
					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"ec2",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"instance/" + aws.ToString(item.InstanceId),
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
	}
}

func (x *TableAwsEc2InstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
