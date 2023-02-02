package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	sfn "github.com/aws/aws-sdk-go-v2/service/sfn"
	gomock "github.com/golang/mock/gomock"
)

type MockSfnClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSfnClientMockRecorder
}

type MockSfnClientMockRecorder struct {
	mock *MockSfnClient
}

func NewMockSfnClient(ctrl *gomock.Controller) *MockSfnClient {
	mock := &MockSfnClient{ctrl: ctrl}
	mock.recorder = &MockSfnClientMockRecorder{mock}
	return mock
}

func (m *MockSfnClient) EXPECT() *MockSfnClientMockRecorder {
	return m.recorder
}

func (m *MockSfnClient) DescribeActivity(arg0 context.Context, arg1 *sfn.DescribeActivityInput, arg2 ...func(*sfn.Options)) (*sfn.DescribeActivityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeActivity, varargs...)
	ret0, _ := ret[0].(*sfn.DescribeActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) DescribeActivity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeActivity, reflect.TypeOf((*MockSfnClient)(nil).DescribeActivity), varargs...)
}

func (m *MockSfnClient) DescribeExecution(arg0 context.Context, arg1 *sfn.DescribeExecutionInput, arg2 ...func(*sfn.Options)) (*sfn.DescribeExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExecution, varargs...)
	ret0, _ := ret[0].(*sfn.DescribeExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) DescribeExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExecution, reflect.TypeOf((*MockSfnClient)(nil).DescribeExecution), varargs...)
}

func (m *MockSfnClient) DescribeMapRun(arg0 context.Context, arg1 *sfn.DescribeMapRunInput, arg2 ...func(*sfn.Options)) (*sfn.DescribeMapRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMapRun, varargs...)
	ret0, _ := ret[0].(*sfn.DescribeMapRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) DescribeMapRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMapRun, reflect.TypeOf((*MockSfnClient)(nil).DescribeMapRun), varargs...)
}

func (m *MockSfnClient) DescribeStateMachine(arg0 context.Context, arg1 *sfn.DescribeStateMachineInput, arg2 ...func(*sfn.Options)) (*sfn.DescribeStateMachineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStateMachine, varargs...)
	ret0, _ := ret[0].(*sfn.DescribeStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) DescribeStateMachine(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStateMachine, reflect.TypeOf((*MockSfnClient)(nil).DescribeStateMachine), varargs...)
}

func (m *MockSfnClient) DescribeStateMachineForExecution(arg0 context.Context, arg1 *sfn.DescribeStateMachineForExecutionInput, arg2 ...func(*sfn.Options)) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStateMachineForExecution, varargs...)
	ret0, _ := ret[0].(*sfn.DescribeStateMachineForExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) DescribeStateMachineForExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStateMachineForExecution, reflect.TypeOf((*MockSfnClient)(nil).DescribeStateMachineForExecution), varargs...)
}

func (m *MockSfnClient) GetActivityTask(arg0 context.Context, arg1 *sfn.GetActivityTaskInput, arg2 ...func(*sfn.Options)) (*sfn.GetActivityTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetActivityTask, varargs...)
	ret0, _ := ret[0].(*sfn.GetActivityTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) GetActivityTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetActivityTask, reflect.TypeOf((*MockSfnClient)(nil).GetActivityTask), varargs...)
}

func (m *MockSfnClient) GetExecutionHistory(arg0 context.Context, arg1 *sfn.GetExecutionHistoryInput, arg2 ...func(*sfn.Options)) (*sfn.GetExecutionHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetExecutionHistory, varargs...)
	ret0, _ := ret[0].(*sfn.GetExecutionHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) GetExecutionHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetExecutionHistory, reflect.TypeOf((*MockSfnClient)(nil).GetExecutionHistory), varargs...)
}

func (m *MockSfnClient) ListActivities(arg0 context.Context, arg1 *sfn.ListActivitiesInput, arg2 ...func(*sfn.Options)) (*sfn.ListActivitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActivities, varargs...)
	ret0, _ := ret[0].(*sfn.ListActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) ListActivities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActivities, reflect.TypeOf((*MockSfnClient)(nil).ListActivities), varargs...)
}

func (m *MockSfnClient) ListExecutions(arg0 context.Context, arg1 *sfn.ListExecutionsInput, arg2 ...func(*sfn.Options)) (*sfn.ListExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExecutions, varargs...)
	ret0, _ := ret[0].(*sfn.ListExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) ListExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExecutions, reflect.TypeOf((*MockSfnClient)(nil).ListExecutions), varargs...)
}

func (m *MockSfnClient) ListMapRuns(arg0 context.Context, arg1 *sfn.ListMapRunsInput, arg2 ...func(*sfn.Options)) (*sfn.ListMapRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMapRuns, varargs...)
	ret0, _ := ret[0].(*sfn.ListMapRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) ListMapRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMapRuns, reflect.TypeOf((*MockSfnClient)(nil).ListMapRuns), varargs...)
}

func (m *MockSfnClient) ListStateMachines(arg0 context.Context, arg1 *sfn.ListStateMachinesInput, arg2 ...func(*sfn.Options)) (*sfn.ListStateMachinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStateMachines, varargs...)
	ret0, _ := ret[0].(*sfn.ListStateMachinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) ListStateMachines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStateMachines, reflect.TypeOf((*MockSfnClient)(nil).ListStateMachines), varargs...)
}

func (m *MockSfnClient) ListTagsForResource(arg0 context.Context, arg1 *sfn.ListTagsForResourceInput, arg2 ...func(*sfn.Options)) (*sfn.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*sfn.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSfnClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSfnClient)(nil).ListTagsForResource), varargs...)
}
