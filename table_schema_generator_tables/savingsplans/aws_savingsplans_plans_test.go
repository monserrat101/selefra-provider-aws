package savingsplans





import (
	"testing"









	"github.com/aws/aws-sdk-go-v2/aws"




	"github.com/aws/aws-sdk-go-v2/service/savingsplans"




	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)





func buildSavingPlansPlans(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockSavingsplansClient(ctrl)







	var s types.SavingsPlan




	if err := faker.FakeObject(&s); err != nil {




		t.Fatal(err)
	}


	m.EXPECT().DescribeSavingsPlans(


		gomock.Any(),
		&savingsplans.DescribeSavingsPlansInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),




	).AnyTimes().Return(
		&savingsplans.DescribeSavingsPlansOutput{


			SavingsPlans: []types.SavingsPlan{s},
		},




		nil,


	)





	return aws_client.AwsServices{




		Savingsplans: m,


	}




}





func TestSavingsplansPlans(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSavingsplansPlansGenerator{}), buildSavingPlansPlans, aws_client.TestOptions{})


}
