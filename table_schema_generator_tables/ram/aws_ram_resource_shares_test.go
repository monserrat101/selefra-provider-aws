



package ram





import (


	"github.com/selefra/selefra-provider-aws/constants"


	"fmt"
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/ram"


	"github.com/aws/aws-sdk-go-v2/service/ram/types"


	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)





func buildRamResourceSharesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockRamClient(ctrl)
	object := types.ResourceShare{}




	err := faker.FakeObject(&object)


	if err != nil {


		t.Fatal(err)
	}



	m.EXPECT().GetResourceShares(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ram.GetResourceSharesOutput{ResourceShares: []types.ResourceShare{object}}, nil).MinTimes(1)



	summary := types.ResourceSharePermissionSummary{}


	err = faker.FakeObject(&summary)
	if err != nil {


		t.Fatal(err)




	}







	var version int32


	err = faker.FakeObject(&version)




	if err != nil {
		t.Fatal(err)




	}


	verStr := fmt.Sprint(version)




	summary.Version = &verStr



	m.EXPECT().ListResourceSharePermissions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ram.ListResourceSharePermissionsOutput{Permissions: []types.ResourceSharePermissionSummary{summary}}, nil).MinTimes(1)





	detail := constants.Constants_47







	m.EXPECT().GetPermission(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ram.GetPermissionOutput{Permission: &types.ResourceSharePermissionDetail{Permission: &detail}}, nil).MinTimes(1)





	return aws_client.AwsServices{RAM: m}


}





func TestRamResourceShares(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRamResourceSharesGenerator{}), buildRamResourceSharesMock, aws_client.TestOptions{})
}


