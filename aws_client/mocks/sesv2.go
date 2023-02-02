package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	sesv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	gomock "github.com/golang/mock/gomock"
)

type MockSesv2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockSesv2ClientMockRecorder
}

type MockSesv2ClientMockRecorder struct {
	mock *MockSesv2Client
}

func NewMockSesv2Client(ctrl *gomock.Controller) *MockSesv2Client {
	mock := &MockSesv2Client{ctrl: ctrl}
	mock.recorder = &MockSesv2ClientMockRecorder{mock}
	return mock
}

func (m *MockSesv2Client) EXPECT() *MockSesv2ClientMockRecorder {
	return m.recorder
}

func (m *MockSesv2Client) BatchGetMetricData(arg0 context.Context, arg1 *sesv2.BatchGetMetricDataInput, arg2 ...func(*sesv2.Options)) (*sesv2.BatchGetMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetMetricData, varargs...)
	ret0, _ := ret[0].(*sesv2.BatchGetMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) BatchGetMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetMetricData, reflect.TypeOf((*MockSesv2Client)(nil).BatchGetMetricData), varargs...)
}

func (m *MockSesv2Client) GetAccount(arg0 context.Context, arg1 *sesv2.GetAccountInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccount, varargs...)
	ret0, _ := ret[0].(*sesv2.GetAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccount, reflect.TypeOf((*MockSesv2Client)(nil).GetAccount), varargs...)
}

func (m *MockSesv2Client) GetBlacklistReports(arg0 context.Context, arg1 *sesv2.GetBlacklistReportsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetBlacklistReportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlacklistReports, varargs...)
	ret0, _ := ret[0].(*sesv2.GetBlacklistReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetBlacklistReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlacklistReports, reflect.TypeOf((*MockSesv2Client)(nil).GetBlacklistReports), varargs...)
}

func (m *MockSesv2Client) GetConfigurationSet(arg0 context.Context, arg1 *sesv2.GetConfigurationSetInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConfigurationSet, varargs...)
	ret0, _ := ret[0].(*sesv2.GetConfigurationSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetConfigurationSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConfigurationSet, reflect.TypeOf((*MockSesv2Client)(nil).GetConfigurationSet), varargs...)
}

func (m *MockSesv2Client) GetConfigurationSetEventDestinations(arg0 context.Context, arg1 *sesv2.GetConfigurationSetEventDestinationsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConfigurationSetEventDestinations, varargs...)
	ret0, _ := ret[0].(*sesv2.GetConfigurationSetEventDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetConfigurationSetEventDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConfigurationSetEventDestinations, reflect.TypeOf((*MockSesv2Client)(nil).GetConfigurationSetEventDestinations), varargs...)
}

func (m *MockSesv2Client) GetContact(arg0 context.Context, arg1 *sesv2.GetContactInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetContactOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContact, varargs...)
	ret0, _ := ret[0].(*sesv2.GetContactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetContact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContact, reflect.TypeOf((*MockSesv2Client)(nil).GetContact), varargs...)
}

func (m *MockSesv2Client) GetContactList(arg0 context.Context, arg1 *sesv2.GetContactListInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetContactListOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContactList, varargs...)
	ret0, _ := ret[0].(*sesv2.GetContactListOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetContactList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContactList, reflect.TypeOf((*MockSesv2Client)(nil).GetContactList), varargs...)
}

func (m *MockSesv2Client) GetCustomVerificationEmailTemplate(arg0 context.Context, arg1 *sesv2.GetCustomVerificationEmailTemplateInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCustomVerificationEmailTemplate, varargs...)
	ret0, _ := ret[0].(*sesv2.GetCustomVerificationEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetCustomVerificationEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCustomVerificationEmailTemplate, reflect.TypeOf((*MockSesv2Client)(nil).GetCustomVerificationEmailTemplate), varargs...)
}

func (m *MockSesv2Client) GetDedicatedIp(arg0 context.Context, arg1 *sesv2.GetDedicatedIpInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDedicatedIp, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDedicatedIpOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDedicatedIp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDedicatedIp, reflect.TypeOf((*MockSesv2Client)(nil).GetDedicatedIp), varargs...)
}

func (m *MockSesv2Client) GetDedicatedIpPool(arg0 context.Context, arg1 *sesv2.GetDedicatedIpPoolInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpPoolOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDedicatedIpPool, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDedicatedIpPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDedicatedIpPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDedicatedIpPool, reflect.TypeOf((*MockSesv2Client)(nil).GetDedicatedIpPool), varargs...)
}

