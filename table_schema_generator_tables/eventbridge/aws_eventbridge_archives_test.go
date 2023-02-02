

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







func buildEventbridgeArchivesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockEventbridgeClient(ctrl)




	object := types.Archive{}
	err := faker.FakeObject(&object)


	if err != nil {




		t.Fatal(err)
	}



	m.EXPECT().ListArchives(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&eventbridge.ListArchivesOutput{




			Archives: []types.Archive{object},
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









func TestEventbridgeArchives(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEventbridgeArchivesGenerator{}), buildEventbridgeArchivesMock, aws_client.TestOptions{})




}




