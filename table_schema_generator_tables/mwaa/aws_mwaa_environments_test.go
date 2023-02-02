



package mwaa





import (




	"testing"



	"github.com/aws/aws-sdk-go-v2/service/mwaa"


	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildMwaaEnvironments(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockMwaaClient(ctrl)


	g := types.Environment{}




	err := faker.FakeObject(&g)


	if err != nil {




		t.Fatal(err)


	}



	m.EXPECT().ListEnvironments(gomock.Any(), gomock.Any()).AnyTimes().Return(


		&mwaa.ListEnvironmentsOutput{


			Environments: []string{*g.Name},
		}, nil)
	m.EXPECT().GetEnvironment(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&mwaa.GetEnvironmentOutput{




			Environment: &g,
		}, nil)
	return aws_client.AwsServices{




		Mwaa: m,




	}


}

func TestMwaaEnvironments(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsMwaaEnvironmentsGenerator{}), buildMwaaEnvironments, aws_client.TestOptions{})




}




