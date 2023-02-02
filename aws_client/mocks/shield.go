package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	shield "github.com/aws/aws-sdk-go-v2/service/shield"
	gomock "github.com/golang/mock/gomock"
)

type MockShieldClient struct {
	ctrl		*gomock.Controller
	recorder	*MockShieldClientMockRecorder
}

type MockShieldClientMockRecorder struct {
	mock *MockShieldClient
}

func NewMockShieldClient(ctrl *gomock.Controller) *MockShieldClient {
	mock := &MockShieldClient{ctrl: ctrl}
	mock.recorder = &MockShieldClientMockRecorder{mock}
	return mock
}

func (m *MockShieldClient) EXPECT() *MockShieldClientMockRecorder {
	return m.recorder
}

func (m *MockShieldClient) DescribeAttack(arg0 context.Context, arg1 *shield.DescribeAttackInput, arg2 ...func(*shield.Options)) (*shield.DescribeAttackOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAttack, varargs...)
	ret0, _ := ret[0].(*shield.DescribeAttackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeAttack(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAttack, reflect.TypeOf((*MockShieldClient)(nil).DescribeAttack), varargs...)
}

func (m *MockShieldClient) DescribeAttackStatistics(arg0 context.Context, arg1 *shield.DescribeAttackStatisticsInput, arg2 ...func(*shield.Options)) (*shield.DescribeAttackStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAttackStatistics, varargs...)
	ret0, _ := ret[0].(*shield.DescribeAttackStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeAttackStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAttackStatistics, reflect.TypeOf((*MockShieldClient)(nil).DescribeAttackStatistics), varargs...)
}

func (m *MockShieldClient) DescribeDRTAccess(arg0 context.Context, arg1 *shield.DescribeDRTAccessInput, arg2 ...func(*shield.Options)) (*shield.DescribeDRTAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDRTAccess, varargs...)
	ret0, _ := ret[0].(*shield.DescribeDRTAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeDRTAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDRTAccess, reflect.TypeOf((*MockShieldClient)(nil).DescribeDRTAccess), varargs...)
}

func (m *MockShieldClient) DescribeEmergencyContactSettings(arg0 context.Context, arg1 *shield.DescribeEmergencyContactSettingsInput, arg2 ...func(*shield.Options)) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEmergencyContactSettings, varargs...)
	ret0, _ := ret[0].(*shield.DescribeEmergencyContactSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeEmergencyContactSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEmergencyContactSettings, reflect.TypeOf((*MockShieldClient)(nil).DescribeEmergencyContactSettings), varargs...)
}

func (m *MockShieldClient) DescribeProtection(arg0 context.Context, arg1 *shield.DescribeProtectionInput, arg2 ...func(*shield.Options)) (*shield.DescribeProtectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProtection, varargs...)
	ret0, _ := ret[0].(*shield.DescribeProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeProtection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProtection, reflect.TypeOf((*MockShieldClient)(nil).DescribeProtection), varargs...)
}

func (m *MockShieldClient) DescribeProtectionGroup(arg0 context.Context, arg1 *shield.DescribeProtectionGroupInput, arg2 ...func(*shield.Options)) (*shield.DescribeProtectionGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProtectionGroup, varargs...)
	ret0, _ := ret[0].(*shield.DescribeProtectionGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeProtectionGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProtectionGroup, reflect.TypeOf((*MockShieldClient)(nil).DescribeProtectionGroup), varargs...)
}

func (m *MockShieldClient) DescribeSubscription(arg0 context.Context, arg1 *shield.DescribeSubscriptionInput, arg2 ...func(*shield.Options)) (*shield.DescribeSubscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSubscription, varargs...)
	ret0, _ := ret[0].(*shield.DescribeSubscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) DescribeSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSubscription, reflect.TypeOf((*MockShieldClient)(nil).DescribeSubscription), varargs...)
}

func (m *MockShieldClient) GetSubscriptionState(arg0 context.Context, arg1 *shield.GetSubscriptionStateInput, arg2 ...func(*shield.Options)) (*shield.GetSubscriptionStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSubscriptionState, varargs...)
	ret0, _ := ret[0].(*shield.GetSubscriptionStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) GetSubscriptionState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSubscriptionState, reflect.TypeOf((*MockShieldClient)(nil).GetSubscriptionState), varargs...)
}

func (m *MockShieldClient) ListAttacks(arg0 context.Context, arg1 *shield.ListAttacksInput, arg2 ...func(*shield.Options)) (*shield.ListAttacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttacks, varargs...)
	ret0, _ := ret[0].(*shield.ListAttacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) ListAttacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttacks, reflect.TypeOf((*MockShieldClient)(nil).ListAttacks), varargs...)
}

func (m *MockShieldClient) ListProtectionGroups(arg0 context.Context, arg1 *shield.ListProtectionGroupsInput, arg2 ...func(*shield.Options)) (*shield.ListProtectionGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProtectionGroups, varargs...)
	ret0, _ := ret[0].(*shield.ListProtectionGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) ListProtectionGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProtectionGroups, reflect.TypeOf((*MockShieldClient)(nil).ListProtectionGroups), varargs...)
}

func (m *MockShieldClient) ListProtections(arg0 context.Context, arg1 *shield.ListProtectionsInput, arg2 ...func(*shield.Options)) (*shield.ListProtectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProtections, varargs...)
	ret0, _ := ret[0].(*shield.ListProtectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) ListProtections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProtections, reflect.TypeOf((*MockShieldClient)(nil).ListProtections), varargs...)
}

func (m *MockShieldClient) ListResourcesInProtectionGroup(arg0 context.Context, arg1 *shield.ListResourcesInProtectionGroupInput, arg2 ...func(*shield.Options)) (*shield.ListResourcesInProtectionGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourcesInProtectionGroup, varargs...)
	ret0, _ := ret[0].(*shield.ListResourcesInProtectionGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) ListResourcesInProtectionGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourcesInProtectionGroup, reflect.TypeOf((*MockShieldClient)(nil).ListResourcesInProtectionGroup), varargs...)
}

func (m *MockShieldClient) ListTagsForResource(arg0 context.Context, arg1 *shield.ListTagsForResourceInput, arg2 ...func(*shield.Options)) (*shield.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*shield.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockShieldClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockShieldClient)(nil).ListTagsForResource), varargs...)
}
