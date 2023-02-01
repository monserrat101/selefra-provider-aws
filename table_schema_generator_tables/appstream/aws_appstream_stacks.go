package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppstreamStacksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppstreamStacksGenerator{}

func (x *TableAwsAppstreamStacksGenerator) GetTableName() string {
	return "aws_appstream_stacks"
}

func (x *TableAwsAppstreamStacksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppstreamStacksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppstreamStacksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAppstreamStacksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input appstream.DescribeStacksInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appstream
			for {
				response, err := svc.DescribeStacks(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Stacks
				if response.NextToken == nil {
					break
				}
				input.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsAppstreamStacksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appstream2")
}

func (x *TableAwsAppstreamStacksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_endpoints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AccessEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack_errors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StackErrors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ApplicationSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("feedback_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FeedbackURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redirect_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RedirectURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("embed_host_domains").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EmbedHostDomains")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_connectors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StorageConnectors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("streaming_experience_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StreamingExperienceSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserSettings")).Build(),
	}
}

func (x *TableAwsAppstreamStacksGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAppstreamStackEntitlementsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsAppstreamStackUserAssociationsGenerator{}),
	}
}
