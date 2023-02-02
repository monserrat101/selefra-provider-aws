



package neptune

import (
	"github.com/selefra/selefra-provider-aws/constants"




	"testing"



	"github.com/aws/aws-sdk-go-v2/aws"


	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)







func buildNeptuneEventSubscriptions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	mock := mocks.NewMockNeptuneClient(ctrl)
	var s types.EventSubscription




	if err := faker.FakeObject(&s); err != nil {


		t.Fatal(err)


	}




	mock.EXPECT().DescribeEventSubscriptions(gomock.Any(), &neptune.DescribeEventSubscriptionsInput{
		Filters: []types.Filter{{Name: aws.String(constants.Engine), Values: []string{constants.Neptune}}},
	}, gomock.Any()).AnyTimes().Return(




		&neptune.DescribeEventSubscriptionsOutput{EventSubscriptionsList: []types.EventSubscription{s}},
		nil,




	)







	mock.EXPECT().ListTagsForResource(


		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn},
		gomock.Any(),




	).AnyTimes().Return(




		&neptune.ListTagsForResourceOutput{


			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},




		},




		nil,
	)
	return aws_client.AwsServices{Neptune: mock}
}





func TestNeptuneEventSubscriptions(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsNeptuneEventSubscriptionsGenerator{}), buildNeptuneEventSubscriptions, aws_client.TestOptions{})


}


