



package elasticbeanstalk

import (




	"testing"



	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"




	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)







func buildElasticbeanstalkApplications(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockElasticbeanstalkClient(ctrl)









	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeObject(&la)
	if err != nil {
		t.Fatal(err)




	}









	tag := elasticbeanstalkTypes.Tag{}
	err = faker.FakeObject(&tag)




	if err != nil {


		t.Fatal(err)


	}

	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&elasticbeanstalk.DescribeApplicationsOutput{




			Applications: []elasticbeanstalkTypes.ApplicationDescription{la},




		}, nil)









	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&elasticbeanstalk.ListTagsForResourceOutput{




			ResourceTags:	[]elasticbeanstalkTypes.Tag{tag},




			ResourceArn:	la.ApplicationArn,


		}, nil)



	return aws_client.AwsServices{


		Elasticbeanstalk: m,


	}
}







func TestElasticbeanstalkApplications(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkApplicationsGenerator{}), buildElasticbeanstalkApplications, aws_client.TestOptions{})
}


