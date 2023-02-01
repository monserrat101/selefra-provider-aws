package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsMqBrokersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsMqBrokersGenerator{}

func (x *TableAwsMqBrokersGenerator) GetTableName() string {
	return "aws_mq_brokers"
}

func (x *TableAwsMqBrokersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsMqBrokersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsMqBrokersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsMqBrokersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config mq.ListBrokersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Mq
			for {
				response, err := svc.ListBrokers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.BrokerSummaries, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Mq
					bs := result.(types.BrokerSummary)

					output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId})
					if err != nil {
						return nil, err
					}
					return output, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsMqBrokersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("mq")
}

func (x *TableAwsMqBrokersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("engine_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_authentication_strategy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PendingAuthenticationStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ldap_server_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LdapServerMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Logs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_minor_version_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoMinorVersionUpgrade")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Configurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_ldap_server_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingLdapServerMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_groups").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Users")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BrokerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BrokerState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HostInstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PendingEngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_security_groups").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("PendingSecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BrokerArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_window_start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaintenanceWindowStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BrokerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_instances").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BrokerInstances")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authentication_strategy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthenticationStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BrokerArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_host_instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PendingHostInstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("publicly_accessible").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PubliclyAccessible")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SubnetIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actions_required").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ActionsRequired")).Build(),
	}
}

func (x *TableAwsMqBrokersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsMqBrokerUsersGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsMqBrokerConfigurationsGenerator{}),
	}
}
