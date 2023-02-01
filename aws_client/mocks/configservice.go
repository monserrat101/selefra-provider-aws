package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	configservice "github.com/aws/aws-sdk-go-v2/service/configservice"
	gomock "github.com/golang/mock/gomock"
)

type MockConfigserviceClient struct {
	ctrl		*gomock.Controller
	recorder	*MockConfigserviceClientMockRecorder
}

type MockConfigserviceClientMockRecorder struct {
	mock *MockConfigserviceClient
}

func NewMockConfigserviceClient(ctrl *gomock.Controller) *MockConfigserviceClient {
	mock := &MockConfigserviceClient{ctrl: ctrl}
	mock.recorder = &MockConfigserviceClientMockRecorder{mock}
	return mock
}

func (m *MockConfigserviceClient) EXPECT() *MockConfigserviceClientMockRecorder {
	return m.recorder
}

func (m *MockConfigserviceClient) BatchGetAggregateResourceConfig(arg0 context.Context, arg1 *configservice.BatchGetAggregateResourceConfigInput, arg2 ...func(*configservice.Options)) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetAggregateResourceConfig, varargs...)
	ret0, _ := ret[0].(*configservice.BatchGetAggregateResourceConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) BatchGetAggregateResourceConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetAggregateResourceConfig, reflect.TypeOf((*MockConfigserviceClient)(nil).BatchGetAggregateResourceConfig), varargs...)
}

func (m *MockConfigserviceClient) BatchGetResourceConfig(arg0 context.Context, arg1 *configservice.BatchGetResourceConfigInput, arg2 ...func(*configservice.Options)) (*configservice.BatchGetResourceConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetResourceConfig, varargs...)
	ret0, _ := ret[0].(*configservice.BatchGetResourceConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) BatchGetResourceConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetResourceConfig, reflect.TypeOf((*MockConfigserviceClient)(nil).BatchGetResourceConfig), varargs...)
}

func (m *MockConfigserviceClient) DescribeAggregateComplianceByConfigRules(arg0 context.Context, arg1 *configservice.DescribeAggregateComplianceByConfigRulesInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAggregateComplianceByConfigRules, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeAggregateComplianceByConfigRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeAggregateComplianceByConfigRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAggregateComplianceByConfigRules, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeAggregateComplianceByConfigRules), varargs...)
}

func (m *MockConfigserviceClient) DescribeAggregateComplianceByConformancePacks(arg0 context.Context, arg1 *configservice.DescribeAggregateComplianceByConformancePacksInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeAggregateComplianceByConformancePacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAggregateComplianceByConformancePacks, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeAggregateComplianceByConformancePacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeAggregateComplianceByConformancePacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAggregateComplianceByConformancePacks, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeAggregateComplianceByConformancePacks), varargs...)
}

func (m *MockConfigserviceClient) DescribeAggregationAuthorizations(arg0 context.Context, arg1 *configservice.DescribeAggregationAuthorizationsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAggregationAuthorizations, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeAggregationAuthorizationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeAggregationAuthorizations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAggregationAuthorizations, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeAggregationAuthorizations), varargs...)
}

func (m *MockConfigserviceClient) DescribeComplianceByConfigRule(arg0 context.Context, arg1 *configservice.DescribeComplianceByConfigRuleInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeComplianceByConfigRule, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeComplianceByConfigRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeComplianceByConfigRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeComplianceByConfigRule, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeComplianceByConfigRule), varargs...)
}

func (m *MockConfigserviceClient) DescribeComplianceByResource(arg0 context.Context, arg1 *configservice.DescribeComplianceByResourceInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeComplianceByResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeComplianceByResource, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeComplianceByResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeComplianceByResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeComplianceByResource, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeComplianceByResource), varargs...)
}

func (m *MockConfigserviceClient) DescribeConfigRuleEvaluationStatus(arg0 context.Context, arg1 *configservice.DescribeConfigRuleEvaluationStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigRuleEvaluationStatus, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigRuleEvaluationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConfigRuleEvaluationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigRuleEvaluationStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConfigRuleEvaluationStatus), varargs...)
}

