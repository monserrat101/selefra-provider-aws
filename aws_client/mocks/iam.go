package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	iam "github.com/aws/aws-sdk-go-v2/service/iam"
	gomock "github.com/golang/mock/gomock"
)

type MockIamClient struct {
	ctrl		*gomock.Controller
	recorder	*MockIamClientMockRecorder
}

type MockIamClientMockRecorder struct {
	mock *MockIamClient
}

func NewMockIamClient(ctrl *gomock.Controller) *MockIamClient {
	mock := &MockIamClient{ctrl: ctrl}
	mock.recorder = &MockIamClientMockRecorder{mock}
	return mock
}

func (m *MockIamClient) EXPECT() *MockIamClientMockRecorder {
	return m.recorder
}

func (m *MockIamClient) GenerateCredentialReport(arg0 context.Context, arg1 *iam.GenerateCredentialReportInput, arg2 ...func(*iam.Options)) (*iam.GenerateCredentialReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GenerateCredentialReport, varargs...)
	ret0, _ := ret[0].(*iam.GenerateCredentialReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GenerateCredentialReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GenerateCredentialReport, reflect.TypeOf((*MockIamClient)(nil).GenerateCredentialReport), varargs...)
}

func (m *MockIamClient) GetAccessKeyLastUsed(arg0 context.Context, arg1 *iam.GetAccessKeyLastUsedInput, arg2 ...func(*iam.Options)) (*iam.GetAccessKeyLastUsedOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessKeyLastUsed, varargs...)
	ret0, _ := ret[0].(*iam.GetAccessKeyLastUsedOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccessKeyLastUsed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessKeyLastUsed, reflect.TypeOf((*MockIamClient)(nil).GetAccessKeyLastUsed), varargs...)
}

func (m *MockIamClient) GetAccountAuthorizationDetails(arg0 context.Context, arg1 *iam.GetAccountAuthorizationDetailsInput, arg2 ...func(*iam.Options)) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountAuthorizationDetails, varargs...)
	ret0, _ := ret[0].(*iam.GetAccountAuthorizationDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccountAuthorizationDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountAuthorizationDetails, reflect.TypeOf((*MockIamClient)(nil).GetAccountAuthorizationDetails), varargs...)
}

func (m *MockIamClient) GetAccountPasswordPolicy(arg0 context.Context, arg1 *iam.GetAccountPasswordPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountPasswordPolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetAccountPasswordPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccountPasswordPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountPasswordPolicy, reflect.TypeOf((*MockIamClient)(nil).GetAccountPasswordPolicy), varargs...)
}

func (m *MockIamClient) GetAccountSummary(arg0 context.Context, arg1 *iam.GetAccountSummaryInput, arg2 ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountSummary, varargs...)
	ret0, _ := ret[0].(*iam.GetAccountSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccountSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountSummary, reflect.TypeOf((*MockIamClient)(nil).GetAccountSummary), varargs...)
}

func (m *MockIamClient) GetContextKeysForCustomPolicy(arg0 context.Context, arg1 *iam.GetContextKeysForCustomPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetContextKeysForCustomPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContextKeysForCustomPolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetContextKeysForCustomPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetContextKeysForCustomPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContextKeysForCustomPolicy, reflect.TypeOf((*MockIamClient)(nil).GetContextKeysForCustomPolicy), varargs...)
}

func (m *MockIamClient) GetContextKeysForPrincipalPolicy(arg0 context.Context, arg1 *iam.GetContextKeysForPrincipalPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetContextKeysForPrincipalPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContextKeysForPrincipalPolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetContextKeysForPrincipalPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetContextKeysForPrincipalPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContextKeysForPrincipalPolicy, reflect.TypeOf((*MockIamClient)(nil).GetContextKeysForPrincipalPolicy), varargs...)
}

func (m *MockIamClient) GetCredentialReport(arg0 context.Context, arg1 *iam.GetCredentialReportInput, arg2 ...func(*iam.Options)) (*iam.GetCredentialReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCredentialReport, varargs...)
	ret0, _ := ret[0].(*iam.GetCredentialReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetCredentialReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCredentialReport, reflect.TypeOf((*MockIamClient)(nil).GetCredentialReport), varargs...)
}

