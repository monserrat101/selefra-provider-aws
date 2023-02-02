

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

func buildRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafregionalClient(ctrl)







	var g types.RuleGroup




	if err := faker.FakeObject(&g); err != nil {




		t.Fatal(err)




	}




	m.EXPECT().ListRuleGroups(




		gomock.Any(),


		&wafregional.ListRuleGroupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupSummary{{RuleGroupId: g.RuleGroupId}},


		},
		nil,




	)



	m.EXPECT().GetRuleGroup(




		gomock.Any(),
		&wafregional.GetRuleGroupInput{RuleGroupId: g.RuleGroupId},




		gomock.Any(),


	).AnyTimes().Return(


		&wafregional.GetRuleGroupOutput{RuleGroup: &g},


		nil,


	)





	m.EXPECT().ListTagsForResource(




		gomock.Any(),




		&wafregional.ListTagsForResourceInput{


			ResourceARN: aws.String(fmt.Sprintf(constants.ArnawswafregionaluseasttestAccountrulegroupv, *g.RuleGroupId)),




		},




		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListTagsForResourceOutput{},


		nil,
	)





	return aws_client.AwsServices{Wafregional: m}


}



func TestRuleGroups(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafregionalRuleGroupsGenerator{}), buildRuleGroupsMock, aws_client.TestOptions{})
}




