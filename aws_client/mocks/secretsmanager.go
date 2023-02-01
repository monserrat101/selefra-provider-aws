package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

type MockSecretsmanagerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSecretsmanagerClientMockRecorder
}

type MockSecretsmanagerClientMockRecorder struct {
	mock *MockSecretsmanagerClient
}

func NewMockSecretsmanagerClient(ctrl *gomock.Controller) *MockSecretsmanagerClient {
	mock := &MockSecretsmanagerClient{ctrl: ctrl}
	mock.recorder = &MockSecretsmanagerClientMockRecorder{mock}
	return mock
}

func (m *MockSecretsmanagerClient) EXPECT() *MockSecretsmanagerClientMockRecorder {
	return m.recorder
}

func (m *MockSecretsmanagerClient) DescribeSecret(arg0 context.Context, arg1 *secretsmanager.DescribeSecretInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecret, varargs...)
	ret0, _ := ret[0].(*secretsmanager.DescribeSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretsmanagerClientMockRecorder) DescribeSecret(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecret, reflect.TypeOf((*MockSecretsmanagerClient)(nil).DescribeSecret), varargs...)
}

func (m *MockSecretsmanagerClient) GetRandomPassword(arg0 context.Context, arg1 *secretsmanager.GetRandomPasswordInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetRandomPasswordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRandomPassword, varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetRandomPasswordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretsmanagerClientMockRecorder) GetRandomPassword(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRandomPassword, reflect.TypeOf((*MockSecretsmanagerClient)(nil).GetRandomPassword), varargs...)
}

func (m *MockSecretsmanagerClient) GetResourcePolicy(arg0 context.Context, arg1 *secretsmanager.GetResourcePolicyInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourcePolicy, varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretsmanagerClientMockRecorder) GetResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourcePolicy, reflect.TypeOf((*MockSecretsmanagerClient)(nil).GetResourcePolicy), varargs...)
}

func (m *MockSecretsmanagerClient) GetSecretValue(arg0 context.Context, arg1 *secretsmanager.GetSecretValueInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSecretValue, varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretsmanagerClientMockRecorder) GetSecretValue(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSecretValue, reflect.TypeOf((*MockSecretsmanagerClient)(nil).GetSecretValue), varargs...)
}

func (m *MockSecretsmanagerClient) ListSecretVersionIds(arg0 context.Context, arg1 *secretsmanager.ListSecretVersionIdsInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSecretVersionIds, varargs...)
	ret0, _ := ret[0].(*secretsmanager.ListSecretVersionIdsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretsmanagerClientMockRecorder) ListSecretVersionIds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSecretVersionIds, reflect.TypeOf((*MockSecretsmanagerClient)(nil).ListSecretVersionIds), varargs...)
}

func (m *MockSecretsmanagerClient) ListSecrets(arg0 context.Context, arg1 *secretsmanager.ListSecretsInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSecrets, varargs...)
	ret0, _ := ret[0].(*secretsmanager.ListSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretsmanagerClientMockRecorder) ListSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSecrets, reflect.TypeOf((*MockSecretsmanagerClient)(nil).ListSecrets), varargs...)
}
