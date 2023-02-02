package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	waf "github.com/aws/aws-sdk-go-v2/service/waf"
	gomock "github.com/golang/mock/gomock"
)

type MockWafClient struct {
	ctrl		*gomock.Controller
	recorder	*MockWafClientMockRecorder
}

type MockWafClientMockRecorder struct {
	mock *MockWafClient
}

func NewMockWafClient(ctrl *gomock.Controller) *MockWafClient {
	mock := &MockWafClient{ctrl: ctrl}
	mock.recorder = &MockWafClientMockRecorder{mock}
	return mock
}

func (m *MockWafClient) EXPECT() *MockWafClientMockRecorder {
	return m.recorder
}

func (m *MockWafClient) GetByteMatchSet(arg0 context.Context, arg1 *waf.GetByteMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetByteMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetByteMatchSet, varargs...)
	ret0, _ := ret[0].(*waf.GetByteMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetByteMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetByteMatchSet, reflect.TypeOf((*MockWafClient)(nil).GetByteMatchSet), varargs...)
}

func (m *MockWafClient) GetChangeToken(arg0 context.Context, arg1 *waf.GetChangeTokenInput, arg2 ...func(*waf.Options)) (*waf.GetChangeTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetChangeToken, varargs...)
	ret0, _ := ret[0].(*waf.GetChangeTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetChangeToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetChangeToken, reflect.TypeOf((*MockWafClient)(nil).GetChangeToken), varargs...)
}

func (m *MockWafClient) GetChangeTokenStatus(arg0 context.Context, arg1 *waf.GetChangeTokenStatusInput, arg2 ...func(*waf.Options)) (*waf.GetChangeTokenStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetChangeTokenStatus, varargs...)
	ret0, _ := ret[0].(*waf.GetChangeTokenStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetChangeTokenStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetChangeTokenStatus, reflect.TypeOf((*MockWafClient)(nil).GetChangeTokenStatus), varargs...)
}

func (m *MockWafClient) GetGeoMatchSet(arg0 context.Context, arg1 *waf.GetGeoMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetGeoMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGeoMatchSet, varargs...)
	ret0, _ := ret[0].(*waf.GetGeoMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetGeoMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGeoMatchSet, reflect.TypeOf((*MockWafClient)(nil).GetGeoMatchSet), varargs...)
}

func (m *MockWafClient) GetIPSet(arg0 context.Context, arg1 *waf.GetIPSetInput, arg2 ...func(*waf.Options)) (*waf.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIPSet, varargs...)
	ret0, _ := ret[0].(*waf.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIPSet, reflect.TypeOf((*MockWafClient)(nil).GetIPSet), varargs...)
}

func (m *MockWafClient) GetLoggingConfiguration(arg0 context.Context, arg1 *waf.GetLoggingConfigurationInput, arg2 ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoggingConfiguration, varargs...)
	ret0, _ := ret[0].(*waf.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoggingConfiguration, reflect.TypeOf((*MockWafClient)(nil).GetLoggingConfiguration), varargs...)
}

func (m *MockWafClient) GetPermissionPolicy(arg0 context.Context, arg1 *waf.GetPermissionPolicyInput, arg2 ...func(*waf.Options)) (*waf.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPermissionPolicy, varargs...)
	ret0, _ := ret[0].(*waf.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPermissionPolicy, reflect.TypeOf((*MockWafClient)(nil).GetPermissionPolicy), varargs...)
}

func (m *MockWafClient) GetRateBasedRule(arg0 context.Context, arg1 *waf.GetRateBasedRuleInput, arg2 ...func(*waf.Options)) (*waf.GetRateBasedRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRateBasedRule, varargs...)
	ret0, _ := ret[0].(*waf.GetRateBasedRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetRateBasedRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRateBasedRule, reflect.TypeOf((*MockWafClient)(nil).GetRateBasedRule), varargs...)
}

func (m *MockWafClient) GetRateBasedRuleManagedKeys(arg0 context.Context, arg1 *waf.GetRateBasedRuleManagedKeysInput, arg2 ...func(*waf.Options)) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRateBasedRuleManagedKeys, varargs...)
	ret0, _ := ret[0].(*waf.GetRateBasedRuleManagedKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetRateBasedRuleManagedKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRateBasedRuleManagedKeys, reflect.TypeOf((*MockWafClient)(nil).GetRateBasedRuleManagedKeys), varargs...)
}

func (m *MockWafClient) GetRegexMatchSet(arg0 context.Context, arg1 *waf.GetRegexMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetRegexMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegexMatchSet, varargs...)
	ret0, _ := ret[0].(*waf.GetRegexMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetRegexMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegexMatchSet, reflect.TypeOf((*MockWafClient)(nil).GetRegexMatchSet), varargs...)
}

