package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticsearchDomainsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticsearchDomainsGenerator{}

func (x *TableAwsElasticsearchDomainsGenerator) GetTableName() string {
	return "aws_elasticsearch_domains"
}

func (x *TableAwsElasticsearchDomainsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticsearchDomainsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticsearchDomainsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElasticsearchDomainsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

			out, err := svc.ListDomainNames(ctx, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			aws_client.SendResults(resultChannel, out.DomainNames, func(result any) (any, error) {
				svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

				out, err := svc.DescribeElasticsearchDomain(ctx,
					&elasticsearchservice.DescribeElasticsearchDomainInput{
						DomainName: result.(types.DomainInfo).DomainName,
					},
				)
				if err != nil {
					return nil, err
				}
				return out.DomainStatus, nil

			})
			return nil
		},
	}
}

func (x *TableAwsElasticsearchDomainsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("es")
}

func (x *TableAwsElasticsearchDomainsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("change_progress_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ChangeProgressDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VPCOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_policies").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AccessPolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elasticsearch_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ElasticsearchVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_software_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceSoftwareOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_security_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdvancedSecurityOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_at_rest_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionAtRestOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorized_principals").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					svc := client.(*aws_client.Client).AwsServices().Elasticsearchservice

					input := &elasticsearchservice.ListVpcEndpointAccessInput{
						DomainName: result.(*types.ElasticsearchDomainStatus).DomainName,
					}

					var principals []types.AuthorizedPrincipal
					for {
						out, err := svc.ListVpcEndpointAccess(ctx, input)
						if err != nil {
							return nil, err
						}

						principals = append(principals, out.AuthorizedPrincipalList...)

						if out.NextToken == nil {
							break
						}

						input.NextToken = out.NextToken
					}

					return principals, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ebs_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EBSOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processing").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Processing")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_publishing_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogPublishingOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elasticsearch_cluster_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ElasticsearchClusterConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdvancedOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_tune_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AutoTuneOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cognito_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CognitoOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Deleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Endpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SnapshotOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_endpoint_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DomainEndpointOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_to_node_encryption_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodeToNodeEncryptionOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upgrade_processing").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("UpgradeProcessing")).Build(),
	}
}

func (x *TableAwsElasticsearchDomainsGenerator) GetSubTables() []*schema.Table {
	return nil
}
