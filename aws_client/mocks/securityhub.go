package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	securityhub "github.com/aws/aws-sdk-go-v2/service/securityhub"
	gomock "github.com/golang/mock/gomock"
)

type MockSecurityhubClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSecurityhubClientMockRecorder
}

type MockSecurityhubClientMockRecorder struct {
	mock *MockSecurityhubClient
}

func NewMockSecurityhubClient(ctrl *gomock.Controller) *MockSecurityhubClient {
	mock := &MockSecurityhubClient{ctrl: ctrl}
	mock.recorder = &MockSecurityhubClientMockRecorder{mock}
	return mock
}

func (m *MockSecurityhubClient) EXPECT() *MockSecurityhubClientMockRecorder {
	return m.recorder
}

func (m *MockSecurityhubClient) DescribeActionTargets(arg0 context.Context, arg1 *securityhub.DescribeActionTargetsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeActionTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeActionTargets, varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeActionTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) DescribeActionTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeActionTargets, reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeActionTargets), varargs...)
}

func (m *MockSecurityhubClient) DescribeHub(arg0 context.Context, arg1 *securityhub.DescribeHubInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeHubOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHub, varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeHubOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) DescribeHub(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHub, reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeHub), varargs...)
}

func (m *MockSecurityhubClient) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *securityhub.DescribeOrganizationConfigurationInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeOrganizationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConfiguration, varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConfiguration, reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeOrganizationConfiguration), varargs...)
}

func (m *MockSecurityhubClient) DescribeProducts(arg0 context.Context, arg1 *securityhub.DescribeProductsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeProductsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProducts, varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeProductsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) DescribeProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProducts, reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeProducts), varargs...)
}

func (m *MockSecurityhubClient) DescribeStandards(arg0 context.Context, arg1 *securityhub.DescribeStandardsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeStandardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStandards, varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeStandardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) DescribeStandards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStandards, reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeStandards), varargs...)
}

func (m *MockSecurityhubClient) DescribeStandardsControls(arg0 context.Context, arg1 *securityhub.DescribeStandardsControlsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeStandardsControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStandardsControls, varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeStandardsControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) DescribeStandardsControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStandardsControls, reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeStandardsControls), varargs...)
}

func (m *MockSecurityhubClient) GetAdministratorAccount(arg0 context.Context, arg1 *securityhub.GetAdministratorAccountInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetAdministratorAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAdministratorAccount, varargs...)
	ret0, _ := ret[0].(*securityhub.GetAdministratorAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetAdministratorAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAdministratorAccount, reflect.TypeOf((*MockSecurityhubClient)(nil).GetAdministratorAccount), varargs...)
}

func (m *MockSecurityhubClient) GetEnabledStandards(arg0 context.Context, arg1 *securityhub.GetEnabledStandardsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetEnabledStandardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEnabledStandards, varargs...)
	ret0, _ := ret[0].(*securityhub.GetEnabledStandardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetEnabledStandards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEnabledStandards, reflect.TypeOf((*MockSecurityhubClient)(nil).GetEnabledStandards), varargs...)
}

func (m *MockSecurityhubClient) GetFindingAggregator(arg0 context.Context, arg1 *securityhub.GetFindingAggregatorInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetFindingAggregatorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFindingAggregator, varargs...)
	ret0, _ := ret[0].(*securityhub.GetFindingAggregatorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetFindingAggregator(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFindingAggregator, reflect.TypeOf((*MockSecurityhubClient)(nil).GetFindingAggregator), varargs...)
}

func (m *MockSecurityhubClient) GetFindings(arg0 context.Context, arg1 *securityhub.GetFindingsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFindings, varargs...)
	ret0, _ := ret[0].(*securityhub.GetFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFindings, reflect.TypeOf((*MockSecurityhubClient)(nil).GetFindings), varargs...)
}

func (m *MockSecurityhubClient) GetInsightResults(arg0 context.Context, arg1 *securityhub.GetInsightResultsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetInsightResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsightResults, varargs...)
	ret0, _ := ret[0].(*securityhub.GetInsightResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetInsightResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsightResults, reflect.TypeOf((*MockSecurityhubClient)(nil).GetInsightResults), varargs...)
}

