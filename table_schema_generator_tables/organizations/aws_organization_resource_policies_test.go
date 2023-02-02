package organizations







import (




	"testing"



	"github.com/aws/aws-sdk-go-v2/service/organizations"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildResourcePolicy(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockOrganizationsClient(ctrl)

	o := organizations.DescribeResourcePolicyOutput{}
	err := faker.FakeObject(&o)




	if err != nil {


		t.Fatal(err)


	}





	m.EXPECT().DescribeResourcePolicy(gomock.Any(), gomock.Any()).AnyTimes().Return(&o, nil)



	return aws_client.AwsServices{




		Organizations: m,




	}




}



func TestResourcePolicies(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsOrganizationResourcePoliciesGenerator{}), buildResourcePolicy, aws_client.TestOptions{})
}


