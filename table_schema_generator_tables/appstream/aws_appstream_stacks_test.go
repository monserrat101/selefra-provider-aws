

package appstream







import (


	"testing"



	"github.com/aws/aws-sdk-go-v2/service/appstream"


	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildAppstreamStacksMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAppstreamClient(ctrl)




	stack := types.Stack{}




	err := faker.FakeObject(&stack)
	if err != nil {
		t.Fatal(err)
	}





	stackEntitlements := types.Entitlement{}


	if faker.FakeObject(&stackEntitlements) != nil {


		t.Fatal(err)


	}



	stackUserAssociations := types.UserStackAssociation{}




	if faker.FakeObject(&stackUserAssociations) != nil {


		t.Fatal(err)


	}

	m.EXPECT().DescribeStacks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&appstream.DescribeStacksOutput{


			Stacks: []types.Stack{stack},


		}, nil)







	m.EXPECT().DescribeEntitlements(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&appstream.DescribeEntitlementsOutput{
			Entitlements: []types.Entitlement{stackEntitlements},




		}, nil)







	m.EXPECT().DescribeUserStackAssociations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&appstream.DescribeUserStackAssociationsOutput{


			UserStackAssociations: []types.UserStackAssociation{stackUserAssociations},




		}, nil)









	return aws_client.AwsServices{
		Appstream: m,


	}




}

func TestAppstreamStacks(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAppstreamStacksGenerator{}), buildAppstreamStacksMock, aws_client.TestOptions{})




}
