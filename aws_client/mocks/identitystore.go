package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	identitystore "github.com/aws/aws-sdk-go-v2/service/identitystore"
	gomock "github.com/golang/mock/gomock"
)

type MockIdentitystoreClient struct {
	ctrl		*gomock.Controller
	recorder	*MockIdentitystoreClientMockRecorder
}

type MockIdentitystoreClientMockRecorder struct {
	mock *MockIdentitystoreClient
}

func NewMockIdentitystoreClient(ctrl *gomock.Controller) *MockIdentitystoreClient {
	mock := &MockIdentitystoreClient{ctrl: ctrl}
	mock.recorder = &MockIdentitystoreClientMockRecorder{mock}
	return mock
}

func (m *MockIdentitystoreClient) EXPECT() *MockIdentitystoreClientMockRecorder {
	return m.recorder
}

func (m *MockIdentitystoreClient) DescribeGroup(arg0 context.Context, arg1 *identitystore.DescribeGroupInput, arg2 ...func(*identitystore.Options)) (*identitystore.DescribeGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGroup, varargs...)
	ret0, _ := ret[0].(*identitystore.DescribeGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) DescribeGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGroup, reflect.TypeOf((*MockIdentitystoreClient)(nil).DescribeGroup), varargs...)
}

func (m *MockIdentitystoreClient) DescribeGroupMembership(arg0 context.Context, arg1 *identitystore.DescribeGroupMembershipInput, arg2 ...func(*identitystore.Options)) (*identitystore.DescribeGroupMembershipOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGroupMembership, varargs...)
	ret0, _ := ret[0].(*identitystore.DescribeGroupMembershipOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) DescribeGroupMembership(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGroupMembership, reflect.TypeOf((*MockIdentitystoreClient)(nil).DescribeGroupMembership), varargs...)
}

func (m *MockIdentitystoreClient) DescribeUser(arg0 context.Context, arg1 *identitystore.DescribeUserInput, arg2 ...func(*identitystore.Options)) (*identitystore.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUser, varargs...)
	ret0, _ := ret[0].(*identitystore.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) DescribeUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUser, reflect.TypeOf((*MockIdentitystoreClient)(nil).DescribeUser), varargs...)
}

func (m *MockIdentitystoreClient) GetGroupId(arg0 context.Context, arg1 *identitystore.GetGroupIdInput, arg2 ...func(*identitystore.Options)) (*identitystore.GetGroupIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroupId, varargs...)
	ret0, _ := ret[0].(*identitystore.GetGroupIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) GetGroupId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroupId, reflect.TypeOf((*MockIdentitystoreClient)(nil).GetGroupId), varargs...)
}

func (m *MockIdentitystoreClient) GetGroupMembershipId(arg0 context.Context, arg1 *identitystore.GetGroupMembershipIdInput, arg2 ...func(*identitystore.Options)) (*identitystore.GetGroupMembershipIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroupMembershipId, varargs...)
	ret0, _ := ret[0].(*identitystore.GetGroupMembershipIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) GetGroupMembershipId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroupMembershipId, reflect.TypeOf((*MockIdentitystoreClient)(nil).GetGroupMembershipId), varargs...)
}

func (m *MockIdentitystoreClient) GetUserId(arg0 context.Context, arg1 *identitystore.GetUserIdInput, arg2 ...func(*identitystore.Options)) (*identitystore.GetUserIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUserId, varargs...)
	ret0, _ := ret[0].(*identitystore.GetUserIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) GetUserId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUserId, reflect.TypeOf((*MockIdentitystoreClient)(nil).GetUserId), varargs...)
}

func (m *MockIdentitystoreClient) ListGroupMemberships(arg0 context.Context, arg1 *identitystore.ListGroupMembershipsInput, arg2 ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroupMemberships, varargs...)
	ret0, _ := ret[0].(*identitystore.ListGroupMembershipsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) ListGroupMemberships(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroupMemberships, reflect.TypeOf((*MockIdentitystoreClient)(nil).ListGroupMemberships), varargs...)
}

func (m *MockIdentitystoreClient) ListGroupMembershipsForMember(arg0 context.Context, arg1 *identitystore.ListGroupMembershipsForMemberInput, arg2 ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsForMemberOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroupMembershipsForMember, varargs...)
	ret0, _ := ret[0].(*identitystore.ListGroupMembershipsForMemberOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) ListGroupMembershipsForMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroupMembershipsForMember, reflect.TypeOf((*MockIdentitystoreClient)(nil).ListGroupMembershipsForMember), varargs...)
}

func (m *MockIdentitystoreClient) ListGroups(arg0 context.Context, arg1 *identitystore.ListGroupsInput, arg2 ...func(*identitystore.Options)) (*identitystore.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroups, varargs...)
	ret0, _ := ret[0].(*identitystore.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroups, reflect.TypeOf((*MockIdentitystoreClient)(nil).ListGroups), varargs...)
}

func (m *MockIdentitystoreClient) ListUsers(arg0 context.Context, arg1 *identitystore.ListUsersInput, arg2 ...func(*identitystore.Options)) (*identitystore.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsers, varargs...)
	ret0, _ := ret[0].(*identitystore.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIdentitystoreClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsers, reflect.TypeOf((*MockIdentitystoreClient)(nil).ListUsers), varargs...)
}
