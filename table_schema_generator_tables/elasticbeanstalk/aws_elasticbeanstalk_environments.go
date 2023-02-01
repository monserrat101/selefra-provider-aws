package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticbeanstalkEnvironmentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkEnvironmentsGenerator{}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_environments"
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"id",
		},
	}
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config elasticbeanstalk.DescribeEnvironmentsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Elasticbeanstalk
			for {
				response, err := svc.DescribeEnvironments(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Environments
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("environment_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Health")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_label").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VersionLabel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("abortable_operation_in_progress").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AbortableOperationInProgress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listeners").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(types.EnvironmentDescription)
					if p.Resources == nil || p.Resources.LoadBalancer == nil {
						return nil, nil
					}
					listeners := make(map[int32]*string, len(p.Resources.LoadBalancer.Listeners))
					for _, l := range p.Resources.LoadBalancer.Listeners {
						listeners[l.Port] = l.Protocol
					}
					return listeners, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplicationName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_updated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DateUpdated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("solution_stack_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SolutionStackName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EndpointURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cname").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CNAME")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_links").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EnvironmentLinks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TemplateName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_created").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DateCreated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HealthStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operations_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationsRole")).Build(),
	}
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkConfigurationSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkConfigurationOptionsGenerator{}),
	}
}
