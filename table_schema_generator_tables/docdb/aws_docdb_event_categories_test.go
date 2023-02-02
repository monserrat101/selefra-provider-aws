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



func buildEventCategoriesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockDocdbClient(ctrl)


	services := aws_client.AwsServices{




		Docdb: m,
	}
	var ec docdb.DescribeEventCategoriesOutput




	if err := faker.FakeObject(&ec); err != nil {


		t.Fatal(err)
	}




	m.EXPECT().DescribeEventCategories(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&ec,


		nil,




	)







	return services


}





func TestEventCategories(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDocdbEventCategoriesGenerator{}), buildEventCategoriesMock, aws_client.TestOptions{})
}
