package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ssoadmin "github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	gomock "github.com/golang/mock/gomock"
)

type MockSsoadminClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSsoadminClientMockRecorder
}

type MockSsoadminClientMockRecorder struct {
	mock *MockSsoadminClient
}

func NewMockSsoadminClient(ctrl *gomock.Controller) *MockSsoadminClient {
	mock := &MockSsoadminClient{ctrl: ctrl}
	mock.recorder = &MockSsoadminClientMockRecorder{mock}
	return mock
}

func (m *MockSsoadminClient) EXPECT() *MockSsoadminClientMockRecorder {
	return m.recorder
}

func (m *MockSsoadminClient) DescribeAccountAssignmentCreationStatus(arg0 context.Context, arg1 *ssoadmin.DescribeAccountAssignmentCreationStatusInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAssignmentCreationStatus, varargs...)
	ret0, _ := ret[0].(*ssoadmin.DescribeAccountAssignmentCreationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) DescribeAccountAssignmentCreationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAssignmentCreationStatus, reflect.TypeOf((*MockSsoadminClient)(nil).DescribeAccountAssignmentCreationStatus), varargs...)
}

func (m *MockSsoadminClient) DescribeAccountAssignmentDeletionStatus(arg0 context.Context, arg1 *ssoadmin.DescribeAccountAssignmentDeletionStatusInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAssignmentDeletionStatus, varargs...)
	ret0, _ := ret[0].(*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) DescribeAccountAssignmentDeletionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAssignmentDeletionStatus, reflect.TypeOf((*MockSsoadminClient)(nil).DescribeAccountAssignmentDeletionStatus), varargs...)
}

func (m *MockSsoadminClient) DescribeInstanceAccessControlAttributeConfiguration(arg0 context.Context, arg1 *ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceAccessControlAttributeConfiguration, varargs...)
	ret0, _ := ret[0].(*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) DescribeInstanceAccessControlAttributeConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceAccessControlAttributeConfiguration, reflect.TypeOf((*MockSsoadminClient)(nil).DescribeInstanceAccessControlAttributeConfiguration), varargs...)
}

func (m *MockSsoadminClient) DescribePermissionSet(arg0 context.Context, arg1 *ssoadmin.DescribePermissionSetInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.DescribePermissionSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePermissionSet, varargs...)
	ret0, _ := ret[0].(*ssoadmin.DescribePermissionSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) DescribePermissionSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePermissionSet, reflect.TypeOf((*MockSsoadminClient)(nil).DescribePermissionSet), varargs...)
}

func (m *MockSsoadminClient) DescribePermissionSetProvisioningStatus(arg0 context.Context, arg1 *ssoadmin.DescribePermissionSetProvisioningStatusInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePermissionSetProvisioningStatus, varargs...)
	ret0, _ := ret[0].(*ssoadmin.DescribePermissionSetProvisioningStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) DescribePermissionSetProvisioningStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePermissionSetProvisioningStatus, reflect.TypeOf((*MockSsoadminClient)(nil).DescribePermissionSetProvisioningStatus), varargs...)
}

func (m *MockSsoadminClient) GetInlinePolicyForPermissionSet(arg0 context.Context, arg1 *ssoadmin.GetInlinePolicyForPermissionSetInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInlinePolicyForPermissionSet, varargs...)
	ret0, _ := ret[0].(*ssoadmin.GetInlinePolicyForPermissionSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) GetInlinePolicyForPermissionSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInlinePolicyForPermissionSet, reflect.TypeOf((*MockSsoadminClient)(nil).GetInlinePolicyForPermissionSet), varargs...)
}

func (m *MockSsoadminClient) GetPermissionsBoundaryForPermissionSet(arg0 context.Context, arg1 *ssoadmin.GetPermissionsBoundaryForPermissionSetInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPermissionsBoundaryForPermissionSet, varargs...)
	ret0, _ := ret[0].(*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) GetPermissionsBoundaryForPermissionSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPermissionsBoundaryForPermissionSet, reflect.TypeOf((*MockSsoadminClient)(nil).GetPermissionsBoundaryForPermissionSet), varargs...)
}

