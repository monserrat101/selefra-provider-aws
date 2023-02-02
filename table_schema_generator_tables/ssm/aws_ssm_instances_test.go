package ssm

import (


	"github.com/selefra/selefra-provider-aws/constants"




	"testing"





	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)









func buildSSMInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	mock := mocks.NewMockSsmClient(ctrl)



	var i types.InstanceInformation


	if err := faker.FakeObject(&i); err != nil {


		t.Fatal(err)


	}




	i.IPAddress = aws.String(constants.Constants_65)


	mock.EXPECT().DescribeInstanceInformation(


		gomock.Any(),




		&ssm.DescribeInstanceInformationInput{},




		gomock.Any(),


	).AnyTimes().Return(


		&ssm.DescribeInstanceInformationOutput{InstanceInformationList: []types.InstanceInformation{i}},
		nil,




	)





	var c types.ComplianceItem
	if err := faker.FakeObject(&c); err != nil {


		t.Fatal(err)




	}




	mock.EXPECT().ListComplianceItems(gomock.Any(),
		&ssm.ListComplianceItemsInput{ResourceIds: []string{*i.InstanceId}},
		gomock.Any(),




	).AnyTimes().Return(




		&ssm.ListComplianceItemsOutput{ComplianceItems: []types.ComplianceItem{c}},




		nil,
	)









	var p types.PatchComplianceData




	if err := faker.FakeObject(&p); err != nil {


		t.Fatal(err)




	}


	mock.EXPECT().DescribeInstancePatches(gomock.Any(),
		&ssm.DescribeInstancePatchesInput{InstanceId: i.InstanceId},


		gomock.Any(),


	).AnyTimes().Return(




		&ssm.DescribeInstancePatchesOutput{Patches: []types.PatchComplianceData{p}},
		nil,


	)





	return aws_client.AwsServices{Ssm: mock}




}



func TestSSMInstances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSsmInstancesGenerator{}), buildSSMInstances, aws_client.TestOptions{})
}




