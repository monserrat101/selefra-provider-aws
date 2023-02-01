



package ses





import (
	"testing"



	"github.com/aws/aws-sdk-go-v2/service/sesv2"




	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)









func buildContactLists(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	sesClient := mocks.NewMockSesv2Client(ctrl)





	cs := sesv2.GetContactListOutput{}
	if err := faker.FakeObject(&cs); err != nil {


		t.Fatal(err)
	}









	sesClient.EXPECT().ListContactLists(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&sesv2.ListContactListsOutput{ContactLists: []types.ContactList{{ContactListName: cs.ContactListName}}},




		nil,


	)


	sesClient.EXPECT().GetContactList(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&cs,
		nil,


	)

	return aws_client.AwsServices{


		Sesv2: sesClient,


	}
}







func TestContactLists(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSesContactListsGenerator{}), buildContactLists, aws_client.TestOptions{})
}


