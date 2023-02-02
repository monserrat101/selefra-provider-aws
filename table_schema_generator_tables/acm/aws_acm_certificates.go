package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAcmCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAcmCertificatesGenerator{}

func (x *TableAwsAcmCertificatesGenerator) GetTableName() string {
	return "aws_acm_certificates"
}

func (x *TableAwsAcmCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAcmCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAcmCertificatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAcmCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Acm
			var input acm.ListCertificatesInput
			paginator := acm.NewListCertificatesPaginator(svc, &input)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.CertificateSummaryList, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Acm
					input := acm.DescribeCertificateInput{CertificateArn: result.(types.CertificateSummary).CertificateArn}
					output, err := svc.DescribeCertificate(ctx, &input)
					if err != nil {
						return nil, err
					}
					return output.Certificate, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsAcmCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("acm")
}

func (x *TableAwsAcmCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("imported_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ImportedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("not_before").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("NotBefore")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revocation_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RevocationReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject_alternative_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SubjectAlternativeNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_validation_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DomainValidationOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("in_use_by").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("InUseBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("renewal_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RenewalSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("signature_algorithm").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SignatureAlgorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_authority_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateAuthorityArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Serial")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_usages").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("KeyUsages")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("not_after").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("NotAfter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("renewal_eligibility").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RenewalEligibility")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issued_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("IssuedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_algorithm").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyAlgorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Options")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Subject")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_key_usages").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExtendedKeyUsages")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Issuer")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revoked_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("RevokedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
	}
}

func (x *TableAwsAcmCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
