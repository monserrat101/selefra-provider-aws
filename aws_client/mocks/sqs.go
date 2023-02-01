package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	gomock "github.com/golang/mock/gomock"
)

type MockSqsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSqsClientMockRecorder
}

type MockSqsClientMockRecorder struct {
	mock *MockSqsClient
}

func NewMockSqsClient(ctrl *gomock.Controller) *MockSqsClient {
	mock := &MockSqsClient{ctrl: ctrl}
	mock.recorder = &MockSqsClientMockRecorder{mock}
	return mock
}

func (m *MockSqsClient) EXPECT() *MockSqsClientMockRecorder {
	return m.recorder
}

func (m *MockSqsClient) GetQueueAttributes(arg0 context.Context, arg1 *sqs.GetQueueAttributesInput, arg2 ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueueAttributes, varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSqsClientMockRecorder) GetQueueAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueueAttributes, reflect.TypeOf((*MockSqsClient)(nil).GetQueueAttributes), varargs...)
}

func (m *MockSqsClient) GetQueueUrl(arg0 context.Context, arg1 *sqs.GetQueueUrlInput, arg2 ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueueUrl, varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueUrlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSqsClientMockRecorder) GetQueueUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueueUrl, reflect.TypeOf((*MockSqsClient)(nil).GetQueueUrl), varargs...)
}

func (m *MockSqsClient) ListDeadLetterSourceQueues(arg0 context.Context, arg1 *sqs.ListDeadLetterSourceQueuesInput, arg2 ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDeadLetterSourceQueues, varargs...)
	ret0, _ := ret[0].(*sqs.ListDeadLetterSourceQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSqsClientMockRecorder) ListDeadLetterSourceQueues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDeadLetterSourceQueues, reflect.TypeOf((*MockSqsClient)(nil).ListDeadLetterSourceQueues), varargs...)
}

func (m *MockSqsClient) ListQueueTags(arg0 context.Context, arg1 *sqs.ListQueueTagsInput, arg2 ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListQueueTags, varargs...)
	ret0, _ := ret[0].(*sqs.ListQueueTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSqsClientMockRecorder) ListQueueTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListQueueTags, reflect.TypeOf((*MockSqsClient)(nil).ListQueueTags), varargs...)
}

func (m *MockSqsClient) ListQueues(arg0 context.Context, arg1 *sqs.ListQueuesInput, arg2 ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListQueues, varargs...)
	ret0, _ := ret[0].(*sqs.ListQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSqsClientMockRecorder) ListQueues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListQueues, reflect.TypeOf((*MockSqsClient)(nil).ListQueues), varargs...)
}
