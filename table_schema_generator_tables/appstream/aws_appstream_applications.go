package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppstreamApplicationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamApplicationsGenerator{}

func (x *TableAwsAppstreamApplicationsGenerator) GetTableName() string {
	return "aws_appstream_applications"
}

func (x *TableAwsAppstreamApplicationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamApplicationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamApplicationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAppstreamApplicationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input appstream.DescribeApplicationsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeApplications(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Applications
				if response.NextToken == nil {
					break
				}
				input.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsAppstreamApplicationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamApplicationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LaunchPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("icon_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IconURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_parameters").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LaunchParameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_block_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AppBlockArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platforms").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Platforms")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("working_directory").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkingDirectory")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("icon_s3_location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IconS3Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_families").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("InstanceFamilies")).Build(),
	}
}

func (x *TableAwsAppstreamApplicationsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAppstreamApplicationFleetAssociationsGenerator{}),
	}
}
