



package securityhub

import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/securityhub"


	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"




	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)





func buildFindings(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	shMock := mocks.NewMockSecurityhubClient(ctrl)


	findings := types.AwsSecurityFinding{}
	err := faker.FakeObject(&findings)


	if err != nil {
		t.Fatal(err)


	}





	shMock.EXPECT().GetFindings(


		gomock.Any(),




		&securityhub.GetFindingsInput{


			MaxResults: 100,


		},




	).AnyTimes().Return(


		&securityhub.GetFindingsOutput{


			Findings: []types.AwsSecurityFinding{findings},




		},




		nil,


	)







	return aws_client.AwsServices{Securityhub: shMock}
}





func TestFindings(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSecurityhubFindingsGenerator{}), buildFindings, aws_client.TestOptions{})
}


