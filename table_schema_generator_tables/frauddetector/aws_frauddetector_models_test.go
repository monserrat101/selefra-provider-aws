



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









func buildModelVersions(t *testing.T, client *mocks.MockFrauddetectorClient) {




	data := types.ModelVersionDetail{}




	err := faker.FakeObject(&data)
	if err != nil {


		t.Fatal(err)


	}

	client.EXPECT().DescribeModelVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.DescribeModelVersionsOutput{ModelVersionDetails: []types.ModelVersionDetail{data}}, nil,




	)




}



func buildModels(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	fdClient := mocks.NewMockFrauddetectorClient(ctrl)







	data := types.Model{}




	err := faker.FakeObject(&data)
	if err != nil {


		t.Fatal(err)
	}





	fdClient.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&frauddetector.GetModelsOutput{Models: []types.Model{data}}, nil,
	)

	buildModelVersions(t, fdClient)





	return aws_client.AwsServices{


		Frauddetector: fdClient,


	}




}



func TestModels(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFrauddetectorModelsGenerator{}), buildModels, aws_client.TestOptions{})




}
