package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	elasticbeanstalk "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticbeanstalkClient struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticbeanstalkClientMockRecorder
}

type MockElasticbeanstalkClientMockRecorder struct {
	mock *MockElasticbeanstalkClient
}

func NewMockElasticbeanstalkClient(ctrl *gomock.Controller) *MockElasticbeanstalkClient {
	mock := &MockElasticbeanstalkClient{ctrl: ctrl}
	mock.recorder = &MockElasticbeanstalkClientMockRecorder{mock}
	return mock
}

func (m *MockElasticbeanstalkClient) EXPECT() *MockElasticbeanstalkClientMockRecorder {
	return m.recorder
}

func (m *MockElasticbeanstalkClient) DescribeAccountAttributes(arg0 context.Context, arg1 *elasticbeanstalk.DescribeAccountAttributesInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAttributes, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeAccountAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAttributes, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeAccountAttributes), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeApplicationVersions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeApplicationVersionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApplicationVersions, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeApplicationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeApplicationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApplicationVersions, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeApplicationVersions), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeApplications(arg0 context.Context, arg1 *elasticbeanstalk.DescribeApplicationsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApplications, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApplications, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeApplications), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeConfigurationOptions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeConfigurationOptionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationOptions, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeConfigurationOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeConfigurationOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationOptions, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeConfigurationOptions), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeConfigurationSettings(arg0 context.Context, arg1 *elasticbeanstalk.DescribeConfigurationSettingsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationSettings, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeConfigurationSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeConfigurationSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationSettings, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeConfigurationSettings), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEnvironmentHealth(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentHealthInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEnvironmentHealth, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEnvironmentHealth, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentHealth), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEnvironmentManagedActionHistory(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEnvironmentManagedActionHistory, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentManagedActionHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEnvironmentManagedActionHistory, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentManagedActionHistory), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEnvironmentManagedActions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentManagedActionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEnvironmentManagedActions, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentManagedActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEnvironmentManagedActions, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentManagedActions), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEnvironmentResources(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentResourcesInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEnvironmentResources, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEnvironmentResources, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentResources), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEnvironments(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEnvironments, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEnvironments, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironments), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEvents(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEventsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEvents), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeInstancesHealth(arg0 context.Context, arg1 *elasticbeanstalk.DescribeInstancesHealthInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstancesHealth, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeInstancesHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeInstancesHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstancesHealth, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeInstancesHealth), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribePlatformVersion(arg0 context.Context, arg1 *elasticbeanstalk.DescribePlatformVersionInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePlatformVersion, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribePlatformVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribePlatformVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePlatformVersion, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribePlatformVersion), varargs...)
}

func (m *MockElasticbeanstalkClient) ListAvailableSolutionStacks(arg0 context.Context, arg1 *elasticbeanstalk.ListAvailableSolutionStacksInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAvailableSolutionStacks, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListAvailableSolutionStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) ListAvailableSolutionStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAvailableSolutionStacks, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListAvailableSolutionStacks), varargs...)
}

func (m *MockElasticbeanstalkClient) ListPlatformBranches(arg0 context.Context, arg1 *elasticbeanstalk.ListPlatformBranchesInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPlatformBranches, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListPlatformBranchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) ListPlatformBranches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPlatformBranches, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListPlatformBranches), varargs...)
}

func (m *MockElasticbeanstalkClient) ListPlatformVersions(arg0 context.Context, arg1 *elasticbeanstalk.ListPlatformVersionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPlatformVersions, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListPlatformVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) ListPlatformVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPlatformVersions, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListPlatformVersions), varargs...)
}

func (m *MockElasticbeanstalkClient) ListTagsForResource(arg0 context.Context, arg1 *elasticbeanstalk.ListTagsForResourceInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListTagsForResource), varargs...)
}
