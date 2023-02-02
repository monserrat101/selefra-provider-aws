

package dms







import (
	"github.com/selefra/selefra-provider-aws/constants"


	"testing"





	"github.com/aws/aws-sdk-go-v2/aws"


	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildDmsReplicationInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockDatabasemigrationserviceClient(ctrl)


	l := types.ReplicationInstance{}
	if err := faker.FakeObject(&l); err != nil {




		t.Fatal(err)
	}




	l.ReplicationInstancePrivateIpAddress = aws.String(constants.Constants_38)
	l.ReplicationInstancePrivateIpAddresses = []string{constants.Constants_39}


	l.ReplicationInstancePublicIpAddress = aws.String(constants.Constants_40)




	l.ReplicationInstancePublicIpAddresses = []string{constants.Constants_41}
	m.EXPECT().DescribeReplicationInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&databasemigrationservice.DescribeReplicationInstancesOutput{
			ReplicationInstances: []types.ReplicationInstance{l},
		}, nil)
	lt := types.Tag{}


	if err := faker.FakeObject(&lt); err != nil {


		t.Fatal(err)


	}
	lt.ResourceArn = l.ReplicationInstanceArn
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&databasemigrationservice.ListTagsForResourceOutput{
			TagList: []types.Tag{lt},
		}, nil)
	return aws_client.AwsServices{
		Databasemigrationservice: m,




	}




}

func TestDmsReplicationInstances(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDmsReplicationInstancesGenerator{}), buildDmsReplicationInstances, aws_client.TestOptions{})




}
