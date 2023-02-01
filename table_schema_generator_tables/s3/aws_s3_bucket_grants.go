package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsS3BucketGrantsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsS3BucketGrantsGenerator{}

func (x *TableAwsS3BucketGrantsGenerator) GetTableName() string {
	return "aws_s3_bucket_grants"
}

func (x *TableAwsS3BucketGrantsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsS3BucketGrantsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsS3BucketGrantsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"bucket_arn",
			"grantee_type",
			"grantee_id",
		},
	}
}

func (x *TableAwsS3BucketGrantsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*WrappedBucket)
			svc := client.(*aws_client.Client).AwsServices().S3
			region := task.ParentRow.GetOrDefault("region", "")
			if region == nil {
				return nil
			}
			aclOutput, err := svc.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: r.Name}, func(options *s3.Options) {
				options.Region = region.(string)
			})
			if err != nil {
				if aws_client.IsAWSError(err, "NoSuchBucket") {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- aclOutput.Grants
			return nil
		},
	}
}

func (x *TableAwsS3BucketGrantsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsS3BucketGrantsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("grantee_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Grantee.Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grantee_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					grantee := result.(types.Grant).Grantee
					switch grantee.Type {
					case types.TypeCanonicalUser:
						return *grantee.ID, nil
					case types.TypeAmazonCustomerByEmail:
						return *grantee.EmailAddress, nil
					case types.TypeGroup:
						return *grantee.URI, nil
					default:
						return nil, fmt.Errorf("unsupported grantee type %q", grantee.Type)
					}
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
		table_schema_generator.NewColumnBuilder().ColumnName("aws_s3_buckets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grantee").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Grantee")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Permission")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsS3BucketGrantsGenerator) GetSubTables() []*schema.Table {
	return nil
}
