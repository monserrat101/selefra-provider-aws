



package wafregional



import (
	"github.com/selefra/selefra-provider-aws/constants"


	"fmt"




	"testing"





	"github.com/aws/aws-sdk-go-v2/aws"




	"github.com/aws/aws-sdk-go-v2/service/wafregional"


	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildRateBasedRulesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafregionalClient(ctrl)



	var rule types.RateBasedRule
	if err := faker.FakeObject(&rule); err != nil {


		t.Fatal(err)


	}
	m.EXPECT().ListRateBasedRules(




		gomock.Any(),
		&wafregional.ListRateBasedRulesInput{},




		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListRateBasedRulesOutput{
			Rules: []types.RuleSummary{{RuleId: rule.RuleId}},




		},
		nil,
	)

	m.EXPECT().GetRateBasedRule(


		gomock.Any(),
		&wafregional.GetRateBasedRuleInput{RuleId: rule.RuleId},
		gomock.Any(),


	).AnyTimes().Return(


		&wafregional.GetRateBasedRuleOutput{Rule: &rule},




		nil,


	)









	m.EXPECT().ListTagsForResource(


		gomock.Any(),


		&wafregional.ListTagsForResourceInput{


			ResourceARN: aws.String(fmt.Sprintf(constants.ArnawswafregionaluseasttestAccountratebasedrulev, *rule.RuleId)),


		},




		gomock.Any(),
	).AnyTimes().Return(


		&wafregional.ListTagsForResourceOutput{},


		nil,


	)









	return aws_client.AwsServices{Wafregional: m}
}



func TestRateBasedRules(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafregionalRateBasedRulesGenerator{}), buildRateBasedRulesMock, aws_client.TestOptions{})
}
