

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



func buildUsers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	mIdentity := mocks.NewMockIdentitystoreClient(ctrl)


	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}


	err := faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)


	}
	users := iTypes.User{}




	err = faker.FakeObject(&users)


	if err != nil {


		t.Fatal(err)


	}







	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ssoadmin.ListInstancesOutput{
			Instances: []types.InstanceMetadata{im},
		}, nil)




	mIdentity.EXPECT().ListUsers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&identitystore.ListUsersOutput{


			Users: []iTypes.User{users},


		}, nil)





	return aws_client.AwsServices{


		Ssoadmin:	mSSOAdmin,
		Identitystore:	mIdentity,


	}




}

func TestIdentityStoreUsers(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIdentitystoreUsersGenerator{}), buildUsers, aws_client.TestOptions{})
}


