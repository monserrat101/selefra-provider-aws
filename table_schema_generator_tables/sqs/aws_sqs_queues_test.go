package sqs







import (




	"github.com/selefra/selefra-provider-aws/constants"


	"testing"



	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)





func buildSQSQueues(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	sqsMock := mocks.NewMockSqsClient(ctrl)







	var queueURL = constants.Httpsurl
	sqsMock.EXPECT().ListQueues(




		gomock.Any(),
		&sqs.ListQueuesInput{},


	).AnyTimes().Return(
		&sqs.ListQueuesOutput{QueueUrls: []string{queueURL}},




		nil,
	)

	sqsMock.EXPECT().GetQueueAttributes(




		gomock.Any(),
		&sqs.GetQueueAttributesInput{QueueUrl: &queueURL, AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll}},
	).AnyTimes().Return(




		&sqs.GetQueueAttributesOutput{




			Attributes: map[string]string{


				constants.Policy:					`{"field1":1}`,




				constants.VisibilityTimeout:			constants.Constants_54,




				constants.MaximumMessageSize:			constants.Constants_55,




				constants.MessageRetentionPeriod:			constants.Constants_56,




				constants.ApproximateNumberOfMessages:		constants.Constants_57,




				constants.ApproximateNumberOfMessagesNotVisible:	constants.Constants_58,


				constants.CreatedTimestamp:			constants.Constants_59,


				constants.LastModifiedTimestamp:			constants.Constants_60,
				constants.QueueArn:				constants.Arnawssqsuseastterraformexamplequeue,




				constants.ApproximateNumberOfMessagesDelayed:	constants.Constants_61,


				constants.DelaySeconds:				constants.Constants_62,
				constants.ReceiveMessageWaitTimeSeconds:		constants.Constants_63,
				constants.RedrivePolicy:				`{"field2":2}`,




				constants.FifoQueue:				constants.True,
				constants.ContentBasedDeduplication:		constants.False,


				constants.KmsMasterKeyId:				"key",
				constants.KmsDataKeyReusePeriodSeconds:		constants.Constants_64,




				constants.SqsManagedSseEnabled:			constants.True,
				constants.DeduplicationScope:			constants.MessageGroup,
				constants.FifoThroughputLimit:			constants.Queue,




				constants.RedriveAllowPolicy:			`{"field3":3}`,



				constants.UnexpectedField:	constants.SomeValue,
			},




		},
		nil,


	)







	sqsMock.EXPECT().ListQueueTags(
		gomock.Any(),




		&sqs.ListQueueTagsInput{QueueUrl: &queueURL},




	).AnyTimes().Return(
		&sqs.ListQueueTagsOutput{Tags: map[string]string{"tag": "value"}},
		nil,


	)
	return aws_client.AwsServices{Sqs: sqsMock}


}







func TestQueues(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSqsQueuesGenerator{}), buildSQSQueues, aws_client.TestOptions{})




}


