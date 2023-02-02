package identitystore





import (




	"testing"







	"github.com/aws/aws-sdk-go-v2/service/identitystore"




	iTypes "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)







func buildGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	mIdentity := mocks.NewMockIdentitystoreClient(ctrl)




	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}




	err := faker.FakeObject(&im)




	if err != nil {


		t.Fatal(err)
	}
	group := iTypes.Group{}




	err = faker.FakeObject(&group)
	if err != nil {
		t.Fatal(err)


	}


	groupMembership := iTypes.GroupMembership{}
	err = faker.FakeObject(&groupMembership)


	if err != nil {
		t.Fatal(err)


	}



	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ssoadmin.ListInstancesOutput{


			Instances: []types.InstanceMetadata{im},
		}, nil)




	mIdentity.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&identitystore.ListGroupsOutput{
			Groups: []iTypes.Group{group},
		}, nil)




	mIdentity.EXPECT().ListGroupMemberships(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&identitystore.ListGroupMembershipsOutput{


			GroupMemberships: []iTypes.GroupMembership{groupMembership},




		}, nil)



	return aws_client.AwsServices{
		Ssoadmin:	mSSOAdmin,




		Identitystore:	mIdentity,




	}
}









func TestIdentityStoreGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIdentitystoreGroupsGenerator{}), buildGroups, aws_client.TestOptions{})




}


