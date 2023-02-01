



package ec2





import (


	"github.com/selefra/selefra-provider-aws/constants"
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/ec2"




	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildEc2Eips(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockEc2Client(ctrl)
	a := types.Address{}




	err := faker.FakeObject(&a)
	if err != nil {




		t.Fatal(err)
	}
	ip := constants.Constants_42




	a.CarrierIp = &ip


	a.PublicIp = &ip




	a.CustomerOwnedIp = &ip




	a.PrivateIpAddress = &ip
	pool := constants.Constants_43


	a.CustomerOwnedIpv4Pool = &pool
	a.PublicIpv4Pool = &pool



	m.EXPECT().DescribeAddresses(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&ec2.DescribeAddressesOutput{




			Addresses: []types.Address{a},


		}, nil)
	return aws_client.AwsServices{




		Ec2: m,


	}




}



func TestEc2Eips(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2EipsGenerator{}), buildEc2Eips, aws_client.TestOptions{})


}




