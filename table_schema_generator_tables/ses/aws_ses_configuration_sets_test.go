

package ses





import (
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/sesv2"


	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)



func buildConfigurationSets(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	sesClient := mocks.NewMockSesv2Client(ctrl)









	cs := sesv2.GetConfigurationSetOutput{}
	if err := faker.FakeObject(&cs); err != nil {


		t.Fatal(err)
	}





	ed := types.EventDestination{}
	if err := faker.FakeObject(&ed); err != nil {




		t.Fatal(err)




	}





	sesClient.EXPECT().ListConfigurationSets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sesv2.ListConfigurationSetsOutput{ConfigurationSets: []string{*cs.ConfigurationSetName}},
		nil,


	)


	sesClient.EXPECT().GetConfigurationSet(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&cs,
		nil,


	)









	sesClient.EXPECT().GetConfigurationSetEventDestinations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&sesv2.GetConfigurationSetEventDestinationsOutput{EventDestinations: []types.EventDestination{ed}},


		nil,




	)





	return aws_client.AwsServices{




		Sesv2: sesClient,
	}




}



func TestConfigurationSets(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSesConfigurationSetsGenerator{}), buildConfigurationSets, aws_client.TestOptions{})
}
