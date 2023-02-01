



package elastictranscoder







import (


	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"


	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildElastictranscoderPresetsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockElastictranscoderClient(ctrl)




	object := types.Preset{}




	err := faker.FakeObject(&object)


	if err != nil {


		t.Fatal(err)
	}



	m.EXPECT().ListPresets(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elastictranscoder.ListPresetsOutput{Presets: []types.Preset{object}},


		nil,


	)









	return aws_client.AwsServices{
		Elastictranscoder: m,
	}
}





func TestElastictranscoderPresets(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElastictranscoderPresetsGenerator{}), buildElastictranscoderPresetsMock, aws_client.TestOptions{})
}
