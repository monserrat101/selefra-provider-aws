package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	mwaa "github.com/aws/aws-sdk-go-v2/service/mwaa"
	gomock "github.com/golang/mock/gomock"
)

type MockMwaaClient struct {
	ctrl		*gomock.Controller
	recorder	*MockMwaaClientMockRecorder
}

type MockMwaaClientMockRecorder struct {
	mock *MockMwaaClient
}

func NewMockMwaaClient(ctrl *gomock.Controller) *MockMwaaClient {
	mock := &MockMwaaClient{ctrl: ctrl}
	mock.recorder = &MockMwaaClientMockRecorder{mock}
	return mock
}

func (m *MockMwaaClient) EXPECT() *MockMwaaClientMockRecorder {
	return m.recorder
}

func (m *MockMwaaClient) GetEnvironment(arg0 context.Context, arg1 *mwaa.GetEnvironmentInput, arg2 ...func(*mwaa.Options)) (*mwaa.GetEnvironmentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEnvironment, varargs...)
	ret0, _ := ret[0].(*mwaa.GetEnvironmentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMwaaClientMockRecorder) GetEnvironment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEnvironment, reflect.TypeOf((*MockMwaaClient)(nil).GetEnvironment), varargs...)
}

func (m *MockMwaaClient) ListEnvironments(arg0 context.Context, arg1 *mwaa.ListEnvironmentsInput, arg2 ...func(*mwaa.Options)) (*mwaa.ListEnvironmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEnvironments, varargs...)
	ret0, _ := ret[0].(*mwaa.ListEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMwaaClientMockRecorder) ListEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEnvironments, reflect.TypeOf((*MockMwaaClient)(nil).ListEnvironments), varargs...)
}

func (m *MockMwaaClient) ListTagsForResource(arg0 context.Context, arg1 *mwaa.ListTagsForResourceInput, arg2 ...func(*mwaa.Options)) (*mwaa.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*mwaa.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMwaaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockMwaaClient)(nil).ListTagsForResource), varargs...)
}
