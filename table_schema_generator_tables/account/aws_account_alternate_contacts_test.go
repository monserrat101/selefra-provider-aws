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









func buildAlternateContacts(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	mock := mocks.NewMockAccountClient(ctrl)



	var ac types.AlternateContact
	if err := faker.FakeObject(&ac); err != nil {


		t.Fatal(err)




	}
	mock.EXPECT().GetAlternateContact(


		gomock.Any(),




		gomock.Any(),
		gomock.Any(),




	).AnyTimes().Return(


		&account.GetAlternateContactOutput{AlternateContact: &ac},


		nil,




	).MinTimes(1)

	return aws_client.AwsServices{Account: mock}
}





func TestAlternateContacts(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAccountAlternateContactsGenerator{}), buildAlternateContacts, aws_client.TestOptions{})
}




