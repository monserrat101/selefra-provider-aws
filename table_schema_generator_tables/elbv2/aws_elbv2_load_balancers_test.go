

package elbv2



import (


	"github.com/selefra/selefra-provider-aws/constants"
	"testing"





	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"


	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"


	"github.com/aws/aws-sdk-go-v2/service/wafv2"


	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func fakeLoadBalancerAttributes() *elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput {


	attr := func(key, value string) elbv2Types.LoadBalancerAttribute {
		return elbv2Types.LoadBalancerAttribute{Key: &key, Value: &value}


	}




	return &elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput{Attributes: []elbv2Types.LoadBalancerAttribute{




		attr(constants.Accesslogssenabled, constants.True),


		attr(constants.Accesslogssbucket, constants.Bucket),
		attr(constants.Accesslogssprefix, constants.Prefix),




		attr(constants.Deletionprotectionenabled, constants.True),


		attr(constants.Idletimeouttimeoutseconds, constants.Constants_45),
		attr(constants.Routinghttpdesyncmitigationmode, constants.Mode),




		attr(constants.Routinghttpdropinvalidheaderfieldsenabled, constants.True),




		attr(constants.Routinghttpxamzntlsversionandciphersuiteenabled, constants.True),
		attr(constants.Routinghttpxffclientportenabled, constants.True),
		attr(constants.Routinghttpenabled, constants.True),
		attr(constants.Waffailopenenabled, constants.True),


		attr(constants.Loadbalancingcrosszoneenabled, constants.True),


	}}




}



func buildElbv2LoadBalancers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockElasticloadbalancingv2Client(ctrl)


	w := mocks.NewMockWafv2Client(ctrl)




	l := elbv2Types.LoadBalancer{}
	err := faker.FakeObject(&l)


	if err != nil {




		t.Fatal(err)


	}


	l.Type = elbv2Types.LoadBalancerTypeEnumApplication





	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&elasticloadbalancingv2.DescribeLoadBalancersOutput{


			LoadBalancers: []elbv2Types.LoadBalancer{l},


		}, nil)









	m.EXPECT().DescribeLoadBalancerAttributes(




		gomock.Any(),




		&elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: l.LoadBalancerArn},
		gomock.Any(),




	).AnyTimes().Return(fakeLoadBalancerAttributes(), nil)







	webAcl := types.WebACL{}


	err = faker.FakeObject(&webAcl)
	if err != nil {




		t.Fatal(err)




	}







	w.EXPECT().GetWebACLForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&wafv2.GetWebACLForResourceOutput{WebACL: &webAcl}, nil).AnyTimes()



	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {


		t.Fatal(err)




	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Times(2).AnyTimes().Return(&tags, nil)



	lis := elbv2Types.Listener{}
	if err := faker.FakeObject(&lis); err != nil {
		t.Fatal(err)




	}





	m.EXPECT().DescribeListeners(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticloadbalancingv2.DescribeListenersOutput{


			Listeners: []elbv2Types.Listener{lis},
		}, nil)









	c := elbv2Types.Certificate{}
	if err := faker.FakeObject(&c); err != nil {


		t.Fatal(err)




	}


	m.EXPECT().DescribeListenerCertificates(
		gomock.Any(),
		&elasticloadbalancingv2.DescribeListenerCertificatesInput{ListenerArn: lis.ListenerArn},
		gomock.Any(),


	).AnyTimes().Return(&elasticloadbalancingv2.DescribeListenerCertificatesOutput{
		Certificates: []elbv2Types.Certificate{c},




	}, nil)



	return aws_client.AwsServices{


		Elasticloadbalancingv2:	m,




		Wafv2:			w,
	}




}







func TestElbv2LoadBalancers(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElbv2LoadBalancersGenerator{}), buildElbv2LoadBalancers, aws_client.TestOptions{})


}




