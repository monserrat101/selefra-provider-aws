package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	kinesis "github.com/aws/aws-sdk-go-v2/service/kinesis"
	gomock "github.com/golang/mock/gomock"
)

type MockKinesisClient struct {
	ctrl		*gomock.Controller
	recorder	*MockKinesisClientMockRecorder
}

type MockKinesisClientMockRecorder struct {
	mock *MockKinesisClient
}

func NewMockKinesisClient(ctrl *gomock.Controller) *MockKinesisClient {
	mock := &MockKinesisClient{ctrl: ctrl}
	mock.recorder = &MockKinesisClientMockRecorder{mock}
	return mock
}

func (m *MockKinesisClient) EXPECT() *MockKinesisClientMockRecorder {
	return m.recorder
}

func (m *MockKinesisClient) DescribeLimits(arg0 context.Context, arg1 *kinesis.DescribeLimitsInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLimits, varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) DescribeLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLimits, reflect.TypeOf((*MockKinesisClient)(nil).DescribeLimits), varargs...)
}

func (m *MockKinesisClient) DescribeStream(arg0 context.Context, arg1 *kinesis.DescribeStreamInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStream, varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) DescribeStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStream, reflect.TypeOf((*MockKinesisClient)(nil).DescribeStream), varargs...)
}

func (m *MockKinesisClient) DescribeStreamConsumer(arg0 context.Context, arg1 *kinesis.DescribeStreamConsumerInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeStreamConsumerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStreamConsumer, varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeStreamConsumerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) DescribeStreamConsumer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStreamConsumer, reflect.TypeOf((*MockKinesisClient)(nil).DescribeStreamConsumer), varargs...)
}

func (m *MockKinesisClient) DescribeStreamSummary(arg0 context.Context, arg1 *kinesis.DescribeStreamSummaryInput, arg2 ...func(*kinesis.Options)) (*kinesis.DescribeStreamSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStreamSummary, varargs...)
	ret0, _ := ret[0].(*kinesis.DescribeStreamSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) DescribeStreamSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStreamSummary, reflect.TypeOf((*MockKinesisClient)(nil).DescribeStreamSummary), varargs...)
}

func (m *MockKinesisClient) GetRecords(arg0 context.Context, arg1 *kinesis.GetRecordsInput, arg2 ...func(*kinesis.Options)) (*kinesis.GetRecordsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRecords, varargs...)
	ret0, _ := ret[0].(*kinesis.GetRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) GetRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRecords, reflect.TypeOf((*MockKinesisClient)(nil).GetRecords), varargs...)
}

func (m *MockKinesisClient) GetShardIterator(arg0 context.Context, arg1 *kinesis.GetShardIteratorInput, arg2 ...func(*kinesis.Options)) (*kinesis.GetShardIteratorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetShardIterator, varargs...)
	ret0, _ := ret[0].(*kinesis.GetShardIteratorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) GetShardIterator(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetShardIterator, reflect.TypeOf((*MockKinesisClient)(nil).GetShardIterator), varargs...)
}

func (m *MockKinesisClient) ListShards(arg0 context.Context, arg1 *kinesis.ListShardsInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListShardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListShards, varargs...)
	ret0, _ := ret[0].(*kinesis.ListShardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) ListShards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListShards, reflect.TypeOf((*MockKinesisClient)(nil).ListShards), varargs...)
}

func (m *MockKinesisClient) ListStreamConsumers(arg0 context.Context, arg1 *kinesis.ListStreamConsumersInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStreamConsumers, varargs...)
	ret0, _ := ret[0].(*kinesis.ListStreamConsumersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) ListStreamConsumers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStreamConsumers, reflect.TypeOf((*MockKinesisClient)(nil).ListStreamConsumers), varargs...)
}

func (m *MockKinesisClient) ListStreams(arg0 context.Context, arg1 *kinesis.ListStreamsInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStreams, varargs...)
	ret0, _ := ret[0].(*kinesis.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) ListStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStreams, reflect.TypeOf((*MockKinesisClient)(nil).ListStreams), varargs...)
}

func (m *MockKinesisClient) ListTagsForStream(arg0 context.Context, arg1 *kinesis.ListTagsForStreamInput, arg2 ...func(*kinesis.Options)) (*kinesis.ListTagsForStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForStream, varargs...)
	ret0, _ := ret[0].(*kinesis.ListTagsForStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKinesisClientMockRecorder) ListTagsForStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForStream, reflect.TypeOf((*MockKinesisClient)(nil).ListTagsForStream), varargs...)
}
