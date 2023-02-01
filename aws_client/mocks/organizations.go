package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	organizations "github.com/aws/aws-sdk-go-v2/service/organizations"
	gomock "github.com/golang/mock/gomock"
)

type MockOrganizationsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockOrganizationsClientMockRecorder
}

type MockOrganizationsClientMockRecorder struct {
	mock *MockOrganizationsClient
}

func NewMockOrganizationsClient(ctrl *gomock.Controller) *MockOrganizationsClient {
	mock := &MockOrganizationsClient{ctrl: ctrl}
	mock.recorder = &MockOrganizationsClientMockRecorder{mock}
	return mock
}

func (m *MockOrganizationsClient) EXPECT() *MockOrganizationsClientMockRecorder {
	return m.recorder
}

func (m *MockOrganizationsClient) DescribeAccount(arg0 context.Context, arg1 *organizations.DescribeAccountInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccount, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccount, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeAccount), varargs...)
}

func (m *MockOrganizationsClient) DescribeCreateAccountStatus(arg0 context.Context, arg1 *organizations.DescribeCreateAccountStatusInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeCreateAccountStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCreateAccountStatus, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeCreateAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeCreateAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCreateAccountStatus, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeCreateAccountStatus), varargs...)
}

func (m *MockOrganizationsClient) DescribeEffectivePolicy(arg0 context.Context, arg1 *organizations.DescribeEffectivePolicyInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeEffectivePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEffectivePolicy, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeEffectivePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeEffectivePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEffectivePolicy, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeEffectivePolicy), varargs...)
}

func (m *MockOrganizationsClient) DescribeHandshake(arg0 context.Context, arg1 *organizations.DescribeHandshakeInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeHandshakeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHandshake, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeHandshakeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeHandshake(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHandshake, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeHandshake), varargs...)
}

func (m *MockOrganizationsClient) DescribeOrganization(arg0 context.Context, arg1 *organizations.DescribeOrganizationInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeOrganizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganization, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganization, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeOrganization), varargs...)
}

func (m *MockOrganizationsClient) DescribeOrganizationalUnit(arg0 context.Context, arg1 *organizations.DescribeOrganizationalUnitInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeOrganizationalUnitOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationalUnit, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeOrganizationalUnitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeOrganizationalUnit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationalUnit, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeOrganizationalUnit), varargs...)
}

func (m *MockOrganizationsClient) DescribePolicy(arg0 context.Context, arg1 *organizations.DescribePolicyInput, arg2 ...func(*organizations.Options)) (*organizations.DescribePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePolicy, varargs...)
	ret0, _ := ret[0].(*organizations.DescribePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePolicy, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribePolicy), varargs...)
}

func (m *MockOrganizationsClient) DescribeResourcePolicy(arg0 context.Context, arg1 *organizations.DescribeResourcePolicyInput, arg2 ...func(*organizations.Options)) (*organizations.DescribeResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeResourcePolicy, varargs...)
	ret0, _ := ret[0].(*organizations.DescribeResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) DescribeResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeResourcePolicy, reflect.TypeOf((*MockOrganizationsClient)(nil).DescribeResourcePolicy), varargs...)
}

func (m *MockOrganizationsClient) ListAWSServiceAccessForOrganization(arg0 context.Context, arg1 *organizations.ListAWSServiceAccessForOrganizationInput, arg2 ...func(*organizations.Options)) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAWSServiceAccessForOrganization, varargs...)
	ret0, _ := ret[0].(*organizations.ListAWSServiceAccessForOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListAWSServiceAccessForOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAWSServiceAccessForOrganization, reflect.TypeOf((*MockOrganizationsClient)(nil).ListAWSServiceAccessForOrganization), varargs...)
}

func (m *MockOrganizationsClient) ListAccounts(arg0 context.Context, arg1 *organizations.ListAccountsInput, arg2 ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccounts, varargs...)
	ret0, _ := ret[0].(*organizations.ListAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccounts, reflect.TypeOf((*MockOrganizationsClient)(nil).ListAccounts), varargs...)
}

func (m *MockOrganizationsClient) ListAccountsForParent(arg0 context.Context, arg1 *organizations.ListAccountsForParentInput, arg2 ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountsForParent, varargs...)
	ret0, _ := ret[0].(*organizations.ListAccountsForParentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListAccountsForParent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountsForParent, reflect.TypeOf((*MockOrganizationsClient)(nil).ListAccountsForParent), varargs...)
}

func (m *MockOrganizationsClient) ListChildren(arg0 context.Context, arg1 *organizations.ListChildrenInput, arg2 ...func(*organizations.Options)) (*organizations.ListChildrenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListChildren, varargs...)
	ret0, _ := ret[0].(*organizations.ListChildrenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListChildren(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListChildren, reflect.TypeOf((*MockOrganizationsClient)(nil).ListChildren), varargs...)
}

func (m *MockOrganizationsClient) ListCreateAccountStatus(arg0 context.Context, arg1 *organizations.ListCreateAccountStatusInput, arg2 ...func(*organizations.Options)) (*organizations.ListCreateAccountStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCreateAccountStatus, varargs...)
	ret0, _ := ret[0].(*organizations.ListCreateAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListCreateAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCreateAccountStatus, reflect.TypeOf((*MockOrganizationsClient)(nil).ListCreateAccountStatus), varargs...)
}

