



package rds



import (




	"github.com/selefra/selefra-provider-aws/constants"


	"testing"





	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)





func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockRdsClient(ctrl)
	services := aws_client.AwsServices{


		Rds: m,




	}


	var ev rds.DescribeDBEngineVersionsOutput
	if err := faker.FakeObject(&ev); err != nil {




		t.Fatal(err)




	}


	ev.Marker = nil




	var aurora types.DBEngineVersion
	if err := faker.FakeObject(&aurora); err != nil {




		t.Fatal(err)


	}




	aurora.DBParameterGroupFamily = aws.String(constants.Auroramysql)
	ev.DBEngineVersions = append(ev.DBEngineVersions, aurora)
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&ev,
		nil,


	)

	var parameters rds.DescribeEngineDefaultClusterParametersOutput




	if err := faker.FakeObject(&parameters); err != nil {




		t.Fatal(err)
	}


	m.EXPECT().DescribeEngineDefaultClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&parameters,




		nil,




	)





	return services




}









func TestEngineVersions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsEngineVersionsGenerator{}), buildEngineVersionsMock, aws_client.TestOptions{})
}


