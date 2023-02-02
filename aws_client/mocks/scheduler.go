package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	scheduler "github.com/aws/aws-sdk-go-v2/service/scheduler"
	gomock "github.com/golang/mock/gomock"
)

type MockSchedulerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSchedulerClientMockRecorder
}

type MockSchedulerClientMockRecorder struct {
	mock *MockSchedulerClient
}

func NewMockSchedulerClient(ctrl *gomock.Controller) *MockSchedulerClient {
	mock := &MockSchedulerClient{ctrl: ctrl}
	mock.recorder = &MockSchedulerClientMockRecorder{mock}
	return mock
}

func (m *MockSchedulerClient) EXPECT() *MockSchedulerClientMockRecorder {
	return m.recorder
}

func (m *MockSchedulerClient) GetSchedule(arg0 context.Context, arg1 *scheduler.GetScheduleInput, arg2 ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSchedule, varargs...)
	ret0, _ := ret[0].(*scheduler.GetScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSchedulerClientMockRecorder) GetSchedule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSchedule, reflect.TypeOf((*MockSchedulerClient)(nil).GetSchedule), varargs...)
}

func (m *MockSchedulerClient) GetScheduleGroup(arg0 context.Context, arg1 *scheduler.GetScheduleGroupInput, arg2 ...func(*scheduler.Options)) (*scheduler.GetScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetScheduleGroup, varargs...)
	ret0, _ := ret[0].(*scheduler.GetScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSchedulerClientMockRecorder) GetScheduleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetScheduleGroup, reflect.TypeOf((*MockSchedulerClient)(nil).GetScheduleGroup), varargs...)
}

func (m *MockSchedulerClient) ListScheduleGroups(arg0 context.Context, arg1 *scheduler.ListScheduleGroupsInput, arg2 ...func(*scheduler.Options)) (*scheduler.ListScheduleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListScheduleGroups, varargs...)
	ret0, _ := ret[0].(*scheduler.ListScheduleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSchedulerClientMockRecorder) ListScheduleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListScheduleGroups, reflect.TypeOf((*MockSchedulerClient)(nil).ListScheduleGroups), varargs...)
}

func (m *MockSchedulerClient) ListSchedules(arg0 context.Context, arg1 *scheduler.ListSchedulesInput, arg2 ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSchedules, varargs...)
	ret0, _ := ret[0].(*scheduler.ListSchedulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSchedulerClientMockRecorder) ListSchedules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSchedules, reflect.TypeOf((*MockSchedulerClient)(nil).ListSchedules), varargs...)
}

func (m *MockSchedulerClient) ListTagsForResource(arg0 context.Context, arg1 *scheduler.ListTagsForResourceInput, arg2 ...func(*scheduler.Options)) (*scheduler.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*scheduler.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSchedulerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSchedulerClient)(nil).ListTagsForResource), varargs...)
}
