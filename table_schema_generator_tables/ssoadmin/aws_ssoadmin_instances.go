package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSsoadminInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSsoadminInstancesGenerator{}

func (x *TableAwsSsoadminInstancesGenerator) GetTableName() string {
	return "aws_ssoadmin_instances"
}

func (x *TableAwsSsoadminInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSsoadminInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSsoadminInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsSsoadminInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Ssoadmin
			config := ssoadmin.ListInstancesInput{}
			response, err := svc.ListInstances(ctx, &config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for _, i := range response.Instances {
				resultChannel <- i
			}
			return nil
		},
	}
}

func (x *TableAwsSsoadminInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsSsoadminInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_store_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityStoreId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceArn")).Build(),
	}
}

func (x *TableAwsSsoadminInstancesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsSsoadminPermissionSetsGenerator{}),
	}
}
