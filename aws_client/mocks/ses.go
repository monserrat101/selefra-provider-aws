package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ses "github.com/aws/aws-sdk-go-v2/service/ses"
	gomock "github.com/golang/mock/gomock"
)

type MockSesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSesClientMockRecorder
}

type MockSesClientMockRecorder struct {
	mock *MockSesClient
}

func NewMockSesClient(ctrl *gomock.Controller) *MockSesClient {
	mock := &MockSesClient{ctrl: ctrl}
	mock.recorder = &MockSesClientMockRecorder{mock}
	return mock
}

func (m *MockSesClient) EXPECT() *MockSesClientMockRecorder {
	return m.recorder
}

func (m *MockSesClient) DescribeActiveReceiptRuleSet(arg0 context.Context, arg1 *ses.DescribeActiveReceiptRuleSetInput, arg2 ...func(*ses.Options)) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeActiveReceiptRuleSet, varargs...)
	ret0, _ := ret[0].(*ses.DescribeActiveReceiptRuleSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) DescribeActiveReceiptRuleSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeActiveReceiptRuleSet, reflect.TypeOf((*MockSesClient)(nil).DescribeActiveReceiptRuleSet), varargs...)
}

func (m *MockSesClient) DescribeConfigurationSet(arg0 context.Context, arg1 *ses.DescribeConfigurationSetInput, arg2 ...func(*ses.Options)) (*ses.DescribeConfigurationSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationSet, varargs...)
	ret0, _ := ret[0].(*ses.DescribeConfigurationSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) DescribeConfigurationSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationSet, reflect.TypeOf((*MockSesClient)(nil).DescribeConfigurationSet), varargs...)
}

func (m *MockSesClient) DescribeReceiptRule(arg0 context.Context, arg1 *ses.DescribeReceiptRuleInput, arg2 ...func(*ses.Options)) (*ses.DescribeReceiptRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReceiptRule, varargs...)
	ret0, _ := ret[0].(*ses.DescribeReceiptRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) DescribeReceiptRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReceiptRule, reflect.TypeOf((*MockSesClient)(nil).DescribeReceiptRule), varargs...)
}

func (m *MockSesClient) DescribeReceiptRuleSet(arg0 context.Context, arg1 *ses.DescribeReceiptRuleSetInput, arg2 ...func(*ses.Options)) (*ses.DescribeReceiptRuleSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReceiptRuleSet, varargs...)
	ret0, _ := ret[0].(*ses.DescribeReceiptRuleSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) DescribeReceiptRuleSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReceiptRuleSet, reflect.TypeOf((*MockSesClient)(nil).DescribeReceiptRuleSet), varargs...)
}

func (m *MockSesClient) GetAccountSendingEnabled(arg0 context.Context, arg1 *ses.GetAccountSendingEnabledInput, arg2 ...func(*ses.Options)) (*ses.GetAccountSendingEnabledOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountSendingEnabled, varargs...)
	ret0, _ := ret[0].(*ses.GetAccountSendingEnabledOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetAccountSendingEnabled(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountSendingEnabled, reflect.TypeOf((*MockSesClient)(nil).GetAccountSendingEnabled), varargs...)
}

func (m *MockSesClient) GetCustomVerificationEmailTemplate(arg0 context.Context, arg1 *ses.GetCustomVerificationEmailTemplateInput, arg2 ...func(*ses.Options)) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCustomVerificationEmailTemplate, varargs...)
	ret0, _ := ret[0].(*ses.GetCustomVerificationEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetCustomVerificationEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCustomVerificationEmailTemplate, reflect.TypeOf((*MockSesClient)(nil).GetCustomVerificationEmailTemplate), varargs...)
}

func (m *MockSesClient) GetIdentityDkimAttributes(arg0 context.Context, arg1 *ses.GetIdentityDkimAttributesInput, arg2 ...func(*ses.Options)) (*ses.GetIdentityDkimAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityDkimAttributes, varargs...)
	ret0, _ := ret[0].(*ses.GetIdentityDkimAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetIdentityDkimAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityDkimAttributes, reflect.TypeOf((*MockSesClient)(nil).GetIdentityDkimAttributes), varargs...)
}

func (m *MockSesClient) GetIdentityMailFromDomainAttributes(arg0 context.Context, arg1 *ses.GetIdentityMailFromDomainAttributesInput, arg2 ...func(*ses.Options)) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityMailFromDomainAttributes, varargs...)
	ret0, _ := ret[0].(*ses.GetIdentityMailFromDomainAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetIdentityMailFromDomainAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityMailFromDomainAttributes, reflect.TypeOf((*MockSesClient)(nil).GetIdentityMailFromDomainAttributes), varargs...)
}

func (m *MockSesClient) GetIdentityNotificationAttributes(arg0 context.Context, arg1 *ses.GetIdentityNotificationAttributesInput, arg2 ...func(*ses.Options)) (*ses.GetIdentityNotificationAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityNotificationAttributes, varargs...)
	ret0, _ := ret[0].(*ses.GetIdentityNotificationAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetIdentityNotificationAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityNotificationAttributes, reflect.TypeOf((*MockSesClient)(nil).GetIdentityNotificationAttributes), varargs...)
}

func (m *MockSesClient) GetIdentityPolicies(arg0 context.Context, arg1 *ses.GetIdentityPoliciesInput, arg2 ...func(*ses.Options)) (*ses.GetIdentityPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityPolicies, varargs...)
	ret0, _ := ret[0].(*ses.GetIdentityPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetIdentityPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityPolicies, reflect.TypeOf((*MockSesClient)(nil).GetIdentityPolicies), varargs...)
}

