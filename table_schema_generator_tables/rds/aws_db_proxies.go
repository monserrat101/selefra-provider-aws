package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDbProxiesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDbProxiesGenerator{}

func (x *TableAwsDbProxiesGenerator) GetTableName() string {
	return "aws_db_proxies"
}

func (x *TableAwsDbProxiesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDbProxiesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDbProxiesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsDbProxiesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Rds
			input := rds.DescribeDBProxiesInput{}
			for {
				output, err := svc.DescribeDBProxies(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.DBProxies
				if aws.ToString(output.Marker) == "" {
					break
				}
				input.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsDbProxiesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("rds")
}

func (x *TableAwsDbProxiesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("db_proxy_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBProxyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBProxyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_subnet_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("VpcSubnetIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_proxy_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBProxyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("debug_logging").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DebugLogging")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Auth")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("idle_client_timeout").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("IdleClientTimeout")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_tls").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RequireTLS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_security_group_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("VpcSecurityGroupIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsDbProxiesGenerator) GetSubTables() []*schema.Table {
	return nil
}
