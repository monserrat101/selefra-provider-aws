package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	wafregional "github.com/aws/aws-sdk-go-v2/service/wafregional"
	gomock "github.com/golang/mock/gomock"
)

type MockWafregionalClient struct {
	ctrl		*gomock.Controller
	recorder	*MockWafregionalClientMockRecorder
}

type MockWafregionalClientMockRecorder struct {
	mock *MockWafregionalClient
}

func NewMockWafregionalClient(ctrl *gomock.Controller) *MockWafregionalClient {
	mock := &MockWafregionalClient{ctrl: ctrl}
	mock.recorder = &MockWafregionalClientMockRecorder{mock}
	return mock
}

func (m *MockWafregionalClient) EXPECT() *MockWafregionalClientMockRecorder {
	return m.recorder
}

func (m *MockWafregionalClient) GetByteMatchSet(arg0 context.Context, arg1 *wafregional.GetByteMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetByteMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetByteMatchSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetByteMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetByteMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetByteMatchSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetByteMatchSet), varargs...)
}

func (m *MockWafregionalClient) GetChangeToken(arg0 context.Context, arg1 *wafregional.GetChangeTokenInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetChangeTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetChangeToken, varargs...)
	ret0, _ := ret[0].(*wafregional.GetChangeTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetChangeToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetChangeToken, reflect.TypeOf((*MockWafregionalClient)(nil).GetChangeToken), varargs...)
}

func (m *MockWafregionalClient) GetChangeTokenStatus(arg0 context.Context, arg1 *wafregional.GetChangeTokenStatusInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetChangeTokenStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetChangeTokenStatus, varargs...)
	ret0, _ := ret[0].(*wafregional.GetChangeTokenStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetChangeTokenStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetChangeTokenStatus, reflect.TypeOf((*MockWafregionalClient)(nil).GetChangeTokenStatus), varargs...)
}

func (m *MockWafregionalClient) GetGeoMatchSet(arg0 context.Context, arg1 *wafregional.GetGeoMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetGeoMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGeoMatchSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetGeoMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetGeoMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGeoMatchSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetGeoMatchSet), varargs...)
}

func (m *MockWafregionalClient) GetIPSet(arg0 context.Context, arg1 *wafregional.GetIPSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIPSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIPSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetIPSet), varargs...)
}

func (m *MockWafregionalClient) GetLoggingConfiguration(arg0 context.Context, arg1 *wafregional.GetLoggingConfigurationInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoggingConfiguration, varargs...)
	ret0, _ := ret[0].(*wafregional.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoggingConfiguration, reflect.TypeOf((*MockWafregionalClient)(nil).GetLoggingConfiguration), varargs...)
}

func (m *MockWafregionalClient) GetPermissionPolicy(arg0 context.Context, arg1 *wafregional.GetPermissionPolicyInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPermissionPolicy, varargs...)
	ret0, _ := ret[0].(*wafregional.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPermissionPolicy, reflect.TypeOf((*MockWafregionalClient)(nil).GetPermissionPolicy), varargs...)
}

func (m *MockWafregionalClient) GetRateBasedRule(arg0 context.Context, arg1 *wafregional.GetRateBasedRuleInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRateBasedRule, varargs...)
	ret0, _ := ret[0].(*wafregional.GetRateBasedRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetRateBasedRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRateBasedRule, reflect.TypeOf((*MockWafregionalClient)(nil).GetRateBasedRule), varargs...)
}

func (m *MockWafregionalClient) GetRateBasedRuleManagedKeys(arg0 context.Context, arg1 *wafregional.GetRateBasedRuleManagedKeysInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleManagedKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRateBasedRuleManagedKeys, varargs...)
	ret0, _ := ret[0].(*wafregional.GetRateBasedRuleManagedKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetRateBasedRuleManagedKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRateBasedRuleManagedKeys, reflect.TypeOf((*MockWafregionalClient)(nil).GetRateBasedRuleManagedKeys), varargs...)
}

func (m *MockWafregionalClient) GetRegexMatchSet(arg0 context.Context, arg1 *wafregional.GetRegexMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRegexMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegexMatchSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetRegexMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetRegexMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegexMatchSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetRegexMatchSet), varargs...)
}

