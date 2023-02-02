



package ram







import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/ram"




	"github.com/aws/aws-sdk-go-v2/service/ram/types"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)







func buildRamResourcesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRamClient(ctrl)




	object := types.Resource{}




	err := faker.FakeObject(&object)




	if err != nil {
		t.Fatal(err)


	}





	m.EXPECT().ListResources(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ram.ListResourcesOutput{
			Resources: []types.Resource{object},


		}, nil).MinTimes(1)





	return aws_client.AwsServices{
		RAM: m,
	}


}





func TestRamResources(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRamResourcesGenerator{}), buildRamResourcesMock, aws_client.TestOptions{})


}




