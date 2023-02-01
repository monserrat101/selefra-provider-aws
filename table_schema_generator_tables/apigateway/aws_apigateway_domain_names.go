package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayDomainNamesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayDomainNamesGenerator{}

func (x *TableAwsApigatewayDomainNamesGenerator) GetTableName() string {
	return "aws_apigateway_domain_names"
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apigateway.GetDomainNamesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			for p := apigateway.NewGetDomainNamesPaginator(svc, &config); p.HasMorePages(); {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("regional_hosted_zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegionalHostedZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_upload_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CertificateUploadDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ownership_verification_certificate_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnershipVerificationCertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("regional_certificate_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegionalCertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("distribution_domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DistributionDomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name_status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainNameStatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mutual_tls_authentication").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MutualTlsAuthentication")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("distribution_hosted_zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DistributionHostedZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("regional_domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegionalDomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					domain := result.(types.DomainName)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/domainnames/%s", aws.ToString(domain.DomainName)),
					}.String(), nil
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
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainNameStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndpointConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("regional_certificate_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegionalCertificateName")).Build(),
	}
}

func (x *TableAwsApigatewayDomainNamesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsApigatewayDomainNameBasePathMappingsGenerator{}),
	}
}
