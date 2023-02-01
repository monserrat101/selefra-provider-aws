



package cloudwatchlogs





import (


	"github.com/selefra/selefra-provider-aws/constants"




	"testing"







	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"


	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)

func buildResourcePolicies(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockCloudwatchlogsClient(ctrl)
	rp := types.ResourcePolicy{}




	err := faker.FakeObject(&rp)
	if err != nil {


		t.Fatal(err)




	}


	policyDocument := constants.Constants_37


	rp.PolicyDocument = &policyDocument
	m.EXPECT().DescribeResourcePolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&cloudwatchlogs.DescribeResourcePoliciesOutput{


			ResourcePolicies: []types.ResourcePolicy{rp},




		}, nil)


	return aws_client.AwsServices{


		Cloudwatchlogs: m,




	}




}



func TestResourcePolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudwatchlogsResourcePoliciesGenerator{}), buildResourcePolicies, aws_client.TestOptions{})
}