func (m *MockIamClient) GetGroup(arg0 context.Context, arg1 *iam.GetGroupInput, arg2 ...func(*iam.Options)) (*iam.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroup, varargs...)
	ret0, _ := ret[0].(*iam.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroup, reflect.TypeOf((*MockIamClient)(nil).GetGroup), varargs...)
}

func (m *MockIamClient) GetGroupPolicy(arg0 context.Context, arg1 *iam.GetGroupPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetGroupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroupPolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetGroupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetGroupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroupPolicy, reflect.TypeOf((*MockIamClient)(nil).GetGroupPolicy), varargs...)
}

func (m *MockIamClient) GetInstanceProfile(arg0 context.Context, arg1 *iam.GetInstanceProfileInput, arg2 ...func(*iam.Options)) (*iam.GetInstanceProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceProfile, varargs...)
	ret0, _ := ret[0].(*iam.GetInstanceProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetInstanceProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceProfile, reflect.TypeOf((*MockIamClient)(nil).GetInstanceProfile), varargs...)
}

func (m *MockIamClient) GetLoginProfile(arg0 context.Context, arg1 *iam.GetLoginProfileInput, arg2 ...func(*iam.Options)) (*iam.GetLoginProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoginProfile, varargs...)
	ret0, _ := ret[0].(*iam.GetLoginProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetLoginProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoginProfile, reflect.TypeOf((*MockIamClient)(nil).GetLoginProfile), varargs...)
}

func (m *MockIamClient) GetOpenIDConnectProvider(arg0 context.Context, arg1 *iam.GetOpenIDConnectProviderInput, arg2 ...func(*iam.Options)) (*iam.GetOpenIDConnectProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOpenIDConnectProvider, varargs...)
	ret0, _ := ret[0].(*iam.GetOpenIDConnectProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetOpenIDConnectProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOpenIDConnectProvider, reflect.TypeOf((*MockIamClient)(nil).GetOpenIDConnectProvider), varargs...)
}

func (m *MockIamClient) GetOrganizationsAccessReport(arg0 context.Context, arg1 *iam.GetOrganizationsAccessReportInput, arg2 ...func(*iam.Options)) (*iam.GetOrganizationsAccessReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOrganizationsAccessReport, varargs...)
	ret0, _ := ret[0].(*iam.GetOrganizationsAccessReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetOrganizationsAccessReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOrganizationsAccessReport, reflect.TypeOf((*MockIamClient)(nil).GetOrganizationsAccessReport), varargs...)
}

func (m *MockIamClient) GetPolicy(arg0 context.Context, arg1 *iam.GetPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPolicy, reflect.TypeOf((*MockIamClient)(nil).GetPolicy), varargs...)
}

func (m *MockIamClient) GetPolicyVersion(arg0 context.Context, arg1 *iam.GetPolicyVersionInput, arg2 ...func(*iam.Options)) (*iam.GetPolicyVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPolicyVersion, varargs...)
	ret0, _ := ret[0].(*iam.GetPolicyVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetPolicyVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPolicyVersion, reflect.TypeOf((*MockIamClient)(nil).GetPolicyVersion), varargs...)
}

func (m *MockIamClient) GetRole(arg0 context.Context, arg1 *iam.GetRoleInput, arg2 ...func(*iam.Options)) (*iam.GetRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRole, varargs...)
	ret0, _ := ret[0].(*iam.GetRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRole, reflect.TypeOf((*MockIamClient)(nil).GetRole), varargs...)
}

func (m *MockIamClient) GetRolePolicy(arg0 context.Context, arg1 *iam.GetRolePolicyInput, arg2 ...func(*iam.Options)) (*iam.GetRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRolePolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetRolePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRolePolicy, reflect.TypeOf((*MockIamClient)(nil).GetRolePolicy), varargs...)
}

func (m *MockIamClient) GetSAMLProvider(arg0 context.Context, arg1 *iam.GetSAMLProviderInput, arg2 ...func(*iam.Options)) (*iam.GetSAMLProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSAMLProvider, varargs...)
	ret0, _ := ret[0].(*iam.GetSAMLProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetSAMLProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSAMLProvider, reflect.TypeOf((*MockIamClient)(nil).GetSAMLProvider), varargs...)
}

