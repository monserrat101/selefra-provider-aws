



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





func buildBatchImports(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)





	data := types.BatchImport{}




	err := faker.FakeObject(&data)


	if err != nil {




		t.Fatal(err)




	}





	fdClient.EXPECT().GetBatchImportJobs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&frauddetector.GetBatchImportJobsOutput{BatchImports: []types.BatchImport{data}}, nil,


	)



	return aws_client.AwsServices{


		Frauddetector: fdClient,
	}
}

func TestBatchImports(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFrauddetectorBatchImportsGenerator{}), buildBatchImports, aws_client.TestOptions{})
}




