package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQldbLedgerJournalS3ExportsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQldbLedgerJournalS3ExportsGenerator{}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetTableName() string {
	return "aws_qldb_ledger_journal_s3_exports"
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			ledger := task.ParentRawResult.(*qldb.DescribeLedgerOutput)
			cl := client.(*aws_client.Client)
			config := &qldb.ListJournalS3ExportsForLedgerInput{
				Name:		ledger.Name,
				MaxResults:	aws.Int32(100),
			}
			for {
				response, err := cl.AwsServices().Qldb.ListJournalS3ExportsForLedger(ctx, config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.JournalS3Exports
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("qldb")
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("s3_export_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("S3ExportConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ledger_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exclusive_end_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ExclusiveEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("export_creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ExportCreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inclusive_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("InclusiveStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("export_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExportId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ledger_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LedgerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_format").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutputFormat")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_qldb_ledgers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_qldb_ledgers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsQldbLedgerJournalS3ExportsGenerator) GetSubTables() []*schema.Table {
	return nil
}
