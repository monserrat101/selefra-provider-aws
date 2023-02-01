package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	wafv2 "github.com/aws/aws-sdk-go-v2/service/wafv2"
	gomock "github.com/golang/mock/gomock"
)

type MockWafv2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockWafv2ClientMockRecorder
}

type MockWafv2ClientMockRecorder struct {
	mock *MockWafv2Client
}

func NewMockWafv2Client(ctrl *gomock.Controller) *MockWafv2Client {
	mock := &MockWafv2Client{ctrl: ctrl}
	mock.recorder = &MockWafv2ClientMockRecorder{mock}
	return mock
}

func (m *MockWafv2Client) EXPECT() *MockWafv2ClientMockRecorder {
	return m.recorder
}

func (m *MockWafv2Client) DescribeManagedRuleGroup(arg0 context.Context, arg1 *wafv2.DescribeManagedRuleGroupInput, arg2 ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeManagedRuleGroup, varargs...)
	ret0, _ := ret[0].(*wafv2.DescribeManagedRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) DescribeManagedRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeManagedRuleGroup, reflect.TypeOf((*MockWafv2Client)(nil).DescribeManagedRuleGroup), varargs...)
}

func (m *MockWafv2Client) GetIPSet(arg0 context.Context, arg1 *wafv2.GetIPSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIPSet, varargs...)
	ret0, _ := ret[0].(*wafv2.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIPSet, reflect.TypeOf((*MockWafv2Client)(nil).GetIPSet), varargs...)
}

func (m *MockWafv2Client) GetLoggingConfiguration(arg0 context.Context, arg1 *wafv2.GetLoggingConfigurationInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoggingConfiguration, varargs...)
	ret0, _ := ret[0].(*wafv2.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoggingConfiguration, reflect.TypeOf((*MockWafv2Client)(nil).GetLoggingConfiguration), varargs...)
}

func (m *MockWafv2Client) GetManagedRuleSet(arg0 context.Context, arg1 *wafv2.GetManagedRuleSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetManagedRuleSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetManagedRuleSet, varargs...)
	ret0, _ := ret[0].(*wafv2.GetManagedRuleSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetManagedRuleSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetManagedRuleSet, reflect.TypeOf((*MockWafv2Client)(nil).GetManagedRuleSet), varargs...)
}

func (m *MockWafv2Client) GetMobileSdkRelease(arg0 context.Context, arg1 *wafv2.GetMobileSdkReleaseInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetMobileSdkReleaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMobileSdkRelease, varargs...)
	ret0, _ := ret[0].(*wafv2.GetMobileSdkReleaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetMobileSdkRelease(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMobileSdkRelease, reflect.TypeOf((*MockWafv2Client)(nil).GetMobileSdkRelease), varargs...)
}

func (m *MockWafv2Client) GetPermissionPolicy(arg0 context.Context, arg1 *wafv2.GetPermissionPolicyInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPermissionPolicy, varargs...)
	ret0, _ := ret[0].(*wafv2.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPermissionPolicy, reflect.TypeOf((*MockWafv2Client)(nil).GetPermissionPolicy), varargs...)
}

func (m *MockWafv2Client) GetRateBasedStatementManagedKeys(arg0 context.Context, arg1 *wafv2.GetRateBasedStatementManagedKeysInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRateBasedStatementManagedKeys, varargs...)
	ret0, _ := ret[0].(*wafv2.GetRateBasedStatementManagedKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetRateBasedStatementManagedKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRateBasedStatementManagedKeys, reflect.TypeOf((*MockWafv2Client)(nil).GetRateBasedStatementManagedKeys), varargs...)
}

func (m *MockWafv2Client) GetRegexPatternSet(arg0 context.Context, arg1 *wafv2.GetRegexPatternSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegexPatternSet, varargs...)
	ret0, _ := ret[0].(*wafv2.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegexPatternSet, reflect.TypeOf((*MockWafv2Client)(nil).GetRegexPatternSet), varargs...)
}

func (m *MockWafv2Client) GetRuleGroup(arg0 context.Context, arg1 *wafv2.GetRuleGroupInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRuleGroup, varargs...)
	ret0, _ := ret[0].(*wafv2.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRuleGroup, reflect.TypeOf((*MockWafv2Client)(nil).GetRuleGroup), varargs...)
}

func (m *MockWafv2Client) GetSampledRequests(arg0 context.Context, arg1 *wafv2.GetSampledRequestsInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetSampledRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSampledRequests, varargs...)
	ret0, _ := ret[0].(*wafv2.GetSampledRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetSampledRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSampledRequests, reflect.TypeOf((*MockWafv2Client)(nil).GetSampledRequests), varargs...)
}