func (m *MockSecurityhubClient) GetInsights(arg0 context.Context, arg1 *securityhub.GetInsightsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetInsightsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsights, varargs...)
	ret0, _ := ret[0].(*securityhub.GetInsightsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetInsights(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsights, reflect.TypeOf((*MockSecurityhubClient)(nil).GetInsights), varargs...)
}

func (m *MockSecurityhubClient) GetInvitationsCount(arg0 context.Context, arg1 *securityhub.GetInvitationsCountInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetInvitationsCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInvitationsCount, varargs...)
	ret0, _ := ret[0].(*securityhub.GetInvitationsCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetInvitationsCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInvitationsCount, reflect.TypeOf((*MockSecurityhubClient)(nil).GetInvitationsCount), varargs...)
}

func (m *MockSecurityhubClient) GetMasterAccount(arg0 context.Context, arg1 *securityhub.GetMasterAccountInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetMasterAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMasterAccount, varargs...)
	ret0, _ := ret[0].(*securityhub.GetMasterAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetMasterAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMasterAccount, reflect.TypeOf((*MockSecurityhubClient)(nil).GetMasterAccount), varargs...)
}

func (m *MockSecurityhubClient) GetMembers(arg0 context.Context, arg1 *securityhub.GetMembersInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMembers, varargs...)
	ret0, _ := ret[0].(*securityhub.GetMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) GetMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMembers, reflect.TypeOf((*MockSecurityhubClient)(nil).GetMembers), varargs...)
}

func (m *MockSecurityhubClient) ListEnabledProductsForImport(arg0 context.Context, arg1 *securityhub.ListEnabledProductsForImportInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListEnabledProductsForImportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEnabledProductsForImport, varargs...)
	ret0, _ := ret[0].(*securityhub.ListEnabledProductsForImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) ListEnabledProductsForImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEnabledProductsForImport, reflect.TypeOf((*MockSecurityhubClient)(nil).ListEnabledProductsForImport), varargs...)
}

func (m *MockSecurityhubClient) ListFindingAggregators(arg0 context.Context, arg1 *securityhub.ListFindingAggregatorsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListFindingAggregatorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFindingAggregators, varargs...)
	ret0, _ := ret[0].(*securityhub.ListFindingAggregatorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) ListFindingAggregators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFindingAggregators, reflect.TypeOf((*MockSecurityhubClient)(nil).ListFindingAggregators), varargs...)
}

func (m *MockSecurityhubClient) ListInvitations(arg0 context.Context, arg1 *securityhub.ListInvitationsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListInvitationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInvitations, varargs...)
	ret0, _ := ret[0].(*securityhub.ListInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) ListInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInvitations, reflect.TypeOf((*MockSecurityhubClient)(nil).ListInvitations), varargs...)
}

func (m *MockSecurityhubClient) ListMembers(arg0 context.Context, arg1 *securityhub.ListMembersInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMembers, varargs...)
	ret0, _ := ret[0].(*securityhub.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMembers, reflect.TypeOf((*MockSecurityhubClient)(nil).ListMembers), varargs...)
}

func (m *MockSecurityhubClient) ListOrganizationAdminAccounts(arg0 context.Context, arg1 *securityhub.ListOrganizationAdminAccountsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListOrganizationAdminAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOrganizationAdminAccounts, varargs...)
	ret0, _ := ret[0].(*securityhub.ListOrganizationAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) ListOrganizationAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOrganizationAdminAccounts, reflect.TypeOf((*MockSecurityhubClient)(nil).ListOrganizationAdminAccounts), varargs...)
}

func (m *MockSecurityhubClient) ListTagsForResource(arg0 context.Context, arg1 *securityhub.ListTagsForResourceInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*securityhub.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityhubClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSecurityhubClient)(nil).ListTagsForResource), varargs...)
}