func (m *MockWafregionalClient) GetRegexPatternSet(arg0 context.Context, arg1 *wafregional.GetRegexPatternSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegexPatternSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegexPatternSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetRegexPatternSet), varargs...)
}

func (m *MockWafregionalClient) GetRule(arg0 context.Context, arg1 *wafregional.GetRuleInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRule, varargs...)
	ret0, _ := ret[0].(*wafregional.GetRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRule, reflect.TypeOf((*MockWafregionalClient)(nil).GetRule), varargs...)
}

func (m *MockWafregionalClient) GetRuleGroup(arg0 context.Context, arg1 *wafregional.GetRuleGroupInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRuleGroup, varargs...)
	ret0, _ := ret[0].(*wafregional.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRuleGroup, reflect.TypeOf((*MockWafregionalClient)(nil).GetRuleGroup), varargs...)
}

func (m *MockWafregionalClient) GetSampledRequests(arg0 context.Context, arg1 *wafregional.GetSampledRequestsInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetSampledRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSampledRequests, varargs...)
	ret0, _ := ret[0].(*wafregional.GetSampledRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetSampledRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSampledRequests, reflect.TypeOf((*MockWafregionalClient)(nil).GetSampledRequests), varargs...)
}

func (m *MockWafregionalClient) GetSizeConstraintSet(arg0 context.Context, arg1 *wafregional.GetSizeConstraintSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetSizeConstraintSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSizeConstraintSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetSizeConstraintSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetSizeConstraintSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSizeConstraintSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetSizeConstraintSet), varargs...)
}

func (m *MockWafregionalClient) GetSqlInjectionMatchSet(arg0 context.Context, arg1 *wafregional.GetSqlInjectionMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetSqlInjectionMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSqlInjectionMatchSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetSqlInjectionMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetSqlInjectionMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSqlInjectionMatchSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetSqlInjectionMatchSet), varargs...)
}

func (m *MockWafregionalClient) GetWebACL(arg0 context.Context, arg1 *wafregional.GetWebACLInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWebACL, varargs...)
	ret0, _ := ret[0].(*wafregional.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWebACL, reflect.TypeOf((*MockWafregionalClient)(nil).GetWebACL), varargs...)
}

func (m *MockWafregionalClient) GetWebACLForResource(arg0 context.Context, arg1 *wafregional.GetWebACLForResourceInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetWebACLForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWebACLForResource, varargs...)
	ret0, _ := ret[0].(*wafregional.GetWebACLForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetWebACLForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWebACLForResource, reflect.TypeOf((*MockWafregionalClient)(nil).GetWebACLForResource), varargs...)
}

func (m *MockWafregionalClient) GetXssMatchSet(arg0 context.Context, arg1 *wafregional.GetXssMatchSetInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetXssMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetXssMatchSet, varargs...)
	ret0, _ := ret[0].(*wafregional.GetXssMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) GetXssMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetXssMatchSet, reflect.TypeOf((*MockWafregionalClient)(nil).GetXssMatchSet), varargs...)
}

func (m *MockWafregionalClient) ListActivatedRulesInRuleGroup(arg0 context.Context, arg1 *wafregional.ListActivatedRulesInRuleGroupInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListActivatedRulesInRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActivatedRulesInRuleGroup, varargs...)
	ret0, _ := ret[0].(*wafregional.ListActivatedRulesInRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListActivatedRulesInRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActivatedRulesInRuleGroup, reflect.TypeOf((*MockWafregionalClient)(nil).ListActivatedRulesInRuleGroup), varargs...)
}

func (m *MockWafregionalClient) ListByteMatchSets(arg0 context.Context, arg1 *wafregional.ListByteMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListByteMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListByteMatchSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListByteMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListByteMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListByteMatchSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListByteMatchSets), varargs...)
}

func (m *MockWafregionalClient) ListGeoMatchSets(arg0 context.Context, arg1 *wafregional.ListGeoMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListGeoMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGeoMatchSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListGeoMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListGeoMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGeoMatchSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListGeoMatchSets), varargs...)
}

func (m *MockWafregionalClient) ListIPSets(arg0 context.Context, arg1 *wafregional.ListIPSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIPSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIPSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListIPSets), varargs...)
}

