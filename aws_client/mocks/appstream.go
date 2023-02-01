package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	appstream "github.com/aws/aws-sdk-go-v2/service/appstream"
	gomock "github.com/golang/mock/gomock"
)

type MockAppstreamClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAppstreamClientMockRecorder
}

type MockAppstreamClientMockRecorder struct {
	mock *MockAppstreamClient
}

func NewMockAppstreamClient(ctrl *gomock.Controller) *MockAppstreamClient {
	mock := &MockAppstreamClient{ctrl: ctrl}
	mock.recorder = &MockAppstreamClientMockRecorder{mock}
	return mock
}

func (m *MockAppstreamClient) EXPECT() *MockAppstreamClientMockRecorder {
	return m.recorder
}

func (m *MockAppstreamClient) DescribeAppBlocks(arg0 context.Context, arg1 *appstream.DescribeAppBlocksInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeAppBlocksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAppBlocks, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeAppBlocksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeAppBlocks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAppBlocks, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeAppBlocks), varargs...)
}

func (m *MockAppstreamClient) DescribeApplicationFleetAssociations(arg0 context.Context, arg1 *appstream.DescribeApplicationFleetAssociationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeApplicationFleetAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApplicationFleetAssociations, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeApplicationFleetAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeApplicationFleetAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApplicationFleetAssociations, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeApplicationFleetAssociations), varargs...)
}

func (m *MockAppstreamClient) DescribeApplications(arg0 context.Context, arg1 *appstream.DescribeApplicationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApplications, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApplications, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeApplications), varargs...)
}

func (m *MockAppstreamClient) DescribeDirectoryConfigs(arg0 context.Context, arg1 *appstream.DescribeDirectoryConfigsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeDirectoryConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDirectoryConfigs, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeDirectoryConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeDirectoryConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDirectoryConfigs, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeDirectoryConfigs), varargs...)
}

func (m *MockAppstreamClient) DescribeEntitlements(arg0 context.Context, arg1 *appstream.DescribeEntitlementsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeEntitlementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEntitlements, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeEntitlementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeEntitlements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEntitlements, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeEntitlements), varargs...)
}

func (m *MockAppstreamClient) DescribeFleets(arg0 context.Context, arg1 *appstream.DescribeFleetsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleets, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleets, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeFleets), varargs...)
}

func (m *MockAppstreamClient) DescribeImageBuilders(arg0 context.Context, arg1 *appstream.DescribeImageBuildersInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeImageBuildersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImageBuilders, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeImageBuildersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeImageBuilders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImageBuilders, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeImageBuilders), varargs...)
}

func (m *MockAppstreamClient) DescribeImagePermissions(arg0 context.Context, arg1 *appstream.DescribeImagePermissionsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeImagePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImagePermissions, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeImagePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeImagePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImagePermissions, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeImagePermissions), varargs...)
}

func (m *MockAppstreamClient) DescribeImages(arg0 context.Context, arg1 *appstream.DescribeImagesInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImages, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImages, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeImages), varargs...)
}

func (m *MockAppstreamClient) DescribeSessions(arg0 context.Context, arg1 *appstream.DescribeSessionsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSessions, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSessions, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeSessions), varargs...)
}

func (m *MockAppstreamClient) DescribeStacks(arg0 context.Context, arg1 *appstream.DescribeStacksInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStacks, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStacks, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeStacks), varargs...)
}

func (m *MockAppstreamClient) DescribeUsageReportSubscriptions(arg0 context.Context, arg1 *appstream.DescribeUsageReportSubscriptionsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUsageReportSubscriptions, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeUsageReportSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeUsageReportSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUsageReportSubscriptions, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeUsageReportSubscriptions), varargs...)
}

func (m *MockAppstreamClient) DescribeUserStackAssociations(arg0 context.Context, arg1 *appstream.DescribeUserStackAssociationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeUserStackAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserStackAssociations, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeUserStackAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeUserStackAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserStackAssociations, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeUserStackAssociations), varargs...)
}

func (m *MockAppstreamClient) DescribeUsers(arg0 context.Context, arg1 *appstream.DescribeUsersInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUsers, varargs...)
	ret0, _ := ret[0].(*appstream.DescribeUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) DescribeUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUsers, reflect.TypeOf((*MockAppstreamClient)(nil).DescribeUsers), varargs...)
}

func (m *MockAppstreamClient) ListAssociatedFleets(arg0 context.Context, arg1 *appstream.ListAssociatedFleetsInput, arg2 ...func(*appstream.Options)) (*appstream.ListAssociatedFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociatedFleets, varargs...)
	ret0, _ := ret[0].(*appstream.ListAssociatedFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) ListAssociatedFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociatedFleets, reflect.TypeOf((*MockAppstreamClient)(nil).ListAssociatedFleets), varargs...)
}

func (m *MockAppstreamClient) ListAssociatedStacks(arg0 context.Context, arg1 *appstream.ListAssociatedStacksInput, arg2 ...func(*appstream.Options)) (*appstream.ListAssociatedStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociatedStacks, varargs...)
	ret0, _ := ret[0].(*appstream.ListAssociatedStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) ListAssociatedStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociatedStacks, reflect.TypeOf((*MockAppstreamClient)(nil).ListAssociatedStacks), varargs...)
}

func (m *MockAppstreamClient) ListEntitledApplications(arg0 context.Context, arg1 *appstream.ListEntitledApplicationsInput, arg2 ...func(*appstream.Options)) (*appstream.ListEntitledApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEntitledApplications, varargs...)
	ret0, _ := ret[0].(*appstream.ListEntitledApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) ListEntitledApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEntitledApplications, reflect.TypeOf((*MockAppstreamClient)(nil).ListEntitledApplications), varargs...)
}

func (m *MockAppstreamClient) ListTagsForResource(arg0 context.Context, arg1 *appstream.ListTagsForResourceInput, arg2 ...func(*appstream.Options)) (*appstream.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*appstream.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppstreamClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockAppstreamClient)(nil).ListTagsForResource), varargs...)
}
