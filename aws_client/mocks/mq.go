package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	mq "github.com/aws/aws-sdk-go-v2/service/mq"
	gomock "github.com/golang/mock/gomock"
)

type MockMqClient struct {
	ctrl		*gomock.Controller
	recorder	*MockMqClientMockRecorder
}

type MockMqClientMockRecorder struct {
	mock *MockMqClient
}

func NewMockMqClient(ctrl *gomock.Controller) *MockMqClient {
	mock := &MockMqClient{ctrl: ctrl}
	mock.recorder = &MockMqClientMockRecorder{mock}
	return mock
}

func (m *MockMqClient) EXPECT() *MockMqClientMockRecorder {
	return m.recorder
}

func (m *MockMqClient) DescribeBroker(arg0 context.Context, arg1 *mq.DescribeBrokerInput, arg2 ...func(*mq.Options)) (*mq.DescribeBrokerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBroker, varargs...)
	ret0, _ := ret[0].(*mq.DescribeBrokerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) DescribeBroker(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBroker, reflect.TypeOf((*MockMqClient)(nil).DescribeBroker), varargs...)
}

func (m *MockMqClient) DescribeBrokerEngineTypes(arg0 context.Context, arg1 *mq.DescribeBrokerEngineTypesInput, arg2 ...func(*mq.Options)) (*mq.DescribeBrokerEngineTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBrokerEngineTypes, varargs...)
	ret0, _ := ret[0].(*mq.DescribeBrokerEngineTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) DescribeBrokerEngineTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBrokerEngineTypes, reflect.TypeOf((*MockMqClient)(nil).DescribeBrokerEngineTypes), varargs...)
}

func (m *MockMqClient) DescribeBrokerInstanceOptions(arg0 context.Context, arg1 *mq.DescribeBrokerInstanceOptionsInput, arg2 ...func(*mq.Options)) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBrokerInstanceOptions, varargs...)
	ret0, _ := ret[0].(*mq.DescribeBrokerInstanceOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) DescribeBrokerInstanceOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBrokerInstanceOptions, reflect.TypeOf((*MockMqClient)(nil).DescribeBrokerInstanceOptions), varargs...)
}

func (m *MockMqClient) DescribeConfiguration(arg0 context.Context, arg1 *mq.DescribeConfigurationInput, arg2 ...func(*mq.Options)) (*mq.DescribeConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfiguration, varargs...)
	ret0, _ := ret[0].(*mq.DescribeConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) DescribeConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfiguration, reflect.TypeOf((*MockMqClient)(nil).DescribeConfiguration), varargs...)
}

func (m *MockMqClient) DescribeConfigurationRevision(arg0 context.Context, arg1 *mq.DescribeConfigurationRevisionInput, arg2 ...func(*mq.Options)) (*mq.DescribeConfigurationRevisionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationRevision, varargs...)
	ret0, _ := ret[0].(*mq.DescribeConfigurationRevisionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) DescribeConfigurationRevision(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationRevision, reflect.TypeOf((*MockMqClient)(nil).DescribeConfigurationRevision), varargs...)
}

func (m *MockMqClient) DescribeUser(arg0 context.Context, arg1 *mq.DescribeUserInput, arg2 ...func(*mq.Options)) (*mq.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUser, varargs...)
	ret0, _ := ret[0].(*mq.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) DescribeUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUser, reflect.TypeOf((*MockMqClient)(nil).DescribeUser), varargs...)
}

func (m *MockMqClient) ListBrokers(arg0 context.Context, arg1 *mq.ListBrokersInput, arg2 ...func(*mq.Options)) (*mq.ListBrokersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBrokers, varargs...)
	ret0, _ := ret[0].(*mq.ListBrokersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) ListBrokers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBrokers, reflect.TypeOf((*MockMqClient)(nil).ListBrokers), varargs...)
}

func (m *MockMqClient) ListConfigurationRevisions(arg0 context.Context, arg1 *mq.ListConfigurationRevisionsInput, arg2 ...func(*mq.Options)) (*mq.ListConfigurationRevisionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConfigurationRevisions, varargs...)
	ret0, _ := ret[0].(*mq.ListConfigurationRevisionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) ListConfigurationRevisions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConfigurationRevisions, reflect.TypeOf((*MockMqClient)(nil).ListConfigurationRevisions), varargs...)
}

func (m *MockMqClient) ListConfigurations(arg0 context.Context, arg1 *mq.ListConfigurationsInput, arg2 ...func(*mq.Options)) (*mq.ListConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConfigurations, varargs...)
	ret0, _ := ret[0].(*mq.ListConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) ListConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConfigurations, reflect.TypeOf((*MockMqClient)(nil).ListConfigurations), varargs...)
}

func (m *MockMqClient) ListTags(arg0 context.Context, arg1 *mq.ListTagsInput, arg2 ...func(*mq.Options)) (*mq.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*mq.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockMqClient)(nil).ListTags), varargs...)
}

func (m *MockMqClient) ListUsers(arg0 context.Context, arg1 *mq.ListUsersInput, arg2 ...func(*mq.Options)) (*mq.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsers, varargs...)
	ret0, _ := ret[0].(*mq.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMqClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsers, reflect.TypeOf((*MockMqClient)(nil).ListUsers), varargs...)
}
