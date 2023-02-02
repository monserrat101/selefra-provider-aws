package quicksight

import (




	"testing"



	"github.com/aws/aws-sdk-go-v2/service/quicksight"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)





func buildFoldersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockQuicksightClient(ctrl)







	var lo quicksight.ListFoldersOutput
	if err := faker.FakeObject(&lo); err != nil {




		t.Fatal(err)


	}


	lo.NextToken = nil




	m.EXPECT().ListFolders(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&lo, nil)





	var qs quicksight.DescribeFolderOutput


	if err := faker.FakeObject(&qs); err != nil {
		t.Fatal(err)




	}
	m.EXPECT().DescribeFolder(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&qs, nil)







	var to quicksight.ListTagsForResourceOutput


	if err := faker.FakeObject(&to); err != nil {




		t.Fatal(err)


	}


	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&to, nil)

	return aws_client.AwsServices{


		Quicksight: m,


	}
}







func TestQuicksightFolders(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsQuicksightFoldersGenerator{}), buildFoldersMock, aws_client.TestOptions{})


}
