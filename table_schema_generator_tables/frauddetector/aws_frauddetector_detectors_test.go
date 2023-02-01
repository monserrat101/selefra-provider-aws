



package frauddetector



import (
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"




	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)



func buildRules(t *testing.T, client *mocks.MockFrauddetectorClient) {




	data := types.RuleDetail{}
	err := faker.FakeObject(&data)


	if err != nil {




		t.Fatal(err)


	}





	client.EXPECT().GetRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetRulesOutput{RuleDetails: []types.RuleDetail{data}}, nil,




	)


}
func addTagsCall(t *testing.T, client *mocks.MockFrauddetectorClient) {
	var data []types.Tag
	err := faker.FakeObject(&data)
	if err != nil {


		t.Fatal(err)
	}





	client.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.ListTagsForResourceOutput{Tags: data}, nil,




	)




}







func buildDetectors(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)





	data := types.Detector{}
	err := faker.FakeObject(&data)


	if err != nil {




		t.Fatal(err)




	}







	fdClient.EXPECT().GetDetectors(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&frauddetector.GetDetectorsOutput{Detectors: []types.Detector{data}}, nil,


	)







	buildRules(t, fdClient)


	addTagsCall(t, fdClient)



	return aws_client.AwsServices{


		Frauddetector: fdClient,




	}




}





func TestDetectors(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFrauddetectorDetectorsGenerator{}), buildDetectors, aws_client.TestOptions{})
}




