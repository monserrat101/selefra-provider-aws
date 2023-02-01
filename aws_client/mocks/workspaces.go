package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	workspaces "github.com/aws/aws-sdk-go-v2/service/workspaces"
	gomock "github.com/golang/mock/gomock"
)

type MockWorkspacesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockWorkspacesClientMockRecorder
}

type MockWorkspacesClientMockRecorder struct {
	mock *MockWorkspacesClient
}

func NewMockWorkspacesClient(ctrl *gomock.Controller) *MockWorkspacesClient {
	mock := &MockWorkspacesClient{ctrl: ctrl}
	mock.recorder = &MockWorkspacesClientMockRecorder{mock}
	return mock
}

func (m *MockWorkspacesClient) EXPECT() *MockWorkspacesClientMockRecorder {
	return m.recorder
}

func (m *MockWorkspacesClient) DescribeAccount(arg0 context.Context, arg1 *workspaces.DescribeAccountInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccount, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccount, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeAccount), varargs...)
}

func (m *MockWorkspacesClient) DescribeAccountModifications(arg0 context.Context, arg1 *workspaces.DescribeAccountModificationsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeAccountModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountModifications, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeAccountModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeAccountModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountModifications, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeAccountModifications), varargs...)
}

func (m *MockWorkspacesClient) DescribeClientBranding(arg0 context.Context, arg1 *workspaces.DescribeClientBrandingInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeClientBrandingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientBranding, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeClientBrandingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeClientBranding(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientBranding, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeClientBranding), varargs...)
}

func (m *MockWorkspacesClient) DescribeClientProperties(arg0 context.Context, arg1 *workspaces.DescribeClientPropertiesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeClientPropertiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientProperties, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeClientPropertiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeClientProperties(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientProperties, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeClientProperties), varargs...)
}

func (m *MockWorkspacesClient) DescribeConnectClientAddIns(arg0 context.Context, arg1 *workspaces.DescribeConnectClientAddInsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeConnectClientAddInsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnectClientAddIns, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeConnectClientAddInsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeConnectClientAddIns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnectClientAddIns, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeConnectClientAddIns), varargs...)
}

func (m *MockWorkspacesClient) DescribeConnectionAliasPermissions(arg0 context.Context, arg1 *workspaces.DescribeConnectionAliasPermissionsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnectionAliasPermissions, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeConnectionAliasPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeConnectionAliasPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnectionAliasPermissions, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeConnectionAliasPermissions), varargs...)
}

func (m *MockWorkspacesClient) DescribeConnectionAliases(arg0 context.Context, arg1 *workspaces.DescribeConnectionAliasesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnectionAliases, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeConnectionAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeConnectionAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnectionAliases, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeConnectionAliases), varargs...)
}

func (m *MockWorkspacesClient) DescribeIpGroups(arg0 context.Context, arg1 *workspaces.DescribeIpGroupsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeIpGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIpGroups, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeIpGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeIpGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIpGroups, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeIpGroups), varargs...)
}

func (m *MockWorkspacesClient) DescribeTags(arg0 context.Context, arg1 *workspaces.DescribeTagsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeTags), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaceBundles(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceBundlesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspaceBundles, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspaceBundles, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceBundles), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaceDirectories(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceDirectoriesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspaceDirectories, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceDirectoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceDirectories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspaceDirectories, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceDirectories), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaceImagePermissions(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceImagePermissionsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspaceImagePermissions, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceImagePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceImagePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspaceImagePermissions, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceImagePermissions), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaceImages(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceImagesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspaceImages, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspaceImages, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceImages), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaceSnapshots(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceSnapshotsInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspaceSnapshots, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspaceSnapshots, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceSnapshots), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaces(arg0 context.Context, arg1 *workspaces.DescribeWorkspacesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspaces, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspaces, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaces), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspacesConnectionStatus(arg0 context.Context, arg1 *workspaces.DescribeWorkspacesConnectionStatusInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkspacesConnectionStatus, varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspacesConnectionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspacesConnectionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkspacesConnectionStatus, reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspacesConnectionStatus), varargs...)
}

func (m *MockWorkspacesClient) ListAvailableManagementCidrRanges(arg0 context.Context, arg1 *workspaces.ListAvailableManagementCidrRangesInput, arg2 ...func(*workspaces.Options)) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAvailableManagementCidrRanges, varargs...)
	ret0, _ := ret[0].(*workspaces.ListAvailableManagementCidrRangesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) ListAvailableManagementCidrRanges(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAvailableManagementCidrRanges, reflect.TypeOf((*MockWorkspacesClient)(nil).ListAvailableManagementCidrRanges), varargs...)
}
