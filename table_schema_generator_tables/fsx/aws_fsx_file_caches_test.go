package fsx









import (


	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


	"github.com/stretchr/testify/require"


)





func buildFileCachesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)





	var fc types.FileCache
	require.NoError(t, faker.FakeObject(&fc))


	m.EXPECT().DescribeFileCaches(
		gomock.Any(),
		gomock.Any(),




	).AnyTimes().Return(
		&fsx.DescribeFileCachesOutput{FileCaches: []types.FileCache{fc}},
		nil,


	)





	var tags []types.Tag
	require.NoError(t, faker.FakeObject(&tags))





	m.EXPECT().ListTagsForResource(


		gomock.Any(),




		gomock.Any(),




	).AnyTimes().Return(


		&fsx.ListTagsForResourceOutput{Tags: tags},




		nil,




	)









	return aws_client.AwsServices{
		Fsx: m,




	}


}



func TestFileCaches(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxFileCachesGenerator{}), buildFileCachesMock, aws_client.TestOptions{})


}
