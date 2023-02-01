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









func buildApprunnerAutoScalingConfigurationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockApprunnerClient(ctrl)




	as := types.AutoScalingConfiguration{}
	err := faker.FakeObject(&as)
	if err != nil {


		t.Fatal(err)
	}





	m.EXPECT().ListAutoScalingConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apprunner.ListAutoScalingConfigurationsOutput{




			AutoScalingConfigurationSummaryList: []types.AutoScalingConfigurationSummary{
				{AutoScalingConfigurationArn: as.AutoScalingConfigurationArn},
			},




		}, nil)





	m.EXPECT().DescribeAutoScalingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apprunner.DescribeAutoScalingConfigurationOutput{


			AutoScalingConfiguration: &as,


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



func TestApprunnerAutoScalingConfigurations(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApprunnerAutoScalingConfigurationsGenerator{}), buildApprunnerAutoScalingConfigurationsMock, aws_client.TestOptions{})
}


