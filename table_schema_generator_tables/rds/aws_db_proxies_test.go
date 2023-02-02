



package rds







import (
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/rds"




	"github.com/aws/aws-sdk-go-v2/service/rds/types"




	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)





func buildRdsDbProxiesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockRdsClient(ctrl)
	proxy := types.DBProxy{}


	err := faker.FakeObject(&proxy)


	if err != nil {


		t.Fatal(err)
	}



	tags := rds.ListTagsForResourceOutput{}


	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}





	m.EXPECT().DescribeDBProxies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&rds.DescribeDBProxiesOutput{
			DBProxies: []types.DBProxy{proxy},
		}, nil)




	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&tags, nil)





	return aws_client.AwsServices{
		Rds: m,




	}




}





func TestRdsDbProxues(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDbProxiesGenerator{}), buildRdsDbProxiesMock, aws_client.TestOptions{})


}




