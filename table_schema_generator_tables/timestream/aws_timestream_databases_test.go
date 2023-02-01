package timestream







import (


	"testing"







	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildTimestreamDatabasesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockTimestreamwriteClient(ctrl)
	database := types.Database{}
	err := faker.FakeObject(&database)


	if err != nil {


		t.Fatal(err)
	}







	m.EXPECT().ListDatabases(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&timestreamwrite.ListDatabasesOutput{Databases: []types.Database{database}}, nil)



	table := types.Table{}




	err = faker.FakeObject(&table)
	if err != nil {
		t.Fatal(err)
	}





	m.EXPECT().ListTables(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&timestreamwrite.ListTablesOutput{Tables: []types.Table{table}}, nil)





	var tags []types.Tag


	err = faker.FakeObject(&tags)
	if err != nil {


		t.Fatal(err)


	}





	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&timestreamwrite.ListTagsForResourceOutput{Tags: tags}, nil)









	return aws_client.AwsServices{Timestreamwrite: m}
}



func TestTimestreamDatabases(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsTimestreamDatabasesGenerator{}), buildTimestreamDatabasesMock, aws_client.TestOptions{})


}


