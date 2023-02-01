

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





func buildKafkaConfigurationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockKafkaClient(ctrl)
	object := types.Configuration{}




	err := faker.FakeObject(&object)


	if err != nil {
		t.Fatal(err)




	}

	m.EXPECT().ListConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&kafka.ListConfigurationsOutput{
			Configurations: []types.Configuration{object},




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

func TestKafkaConfigurations(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsKafkaConfigurationsGenerator{}), buildKafkaConfigurationsMock, aws_client.TestOptions{})
}


