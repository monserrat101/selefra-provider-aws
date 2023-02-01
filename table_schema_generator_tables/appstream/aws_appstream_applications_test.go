

package appstream





import (


	"testing"







	"github.com/aws/aws-sdk-go-v2/service/appstream"


	"github.com/aws/aws-sdk-go-v2/service/appstream/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)



func buildAppstreamApplicationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockAppstreamClient(ctrl)




	application := types.Application{}
	err := faker.FakeObject(&application)
	if err != nil {


		t.Fatal(err)


	}





	applicationFleetAssocition := types.ApplicationFleetAssociation{}
	if faker.FakeObject(&applicationFleetAssocition) != nil {
		t.Fatal(err)




	}







	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&appstream.DescribeApplicationsOutput{


			Applications: []types.Application{application},




		}, nil)



	m.EXPECT().DescribeApplicationFleetAssociations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&appstream.DescribeApplicationFleetAssociationsOutput{




			ApplicationFleetAssociations: []types.ApplicationFleetAssociation{applicationFleetAssocition},


		}, nil)









	return aws_client.AwsServices{


		Appstream: m,




	}
}



func TestAppstreamApplications(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAppstreamApplicationsGenerator{}), buildAppstreamApplicationsMock, aws_client.TestOptions{})




}




