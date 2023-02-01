package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cloudformation "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudformationClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudformationClientMockRecorder
}

type MockCloudformationClientMockRecorder struct {
	mock *MockCloudformationClient
}

func NewMockCloudformationClient(ctrl *gomock.Controller) *MockCloudformationClient {
	mock := &MockCloudformationClient{ctrl: ctrl}
	mock.recorder = &MockCloudformationClientMockRecorder{mock}
	return mock
}

func (m *MockCloudformationClient) EXPECT() *MockCloudformationClientMockRecorder {
	return m.recorder
}

func (m *MockCloudformationClient) DescribeAccountLimits(arg0 context.Context, arg1 *cloudformation.DescribeAccountLimitsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeAccountLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountLimits, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeAccountLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeAccountLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountLimits, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeAccountLimits), varargs...)
}

func (m *MockCloudformationClient) DescribeChangeSet(arg0 context.Context, arg1 *cloudformation.DescribeChangeSetInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeChangeSet, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeChangeSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeChangeSet, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeChangeSet), varargs...)
}

func (m *MockCloudformationClient) DescribeChangeSetHooks(arg0 context.Context, arg1 *cloudformation.DescribeChangeSetHooksInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetHooksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeChangeSetHooks, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeChangeSetHooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeChangeSetHooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeChangeSetHooks, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeChangeSetHooks), varargs...)
}

func (m *MockCloudformationClient) DescribePublisher(arg0 context.Context, arg1 *cloudformation.DescribePublisherInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribePublisherOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePublisher, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribePublisherOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribePublisher(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePublisher, reflect.TypeOf((*MockCloudformationClient)(nil).DescribePublisher), varargs...)
}

func (m *MockCloudformationClient) DescribeStackDriftDetectionStatus(arg0 context.Context, arg1 *cloudformation.DescribeStackDriftDetectionStatusInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackDriftDetectionStatus, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackDriftDetectionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackDriftDetectionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackDriftDetectionStatus, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackDriftDetectionStatus), varargs...)
}

func (m *MockCloudformationClient) DescribeStackEvents(arg0 context.Context, arg1 *cloudformation.DescribeStackEventsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackEvents, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackEvents, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackEvents), varargs...)
}

func (m *MockCloudformationClient) DescribeStackInstance(arg0 context.Context, arg1 *cloudformation.DescribeStackInstanceInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackInstanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackInstance, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackInstance, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackInstance), varargs...)
}

func (m *MockCloudformationClient) DescribeStackResource(arg0 context.Context, arg1 *cloudformation.DescribeStackResourceInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackResource, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackResource, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackResource), varargs...)
}

func (m *MockCloudformationClient) DescribeStackResourceDrifts(arg0 context.Context, arg1 *cloudformation.DescribeStackResourceDriftsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackResourceDrifts, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackResourceDriftsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackResourceDrifts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackResourceDrifts, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackResourceDrifts), varargs...)
}

func (m *MockCloudformationClient) DescribeStackResources(arg0 context.Context, arg1 *cloudformation.DescribeStackResourcesInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackResources, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackResources, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackResources), varargs...)
}

func (m *MockCloudformationClient) DescribeStackSet(arg0 context.Context, arg1 *cloudformation.DescribeStackSetInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackSet, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackSet, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackSet), varargs...)
}

func (m *MockCloudformationClient) DescribeStackSetOperation(arg0 context.Context, arg1 *cloudformation.DescribeStackSetOperationInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOperationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStackSetOperation, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStackSetOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStackSetOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStackSetOperation, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStackSetOperation), varargs...)
}

func (m *MockCloudformationClient) DescribeStacks(arg0 context.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStacks, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStacks, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeStacks), varargs...)
}

func (m *MockCloudformationClient) DescribeType(arg0 context.Context, arg1 *cloudformation.DescribeTypeInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeType, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeType, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeType), varargs...)
}

func (m *MockCloudformationClient) DescribeTypeRegistration(arg0 context.Context, arg1 *cloudformation.DescribeTypeRegistrationInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeRegistrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTypeRegistration, varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeTypeRegistrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) DescribeTypeRegistration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTypeRegistration, reflect.TypeOf((*MockCloudformationClient)(nil).DescribeTypeRegistration), varargs...)
}

func (m *MockCloudformationClient) GetStackPolicy(arg0 context.Context, arg1 *cloudformation.GetStackPolicyInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.GetStackPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStackPolicy, varargs...)
	ret0, _ := ret[0].(*cloudformation.GetStackPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) GetStackPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStackPolicy, reflect.TypeOf((*MockCloudformationClient)(nil).GetStackPolicy), varargs...)
}

