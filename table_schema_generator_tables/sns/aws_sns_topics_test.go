package sns





import (




	"github.com/selefra/selefra-provider-aws/constants"


	"testing"







	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)



func buildSnsTopics(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockSnsClient(ctrl)




	topic := types.Topic{}




	tag := types.Tag{}
	err := faker.FakeObject(&topic)
	if err != nil {




		t.Fatal(err)
	}
	tagerr := faker.FakeObject(&tag)


	if tagerr != nil {
		t.Fatal(tagerr)




	}





	m.EXPECT().ListTopics(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sns.ListTopicsOutput{




			Topics: []types.Topic{topic},




		}, nil)




	m.EXPECT().GetTopicAttributes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&sns.GetTopicAttributesOutput{


			Attributes: map[string]string{


				constants.SubscriptionsConfirmed:		constants.Constants_51,


				constants.SubscriptionsDeleted:		constants.Constants_52,
				constants.SubscriptionsPending:		constants.Constants_53,
				constants.FifoTopic:			constants.False,




				constants.ContentBasedDeduplication:	constants.True,
				constants.DisplayName:			constants.Selefra,


				constants.KmsMasterKeyId:			constants.Testkey,


				constants.Owner:				"owner",




				constants.Policy:				`{"stuff": 3}`,




				constants.DeliveryPolicy:			`{"stuff": 3}`,




				constants.EffectiveDeliveryPolicy:	`{"stuff": 3}`,


				constants.WeirdAndUnexpectedField:	constants.Needsupdating,


			},


		}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&sns.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		}, nil)
	return aws_client.AwsServices{


		Sns: m,
	}




}

func TestSnsTopics(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSnsTopicsGenerator{}), buildSnsTopics, aws_client.TestOptions{})
}