func (m *MockIamClient) GetSSHPublicKey(arg0 context.Context, arg1 *iam.GetSSHPublicKeyInput, arg2 ...func(*iam.Options)) (*iam.GetSSHPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSSHPublicKey, varargs...)
	ret0, _ := ret[0].(*iam.GetSSHPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetSSHPublicKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSSHPublicKey, reflect.TypeOf((*MockIamClient)(nil).GetSSHPublicKey), varargs...)
}

func (m *MockIamClient) GetServerCertificate(arg0 context.Context, arg1 *iam.GetServerCertificateInput, arg2 ...func(*iam.Options)) (*iam.GetServerCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServerCertificate, varargs...)
	ret0, _ := ret[0].(*iam.GetServerCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetServerCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServerCertificate, reflect.TypeOf((*MockIamClient)(nil).GetServerCertificate), varargs...)
}

func (m *MockIamClient) GetServiceLastAccessedDetails(arg0 context.Context, arg1 *iam.GetServiceLastAccessedDetailsInput, arg2 ...func(*iam.Options)) (*iam.GetServiceLastAccessedDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceLastAccessedDetails, varargs...)
	ret0, _ := ret[0].(*iam.GetServiceLastAccessedDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetServiceLastAccessedDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceLastAccessedDetails, reflect.TypeOf((*MockIamClient)(nil).GetServiceLastAccessedDetails), varargs...)
}

func (m *MockIamClient) GetServiceLastAccessedDetailsWithEntities(arg0 context.Context, arg1 *iam.GetServiceLastAccessedDetailsWithEntitiesInput, arg2 ...func(*iam.Options)) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceLastAccessedDetailsWithEntities, varargs...)
	ret0, _ := ret[0].(*iam.GetServiceLastAccessedDetailsWithEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetServiceLastAccessedDetailsWithEntities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceLastAccessedDetailsWithEntities, reflect.TypeOf((*MockIamClient)(nil).GetServiceLastAccessedDetailsWithEntities), varargs...)
}

func (m *MockIamClient) GetServiceLinkedRoleDeletionStatus(arg0 context.Context, arg1 *iam.GetServiceLinkedRoleDeletionStatusInput, arg2 ...func(*iam.Options)) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceLinkedRoleDeletionStatus, varargs...)
	ret0, _ := ret[0].(*iam.GetServiceLinkedRoleDeletionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetServiceLinkedRoleDeletionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceLinkedRoleDeletionStatus, reflect.TypeOf((*MockIamClient)(nil).GetServiceLinkedRoleDeletionStatus), varargs...)
}

func (m *MockIamClient) GetUser(arg0 context.Context, arg1 *iam.GetUserInput, arg2 ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUser, varargs...)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUser, reflect.TypeOf((*MockIamClient)(nil).GetUser), varargs...)
}

func (m *MockIamClient) GetUserPolicy(arg0 context.Context, arg1 *iam.GetUserPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUserPolicy, varargs...)
	ret0, _ := ret[0].(*iam.GetUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetUserPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUserPolicy, reflect.TypeOf((*MockIamClient)(nil).GetUserPolicy), varargs...)
}

func (m *MockIamClient) ListAccessKeys(arg0 context.Context, arg1 *iam.ListAccessKeysInput, arg2 ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccessKeys, varargs...)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAccessKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccessKeys, reflect.TypeOf((*MockIamClient)(nil).ListAccessKeys), varargs...)
}

func (m *MockIamClient) ListAccountAliases(arg0 context.Context, arg1 *iam.ListAccountAliasesInput, arg2 ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountAliases, varargs...)
	ret0, _ := ret[0].(*iam.ListAccountAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAccountAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountAliases, reflect.TypeOf((*MockIamClient)(nil).ListAccountAliases), varargs...)
}

func (m *MockIamClient) ListAttachedGroupPolicies(arg0 context.Context, arg1 *iam.ListAttachedGroupPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedGroupPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttachedGroupPolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedGroupPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAttachedGroupPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttachedGroupPolicies, reflect.TypeOf((*MockIamClient)(nil).ListAttachedGroupPolicies), varargs...)
}

func (m *MockIamClient) ListAttachedRolePolicies(arg0 context.Context, arg1 *iam.ListAttachedRolePoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttachedRolePolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAttachedRolePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttachedRolePolicies, reflect.TypeOf((*MockIamClient)(nil).ListAttachedRolePolicies), varargs...)
}

