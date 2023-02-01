

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







func buildSchedulerSchedulesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockSchedulerClient(ctrl)




	object := types.ScheduleSummary{}


	err := faker.FakeObject(&object)


	if err != nil {


		t.Fatal(err)


	}







	m.EXPECT().ListSchedules(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&scheduler.ListSchedulesOutput{




			Schedules: []types.ScheduleSummary{object},




		}, nil)
	object2 := scheduler.GetScheduleOutput{}


	err = faker.FakeObject(&object2)
	if err != nil {




		t.Fatal(err)


	}





	m.EXPECT().GetSchedule(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&object2, nil)







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



func TestSchedulesMock(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSchedulerSchedulesGenerator{}), buildSchedulerSchedulesMock, aws_client.TestOptions{})
}




