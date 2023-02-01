package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	elastictranscoder "github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	gomock "github.com/golang/mock/gomock"
)

type MockElastictranscoderClient struct {
	ctrl		*gomock.Controller
	recorder	*MockElastictranscoderClientMockRecorder
}

type MockElastictranscoderClientMockRecorder struct {
	mock *MockElastictranscoderClient
}

func NewMockElastictranscoderClient(ctrl *gomock.Controller) *MockElastictranscoderClient {
	mock := &MockElastictranscoderClient{ctrl: ctrl}
	mock.recorder = &MockElastictranscoderClientMockRecorder{mock}
	return mock
}

func (m *MockElastictranscoderClient) EXPECT() *MockElastictranscoderClientMockRecorder {
	return m.recorder
}

func (m *MockElastictranscoderClient) ListJobsByPipeline(arg0 context.Context, arg1 *elastictranscoder.ListJobsByPipelineInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobsByPipeline, varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListJobsByPipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastictranscoderClientMockRecorder) ListJobsByPipeline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobsByPipeline, reflect.TypeOf((*MockElastictranscoderClient)(nil).ListJobsByPipeline), varargs...)
}

func (m *MockElastictranscoderClient) ListJobsByStatus(arg0 context.Context, arg1 *elastictranscoder.ListJobsByStatusInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobsByStatus, varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListJobsByStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastictranscoderClientMockRecorder) ListJobsByStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobsByStatus, reflect.TypeOf((*MockElastictranscoderClient)(nil).ListJobsByStatus), varargs...)
}

func (m *MockElastictranscoderClient) ListPipelines(arg0 context.Context, arg1 *elastictranscoder.ListPipelinesInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPipelinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelines, varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListPipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastictranscoderClientMockRecorder) ListPipelines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelines, reflect.TypeOf((*MockElastictranscoderClient)(nil).ListPipelines), varargs...)
}

func (m *MockElastictranscoderClient) ListPresets(arg0 context.Context, arg1 *elastictranscoder.ListPresetsInput, arg2 ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPresetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPresets, varargs...)
	ret0, _ := ret[0].(*elastictranscoder.ListPresetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastictranscoderClientMockRecorder) ListPresets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPresets, reflect.TypeOf((*MockElastictranscoderClient)(nil).ListPresets), varargs...)
}
