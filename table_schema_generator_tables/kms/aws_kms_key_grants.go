package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsKmsKeyGrantsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKmsKeyGrantsGenerator{}

func (x *TableAwsKmsKeyGrantsGenerator) GetTableName() string {
	return "aws_kms_key_grants"
}

func (x *TableAwsKmsKeyGrantsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKmsKeyGrantsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKmsKeyGrantsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"key_arn",
			"grant_id",
		},
	}
}

func (x *TableAwsKmsKeyGrantsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			k := task.ParentRawResult.(*types.KeyMetadata)
			config := kms.ListGrantsInput{
				KeyId:	k.Arn,
				Limit:	aws.Int32(100),
			}

			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kms
			p := kms.NewListGrantsPaginator(svc, &config)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Grants
			}
			return nil
		},
	}
}

func (x *TableAwsKmsKeyGrantsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kms")
}

func (x *TableAwsKmsKeyGrantsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("constraints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Constraints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retiring_principal").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RetiringPrincipal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_kms_keys_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_kms_keys.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operations").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Operations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grant_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grantee_principal").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GranteePrincipal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuing_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IssuingAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableAwsKmsKeyGrantsGenerator) GetSubTables() []*schema.Table {
	return nil
}
