package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	xray "github.com/aws/aws-sdk-go-v2/service/xray"
	gomock "github.com/golang/mock/gomock"
)

type MockXrayClient struct {
	ctrl		*gomock.Controller
	recorder	*MockXrayClientMockRecorder
}

type MockXrayClientMockRecorder struct {
	mock *MockXrayClient
}

func NewMockXrayClient(ctrl *gomock.Controller) *MockXrayClient {
	mock := &MockXrayClient{ctrl: ctrl}
	mock.recorder = &MockXrayClientMockRecorder{mock}
	return mock
}

func (m *MockXrayClient) EXPECT() *MockXrayClientMockRecorder {
	return m.recorder
}

func (m *MockXrayClient) BatchGetTraces(arg0 context.Context, arg1 *xray.BatchGetTracesInput, arg2 ...func(*xray.Options)) (*xray.BatchGetTracesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetTraces, varargs...)
	ret0, _ := ret[0].(*xray.BatchGetTracesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) BatchGetTraces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetTraces, reflect.TypeOf((*MockXrayClient)(nil).BatchGetTraces), varargs...)
}

func (m *MockXrayClient) GetEncryptionConfig(arg0 context.Context, arg1 *xray.GetEncryptionConfigInput, arg2 ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEncryptionConfig, varargs...)
	ret0, _ := ret[0].(*xray.GetEncryptionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetEncryptionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEncryptionConfig, reflect.TypeOf((*MockXrayClient)(nil).GetEncryptionConfig), varargs...)
}

func (m *MockXrayClient) GetGroup(arg0 context.Context, arg1 *xray.GetGroupInput, arg2 ...func(*xray.Options)) (*xray.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroup, varargs...)
	ret0, _ := ret[0].(*xray.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroup, reflect.TypeOf((*MockXrayClient)(nil).GetGroup), varargs...)
}

func (m *MockXrayClient) GetGroups(arg0 context.Context, arg1 *xray.GetGroupsInput, arg2 ...func(*xray.Options)) (*xray.GetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroups, varargs...)
	ret0, _ := ret[0].(*xray.GetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroups, reflect.TypeOf((*MockXrayClient)(nil).GetGroups), varargs...)
}

func (m *MockXrayClient) GetInsight(arg0 context.Context, arg1 *xray.GetInsightInput, arg2 ...func(*xray.Options)) (*xray.GetInsightOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsight, varargs...)
	ret0, _ := ret[0].(*xray.GetInsightOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetInsight(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsight, reflect.TypeOf((*MockXrayClient)(nil).GetInsight), varargs...)
}

func (m *MockXrayClient) GetInsightEvents(arg0 context.Context, arg1 *xray.GetInsightEventsInput, arg2 ...func(*xray.Options)) (*xray.GetInsightEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsightEvents, varargs...)
	ret0, _ := ret[0].(*xray.GetInsightEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetInsightEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsightEvents, reflect.TypeOf((*MockXrayClient)(nil).GetInsightEvents), varargs...)
}

func (m *MockXrayClient) GetInsightImpactGraph(arg0 context.Context, arg1 *xray.GetInsightImpactGraphInput, arg2 ...func(*xray.Options)) (*xray.GetInsightImpactGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsightImpactGraph, varargs...)
	ret0, _ := ret[0].(*xray.GetInsightImpactGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetInsightImpactGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsightImpactGraph, reflect.TypeOf((*MockXrayClient)(nil).GetInsightImpactGraph), varargs...)
}

func (m *MockXrayClient) GetInsightSummaries(arg0 context.Context, arg1 *xray.GetInsightSummariesInput, arg2 ...func(*xray.Options)) (*xray.GetInsightSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsightSummaries, varargs...)
	ret0, _ := ret[0].(*xray.GetInsightSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetInsightSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsightSummaries, reflect.TypeOf((*MockXrayClient)(nil).GetInsightSummaries), varargs...)
}

func (m *MockXrayClient) GetSamplingRules(arg0 context.Context, arg1 *xray.GetSamplingRulesInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSamplingRules, varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetSamplingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSamplingRules, reflect.TypeOf((*MockXrayClient)(nil).GetSamplingRules), varargs...)
}

func (m *MockXrayClient) GetSamplingStatisticSummaries(arg0 context.Context, arg1 *xray.GetSamplingStatisticSummariesInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingStatisticSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSamplingStatisticSummaries, varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingStatisticSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetSamplingStatisticSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSamplingStatisticSummaries, reflect.TypeOf((*MockXrayClient)(nil).GetSamplingStatisticSummaries), varargs...)
}

func (m *MockXrayClient) GetSamplingTargets(arg0 context.Context, arg1 *xray.GetSamplingTargetsInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSamplingTargets, varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetSamplingTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSamplingTargets, reflect.TypeOf((*MockXrayClient)(nil).GetSamplingTargets), varargs...)
}

func (m *MockXrayClient) GetServiceGraph(arg0 context.Context, arg1 *xray.GetServiceGraphInput, arg2 ...func(*xray.Options)) (*xray.GetServiceGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceGraph, varargs...)
	ret0, _ := ret[0].(*xray.GetServiceGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetServiceGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceGraph, reflect.TypeOf((*MockXrayClient)(nil).GetServiceGraph), varargs...)
}

func (m *MockXrayClient) GetTimeSeriesServiceStatistics(arg0 context.Context, arg1 *xray.GetTimeSeriesServiceStatisticsInput, arg2 ...func(*xray.Options)) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTimeSeriesServiceStatistics, varargs...)
	ret0, _ := ret[0].(*xray.GetTimeSeriesServiceStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetTimeSeriesServiceStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTimeSeriesServiceStatistics, reflect.TypeOf((*MockXrayClient)(nil).GetTimeSeriesServiceStatistics), varargs...)
}

func (m *MockXrayClient) GetTraceGraph(arg0 context.Context, arg1 *xray.GetTraceGraphInput, arg2 ...func(*xray.Options)) (*xray.GetTraceGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTraceGraph, varargs...)
	ret0, _ := ret[0].(*xray.GetTraceGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetTraceGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTraceGraph, reflect.TypeOf((*MockXrayClient)(nil).GetTraceGraph), varargs...)
}

func (m *MockXrayClient) GetTraceSummaries(arg0 context.Context, arg1 *xray.GetTraceSummariesInput, arg2 ...func(*xray.Options)) (*xray.GetTraceSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTraceSummaries, varargs...)
	ret0, _ := ret[0].(*xray.GetTraceSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetTraceSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTraceSummaries, reflect.TypeOf((*MockXrayClient)(nil).GetTraceSummaries), varargs...)
}

func (m *MockXrayClient) ListResourcePolicies(arg0 context.Context, arg1 *xray.ListResourcePoliciesInput, arg2 ...func(*xray.Options)) (*xray.ListResourcePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourcePolicies, varargs...)
	ret0, _ := ret[0].(*xray.ListResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) ListResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourcePolicies, reflect.TypeOf((*MockXrayClient)(nil).ListResourcePolicies), varargs...)
}

func (m *MockXrayClient) ListTagsForResource(arg0 context.Context, arg1 *xray.ListTagsForResourceInput, arg2 ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*xray.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockXrayClient)(nil).ListTagsForResource), varargs...)
}
