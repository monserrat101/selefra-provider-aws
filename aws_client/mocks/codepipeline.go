package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	codepipeline "github.com/aws/aws-sdk-go-v2/service/codepipeline"
	gomock "github.com/golang/mock/gomock"
)

type MockCodepipelineClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCodepipelineClientMockRecorder
}

type MockCodepipelineClientMockRecorder struct {
	mock *MockCodepipelineClient
}

func NewMockCodepipelineClient(ctrl *gomock.Controller) *MockCodepipelineClient {
	mock := &MockCodepipelineClient{ctrl: ctrl}
	mock.recorder = &MockCodepipelineClientMockRecorder{mock}
	return mock
}

func (m *MockCodepipelineClient) EXPECT() *MockCodepipelineClientMockRecorder {
	return m.recorder
}

func (m *MockCodepipelineClient) GetActionType(arg0 context.Context, arg1 *codepipeline.GetActionTypeInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetActionTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetActionType, varargs...)
	ret0, _ := ret[0].(*codepipeline.GetActionTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) GetActionType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetActionType, reflect.TypeOf((*MockCodepipelineClient)(nil).GetActionType), varargs...)
}

func (m *MockCodepipelineClient) GetJobDetails(arg0 context.Context, arg1 *codepipeline.GetJobDetailsInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetJobDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobDetails, varargs...)
	ret0, _ := ret[0].(*codepipeline.GetJobDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) GetJobDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobDetails, reflect.TypeOf((*MockCodepipelineClient)(nil).GetJobDetails), varargs...)
}

func (m *MockCodepipelineClient) GetPipeline(arg0 context.Context, arg1 *codepipeline.GetPipelineInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetPipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPipeline, varargs...)
	ret0, _ := ret[0].(*codepipeline.GetPipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) GetPipeline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPipeline, reflect.TypeOf((*MockCodepipelineClient)(nil).GetPipeline), varargs...)
}

func (m *MockCodepipelineClient) GetPipelineExecution(arg0 context.Context, arg1 *codepipeline.GetPipelineExecutionInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetPipelineExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPipelineExecution, varargs...)
	ret0, _ := ret[0].(*codepipeline.GetPipelineExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) GetPipelineExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPipelineExecution, reflect.TypeOf((*MockCodepipelineClient)(nil).GetPipelineExecution), varargs...)
}

func (m *MockCodepipelineClient) GetPipelineState(arg0 context.Context, arg1 *codepipeline.GetPipelineStateInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetPipelineStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPipelineState, varargs...)
	ret0, _ := ret[0].(*codepipeline.GetPipelineStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) GetPipelineState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPipelineState, reflect.TypeOf((*MockCodepipelineClient)(nil).GetPipelineState), varargs...)
}

func (m *MockCodepipelineClient) GetThirdPartyJobDetails(arg0 context.Context, arg1 *codepipeline.GetThirdPartyJobDetailsInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetThirdPartyJobDetails, varargs...)
	ret0, _ := ret[0].(*codepipeline.GetThirdPartyJobDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) GetThirdPartyJobDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetThirdPartyJobDetails, reflect.TypeOf((*MockCodepipelineClient)(nil).GetThirdPartyJobDetails), varargs...)
}

func (m *MockCodepipelineClient) ListActionExecutions(arg0 context.Context, arg1 *codepipeline.ListActionExecutionsInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListActionExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActionExecutions, varargs...)
	ret0, _ := ret[0].(*codepipeline.ListActionExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) ListActionExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActionExecutions, reflect.TypeOf((*MockCodepipelineClient)(nil).ListActionExecutions), varargs...)
}

func (m *MockCodepipelineClient) ListActionTypes(arg0 context.Context, arg1 *codepipeline.ListActionTypesInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListActionTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActionTypes, varargs...)
	ret0, _ := ret[0].(*codepipeline.ListActionTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) ListActionTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActionTypes, reflect.TypeOf((*MockCodepipelineClient)(nil).ListActionTypes), varargs...)
}

func (m *MockCodepipelineClient) ListPipelineExecutions(arg0 context.Context, arg1 *codepipeline.ListPipelineExecutionsInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListPipelineExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelineExecutions, varargs...)
	ret0, _ := ret[0].(*codepipeline.ListPipelineExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) ListPipelineExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelineExecutions, reflect.TypeOf((*MockCodepipelineClient)(nil).ListPipelineExecutions), varargs...)
}

func (m *MockCodepipelineClient) ListPipelines(arg0 context.Context, arg1 *codepipeline.ListPipelinesInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListPipelinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelines, varargs...)
	ret0, _ := ret[0].(*codepipeline.ListPipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) ListPipelines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelines, reflect.TypeOf((*MockCodepipelineClient)(nil).ListPipelines), varargs...)
}

func (m *MockCodepipelineClient) ListTagsForResource(arg0 context.Context, arg1 *codepipeline.ListTagsForResourceInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*codepipeline.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockCodepipelineClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockCodepipelineClient) ListWebhooks(arg0 context.Context, arg1 *codepipeline.ListWebhooksInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListWebhooksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWebhooks, varargs...)
	ret0, _ := ret[0].(*codepipeline.ListWebhooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodepipelineClientMockRecorder) ListWebhooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWebhooks, reflect.TypeOf((*MockCodepipelineClient)(nil).ListWebhooks), varargs...)
}