func (m *MockSesClient) GetIdentityVerificationAttributes(arg0 context.Context, arg1 *ses.GetIdentityVerificationAttributesInput, arg2 ...func(*ses.Options)) (*ses.GetIdentityVerificationAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityVerificationAttributes, varargs...)
	ret0, _ := ret[0].(*ses.GetIdentityVerificationAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetIdentityVerificationAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityVerificationAttributes, reflect.TypeOf((*MockSesClient)(nil).GetIdentityVerificationAttributes), varargs...)
}

func (m *MockSesClient) GetSendQuota(arg0 context.Context, arg1 *ses.GetSendQuotaInput, arg2 ...func(*ses.Options)) (*ses.GetSendQuotaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSendQuota, varargs...)
	ret0, _ := ret[0].(*ses.GetSendQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetSendQuota(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSendQuota, reflect.TypeOf((*MockSesClient)(nil).GetSendQuota), varargs...)
}

func (m *MockSesClient) GetSendStatistics(arg0 context.Context, arg1 *ses.GetSendStatisticsInput, arg2 ...func(*ses.Options)) (*ses.GetSendStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSendStatistics, varargs...)
	ret0, _ := ret[0].(*ses.GetSendStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetSendStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSendStatistics, reflect.TypeOf((*MockSesClient)(nil).GetSendStatistics), varargs...)
}

func (m *MockSesClient) GetTemplate(arg0 context.Context, arg1 *ses.GetTemplateInput, arg2 ...func(*ses.Options)) (*ses.GetTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTemplate, varargs...)
	ret0, _ := ret[0].(*ses.GetTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) GetTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTemplate, reflect.TypeOf((*MockSesClient)(nil).GetTemplate), varargs...)
}

func (m *MockSesClient) ListConfigurationSets(arg0 context.Context, arg1 *ses.ListConfigurationSetsInput, arg2 ...func(*ses.Options)) (*ses.ListConfigurationSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConfigurationSets, varargs...)
	ret0, _ := ret[0].(*ses.ListConfigurationSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListConfigurationSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConfigurationSets, reflect.TypeOf((*MockSesClient)(nil).ListConfigurationSets), varargs...)
}

func (m *MockSesClient) ListCustomVerificationEmailTemplates(arg0 context.Context, arg1 *ses.ListCustomVerificationEmailTemplatesInput, arg2 ...func(*ses.Options)) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCustomVerificationEmailTemplates, varargs...)
	ret0, _ := ret[0].(*ses.ListCustomVerificationEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListCustomVerificationEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCustomVerificationEmailTemplates, reflect.TypeOf((*MockSesClient)(nil).ListCustomVerificationEmailTemplates), varargs...)
}

func (m *MockSesClient) ListIdentities(arg0 context.Context, arg1 *ses.ListIdentitiesInput, arg2 ...func(*ses.Options)) (*ses.ListIdentitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIdentities, varargs...)
	ret0, _ := ret[0].(*ses.ListIdentitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListIdentities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIdentities, reflect.TypeOf((*MockSesClient)(nil).ListIdentities), varargs...)
}

func (m *MockSesClient) ListIdentityPolicies(arg0 context.Context, arg1 *ses.ListIdentityPoliciesInput, arg2 ...func(*ses.Options)) (*ses.ListIdentityPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIdentityPolicies, varargs...)
	ret0, _ := ret[0].(*ses.ListIdentityPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListIdentityPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIdentityPolicies, reflect.TypeOf((*MockSesClient)(nil).ListIdentityPolicies), varargs...)
}

func (m *MockSesClient) ListReceiptFilters(arg0 context.Context, arg1 *ses.ListReceiptFiltersInput, arg2 ...func(*ses.Options)) (*ses.ListReceiptFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReceiptFilters, varargs...)
	ret0, _ := ret[0].(*ses.ListReceiptFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListReceiptFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReceiptFilters, reflect.TypeOf((*MockSesClient)(nil).ListReceiptFilters), varargs...)
}

func (m *MockSesClient) ListReceiptRuleSets(arg0 context.Context, arg1 *ses.ListReceiptRuleSetsInput, arg2 ...func(*ses.Options)) (*ses.ListReceiptRuleSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReceiptRuleSets, varargs...)
	ret0, _ := ret[0].(*ses.ListReceiptRuleSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListReceiptRuleSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReceiptRuleSets, reflect.TypeOf((*MockSesClient)(nil).ListReceiptRuleSets), varargs...)
}

func (m *MockSesClient) ListTemplates(arg0 context.Context, arg1 *ses.ListTemplatesInput, arg2 ...func(*ses.Options)) (*ses.ListTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTemplates, varargs...)
	ret0, _ := ret[0].(*ses.ListTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTemplates, reflect.TypeOf((*MockSesClient)(nil).ListTemplates), varargs...)
}

func (m *MockSesClient) ListVerifiedEmailAddresses(arg0 context.Context, arg1 *ses.ListVerifiedEmailAddressesInput, arg2 ...func(*ses.Options)) (*ses.ListVerifiedEmailAddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVerifiedEmailAddresses, varargs...)
	ret0, _ := ret[0].(*ses.ListVerifiedEmailAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSesClientMockRecorder) ListVerifiedEmailAddresses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVerifiedEmailAddresses, reflect.TypeOf((*MockSesClient)(nil).ListVerifiedEmailAddresses), varargs...)
}
