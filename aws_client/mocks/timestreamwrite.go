package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	timestreamwrite "github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	gomock "github.com/golang/mock/gomock"
)

type MockTimestreamwriteClient struct {
	ctrl		*gomock.Controller
	recorder	*MockTimestreamwriteClientMockRecorder
}

type MockTimestreamwriteClientMockRecorder struct {
	mock *MockTimestreamwriteClient
}

func NewMockTimestreamwriteClient(ctrl *gomock.Controller) *MockTimestreamwriteClient {
	mock := &MockTimestreamwriteClient{ctrl: ctrl}
	mock.recorder = &MockTimestreamwriteClientMockRecorder{mock}
	return mock
}

func (m *MockTimestreamwriteClient) EXPECT() *MockTimestreamwriteClientMockRecorder {
	return m.recorder
}

func (m *MockTimestreamwriteClient) DescribeDatabase(arg0 context.Context, arg1 *timestreamwrite.DescribeDatabaseInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeDatabaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDatabase, varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeDatabaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTimestreamwriteClientMockRecorder) DescribeDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDatabase, reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeDatabase), varargs...)
}

func (m *MockTimestreamwriteClient) DescribeEndpoints(arg0 context.Context, arg1 *timestreamwrite.DescribeEndpointsInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpoints, varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTimestreamwriteClientMockRecorder) DescribeEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpoints, reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeEndpoints), varargs...)
}

func (m *MockTimestreamwriteClient) DescribeTable(arg0 context.Context, arg1 *timestreamwrite.DescribeTableInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTable, varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTimestreamwriteClientMockRecorder) DescribeTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTable, reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeTable), varargs...)
}

func (m *MockTimestreamwriteClient) ListDatabases(arg0 context.Context, arg1 *timestreamwrite.ListDatabasesInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDatabases, varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTimestreamwriteClientMockRecorder) ListDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDatabases, reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListDatabases), varargs...)
}

func (m *MockTimestreamwriteClient) ListTables(arg0 context.Context, arg1 *timestreamwrite.ListTablesInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTables, varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTimestreamwriteClientMockRecorder) ListTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTables, reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListTables), varargs...)
}

func (m *MockTimestreamwriteClient) ListTagsForResource(arg0 context.Context, arg1 *timestreamwrite.ListTagsForResourceInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTimestreamwriteClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListTagsForResource), varargs...)
}
