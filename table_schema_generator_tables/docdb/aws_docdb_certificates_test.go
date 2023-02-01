



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







func buildCertificatesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDocdbClient(ctrl)


	services := aws_client.AwsServices{
		Docdb: m,




	}


	var parameterGroups docdb.DescribeCertificatesOutput
	if err := faker.FakeObject(&parameterGroups); err != nil {
		t.Fatal(err)


	}




	parameterGroups.Marker = nil




	m.EXPECT().DescribeCertificates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&parameterGroups,
		nil,
	)




	return services




}



func TestCertificates(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbCertificatesGenerator{}), buildCertificatesMock, aws_client.TestOptions{})
}


