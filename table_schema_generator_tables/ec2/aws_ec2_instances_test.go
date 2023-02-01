package ec2



import (
	"github.com/selefra/selefra-provider-aws/constants"
	"testing"



	"github.com/aws/aws-sdk-go-v2/aws"




	"github.com/aws/aws-sdk-go-v2/service/ec2"


	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)

func buildEc2Instances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockEc2Client(ctrl)


	l := types.Reservation{}




	err := faker.FakeObject(&l)


	if err != nil {
		t.Fatal(err)


	}
	l.Instances[0].StateTransitionReason = aws.String(constants.UserinitiatedGMT)




	creationDate := constants.T


	l.Instances[0].ElasticGpuAssociations[0].ElasticGpuAssociationTime = &creationDate
	nextToken := constants.Test

	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).AnyTimes().Return(


		&ec2.DescribeInstancesOutput{


			Reservations:	[]types.Reservation{},




			NextToken:	&nextToken,




		}, nil)


	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).AnyTimes().Return(


		&ec2.DescribeInstancesOutput{




			Reservations:	[]types.Reservation{l},




			NextToken:	nil,




		}, nil)


	return aws_client.AwsServices{
		Ec2: m,




	}


}







func TestEc2Instances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2InstancesGenerator{}), buildEc2Instances, aws_client.TestOptions{})




}




