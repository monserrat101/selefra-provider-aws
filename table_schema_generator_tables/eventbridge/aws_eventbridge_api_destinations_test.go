



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



func buildEventbridgeApiDestinationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEventbridgeClient(ctrl)
	object := types.ApiDestination{}




	err := faker.FakeObject(&object)




	if err != nil {
		t.Fatal(err)
	}





	m.EXPECT().ListApiDestinations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&eventbridge.ListApiDestinationsOutput{


			ApiDestinations: []types.ApiDestination{object},


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



func TestEventbridgeApiDestinations(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEventbridgeApiDestinationsGenerator{}), buildEventbridgeApiDestinationsMock, aws_client.TestOptions{})




}
