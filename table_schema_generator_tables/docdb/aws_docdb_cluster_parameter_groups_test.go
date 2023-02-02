



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





func buildClusterParameterGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockDocdbClient(ctrl)




	services := aws_client.AwsServices{




		Docdb: m,
	}


	var parameterGroups docdb.DescribeDBClusterParameterGroupsOutput


	if err := faker.FakeObject(&parameterGroups); err != nil {
		t.Fatal(err)
	}
	parameterGroups.Marker = nil




	m.EXPECT().DescribeDBClusterParameterGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&parameterGroups,


		nil,


	)









	var parameters docdb.DescribeDBClusterParametersOutput
	if err := faker.FakeObject(&parameters); err != nil {
		t.Fatal(err)
	}




	parameters.Marker = nil




	m.EXPECT().DescribeDBClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&parameters,




		nil,


	)





	var tags docdb.ListTagsForResourceOutput


	if err := faker.FakeObject(&tags); err != nil {




		t.Fatal(err)
	}


	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&tags,




		nil,


	)







	return services
}









func TestClusterParameterGroups(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbClusterParameterGroupsGenerator{}), buildClusterParameterGroupsMock, aws_client.TestOptions{})


}