func (m *MockSesv2Client) GetDedicatedIps(arg0 context.Context, arg1 *sesv2.GetDedicatedIpsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDedicatedIps, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDedicatedIpsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDedicatedIps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDedicatedIps, reflect.TypeOf((*MockSesv2Client)(nil).GetDedicatedIps), varargs...)
}

func (m *MockSesv2Client) GetDeliverabilityDashboardOptions(arg0 context.Context, arg1 *sesv2.GetDeliverabilityDashboardOptionsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeliverabilityDashboardOptions, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDeliverabilityDashboardOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDeliverabilityDashboardOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeliverabilityDashboardOptions, reflect.TypeOf((*MockSesv2Client)(nil).GetDeliverabilityDashboardOptions), varargs...)
}

func (m *MockSesv2Client) GetDeliverabilityTestReport(arg0 context.Context, arg1 *sesv2.GetDeliverabilityTestReportInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityTestReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeliverabilityTestReport, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDeliverabilityTestReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDeliverabilityTestReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeliverabilityTestReport, reflect.TypeOf((*MockSesv2Client)(nil).GetDeliverabilityTestReport), varargs...)
}

func (m *MockSesv2Client) GetDomainDeliverabilityCampaign(arg0 context.Context, arg1 *sesv2.GetDomainDeliverabilityCampaignInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainDeliverabilityCampaign, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDomainDeliverabilityCampaignOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDomainDeliverabilityCampaign(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainDeliverabilityCampaign, reflect.TypeOf((*MockSesv2Client)(nil).GetDomainDeliverabilityCampaign), varargs...)
}

func (m *MockSesv2Client) GetDomainStatisticsReport(arg0 context.Context, arg1 *sesv2.GetDomainStatisticsReportInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDomainStatisticsReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainStatisticsReport, varargs...)
	ret0, _ := ret[0].(*sesv2.GetDomainStatisticsReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetDomainStatisticsReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainStatisticsReport, reflect.TypeOf((*MockSesv2Client)(nil).GetDomainStatisticsReport), varargs...)
}

func (m *MockSesv2Client) GetEmailIdentity(arg0 context.Context, arg1 *sesv2.GetEmailIdentityInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEmailIdentity, varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetEmailIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEmailIdentity, reflect.TypeOf((*MockSesv2Client)(nil).GetEmailIdentity), varargs...)
}

func (m *MockSesv2Client) GetEmailIdentityPolicies(arg0 context.Context, arg1 *sesv2.GetEmailIdentityPoliciesInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEmailIdentityPolicies, varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailIdentityPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetEmailIdentityPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEmailIdentityPolicies, reflect.TypeOf((*MockSesv2Client)(nil).GetEmailIdentityPolicies), varargs...)
}

func (m *MockSesv2Client) GetEmailTemplate(arg0 context.Context, arg1 *sesv2.GetEmailTemplateInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEmailTemplate, varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEmailTemplate, reflect.TypeOf((*MockSesv2Client)(nil).GetEmailTemplate), varargs...)
}

func (m *MockSesv2Client) GetImportJob(arg0 context.Context, arg1 *sesv2.GetImportJobInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetImportJob, varargs...)
	ret0, _ := ret[0].(*sesv2.GetImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetImportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetImportJob, reflect.TypeOf((*MockSesv2Client)(nil).GetImportJob), varargs...)
}

func (m *MockSesv2Client) GetSuppressedDestination(arg0 context.Context, arg1 *sesv2.GetSuppressedDestinationInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetSuppressedDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSuppressedDestination, varargs...)
	ret0, _ := ret[0].(*sesv2.GetSuppressedDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) GetSuppressedDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSuppressedDestination, reflect.TypeOf((*MockSesv2Client)(nil).GetSuppressedDestination), varargs...)
}

func (m *MockSesv2Client) ListConfigurationSets(arg0 context.Context, arg1 *sesv2.ListConfigurationSetsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListConfigurationSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConfigurationSets, varargs...)
	ret0, _ := ret[0].(*sesv2.ListConfigurationSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListConfigurationSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConfigurationSets, reflect.TypeOf((*MockSesv2Client)(nil).ListConfigurationSets), varargs...)
}

func (m *MockSesv2Client) ListContactLists(arg0 context.Context, arg1 *sesv2.ListContactListsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListContactListsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListContactLists, varargs...)
	ret0, _ := ret[0].(*sesv2.ListContactListsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListContactLists(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListContactLists, reflect.TypeOf((*MockSesv2Client)(nil).ListContactLists), varargs...)
}

