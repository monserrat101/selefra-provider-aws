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



func buildPendingMaintenanceActionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockDocdbClient(ctrl)
	services := aws_client.AwsServices{




		Docdb: m,




	}
	var ev docdb.DescribePendingMaintenanceActionsOutput
	if err := faker.FakeObject(&ev); err != nil {


		t.Fatal(err)
	}
	ev.Marker = nil




	m.EXPECT().DescribePendingMaintenanceActions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ev,


		nil,


	)




	return services
}



func TestPendingMaintenanceActions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbPendingMaintenanceActionsGenerator{}), buildPendingMaintenanceActionsMock, aws_client.TestOptions{})


}
