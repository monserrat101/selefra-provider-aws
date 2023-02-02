

package iam









import (




	"testing"









	"github.com/aws/aws-sdk-go-v2/service/iam"


	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"


	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

func buildIamInstanceProfiles(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockIamClient(ctrl)


	p := iamTypes.InstanceProfile{}
	err := faker.FakeObject(&p)


	if err != nil {
		t.Fatal(err)




	}







	m.EXPECT().ListInstanceProfiles(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListInstanceProfilesOutput{


			InstanceProfiles: []iamTypes.InstanceProfile{p},
		}, nil)







	tag := iamTypes.Tag{}


	err = faker.FakeObject(&tag)
	if err != nil {




		t.Fatal(err)
	}
	m.EXPECT().ListInstanceProfileTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListInstanceProfileTagsOutput{
			Tags: []iamTypes.Tag{




				tag,
			},


		}, nil)







	return aws_client.AwsServices{




		IAM: m,




	}
}



func TestIamInstanceProfiles(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamInstanceProfilesGenerator{}), buildIamInstanceProfiles, aws_client.TestOptions{})




}