func (m *MockWafv2Client) GetWebACL(arg0 context.Context, arg1 *wafv2.GetWebACLInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWebACL, varargs...)
	ret0, _ := ret[0].(*wafv2.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWebACL, reflect.TypeOf((*MockWafv2Client)(nil).GetWebACL), varargs...)
}

func (m *MockWafv2Client) GetWebACLForResource(arg0 context.Context, arg1 *wafv2.GetWebACLForResourceInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWebACLForResource, varargs...)
	ret0, _ := ret[0].(*wafv2.GetWebACLForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) GetWebACLForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWebACLForResource, reflect.TypeOf((*MockWafv2Client)(nil).GetWebACLForResource), varargs...)
}

func (m *MockWafv2Client) ListAvailableManagedRuleGroupVersions(arg0 context.Context, arg1 *wafv2.ListAvailableManagedRuleGroupVersionsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAvailableManagedRuleGroupVersions, varargs...)
	ret0, _ := ret[0].(*wafv2.ListAvailableManagedRuleGroupVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListAvailableManagedRuleGroupVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAvailableManagedRuleGroupVersions, reflect.TypeOf((*MockWafv2Client)(nil).ListAvailableManagedRuleGroupVersions), varargs...)
}

func (m *MockWafv2Client) ListAvailableManagedRuleGroups(arg0 context.Context, arg1 *wafv2.ListAvailableManagedRuleGroupsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAvailableManagedRuleGroups, varargs...)
	ret0, _ := ret[0].(*wafv2.ListAvailableManagedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListAvailableManagedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAvailableManagedRuleGroups, reflect.TypeOf((*MockWafv2Client)(nil).ListAvailableManagedRuleGroups), varargs...)
}

func (m *MockWafv2Client) ListIPSets(arg0 context.Context, arg1 *wafv2.ListIPSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIPSets, varargs...)
	ret0, _ := ret[0].(*wafv2.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIPSets, reflect.TypeOf((*MockWafv2Client)(nil).ListIPSets), varargs...)
}

func (m *MockWafv2Client) ListLoggingConfigurations(arg0 context.Context, arg1 *wafv2.ListLoggingConfigurationsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLoggingConfigurations, varargs...)
	ret0, _ := ret[0].(*wafv2.ListLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLoggingConfigurations, reflect.TypeOf((*MockWafv2Client)(nil).ListLoggingConfigurations), varargs...)
}

func (m *MockWafv2Client) ListManagedRuleSets(arg0 context.Context, arg1 *wafv2.ListManagedRuleSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListManagedRuleSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListManagedRuleSets, varargs...)
	ret0, _ := ret[0].(*wafv2.ListManagedRuleSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListManagedRuleSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListManagedRuleSets, reflect.TypeOf((*MockWafv2Client)(nil).ListManagedRuleSets), varargs...)
}

func (m *MockWafv2Client) ListMobileSdkReleases(arg0 context.Context, arg1 *wafv2.ListMobileSdkReleasesInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListMobileSdkReleasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMobileSdkReleases, varargs...)
	ret0, _ := ret[0].(*wafv2.ListMobileSdkReleasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListMobileSdkReleases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMobileSdkReleases, reflect.TypeOf((*MockWafv2Client)(nil).ListMobileSdkReleases), varargs...)
}

func (m *MockWafv2Client) ListRegexPatternSets(arg0 context.Context, arg1 *wafv2.ListRegexPatternSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegexPatternSets, varargs...)
	ret0, _ := ret[0].(*wafv2.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegexPatternSets, reflect.TypeOf((*MockWafv2Client)(nil).ListRegexPatternSets), varargs...)
}

func (m *MockWafv2Client) ListResourcesForWebACL(arg0 context.Context, arg1 *wafv2.ListResourcesForWebACLInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourcesForWebACL, varargs...)
	ret0, _ := ret[0].(*wafv2.ListResourcesForWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListResourcesForWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourcesForWebACL, reflect.TypeOf((*MockWafv2Client)(nil).ListResourcesForWebACL), varargs...)
}

func (m *MockWafv2Client) ListRuleGroups(arg0 context.Context, arg1 *wafv2.ListRuleGroupsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRuleGroups, varargs...)
	ret0, _ := ret[0].(*wafv2.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRuleGroups, reflect.TypeOf((*MockWafv2Client)(nil).ListRuleGroups), varargs...)
}

func (m *MockWafv2Client) ListTagsForResource(arg0 context.Context, arg1 *wafv2.ListTagsForResourceInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*wafv2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockWafv2Client)(nil).ListTagsForResource), varargs...)
}

func (m *MockWafv2Client) ListWebACLs(arg0 context.Context, arg1 *wafv2.ListWebACLsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWebACLs, varargs...)
	ret0, _ := ret[0].(*wafv2.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafv2ClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWebACLs, reflect.TypeOf((*MockWafv2Client)(nil).ListWebACLs), varargs...)
}
