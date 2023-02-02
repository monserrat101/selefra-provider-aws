package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	resourcegroups "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	gomock "github.com/golang/mock/gomock"
)

type MockResourcegroupsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockResourcegroupsClientMockRecorder
}

type MockResourcegroupsClientMockRecorder struct {
	mock *MockResourcegroupsClient
}

func NewMockResourcegroupsClient(ctrl *gomock.Controller) *MockResourcegroupsClient {
	mock := &MockResourcegroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourcegroupsClientMockRecorder{mock}
	return mock
}

func (m *MockResourcegroupsClient) EXPECT() *MockResourcegroupsClientMockRecorder {
	return m.recorder
}

func (m *MockResourcegroupsClient) GetGroup(arg0 context.Context, arg1 *resourcegroups.GetGroupInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroup, varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroup, reflect.TypeOf((*MockResourcegroupsClient)(nil).GetGroup), varargs...)
}

func (m *MockResourcegroupsClient) GetGroupConfiguration(arg0 context.Context, arg1 *resourcegroups.GetGroupConfigurationInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroupConfiguration, varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) GetGroupConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroupConfiguration, reflect.TypeOf((*MockResourcegroupsClient)(nil).GetGroupConfiguration), varargs...)
}

func (m *MockResourcegroupsClient) GetGroupQuery(arg0 context.Context, arg1 *resourcegroups.GetGroupQueryInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroupQuery, varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) GetGroupQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroupQuery, reflect.TypeOf((*MockResourcegroupsClient)(nil).GetGroupQuery), varargs...)
}

func (m *MockResourcegroupsClient) GetTags(arg0 context.Context, arg1 *resourcegroups.GetTagsInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTags, varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTags, reflect.TypeOf((*MockResourcegroupsClient)(nil).GetTags), varargs...)
}

func (m *MockResourcegroupsClient) ListGroupResources(arg0 context.Context, arg1 *resourcegroups.ListGroupResourcesInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroupResources, varargs...)
	ret0, _ := ret[0].(*resourcegroups.ListGroupResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) ListGroupResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroupResources, reflect.TypeOf((*MockResourcegroupsClient)(nil).ListGroupResources), varargs...)
}

func (m *MockResourcegroupsClient) ListGroups(arg0 context.Context, arg1 *resourcegroups.ListGroupsInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroups, varargs...)
	ret0, _ := ret[0].(*resourcegroups.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroups, reflect.TypeOf((*MockResourcegroupsClient)(nil).ListGroups), varargs...)
}

func (m *MockResourcegroupsClient) SearchResources(arg0 context.Context, arg1 *resourcegroups.SearchResourcesInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchResources, varargs...)
	ret0, _ := ret[0].(*resourcegroups.SearchResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcegroupsClientMockRecorder) SearchResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchResources, reflect.TypeOf((*MockResourcegroupsClient)(nil).SearchResources), varargs...)
}
