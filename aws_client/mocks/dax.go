package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	dax "github.com/aws/aws-sdk-go-v2/service/dax"
	gomock "github.com/golang/mock/gomock"
)

type MockDaxClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDaxClientMockRecorder
}

type MockDaxClientMockRecorder struct {
	mock *MockDaxClient
}

func NewMockDaxClient(ctrl *gomock.Controller) *MockDaxClient {
	mock := &MockDaxClient{ctrl: ctrl}
	mock.recorder = &MockDaxClientMockRecorder{mock}
	return mock
}

func (m *MockDaxClient) EXPECT() *MockDaxClientMockRecorder {
	return m.recorder
}

func (m *MockDaxClient) DescribeClusters(arg0 context.Context, arg1 *dax.DescribeClustersInput, arg2 ...func(*dax.Options)) (*dax.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusters, varargs...)
	ret0, _ := ret[0].(*dax.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusters, reflect.TypeOf((*MockDaxClient)(nil).DescribeClusters), varargs...)
}

func (m *MockDaxClient) DescribeDefaultParameters(arg0 context.Context, arg1 *dax.DescribeDefaultParametersInput, arg2 ...func(*dax.Options)) (*dax.DescribeDefaultParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDefaultParameters, varargs...)
	ret0, _ := ret[0].(*dax.DescribeDefaultParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) DescribeDefaultParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDefaultParameters, reflect.TypeOf((*MockDaxClient)(nil).DescribeDefaultParameters), varargs...)
}

func (m *MockDaxClient) DescribeEvents(arg0 context.Context, arg1 *dax.DescribeEventsInput, arg2 ...func(*dax.Options)) (*dax.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*dax.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockDaxClient)(nil).DescribeEvents), varargs...)
}

func (m *MockDaxClient) DescribeParameterGroups(arg0 context.Context, arg1 *dax.DescribeParameterGroupsInput, arg2 ...func(*dax.Options)) (*dax.DescribeParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeParameterGroups, varargs...)
	ret0, _ := ret[0].(*dax.DescribeParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) DescribeParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeParameterGroups, reflect.TypeOf((*MockDaxClient)(nil).DescribeParameterGroups), varargs...)
}

func (m *MockDaxClient) DescribeParameters(arg0 context.Context, arg1 *dax.DescribeParametersInput, arg2 ...func(*dax.Options)) (*dax.DescribeParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeParameters, varargs...)
	ret0, _ := ret[0].(*dax.DescribeParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) DescribeParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeParameters, reflect.TypeOf((*MockDaxClient)(nil).DescribeParameters), varargs...)
}

func (m *MockDaxClient) DescribeSubnetGroups(arg0 context.Context, arg1 *dax.DescribeSubnetGroupsInput, arg2 ...func(*dax.Options)) (*dax.DescribeSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSubnetGroups, varargs...)
	ret0, _ := ret[0].(*dax.DescribeSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) DescribeSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSubnetGroups, reflect.TypeOf((*MockDaxClient)(nil).DescribeSubnetGroups), varargs...)
}

func (m *MockDaxClient) ListTags(arg0 context.Context, arg1 *dax.ListTagsInput, arg2 ...func(*dax.Options)) (*dax.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*dax.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaxClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockDaxClient)(nil).ListTags), varargs...)
}
