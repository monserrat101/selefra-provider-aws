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



func buildElasticSearchPackages(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockElasticsearchserviceClient(ctrl)





	var pkg types.PackageDetails




	if err := faker.FakeObject(&pkg); err != nil {


		t.Fatal(err)
	}
	m.EXPECT().DescribePackages(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&elasticsearchservice.DescribePackagesOutput{


			PackageDetailsList: []types.PackageDetails{pkg},




		},
		nil,
	)

	return aws_client.AwsServices{Elasticsearchservice: m}


}









func TestElasticSearchPackages(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticsearchPackagesGenerator{}), buildElasticSearchPackages, aws_client.TestOptions{})




}
