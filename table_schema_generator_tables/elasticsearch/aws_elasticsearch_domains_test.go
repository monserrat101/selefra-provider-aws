



package elasticsearch









import (


	"testing"





	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"




	"github.com/golang/mock/gomock"




	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)









func buildElasticSearchDomains(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockElasticsearchserviceClient(ctrl)







	var info types.DomainInfo




	if err := faker.FakeObject(&info); err != nil {




		t.Fatal(err)


	}




	m.EXPECT().ListDomainNames(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&elasticsearchservice.ListDomainNamesOutput{




			DomainNames: []types.DomainInfo{info},




		},




		nil,
	)

	var ds types.ElasticsearchDomainStatus


	if err := faker.FakeObject(&ds); err != nil {




		t.Fatal(err)


	}
	m.EXPECT().DescribeElasticsearchDomain(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&elasticsearchservice.DescribeElasticsearchDomainOutput{


			DomainStatus: &ds,




		},
		nil,
	)

	var principal types.AuthorizedPrincipal
	if err := faker.FakeObject(&principal); err != nil {


		t.Fatal(err)




	}




	m.EXPECT().ListVpcEndpointAccess(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticsearchservice.ListVpcEndpointAccessOutput{
			AuthorizedPrincipalList: []types.AuthorizedPrincipal{principal},


		},
		nil,
	)



	var tags elasticsearchservice.ListTagsOutput
	if err := faker.FakeObject(&tags); err != nil {




		t.Fatal(err)




	}


	m.EXPECT().ListTags(gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)



	return aws_client.AwsServices{Elasticsearchservice: m}


}





func TestElasticSearchDomains(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticsearchDomainsGenerator{}), buildElasticSearchDomains, aws_client.TestOptions{})




}




