package route53

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator{}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetTableName() string {
	return "aws_route53_hosted_zone_traffic_policy_instances"
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*Route53HostedZoneWrapper)
			config := route53.ListTrafficPolicyInstancesByHostedZoneInput{HostedZoneId: r.Id}
			svc := client.(*aws_client.Client).AwsServices().Route53
			for {
				response, err := svc.ListTrafficPolicyInstancesByHostedZone(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.TrafficPolicyInstances
				if aws.ToString(response.TrafficPolicyInstanceNameMarker) == "" {
					break
				}
				config.TrafficPolicyInstanceNameMarker = response.TrafficPolicyInstanceNameMarker
			}
			return nil
		},
	}
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic_policy_version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TrafficPolicyVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic_policy_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrafficPolicyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_route53_hosted_zones_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_route53_hosted_zones.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("`Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.`").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					tp := result.(types.TrafficPolicyInstance)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("route53"),
						Region:		"",
						AccountID:	"",
						Resource:	fmt.Sprintf("trafficpolicyinstance/%s", aws.ToString(tp.Id)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic_policy_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrafficPolicyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HostedZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TTL")).Build(),
	}
}

func (x *TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
