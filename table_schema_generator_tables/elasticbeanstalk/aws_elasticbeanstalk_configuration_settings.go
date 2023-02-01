package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticbeanstalkConfigurationSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkConfigurationSettingsGenerator{}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_configuration_settings"
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(types.EnvironmentDescription)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Elasticbeanstalk

			configOptionsIn := elasticbeanstalk.DescribeConfigurationSettingsInput{
				ApplicationName:	p.ApplicationName,
				EnvironmentName:	p.EnvironmentName,
			}
			output, err := svc.DescribeConfigurationSettings(ctx, &configOptionsIn)
			if err != nil {

				if aws_client.IsInvalidParameterValueError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			arnStr := arn.ARN{
				Partition:	cl.Partition,
				Service:	"elasticbeanstalk",
				Region:		cl.Region,
				AccountID:	cl.AccountID,
				Resource:	fmt.Sprintf("application/%s", aws.ToString(p.ApplicationName)),
			}.String()

			for _, option := range output.ConfigurationSettings {
				resultChannel <- ConfigurationSettingsDescriptionWrapper{
					ConfigurationSettingsDescription:	option, ApplicationArn: arnStr,
				}
			}

			return nil
		},
	}
}

type ConfigurationSettingsDescriptionWrapper struct {
	types.ConfigurationSettingsDescription
	ApplicationArn	string
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elasticbeanstalk_environments_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elasticbeanstalk_environments.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
