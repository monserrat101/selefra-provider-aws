

package xray

import (
	"testing"







	"github.com/aws/aws-sdk-go-v2/service/xray"


	"github.com/aws/aws-sdk-go-v2/service/xray/types"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)



func buildResourcePolicies(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	mock := mocks.NewMockXrayClient(ctrl)



	var pols types.ResourcePolicy




	if err := faker.FakeObject(&pols); err != nil {


		t.Fatal(err)


	}





	mock.EXPECT().ListResourcePolicies(
		gomock.Any(),




		&xray.ListResourcePoliciesInput{},


		gomock.Any(),


	).AnyTimes().Return(
		&xray.ListResourcePoliciesOutput{


			ResourcePolicies: []types.ResourcePolicy{


				pols,
			},


		},
		nil,
	)









	return aws_client.AwsServices{Xray: mock}
}





func TestResourcePolicies(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsXrayResourcePoliciesGenerator{}), buildResourcePolicies, aws_client.TestOptions{})
}


