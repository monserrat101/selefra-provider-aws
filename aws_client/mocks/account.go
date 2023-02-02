package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	account "github.com/aws/aws-sdk-go-v2/service/account"
	gomock "github.com/golang/mock/gomock"
)

type MockAccountClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAccountClientMockRecorder
}

type MockAccountClientMockRecorder struct {
	mock *MockAccountClient
}

func NewMockAccountClient(ctrl *gomock.Controller) *MockAccountClient {
	mock := &MockAccountClient{ctrl: ctrl}
	mock.recorder = &MockAccountClientMockRecorder{mock}
	return mock
}

func (m *MockAccountClient) EXPECT() *MockAccountClientMockRecorder {
	return m.recorder
}

func (m *MockAccountClient) GetAlternateContact(arg0 context.Context, arg1 *account.GetAlternateContactInput, arg2 ...func(*account.Options)) (*account.GetAlternateContactOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAlternateContact, varargs...)
	ret0, _ := ret[0].(*account.GetAlternateContactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccountClientMockRecorder) GetAlternateContact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAlternateContact, reflect.TypeOf((*MockAccountClient)(nil).GetAlternateContact), varargs...)
}

func (m *MockAccountClient) GetContactInformation(arg0 context.Context, arg1 *account.GetContactInformationInput, arg2 ...func(*account.Options)) (*account.GetContactInformationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContactInformation, varargs...)
	ret0, _ := ret[0].(*account.GetContactInformationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccountClientMockRecorder) GetContactInformation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContactInformation, reflect.TypeOf((*MockAccountClient)(nil).GetContactInformation), varargs...)
}
