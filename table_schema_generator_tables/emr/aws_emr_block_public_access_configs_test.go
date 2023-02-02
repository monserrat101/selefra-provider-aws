package emr



import (


	"github.com/selefra/selefra-provider-aws/constants"


	"testing"
	"time"









	"github.com/aws/aws-sdk-go-v2/aws"




	"github.com/aws/aws-sdk-go-v2/service/emr"


	"github.com/aws/aws-sdk-go-v2/service/emr/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildEMRClient(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	emrmock := mocks.NewMockEmrClient(ctrl)







	out := &emr.GetBlockPublicAccessConfigurationOutput{
		BlockPublicAccessConfiguration: &types.BlockPublicAccessConfiguration{


			Classification:	aws.String(constants.Classification),
			Configurations:	[]types.Configuration{},


			PermittedPublicSecurityGroupRuleRanges: []types.PortRange{


				{
					MinRange:	aws.Int32(1024),


					MaxRange:	aws.Int32(2048),
				},
			},


			Properties: map[string]string{


				"key": "value",




			},




		},


		BlockPublicAccessConfigurationMetadata: &types.BlockPublicAccessConfigurationMetadata{




			CreatedByArn:		aws.String(constants.Justsomevalue),




			CreationDateTime:	aws.Time(time.Now()),


		},




	}
	emrmock.EXPECT().GetBlockPublicAccessConfiguration(
		gomock.Any(),




		&emr.GetBlockPublicAccessConfigurationInput{},


		gomock.Any(),
	).AnyTimes().Return(out, nil)
	return aws_client.AwsServices{Emr: emrmock}


}





func TestEMRBlockPublicAccessConfigs(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEmrBlockPublicAccessConfigsGenerator{}), buildEMRClient, aws_client.TestOptions{})




}


