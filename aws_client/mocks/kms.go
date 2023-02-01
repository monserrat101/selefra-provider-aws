package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	kms "github.com/aws/aws-sdk-go-v2/service/kms"
	gomock "github.com/golang/mock/gomock"
)

type MockKmsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockKmsClientMockRecorder
}

type MockKmsClientMockRecorder struct {
	mock *MockKmsClient
}

func NewMockKmsClient(ctrl *gomock.Controller) *MockKmsClient {
	mock := &MockKmsClient{ctrl: ctrl}
	mock.recorder = &MockKmsClientMockRecorder{mock}
	return mock
}

func (m *MockKmsClient) EXPECT() *MockKmsClientMockRecorder {
	return m.recorder
}

func (m *MockKmsClient) DescribeCustomKeyStores(arg0 context.Context, arg1 *kms.DescribeCustomKeyStoresInput, arg2 ...func(*kms.Options)) (*kms.DescribeCustomKeyStoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCustomKeyStores, varargs...)
	ret0, _ := ret[0].(*kms.DescribeCustomKeyStoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) DescribeCustomKeyStores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCustomKeyStores, reflect.TypeOf((*MockKmsClient)(nil).DescribeCustomKeyStores), varargs...)
}

func (m *MockKmsClient) DescribeKey(arg0 context.Context, arg1 *kms.DescribeKeyInput, arg2 ...func(*kms.Options)) (*kms.DescribeKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeKey, varargs...)
	ret0, _ := ret[0].(*kms.DescribeKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) DescribeKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeKey, reflect.TypeOf((*MockKmsClient)(nil).DescribeKey), varargs...)
}

func (m *MockKmsClient) GetKeyPolicy(arg0 context.Context, arg1 *kms.GetKeyPolicyInput, arg2 ...func(*kms.Options)) (*kms.GetKeyPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKeyPolicy, varargs...)
	ret0, _ := ret[0].(*kms.GetKeyPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) GetKeyPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKeyPolicy, reflect.TypeOf((*MockKmsClient)(nil).GetKeyPolicy), varargs...)
}

func (m *MockKmsClient) GetKeyRotationStatus(arg0 context.Context, arg1 *kms.GetKeyRotationStatusInput, arg2 ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKeyRotationStatus, varargs...)
	ret0, _ := ret[0].(*kms.GetKeyRotationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) GetKeyRotationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKeyRotationStatus, reflect.TypeOf((*MockKmsClient)(nil).GetKeyRotationStatus), varargs...)
}

func (m *MockKmsClient) GetParametersForImport(arg0 context.Context, arg1 *kms.GetParametersForImportInput, arg2 ...func(*kms.Options)) (*kms.GetParametersForImportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetParametersForImport, varargs...)
	ret0, _ := ret[0].(*kms.GetParametersForImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) GetParametersForImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetParametersForImport, reflect.TypeOf((*MockKmsClient)(nil).GetParametersForImport), varargs...)
}

func (m *MockKmsClient) GetPublicKey(arg0 context.Context, arg1 *kms.GetPublicKeyInput, arg2 ...func(*kms.Options)) (*kms.GetPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPublicKey, varargs...)
	ret0, _ := ret[0].(*kms.GetPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) GetPublicKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPublicKey, reflect.TypeOf((*MockKmsClient)(nil).GetPublicKey), varargs...)
}

func (m *MockKmsClient) ListAliases(arg0 context.Context, arg1 *kms.ListAliasesInput, arg2 ...func(*kms.Options)) (*kms.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAliases, varargs...)
	ret0, _ := ret[0].(*kms.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAliases, reflect.TypeOf((*MockKmsClient)(nil).ListAliases), varargs...)
}

func (m *MockKmsClient) ListGrants(arg0 context.Context, arg1 *kms.ListGrantsInput, arg2 ...func(*kms.Options)) (*kms.ListGrantsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGrants, varargs...)
	ret0, _ := ret[0].(*kms.ListGrantsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListGrants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGrants, reflect.TypeOf((*MockKmsClient)(nil).ListGrants), varargs...)
}

func (m *MockKmsClient) ListKeyPolicies(arg0 context.Context, arg1 *kms.ListKeyPoliciesInput, arg2 ...func(*kms.Options)) (*kms.ListKeyPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListKeyPolicies, varargs...)
	ret0, _ := ret[0].(*kms.ListKeyPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListKeyPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListKeyPolicies, reflect.TypeOf((*MockKmsClient)(nil).ListKeyPolicies), varargs...)
}

func (m *MockKmsClient) ListKeys(arg0 context.Context, arg1 *kms.ListKeysInput, arg2 ...func(*kms.Options)) (*kms.ListKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListKeys, varargs...)
	ret0, _ := ret[0].(*kms.ListKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListKeys, reflect.TypeOf((*MockKmsClient)(nil).ListKeys), varargs...)
}

func (m *MockKmsClient) ListResourceTags(arg0 context.Context, arg1 *kms.ListResourceTagsInput, arg2 ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceTags, varargs...)
	ret0, _ := ret[0].(*kms.ListResourceTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListResourceTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceTags, reflect.TypeOf((*MockKmsClient)(nil).ListResourceTags), varargs...)
}

func (m *MockKmsClient) ListRetirableGrants(arg0 context.Context, arg1 *kms.ListRetirableGrantsInput, arg2 ...func(*kms.Options)) (*kms.ListRetirableGrantsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRetirableGrants, varargs...)
	ret0, _ := ret[0].(*kms.ListRetirableGrantsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListRetirableGrants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRetirableGrants, reflect.TypeOf((*MockKmsClient)(nil).ListRetirableGrants), varargs...)
}
