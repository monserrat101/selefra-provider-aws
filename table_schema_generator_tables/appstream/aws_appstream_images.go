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

type TableAwsAppstreamImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamImagesGenerator{}

func (x *TableAwsAppstreamImagesGenerator) GetTableName() string {
	return "aws_appstream_images"
}

func (x *TableAwsAppstreamImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamImagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAppstreamImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := appstream.DescribeImagesInput{MaxResults: aws.Int32(25)}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeImages(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Images

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAppstreamImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("applications").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Applications")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_builder_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageBuilderName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_permissions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImagePermissions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Visibility")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_builder_supported").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ImageBuilderSupported")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Platform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_base_image_released_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("PublicBaseImageReleasedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_change_reason").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateChangeReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("appstream_agent_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AppstreamAgentVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("base_image_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BaseImageArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_errors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageErrors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsAppstreamImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
