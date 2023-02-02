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







func buildCustomVerificationEmailTemplates(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	sesClient := mocks.NewMockSesv2Client(ctrl)





	metadata := types.CustomVerificationEmailTemplateMetadata{}


	if err := faker.FakeObject(&metadata); err != nil {




		t.Fatal(err)


	}







	get := new(sesv2.GetCustomVerificationEmailTemplateOutput)
	if err := faker.FakeObject(get); err != nil {


		t.Fatal(err)


	}


	metadata.TemplateName = get.TemplateName







	sesClient.EXPECT().ListCustomVerificationEmailTemplates(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&sesv2.ListCustomVerificationEmailTemplatesOutput{




			CustomVerificationEmailTemplates: []types.CustomVerificationEmailTemplateMetadata{metadata},


		},
		nil,
	)
	sesClient.EXPECT().GetCustomVerificationEmailTemplate(gomock.Any(), gomock.Any()).AnyTimes().Return(




		get,
		nil,
	)







	return aws_client.AwsServices{


		Sesv2: sesClient,
	}
}



func TestCustomVerificationEmailTemplates(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSesCustomVerificationEmailTemplatesGenerator{}), buildCustomVerificationEmailTemplates, aws_client.TestOptions{})




}
