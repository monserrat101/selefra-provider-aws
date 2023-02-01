package docdb





import (




	"testing"









	"github.com/aws/aws-sdk-go-v2/service/docdb"


	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)







func buildClustersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockDocdbClient(ctrl)
	services := aws_client.AwsServices{
		Docdb: m,




	}




	var clusters docdb.DescribeDBClustersOutput




	if err := faker.FakeObject(&clusters); err != nil {




		t.Fatal(err)


	}




	clusters.Marker = nil
	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&clusters,
		nil,


	)





	var clusterSnapshots docdb.DescribeDBClusterSnapshotsOutput
	if err := faker.FakeObject(&clusterSnapshots); err != nil {
		t.Fatal(err)


	}
	clusterSnapshots.Marker = nil




	m.EXPECT().DescribeDBClusterSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&clusterSnapshots,


		nil,


	)

	var snapshotAttributes docdb.DescribeDBClusterSnapshotAttributesOutput


	if err := faker.FakeObject(&snapshotAttributes); err != nil {


		t.Fatal(err)




	}
	m.EXPECT().DescribeDBClusterSnapshotAttributes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&snapshotAttributes,
		nil,


	)

	var ev docdb.DescribeDBInstancesOutput


	if err := faker.FakeObject(&ev); err != nil {




		t.Fatal(err)
	}




	ev.Marker = nil




	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&ev,


		nil,
	)







	var tags docdb.ListTagsForResourceOutput




	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)




	}




	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&tags,
		nil,


	).AnyTimes()





	return services


}







func TestClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbClustersGenerator{}), buildClustersMock, aws_client.TestOptions{})
}
