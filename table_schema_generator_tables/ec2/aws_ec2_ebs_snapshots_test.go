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





func buildEc2EbsSnapshots(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockEc2Client(ctrl)
	s := types.Snapshot{}


	userId := constants.UserId




	sa := types.CreateVolumePermission{


		Group:	constants.Test,




		UserId:	&userId,
	}
	err := faker.FakeObject(&s)


	if err != nil {




		t.Fatal(err)


	}
	m.EXPECT().DescribeSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeSnapshotsOutput{




			Snapshots: []types.Snapshot{s},
		}, nil)




	m.EXPECT().DescribeSnapshotAttribute(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeSnapshotAttributeOutput{


			CreateVolumePermissions: []types.CreateVolumePermission{sa},




		}, nil)
	return aws_client.AwsServices{




		Ec2: m,


	}




}







func TestEc2EbsSnapshots(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2EbsSnapshotsGenerator{}), buildEc2EbsSnapshots, aws_client.TestOptions{})


}




