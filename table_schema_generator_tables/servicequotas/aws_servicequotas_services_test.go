

package servicequotas





import (


	"testing"







	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)





func buildServices(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockServicequotasClient(ctrl)







	services := servicequotas.ListServicesOutput{}


	err := faker.FakeObject(&services)




	if err != nil {




		t.Fatal(err)


	}
	services.NextToken = nil
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&services, nil)





	quotas := servicequotas.ListServiceQuotasOutput{}




	err = faker.FakeObject(&quotas)
	if err != nil {


		t.Fatal(err)


	}



	quotas.NextToken = nil




	m.EXPECT().ListServiceQuotas(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&quotas, nil).AnyTimes()



	return aws_client.AwsServices{




		Servicequotas: m,
	}
}



func TestQuotas(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsServicequotasServicesGenerator{}), buildServices, aws_client.TestOptions{})


}


