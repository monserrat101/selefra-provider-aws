package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	apprunner "github.com/aws/aws-sdk-go-v2/service/apprunner"
	gomock "github.com/golang/mock/gomock"
)

type MockApprunnerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockApprunnerClientMockRecorder
}

type MockApprunnerClientMockRecorder struct {
	mock *MockApprunnerClient
}

func NewMockApprunnerClient(ctrl *gomock.Controller) *MockApprunnerClient {
	mock := &MockApprunnerClient{ctrl: ctrl}
	mock.recorder = &MockApprunnerClientMockRecorder{mock}
	return mock
}

func (m *MockApprunnerClient) EXPECT() *MockApprunnerClientMockRecorder {
	return m.recorder
}

func (m *MockApprunnerClient) DescribeAutoScalingConfiguration(arg0 context.Context, arg1 *apprunner.DescribeAutoScalingConfigurationInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeAutoScalingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutoScalingConfiguration, varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeAutoScalingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) DescribeAutoScalingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutoScalingConfiguration, reflect.TypeOf((*MockApprunnerClient)(nil).DescribeAutoScalingConfiguration), varargs...)
}

func (m *MockApprunnerClient) DescribeCustomDomains(arg0 context.Context, arg1 *apprunner.DescribeCustomDomainsInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeCustomDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCustomDomains, varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeCustomDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) DescribeCustomDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCustomDomains, reflect.TypeOf((*MockApprunnerClient)(nil).DescribeCustomDomains), varargs...)
}

func (m *MockApprunnerClient) DescribeObservabilityConfiguration(arg0 context.Context, arg1 *apprunner.DescribeObservabilityConfigurationInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeObservabilityConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeObservabilityConfiguration, varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeObservabilityConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) DescribeObservabilityConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeObservabilityConfiguration, reflect.TypeOf((*MockApprunnerClient)(nil).DescribeObservabilityConfiguration), varargs...)
}

func (m *MockApprunnerClient) DescribeService(arg0 context.Context, arg1 *apprunner.DescribeServiceInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeService, varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) DescribeService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeService, reflect.TypeOf((*MockApprunnerClient)(nil).DescribeService), varargs...)
}

func (m *MockApprunnerClient) DescribeVpcConnector(arg0 context.Context, arg1 *apprunner.DescribeVpcConnectorInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeVpcConnectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcConnector, varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeVpcConnectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) DescribeVpcConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcConnector, reflect.TypeOf((*MockApprunnerClient)(nil).DescribeVpcConnector), varargs...)
}

func (m *MockApprunnerClient) DescribeVpcIngressConnection(arg0 context.Context, arg1 *apprunner.DescribeVpcIngressConnectionInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeVpcIngressConnectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcIngressConnection, varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeVpcIngressConnectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) DescribeVpcIngressConnection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcIngressConnection, reflect.TypeOf((*MockApprunnerClient)(nil).DescribeVpcIngressConnection), varargs...)
}

func (m *MockApprunnerClient) ListAutoScalingConfigurations(arg0 context.Context, arg1 *apprunner.ListAutoScalingConfigurationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListAutoScalingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAutoScalingConfigurations, varargs...)
	ret0, _ := ret[0].(*apprunner.ListAutoScalingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListAutoScalingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAutoScalingConfigurations, reflect.TypeOf((*MockApprunnerClient)(nil).ListAutoScalingConfigurations), varargs...)
}

func (m *MockApprunnerClient) ListConnections(arg0 context.Context, arg1 *apprunner.ListConnectionsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConnections, varargs...)
	ret0, _ := ret[0].(*apprunner.ListConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConnections, reflect.TypeOf((*MockApprunnerClient)(nil).ListConnections), varargs...)
}

func (m *MockApprunnerClient) ListObservabilityConfigurations(arg0 context.Context, arg1 *apprunner.ListObservabilityConfigurationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListObservabilityConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListObservabilityConfigurations, varargs...)
	ret0, _ := ret[0].(*apprunner.ListObservabilityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListObservabilityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListObservabilityConfigurations, reflect.TypeOf((*MockApprunnerClient)(nil).ListObservabilityConfigurations), varargs...)
}

func (m *MockApprunnerClient) ListOperations(arg0 context.Context, arg1 *apprunner.ListOperationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOperations, varargs...)
	ret0, _ := ret[0].(*apprunner.ListOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOperations, reflect.TypeOf((*MockApprunnerClient)(nil).ListOperations), varargs...)
}

func (m *MockApprunnerClient) ListServices(arg0 context.Context, arg1 *apprunner.ListServicesInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServices, varargs...)
	ret0, _ := ret[0].(*apprunner.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServices, reflect.TypeOf((*MockApprunnerClient)(nil).ListServices), varargs...)
}

func (m *MockApprunnerClient) ListTagsForResource(arg0 context.Context, arg1 *apprunner.ListTagsForResourceInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*apprunner.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockApprunnerClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockApprunnerClient) ListVpcConnectors(arg0 context.Context, arg1 *apprunner.ListVpcConnectorsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListVpcConnectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVpcConnectors, varargs...)
	ret0, _ := ret[0].(*apprunner.ListVpcConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListVpcConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVpcConnectors, reflect.TypeOf((*MockApprunnerClient)(nil).ListVpcConnectors), varargs...)
}

func (m *MockApprunnerClient) ListVpcIngressConnections(arg0 context.Context, arg1 *apprunner.ListVpcIngressConnectionsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListVpcIngressConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVpcIngressConnections, varargs...)
	ret0, _ := ret[0].(*apprunner.ListVpcIngressConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApprunnerClientMockRecorder) ListVpcIngressConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVpcIngressConnections, reflect.TypeOf((*MockApprunnerClient)(nil).ListVpcIngressConnections), varargs...)
}
