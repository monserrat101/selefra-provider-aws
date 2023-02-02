package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	applicationautoscaling "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	gomock "github.com/golang/mock/gomock"
)

type MockApplicationautoscalingClient struct {
	ctrl		*gomock.Controller
	recorder	*MockApplicationautoscalingClientMockRecorder
}

type MockApplicationautoscalingClientMockRecorder struct {
	mock *MockApplicationautoscalingClient
}

func NewMockApplicationautoscalingClient(ctrl *gomock.Controller) *MockApplicationautoscalingClient {
	mock := &MockApplicationautoscalingClient{ctrl: ctrl}
	mock.recorder = &MockApplicationautoscalingClientMockRecorder{mock}
	return mock
}

func (m *MockApplicationautoscalingClient) EXPECT() *MockApplicationautoscalingClientMockRecorder {
	return m.recorder
}

func (m *MockApplicationautoscalingClient) DescribeScalableTargets(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalableTargetsInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScalableTargets, varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalableTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScalableTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScalableTargets, reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScalableTargets), varargs...)
}

func (m *MockApplicationautoscalingClient) DescribeScalingActivities(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingActivitiesInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScalingActivities, varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScalingActivities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScalingActivities, reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScalingActivities), varargs...)
}

func (m *MockApplicationautoscalingClient) DescribeScalingPolicies(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingPoliciesInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScalingPolicies, varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScalingPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScalingPolicies, reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScalingPolicies), varargs...)
}

func (m *MockApplicationautoscalingClient) DescribeScheduledActions(arg0 context.Context, arg1 *applicationautoscaling.DescribeScheduledActionsInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScheduledActions, varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApplicationautoscalingClientMockRecorder) DescribeScheduledActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScheduledActions, reflect.TypeOf((*MockApplicationautoscalingClient)(nil).DescribeScheduledActions), varargs...)
}
