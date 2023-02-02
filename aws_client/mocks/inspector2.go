package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	inspector2 "github.com/aws/aws-sdk-go-v2/service/inspector2"
	gomock "github.com/golang/mock/gomock"
)

type MockInspector2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockInspector2ClientMockRecorder
}

type MockInspector2ClientMockRecorder struct {
	mock *MockInspector2Client
}

func NewMockInspector2Client(ctrl *gomock.Controller) *MockInspector2Client {
	mock := &MockInspector2Client{ctrl: ctrl}
	mock.recorder = &MockInspector2ClientMockRecorder{mock}
	return mock
}

func (m *MockInspector2Client) EXPECT() *MockInspector2ClientMockRecorder {
	return m.recorder
}

func (m *MockInspector2Client) BatchGetAccountStatus(arg0 context.Context, arg1 *inspector2.BatchGetAccountStatusInput, arg2 ...func(*inspector2.Options)) (*inspector2.BatchGetAccountStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetAccountStatus, varargs...)
	ret0, _ := ret[0].(*inspector2.BatchGetAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) BatchGetAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetAccountStatus, reflect.TypeOf((*MockInspector2Client)(nil).BatchGetAccountStatus), varargs...)
}

func (m *MockInspector2Client) BatchGetFreeTrialInfo(arg0 context.Context, arg1 *inspector2.BatchGetFreeTrialInfoInput, arg2 ...func(*inspector2.Options)) (*inspector2.BatchGetFreeTrialInfoOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetFreeTrialInfo, varargs...)
	ret0, _ := ret[0].(*inspector2.BatchGetFreeTrialInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) BatchGetFreeTrialInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetFreeTrialInfo, reflect.TypeOf((*MockInspector2Client)(nil).BatchGetFreeTrialInfo), varargs...)
}

func (m *MockInspector2Client) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *inspector2.DescribeOrganizationConfigurationInput, arg2 ...func(*inspector2.Options)) (*inspector2.DescribeOrganizationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConfiguration, varargs...)
	ret0, _ := ret[0].(*inspector2.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConfiguration, reflect.TypeOf((*MockInspector2Client)(nil).DescribeOrganizationConfiguration), varargs...)
}

func (m *MockInspector2Client) GetConfiguration(arg0 context.Context, arg1 *inspector2.GetConfigurationInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConfiguration, varargs...)
	ret0, _ := ret[0].(*inspector2.GetConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) GetConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConfiguration, reflect.TypeOf((*MockInspector2Client)(nil).GetConfiguration), varargs...)
}

func (m *MockInspector2Client) GetDelegatedAdminAccount(arg0 context.Context, arg1 *inspector2.GetDelegatedAdminAccountInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetDelegatedAdminAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDelegatedAdminAccount, varargs...)
	ret0, _ := ret[0].(*inspector2.GetDelegatedAdminAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) GetDelegatedAdminAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDelegatedAdminAccount, reflect.TypeOf((*MockInspector2Client)(nil).GetDelegatedAdminAccount), varargs...)
}

func (m *MockInspector2Client) GetFindingsReportStatus(arg0 context.Context, arg1 *inspector2.GetFindingsReportStatusInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetFindingsReportStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFindingsReportStatus, varargs...)
	ret0, _ := ret[0].(*inspector2.GetFindingsReportStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) GetFindingsReportStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFindingsReportStatus, reflect.TypeOf((*MockInspector2Client)(nil).GetFindingsReportStatus), varargs...)
}

func (m *MockInspector2Client) GetMember(arg0 context.Context, arg1 *inspector2.GetMemberInput, arg2 ...func(*inspector2.Options)) (*inspector2.GetMemberOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMember, varargs...)
	ret0, _ := ret[0].(*inspector2.GetMemberOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) GetMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMember, reflect.TypeOf((*MockInspector2Client)(nil).GetMember), varargs...)
}

func (m *MockInspector2Client) ListAccountPermissions(arg0 context.Context, arg1 *inspector2.ListAccountPermissionsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListAccountPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountPermissions, varargs...)
	ret0, _ := ret[0].(*inspector2.ListAccountPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListAccountPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountPermissions, reflect.TypeOf((*MockInspector2Client)(nil).ListAccountPermissions), varargs...)
}

func (m *MockInspector2Client) ListCoverage(arg0 context.Context, arg1 *inspector2.ListCoverageInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListCoverageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCoverage, varargs...)
	ret0, _ := ret[0].(*inspector2.ListCoverageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListCoverage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCoverage, reflect.TypeOf((*MockInspector2Client)(nil).ListCoverage), varargs...)
}

func (m *MockInspector2Client) ListCoverageStatistics(arg0 context.Context, arg1 *inspector2.ListCoverageStatisticsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListCoverageStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCoverageStatistics, varargs...)
	ret0, _ := ret[0].(*inspector2.ListCoverageStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListCoverageStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCoverageStatistics, reflect.TypeOf((*MockInspector2Client)(nil).ListCoverageStatistics), varargs...)
}

func (m *MockInspector2Client) ListDelegatedAdminAccounts(arg0 context.Context, arg1 *inspector2.ListDelegatedAdminAccountsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListDelegatedAdminAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDelegatedAdminAccounts, varargs...)
	ret0, _ := ret[0].(*inspector2.ListDelegatedAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListDelegatedAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDelegatedAdminAccounts, reflect.TypeOf((*MockInspector2Client)(nil).ListDelegatedAdminAccounts), varargs...)
}

func (m *MockInspector2Client) ListFilters(arg0 context.Context, arg1 *inspector2.ListFiltersInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFilters, varargs...)
	ret0, _ := ret[0].(*inspector2.ListFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFilters, reflect.TypeOf((*MockInspector2Client)(nil).ListFilters), varargs...)
}

func (m *MockInspector2Client) ListFindingAggregations(arg0 context.Context, arg1 *inspector2.ListFindingAggregationsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFindingAggregationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFindingAggregations, varargs...)
	ret0, _ := ret[0].(*inspector2.ListFindingAggregationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListFindingAggregations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFindingAggregations, reflect.TypeOf((*MockInspector2Client)(nil).ListFindingAggregations), varargs...)
}

func (m *MockInspector2Client) ListFindings(arg0 context.Context, arg1 *inspector2.ListFindingsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFindings, varargs...)
	ret0, _ := ret[0].(*inspector2.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFindings, reflect.TypeOf((*MockInspector2Client)(nil).ListFindings), varargs...)
}

func (m *MockInspector2Client) ListMembers(arg0 context.Context, arg1 *inspector2.ListMembersInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMembers, varargs...)
	ret0, _ := ret[0].(*inspector2.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMembers, reflect.TypeOf((*MockInspector2Client)(nil).ListMembers), varargs...)
}

func (m *MockInspector2Client) ListTagsForResource(arg0 context.Context, arg1 *inspector2.ListTagsForResourceInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*inspector2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockInspector2Client)(nil).ListTagsForResource), varargs...)
}

func (m *MockInspector2Client) ListUsageTotals(arg0 context.Context, arg1 *inspector2.ListUsageTotalsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListUsageTotalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsageTotals, varargs...)
	ret0, _ := ret[0].(*inspector2.ListUsageTotalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspector2ClientMockRecorder) ListUsageTotals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsageTotals, reflect.TypeOf((*MockInspector2Client)(nil).ListUsageTotals), varargs...)
}
