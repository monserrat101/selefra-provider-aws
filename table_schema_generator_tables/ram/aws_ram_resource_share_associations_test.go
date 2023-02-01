

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

func buildRamResourceShareAssociationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockRamClient(ctrl)

	object := types.ResourceShareAssociation{}


	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}





	m.EXPECT().GetResourceShareAssociations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ram.GetResourceShareAssociationsOutput{ResourceShareAssociations: []types.ResourceShareAssociation{object}}, nil).MinTimes(1)




	return aws_client.AwsServices{RAM: m}




}









func TestRamResourceShareAssociatedResources(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRamResourceShareAssociationsGenerator{}), buildRamResourceShareAssociationsMock, aws_client.TestOptions{})
}
