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





func buildApprunnerVpcConnectorsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockApprunnerClient(ctrl)


	vc := types.VpcConnector{}


	err := faker.FakeObject(&vc)




	if err != nil {




		t.Fatal(err)




	}



	m.EXPECT().ListVpcConnectors(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&apprunner.ListVpcConnectorsOutput{


			VpcConnectors: []types.VpcConnector{vc},
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

func TestApprunnerVpcConnector(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApprunnerVpcConnectorsGenerator{}), buildApprunnerVpcConnectorsMock, aws_client.TestOptions{})
}




