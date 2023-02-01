package kafka



import (


	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"


	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)



func buildKafkaNodesMock(t *testing.T, m *mocks.MockKafkaClient) {
	object := types.NodeInfo{}




	err := faker.FakeObject(&object)


	if err != nil {
		t.Fatal(err)
	}



	m.EXPECT().ListNodes(gomock.Any(), gomock.Any(), gomock.Any()).Return(




		&kafka.ListNodesOutput{




			NodeInfoList: []types.NodeInfo{object},
		}, nil)




}




func buildKafkaClusterOperationsMock(t *testing.T, m *mocks.MockKafkaClient) {
	object := types.ClusterOperationInfo{}
	err := faker.FakeObject(&object)


	if err != nil {


		t.Fatal(err)
	}







	m.EXPECT().ListClusterOperations(gomock.Any(), gomock.Any(), gomock.Any()).Return(




		&kafka.ListClusterOperationsOutput{
			ClusterOperationInfoList: []types.ClusterOperationInfo{object},
		}, nil)




}



func buildKafkaClustersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockKafkaClient(ctrl)




	object := types.Cluster{}
	err := faker.FakeObject(&object)


	if err != nil {


		t.Fatal(err)




	}




	buildKafkaNodesMock(t, m)
	buildKafkaClusterOperationsMock(t, m)



	m.EXPECT().ListClustersV2(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&kafka.ListClustersV2Output{




			ClusterInfoList: []types.Cluster{object},




		}, nil)





	m.EXPECT().DescribeClusterV2(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&kafka.DescribeClusterV2Output{




			ClusterInfo: &object,
		}, nil)





	tagsOutput := kafka.ListTagsForResourceOutput{}


	err = faker.FakeObject(&tagsOutput)
	if err != nil {
		t.Fatal(err)




	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).AnyTimes().Return(&tagsOutput, nil).AnyTimes()

	return aws_client.AwsServices{
		Kafka: m,
	}
}









func TestKafkaClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsKafkaClustersGenerator{}), buildKafkaClustersMock, aws_client.TestOptions{})
}
