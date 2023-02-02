package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	elasticloadbalancing "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticloadbalancingClient struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticloadbalancingClientMockRecorder
}

type MockElasticloadbalancingClientMockRecorder struct {
	mock *MockElasticloadbalancingClient
}

func NewMockElasticloadbalancingClient(ctrl *gomock.Controller) *MockElasticloadbalancingClient {
	mock := &MockElasticloadbalancingClient{ctrl: ctrl}
	mock.recorder = &MockElasticloadbalancingClientMockRecorder{mock}
	return mock
}

func (m *MockElasticloadbalancingClient) EXPECT() *MockElasticloadbalancingClientMockRecorder {
	return m.recorder
}

func (m *MockElasticloadbalancingClient) DescribeAccountLimits(arg0 context.Context, arg1 *elasticloadbalancing.DescribeAccountLimitsInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeAccountLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountLimits, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeAccountLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeAccountLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountLimits, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeAccountLimits), varargs...)
}

func (m *MockElasticloadbalancingClient) DescribeInstanceHealth(arg0 context.Context, arg1 *elasticloadbalancing.DescribeInstanceHealthInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeInstanceHealthOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceHealth, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeInstanceHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeInstanceHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceHealth, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeInstanceHealth), varargs...)
}

func (m *MockElasticloadbalancingClient) DescribeLoadBalancerAttributes(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancerAttributesInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancerAttributes, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeLoadBalancerAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancerAttributes, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeLoadBalancerAttributes), varargs...)
}

func (m *MockElasticloadbalancingClient) DescribeLoadBalancerPolicies(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancerPoliciesInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancerPolicies, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancerPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeLoadBalancerPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancerPolicies, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeLoadBalancerPolicies), varargs...)
}

func (m *MockElasticloadbalancingClient) DescribeLoadBalancerPolicyTypes(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancerPolicyTypesInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerPolicyTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancerPolicyTypes, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancerPolicyTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeLoadBalancerPolicyTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancerPolicyTypes, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeLoadBalancerPolicyTypes), varargs...)
}

func (m *MockElasticloadbalancingClient) DescribeLoadBalancers(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancersInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancers, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancers, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeLoadBalancers), varargs...)
}

func (m *MockElasticloadbalancingClient) DescribeTags(arg0 context.Context, arg1 *elasticloadbalancing.DescribeTagsInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockElasticloadbalancingClient)(nil).DescribeTags), varargs...)
}
