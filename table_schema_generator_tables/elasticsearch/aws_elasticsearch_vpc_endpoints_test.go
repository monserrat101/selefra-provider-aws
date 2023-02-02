



package elasticsearch









import (
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"


	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"




	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildElasticSearchVpcEndpoints(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockElasticsearchserviceClient(ctrl)

	var summary types.VpcEndpointSummary


	if err := faker.FakeObject(&summary); err != nil {
		t.Fatal(err)


	}




	m.EXPECT().ListVpcEndpoints(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticsearchservice.ListVpcEndpointsOutput{


			VpcEndpointSummaryList: []types.VpcEndpointSummary{summary},
		},




		nil,


	)





	var endpoint types.VpcEndpoint




	if err := faker.FakeObject(&endpoint); err != nil {


		t.Fatal(err)




	}


	endpoint.VpcEndpointId = summary.VpcEndpointId

	m.EXPECT().DescribeVpcEndpoints(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticsearchservice.DescribeVpcEndpointsOutput{




			VpcEndpoints: []types.VpcEndpoint{endpoint},




		},
		nil,
	)



	return aws_client.AwsServices{Elasticsearchservice: m}
}





func TestElasticSearchVpcEndpoints(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticsearchVpcEndpointsGenerator{}), buildElasticSearchVpcEndpoints, aws_client.TestOptions{})




}
