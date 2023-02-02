package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppstreamFleetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamFleetsGenerator{}

func (x *TableAwsAppstreamFleetsGenerator) GetTableName() string {
	return "aws_appstream_fleets"
}

func (x *TableAwsAppstreamFleetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamFleetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamFleetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAppstreamFleetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input appstream.DescribeFleetsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeFleets(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Fleets

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAppstreamFleetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamFleetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("session_script_s3_location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SessionScriptS3Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_concurrent_sessions").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxConcurrentSessions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("idle_disconnect_timeout_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("IdleDisconnectTimeoutInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disconnect_timeout_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DisconnectTimeoutInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_join_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DomainJoinInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IamRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stream_view").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StreamView")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fleet_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FleetType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Platform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usb_device_filter_strings").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("UsbDeviceFilterStrings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_user_duration_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxUserDurationInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compute_capacity_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ComputeCapacityStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_default_internet_access").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableDefaultInternetAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fleet_errors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FleetErrors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageArn")).Build(),
	}
}

func (x *TableAwsAppstreamFleetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
