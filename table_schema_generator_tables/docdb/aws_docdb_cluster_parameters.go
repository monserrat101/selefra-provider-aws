package docdb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbClusterParametersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbClusterParametersGenerator{}

func (x *TableAwsDocdbClusterParametersGenerator) GetTableName() string {
	return "aws_docdb_cluster_parameters"
}

func (x *TableAwsDocdbClusterParametersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbClusterParametersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbClusterParametersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDocdbClusterParametersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb
			switch item := task.ParentRawResult.(type) {
			case types.DBClusterParameterGroup:
				maybeError := fetchParameterGroupParameters(ctx, svc, item, resultChannel)
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			case types.DBEngineVersion:
				maybeError := fetchEngineVersionParameters(ctx, svc, item, resultChannel)
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			}
			maybeError := fmt.Errorf("wrong parrent type to fetch cluster parameters")
			return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
		},
	}
}

func fetchEngineVersionParameters(ctx context.Context, svc aws_client.DocdbClient, item types.DBEngineVersion, res chan<- any) error {
	input := &docdb.DescribeEngineDefaultClusterParametersInput{
		DBParameterGroupFamily: item.DBParameterGroupFamily,
	}
	output, err := svc.DescribeEngineDefaultClusterParameters(ctx, input)
	if err != nil {
		return err
	}
	if output.EngineDefaults == nil || len(output.EngineDefaults.Parameters) == 0 {
		return nil
	}
	res <- output.EngineDefaults.Parameters
	return nil
}
func fetchParameterGroupParameters(ctx context.Context, svc aws_client.DocdbClient, item types.DBClusterParameterGroup, res chan<- any) error {
	input := &docdb.DescribeDBClusterParametersInput{
		DBClusterParameterGroupName: item.DBClusterParameterGroupName,
	}
	p := docdb.NewDescribeDBClusterParametersPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Parameters
	}
	return nil
}

func (x *TableAwsDocdbClusterParametersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbClusterParametersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("apply_method").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplyMethod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ParameterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_value").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ParameterValue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_values").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AllowedValues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_modifiable").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsModifiable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MinimumEngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apply_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DataType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Source")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_docdb_engine_versions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_docdb_engine_versions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsDocdbClusterParametersGenerator) GetSubTables() []*schema.Table {
	return nil
}
