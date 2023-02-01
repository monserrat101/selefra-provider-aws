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

type TableAwsEc2ReservedInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2ReservedInstancesGenerator{}

func (x *TableAwsEc2ReservedInstancesGenerator) GetTableName() string {
	return "aws_ec2_reserved_instances"
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ec2.DescribeReservedInstancesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ec2

			output, err := svc.DescribeReservedInstances(ctx, &config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.ReservedInstances
			return nil
		},
	}
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("InstanceCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_tenancy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceTenancy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("offering_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OfferingClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reserved_instances_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReservedInstancesId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("duration").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Duration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("End")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_price").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("UsagePrice")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scope").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Scope")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.ReservedInstances)
					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"ec2",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"reserved-instance/" + aws.ToString(item.ReservedInstancesId),
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
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("currency_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CurrencyCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("offering_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OfferingType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fixed_price").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("FixedPrice")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurring_charges").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RecurringCharges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Start")).Build(),
	}
}

func (x *TableAwsEc2ReservedInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
