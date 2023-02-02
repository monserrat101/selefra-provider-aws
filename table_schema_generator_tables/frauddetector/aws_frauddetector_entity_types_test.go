



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



func buildEntityTypes(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	fdClient := mocks.NewMockFrauddetectorClient(ctrl)



	data := types.EntityType{}


	err := faker.FakeObject(&data)
	if err != nil {


		t.Fatal(err)




	}



	fdClient.EXPECT().GetEntityTypes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&frauddetector.GetEntityTypesOutput{EntityTypes: []types.EntityType{data}}, nil,




	)

	addTagsCall(t, fdClient)







	return aws_client.AwsServices{
		Frauddetector: fdClient,


	}




}







func TestEntityTypes(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFrauddetectorEntityTypesGenerator{}), buildEntityTypes, aws_client.TestOptions{})




}




