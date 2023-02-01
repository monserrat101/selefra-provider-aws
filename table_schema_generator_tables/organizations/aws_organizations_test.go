



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





func buildOrganizations(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockOrganizationsClient(ctrl)





	o := organizations.DescribeOrganizationOutput{}
	err := faker.FakeObject(&o)


	if err != nil {




		t.Fatal(err)




	}



	m.EXPECT().DescribeOrganization(gomock.Any(), gomock.Any()).AnyTimes().Return(&o, nil)







	return aws_client.AwsServices{
		Organizations: m,


	}


}

func TestOrganizations(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsOrganizationsGenerator{}), buildOrganizations, aws_client.TestOptions{})




}
