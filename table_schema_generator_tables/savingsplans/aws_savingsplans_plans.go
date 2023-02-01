package savingsplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSavingsplansPlansGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSavingsplansPlansGenerator{}

func (x *TableAwsSavingsplansPlansGenerator) GetTableName() string {
	return "aws_savingsplans_plans"
}

func (x *TableAwsSavingsplansPlansGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSavingsplansPlansGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSavingsplansPlansGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSavingsplansPlansGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().Savingsplans
			config := savingsplans.DescribeSavingsPlansInput{
				MaxResults: aws.Int32(1000),
			}
			for {
				response, err := svc.DescribeSavingsPlans(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.SavingsPlans
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsSavingsplansPlansGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsSavingsplansPlansGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("currency").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Currency")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ProductTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("savings_plan_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SavingsPlanArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("savings_plan_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SavingsPlanType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upfront_payment_amount").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UpfrontPaymentAmount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("commitment").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Commitment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurring_payment_amount").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecurringPaymentAmount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("term_duration_in_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TermDurationInSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ec2_instance_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ec2InstanceFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("End")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payment_option").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PaymentOption")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("`The Amazon Resource Name (ARN) of the Savings Plan.`").
			Extractor(column_value_extractor.StructSelector("SavingsPlanArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("offering_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OfferingId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("savings_plan_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SavingsPlanId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Start")).Build(),
	}
}

func (x *TableAwsSavingsplansPlansGenerator) GetSubTables() []*schema.Table {
	return nil
}
