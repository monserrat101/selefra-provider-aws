package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	accessanalyzer "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	gomock "github.com/golang/mock/gomock"
)

type MockAccessanalyzerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAccessanalyzerClientMockRecorder
}

type MockAccessanalyzerClientMockRecorder struct {
	mock *MockAccessanalyzerClient
}

func NewMockAccessanalyzerClient(ctrl *gomock.Controller) *MockAccessanalyzerClient {
	mock := &MockAccessanalyzerClient{ctrl: ctrl}
	mock.recorder = &MockAccessanalyzerClientMockRecorder{mock}
	return mock
}

func (m *MockAccessanalyzerClient) EXPECT() *MockAccessanalyzerClientMockRecorder {
	return m.recorder
}

func (m *MockAccessanalyzerClient) GetAccessPreview(arg0 context.Context, arg1 *accessanalyzer.GetAccessPreviewInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAccessPreviewOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPreview, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetAccessPreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) GetAccessPreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPreview, reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetAccessPreview), varargs...)
}

func (m *MockAccessanalyzerClient) GetAnalyzedResource(arg0 context.Context, arg1 *accessanalyzer.GetAnalyzedResourceInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAnalyzedResource, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetAnalyzedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) GetAnalyzedResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAnalyzedResource, reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetAnalyzedResource), varargs...)
}

func (m *MockAccessanalyzerClient) GetAnalyzer(arg0 context.Context, arg1 *accessanalyzer.GetAnalyzerInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAnalyzer, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetAnalyzerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) GetAnalyzer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAnalyzer, reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetAnalyzer), varargs...)
}

func (m *MockAccessanalyzerClient) GetArchiveRule(arg0 context.Context, arg1 *accessanalyzer.GetArchiveRuleInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetArchiveRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetArchiveRule, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetArchiveRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) GetArchiveRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetArchiveRule, reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetArchiveRule), varargs...)
}

func (m *MockAccessanalyzerClient) GetFinding(arg0 context.Context, arg1 *accessanalyzer.GetFindingInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetFindingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFinding, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetFindingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) GetFinding(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFinding, reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetFinding), varargs...)
}

func (m *MockAccessanalyzerClient) GetGeneratedPolicy(arg0 context.Context, arg1 *accessanalyzer.GetGeneratedPolicyInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetGeneratedPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGeneratedPolicy, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetGeneratedPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) GetGeneratedPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGeneratedPolicy, reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetGeneratedPolicy), varargs...)
}

func (m *MockAccessanalyzerClient) ListAccessPreviewFindings(arg0 context.Context, arg1 *accessanalyzer.ListAccessPreviewFindingsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccessPreviewFindings, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAccessPreviewFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListAccessPreviewFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccessPreviewFindings, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAccessPreviewFindings), varargs...)
}

func (m *MockAccessanalyzerClient) ListAccessPreviews(arg0 context.Context, arg1 *accessanalyzer.ListAccessPreviewsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccessPreviews, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAccessPreviewsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListAccessPreviews(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccessPreviews, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAccessPreviews), varargs...)
}

func (m *MockAccessanalyzerClient) ListAnalyzedResources(arg0 context.Context, arg1 *accessanalyzer.ListAnalyzedResourcesInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAnalyzedResources, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAnalyzedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListAnalyzedResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAnalyzedResources, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAnalyzedResources), varargs...)
}

func (m *MockAccessanalyzerClient) ListAnalyzers(arg0 context.Context, arg1 *accessanalyzer.ListAnalyzersInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAnalyzers, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAnalyzersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListAnalyzers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAnalyzers, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAnalyzers), varargs...)
}

func (m *MockAccessanalyzerClient) ListArchiveRules(arg0 context.Context, arg1 *accessanalyzer.ListArchiveRulesInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListArchiveRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListArchiveRules, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListArchiveRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListArchiveRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListArchiveRules, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListArchiveRules), varargs...)
}

func (m *MockAccessanalyzerClient) ListFindings(arg0 context.Context, arg1 *accessanalyzer.ListFindingsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFindings, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFindings, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListFindings), varargs...)
}

func (m *MockAccessanalyzerClient) ListPolicyGenerations(arg0 context.Context, arg1 *accessanalyzer.ListPolicyGenerationsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListPolicyGenerationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicyGenerations, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListPolicyGenerationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListPolicyGenerations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicyGenerations, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListPolicyGenerations), varargs...)
}

func (m *MockAccessanalyzerClient) ListTagsForResource(arg0 context.Context, arg1 *accessanalyzer.ListTagsForResourceInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAccessanalyzerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListTagsForResource), varargs...)
}
