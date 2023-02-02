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

type TableAwsKmsKeyPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKmsKeyPoliciesGenerator{}

func (x *TableAwsKmsKeyPoliciesGenerator) GetTableName() string {
	return "aws_kms_key_policies"
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"key_arn",
			"name",
		},
	}
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kms

			var policyName = "default"

			k := task.ParentRawResult.(*types.KeyMetadata)
			d, err := svc.GetKeyPolicy(ctx, &kms.GetKeyPolicyInput{
				KeyId:		k.Arn,
				PolicyName:	aws.String(policyName),
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- KeyPolicy{
				Name:	policyName,
				Policy:	d.Policy,
			}
			return nil
		},
	}
}

type KeyPolicy struct {
	Name	string
	Policy	*string
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kms")
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_kms_keys_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_kms_keys.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsKmsKeyPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
