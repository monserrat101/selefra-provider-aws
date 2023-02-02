



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



func buildApprunnerVpcIngressConnectionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockApprunnerClient(ctrl)


	vc := types.VpcIngressConnection{}




	err := faker.FakeObject(&vc)


	if err != nil {
		t.Fatal(err)


	}









	m.EXPECT().ListVpcIngressConnections(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&apprunner.ListVpcIngressConnectionsOutput{




			VpcIngressConnectionSummaryList: []types.VpcIngressConnectionSummary{{ServiceArn: vc.ServiceArn, VpcIngressConnectionArn: vc.VpcIngressConnectionArn}},




		}, nil)
	m.EXPECT().DescribeVpcIngressConnection(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apprunner.DescribeVpcIngressConnectionOutput{VpcIngressConnection: &vc}, nil)





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





func TestApprunnerVpcIngressConnector(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApprunnerVpcIngressConnectionsGenerator{}), buildApprunnerVpcIngressConnectionsMock, aws_client.TestOptions{})


}


