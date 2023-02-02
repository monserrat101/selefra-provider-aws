package apprunner





import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/apprunner"




	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)





func buildObservabilityConfiguration(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApprunnerClient(ctrl)


	s := types.ObservabilityConfiguration{}


	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}



	m.EXPECT().ListObservabilityConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&apprunner.ListObservabilityConfigurationsOutput{




			ObservabilityConfigurationSummaryList: []types.ObservabilityConfigurationSummary{




				{ObservabilityConfigurationArn: s.ObservabilityConfigurationArn},


			},


		}, nil)







	m.EXPECT().DescribeObservabilityConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apprunner.DescribeObservabilityConfigurationOutput{
			ObservabilityConfiguration: &s,
		}, nil)




	tags := types.Tag{}


	err = faker.FakeObject(&tags)


	if err != nil {
		t.Fatal(err)


	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&apprunner.ListTagsForResourceOutput{Tags: []types.Tag{tags}}, nil)


	return aws_client.AwsServices{


		Apprunner: m,




	}
}

func TestObservabilityConfiguration(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApprunnerObservabilityConfigurationsGenerator{}), buildObservabilityConfiguration, aws_client.TestOptions{})




}




