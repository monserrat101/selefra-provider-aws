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
	"golang.org/x/sync/errgroup"
)

type TableAwsEc2ImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2ImagesGenerator{}

func (x *TableAwsEc2ImagesGenerator) GetTableName() string {
	return "aws_ec2_images"
}

func (x *TableAwsEc2ImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2ImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2ImagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2ImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)

			svc := c.AwsServices().Ec2
			g, ctx := errgroup.WithContext(ctx)
			g.Go(func() error {

				pag := ec2.NewDescribeImagesPaginator(svc, &ec2.DescribeImagesInput{
					Owners: []string{"self"},
				})
				for pag.HasMorePages() {
					resp, err := pag.NextPage(ctx)
					if err != nil {
						return err
					}
					resultChannel <- resp.Images
				}
				return nil
			})

			g.Go(func() error {

				pag := ec2.NewDescribeImagesPaginator(svc, &ec2.DescribeImagesInput{
					ExecutableUsers: []string{"self"},
				})
				for pag.HasMorePages() {
					resp, err := pag.NextPage(ctx)
					if err != nil {
						return err
					}
					resultChannel <- resp.Images
				}
				return nil
			})
			maybeError := g.Wait()
			return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
		},
	}
}

func (x *TableAwsEc2ImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2ImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("ena_support").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnaSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hypervisor").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Hypervisor")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kernel_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KernelId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_device_mappings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BlockDeviceMappings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("imds_support").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImdsSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_details").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("root_device_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RootDeviceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architecture").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Architecture")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageLocation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Public")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtualization_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VirtualizationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Platform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_codes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProductCodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("root_device_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RootDeviceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sriov_net_support").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SriovNetSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_owner_alias").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageOwnerAlias")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ramdisk_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RamdiskId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_operation").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UsageOperation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deprecation_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeprecationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tpm_support").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TpmSupport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.Image)
					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"ec2",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"image/" + aws.ToString(item.ImageId),
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
		table_schema_generator.NewColumnBuilder().ColumnName("boot_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BootMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
	}
}

func (x *TableAwsEc2ImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
