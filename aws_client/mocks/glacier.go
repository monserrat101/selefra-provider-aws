package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	glacier "github.com/aws/aws-sdk-go-v2/service/glacier"
	gomock "github.com/golang/mock/gomock"
)

type MockGlacierClient struct {
	ctrl		*gomock.Controller
	recorder	*MockGlacierClientMockRecorder
}

type MockGlacierClientMockRecorder struct {
	mock *MockGlacierClient
}

func NewMockGlacierClient(ctrl *gomock.Controller) *MockGlacierClient {
	mock := &MockGlacierClient{ctrl: ctrl}
	mock.recorder = &MockGlacierClientMockRecorder{mock}
	return mock
}

func (m *MockGlacierClient) EXPECT() *MockGlacierClientMockRecorder {
	return m.recorder
}

func (m *MockGlacierClient) DescribeJob(arg0 context.Context, arg1 *glacier.DescribeJobInput, arg2 ...func(*glacier.Options)) (*glacier.DescribeJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJob, varargs...)
	ret0, _ := ret[0].(*glacier.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) DescribeJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJob, reflect.TypeOf((*MockGlacierClient)(nil).DescribeJob), varargs...)
}

func (m *MockGlacierClient) DescribeVault(arg0 context.Context, arg1 *glacier.DescribeVaultInput, arg2 ...func(*glacier.Options)) (*glacier.DescribeVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVault, varargs...)
	ret0, _ := ret[0].(*glacier.DescribeVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) DescribeVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVault, reflect.TypeOf((*MockGlacierClient)(nil).DescribeVault), varargs...)
}

func (m *MockGlacierClient) GetDataRetrievalPolicy(arg0 context.Context, arg1 *glacier.GetDataRetrievalPolicyInput, arg2 ...func(*glacier.Options)) (*glacier.GetDataRetrievalPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataRetrievalPolicy, varargs...)
	ret0, _ := ret[0].(*glacier.GetDataRetrievalPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetDataRetrievalPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataRetrievalPolicy, reflect.TypeOf((*MockGlacierClient)(nil).GetDataRetrievalPolicy), varargs...)
}

func (m *MockGlacierClient) GetJobOutput(arg0 context.Context, arg1 *glacier.GetJobOutputInput, arg2 ...func(*glacier.Options)) (*glacier.GetJobOutputOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobOutput, varargs...)
	ret0, _ := ret[0].(*glacier.GetJobOutputOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetJobOutput(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobOutput, reflect.TypeOf((*MockGlacierClient)(nil).GetJobOutput), varargs...)
}

func (m *MockGlacierClient) GetVaultAccessPolicy(arg0 context.Context, arg1 *glacier.GetVaultAccessPolicyInput, arg2 ...func(*glacier.Options)) (*glacier.GetVaultAccessPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVaultAccessPolicy, varargs...)
	ret0, _ := ret[0].(*glacier.GetVaultAccessPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetVaultAccessPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVaultAccessPolicy, reflect.TypeOf((*MockGlacierClient)(nil).GetVaultAccessPolicy), varargs...)
}

func (m *MockGlacierClient) GetVaultLock(arg0 context.Context, arg1 *glacier.GetVaultLockInput, arg2 ...func(*glacier.Options)) (*glacier.GetVaultLockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVaultLock, varargs...)
	ret0, _ := ret[0].(*glacier.GetVaultLockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetVaultLock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVaultLock, reflect.TypeOf((*MockGlacierClient)(nil).GetVaultLock), varargs...)
}

func (m *MockGlacierClient) GetVaultNotifications(arg0 context.Context, arg1 *glacier.GetVaultNotificationsInput, arg2 ...func(*glacier.Options)) (*glacier.GetVaultNotificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVaultNotifications, varargs...)
	ret0, _ := ret[0].(*glacier.GetVaultNotificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetVaultNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVaultNotifications, reflect.TypeOf((*MockGlacierClient)(nil).GetVaultNotifications), varargs...)
}

func (m *MockGlacierClient) ListJobs(arg0 context.Context, arg1 *glacier.ListJobsInput, arg2 ...func(*glacier.Options)) (*glacier.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobs, varargs...)
	ret0, _ := ret[0].(*glacier.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobs, reflect.TypeOf((*MockGlacierClient)(nil).ListJobs), varargs...)
}

func (m *MockGlacierClient) ListMultipartUploads(arg0 context.Context, arg1 *glacier.ListMultipartUploadsInput, arg2 ...func(*glacier.Options)) (*glacier.ListMultipartUploadsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMultipartUploads, varargs...)
	ret0, _ := ret[0].(*glacier.ListMultipartUploadsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListMultipartUploads(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMultipartUploads, reflect.TypeOf((*MockGlacierClient)(nil).ListMultipartUploads), varargs...)
}

func (m *MockGlacierClient) ListParts(arg0 context.Context, arg1 *glacier.ListPartsInput, arg2 ...func(*glacier.Options)) (*glacier.ListPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListParts, varargs...)
	ret0, _ := ret[0].(*glacier.ListPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListParts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListParts, reflect.TypeOf((*MockGlacierClient)(nil).ListParts), varargs...)
}

func (m *MockGlacierClient) ListProvisionedCapacity(arg0 context.Context, arg1 *glacier.ListProvisionedCapacityInput, arg2 ...func(*glacier.Options)) (*glacier.ListProvisionedCapacityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisionedCapacity, varargs...)
	ret0, _ := ret[0].(*glacier.ListProvisionedCapacityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListProvisionedCapacity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisionedCapacity, reflect.TypeOf((*MockGlacierClient)(nil).ListProvisionedCapacity), varargs...)
}

func (m *MockGlacierClient) ListTagsForVault(arg0 context.Context, arg1 *glacier.ListTagsForVaultInput, arg2 ...func(*glacier.Options)) (*glacier.ListTagsForVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForVault, varargs...)
	ret0, _ := ret[0].(*glacier.ListTagsForVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListTagsForVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForVault, reflect.TypeOf((*MockGlacierClient)(nil).ListTagsForVault), varargs...)
}

func (m *MockGlacierClient) ListVaults(arg0 context.Context, arg1 *glacier.ListVaultsInput, arg2 ...func(*glacier.Options)) (*glacier.ListVaultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVaults, varargs...)
	ret0, _ := ret[0].(*glacier.ListVaultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListVaults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVaults, reflect.TypeOf((*MockGlacierClient)(nil).ListVaults), varargs...)
}
