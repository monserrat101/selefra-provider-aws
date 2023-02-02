package servicecatalog







import (




	"testing"



	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"


	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"




	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)



func buildPortfolios(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mk := mocks.NewMockServicecatalogClient(ctrl)


	ma := mocks.NewMockServicecatalogappregistryClient(ctrl)





	o := servicecatalog.ListPortfoliosOutput{}
	if err := faker.FakeObject(&o); err != nil {




		t.Fatal(err)
	}


	o.NextPageToken = nil







	mk.EXPECT().ListPortfolios(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&o,
		nil,




	)



	to := servicecatalogappregistry.ListTagsForResourceOutput{}




	if err := faker.FakeObject(&to); err != nil {




		t.Fatal(err)
	}







	ma.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&to,


		nil,


	)



	return aws_client.AwsServices{


		Servicecatalog:			mk,




		Servicecatalogappregistry:	ma,
	}


}







func TestPortfolios(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsServicecatalogPortfoliosGenerator{}), buildPortfolios, aws_client.TestOptions{})
}


