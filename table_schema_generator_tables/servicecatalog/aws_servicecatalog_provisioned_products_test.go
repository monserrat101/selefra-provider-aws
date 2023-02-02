

package servicecatalog



import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"


	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)









func buildProvisionedProducts(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mk := mocks.NewMockServicecatalogClient(ctrl)



	o := servicecatalog.SearchProvisionedProductsOutput{}


	if err := faker.FakeObject(&o); err != nil {




		t.Fatal(err)
	}
	o.NextPageToken = nil







	mk.EXPECT().SearchProvisionedProducts(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&o,


		nil,




	)







	return aws_client.AwsServices{
		Servicecatalog: mk,
	}


}





func TestProvisionedProducts(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsServicecatalogProvisionedProductsGenerator{}), buildProvisionedProducts, aws_client.TestOptions{})


}
