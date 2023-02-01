package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsSesCustomVerificationEmailTemplatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSesCustomVerificationEmailTemplatesGenerator{}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetTableName() string {
	return "aws_ses_custom_verification_email_templates"
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sesv2

			p := sesv2.NewListCustomVerificationEmailTemplatesPaginator(svc, nil)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.CustomVerificationEmailTemplates, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sesv2
					name := result.(types.CustomVerificationEmailTemplateMetadata).TemplateName

					getOutput, err := svc.GetCustomVerificationEmailTemplate(ctx, &sesv2.GetCustomVerificationEmailTemplateInput{TemplateName: name})
					if err != nil {
						return nil, err
					}
					return getOutput, nil

				})
			}

			return nil
		},
	}
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("email")
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("success_redirection_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SuccessRedirectionURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_redirection_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureRedirectionURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("from_email_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FromEmailAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_subject").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TemplateSubject")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					return client.(*aws_client.Client).ARN("custom-verification-email-template", *result.(*sesv2.GetCustomVerificationEmailTemplateOutput).TemplateName), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_content").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TemplateContent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TemplateName")).Build(),
	}
}

func (x *TableAwsSesCustomVerificationEmailTemplatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
