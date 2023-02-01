

package cloudtrail



import (


	"github.com/selefra/selefra-provider-aws/constants"
	"testing"







	"github.com/aws/aws-sdk-go-v2/aws"


	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)





func buildCloudtrailTrailsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockCloudtrailClient(ctrl)


	services := aws_client.AwsServices{




		Cloudtrail: m,




	}


	trail := types.Trail{}




	err := faker.FakeObject(&trail)
	if err != nil {
		t.Fatal(err)




	}





	trail.TrailARN = aws.String(constants.ArnawscloudtraileucentraltestAccounttrailtesttrail)


	trail.CloudWatchLogsLogGroupArn = aws.String(constants.Arnawslogseucentralloggrouptestgroup)





	trailStatus := cloudtrail.GetTrailStatusOutput{}
	err = faker.FakeObject(&trailStatus)


	if err != nil {


		t.Fatal(err)




	}
	eventSelector := types.EventSelector{}
	err = faker.FakeObject(&eventSelector)


	if err != nil {


		t.Fatal(err)




	}




	m.EXPECT().DescribeTrails(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&cloudtrail.DescribeTrailsOutput{




			TrailList: []types.Trail{trail},
		},




		nil,
	)


	m.EXPECT().GetTrailStatus(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&trailStatus,




		nil,




	)
	m.EXPECT().GetEventSelectors(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&cloudtrail.GetEventSelectorsOutput{
			EventSelectors: []types.EventSelector{eventSelector},


		},
		nil,


	)




	tags := cloudtrail.ListTagsOutput{}




	err = faker.FakeObject(&tags)
	if err != nil {


		t.Fatal(err)
	}


	tags.ResourceTagList[0].ResourceId = trail.TrailARN


	tags.NextToken = nil
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)



	return services


}





func TestCloudtrailTrails(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudtrailTrailsGenerator{}), buildCloudtrailTrailsMock, aws_client.TestOptions{})


}




