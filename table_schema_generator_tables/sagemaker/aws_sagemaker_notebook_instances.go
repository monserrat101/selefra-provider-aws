package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSagemakerNotebookInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSagemakerNotebookInstancesGenerator{}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetTableName() string {
	return "aws_sagemaker_notebook_instances"
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sagemaker
			config := sagemaker.ListNotebookInstancesInput{}
			for {
				response, err := svc.ListNotebookInstances(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.NotebookInstances, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sagemaker
					n := result.(types.NotebookInstanceSummary)

					response, err := svc.DescribeNotebookInstance(ctx, &sagemaker.DescribeNotebookInstanceInput{
						NotebookInstanceName: n.NotebookInstanceName,
					})
					if err != nil {
						return nil, err
					}
					return &WrappedSageMakerNotebookInstance{
						DescribeNotebookInstanceOutput:	response,
						NotebookInstanceArn:		*n.NotebookInstanceArn,
						NotebookInstanceName:		*n.NotebookInstanceName,
					}, nil

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

type WrappedSageMakerNotebookInstance struct {
	*sagemaker.DescribeNotebookInstanceOutput
	NotebookInstanceArn	string
	NotebookInstanceName	string
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("api.sagemaker")
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("notebook_instance_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NotebookInstanceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NotebookInstanceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The tags associated with the notebook instance.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("describe_notebook_instance_output").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DescribeNotebookInstanceOutput")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notebook_instance_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NotebookInstanceArn")).Build(),
	}
}

func (x *TableAwsSagemakerNotebookInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