func (m *MockConfigserviceClient) DescribeConfigRules(arg0 context.Context, arg1 *configservice.DescribeConfigRulesInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigRules, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConfigRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigRules, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConfigRules), varargs...)
}

func (m *MockConfigserviceClient) DescribeConfigurationAggregatorSourcesStatus(arg0 context.Context, arg1 *configservice.DescribeConfigurationAggregatorSourcesStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationAggregatorSourcesStatus, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConfigurationAggregatorSourcesStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationAggregatorSourcesStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConfigurationAggregatorSourcesStatus), varargs...)
}

func (m *MockConfigserviceClient) DescribeConfigurationAggregators(arg0 context.Context, arg1 *configservice.DescribeConfigurationAggregatorsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationAggregators, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigurationAggregatorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConfigurationAggregators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationAggregators, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConfigurationAggregators), varargs...)
}

func (m *MockConfigserviceClient) DescribeConfigurationRecorderStatus(arg0 context.Context, arg1 *configservice.DescribeConfigurationRecorderStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationRecorderStatus, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigurationRecorderStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConfigurationRecorderStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationRecorderStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConfigurationRecorderStatus), varargs...)
}

func (m *MockConfigserviceClient) DescribeConfigurationRecorders(arg0 context.Context, arg1 *configservice.DescribeConfigurationRecordersInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecordersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationRecorders, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigurationRecordersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConfigurationRecorders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationRecorders, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConfigurationRecorders), varargs...)
}

func (m *MockConfigserviceClient) DescribeConformancePackCompliance(arg0 context.Context, arg1 *configservice.DescribeConformancePackComplianceInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConformancePackComplianceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConformancePackCompliance, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConformancePackComplianceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConformancePackCompliance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConformancePackCompliance, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConformancePackCompliance), varargs...)
}

func (m *MockConfigserviceClient) DescribeConformancePackStatus(arg0 context.Context, arg1 *configservice.DescribeConformancePackStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConformancePackStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConformancePackStatus, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConformancePackStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConformancePackStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConformancePackStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConformancePackStatus), varargs...)
}

func (m *MockConfigserviceClient) DescribeConformancePacks(arg0 context.Context, arg1 *configservice.DescribeConformancePacksInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConformancePacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConformancePacks, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConformancePacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeConformancePacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConformancePacks, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeConformancePacks), varargs...)
}

func (m *MockConfigserviceClient) DescribeDeliveryChannelStatus(arg0 context.Context, arg1 *configservice.DescribeDeliveryChannelStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDeliveryChannelStatus, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeDeliveryChannelStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeDeliveryChannelStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDeliveryChannelStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeDeliveryChannelStatus), varargs...)
}

func (m *MockConfigserviceClient) DescribeDeliveryChannels(arg0 context.Context, arg1 *configservice.DescribeDeliveryChannelsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeDeliveryChannelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDeliveryChannels, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeDeliveryChannelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeDeliveryChannels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDeliveryChannels, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeDeliveryChannels), varargs...)
}

func (m *MockConfigserviceClient) DescribeOrganizationConfigRuleStatuses(arg0 context.Context, arg1 *configservice.DescribeOrganizationConfigRuleStatusesInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConfigRuleStatuses, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeOrganizationConfigRuleStatusesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeOrganizationConfigRuleStatuses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConfigRuleStatuses, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeOrganizationConfigRuleStatuses), varargs...)
}

func (m *MockConfigserviceClient) DescribeOrganizationConfigRules(arg0 context.Context, arg1 *configservice.DescribeOrganizationConfigRulesInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConfigRules, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeOrganizationConfigRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeOrganizationConfigRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConfigRules, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeOrganizationConfigRules), varargs...)
}

