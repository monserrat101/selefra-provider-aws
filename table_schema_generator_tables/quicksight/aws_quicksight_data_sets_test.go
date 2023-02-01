

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





func buildDataSetsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockQuicksightClient(ctrl)





	var ld quicksight.ListDataSetsOutput


	if err := faker.FakeObject(&ld); err != nil {


		t.Fatal(err)


	}
	ld.NextToken = nil
	m.EXPECT().ListDataSets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&ld, nil)







	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {




		t.Fatal(err)


	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&to, nil)





	var ig quicksight.ListIngestionsOutput
	if err := faker.FakeObject(&ig); err != nil {




		t.Fatal(err)


	}


	ig.NextToken = nil
	m.EXPECT().ListIngestions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&ig, nil)







	var to2 quicksight.ListTagsForResourceOutput


	if err := faker.FakeObject(&to2); err != nil {




		t.Fatal(err)




	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&to2, nil)





	return aws_client.AwsServices{




		Quicksight: m,
	}




}





func TestQuicksightDataSets(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsQuicksightDataSetsGenerator{}), buildDataSetsMock, aws_client.TestOptions{})




}
