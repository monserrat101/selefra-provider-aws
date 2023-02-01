

package elasticache



import (




	"testing"



	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)









func buildElasticacheClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	mockElasticache := mocks.NewMockElasticacheClient(ctrl)




	output := elasticache.DescribeCacheClustersOutput{}




	err := faker.FakeObject(&output)




	output.Marker = nil




	if err != nil {




		t.Fatal(err)




	}







	var ta types.Tag


	if err = faker.FakeObject(&ta); err != nil {


		t.Fatal(err)
	}





	mockElasticache.EXPECT().DescribeCacheClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&output, nil)
	mockElasticache.EXPECT().ListTagsForResource(gomock.Any(), &elasticache.ListTagsForResourceInput{ResourceName: output.CacheClusters[0].ARN}, gomock.Any()).AnyTimes().Return(&elasticache.ListTagsForResourceOutput{TagList: []types.Tag{ta}}, nil)









	return aws_client.AwsServices{


		Elasticache: mockElasticache,


	}


}









func TestElasticacheClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticacheClustersGenerator{}), buildElasticacheClusters, aws_client.TestOptions{})
}