func (m *MockConfigserviceClient) DescribeOrganizationConformancePackStatuses(arg0 context.Context, arg1 *configservice.DescribeOrganizationConformancePackStatusesInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConformancePackStatuses, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeOrganizationConformancePackStatusesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeOrganizationConformancePackStatuses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConformancePackStatuses, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeOrganizationConformancePackStatuses), varargs...)
}

func (m *MockConfigserviceClient) DescribeOrganizationConformancePacks(arg0 context.Context, arg1 *configservice.DescribeOrganizationConformancePacksInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrganizationConformancePacks, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeOrganizationConformancePacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeOrganizationConformancePacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrganizationConformancePacks, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeOrganizationConformancePacks), varargs...)
}

func (m *MockConfigserviceClient) DescribePendingAggregationRequests(arg0 context.Context, arg1 *configservice.DescribePendingAggregationRequestsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribePendingAggregationRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePendingAggregationRequests, varargs...)
	ret0, _ := ret[0].(*configservice.DescribePendingAggregationRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribePendingAggregationRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePendingAggregationRequests, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribePendingAggregationRequests), varargs...)
}

func (m *MockConfigserviceClient) DescribeRemediationConfigurations(arg0 context.Context, arg1 *configservice.DescribeRemediationConfigurationsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeRemediationConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRemediationConfigurations, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeRemediationConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeRemediationConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRemediationConfigurations, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeRemediationConfigurations), varargs...)
}

func (m *MockConfigserviceClient) DescribeRemediationExceptions(arg0 context.Context, arg1 *configservice.DescribeRemediationExceptionsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeRemediationExceptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRemediationExceptions, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeRemediationExceptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeRemediationExceptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRemediationExceptions, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeRemediationExceptions), varargs...)
}

func (m *MockConfigserviceClient) DescribeRemediationExecutionStatus(arg0 context.Context, arg1 *configservice.DescribeRemediationExecutionStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRemediationExecutionStatus, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeRemediationExecutionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeRemediationExecutionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRemediationExecutionStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeRemediationExecutionStatus), varargs...)
}

func (m *MockConfigserviceClient) DescribeRetentionConfigurations(arg0 context.Context, arg1 *configservice.DescribeRetentionConfigurationsInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeRetentionConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRetentionConfigurations, varargs...)
	ret0, _ := ret[0].(*configservice.DescribeRetentionConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) DescribeRetentionConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRetentionConfigurations, reflect.TypeOf((*MockConfigserviceClient)(nil).DescribeRetentionConfigurations), varargs...)
}

func (m *MockConfigserviceClient) GetAggregateComplianceDetailsByConfigRule(arg0 context.Context, arg1 *configservice.GetAggregateComplianceDetailsByConfigRuleInput, arg2 ...func(*configservice.Options)) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAggregateComplianceDetailsByConfigRule, varargs...)
	ret0, _ := ret[0].(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetAggregateComplianceDetailsByConfigRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAggregateComplianceDetailsByConfigRule, reflect.TypeOf((*MockConfigserviceClient)(nil).GetAggregateComplianceDetailsByConfigRule), varargs...)
}

func (m *MockConfigserviceClient) GetAggregateConfigRuleComplianceSummary(arg0 context.Context, arg1 *configservice.GetAggregateConfigRuleComplianceSummaryInput, arg2 ...func(*configservice.Options)) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAggregateConfigRuleComplianceSummary, varargs...)
	ret0, _ := ret[0].(*configservice.GetAggregateConfigRuleComplianceSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetAggregateConfigRuleComplianceSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAggregateConfigRuleComplianceSummary, reflect.TypeOf((*MockConfigserviceClient)(nil).GetAggregateConfigRuleComplianceSummary), varargs...)
}

