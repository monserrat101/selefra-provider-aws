package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cloudwatch "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudwatchClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudwatchClientMockRecorder
}

type MockCloudwatchClientMockRecorder struct {
	mock *MockCloudwatchClient
}

func NewMockCloudwatchClient(ctrl *gomock.Controller) *MockCloudwatchClient {
	mock := &MockCloudwatchClient{ctrl: ctrl}
	mock.recorder = &MockCloudwatchClientMockRecorder{mock}
	return mock
}

func (m *MockCloudwatchClient) EXPECT() *MockCloudwatchClientMockRecorder {
	return m.recorder
}

func (m *MockCloudwatchClient) DescribeAlarmHistory(arg0 context.Context, arg1 *cloudwatch.DescribeAlarmHistoryInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAlarmHistory, varargs...)
	ret0, _ := ret[0].(*cloudwatch.DescribeAlarmHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) DescribeAlarmHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAlarmHistory, reflect.TypeOf((*MockCloudwatchClient)(nil).DescribeAlarmHistory), varargs...)
}

func (m *MockCloudwatchClient) DescribeAlarms(arg0 context.Context, arg1 *cloudwatch.DescribeAlarmsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAlarms, varargs...)
	ret0, _ := ret[0].(*cloudwatch.DescribeAlarmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) DescribeAlarms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAlarms, reflect.TypeOf((*MockCloudwatchClient)(nil).DescribeAlarms), varargs...)
}

func (m *MockCloudwatchClient) DescribeAlarmsForMetric(arg0 context.Context, arg1 *cloudwatch.DescribeAlarmsForMetricInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAlarmsForMetric, varargs...)
	ret0, _ := ret[0].(*cloudwatch.DescribeAlarmsForMetricOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) DescribeAlarmsForMetric(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAlarmsForMetric, reflect.TypeOf((*MockCloudwatchClient)(nil).DescribeAlarmsForMetric), varargs...)
}

func (m *MockCloudwatchClient) DescribeAnomalyDetectors(arg0 context.Context, arg1 *cloudwatch.DescribeAnomalyDetectorsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAnomalyDetectors, varargs...)
	ret0, _ := ret[0].(*cloudwatch.DescribeAnomalyDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) DescribeAnomalyDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAnomalyDetectors, reflect.TypeOf((*MockCloudwatchClient)(nil).DescribeAnomalyDetectors), varargs...)
}

func (m *MockCloudwatchClient) DescribeInsightRules(arg0 context.Context, arg1 *cloudwatch.DescribeInsightRulesInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.DescribeInsightRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInsightRules, varargs...)
	ret0, _ := ret[0].(*cloudwatch.DescribeInsightRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) DescribeInsightRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInsightRules, reflect.TypeOf((*MockCloudwatchClient)(nil).DescribeInsightRules), varargs...)
}

func (m *MockCloudwatchClient) GetDashboard(arg0 context.Context, arg1 *cloudwatch.GetDashboardInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.GetDashboardOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDashboard, varargs...)
	ret0, _ := ret[0].(*cloudwatch.GetDashboardOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) GetDashboard(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDashboard, reflect.TypeOf((*MockCloudwatchClient)(nil).GetDashboard), varargs...)
}

func (m *MockCloudwatchClient) GetInsightRuleReport(arg0 context.Context, arg1 *cloudwatch.GetInsightRuleReportInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.GetInsightRuleReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsightRuleReport, varargs...)
	ret0, _ := ret[0].(*cloudwatch.GetInsightRuleReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) GetInsightRuleReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsightRuleReport, reflect.TypeOf((*MockCloudwatchClient)(nil).GetInsightRuleReport), varargs...)
}

func (m *MockCloudwatchClient) GetMetricData(arg0 context.Context, arg1 *cloudwatch.GetMetricDataInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMetricData, varargs...)
	ret0, _ := ret[0].(*cloudwatch.GetMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) GetMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMetricData, reflect.TypeOf((*MockCloudwatchClient)(nil).GetMetricData), varargs...)
}

func (m *MockCloudwatchClient) GetMetricStatistics(arg0 context.Context, arg1 *cloudwatch.GetMetricStatisticsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMetricStatistics, varargs...)
	ret0, _ := ret[0].(*cloudwatch.GetMetricStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) GetMetricStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMetricStatistics, reflect.TypeOf((*MockCloudwatchClient)(nil).GetMetricStatistics), varargs...)
}

func (m *MockCloudwatchClient) GetMetricStream(arg0 context.Context, arg1 *cloudwatch.GetMetricStreamInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMetricStream, varargs...)
	ret0, _ := ret[0].(*cloudwatch.GetMetricStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) GetMetricStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMetricStream, reflect.TypeOf((*MockCloudwatchClient)(nil).GetMetricStream), varargs...)
}

func (m *MockCloudwatchClient) GetMetricWidgetImage(arg0 context.Context, arg1 *cloudwatch.GetMetricWidgetImageInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMetricWidgetImage, varargs...)
	ret0, _ := ret[0].(*cloudwatch.GetMetricWidgetImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) GetMetricWidgetImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMetricWidgetImage, reflect.TypeOf((*MockCloudwatchClient)(nil).GetMetricWidgetImage), varargs...)
}

func (m *MockCloudwatchClient) ListDashboards(arg0 context.Context, arg1 *cloudwatch.ListDashboardsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.ListDashboardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDashboards, varargs...)
	ret0, _ := ret[0].(*cloudwatch.ListDashboardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) ListDashboards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDashboards, reflect.TypeOf((*MockCloudwatchClient)(nil).ListDashboards), varargs...)
}

func (m *MockCloudwatchClient) ListManagedInsightRules(arg0 context.Context, arg1 *cloudwatch.ListManagedInsightRulesInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.ListManagedInsightRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListManagedInsightRules, varargs...)
	ret0, _ := ret[0].(*cloudwatch.ListManagedInsightRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) ListManagedInsightRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListManagedInsightRules, reflect.TypeOf((*MockCloudwatchClient)(nil).ListManagedInsightRules), varargs...)
}

func (m *MockCloudwatchClient) ListMetricStreams(arg0 context.Context, arg1 *cloudwatch.ListMetricStreamsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMetricStreams, varargs...)
	ret0, _ := ret[0].(*cloudwatch.ListMetricStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) ListMetricStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMetricStreams, reflect.TypeOf((*MockCloudwatchClient)(nil).ListMetricStreams), varargs...)
}

func (m *MockCloudwatchClient) ListMetrics(arg0 context.Context, arg1 *cloudwatch.ListMetricsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMetrics, varargs...)
	ret0, _ := ret[0].(*cloudwatch.ListMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) ListMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMetrics, reflect.TypeOf((*MockCloudwatchClient)(nil).ListMetrics), varargs...)
}

func (m *MockCloudwatchClient) ListTagsForResource(arg0 context.Context, arg1 *cloudwatch.ListTagsForResourceInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*cloudwatch.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockCloudwatchClient)(nil).ListTagsForResource), varargs...)
}
