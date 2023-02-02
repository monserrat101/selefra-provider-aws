package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	guardduty "github.com/aws/aws-sdk-go-v2/service/guardduty"
	gomock "github.com/golang/mock/gomock"
)

type MockGuarddutyClient struct {
	ctrl		*gomock.Controller
	recorder	*MockGuarddutyClientMockRecorder
}

type MockGuarddutyClientMockRecorder struct {
	mock *MockGuarddutyClient
}

func NewMockGuarddutyClient(ctrl *gomock.Controller) *MockGuarddutyClient {
	mock := &MockGuarddutyClient{ctrl: ctrl}
	mock.recorder = &MockGuarddutyClientMockRecorder{mock}
	return mock
}

func (m *MockGuarddutyClient) EXPECT() *MockGuarddutyClientMockRecorder {
	return m.recorder
}

func (m *MockGuarddutyClient) DescribeMalwareScans(arg0 context.Context, arg1 *guardduty.DescribeMalwareScansInput, arg2 ...func(*guardduty.Options)) (*guardduty.DescribeMalwareScansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMalwareScans, varargs...)
	ret0, _ := ret[0].(*guardduty.DescribeMalwareScansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) DescribeMalwareScans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMalwareScans, reflect.TypeOf((*MockGuarddutyClient)(nil).DescribeMalwareScans), varargs...)
}

func (m *MockGuarddutyClient) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *guardduty.DescribeOrganizationConfigurationInput, arg2 ...func(*guardduty.Options)) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConfiguration, varargs...)
	ret0, _ := ret[0].(*guardduty.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConfiguration, reflect.TypeOf((*MockGuarddutyClient)(nil).DescribeOrganizationConfiguration), varargs...)
}

func (m *MockGuarddutyClient) DescribePublishingDestination(arg0 context.Context, arg1 *guardduty.DescribePublishingDestinationInput, arg2 ...func(*guardduty.Options)) (*guardduty.DescribePublishingDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePublishingDestination, varargs...)
	ret0, _ := ret[0].(*guardduty.DescribePublishingDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) DescribePublishingDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePublishingDestination, reflect.TypeOf((*MockGuarddutyClient)(nil).DescribePublishingDestination), varargs...)
}

func (m *MockGuarddutyClient) GetAdministratorAccount(arg0 context.Context, arg1 *guardduty.GetAdministratorAccountInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetAdministratorAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAdministratorAccount, varargs...)
	ret0, _ := ret[0].(*guardduty.GetAdministratorAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetAdministratorAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAdministratorAccount, reflect.TypeOf((*MockGuarddutyClient)(nil).GetAdministratorAccount), varargs...)
}

func (m *MockGuarddutyClient) GetDetector(arg0 context.Context, arg1 *guardduty.GetDetectorInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDetector, varargs...)
	ret0, _ := ret[0].(*guardduty.GetDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetDetector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDetector, reflect.TypeOf((*MockGuarddutyClient)(nil).GetDetector), varargs...)
}

func (m *MockGuarddutyClient) GetFilter(arg0 context.Context, arg1 *guardduty.GetFilterInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetFilterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFilter, varargs...)
	ret0, _ := ret[0].(*guardduty.GetFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetFilter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFilter, reflect.TypeOf((*MockGuarddutyClient)(nil).GetFilter), varargs...)
}

func (m *MockGuarddutyClient) GetFindings(arg0 context.Context, arg1 *guardduty.GetFindingsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFindings, varargs...)
	ret0, _ := ret[0].(*guardduty.GetFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFindings, reflect.TypeOf((*MockGuarddutyClient)(nil).GetFindings), varargs...)
}

func (m *MockGuarddutyClient) GetFindingsStatistics(arg0 context.Context, arg1 *guardduty.GetFindingsStatisticsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetFindingsStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFindingsStatistics, varargs...)
	ret0, _ := ret[0].(*guardduty.GetFindingsStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetFindingsStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFindingsStatistics, reflect.TypeOf((*MockGuarddutyClient)(nil).GetFindingsStatistics), varargs...)
}

func (m *MockGuarddutyClient) GetIPSet(arg0 context.Context, arg1 *guardduty.GetIPSetInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIPSet, varargs...)
	ret0, _ := ret[0].(*guardduty.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIPSet, reflect.TypeOf((*MockGuarddutyClient)(nil).GetIPSet), varargs...)
}

func (m *MockGuarddutyClient) GetInvitationsCount(arg0 context.Context, arg1 *guardduty.GetInvitationsCountInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetInvitationsCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInvitationsCount, varargs...)
	ret0, _ := ret[0].(*guardduty.GetInvitationsCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetInvitationsCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInvitationsCount, reflect.TypeOf((*MockGuarddutyClient)(nil).GetInvitationsCount), varargs...)
}

func (m *MockGuarddutyClient) GetMalwareScanSettings(arg0 context.Context, arg1 *guardduty.GetMalwareScanSettingsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMalwareScanSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMalwareScanSettings, varargs...)
	ret0, _ := ret[0].(*guardduty.GetMalwareScanSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetMalwareScanSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMalwareScanSettings, reflect.TypeOf((*MockGuarddutyClient)(nil).GetMalwareScanSettings), varargs...)
}

func (m *MockGuarddutyClient) GetMasterAccount(arg0 context.Context, arg1 *guardduty.GetMasterAccountInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMasterAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMasterAccount, varargs...)
	ret0, _ := ret[0].(*guardduty.GetMasterAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetMasterAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMasterAccount, reflect.TypeOf((*MockGuarddutyClient)(nil).GetMasterAccount), varargs...)
}