func (m *MockConfigserviceClient) GetAggregateConformancePackComplianceSummary(arg0 context.Context, arg1 *configservice.GetAggregateConformancePackComplianceSummaryInput, arg2 ...func(*configservice.Options)) (*configservice.GetAggregateConformancePackComplianceSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAggregateConformancePackComplianceSummary, varargs...)
	ret0, _ := ret[0].(*configservice.GetAggregateConformancePackComplianceSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetAggregateConformancePackComplianceSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAggregateConformancePackComplianceSummary, reflect.TypeOf((*MockConfigserviceClient)(nil).GetAggregateConformancePackComplianceSummary), varargs...)
}

func (m *MockConfigserviceClient) GetAggregateDiscoveredResourceCounts(arg0 context.Context, arg1 *configservice.GetAggregateDiscoveredResourceCountsInput, arg2 ...func(*configservice.Options)) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAggregateDiscoveredResourceCounts, varargs...)
	ret0, _ := ret[0].(*configservice.GetAggregateDiscoveredResourceCountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetAggregateDiscoveredResourceCounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAggregateDiscoveredResourceCounts, reflect.TypeOf((*MockConfigserviceClient)(nil).GetAggregateDiscoveredResourceCounts), varargs...)
}

func (m *MockConfigserviceClient) GetAggregateResourceConfig(arg0 context.Context, arg1 *configservice.GetAggregateResourceConfigInput, arg2 ...func(*configservice.Options)) (*configservice.GetAggregateResourceConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAggregateResourceConfig, varargs...)
	ret0, _ := ret[0].(*configservice.GetAggregateResourceConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetAggregateResourceConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAggregateResourceConfig, reflect.TypeOf((*MockConfigserviceClient)(nil).GetAggregateResourceConfig), varargs...)
}

func (m *MockConfigserviceClient) GetComplianceDetailsByConfigRule(arg0 context.Context, arg1 *configservice.GetComplianceDetailsByConfigRuleInput, arg2 ...func(*configservice.Options)) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetComplianceDetailsByConfigRule, varargs...)
	ret0, _ := ret[0].(*configservice.GetComplianceDetailsByConfigRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetComplianceDetailsByConfigRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetComplianceDetailsByConfigRule, reflect.TypeOf((*MockConfigserviceClient)(nil).GetComplianceDetailsByConfigRule), varargs...)
}

func (m *MockConfigserviceClient) GetComplianceDetailsByResource(arg0 context.Context, arg1 *configservice.GetComplianceDetailsByResourceInput, arg2 ...func(*configservice.Options)) (*configservice.GetComplianceDetailsByResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetComplianceDetailsByResource, varargs...)
	ret0, _ := ret[0].(*configservice.GetComplianceDetailsByResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetComplianceDetailsByResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetComplianceDetailsByResource, reflect.TypeOf((*MockConfigserviceClient)(nil).GetComplianceDetailsByResource), varargs...)
}

func (m *MockConfigserviceClient) GetComplianceSummaryByConfigRule(arg0 context.Context, arg1 *configservice.GetComplianceSummaryByConfigRuleInput, arg2 ...func(*configservice.Options)) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetComplianceSummaryByConfigRule, varargs...)
	ret0, _ := ret[0].(*configservice.GetComplianceSummaryByConfigRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetComplianceSummaryByConfigRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetComplianceSummaryByConfigRule, reflect.TypeOf((*MockConfigserviceClient)(nil).GetComplianceSummaryByConfigRule), varargs...)
}

func (m *MockConfigserviceClient) GetComplianceSummaryByResourceType(arg0 context.Context, arg1 *configservice.GetComplianceSummaryByResourceTypeInput, arg2 ...func(*configservice.Options)) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetComplianceSummaryByResourceType, varargs...)
	ret0, _ := ret[0].(*configservice.GetComplianceSummaryByResourceTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetComplianceSummaryByResourceType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetComplianceSummaryByResourceType, reflect.TypeOf((*MockConfigserviceClient)(nil).GetComplianceSummaryByResourceType), varargs...)
}

