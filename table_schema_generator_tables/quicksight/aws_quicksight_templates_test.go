



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





func buildTemplatesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockQuicksightClient(ctrl)



	var lo quicksight.ListTemplatesOutput
	if err := faker.FakeObject(&lo); err != nil {


		t.Fatal(err)




	}


	lo.NextToken = nil




	m.EXPECT().ListTemplates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&lo, nil)



	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {


		t.Fatal(err)
	}




	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&to, nil)









	return aws_client.AwsServices{
		Quicksight: m,




	}
}

func TestQuicksightTemplates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsQuicksightTemplatesGenerator{}), buildTemplatesMock, aws_client.TestOptions{})
}




