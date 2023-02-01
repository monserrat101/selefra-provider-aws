package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"golang.org/x/sync/errgroup"
)

type TableAwsLightsailDistributionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailDistributionsGenerator{}

func (x *TableAwsLightsailDistributionsGenerator) GetTableName() string {
	return "aws_lightsail_distributions"
}

func (x *TableAwsLightsailDistributionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailDistributionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailDistributionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLightsailDistributionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetDistributionsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetDistributions(ctx, &input, func(options *lightsail.Options) {

					options.Region = "us-east-1"
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				errs, ctx := errgroup.WithContext(ctx)
				errs.SetLimit(MaxGoroutines)
				for _, d := range response.Distributions {
					func(d types.LightsailDistribution) {
						errs.Go(func() error {
							return fetchCacheReset(ctx, resultChannel, c, d)
						})
					}(d)
				}
				err = errs.Wait()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

type DistributionWrapper struct {
	*types.LightsailDistribution
	LatestCacheReset	*lightsail.GetDistributionLatestCacheResetOutput
}

func fetchCacheReset(ctx context.Context, res chan<- any, c *aws_client.Client, d types.LightsailDistribution) error {
	svc := c.AwsServices().Lightsail
	resetInput := lightsail.GetDistributionLatestCacheResetInput{
		DistributionName: d.Name,
	}
	resetResp, err := svc.GetDistributionLatestCacheReset(ctx, &resetInput, func(options *lightsail.Options) {

		options.Region = "us-east-1"
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}
	res <- DistributionWrapper{LightsailDistribution: &d, LatestCacheReset: resetResp}
	return nil
}

func (x *TableAwsLightsailDistributionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailDistributionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lightsail_distribution").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LightsailDistribution")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_cache_reset").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LatestCacheReset")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsLightsailDistributionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
