package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcrRepositoriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcrRepositoriesGenerator{}

func (x *TableAwsEcrRepositoriesGenerator) GetTableName() string {
	return "aws_ecr_repositories"
}

func (x *TableAwsEcrRepositoriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcrRepositoriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcrRepositoriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEcrRepositoriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			maxResults := int32(1000)
			config := ecr.DescribeRepositoriesInput{
				MaxResults: &maxResults,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ecr
			for {
				output, err := svc.DescribeRepositories(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Repositories
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEcrRepositoriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("api.ecr")
}

func (x *TableAwsEcrRepositoriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_tag_mutability").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageTagMutability")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registry_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistryId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_text").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Ecr
					repo := result.(types.Repository)

					input := ecr.GetRepositoryPolicyInput{
						RepositoryName:	repo.RepositoryName,
						RegistryId:	repo.RegistryId,
					}
					output, err := svc.GetRepositoryPolicy(ctx, &input)
					if err != nil {
						return nil, err
					}
					return output.PolicyText, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_scanning_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageScanningConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsEcrRepositoriesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsEcrRepositoryImagesGenerator{}),
	}
}
