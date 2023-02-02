



package wafv2









import (




	"github.com/selefra/selefra-provider-aws/constants"


	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"


	"github.com/aws/aws-sdk-go-v2/service/wafv2"




	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"




	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)







func buildIpsetsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafv2Client(ctrl)



	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {




		var s types.IPSet




		if err := faker.FakeObject(&s); err != nil {




			t.Fatal(err)




		}




		s.Addresses = []string{constants.Constants_66, constants.Db}




		m.EXPECT().ListIPSets(


			gomock.Any(),


			&wafv2.ListIPSetsInput{Scope: scope, Limit: aws.Int32(100)},
			gomock.Any(),




		).AnyTimes().Return(




			&wafv2.ListIPSetsOutput{
				IPSets: []types.IPSetSummary{{Id: s.Id, Name: s.Name}},


			},




			nil,




		)







		m.EXPECT().GetIPSet(


			gomock.Any(),




			&wafv2.GetIPSetInput{Name: s.Name, Id: s.Id, Scope: scope},


			gomock.Any(),


		).AnyTimes().Return(




			&wafv2.GetIPSetOutput{IPSet: &s},
			nil,
		)







		m.EXPECT().ListTagsForResource(
			gomock.Any(),




			&wafv2.ListTagsForResourceInput{ResourceARN: s.ARN},
			gomock.Any(),


		).AnyTimes().Return(


			&wafv2.ListTagsForResourceOutput{
				TagInfoForResource: &types.TagInfoForResource{


					TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
				},




			},




			nil,


		)
	}





	return aws_client.AwsServices{Wafv2: m}


}









func TestWafV2IPSets(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafv2IpsetsGenerator{}), buildIpsetsMock, aws_client.TestOptions{})
}