func (m *MockIamClient) ListAttachedUserPolicies(arg0 context.Context, arg1 *iam.ListAttachedUserPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttachedUserPolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAttachedUserPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttachedUserPolicies, reflect.TypeOf((*MockIamClient)(nil).ListAttachedUserPolicies), varargs...)
}

func (m *MockIamClient) ListEntitiesForPolicy(arg0 context.Context, arg1 *iam.ListEntitiesForPolicyInput, arg2 ...func(*iam.Options)) (*iam.ListEntitiesForPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEntitiesForPolicy, varargs...)
	ret0, _ := ret[0].(*iam.ListEntitiesForPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListEntitiesForPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEntitiesForPolicy, reflect.TypeOf((*MockIamClient)(nil).ListEntitiesForPolicy), varargs...)
}

func (m *MockIamClient) ListGroupPolicies(arg0 context.Context, arg1 *iam.ListGroupPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListGroupPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroupPolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListGroupPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListGroupPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroupPolicies, reflect.TypeOf((*MockIamClient)(nil).ListGroupPolicies), varargs...)
}

func (m *MockIamClient) ListGroups(arg0 context.Context, arg1 *iam.ListGroupsInput, arg2 ...func(*iam.Options)) (*iam.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroups, varargs...)
	ret0, _ := ret[0].(*iam.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroups, reflect.TypeOf((*MockIamClient)(nil).ListGroups), varargs...)
}

func (m *MockIamClient) ListGroupsForUser(arg0 context.Context, arg1 *iam.ListGroupsForUserInput, arg2 ...func(*iam.Options)) (*iam.ListGroupsForUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroupsForUser, varargs...)
	ret0, _ := ret[0].(*iam.ListGroupsForUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListGroupsForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroupsForUser, reflect.TypeOf((*MockIamClient)(nil).ListGroupsForUser), varargs...)
}

func (m *MockIamClient) ListInstanceProfileTags(arg0 context.Context, arg1 *iam.ListInstanceProfileTagsInput, arg2 ...func(*iam.Options)) (*iam.ListInstanceProfileTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstanceProfileTags, varargs...)
	ret0, _ := ret[0].(*iam.ListInstanceProfileTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListInstanceProfileTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstanceProfileTags, reflect.TypeOf((*MockIamClient)(nil).ListInstanceProfileTags), varargs...)
}

func (m *MockIamClient) ListInstanceProfiles(arg0 context.Context, arg1 *iam.ListInstanceProfilesInput, arg2 ...func(*iam.Options)) (*iam.ListInstanceProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstanceProfiles, varargs...)
	ret0, _ := ret[0].(*iam.ListInstanceProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListInstanceProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstanceProfiles, reflect.TypeOf((*MockIamClient)(nil).ListInstanceProfiles), varargs...)
}

func (m *MockIamClient) ListInstanceProfilesForRole(arg0 context.Context, arg1 *iam.ListInstanceProfilesForRoleInput, arg2 ...func(*iam.Options)) (*iam.ListInstanceProfilesForRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstanceProfilesForRole, varargs...)
	ret0, _ := ret[0].(*iam.ListInstanceProfilesForRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListInstanceProfilesForRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstanceProfilesForRole, reflect.TypeOf((*MockIamClient)(nil).ListInstanceProfilesForRole), varargs...)
}

func (m *MockIamClient) ListMFADeviceTags(arg0 context.Context, arg1 *iam.ListMFADeviceTagsInput, arg2 ...func(*iam.Options)) (*iam.ListMFADeviceTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMFADeviceTags, varargs...)
	ret0, _ := ret[0].(*iam.ListMFADeviceTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListMFADeviceTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMFADeviceTags, reflect.TypeOf((*MockIamClient)(nil).ListMFADeviceTags), varargs...)
}

func (m *MockIamClient) ListMFADevices(arg0 context.Context, arg1 *iam.ListMFADevicesInput, arg2 ...func(*iam.Options)) (*iam.ListMFADevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMFADevices, varargs...)
	ret0, _ := ret[0].(*iam.ListMFADevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListMFADevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMFADevices, reflect.TypeOf((*MockIamClient)(nil).ListMFADevices), varargs...)
}

