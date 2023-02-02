package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	inspector "github.com/aws/aws-sdk-go-v2/service/inspector"
	gomock "github.com/golang/mock/gomock"
)

type MockInspectorClient struct {
	ctrl		*gomock.Controller
	recorder	*MockInspectorClientMockRecorder
}

type MockInspectorClientMockRecorder struct {
	mock *MockInspectorClient
}

func NewMockInspectorClient(ctrl *gomock.Controller) *MockInspectorClient {
	mock := &MockInspectorClient{ctrl: ctrl}
	mock.recorder = &MockInspectorClientMockRecorder{mock}
	return mock
}

func (m *MockInspectorClient) EXPECT() *MockInspectorClientMockRecorder {
	return m.recorder
}

func (m *MockInspectorClient) DescribeAssessmentRuns(arg0 context.Context, arg1 *inspector.DescribeAssessmentRunsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeAssessmentRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAssessmentRuns, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeAssessmentRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeAssessmentRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAssessmentRuns, reflect.TypeOf((*MockInspectorClient)(nil).DescribeAssessmentRuns), varargs...)
}

func (m *MockInspectorClient) DescribeAssessmentTargets(arg0 context.Context, arg1 *inspector.DescribeAssessmentTargetsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeAssessmentTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAssessmentTargets, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeAssessmentTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeAssessmentTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAssessmentTargets, reflect.TypeOf((*MockInspectorClient)(nil).DescribeAssessmentTargets), varargs...)
}

func (m *MockInspectorClient) DescribeAssessmentTemplates(arg0 context.Context, arg1 *inspector.DescribeAssessmentTemplatesInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeAssessmentTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAssessmentTemplates, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeAssessmentTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeAssessmentTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAssessmentTemplates, reflect.TypeOf((*MockInspectorClient)(nil).DescribeAssessmentTemplates), varargs...)
}

func (m *MockInspectorClient) DescribeCrossAccountAccessRole(arg0 context.Context, arg1 *inspector.DescribeCrossAccountAccessRoleInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCrossAccountAccessRole, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeCrossAccountAccessRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeCrossAccountAccessRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCrossAccountAccessRole, reflect.TypeOf((*MockInspectorClient)(nil).DescribeCrossAccountAccessRole), varargs...)
}

func (m *MockInspectorClient) DescribeExclusions(arg0 context.Context, arg1 *inspector.DescribeExclusionsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeExclusionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExclusions, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeExclusionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeExclusions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExclusions, reflect.TypeOf((*MockInspectorClient)(nil).DescribeExclusions), varargs...)
}

func (m *MockInspectorClient) DescribeFindings(arg0 context.Context, arg1 *inspector.DescribeFindingsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFindings, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFindings, reflect.TypeOf((*MockInspectorClient)(nil).DescribeFindings), varargs...)
}

func (m *MockInspectorClient) DescribeResourceGroups(arg0 context.Context, arg1 *inspector.DescribeResourceGroupsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeResourceGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeResourceGroups, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeResourceGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeResourceGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeResourceGroups, reflect.TypeOf((*MockInspectorClient)(nil).DescribeResourceGroups), varargs...)
}

func (m *MockInspectorClient) DescribeRulesPackages(arg0 context.Context, arg1 *inspector.DescribeRulesPackagesInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeRulesPackagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRulesPackages, varargs...)
	ret0, _ := ret[0].(*inspector.DescribeRulesPackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeRulesPackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRulesPackages, reflect.TypeOf((*MockInspectorClient)(nil).DescribeRulesPackages), varargs...)
}

func (m *MockInspectorClient) GetAssessmentReport(arg0 context.Context, arg1 *inspector.GetAssessmentReportInput, arg2 ...func(*inspector.Options)) (*inspector.GetAssessmentReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAssessmentReport, varargs...)
	ret0, _ := ret[0].(*inspector.GetAssessmentReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) GetAssessmentReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAssessmentReport, reflect.TypeOf((*MockInspectorClient)(nil).GetAssessmentReport), varargs...)
}

