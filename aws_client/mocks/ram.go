package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ram "github.com/aws/aws-sdk-go-v2/service/ram"
	gomock "github.com/golang/mock/gomock"
)

type MockRamClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRamClientMockRecorder
}

type MockRamClientMockRecorder struct {
	mock *MockRamClient
}

func NewMockRamClient(ctrl *gomock.Controller) *MockRamClient {
	mock := &MockRamClient{ctrl: ctrl}
	mock.recorder = &MockRamClientMockRecorder{mock}
	return mock
}

func (m *MockRamClient) EXPECT() *MockRamClientMockRecorder {
	return m.recorder
}

func (m *MockRamClient) GetPermission(arg0 context.Context, arg1 *ram.GetPermissionInput, arg2 ...func(*ram.Options)) (*ram.GetPermissionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPermission, varargs...)
	ret0, _ := ret[0].(*ram.GetPermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) GetPermission(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPermission, reflect.TypeOf((*MockRamClient)(nil).GetPermission), varargs...)
}

func (m *MockRamClient) GetResourcePolicies(arg0 context.Context, arg1 *ram.GetResourcePoliciesInput, arg2 ...func(*ram.Options)) (*ram.GetResourcePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourcePolicies, varargs...)
	ret0, _ := ret[0].(*ram.GetResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) GetResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourcePolicies, reflect.TypeOf((*MockRamClient)(nil).GetResourcePolicies), varargs...)
}

func (m *MockRamClient) GetResourceShareAssociations(arg0 context.Context, arg1 *ram.GetResourceShareAssociationsInput, arg2 ...func(*ram.Options)) (*ram.GetResourceShareAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourceShareAssociations, varargs...)
	ret0, _ := ret[0].(*ram.GetResourceShareAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) GetResourceShareAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourceShareAssociations, reflect.TypeOf((*MockRamClient)(nil).GetResourceShareAssociations), varargs...)
}

func (m *MockRamClient) GetResourceShareInvitations(arg0 context.Context, arg1 *ram.GetResourceShareInvitationsInput, arg2 ...func(*ram.Options)) (*ram.GetResourceShareInvitationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourceShareInvitations, varargs...)
	ret0, _ := ret[0].(*ram.GetResourceShareInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) GetResourceShareInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourceShareInvitations, reflect.TypeOf((*MockRamClient)(nil).GetResourceShareInvitations), varargs...)
}

func (m *MockRamClient) GetResourceShares(arg0 context.Context, arg1 *ram.GetResourceSharesInput, arg2 ...func(*ram.Options)) (*ram.GetResourceSharesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourceShares, varargs...)
	ret0, _ := ret[0].(*ram.GetResourceSharesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) GetResourceShares(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourceShares, reflect.TypeOf((*MockRamClient)(nil).GetResourceShares), varargs...)
}

func (m *MockRamClient) ListPendingInvitationResources(arg0 context.Context, arg1 *ram.ListPendingInvitationResourcesInput, arg2 ...func(*ram.Options)) (*ram.ListPendingInvitationResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPendingInvitationResources, varargs...)
	ret0, _ := ret[0].(*ram.ListPendingInvitationResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListPendingInvitationResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPendingInvitationResources, reflect.TypeOf((*MockRamClient)(nil).ListPendingInvitationResources), varargs...)
}

func (m *MockRamClient) ListPermissionVersions(arg0 context.Context, arg1 *ram.ListPermissionVersionsInput, arg2 ...func(*ram.Options)) (*ram.ListPermissionVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPermissionVersions, varargs...)
	ret0, _ := ret[0].(*ram.ListPermissionVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListPermissionVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPermissionVersions, reflect.TypeOf((*MockRamClient)(nil).ListPermissionVersions), varargs...)
}

func (m *MockRamClient) ListPermissions(arg0 context.Context, arg1 *ram.ListPermissionsInput, arg2 ...func(*ram.Options)) (*ram.ListPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPermissions, varargs...)
	ret0, _ := ret[0].(*ram.ListPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPermissions, reflect.TypeOf((*MockRamClient)(nil).ListPermissions), varargs...)
}

func (m *MockRamClient) ListPrincipals(arg0 context.Context, arg1 *ram.ListPrincipalsInput, arg2 ...func(*ram.Options)) (*ram.ListPrincipalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPrincipals, varargs...)
	ret0, _ := ret[0].(*ram.ListPrincipalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListPrincipals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPrincipals, reflect.TypeOf((*MockRamClient)(nil).ListPrincipals), varargs...)
}

func (m *MockRamClient) ListResourceSharePermissions(arg0 context.Context, arg1 *ram.ListResourceSharePermissionsInput, arg2 ...func(*ram.Options)) (*ram.ListResourceSharePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceSharePermissions, varargs...)
	ret0, _ := ret[0].(*ram.ListResourceSharePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListResourceSharePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceSharePermissions, reflect.TypeOf((*MockRamClient)(nil).ListResourceSharePermissions), varargs...)
}

func (m *MockRamClient) ListResourceTypes(arg0 context.Context, arg1 *ram.ListResourceTypesInput, arg2 ...func(*ram.Options)) (*ram.ListResourceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceTypes, varargs...)
	ret0, _ := ret[0].(*ram.ListResourceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListResourceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceTypes, reflect.TypeOf((*MockRamClient)(nil).ListResourceTypes), varargs...)
}

func (m *MockRamClient) ListResources(arg0 context.Context, arg1 *ram.ListResourcesInput, arg2 ...func(*ram.Options)) (*ram.ListResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResources, varargs...)
	ret0, _ := ret[0].(*ram.ListResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRamClientMockRecorder) ListResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResources, reflect.TypeOf((*MockRamClient)(nil).ListResources), varargs...)
}
