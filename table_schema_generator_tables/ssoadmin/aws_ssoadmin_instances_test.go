

package ssoadmin







import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"




	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"




	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)







func buildInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)


	im := types.InstanceMetadata{}


	ps := types.PermissionSet{}


	as := types.AccountAssignment{}


	ip := `{"key": "value"}`
	err := faker.FakeObject(&ps)


	if err != nil {


		t.Fatal(err)
	}
	err = faker.FakeObject(&as)
	if err != nil {




		t.Fatal(err)




	}
	err = faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)




	}





	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ssoadmin.ListInstancesOutput{


			Instances: []types.InstanceMetadata{im},
		}, nil)



	mSSOAdmin.EXPECT().ListPermissionSets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ssoadmin.ListPermissionSetsOutput{




			PermissionSets: []string{*ps.Name},


		}, nil)







	mSSOAdmin.EXPECT().DescribePermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ssoadmin.DescribePermissionSetOutput{
			PermissionSet: &ps,


		}, nil)

	mSSOAdmin.EXPECT().ListAccountAssignments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ssoadmin.ListAccountAssignmentsOutput{


			AccountAssignments: []types.AccountAssignment{as},
		}, nil)



	mSSOAdmin.EXPECT().GetInlinePolicyForPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ssoadmin.GetInlinePolicyForPermissionSetOutput{
			InlinePolicy: &ip,




		}, nil)







	return aws_client.AwsServices{


		Ssoadmin: mSSOAdmin,




	}


}









func TestSSOAdminInstances(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSsoadminInstancesGenerator{}), buildInstances, aws_client.TestOptions{})
}


