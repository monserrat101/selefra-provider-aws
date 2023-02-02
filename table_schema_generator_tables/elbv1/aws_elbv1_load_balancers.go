package elbv1

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsElbv1LoadBalancersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv1LoadBalancersGenerator{}

func (x *TableAwsElbv1LoadBalancersGenerator) GetTableName() string {
	return "aws_elbv1_load_balancers"
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Elasticloadbalancing
			processLoadBalancers := func(loadBalancers []types.LoadBalancerDescription) error {
				tagsCfg := &elbv1.DescribeTagsInput{LoadBalancerNames: make([]string, 0, len(loadBalancers))}
				for _, lb := range loadBalancers {
					tagsCfg.LoadBalancerNames = append(tagsCfg.LoadBalancerNames, *lb.LoadBalancerName)
				}
				tagsResponse, err := svc.DescribeTags(ctx, tagsCfg)
				if err != nil {
					return err
				}
				for _, lb := range loadBalancers {
					loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName})
					if err != nil {
						if c.IsNotFoundError(err) {
							continue
						}
						return err
					}

					wrapper := ELBv1LoadBalancerWrapper{
						LoadBalancerDescription:	lb,
						Tags:				aws_client.TagsToMap(getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)),
						Attributes:			loadBalancerAttributes.LoadBalancerAttributes,
					}

					resultChannel <- wrapper
				}
				return nil
			}

			var config elbv1.DescribeLoadBalancersInput
			for {
				response, err := svc.DescribeLoadBalancers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for i := 0; i < len(response.LoadBalancerDescriptions); i += 20 {
					end := i + 20

					if end > len(response.LoadBalancerDescriptions) {
						end = len(response.LoadBalancerDescriptions)
					}
					loadBalancers := response.LoadBalancerDescriptions[i:end]
					if err := processLoadBalancers(loadBalancers); err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
				}

				if aws.ToString(response.NextMarker) == "" {
					break
				}
				config.Marker = response.NextMarker
			}

			return nil
		},
	}
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_name").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LoadBalancerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_description").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LoadBalancerDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{"loadbalancer", *result.(ELBv1LoadBalancerWrapper).LoadBalancerName}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"elasticloadbalancing",
					Region:		cl.Region,
					AccountID:	cl.AccountID,
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsElbv1LoadBalancerPoliciesGenerator{}),
	}
}
