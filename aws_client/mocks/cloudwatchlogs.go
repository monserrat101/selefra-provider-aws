package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudwatchlogsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudwatchlogsClientMockRecorder
}

type MockCloudwatchlogsClientMockRecorder struct {
	mock *MockCloudwatchlogsClient
}

func NewMockCloudwatchlogsClient(ctrl *gomock.Controller) *MockCloudwatchlogsClient {
	mock := &MockCloudwatchlogsClient{ctrl: ctrl}
	mock.recorder = &MockCloudwatchlogsClientMockRecorder{mock}
	return mock
}

func (m *MockCloudwatchlogsClient) EXPECT() *MockCloudwatchlogsClientMockRecorder {
	return m.recorder
}

func (m *MockCloudwatchlogsClient) DescribeDestinations(arg0 context.Context, arg1 *cloudwatchlogs.DescribeDestinationsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDestinations, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDestinations, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeDestinations), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeExportTasks(arg0 context.Context, arg1 *cloudwatchlogs.DescribeExportTasksInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExportTasks, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeExportTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeExportTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExportTasks, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeExportTasks), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeLogGroups(arg0 context.Context, arg1 *cloudwatchlogs.DescribeLogGroupsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLogGroups, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeLogGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLogGroups, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeLogGroups), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeLogStreams(arg0 context.Context, arg1 *cloudwatchlogs.DescribeLogStreamsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLogStreams, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeLogStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLogStreams, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeLogStreams), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeMetricFilters(arg0 context.Context, arg1 *cloudwatchlogs.DescribeMetricFiltersInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMetricFilters, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeMetricFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeMetricFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMetricFilters, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeMetricFilters), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeQueries(arg0 context.Context, arg1 *cloudwatchlogs.DescribeQueriesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeQueries, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeQueries, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeQueries), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeQueryDefinitions(arg0 context.Context, arg1 *cloudwatchlogs.DescribeQueryDefinitionsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeQueryDefinitions, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeQueryDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeQueryDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeQueryDefinitions, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeQueryDefinitions), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeResourcePolicies(arg0 context.Context, arg1 *cloudwatchlogs.DescribeResourcePoliciesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeResourcePolicies, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeResourcePolicies, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeResourcePolicies), varargs...)
}

func (m *MockCloudwatchlogsClient) DescribeSubscriptionFilters(arg0 context.Context, arg1 *cloudwatchlogs.DescribeSubscriptionFiltersInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSubscriptionFilters, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeSubscriptionFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) DescribeSubscriptionFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSubscriptionFilters, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeSubscriptionFilters), varargs...)
}

func (m *MockCloudwatchlogsClient) GetDataProtectionPolicy(arg0 context.Context, arg1 *cloudwatchlogs.GetDataProtectionPolicyInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDataProtectionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataProtectionPolicy, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetDataProtectionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) GetDataProtectionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataProtectionPolicy, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetDataProtectionPolicy), varargs...)
}

func (m *MockCloudwatchlogsClient) GetLogEvents(arg0 context.Context, arg1 *cloudwatchlogs.GetLogEventsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLogEvents, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) GetLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLogEvents, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetLogEvents), varargs...)
}

func (m *MockCloudwatchlogsClient) GetLogGroupFields(arg0 context.Context, arg1 *cloudwatchlogs.GetLogGroupFieldsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLogGroupFields, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogGroupFieldsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) GetLogGroupFields(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLogGroupFields, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetLogGroupFields), varargs...)
}

func (m *MockCloudwatchlogsClient) GetLogRecord(arg0 context.Context, arg1 *cloudwatchlogs.GetLogRecordInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLogRecord, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) GetLogRecord(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLogRecord, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetLogRecord), varargs...)
}

func (m *MockCloudwatchlogsClient) GetQueryResults(arg0 context.Context, arg1 *cloudwatchlogs.GetQueryResultsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetQueryResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueryResults, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) GetQueryResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueryResults, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetQueryResults), varargs...)
}

func (m *MockCloudwatchlogsClient) ListTagsForResource(arg0 context.Context, arg1 *cloudwatchlogs.ListTagsForResourceInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockCloudwatchlogsClient) ListTagsLogGroup(arg0 context.Context, arg1 *cloudwatchlogs.ListTagsLogGroupInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsLogGroup, varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.ListTagsLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchlogsClientMockRecorder) ListTagsLogGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsLogGroup, reflect.TypeOf((*MockCloudwatchlogsClient)(nil).ListTagsLogGroup), varargs...)
}