func (m *MockSsoadminClient) ListAccountAssignmentCreationStatus(arg0 context.Context, arg1 *ssoadmin.ListAccountAssignmentCreationStatusInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountAssignmentCreationStatus, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListAccountAssignmentCreationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListAccountAssignmentCreationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountAssignmentCreationStatus, reflect.TypeOf((*MockSsoadminClient)(nil).ListAccountAssignmentCreationStatus), varargs...)
}

func (m *MockSsoadminClient) ListAccountAssignmentDeletionStatus(arg0 context.Context, arg1 *ssoadmin.ListAccountAssignmentDeletionStatusInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountAssignmentDeletionStatus, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListAccountAssignmentDeletionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListAccountAssignmentDeletionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountAssignmentDeletionStatus, reflect.TypeOf((*MockSsoadminClient)(nil).ListAccountAssignmentDeletionStatus), varargs...)
}

func (m *MockSsoadminClient) ListAccountAssignments(arg0 context.Context, arg1 *ssoadmin.ListAccountAssignmentsInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountAssignments, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListAccountAssignmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListAccountAssignments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountAssignments, reflect.TypeOf((*MockSsoadminClient)(nil).ListAccountAssignments), varargs...)
}

func (m *MockSsoadminClient) ListAccountsForProvisionedPermissionSet(arg0 context.Context, arg1 *ssoadmin.ListAccountsForProvisionedPermissionSetInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountsForProvisionedPermissionSet, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListAccountsForProvisionedPermissionSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountsForProvisionedPermissionSet, reflect.TypeOf((*MockSsoadminClient)(nil).ListAccountsForProvisionedPermissionSet), varargs...)
}

func (m *MockSsoadminClient) ListCustomerManagedPolicyReferencesInPermissionSet(arg0 context.Context, arg1 *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCustomerManagedPolicyReferencesInPermissionSet, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListCustomerManagedPolicyReferencesInPermissionSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCustomerManagedPolicyReferencesInPermissionSet, reflect.TypeOf((*MockSsoadminClient)(nil).ListCustomerManagedPolicyReferencesInPermissionSet), varargs...)
}

func (m *MockSsoadminClient) ListInstances(arg0 context.Context, arg1 *ssoadmin.ListInstancesInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstances, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstances, reflect.TypeOf((*MockSsoadminClient)(nil).ListInstances), varargs...)
}

func (m *MockSsoadminClient) ListManagedPoliciesInPermissionSet(arg0 context.Context, arg1 *ssoadmin.ListManagedPoliciesInPermissionSetInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListManagedPoliciesInPermissionSet, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListManagedPoliciesInPermissionSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListManagedPoliciesInPermissionSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListManagedPoliciesInPermissionSet, reflect.TypeOf((*MockSsoadminClient)(nil).ListManagedPoliciesInPermissionSet), varargs...)
}

func (m *MockSsoadminClient) ListPermissionSetProvisioningStatus(arg0 context.Context, arg1 *ssoadmin.ListPermissionSetProvisioningStatusInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPermissionSetProvisioningStatus, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListPermissionSetProvisioningStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListPermissionSetProvisioningStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPermissionSetProvisioningStatus, reflect.TypeOf((*MockSsoadminClient)(nil).ListPermissionSetProvisioningStatus), varargs...)
}

func (m *MockSsoadminClient) ListPermissionSets(arg0 context.Context, arg1 *ssoadmin.ListPermissionSetsInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPermissionSets, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListPermissionSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListPermissionSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPermissionSets, reflect.TypeOf((*MockSsoadminClient)(nil).ListPermissionSets), varargs...)
}

func (m *MockSsoadminClient) ListPermissionSetsProvisionedToAccount(arg0 context.Context, arg1 *ssoadmin.ListPermissionSetsProvisionedToAccountInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPermissionSetsProvisionedToAccount, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListPermissionSetsProvisionedToAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPermissionSetsProvisionedToAccount, reflect.TypeOf((*MockSsoadminClient)(nil).ListPermissionSetsProvisionedToAccount), varargs...)
}

func (m *MockSsoadminClient) ListTagsForResource(arg0 context.Context, arg1 *ssoadmin.ListTagsForResourceInput, arg2 ...func(*ssoadmin.Options)) (*ssoadmin.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*ssoadmin.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsoadminClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSsoadminClient)(nil).ListTagsForResource), varargs...)
}
