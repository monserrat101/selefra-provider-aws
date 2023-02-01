package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsmPatchBaselinesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsmPatchBaselinesGenerator{}

func (x *TableAwsSsmPatchBaselinesGenerator) GetTableName() string {
	return "aws_ssm_patch_baselines"
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"baseline_id",
		},
	}
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ssm

			paginator := ssm.NewDescribePatchBaselinesPaginator(svc, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.BaselineIdentities
			}
			return nil
		},
	}
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ssm")
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("default_baseline").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DefaultBaseline")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operating_system").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperatingSystem")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("baseline_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BaselineDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("baseline_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("baseline_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BaselineName")).Build(),
	}
}

func (x *TableAwsSsmPatchBaselinesGenerator) GetSubTables() []*schema.Table {
	return nil
}