func (m *MockWafregionalClient) ListLoggingConfigurations(arg0 context.Context, arg1 *wafregional.ListLoggingConfigurationsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLoggingConfigurations, varargs...)
	ret0, _ := ret[0].(*wafregional.ListLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLoggingConfigurations, reflect.TypeOf((*MockWafregionalClient)(nil).ListLoggingConfigurations), varargs...)
}

func (m *MockWafregionalClient) ListRateBasedRules(arg0 context.Context, arg1 *wafregional.ListRateBasedRulesInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRateBasedRules, varargs...)
	ret0, _ := ret[0].(*wafregional.ListRateBasedRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListRateBasedRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRateBasedRules, reflect.TypeOf((*MockWafregionalClient)(nil).ListRateBasedRules), varargs...)
}

func (m *MockWafregionalClient) ListRegexMatchSets(arg0 context.Context, arg1 *wafregional.ListRegexMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRegexMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegexMatchSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListRegexMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListRegexMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegexMatchSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListRegexMatchSets), varargs...)
}

func (m *MockWafregionalClient) ListRegexPatternSets(arg0 context.Context, arg1 *wafregional.ListRegexPatternSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegexPatternSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegexPatternSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListRegexPatternSets), varargs...)
}

func (m *MockWafregionalClient) ListResourcesForWebACL(arg0 context.Context, arg1 *wafregional.ListResourcesForWebACLInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListResourcesForWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourcesForWebACL, varargs...)
	ret0, _ := ret[0].(*wafregional.ListResourcesForWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListResourcesForWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourcesForWebACL, reflect.TypeOf((*MockWafregionalClient)(nil).ListResourcesForWebACL), varargs...)
}

func (m *MockWafregionalClient) ListRuleGroups(arg0 context.Context, arg1 *wafregional.ListRuleGroupsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRuleGroups, varargs...)
	ret0, _ := ret[0].(*wafregional.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRuleGroups, reflect.TypeOf((*MockWafregionalClient)(nil).ListRuleGroups), varargs...)
}

func (m *MockWafregionalClient) ListRules(arg0 context.Context, arg1 *wafregional.ListRulesInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRules, varargs...)
	ret0, _ := ret[0].(*wafregional.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRules, reflect.TypeOf((*MockWafregionalClient)(nil).ListRules), varargs...)
}

func (m *MockWafregionalClient) ListSizeConstraintSets(arg0 context.Context, arg1 *wafregional.ListSizeConstraintSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListSizeConstraintSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSizeConstraintSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListSizeConstraintSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListSizeConstraintSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSizeConstraintSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListSizeConstraintSets), varargs...)
}

func (m *MockWafregionalClient) ListSqlInjectionMatchSets(arg0 context.Context, arg1 *wafregional.ListSqlInjectionMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListSqlInjectionMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSqlInjectionMatchSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListSqlInjectionMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListSqlInjectionMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSqlInjectionMatchSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListSqlInjectionMatchSets), varargs...)
}

func (m *MockWafregionalClient) ListSubscribedRuleGroups(arg0 context.Context, arg1 *wafregional.ListSubscribedRuleGroupsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListSubscribedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSubscribedRuleGroups, varargs...)
	ret0, _ := ret[0].(*wafregional.ListSubscribedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListSubscribedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSubscribedRuleGroups, reflect.TypeOf((*MockWafregionalClient)(nil).ListSubscribedRuleGroups), varargs...)
}

func (m *MockWafregionalClient) ListTagsForResource(arg0 context.Context, arg1 *wafregional.ListTagsForResourceInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*wafregional.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockWafregionalClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockWafregionalClient) ListWebACLs(arg0 context.Context, arg1 *wafregional.ListWebACLsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWebACLs, varargs...)
	ret0, _ := ret[0].(*wafregional.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWebACLs, reflect.TypeOf((*MockWafregionalClient)(nil).ListWebACLs), varargs...)
}

func (m *MockWafregionalClient) ListXssMatchSets(arg0 context.Context, arg1 *wafregional.ListXssMatchSetsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListXssMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListXssMatchSets, varargs...)
	ret0, _ := ret[0].(*wafregional.ListXssMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafregionalClientMockRecorder) ListXssMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListXssMatchSets, reflect.TypeOf((*MockWafregionalClient)(nil).ListXssMatchSets), varargs...)
}
