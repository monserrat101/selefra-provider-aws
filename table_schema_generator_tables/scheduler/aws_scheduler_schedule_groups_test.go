package scheduler







import (


	"testing"







	"github.com/aws/aws-sdk-go-v2/service/scheduler"


	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)



func buildSchedulerScheduleGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockSchedulerClient(ctrl)
	object := types.ScheduleGroupSummary{}


	err := faker.FakeObject(&object)
	if err != nil {




		t.Fatal(err)


	}







	m.EXPECT().ListScheduleGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&scheduler.ListScheduleGroupsOutput{




			ScheduleGroups: []types.ScheduleGroupSummary{object},
		}, nil)



	tagsOutput := scheduler.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagsOutput)




	if err != nil {


		t.Fatal(err)




	}


	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).AnyTimes().Return(&tagsOutput, nil).AnyTimes()


	return aws_client.AwsServices{




		Scheduler: m,
	}




}

func TestSchedulerSchedulerGroups(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSchedulerScheduleGroupsGenerator{}), buildSchedulerScheduleGroupsMock, aws_client.TestOptions{})


}


