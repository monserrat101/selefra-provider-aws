



package account





import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/account"


	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)



func buildContacts(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	mock := mocks.NewMockAccountClient(ctrl)







	var ci types.ContactInformation
	if err := faker.FakeObject(&ci); err != nil {




		t.Fatal(err)


	}
	mock.EXPECT().GetContactInformation(




		gomock.Any(),


		&account.GetContactInformationInput{},


		gomock.Any(),


	).AnyTimes().Return(
		&account.GetContactInformationOutput{ContactInformation: &ci},
		nil,
	)





	return aws_client.AwsServices{Account: mock}


}



func TestContacts(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAccountContactsGenerator{}), buildContacts, aws_client.TestOptions{})




}


