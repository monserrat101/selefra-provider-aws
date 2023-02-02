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

type TableAwsSagemakerModelsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSagemakerModelsGenerator{}

func (x *TableAwsSagemakerModelsGenerator) GetTableName() string {
	return "aws_sagemaker_models"
}

func (x *TableAwsSagemakerModelsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSagemakerModelsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSagemakerModelsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSagemakerModelsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sagemaker
			config := sagemaker.ListModelsInput{}
			for {
				response, err := svc.ListModels(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Models, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sagemaker
					n := result.(types.ModelSummary)

					response, err := svc.DescribeModel(ctx, &sagemaker.DescribeModelInput{
						ModelName: n.ModelName,
					})
					if err != nil {
						return nil, err
					}
					return &WrappedSageMakerModel{
						DescribeModelOutput:	response,
						ModelArn:		n.ModelArn,
						ModelName:		n.ModelName,
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

type WrappedSageMakerModel struct {
	*sagemaker.DescribeModelOutput
	ModelArn	*string
	ModelName	*string
}

func (x *TableAwsSagemakerModelsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("api.sagemaker")
}

func (x *TableAwsSagemakerModelsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("describe_model_output").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DescribeModelOutput")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The tags associated with the model.`").Build(),
	}
}

func (x *TableAwsSagemakerModelsGenerator) GetSubTables() []*schema.Table {
	return nil
}
