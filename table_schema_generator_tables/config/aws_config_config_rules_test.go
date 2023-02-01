



package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"




	"github.com/aws/aws-sdk-go-v2/service/configservice/types"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)







func buildConfigRules(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockConfigserviceClient(ctrl)


	l := types.ConfigRule{}




	if err := faker.FakeObject(&l); err != nil {


		t.Fatal(err)




	}


	sl := types.ComplianceByConfigRule{}




	if err := faker.FakeObject(&sl); err != nil {
		t.Fatal(err)


	}
	m.EXPECT().DescribeConfigRules(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.DescribeConfigRulesOutput{
			ConfigRules: []types.ConfigRule{l},




		}, nil)




	m.EXPECT().DescribeComplianceByConfigRule(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.DescribeComplianceByConfigRuleOutput{




			ComplianceByConfigRules: []types.ComplianceByConfigRule{sl},




		}, nil)
	return aws_client.AwsServices{
		Configservice: m,




	}




}





func TestConfigRules(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsConfigConfigRulesGenerator{}), buildConfigRules, aws_client.TestOptions{})
}