func (m *MockIamClient) ListOpenIDConnectProviderTags(arg0 context.Context, arg1 *iam.ListOpenIDConnectProviderTagsInput, arg2 ...func(*iam.Options)) (*iam.ListOpenIDConnectProviderTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOpenIDConnectProviderTags, varargs...)
	ret0, _ := ret[0].(*iam.ListOpenIDConnectProviderTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListOpenIDConnectProviderTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOpenIDConnectProviderTags, reflect.TypeOf((*MockIamClient)(nil).ListOpenIDConnectProviderTags), varargs...)
}

func (m *MockIamClient) ListOpenIDConnectProviders(arg0 context.Context, arg1 *iam.ListOpenIDConnectProvidersInput, arg2 ...func(*iam.Options)) (*iam.ListOpenIDConnectProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOpenIDConnectProviders, varargs...)
	ret0, _ := ret[0].(*iam.ListOpenIDConnectProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListOpenIDConnectProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOpenIDConnectProviders, reflect.TypeOf((*MockIamClient)(nil).ListOpenIDConnectProviders), varargs...)
}

func (m *MockIamClient) ListPolicies(arg0 context.Context, arg1 *iam.ListPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicies, reflect.TypeOf((*MockIamClient)(nil).ListPolicies), varargs...)
}

func (m *MockIamClient) ListPoliciesGrantingServiceAccess(arg0 context.Context, arg1 *iam.ListPoliciesGrantingServiceAccessInput, arg2 ...func(*iam.Options)) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPoliciesGrantingServiceAccess, varargs...)
	ret0, _ := ret[0].(*iam.ListPoliciesGrantingServiceAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListPoliciesGrantingServiceAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPoliciesGrantingServiceAccess, reflect.TypeOf((*MockIamClient)(nil).ListPoliciesGrantingServiceAccess), varargs...)
}

func (m *MockIamClient) ListPolicyTags(arg0 context.Context, arg1 *iam.ListPolicyTagsInput, arg2 ...func(*iam.Options)) (*iam.ListPolicyTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicyTags, varargs...)
	ret0, _ := ret[0].(*iam.ListPolicyTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListPolicyTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicyTags, reflect.TypeOf((*MockIamClient)(nil).ListPolicyTags), varargs...)
}

func (m *MockIamClient) ListPolicyVersions(arg0 context.Context, arg1 *iam.ListPolicyVersionsInput, arg2 ...func(*iam.Options)) (*iam.ListPolicyVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicyVersions, varargs...)
	ret0, _ := ret[0].(*iam.ListPolicyVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListPolicyVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicyVersions, reflect.TypeOf((*MockIamClient)(nil).ListPolicyVersions), varargs...)
}

func (m *MockIamClient) ListRolePolicies(arg0 context.Context, arg1 *iam.ListRolePoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRolePolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListRolePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRolePolicies, reflect.TypeOf((*MockIamClient)(nil).ListRolePolicies), varargs...)
}

func (m *MockIamClient) ListRoleTags(arg0 context.Context, arg1 *iam.ListRoleTagsInput, arg2 ...func(*iam.Options)) (*iam.ListRoleTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRoleTags, varargs...)
	ret0, _ := ret[0].(*iam.ListRoleTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListRoleTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRoleTags, reflect.TypeOf((*MockIamClient)(nil).ListRoleTags), varargs...)
}

func (m *MockIamClient) ListRoles(arg0 context.Context, arg1 *iam.ListRolesInput, arg2 ...func(*iam.Options)) (*iam.ListRolesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRoles, varargs...)
	ret0, _ := ret[0].(*iam.ListRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRoles, reflect.TypeOf((*MockIamClient)(nil).ListRoles), varargs...)
}

func (m *MockIamClient) ListSAMLProviderTags(arg0 context.Context, arg1 *iam.ListSAMLProviderTagsInput, arg2 ...func(*iam.Options)) (*iam.ListSAMLProviderTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSAMLProviderTags, varargs...)
	ret0, _ := ret[0].(*iam.ListSAMLProviderTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListSAMLProviderTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSAMLProviderTags, reflect.TypeOf((*MockIamClient)(nil).ListSAMLProviderTags), varargs...)
}

