

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





func buildEventSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockDocdbClient(ctrl)
	services := aws_client.AwsServices{


		Docdb: m,




	}
	var ec docdb.DescribeEventSubscriptionsOutput
	if err := faker.FakeObject(&ec); err != nil {
		t.Fatal(err)
	}




	ec.Marker = nil


	m.EXPECT().DescribeEventSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&ec,
		nil,




	)

	return services
}







func TestEventSubscriptions(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbEventSubscriptionsGenerator{}), buildEventSubscriptionsMock, aws_client.TestOptions{})




}




