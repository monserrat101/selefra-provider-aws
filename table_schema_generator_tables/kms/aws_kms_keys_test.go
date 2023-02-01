

package kms





import (




	"github.com/selefra/selefra-provider-aws/constants"




	"testing"



	"github.com/aws/aws-sdk-go-v2/aws"


	"github.com/aws/aws-sdk-go-v2/service/kms"


	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)









func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockKmsClient(ctrl)







	keyListEntry := types.KeyListEntry{}




	if err := faker.FakeObject(&keyListEntry); err != nil {




		t.Fatal(err)


	}





	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&kms.ListKeysOutput{Keys: []types.KeyListEntry{keyListEntry}}, nil)







	tags := kms.ListResourceTagsOutput{}
	if err := faker.FakeObject(&tags); err != nil {




		t.Fatal(err)




	}
	tags.NextMarker = nil




	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)









	key := kms.DescribeKeyOutput{}
	if err := faker.FakeObject(&key); err != nil {
		t.Fatal(err)


	}




	m.EXPECT().DescribeKey(gomock.Any(), &kms.DescribeKeyInput{KeyId: keyListEntry.KeyId}, gomock.Any()).AnyTimes().Return(&key, nil)





	rotation := kms.GetKeyRotationStatusOutput{}


	if err := faker.FakeObject(&rotation); err != nil {


		t.Fatal(err)


	}
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&rotation, nil)



	g := kms.ListGrantsOutput{}




	if err := faker.FakeObject(&g); err != nil {




		t.Fatal(err)




	}




	g.NextMarker = nil
	m.EXPECT().ListGrants(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&g, nil)







	pj := `{"data":["data"]}`




	m.EXPECT().GetKeyPolicy(gomock.Any(), &kms.GetKeyPolicyInput{
		KeyId:		keyListEntry.KeyId,




		PolicyName:	aws.String(constants.Default),
	}, gomock.Any()).AnyTimes().Return(&kms.GetKeyPolicyOutput{Policy: &pj}, nil)



	return aws_client.AwsServices{
		Kms: m,




	}
}

func TestKmsKeys(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsKmsKeysGenerator{}), buildKmsKeys, aws_client.TestOptions{})
}
