package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	amp "github.com/aws/aws-sdk-go-v2/service/amp"
	gomock "github.com/golang/mock/gomock"
)

type MockAmpClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAmpClientMockRecorder
}

type MockAmpClientMockRecorder struct {
	mock *MockAmpClient
}

func NewMockAmpClient(ctrl *gomock.Controller) *MockAmpClient {
	mock := &MockAmpClient{ctrl: ctrl}
	mock.recorder = &MockAmpClientMockRecorder{mock}
	return mock
}

func (m *MockAmpClient) EXPECT() *MockAmpClientMockRecorder {
	return m.recorder
}

func (m *MockAmpClient) DescribeAlertManagerDefinition(arg0 context.Context, arg1 *amp.DescribeAlertManagerDefinitionInput, arg2 ...func(*amp.Options)) (*amp.DescribeAlertManagerDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAlertManagerDefinition, varargs...)
	ret0, _ := ret[0].(*amp.DescribeAlertManagerDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) DescribeAlertManagerDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAlertManagerDefinition, reflect.TypeOf((*MockAmpClient)(nil).DescribeAlertManagerDefinition), varargs...)
}

func (m *MockAmpClient) DescribeLoggingConfiguration(arg0 context.Context, arg1 *amp.DescribeLoggingConfigurationInput, arg2 ...func(*amp.Options)) (*amp.DescribeLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoggingConfiguration, varargs...)
	ret0, _ := ret[0].(*amp.DescribeLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) DescribeLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoggingConfiguration, reflect.TypeOf((*MockAmpClient)(nil).DescribeLoggingConfiguration), varargs...)
}

func (m *MockAmpClient) DescribeRuleGroupsNamespace(arg0 context.Context, arg1 *amp.DescribeRuleGroupsNamespaceInput, arg2 ...func(*amp.Options)) (*amp.DescribeRuleGroupsNamespaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRuleGroupsNamespace, varargs...)
	ret0, _ := ret[0].(*amp.DescribeRuleGroupsNamespaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) DescribeRuleGroupsNamespace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRuleGroupsNamespace, reflect.TypeOf((*MockAmpClient)(nil).DescribeRuleGroupsNamespace), varargs...)
}

func (m *MockAmpClient) DescribeWorkspace(arg0 context.Context, arg1 *amp.DescribeWorkspaceInput, arg2 ...func(*amp.Options)) (*amp.DescribeWorkspaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspace, varargs...)
	ret0, _ := ret[0].(*amp.DescribeWorkspaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) DescribeWorkspace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspace, reflect.TypeOf((*MockAmpClient)(nil).DescribeWorkspace), varargs...)
}

func (m *MockAmpClient) ListRuleGroupsNamespaces(arg0 context.Context, arg1 *amp.ListRuleGroupsNamespacesInput, arg2 ...func(*amp.Options)) (*amp.ListRuleGroupsNamespacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRuleGroupsNamespaces, varargs...)
	ret0, _ := ret[0].(*amp.ListRuleGroupsNamespacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) ListRuleGroupsNamespaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRuleGroupsNamespaces, reflect.TypeOf((*MockAmpClient)(nil).ListRuleGroupsNamespaces), varargs...)
}

func (m *MockAmpClient) ListTagsForResource(arg0 context.Context, arg1 *amp.ListTagsForResourceInput, arg2 ...func(*amp.Options)) (*amp.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*amp.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockAmpClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockAmpClient) ListWorkspaces(arg0 context.Context, arg1 *amp.ListWorkspacesInput, arg2 ...func(*amp.Options)) (*amp.ListWorkspacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWorkspaces, varargs...)
	ret0, _ := ret[0].(*amp.ListWorkspacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAmpClientMockRecorder) ListWorkspaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWorkspaces, reflect.TypeOf((*MockAmpClient)(nil).ListWorkspaces), varargs...)
}
