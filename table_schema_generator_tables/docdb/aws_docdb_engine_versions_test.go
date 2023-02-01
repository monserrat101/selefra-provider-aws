

package docdb







import (


	"testing"





	"github.com/aws/aws-sdk-go-v2/service/docdb"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDocdbClient(ctrl)


	services := aws_client.AwsServices{
		Docdb: m,


	}
	var ev docdb.DescribeDBEngineVersionsOutput
	if err := faker.FakeObject(&ev); err != nil {


		t.Fatal(err)


	}
	ev.Marker = nil
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ev,




		nil,




	)







	var parameters docdb.DescribeEngineDefaultClusterParametersOutput
	if err := faker.FakeObject(&parameters); err != nil {


		t.Fatal(err)


	}
	m.EXPECT().DescribeEngineDefaultClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&parameters,
		nil,


	)



	var instanceOptions docdb.DescribeOrderableDBInstanceOptionsOutput


	if err := faker.FakeObject(&instanceOptions); err != nil {


		t.Fatal(err)


	}


	instanceOptions.Marker = nil




	m.EXPECT().DescribeOrderableDBInstanceOptions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&instanceOptions,




		nil,


	)





	return services
}



func TestEngineVersions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbEngineVersionsGenerator{}), buildEngineVersionsMock, aws_client.TestOptions{})




}
