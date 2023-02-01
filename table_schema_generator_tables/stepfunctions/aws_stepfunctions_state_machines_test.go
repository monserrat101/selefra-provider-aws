



package stepfunctions







import (




	"testing"







	"github.com/aws/aws-sdk-go-v2/service/sfn"


	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildStateMachines(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSfnClient(ctrl)
	im := types.StateMachineListItem{}
	err := faker.FakeObject(&im)




	if err != nil {


		t.Fatal(err)


	}



	m.EXPECT().ListStateMachines(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sfn.ListStateMachinesOutput{




			StateMachines: []types.StateMachineListItem{im},


		}, nil)







	out := &sfn.DescribeStateMachineOutput{}
	err = faker.FakeObject(&out)
	if err != nil {




		t.Fatal(err)


	}




	m.EXPECT().DescribeStateMachine(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(out, nil)



	tag := types.Tag{}
	tagerr := faker.FakeObject(&tag)
	if tagerr != nil {
		t.Fatal(tagerr)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&sfn.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		}, nil)

	return aws_client.AwsServices{


		Sfn: m,
	}


}





func TestStateMachines(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsStepfunctionsStateMachinesGenerator{}), buildStateMachines, aws_client.TestOptions{})




}