func (m *MockCloudformationClient) GetTemplate(arg0 context.Context, arg1 *cloudformation.GetTemplateInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTemplate, varargs...)
	ret0, _ := ret[0].(*cloudformation.GetTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) GetTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTemplate, reflect.TypeOf((*MockCloudformationClient)(nil).GetTemplate), varargs...)
}

func (m *MockCloudformationClient) GetTemplateSummary(arg0 context.Context, arg1 *cloudformation.GetTemplateSummaryInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.GetTemplateSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTemplateSummary, varargs...)
	ret0, _ := ret[0].(*cloudformation.GetTemplateSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) GetTemplateSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTemplateSummary, reflect.TypeOf((*MockCloudformationClient)(nil).GetTemplateSummary), varargs...)
}

func (m *MockCloudformationClient) ListChangeSets(arg0 context.Context, arg1 *cloudformation.ListChangeSetsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListChangeSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListChangeSets, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListChangeSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListChangeSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListChangeSets, reflect.TypeOf((*MockCloudformationClient)(nil).ListChangeSets), varargs...)
}

func (m *MockCloudformationClient) ListExports(arg0 context.Context, arg1 *cloudformation.ListExportsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListExportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExports, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListExportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListExports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExports, reflect.TypeOf((*MockCloudformationClient)(nil).ListExports), varargs...)
}

func (m *MockCloudformationClient) ListImports(arg0 context.Context, arg1 *cloudformation.ListImportsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListImportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImports, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListImportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListImports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImports, reflect.TypeOf((*MockCloudformationClient)(nil).ListImports), varargs...)
}

func (m *MockCloudformationClient) ListStackInstances(arg0 context.Context, arg1 *cloudformation.ListStackInstancesInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStackInstances, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListStackInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStackInstances, reflect.TypeOf((*MockCloudformationClient)(nil).ListStackInstances), varargs...)
}

func (m *MockCloudformationClient) ListStackResources(arg0 context.Context, arg1 *cloudformation.ListStackResourcesInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStackResources, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListStackResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStackResources, reflect.TypeOf((*MockCloudformationClient)(nil).ListStackResources), varargs...)
}

func (m *MockCloudformationClient) ListStackSetOperationResults(arg0 context.Context, arg1 *cloudformation.ListStackSetOperationResultsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStackSetOperationResults, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackSetOperationResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListStackSetOperationResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStackSetOperationResults, reflect.TypeOf((*MockCloudformationClient)(nil).ListStackSetOperationResults), varargs...)
}

func (m *MockCloudformationClient) ListStackSetOperations(arg0 context.Context, arg1 *cloudformation.ListStackSetOperationsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStackSetOperations, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackSetOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListStackSetOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStackSetOperations, reflect.TypeOf((*MockCloudformationClient)(nil).ListStackSetOperations), varargs...)
}

func (m *MockCloudformationClient) ListStackSets(arg0 context.Context, arg1 *cloudformation.ListStackSetsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStackSets, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListStackSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStackSets, reflect.TypeOf((*MockCloudformationClient)(nil).ListStackSets), varargs...)
}

func (m *MockCloudformationClient) ListStacks(arg0 context.Context, arg1 *cloudformation.ListStacksInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStacks, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStacks, reflect.TypeOf((*MockCloudformationClient)(nil).ListStacks), varargs...)
}

func (m *MockCloudformationClient) ListTypeRegistrations(arg0 context.Context, arg1 *cloudformation.ListTypeRegistrationsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListTypeRegistrationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTypeRegistrations, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListTypeRegistrationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListTypeRegistrations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTypeRegistrations, reflect.TypeOf((*MockCloudformationClient)(nil).ListTypeRegistrations), varargs...)
}

func (m *MockCloudformationClient) ListTypeVersions(arg0 context.Context, arg1 *cloudformation.ListTypeVersionsInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListTypeVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTypeVersions, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListTypeVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListTypeVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTypeVersions, reflect.TypeOf((*MockCloudformationClient)(nil).ListTypeVersions), varargs...)
}

func (m *MockCloudformationClient) ListTypes(arg0 context.Context, arg1 *cloudformation.ListTypesInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTypes, varargs...)
	ret0, _ := ret[0].(*cloudformation.ListTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudformationClientMockRecorder) ListTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTypes, reflect.TypeOf((*MockCloudformationClient)(nil).ListTypes), varargs...)
}
