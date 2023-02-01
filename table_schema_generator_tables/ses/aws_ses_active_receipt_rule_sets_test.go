

package ses







import (
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/ses"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildActiveReceiptRuleSets(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	sesClient := mocks.NewMockSesClient(ctrl)





	data := new(ses.DescribeActiveReceiptRuleSetOutput)
	if err := faker.FakeObject(data); err != nil {




		t.Fatal(err)


	}

	sesClient.EXPECT().DescribeActiveReceiptRuleSet(gomock.Any(), gomock.Any()).AnyTimes().Return(data, nil)



	return aws_client.AwsServices{


		Ses: sesClient,
	}
}







func TestActiveReceiptRuleSets(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSesActiveReceiptRuleSetsGenerator{}), buildActiveReceiptRuleSets, aws_client.TestOptions{})


}


