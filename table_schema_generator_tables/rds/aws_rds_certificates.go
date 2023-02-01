package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRdsCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRdsCertificatesGenerator{}

func (x *TableAwsRdsCertificatesGenerator) GetTableName() string {
	return "aws_rds_certificates"
}

func (x *TableAwsRdsCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRdsCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRdsCertificatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRdsCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config rds.DescribeCertificatesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Rds
			for {
				response, err := svc.DescribeCertificates(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Certificates
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRdsCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("rds")
}

func (x *TableAwsRdsCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thumbprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Thumbprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_from").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ValidFrom")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_till").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ValidTill")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_override").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("CustomerOverride")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_override_valid_till").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CustomerOverrideValidTill")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateArn")).Build(),
	}
}

func (x *TableAwsRdsCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