func (m *MockInspectorClient) GetExclusionsPreview(arg0 context.Context, arg1 *inspector.GetExclusionsPreviewInput, arg2 ...func(*inspector.Options)) (*inspector.GetExclusionsPreviewOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetExclusionsPreview, varargs...)
	ret0, _ := ret[0].(*inspector.GetExclusionsPreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) GetExclusionsPreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetExclusionsPreview, reflect.TypeOf((*MockInspectorClient)(nil).GetExclusionsPreview), varargs...)
}

func (m *MockInspectorClient) GetTelemetryMetadata(arg0 context.Context, arg1 *inspector.GetTelemetryMetadataInput, arg2 ...func(*inspector.Options)) (*inspector.GetTelemetryMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTelemetryMetadata, varargs...)
	ret0, _ := ret[0].(*inspector.GetTelemetryMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) GetTelemetryMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTelemetryMetadata, reflect.TypeOf((*MockInspectorClient)(nil).GetTelemetryMetadata), varargs...)
}

func (m *MockInspectorClient) ListAssessmentRunAgents(arg0 context.Context, arg1 *inspector.ListAssessmentRunAgentsInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentRunAgentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssessmentRunAgents, varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentRunAgentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListAssessmentRunAgents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssessmentRunAgents, reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentRunAgents), varargs...)
}

func (m *MockInspectorClient) ListAssessmentRuns(arg0 context.Context, arg1 *inspector.ListAssessmentRunsInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssessmentRuns, varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListAssessmentRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssessmentRuns, reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentRuns), varargs...)
}

func (m *MockInspectorClient) ListAssessmentTargets(arg0 context.Context, arg1 *inspector.ListAssessmentTargetsInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssessmentTargets, varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListAssessmentTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssessmentTargets, reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentTargets), varargs...)
}

func (m *MockInspectorClient) ListAssessmentTemplates(arg0 context.Context, arg1 *inspector.ListAssessmentTemplatesInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssessmentTemplates, varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListAssessmentTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssessmentTemplates, reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentTemplates), varargs...)
}

func (m *MockInspectorClient) ListEventSubscriptions(arg0 context.Context, arg1 *inspector.ListEventSubscriptionsInput, arg2 ...func(*inspector.Options)) (*inspector.ListEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEventSubscriptions, varargs...)
	ret0, _ := ret[0].(*inspector.ListEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEventSubscriptions, reflect.TypeOf((*MockInspectorClient)(nil).ListEventSubscriptions), varargs...)
}

func (m *MockInspectorClient) ListExclusions(arg0 context.Context, arg1 *inspector.ListExclusionsInput, arg2 ...func(*inspector.Options)) (*inspector.ListExclusionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExclusions, varargs...)
	ret0, _ := ret[0].(*inspector.ListExclusionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListExclusions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExclusions, reflect.TypeOf((*MockInspectorClient)(nil).ListExclusions), varargs...)
}

func (m *MockInspectorClient) ListFindings(arg0 context.Context, arg1 *inspector.ListFindingsInput, arg2 ...func(*inspector.Options)) (*inspector.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFindings, varargs...)
	ret0, _ := ret[0].(*inspector.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFindings, reflect.TypeOf((*MockInspectorClient)(nil).ListFindings), varargs...)
}

func (m *MockInspectorClient) ListRulesPackages(arg0 context.Context, arg1 *inspector.ListRulesPackagesInput, arg2 ...func(*inspector.Options)) (*inspector.ListRulesPackagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRulesPackages, varargs...)
	ret0, _ := ret[0].(*inspector.ListRulesPackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListRulesPackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRulesPackages, reflect.TypeOf((*MockInspectorClient)(nil).ListRulesPackages), varargs...)
}

func (m *MockInspectorClient) ListTagsForResource(arg0 context.Context, arg1 *inspector.ListTagsForResourceInput, arg2 ...func(*inspector.Options)) (*inspector.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*inspector.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockInspectorClient)(nil).ListTagsForResource), varargs...)
}
