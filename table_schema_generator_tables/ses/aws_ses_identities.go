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

type TableAwsSesIdentitiesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSesIdentitiesGenerator{}

func (x *TableAwsSesIdentitiesGenerator) GetTableName() string {
	return "aws_ses_identities"
}

func (x *TableAwsSesIdentitiesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSesIdentitiesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSesIdentitiesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSesIdentitiesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sesv2

			p := sesv2.NewListEmailIdentitiesPaginator(svc, nil)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.EmailIdentities, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sesv2
					ei := result.(types.IdentityInfo)

					getOutput, err := svc.GetEmailIdentity(ctx, &sesv2.GetEmailIdentityInput{EmailIdentity: ei.IdentityName})
					if err != nil {
						return nil, err
					}
					return &EmailIdentity{
						IdentityName:		ei.IdentityName,
						SendingEnabled:		ei.SendingEnabled,
						GetEmailIdentityOutput:	getOutput,
					}, nil

				})
			}

			return nil
		},
	}
}

type EmailIdentity struct {
	IdentityName	*string
	SendingEnabled	bool

	*sesv2.GetEmailIdentityOutput
}

func (x *TableAwsSesIdentitiesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("email")
}

func (x *TableAwsSesIdentitiesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("sending_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SendingEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("get_email_identity_output").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GetEmailIdentityOutput")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					return client.(*aws_client.Client).ARN("ses", "identity", *result.(*EmailIdentity).IdentityName), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityName")).Build(),
	}
}

func (x *TableAwsSesIdentitiesGenerator) GetSubTables() []*schema.Table {
	return nil
}
