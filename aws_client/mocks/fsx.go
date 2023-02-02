package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	fsx "github.com/aws/aws-sdk-go-v2/service/fsx"
	gomock "github.com/golang/mock/gomock"
)

type MockFsxClient struct {
	ctrl		*gomock.Controller
	recorder	*MockFsxClientMockRecorder
}

type MockFsxClientMockRecorder struct {
	mock *MockFsxClient
}

func NewMockFsxClient(ctrl *gomock.Controller) *MockFsxClient {
	mock := &MockFsxClient{ctrl: ctrl}
	mock.recorder = &MockFsxClientMockRecorder{mock}
	return mock
}

func (m *MockFsxClient) EXPECT() *MockFsxClientMockRecorder {
	return m.recorder
}

func (m *MockFsxClient) DescribeBackups(arg0 context.Context, arg1 *fsx.DescribeBackupsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBackups, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBackups, reflect.TypeOf((*MockFsxClient)(nil).DescribeBackups), varargs...)
}

func (m *MockFsxClient) DescribeDataRepositoryAssociations(arg0 context.Context, arg1 *fsx.DescribeDataRepositoryAssociationsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataRepositoryAssociations, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeDataRepositoryAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeDataRepositoryAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataRepositoryAssociations, reflect.TypeOf((*MockFsxClient)(nil).DescribeDataRepositoryAssociations), varargs...)
}

func (m *MockFsxClient) DescribeDataRepositoryTasks(arg0 context.Context, arg1 *fsx.DescribeDataRepositoryTasksInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataRepositoryTasks, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeDataRepositoryTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeDataRepositoryTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataRepositoryTasks, reflect.TypeOf((*MockFsxClient)(nil).DescribeDataRepositoryTasks), varargs...)
}

func (m *MockFsxClient) DescribeFileCaches(arg0 context.Context, arg1 *fsx.DescribeFileCachesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileCachesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFileCaches, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileCachesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeFileCaches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFileCaches, reflect.TypeOf((*MockFsxClient)(nil).DescribeFileCaches), varargs...)
}

func (m *MockFsxClient) DescribeFileSystemAliases(arg0 context.Context, arg1 *fsx.DescribeFileSystemAliasesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileSystemAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFileSystemAliases, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileSystemAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeFileSystemAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFileSystemAliases, reflect.TypeOf((*MockFsxClient)(nil).DescribeFileSystemAliases), varargs...)
}

func (m *MockFsxClient) DescribeFileSystems(arg0 context.Context, arg1 *fsx.DescribeFileSystemsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFileSystems, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeFileSystems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFileSystems, reflect.TypeOf((*MockFsxClient)(nil).DescribeFileSystems), varargs...)
}

func (m *MockFsxClient) DescribeSnapshots(arg0 context.Context, arg1 *fsx.DescribeSnapshotsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshots, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshots, reflect.TypeOf((*MockFsxClient)(nil).DescribeSnapshots), varargs...)
}

func (m *MockFsxClient) DescribeStorageVirtualMachines(arg0 context.Context, arg1 *fsx.DescribeStorageVirtualMachinesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeStorageVirtualMachinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStorageVirtualMachines, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeStorageVirtualMachinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeStorageVirtualMachines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStorageVirtualMachines, reflect.TypeOf((*MockFsxClient)(nil).DescribeStorageVirtualMachines), varargs...)
}

func (m *MockFsxClient) DescribeVolumes(arg0 context.Context, arg1 *fsx.DescribeVolumesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeVolumesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVolumes, varargs...)
	ret0, _ := ret[0].(*fsx.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeVolumes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVolumes, reflect.TypeOf((*MockFsxClient)(nil).DescribeVolumes), varargs...)
}

func (m *MockFsxClient) ListTagsForResource(arg0 context.Context, arg1 *fsx.ListTagsForResourceInput, arg2 ...func(*fsx.Options)) (*fsx.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*fsx.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockFsxClient)(nil).ListTagsForResource), varargs...)
}