func (m *MockGuarddutyClient) GetMemberDetectors(arg0 context.Context, arg1 *guardduty.GetMemberDetectorsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMemberDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMemberDetectors, varargs...)
	ret0, _ := ret[0].(*guardduty.GetMemberDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetMemberDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMemberDetectors, reflect.TypeOf((*MockGuarddutyClient)(nil).GetMemberDetectors), varargs...)
}

func (m *MockGuarddutyClient) GetMembers(arg0 context.Context, arg1 *guardduty.GetMembersInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMembers, varargs...)
	ret0, _ := ret[0].(*guardduty.GetMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMembers, reflect.TypeOf((*MockGuarddutyClient)(nil).GetMembers), varargs...)
}

func (m *MockGuarddutyClient) GetRemainingFreeTrialDays(arg0 context.Context, arg1 *guardduty.GetRemainingFreeTrialDaysInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetRemainingFreeTrialDaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRemainingFreeTrialDays, varargs...)
	ret0, _ := ret[0].(*guardduty.GetRemainingFreeTrialDaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetRemainingFreeTrialDays(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRemainingFreeTrialDays, reflect.TypeOf((*MockGuarddutyClient)(nil).GetRemainingFreeTrialDays), varargs...)
}

func (m *MockGuarddutyClient) GetThreatIntelSet(arg0 context.Context, arg1 *guardduty.GetThreatIntelSetInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetThreatIntelSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetThreatIntelSet, varargs...)
	ret0, _ := ret[0].(*guardduty.GetThreatIntelSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetThreatIntelSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetThreatIntelSet, reflect.TypeOf((*MockGuarddutyClient)(nil).GetThreatIntelSet), varargs...)
}

func (m *MockGuarddutyClient) GetUsageStatistics(arg0 context.Context, arg1 *guardduty.GetUsageStatisticsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetUsageStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUsageStatistics, varargs...)
	ret0, _ := ret[0].(*guardduty.GetUsageStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) GetUsageStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUsageStatistics, reflect.TypeOf((*MockGuarddutyClient)(nil).GetUsageStatistics), varargs...)
}

func (m *MockGuarddutyClient) ListDetectors(arg0 context.Context, arg1 *guardduty.ListDetectorsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDetectors, varargs...)
	ret0, _ := ret[0].(*guardduty.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDetectors, reflect.TypeOf((*MockGuarddutyClient)(nil).ListDetectors), varargs...)
}

func (m *MockGuarddutyClient) ListFilters(arg0 context.Context, arg1 *guardduty.ListFiltersInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFilters, varargs...)
	ret0, _ := ret[0].(*guardduty.ListFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFilters, reflect.TypeOf((*MockGuarddutyClient)(nil).ListFilters), varargs...)
}

func (m *MockGuarddutyClient) ListFindings(arg0 context.Context, arg1 *guardduty.ListFindingsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFindings, varargs...)
	ret0, _ := ret[0].(*guardduty.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFindings, reflect.TypeOf((*MockGuarddutyClient)(nil).ListFindings), varargs...)
}

func (m *MockGuarddutyClient) ListIPSets(arg0 context.Context, arg1 *guardduty.ListIPSetsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIPSets, varargs...)
	ret0, _ := ret[0].(*guardduty.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIPSets, reflect.TypeOf((*MockGuarddutyClient)(nil).ListIPSets), varargs...)
}

func (m *MockGuarddutyClient) ListInvitations(arg0 context.Context, arg1 *guardduty.ListInvitationsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListInvitationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInvitations, varargs...)
	ret0, _ := ret[0].(*guardduty.ListInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInvitations, reflect.TypeOf((*MockGuarddutyClient)(nil).ListInvitations), varargs...)
}

func (m *MockGuarddutyClient) ListMembers(arg0 context.Context, arg1 *guardduty.ListMembersInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMembers, varargs...)
	ret0, _ := ret[0].(*guardduty.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMembers, reflect.TypeOf((*MockGuarddutyClient)(nil).ListMembers), varargs...)
}

func (m *MockGuarddutyClient) ListOrganizationAdminAccounts(arg0 context.Context, arg1 *guardduty.ListOrganizationAdminAccountsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOrganizationAdminAccounts, varargs...)
	ret0, _ := ret[0].(*guardduty.ListOrganizationAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListOrganizationAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOrganizationAdminAccounts, reflect.TypeOf((*MockGuarddutyClient)(nil).ListOrganizationAdminAccounts), varargs...)
}

func (m *MockGuarddutyClient) ListPublishingDestinations(arg0 context.Context, arg1 *guardduty.ListPublishingDestinationsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListPublishingDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPublishingDestinations, varargs...)
	ret0, _ := ret[0].(*guardduty.ListPublishingDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListPublishingDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPublishingDestinations, reflect.TypeOf((*MockGuarddutyClient)(nil).ListPublishingDestinations), varargs...)
}

func (m *MockGuarddutyClient) ListTagsForResource(arg0 context.Context, arg1 *guardduty.ListTagsForResourceInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*guardduty.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockGuarddutyClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockGuarddutyClient) ListThreatIntelSets(arg0 context.Context, arg1 *guardduty.ListThreatIntelSetsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListThreatIntelSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThreatIntelSets, varargs...)
	ret0, _ := ret[0].(*guardduty.ListThreatIntelSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuarddutyClientMockRecorder) ListThreatIntelSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThreatIntelSets, reflect.TypeOf((*MockGuarddutyClient)(nil).ListThreatIntelSets), varargs...)
}
