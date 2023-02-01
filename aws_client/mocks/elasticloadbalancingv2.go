package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	elasticloadbalancingv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticloadbalancingv2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticloadbalancingv2ClientMockRecorder
}

type MockElasticloadbalancingv2ClientMockRecorder struct {
	mock *MockElasticloadbalancingv2Client
}

func NewMockElasticloadbalancingv2Client(ctrl *gomock.Controller) *MockElasticloadbalancingv2Client {
	mock := &MockElasticloadbalancingv2Client{ctrl: ctrl}
	mock.recorder = &MockElasticloadbalancingv2ClientMockRecorder{mock}
	return mock
}

func (m *MockElasticloadbalancingv2Client) EXPECT() *MockElasticloadbalancingv2ClientMockRecorder {
	return m.recorder
}

func (m *MockElasticloadbalancingv2Client) DescribeAccountLimits(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeAccountLimitsInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeAccountLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountLimits, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeAccountLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeAccountLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountLimits, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeAccountLimits), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeListenerCertificates(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenerCertificatesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeListenerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeListenerCertificates, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeListenerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeListenerCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeListenerCertificates, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeListenerCertificates), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeListeners(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenersInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeListenersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeListeners, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeListeners(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeListeners, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeListeners), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeLoadBalancerAttributes(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancerAttributesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancerAttributes, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeLoadBalancerAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancerAttributes, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeLoadBalancerAttributes), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeLoadBalancers(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancersInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancers, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancers, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeLoadBalancers), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeRules(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeRulesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRules, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRules, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeRules), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeSSLPolicies(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeSSLPoliciesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeSSLPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSSLPolicies, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeSSLPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeSSLPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSSLPolicies, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeSSLPolicies), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeTags(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTagsInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeTags), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeTargetGroupAttributes(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetGroupAttributesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetGroupAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTargetGroupAttributes, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetGroupAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeTargetGroupAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTargetGroupAttributes, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeTargetGroupAttributes), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeTargetGroups(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetGroupsInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTargetGroups, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeTargetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTargetGroups, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeTargetGroups), varargs...)
}

func (m *MockElasticloadbalancingv2Client) DescribeTargetHealth(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetHealthInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetHealthOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTargetHealth, varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticloadbalancingv2ClientMockRecorder) DescribeTargetHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTargetHealth, reflect.TypeOf((*MockElasticloadbalancingv2Client)(nil).DescribeTargetHealth), varargs...)
}