func (m *MockConfigserviceClient) GetConformancePackComplianceDetails(arg0 context.Context, arg1 *configservice.GetConformancePackComplianceDetailsInput, arg2 ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConformancePackComplianceDetails, varargs...)
	ret0, _ := ret[0].(*configservice.GetConformancePackComplianceDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetConformancePackComplianceDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConformancePackComplianceDetails, reflect.TypeOf((*MockConfigserviceClient)(nil).GetConformancePackComplianceDetails), varargs...)
}

func (m *MockConfigserviceClient) GetConformancePackComplianceSummary(arg0 context.Context, arg1 *configservice.GetConformancePackComplianceSummaryInput, arg2 ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConformancePackComplianceSummary, varargs...)
	ret0, _ := ret[0].(*configservice.GetConformancePackComplianceSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetConformancePackComplianceSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConformancePackComplianceSummary, reflect.TypeOf((*MockConfigserviceClient)(nil).GetConformancePackComplianceSummary), varargs...)
}

func (m *MockConfigserviceClient) GetCustomRulePolicy(arg0 context.Context, arg1 *configservice.GetCustomRulePolicyInput, arg2 ...func(*configservice.Options)) (*configservice.GetCustomRulePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCustomRulePolicy, varargs...)
	ret0, _ := ret[0].(*configservice.GetCustomRulePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetCustomRulePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCustomRulePolicy, reflect.TypeOf((*MockConfigserviceClient)(nil).GetCustomRulePolicy), varargs...)
}

func (m *MockConfigserviceClient) GetDiscoveredResourceCounts(arg0 context.Context, arg1 *configservice.GetDiscoveredResourceCountsInput, arg2 ...func(*configservice.Options)) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDiscoveredResourceCounts, varargs...)
	ret0, _ := ret[0].(*configservice.GetDiscoveredResourceCountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetDiscoveredResourceCounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDiscoveredResourceCounts, reflect.TypeOf((*MockConfigserviceClient)(nil).GetDiscoveredResourceCounts), varargs...)
}

func (m *MockConfigserviceClient) GetOrganizationConfigRuleDetailedStatus(arg0 context.Context, arg1 *configservice.GetOrganizationConfigRuleDetailedStatusInput, arg2 ...func(*configservice.Options)) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOrganizationConfigRuleDetailedStatus, varargs...)
	ret0, _ := ret[0].(*configservice.GetOrganizationConfigRuleDetailedStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetOrganizationConfigRuleDetailedStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOrganizationConfigRuleDetailedStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).GetOrganizationConfigRuleDetailedStatus), varargs...)
}

func (m *MockConfigserviceClient) GetOrganizationConformancePackDetailedStatus(arg0 context.Context, arg1 *configservice.GetOrganizationConformancePackDetailedStatusInput, arg2 ...func(*configservice.Options)) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOrganizationConformancePackDetailedStatus, varargs...)
	ret0, _ := ret[0].(*configservice.GetOrganizationConformancePackDetailedStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetOrganizationConformancePackDetailedStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOrganizationConformancePackDetailedStatus, reflect.TypeOf((*MockConfigserviceClient)(nil).GetOrganizationConformancePackDetailedStatus), varargs...)
}

func (m *MockConfigserviceClient) GetOrganizationCustomRulePolicy(arg0 context.Context, arg1 *configservice.GetOrganizationCustomRulePolicyInput, arg2 ...func(*configservice.Options)) (*configservice.GetOrganizationCustomRulePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOrganizationCustomRulePolicy, varargs...)
	ret0, _ := ret[0].(*configservice.GetOrganizationCustomRulePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetOrganizationCustomRulePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOrganizationCustomRulePolicy, reflect.TypeOf((*MockConfigserviceClient)(nil).GetOrganizationCustomRulePolicy), varargs...)
}

func (m *MockConfigserviceClient) GetResourceConfigHistory(arg0 context.Context, arg1 *configservice.GetResourceConfigHistoryInput, arg2 ...func(*configservice.Options)) (*configservice.GetResourceConfigHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourceConfigHistory, varargs...)
	ret0, _ := ret[0].(*configservice.GetResourceConfigHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetResourceConfigHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourceConfigHistory, reflect.TypeOf((*MockConfigserviceClient)(nil).GetResourceConfigHistory), varargs...)
}

