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







func buildEventsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockDocdbClient(ctrl)
	services := aws_client.AwsServices{




		Docdb: m,


	}
	var e docdb.DescribeEventsOutput
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)


	}
	e.Marker = nil


	m.EXPECT().DescribeEvents(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&e,
		nil,
	)



	return services




}





func TestEvents(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbEventsGenerator{}), buildEventsMock, aws_client.TestOptions{})


}




