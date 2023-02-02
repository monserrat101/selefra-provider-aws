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









func buildAppstreamDirectoryConfigsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.DirectoryConfig{}
	err := faker.FakeObject(&object)




	if err != nil {


		t.Fatal(err)


	}





	m.EXPECT().DescribeDirectoryConfigs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&appstream.DescribeDirectoryConfigsOutput{
			DirectoryConfigs: []types.DirectoryConfig{object},




		}, nil)





	tagsOutput := appstream.ListTagsForResourceOutput{}


	err = faker.FakeObject(&tagsOutput)


	if err != nil {
		t.Fatal(err)


	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).AnyTimes().Return(&tagsOutput, nil).AnyTimes()









	return aws_client.AwsServices{




		Appstream: m,




	}




}





func TestAppstreamDirectoryConfigs(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAppstreamDirectoryConfigsGenerator{}), buildAppstreamDirectoryConfigsMock, aws_client.TestOptions{})
}