func (m *MockConfigserviceClient) GetResourceEvaluationSummary(arg0 context.Context, arg1 *configservice.GetResourceEvaluationSummaryInput, arg2 ...func(*configservice.Options)) (*configservice.GetResourceEvaluationSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourceEvaluationSummary, varargs...)
	ret0, _ := ret[0].(*configservice.GetResourceEvaluationSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetResourceEvaluationSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourceEvaluationSummary, reflect.TypeOf((*MockConfigserviceClient)(nil).GetResourceEvaluationSummary), varargs...)
}

func (m *MockConfigserviceClient) GetStoredQuery(arg0 context.Context, arg1 *configservice.GetStoredQueryInput, arg2 ...func(*configservice.Options)) (*configservice.GetStoredQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStoredQuery, varargs...)
	ret0, _ := ret[0].(*configservice.GetStoredQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) GetStoredQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStoredQuery, reflect.TypeOf((*MockConfigserviceClient)(nil).GetStoredQuery), varargs...)
}

func (m *MockConfigserviceClient) ListAggregateDiscoveredResources(arg0 context.Context, arg1 *configservice.ListAggregateDiscoveredResourcesInput, arg2 ...func(*configservice.Options)) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAggregateDiscoveredResources, varargs...)
	ret0, _ := ret[0].(*configservice.ListAggregateDiscoveredResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) ListAggregateDiscoveredResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAggregateDiscoveredResources, reflect.TypeOf((*MockConfigserviceClient)(nil).ListAggregateDiscoveredResources), varargs...)
}

func (m *MockConfigserviceClient) ListConformancePackComplianceScores(arg0 context.Context, arg1 *configservice.ListConformancePackComplianceScoresInput, arg2 ...func(*configservice.Options)) (*configservice.ListConformancePackComplianceScoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConformancePackComplianceScores, varargs...)
	ret0, _ := ret[0].(*configservice.ListConformancePackComplianceScoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) ListConformancePackComplianceScores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConformancePackComplianceScores, reflect.TypeOf((*MockConfigserviceClient)(nil).ListConformancePackComplianceScores), varargs...)
}

func (m *MockConfigserviceClient) ListDiscoveredResources(arg0 context.Context, arg1 *configservice.ListDiscoveredResourcesInput, arg2 ...func(*configservice.Options)) (*configservice.ListDiscoveredResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDiscoveredResources, varargs...)
	ret0, _ := ret[0].(*configservice.ListDiscoveredResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) ListDiscoveredResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDiscoveredResources, reflect.TypeOf((*MockConfigserviceClient)(nil).ListDiscoveredResources), varargs...)
}

func (m *MockConfigserviceClient) ListResourceEvaluations(arg0 context.Context, arg1 *configservice.ListResourceEvaluationsInput, arg2 ...func(*configservice.Options)) (*configservice.ListResourceEvaluationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceEvaluations, varargs...)
	ret0, _ := ret[0].(*configservice.ListResourceEvaluationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) ListResourceEvaluations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceEvaluations, reflect.TypeOf((*MockConfigserviceClient)(nil).ListResourceEvaluations), varargs...)
}

func (m *MockConfigserviceClient) ListStoredQueries(arg0 context.Context, arg1 *configservice.ListStoredQueriesInput, arg2 ...func(*configservice.Options)) (*configservice.ListStoredQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStoredQueries, varargs...)
	ret0, _ := ret[0].(*configservice.ListStoredQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) ListStoredQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStoredQueries, reflect.TypeOf((*MockConfigserviceClient)(nil).ListStoredQueries), varargs...)
}

func (m *MockConfigserviceClient) ListTagsForResource(arg0 context.Context, arg1 *configservice.ListTagsForResourceInput, arg2 ...func(*configservice.Options)) (*configservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*configservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigserviceClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockConfigserviceClient)(nil).ListTagsForResource), varargs...)
}
