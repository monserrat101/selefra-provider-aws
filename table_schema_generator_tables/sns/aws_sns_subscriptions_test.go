package sns









import (




	"github.com/selefra/selefra-provider-aws/constants"
	"testing"



	"github.com/aws/aws-sdk-go-v2/aws"


	"github.com/aws/aws-sdk-go-v2/service/sns"




	"github.com/aws/aws-sdk-go-v2/service/sns/types"




	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)



func buildSnsSubscriptions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockSnsClient(ctrl)




	sub := types.Subscription{}
	err := faker.FakeObject(&sub)
	if err != nil {


		t.Fatal(err)




	}





	subTemp := types.Subscription{}




	err = faker.FakeObject(&subTemp)




	if err != nil {




		t.Fatal(err)
	}


	emptySub := types.Subscription{
		SubscriptionArn:	aws.String(constants.PendingConfirmation),


		Owner:			subTemp.Owner,




		Protocol:		subTemp.Protocol,


		TopicArn:		subTemp.TopicArn,
		Endpoint:		subTemp.Endpoint,
	}



	m.EXPECT().ListSubscriptions(
		gomock.Any(),
		&sns.ListSubscriptionsInput{},


	).AnyTimes().Return(


		&sns.ListSubscriptionsOutput{


			Subscriptions: []types.Subscription{sub, emptySub},


		}, nil)

	m.EXPECT().GetSubscriptionAttributes(
		gomock.Any(),


		&sns.GetSubscriptionAttributesInput{SubscriptionArn: sub.SubscriptionArn},


	).AnyTimes().Return(




		&sns.GetSubscriptionAttributesOutput{Attributes: map[string]string{


			constants.ConfirmationWasAuthenticated:	constants.True,




			constants.DeliveryPolicy:			constants.Constants_48,


			constants.EffectiveDeliveryPolicy:	constants.Constants_49,




			constants.FilterPolicy:			constants.Constants_50,




			constants.PendingConfirmation:		constants.True,




			constants.RawMessageDelivery:		constants.True,




			constants.RedrivePolicy:			`{"deadLetterTargetArn": "test"}`,
			constants.SubscriptionRoleArn:		constants.Some,


			constants.WeirdAndUnexpectedField:	constants.Needsupdating,


		}},


		nil,


	)







	return aws_client.AwsServices{


		Sns: m,
	}
}





func TestSnsSubscriptions(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSnsSubscriptionsGenerator{}), buildSnsSubscriptions, aws_client.TestOptions{})


}


