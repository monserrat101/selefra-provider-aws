



package eventbridge

import (
	"testing"



	"github.com/aws/aws-sdk-go-v2/service/eventbridge"




	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"


	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)





func buildEventbridgeEventSourcesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEventbridgeClient(ctrl)




	object := types.EventSource{}


	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)




	}









	m.EXPECT().ListEventSources(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&eventbridge.ListEventSourcesOutput{
			EventSources: []types.EventSource{object},




		}, nil)





	tagsOutput := eventbridge.ListTagsForResourceOutput{}




	err = faker.FakeObject(&tagsOutput)
	if err != nil {




		t.Fatal(err)




	}


	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).AnyTimes().Return(&tagsOutput, nil).AnyTimes()




	return aws_client.AwsServices{


		Eventbridge: m,
	}
}



func TestEventbridgeEventSources(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEventbridgeEventSourcesGenerator{}), buildEventbridgeEventSourcesMock, aws_client.TestOptions{})
}




