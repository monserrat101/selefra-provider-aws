package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsIamInstanceProfilesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamInstanceProfilesGenerator{}

func (x *TableAwsIamInstanceProfilesGenerator) GetTableName() string {
	return "aws_iam_instance_profiles"
}

func (x *TableAwsIamInstanceProfilesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamInstanceProfilesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamInstanceProfilesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"id",
		},
	}
}

func (x *TableAwsIamInstanceProfilesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := iam.ListInstanceProfilesInput{}
			svc := client.(*aws_client.Client).AwsServices().IAM
			p := iam.NewListInstanceProfilesPaginator(svc, &config)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.InstanceProfiles
			}
			return nil
		},
	}
}

func (x *TableAwsIamInstanceProfilesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamInstanceProfilesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_profile_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceProfileId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_profile_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceProfileName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceProfileId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreateDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("roles").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Roles")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsIamInstanceProfilesGenerator) GetSubTables() []*schema.Table {
	return nil
}