func (m *MockOrganizationsClient) ListDelegatedAdministrators(arg0 context.Context, arg1 *organizations.ListDelegatedAdministratorsInput, arg2 ...func(*organizations.Options)) (*organizations.ListDelegatedAdministratorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDelegatedAdministrators, varargs...)
	ret0, _ := ret[0].(*organizations.ListDelegatedAdministratorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListDelegatedAdministrators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDelegatedAdministrators, reflect.TypeOf((*MockOrganizationsClient)(nil).ListDelegatedAdministrators), varargs...)
}

func (m *MockOrganizationsClient) ListDelegatedServicesForAccount(arg0 context.Context, arg1 *organizations.ListDelegatedServicesForAccountInput, arg2 ...func(*organizations.Options)) (*organizations.ListDelegatedServicesForAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDelegatedServicesForAccount, varargs...)
	ret0, _ := ret[0].(*organizations.ListDelegatedServicesForAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListDelegatedServicesForAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDelegatedServicesForAccount, reflect.TypeOf((*MockOrganizationsClient)(nil).ListDelegatedServicesForAccount), varargs...)
}

func (m *MockOrganizationsClient) ListHandshakesForAccount(arg0 context.Context, arg1 *organizations.ListHandshakesForAccountInput, arg2 ...func(*organizations.Options)) (*organizations.ListHandshakesForAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHandshakesForAccount, varargs...)
	ret0, _ := ret[0].(*organizations.ListHandshakesForAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListHandshakesForAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHandshakesForAccount, reflect.TypeOf((*MockOrganizationsClient)(nil).ListHandshakesForAccount), varargs...)
}

func (m *MockOrganizationsClient) ListHandshakesForOrganization(arg0 context.Context, arg1 *organizations.ListHandshakesForOrganizationInput, arg2 ...func(*organizations.Options)) (*organizations.ListHandshakesForOrganizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHandshakesForOrganization, varargs...)
	ret0, _ := ret[0].(*organizations.ListHandshakesForOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListHandshakesForOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHandshakesForOrganization, reflect.TypeOf((*MockOrganizationsClient)(nil).ListHandshakesForOrganization), varargs...)
}

func (m *MockOrganizationsClient) ListOrganizationalUnitsForParent(arg0 context.Context, arg1 *organizations.ListOrganizationalUnitsForParentInput, arg2 ...func(*organizations.Options)) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOrganizationalUnitsForParent, varargs...)
	ret0, _ := ret[0].(*organizations.ListOrganizationalUnitsForParentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListOrganizationalUnitsForParent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOrganizationalUnitsForParent, reflect.TypeOf((*MockOrganizationsClient)(nil).ListOrganizationalUnitsForParent), varargs...)
}

func (m *MockOrganizationsClient) ListParents(arg0 context.Context, arg1 *organizations.ListParentsInput, arg2 ...func(*organizations.Options)) (*organizations.ListParentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListParents, varargs...)
	ret0, _ := ret[0].(*organizations.ListParentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListParents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListParents, reflect.TypeOf((*MockOrganizationsClient)(nil).ListParents), varargs...)
}

func (m *MockOrganizationsClient) ListPolicies(arg0 context.Context, arg1 *organizations.ListPoliciesInput, arg2 ...func(*organizations.Options)) (*organizations.ListPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicies, varargs...)
	ret0, _ := ret[0].(*organizations.ListPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicies, reflect.TypeOf((*MockOrganizationsClient)(nil).ListPolicies), varargs...)
}

func (m *MockOrganizationsClient) ListPoliciesForTarget(arg0 context.Context, arg1 *organizations.ListPoliciesForTargetInput, arg2 ...func(*organizations.Options)) (*organizations.ListPoliciesForTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPoliciesForTarget, varargs...)
	ret0, _ := ret[0].(*organizations.ListPoliciesForTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListPoliciesForTarget(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPoliciesForTarget, reflect.TypeOf((*MockOrganizationsClient)(nil).ListPoliciesForTarget), varargs...)
}

func (m *MockOrganizationsClient) ListRoots(arg0 context.Context, arg1 *organizations.ListRootsInput, arg2 ...func(*organizations.Options)) (*organizations.ListRootsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRoots, varargs...)
	ret0, _ := ret[0].(*organizations.ListRootsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListRoots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRoots, reflect.TypeOf((*MockOrganizationsClient)(nil).ListRoots), varargs...)
}

func (m *MockOrganizationsClient) ListTagsForResource(arg0 context.Context, arg1 *organizations.ListTagsForResourceInput, arg2 ...func(*organizations.Options)) (*organizations.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*organizations.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockOrganizationsClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockOrganizationsClient) ListTargetsForPolicy(arg0 context.Context, arg1 *organizations.ListTargetsForPolicyInput, arg2 ...func(*organizations.Options)) (*organizations.ListTargetsForPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTargetsForPolicy, varargs...)
	ret0, _ := ret[0].(*organizations.ListTargetsForPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListTargetsForPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTargetsForPolicy, reflect.TypeOf((*MockOrganizationsClient)(nil).ListTargetsForPolicy), varargs...)
}
