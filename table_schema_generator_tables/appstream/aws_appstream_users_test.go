



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

func buildAppstreamUsersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.User{}




	err := faker.FakeObject(&object)
	if err != nil {




		t.Fatal(err)




	}



	m.EXPECT().DescribeUsers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&appstream.DescribeUsersOutput{


			Users: []types.User{object},




		}, nil)



	return aws_client.AwsServices{


		Appstream: m,
	}


}







func TestAppstreamUsers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAppstreamUsersGenerator{}), buildAppstreamUsersMock, aws_client.TestOptions{})
}


