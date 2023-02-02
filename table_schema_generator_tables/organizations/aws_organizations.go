package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsOrganizationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsOrganizationsGenerator{}

func (x *TableAwsOrganizationsGenerator) GetTableName() string {
	return "aws_organizations"
}

func (x *TableAwsOrganizationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsOrganizationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsOrganizationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsOrganizationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			o, err := c.AwsServices().Organizations.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- o.Organization
			return nil
		},
	}
}

func (x *TableAwsOrganizationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsOrganizationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("feature_set").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FeatureSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_account_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterAccountArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_policy_types").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AvailablePolicyTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_account_email").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterAccountEmail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsOrganizationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
