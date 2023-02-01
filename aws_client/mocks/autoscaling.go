package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	autoscaling "github.com/aws/aws-sdk-go-v2/service/autoscaling"
	gomock "github.com/golang/mock/gomock"
)

type MockAutoscalingClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAutoscalingClientMockRecorder
}

type MockAutoscalingClientMockRecorder struct {
	mock *MockAutoscalingClient
}

func NewMockAutoscalingClient(ctrl *gomock.Controller) *MockAutoscalingClient {
	mock := &MockAutoscalingClient{ctrl: ctrl}
	mock.recorder = &MockAutoscalingClientMockRecorder{mock}
	return mock
}

func (m *MockAutoscalingClient) EXPECT() *MockAutoscalingClientMockRecorder {
	return m.recorder
}

func (m *MockAutoscalingClient) DescribeAccountLimits(arg0 context.Context, arg1 *autoscaling.DescribeAccountLimitsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAccountLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountLimits, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAccountLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeAccountLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountLimits, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAccountLimits), varargs...)
}

func (m *MockAutoscalingClient) DescribeAdjustmentTypes(arg0 context.Context, arg1 *autoscaling.DescribeAdjustmentTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAdjustmentTypes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAdjustmentTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeAdjustmentTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAdjustmentTypes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAdjustmentTypes), varargs...)
}

func (m *MockAutoscalingClient) DescribeAutoScalingGroups(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingGroupsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutoScalingGroups, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutoScalingGroups, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingGroups), varargs...)
}

func (m *MockAutoscalingClient) DescribeAutoScalingInstances(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingInstancesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutoScalingInstances, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutoScalingInstances, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingInstances), varargs...)
}

func (m *MockAutoscalingClient) DescribeAutoScalingNotificationTypes(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingNotificationTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutoScalingNotificationTypes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingNotificationTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingNotificationTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutoScalingNotificationTypes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingNotificationTypes), varargs...)
}

func (m *MockAutoscalingClient) DescribeInstanceRefreshes(arg0 context.Context, arg1 *autoscaling.DescribeInstanceRefreshesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceRefreshes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeInstanceRefreshesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeInstanceRefreshes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceRefreshes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeInstanceRefreshes), varargs...)
}

func (m *MockAutoscalingClient) DescribeLaunchConfigurations(arg0 context.Context, arg1 *autoscaling.DescribeLaunchConfigurationsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLaunchConfigurations, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLaunchConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLaunchConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLaunchConfigurations, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLaunchConfigurations), varargs...)
}

func (m *MockAutoscalingClient) DescribeLifecycleHookTypes(arg0 context.Context, arg1 *autoscaling.DescribeLifecycleHookTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLifecycleHookTypes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLifecycleHookTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLifecycleHookTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLifecycleHookTypes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLifecycleHookTypes), varargs...)
}

func (m *MockAutoscalingClient) DescribeLifecycleHooks(arg0 context.Context, arg1 *autoscaling.DescribeLifecycleHooksInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHooksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLifecycleHooks, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLifecycleHooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLifecycleHooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLifecycleHooks, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLifecycleHooks), varargs...)
}

func (m *MockAutoscalingClient) DescribeLoadBalancerTargetGroups(arg0 context.Context, arg1 *autoscaling.DescribeLoadBalancerTargetGroupsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancerTargetGroups, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLoadBalancerTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLoadBalancerTargetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancerTargetGroups, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLoadBalancerTargetGroups), varargs...)
}

func (m *MockAutoscalingClient) DescribeLoadBalancers(arg0 context.Context, arg1 *autoscaling.DescribeLoadBalancersInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoadBalancers, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoadBalancers, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLoadBalancers), varargs...)
}

func (m *MockAutoscalingClient) DescribeMetricCollectionTypes(arg0 context.Context, arg1 *autoscaling.DescribeMetricCollectionTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMetricCollectionTypes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeMetricCollectionTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeMetricCollectionTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMetricCollectionTypes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeMetricCollectionTypes), varargs...)
}

func (m *MockAutoscalingClient) DescribeNotificationConfigurations(arg0 context.Context, arg1 *autoscaling.DescribeNotificationConfigurationsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNotificationConfigurations, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeNotificationConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeNotificationConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNotificationConfigurations, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeNotificationConfigurations), varargs...)
}

func (m *MockAutoscalingClient) DescribePolicies(arg0 context.Context, arg1 *autoscaling.DescribePoliciesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePolicies, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePolicies, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribePolicies), varargs...)
}

func (m *MockAutoscalingClient) DescribeScalingActivities(arg0 context.Context, arg1 *autoscaling.DescribeScalingActivitiesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingActivitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScalingActivities, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScalingActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeScalingActivities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScalingActivities, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScalingActivities), varargs...)
}

func (m *MockAutoscalingClient) DescribeScalingProcessTypes(arg0 context.Context, arg1 *autoscaling.DescribeScalingProcessTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScalingProcessTypes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScalingProcessTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeScalingProcessTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScalingProcessTypes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScalingProcessTypes), varargs...)
}

func (m *MockAutoscalingClient) DescribeScheduledActions(arg0 context.Context, arg1 *autoscaling.DescribeScheduledActionsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScheduledActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScheduledActions, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeScheduledActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScheduledActions, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScheduledActions), varargs...)
}

func (m *MockAutoscalingClient) DescribeTags(arg0 context.Context, arg1 *autoscaling.DescribeTagsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTags), varargs...)
}

func (m *MockAutoscalingClient) DescribeTerminationPolicyTypes(arg0 context.Context, arg1 *autoscaling.DescribeTerminationPolicyTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTerminationPolicyTypes, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTerminationPolicyTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeTerminationPolicyTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTerminationPolicyTypes, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTerminationPolicyTypes), varargs...)
}

func (m *MockAutoscalingClient) DescribeTrafficSources(arg0 context.Context, arg1 *autoscaling.DescribeTrafficSourcesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTrafficSourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrafficSources, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTrafficSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeTrafficSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrafficSources, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTrafficSources), varargs...)
}

func (m *MockAutoscalingClient) DescribeWarmPool(arg0 context.Context, arg1 *autoscaling.DescribeWarmPoolInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeWarmPoolOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWarmPool, varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeWarmPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeWarmPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWarmPool, reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeWarmPool), varargs...)
}

func (m *MockAutoscalingClient) GetPredictiveScalingForecast(arg0 context.Context, arg1 *autoscaling.GetPredictiveScalingForecastInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.GetPredictiveScalingForecastOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPredictiveScalingForecast, varargs...)
	ret0, _ := ret[0].(*autoscaling.GetPredictiveScalingForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) GetPredictiveScalingForecast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPredictiveScalingForecast, reflect.TypeOf((*MockAutoscalingClient)(nil).GetPredictiveScalingForecast), varargs...)
}
