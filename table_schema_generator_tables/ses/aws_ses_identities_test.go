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



func buildIdentities(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	sesClient := mocks.NewMockSesv2Client(ctrl)







	ei := types.IdentityInfo{}
	if err := faker.FakeObject(&ei); err != nil {
		t.Fatal(err)
	}
	o := sesv2.GetEmailIdentityOutput{}


	if err := faker.FakeObject(&o); err != nil {
		t.Fatal(err)


	}


	o.IdentityType = ei.IdentityType

	sesClient.EXPECT().ListEmailIdentities(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&sesv2.ListEmailIdentitiesOutput{EmailIdentities: []types.IdentityInfo{ei}},
		nil,
	)




	sesClient.EXPECT().GetEmailIdentity(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&o,
		nil,
	)





	return aws_client.AwsServices{




		Sesv2: sesClient,




	}
}



func TestIdentities(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSesIdentitiesGenerator{}), buildIdentities, aws_client.TestOptions{})


}