func (m *MockSesv2Client) ListContacts(arg0 context.Context, arg1 *sesv2.ListContactsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListContactsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListContacts, varargs...)
	ret0, _ := ret[0].(*sesv2.ListContactsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListContacts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListContacts, reflect.TypeOf((*MockSesv2Client)(nil).ListContacts), varargs...)
}

func (m *MockSesv2Client) ListCustomVerificationEmailTemplates(arg0 context.Context, arg1 *sesv2.ListCustomVerificationEmailTemplatesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCustomVerificationEmailTemplates, varargs...)
	ret0, _ := ret[0].(*sesv2.ListCustomVerificationEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListCustomVerificationEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCustomVerificationEmailTemplates, reflect.TypeOf((*MockSesv2Client)(nil).ListCustomVerificationEmailTemplates), varargs...)
}

func (m *MockSesv2Client) ListDedicatedIpPools(arg0 context.Context, arg1 *sesv2.ListDedicatedIpPoolsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListDedicatedIpPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDedicatedIpPools, varargs...)
	ret0, _ := ret[0].(*sesv2.ListDedicatedIpPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListDedicatedIpPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDedicatedIpPools, reflect.TypeOf((*MockSesv2Client)(nil).ListDedicatedIpPools), varargs...)
}

func (m *MockSesv2Client) ListDeliverabilityTestReports(arg0 context.Context, arg1 *sesv2.ListDeliverabilityTestReportsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDeliverabilityTestReports, varargs...)
	ret0, _ := ret[0].(*sesv2.ListDeliverabilityTestReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListDeliverabilityTestReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDeliverabilityTestReports, reflect.TypeOf((*MockSesv2Client)(nil).ListDeliverabilityTestReports), varargs...)
}

func (m *MockSesv2Client) ListDomainDeliverabilityCampaigns(arg0 context.Context, arg1 *sesv2.ListDomainDeliverabilityCampaignsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomainDeliverabilityCampaigns, varargs...)
	ret0, _ := ret[0].(*sesv2.ListDomainDeliverabilityCampaignsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListDomainDeliverabilityCampaigns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomainDeliverabilityCampaigns, reflect.TypeOf((*MockSesv2Client)(nil).ListDomainDeliverabilityCampaigns), varargs...)
}

func (m *MockSesv2Client) ListEmailIdentities(arg0 context.Context, arg1 *sesv2.ListEmailIdentitiesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListEmailIdentitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEmailIdentities, varargs...)
	ret0, _ := ret[0].(*sesv2.ListEmailIdentitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListEmailIdentities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEmailIdentities, reflect.TypeOf((*MockSesv2Client)(nil).ListEmailIdentities), varargs...)
}

func (m *MockSesv2Client) ListEmailTemplates(arg0 context.Context, arg1 *sesv2.ListEmailTemplatesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEmailTemplates, varargs...)
	ret0, _ := ret[0].(*sesv2.ListEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEmailTemplates, reflect.TypeOf((*MockSesv2Client)(nil).ListEmailTemplates), varargs...)
}

func (m *MockSesv2Client) ListImportJobs(arg0 context.Context, arg1 *sesv2.ListImportJobsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListImportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImportJobs, varargs...)
	ret0, _ := ret[0].(*sesv2.ListImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListImportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImportJobs, reflect.TypeOf((*MockSesv2Client)(nil).ListImportJobs), varargs...)
}

func (m *MockSesv2Client) ListRecommendations(arg0 context.Context, arg1 *sesv2.ListRecommendationsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRecommendations, varargs...)
	ret0, _ := ret[0].(*sesv2.ListRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRecommendations, reflect.TypeOf((*MockSesv2Client)(nil).ListRecommendations), varargs...)
}

func (m *MockSesv2Client) ListSuppressedDestinations(arg0 context.Context, arg1 *sesv2.ListSuppressedDestinationsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListSuppressedDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSuppressedDestinations, varargs...)
	ret0, _ := ret[0].(*sesv2.ListSuppressedDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListSuppressedDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSuppressedDestinations, reflect.TypeOf((*MockSesv2Client)(nil).ListSuppressedDestinations), varargs...)
}

func (m *MockSesv2Client) ListTagsForResource(arg0 context.Context, arg1 *sesv2.ListTagsForResourceInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*sesv2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesv2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSesv2Client)(nil).ListTagsForResource), varargs...)
}
