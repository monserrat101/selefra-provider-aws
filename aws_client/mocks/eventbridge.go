package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	eventbridge "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	gomock "github.com/golang/mock/gomock"
)

type MockEventbridgeClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEventbridgeClientMockRecorder
}

type MockEventbridgeClientMockRecorder struct {
	mock *MockEventbridgeClient
}

func NewMockEventbridgeClient(ctrl *gomock.Controller) *MockEventbridgeClient {
	mock := &MockEventbridgeClient{ctrl: ctrl}
	mock.recorder = &MockEventbridgeClientMockRecorder{mock}
	return mock
}

func (m *MockEventbridgeClient) EXPECT() *MockEventbridgeClientMockRecorder {
	return m.recorder
}

func (m *MockEventbridgeClient) DescribeApiDestination(arg0 context.Context, arg1 *eventbridge.DescribeApiDestinationInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeApiDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApiDestination, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeApiDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeApiDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApiDestination, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeApiDestination), varargs...)
}

func (m *MockEventbridgeClient) DescribeArchive(arg0 context.Context, arg1 *eventbridge.DescribeArchiveInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeArchiveOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeArchive, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeArchiveOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeArchive(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeArchive, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeArchive), varargs...)
}

func (m *MockEventbridgeClient) DescribeConnection(arg0 context.Context, arg1 *eventbridge.DescribeConnectionInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeConnectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnection, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeConnectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeConnection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnection, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeConnection), varargs...)
}

func (m *MockEventbridgeClient) DescribeEndpoint(arg0 context.Context, arg1 *eventbridge.DescribeEndpointInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeEndpointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpoint, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeEndpoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpoint, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeEndpoint), varargs...)
}

func (m *MockEventbridgeClient) DescribeEventBus(arg0 context.Context, arg1 *eventbridge.DescribeEventBusInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeEventBusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventBus, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeEventBusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeEventBus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventBus, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeEventBus), varargs...)
}

func (m *MockEventbridgeClient) DescribeEventSource(arg0 context.Context, arg1 *eventbridge.DescribeEventSourceInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeEventSourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventSource, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeEventSourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeEventSource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventSource, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeEventSource), varargs...)
}

func (m *MockEventbridgeClient) DescribePartnerEventSource(arg0 context.Context, arg1 *eventbridge.DescribePartnerEventSourceInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePartnerEventSource, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribePartnerEventSourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribePartnerEventSource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePartnerEventSource, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribePartnerEventSource), varargs...)
}

func (m *MockEventbridgeClient) DescribeReplay(arg0 context.Context, arg1 *eventbridge.DescribeReplayInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeReplayOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplay, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeReplayOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeReplay(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplay, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeReplay), varargs...)
}

func (m *MockEventbridgeClient) DescribeRule(arg0 context.Context, arg1 *eventbridge.DescribeRuleInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.DescribeRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRule, varargs...)
	ret0, _ := ret[0].(*eventbridge.DescribeRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) DescribeRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRule, reflect.TypeOf((*MockEventbridgeClient)(nil).DescribeRule), varargs...)
}

func (m *MockEventbridgeClient) ListApiDestinations(arg0 context.Context, arg1 *eventbridge.ListApiDestinationsInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListApiDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListApiDestinations, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListApiDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListApiDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListApiDestinations, reflect.TypeOf((*MockEventbridgeClient)(nil).ListApiDestinations), varargs...)
}

func (m *MockEventbridgeClient) ListArchives(arg0 context.Context, arg1 *eventbridge.ListArchivesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListArchivesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListArchives, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListArchivesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListArchives(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListArchives, reflect.TypeOf((*MockEventbridgeClient)(nil).ListArchives), varargs...)
}

func (m *MockEventbridgeClient) ListConnections(arg0 context.Context, arg1 *eventbridge.ListConnectionsInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConnections, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConnections, reflect.TypeOf((*MockEventbridgeClient)(nil).ListConnections), varargs...)
}

func (m *MockEventbridgeClient) ListEndpoints(arg0 context.Context, arg1 *eventbridge.ListEndpointsInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEndpoints, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEndpoints, reflect.TypeOf((*MockEventbridgeClient)(nil).ListEndpoints), varargs...)
}

func (m *MockEventbridgeClient) ListEventBuses(arg0 context.Context, arg1 *eventbridge.ListEventBusesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListEventBusesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEventBuses, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListEventBusesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListEventBuses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEventBuses, reflect.TypeOf((*MockEventbridgeClient)(nil).ListEventBuses), varargs...)
}

func (m *MockEventbridgeClient) ListEventSources(arg0 context.Context, arg1 *eventbridge.ListEventSourcesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListEventSourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEventSources, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListEventSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListEventSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEventSources, reflect.TypeOf((*MockEventbridgeClient)(nil).ListEventSources), varargs...)
}

func (m *MockEventbridgeClient) ListPartnerEventSourceAccounts(arg0 context.Context, arg1 *eventbridge.ListPartnerEventSourceAccountsInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPartnerEventSourceAccounts, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListPartnerEventSourceAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListPartnerEventSourceAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPartnerEventSourceAccounts, reflect.TypeOf((*MockEventbridgeClient)(nil).ListPartnerEventSourceAccounts), varargs...)
}

func (m *MockEventbridgeClient) ListPartnerEventSources(arg0 context.Context, arg1 *eventbridge.ListPartnerEventSourcesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPartnerEventSources, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListPartnerEventSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListPartnerEventSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPartnerEventSources, reflect.TypeOf((*MockEventbridgeClient)(nil).ListPartnerEventSources), varargs...)
}

func (m *MockEventbridgeClient) ListReplays(arg0 context.Context, arg1 *eventbridge.ListReplaysInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListReplaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReplays, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListReplaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListReplays(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReplays, reflect.TypeOf((*MockEventbridgeClient)(nil).ListReplays), varargs...)
}

func (m *MockEventbridgeClient) ListRuleNamesByTarget(arg0 context.Context, arg1 *eventbridge.ListRuleNamesByTargetInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRuleNamesByTarget, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListRuleNamesByTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListRuleNamesByTarget(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRuleNamesByTarget, reflect.TypeOf((*MockEventbridgeClient)(nil).ListRuleNamesByTarget), varargs...)
}

func (m *MockEventbridgeClient) ListRules(arg0 context.Context, arg1 *eventbridge.ListRulesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRules, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRules, reflect.TypeOf((*MockEventbridgeClient)(nil).ListRules), varargs...)
}

func (m *MockEventbridgeClient) ListTagsForResource(arg0 context.Context, arg1 *eventbridge.ListTagsForResourceInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockEventbridgeClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockEventbridgeClient) ListTargetsByRule(arg0 context.Context, arg1 *eventbridge.ListTargetsByRuleInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListTargetsByRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTargetsByRule, varargs...)
	ret0, _ := ret[0].(*eventbridge.ListTargetsByRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventbridgeClientMockRecorder) ListTargetsByRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTargetsByRule, reflect.TypeOf((*MockEventbridgeClient)(nil).ListTargetsByRule), varargs...)
}