func (m *MockWafClient) GetRegexPatternSet(arg0 context.Context, arg1 *waf.GetRegexPatternSetInput, arg2 ...func(*waf.Options)) (*waf.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegexPatternSet, varargs...)
	ret0, _ := ret[0].(*waf.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegexPatternSet, reflect.TypeOf((*MockWafClient)(nil).GetRegexPatternSet), varargs...)
}

func (m *MockWafClient) GetRule(arg0 context.Context, arg1 *waf.GetRuleInput, arg2 ...func(*waf.Options)) (*waf.GetRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRule, varargs...)
	ret0, _ := ret[0].(*waf.GetRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRule, reflect.TypeOf((*MockWafClient)(nil).GetRule), varargs...)
}

func (m *MockWafClient) GetRuleGroup(arg0 context.Context, arg1 *waf.GetRuleGroupInput, arg2 ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRuleGroup, varargs...)
	ret0, _ := ret[0].(*waf.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRuleGroup, reflect.TypeOf((*MockWafClient)(nil).GetRuleGroup), varargs...)
}

func (m *MockWafClient) GetSampledRequests(arg0 context.Context, arg1 *waf.GetSampledRequestsInput, arg2 ...func(*waf.Options)) (*waf.GetSampledRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSampledRequests, varargs...)
	ret0, _ := ret[0].(*waf.GetSampledRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetSampledRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSampledRequests, reflect.TypeOf((*MockWafClient)(nil).GetSampledRequests), varargs...)
}

func (m *MockWafClient) GetSizeConstraintSet(arg0 context.Context, arg1 *waf.GetSizeConstraintSetInput, arg2 ...func(*waf.Options)) (*waf.GetSizeConstraintSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSizeConstraintSet, varargs...)
	ret0, _ := ret[0].(*waf.GetSizeConstraintSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetSizeConstraintSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSizeConstraintSet, reflect.TypeOf((*MockWafClient)(nil).GetSizeConstraintSet), varargs...)
}

func (m *MockWafClient) GetSqlInjectionMatchSet(arg0 context.Context, arg1 *waf.GetSqlInjectionMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetSqlInjectionMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSqlInjectionMatchSet, varargs...)
	ret0, _ := ret[0].(*waf.GetSqlInjectionMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetSqlInjectionMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSqlInjectionMatchSet, reflect.TypeOf((*MockWafClient)(nil).GetSqlInjectionMatchSet), varargs...)
}

func (m *MockWafClient) GetWebACL(arg0 context.Context, arg1 *waf.GetWebACLInput, arg2 ...func(*waf.Options)) (*waf.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWebACL, varargs...)
	ret0, _ := ret[0].(*waf.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWebACL, reflect.TypeOf((*MockWafClient)(nil).GetWebACL), varargs...)
}

func (m *MockWafClient) GetXssMatchSet(arg0 context.Context, arg1 *waf.GetXssMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetXssMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetXssMatchSet, varargs...)
	ret0, _ := ret[0].(*waf.GetXssMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) GetXssMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetXssMatchSet, reflect.TypeOf((*MockWafClient)(nil).GetXssMatchSet), varargs...)
}

func (m *MockWafClient) ListActivatedRulesInRuleGroup(arg0 context.Context, arg1 *waf.ListActivatedRulesInRuleGroupInput, arg2 ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActivatedRulesInRuleGroup, varargs...)
	ret0, _ := ret[0].(*waf.ListActivatedRulesInRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListActivatedRulesInRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActivatedRulesInRuleGroup, reflect.TypeOf((*MockWafClient)(nil).ListActivatedRulesInRuleGroup), varargs...)
}

func (m *MockWafClient) ListByteMatchSets(arg0 context.Context, arg1 *waf.ListByteMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListByteMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListByteMatchSets, varargs...)
	ret0, _ := ret[0].(*waf.ListByteMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListByteMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListByteMatchSets, reflect.TypeOf((*MockWafClient)(nil).ListByteMatchSets), varargs...)
}

func (m *MockWafClient) ListGeoMatchSets(arg0 context.Context, arg1 *waf.ListGeoMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListGeoMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGeoMatchSets, varargs...)
	ret0, _ := ret[0].(*waf.ListGeoMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListGeoMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGeoMatchSets, reflect.TypeOf((*MockWafClient)(nil).ListGeoMatchSets), varargs...)
}

func (m *MockWafClient) ListIPSets(arg0 context.Context, arg1 *waf.ListIPSetsInput, arg2 ...func(*waf.Options)) (*waf.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIPSets, varargs...)
	ret0, _ := ret[0].(*waf.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIPSets, reflect.TypeOf((*MockWafClient)(nil).ListIPSets), varargs...)
}