func (m *MockIamClient) ListSAMLProviders(arg0 context.Context, arg1 *iam.ListSAMLProvidersInput, arg2 ...func(*iam.Options)) (*iam.ListSAMLProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSAMLProviders, varargs...)
	ret0, _ := ret[0].(*iam.ListSAMLProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListSAMLProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSAMLProviders, reflect.TypeOf((*MockIamClient)(nil).ListSAMLProviders), varargs...)
}

func (m *MockIamClient) ListSSHPublicKeys(arg0 context.Context, arg1 *iam.ListSSHPublicKeysInput, arg2 ...func(*iam.Options)) (*iam.ListSSHPublicKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSSHPublicKeys, varargs...)
	ret0, _ := ret[0].(*iam.ListSSHPublicKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListSSHPublicKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSSHPublicKeys, reflect.TypeOf((*MockIamClient)(nil).ListSSHPublicKeys), varargs...)
}

func (m *MockIamClient) ListServerCertificateTags(arg0 context.Context, arg1 *iam.ListServerCertificateTagsInput, arg2 ...func(*iam.Options)) (*iam.ListServerCertificateTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServerCertificateTags, varargs...)
	ret0, _ := ret[0].(*iam.ListServerCertificateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListServerCertificateTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServerCertificateTags, reflect.TypeOf((*MockIamClient)(nil).ListServerCertificateTags), varargs...)
}

func (m *MockIamClient) ListServerCertificates(arg0 context.Context, arg1 *iam.ListServerCertificatesInput, arg2 ...func(*iam.Options)) (*iam.ListServerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServerCertificates, varargs...)
	ret0, _ := ret[0].(*iam.ListServerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListServerCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServerCertificates, reflect.TypeOf((*MockIamClient)(nil).ListServerCertificates), varargs...)
}

func (m *MockIamClient) ListServiceSpecificCredentials(arg0 context.Context, arg1 *iam.ListServiceSpecificCredentialsInput, arg2 ...func(*iam.Options)) (*iam.ListServiceSpecificCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServiceSpecificCredentials, varargs...)
	ret0, _ := ret[0].(*iam.ListServiceSpecificCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListServiceSpecificCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServiceSpecificCredentials, reflect.TypeOf((*MockIamClient)(nil).ListServiceSpecificCredentials), varargs...)
}

func (m *MockIamClient) ListSigningCertificates(arg0 context.Context, arg1 *iam.ListSigningCertificatesInput, arg2 ...func(*iam.Options)) (*iam.ListSigningCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSigningCertificates, varargs...)
	ret0, _ := ret[0].(*iam.ListSigningCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListSigningCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSigningCertificates, reflect.TypeOf((*MockIamClient)(nil).ListSigningCertificates), varargs...)
}

func (m *MockIamClient) ListUserPolicies(arg0 context.Context, arg1 *iam.ListUserPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserPolicies, varargs...)
	ret0, _ := ret[0].(*iam.ListUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListUserPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserPolicies, reflect.TypeOf((*MockIamClient)(nil).ListUserPolicies), varargs...)
}

func (m *MockIamClient) ListUserTags(arg0 context.Context, arg1 *iam.ListUserTagsInput, arg2 ...func(*iam.Options)) (*iam.ListUserTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserTags, varargs...)
	ret0, _ := ret[0].(*iam.ListUserTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListUserTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserTags, reflect.TypeOf((*MockIamClient)(nil).ListUserTags), varargs...)
}

func (m *MockIamClient) ListUsers(arg0 context.Context, arg1 *iam.ListUsersInput, arg2 ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsers, varargs...)
	ret0, _ := ret[0].(*iam.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsers, reflect.TypeOf((*MockIamClient)(nil).ListUsers), varargs...)
}

func (m *MockIamClient) ListVirtualMFADevices(arg0 context.Context, arg1 *iam.ListVirtualMFADevicesInput, arg2 ...func(*iam.Options)) (*iam.ListVirtualMFADevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVirtualMFADevices, varargs...)
	ret0, _ := ret[0].(*iam.ListVirtualMFADevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListVirtualMFADevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVirtualMFADevices, reflect.TypeOf((*MockIamClient)(nil).ListVirtualMFADevices), varargs...)
}
