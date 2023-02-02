package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	efs "github.com/aws/aws-sdk-go-v2/service/efs"
	gomock "github.com/golang/mock/gomock"
)

type MockEfsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEfsClientMockRecorder
}

type MockEfsClientMockRecorder struct {
	mock *MockEfsClient
}

func NewMockEfsClient(ctrl *gomock.Controller) *MockEfsClient {
	mock := &MockEfsClient{ctrl: ctrl}
	mock.recorder = &MockEfsClientMockRecorder{mock}
	return mock
}

func (m *MockEfsClient) EXPECT() *MockEfsClientMockRecorder {
	return m.recorder
}

func (m *MockEfsClient) DescribeAccessPoints(arg0 context.Context, arg1 *efs.DescribeAccessPointsInput, arg2 ...func(*efs.Options)) (*efs.DescribeAccessPointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccessPoints, varargs...)
	ret0, _ := ret[0].(*efs.DescribeAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeAccessPoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccessPoints, reflect.TypeOf((*MockEfsClient)(nil).DescribeAccessPoints), varargs...)
}

func (m *MockEfsClient) DescribeAccountPreferences(arg0 context.Context, arg1 *efs.DescribeAccountPreferencesInput, arg2 ...func(*efs.Options)) (*efs.DescribeAccountPreferencesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountPreferences, varargs...)
	ret0, _ := ret[0].(*efs.DescribeAccountPreferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeAccountPreferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountPreferences, reflect.TypeOf((*MockEfsClient)(nil).DescribeAccountPreferences), varargs...)
}

func (m *MockEfsClient) DescribeBackupPolicy(arg0 context.Context, arg1 *efs.DescribeBackupPolicyInput, arg2 ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBackupPolicy, varargs...)
	ret0, _ := ret[0].(*efs.DescribeBackupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeBackupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBackupPolicy, reflect.TypeOf((*MockEfsClient)(nil).DescribeBackupPolicy), varargs...)
}

func (m *MockEfsClient) DescribeFileSystemPolicy(arg0 context.Context, arg1 *efs.DescribeFileSystemPolicyInput, arg2 ...func(*efs.Options)) (*efs.DescribeFileSystemPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFileSystemPolicy, varargs...)
	ret0, _ := ret[0].(*efs.DescribeFileSystemPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeFileSystemPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFileSystemPolicy, reflect.TypeOf((*MockEfsClient)(nil).DescribeFileSystemPolicy), varargs...)
}

func (m *MockEfsClient) DescribeFileSystems(arg0 context.Context, arg1 *efs.DescribeFileSystemsInput, arg2 ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFileSystems, varargs...)
	ret0, _ := ret[0].(*efs.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeFileSystems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFileSystems, reflect.TypeOf((*MockEfsClient)(nil).DescribeFileSystems), varargs...)
}

func (m *MockEfsClient) DescribeLifecycleConfiguration(arg0 context.Context, arg1 *efs.DescribeLifecycleConfigurationInput, arg2 ...func(*efs.Options)) (*efs.DescribeLifecycleConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLifecycleConfiguration, varargs...)
	ret0, _ := ret[0].(*efs.DescribeLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLifecycleConfiguration, reflect.TypeOf((*MockEfsClient)(nil).DescribeLifecycleConfiguration), varargs...)
}

func (m *MockEfsClient) DescribeMountTargetSecurityGroups(arg0 context.Context, arg1 *efs.DescribeMountTargetSecurityGroupsInput, arg2 ...func(*efs.Options)) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMountTargetSecurityGroups, varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeMountTargetSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMountTargetSecurityGroups, reflect.TypeOf((*MockEfsClient)(nil).DescribeMountTargetSecurityGroups), varargs...)
}

func (m *MockEfsClient) DescribeMountTargets(arg0 context.Context, arg1 *efs.DescribeMountTargetsInput, arg2 ...func(*efs.Options)) (*efs.DescribeMountTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMountTargets, varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeMountTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMountTargets, reflect.TypeOf((*MockEfsClient)(nil).DescribeMountTargets), varargs...)
}

func (m *MockEfsClient) DescribeReplicationConfigurations(arg0 context.Context, arg1 *efs.DescribeReplicationConfigurationsInput, arg2 ...func(*efs.Options)) (*efs.DescribeReplicationConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationConfigurations, varargs...)
	ret0, _ := ret[0].(*efs.DescribeReplicationConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeReplicationConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationConfigurations, reflect.TypeOf((*MockEfsClient)(nil).DescribeReplicationConfigurations), varargs...)
}

func (m *MockEfsClient) DescribeTags(arg0 context.Context, arg1 *efs.DescribeTagsInput, arg2 ...func(*efs.Options)) (*efs.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*efs.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockEfsClient)(nil).DescribeTags), varargs...)
}

func (m *MockEfsClient) ListTagsForResource(arg0 context.Context, arg1 *efs.ListTagsForResourceInput, arg2 ...func(*efs.Options)) (*efs.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*efs.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockEfsClient)(nil).ListTagsForResource), varargs...)
}
