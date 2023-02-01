

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









func buildElasticSearchVersions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockElasticsearchserviceClient(ctrl)









	var versions []string
	if err := faker.FakeObject(&versions); err != nil {
		t.Fatal(err)
	}


	m.EXPECT().ListElasticsearchVersions(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticsearchservice.ListElasticsearchVersionsOutput{


			ElasticsearchVersions: versions,




		},


		nil,




	)



	var instanceTypes []types.ESPartitionInstanceType




	if err := faker.FakeObject(&versions); err != nil {




		t.Fatal(err)


	}




	m.EXPECT().ListElasticsearchInstanceTypes(gomock.Any(), gomock.Any()).AnyTimes().Return(


		&elasticsearchservice.ListElasticsearchInstanceTypesOutput{


			ElasticsearchInstanceTypes: instanceTypes,


		},
		nil,
	)



	return aws_client.AwsServices{Elasticsearchservice: m}




}



func TestElasticSearchVersions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticsearchVersionsGenerator{}), buildElasticSearchVersions, aws_client.TestOptions{})




}