func (m *MockWafClient) ListLoggingConfigurations(arg0 context.Context, arg1 *waf.ListLoggingConfigurationsInput, arg2 ...func(*waf.Options)) (*waf.ListLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLoggingConfigurations, varargs...)
	ret0, _ := ret[0].(*waf.ListLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLoggingConfigurations, reflect.TypeOf((*MockWafClient)(nil).ListLoggingConfigurations), varargs...)
}

func (m *MockWafClient) ListRateBasedRules(arg0 context.Context, arg1 *waf.ListRateBasedRulesInput, arg2 ...func(*waf.Options)) (*waf.ListRateBasedRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRateBasedRules, varargs...)
	ret0, _ := ret[0].(*waf.ListRateBasedRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListRateBasedRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRateBasedRules, reflect.TypeOf((*MockWafClient)(nil).ListRateBasedRules), varargs...)
}

func (m *MockWafClient) ListRegexMatchSets(arg0 context.Context, arg1 *waf.ListRegexMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListRegexMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegexMatchSets, varargs...)
	ret0, _ := ret[0].(*waf.ListRegexMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListRegexMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegexMatchSets, reflect.TypeOf((*MockWafClient)(nil).ListRegexMatchSets), varargs...)
}

func (m *MockWafClient) ListRegexPatternSets(arg0 context.Context, arg1 *waf.ListRegexPatternSetsInput, arg2 ...func(*waf.Options)) (*waf.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegexPatternSets, varargs...)
	ret0, _ := ret[0].(*waf.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegexPatternSets, reflect.TypeOf((*MockWafClient)(nil).ListRegexPatternSets), varargs...)
}

func (m *MockWafClient) ListRuleGroups(arg0 context.Context, arg1 *waf.ListRuleGroupsInput, arg2 ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRuleGroups, varargs...)
	ret0, _ := ret[0].(*waf.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRuleGroups, reflect.TypeOf((*MockWafClient)(nil).ListRuleGroups), varargs...)
}

func (m *MockWafClient) ListRules(arg0 context.Context, arg1 *waf.ListRulesInput, arg2 ...func(*waf.Options)) (*waf.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRules, varargs...)
	ret0, _ := ret[0].(*waf.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRules, reflect.TypeOf((*MockWafClient)(nil).ListRules), varargs...)
}

func (m *MockWafClient) ListSizeConstraintSets(arg0 context.Context, arg1 *waf.ListSizeConstraintSetsInput, arg2 ...func(*waf.Options)) (*waf.ListSizeConstraintSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSizeConstraintSets, varargs...)
	ret0, _ := ret[0].(*waf.ListSizeConstraintSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListSizeConstraintSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSizeConstraintSets, reflect.TypeOf((*MockWafClient)(nil).ListSizeConstraintSets), varargs...)
}

func (m *MockWafClient) ListSqlInjectionMatchSets(arg0 context.Context, arg1 *waf.ListSqlInjectionMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSqlInjectionMatchSets, varargs...)
	ret0, _ := ret[0].(*waf.ListSqlInjectionMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListSqlInjectionMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSqlInjectionMatchSets, reflect.TypeOf((*MockWafClient)(nil).ListSqlInjectionMatchSets), varargs...)
}

func (m *MockWafClient) ListSubscribedRuleGroups(arg0 context.Context, arg1 *waf.ListSubscribedRuleGroupsInput, arg2 ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSubscribedRuleGroups, varargs...)
	ret0, _ := ret[0].(*waf.ListSubscribedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListSubscribedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSubscribedRuleGroups, reflect.TypeOf((*MockWafClient)(nil).ListSubscribedRuleGroups), varargs...)
}

func (m *MockWafClient) ListTagsForResource(arg0 context.Context, arg1 *waf.ListTagsForResourceInput, arg2 ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*waf.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockWafClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockWafClient) ListWebACLs(arg0 context.Context, arg1 *waf.ListWebACLsInput, arg2 ...func(*waf.Options)) (*waf.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWebACLs, varargs...)
	ret0, _ := ret[0].(*waf.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWebACLs, reflect.TypeOf((*MockWafClient)(nil).ListWebACLs), varargs...)
}

func (m *MockWafClient) ListXssMatchSets(arg0 context.Context, arg1 *waf.ListXssMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListXssMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListXssMatchSets, varargs...)
	ret0, _ := ret[0].(*waf.ListXssMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafClientMockRecorder) ListXssMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListXssMatchSets, reflect.TypeOf((*MockWafClient)(nil).ListXssMatchSets), varargs...)
}
