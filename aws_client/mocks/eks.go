package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	eks "github.com/aws/aws-sdk-go-v2/service/eks"
	gomock "github.com/golang/mock/gomock"
)

type MockEksClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEksClientMockRecorder
}

type MockEksClientMockRecorder struct {
	mock *MockEksClient
}

func NewMockEksClient(ctrl *gomock.Controller) *MockEksClient {
	mock := &MockEksClient{ctrl: ctrl}
	mock.recorder = &MockEksClientMockRecorder{mock}
	return mock
}

func (m *MockEksClient) EXPECT() *MockEksClientMockRecorder {
	return m.recorder
}

func (m *MockEksClient) DescribeAddon(arg0 context.Context, arg1 *eks.DescribeAddonInput, arg2 ...func(*eks.Options)) (*eks.DescribeAddonOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAddon, varargs...)
	ret0, _ := ret[0].(*eks.DescribeAddonOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeAddon(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAddon, reflect.TypeOf((*MockEksClient)(nil).DescribeAddon), varargs...)
}

func (m *MockEksClient) DescribeAddonConfiguration(arg0 context.Context, arg1 *eks.DescribeAddonConfigurationInput, arg2 ...func(*eks.Options)) (*eks.DescribeAddonConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAddonConfiguration, varargs...)
	ret0, _ := ret[0].(*eks.DescribeAddonConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeAddonConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAddonConfiguration, reflect.TypeOf((*MockEksClient)(nil).DescribeAddonConfiguration), varargs...)
}

func (m *MockEksClient) DescribeAddonVersions(arg0 context.Context, arg1 *eks.DescribeAddonVersionsInput, arg2 ...func(*eks.Options)) (*eks.DescribeAddonVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAddonVersions, varargs...)
	ret0, _ := ret[0].(*eks.DescribeAddonVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeAddonVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAddonVersions, reflect.TypeOf((*MockEksClient)(nil).DescribeAddonVersions), varargs...)
}

func (m *MockEksClient) DescribeCluster(arg0 context.Context, arg1 *eks.DescribeClusterInput, arg2 ...func(*eks.Options)) (*eks.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCluster, varargs...)
	ret0, _ := ret[0].(*eks.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCluster, reflect.TypeOf((*MockEksClient)(nil).DescribeCluster), varargs...)
}

func (m *MockEksClient) DescribeFargateProfile(arg0 context.Context, arg1 *eks.DescribeFargateProfileInput, arg2 ...func(*eks.Options)) (*eks.DescribeFargateProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFargateProfile, varargs...)
	ret0, _ := ret[0].(*eks.DescribeFargateProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeFargateProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFargateProfile, reflect.TypeOf((*MockEksClient)(nil).DescribeFargateProfile), varargs...)
}

func (m *MockEksClient) DescribeIdentityProviderConfig(arg0 context.Context, arg1 *eks.DescribeIdentityProviderConfigInput, arg2 ...func(*eks.Options)) (*eks.DescribeIdentityProviderConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIdentityProviderConfig, varargs...)
	ret0, _ := ret[0].(*eks.DescribeIdentityProviderConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeIdentityProviderConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIdentityProviderConfig, reflect.TypeOf((*MockEksClient)(nil).DescribeIdentityProviderConfig), varargs...)
}

func (m *MockEksClient) DescribeNodegroup(arg0 context.Context, arg1 *eks.DescribeNodegroupInput, arg2 ...func(*eks.Options)) (*eks.DescribeNodegroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNodegroup, varargs...)
	ret0, _ := ret[0].(*eks.DescribeNodegroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeNodegroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNodegroup, reflect.TypeOf((*MockEksClient)(nil).DescribeNodegroup), varargs...)
}

func (m *MockEksClient) DescribeUpdate(arg0 context.Context, arg1 *eks.DescribeUpdateInput, arg2 ...func(*eks.Options)) (*eks.DescribeUpdateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUpdate, varargs...)
	ret0, _ := ret[0].(*eks.DescribeUpdateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUpdate, reflect.TypeOf((*MockEksClient)(nil).DescribeUpdate), varargs...)
}

func (m *MockEksClient) ListAddons(arg0 context.Context, arg1 *eks.ListAddonsInput, arg2 ...func(*eks.Options)) (*eks.ListAddonsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAddons, varargs...)
	ret0, _ := ret[0].(*eks.ListAddonsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListAddons(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAddons, reflect.TypeOf((*MockEksClient)(nil).ListAddons), varargs...)
}

func (m *MockEksClient) ListClusters(arg0 context.Context, arg1 *eks.ListClustersInput, arg2 ...func(*eks.Options)) (*eks.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListClusters, varargs...)
	ret0, _ := ret[0].(*eks.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListClusters, reflect.TypeOf((*MockEksClient)(nil).ListClusters), varargs...)
}

func (m *MockEksClient) ListFargateProfiles(arg0 context.Context, arg1 *eks.ListFargateProfilesInput, arg2 ...func(*eks.Options)) (*eks.ListFargateProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFargateProfiles, varargs...)
	ret0, _ := ret[0].(*eks.ListFargateProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListFargateProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFargateProfiles, reflect.TypeOf((*MockEksClient)(nil).ListFargateProfiles), varargs...)
}

func (m *MockEksClient) ListIdentityProviderConfigs(arg0 context.Context, arg1 *eks.ListIdentityProviderConfigsInput, arg2 ...func(*eks.Options)) (*eks.ListIdentityProviderConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIdentityProviderConfigs, varargs...)
	ret0, _ := ret[0].(*eks.ListIdentityProviderConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListIdentityProviderConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIdentityProviderConfigs, reflect.TypeOf((*MockEksClient)(nil).ListIdentityProviderConfigs), varargs...)
}

func (m *MockEksClient) ListNodegroups(arg0 context.Context, arg1 *eks.ListNodegroupsInput, arg2 ...func(*eks.Options)) (*eks.ListNodegroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNodegroups, varargs...)
	ret0, _ := ret[0].(*eks.ListNodegroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListNodegroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNodegroups, reflect.TypeOf((*MockEksClient)(nil).ListNodegroups), varargs...)
}

func (m *MockEksClient) ListTagsForResource(arg0 context.Context, arg1 *eks.ListTagsForResourceInput, arg2 ...func(*eks.Options)) (*eks.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*eks.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockEksClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockEksClient) ListUpdates(arg0 context.Context, arg1 *eks.ListUpdatesInput, arg2 ...func(*eks.Options)) (*eks.ListUpdatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUpdates, varargs...)
	ret0, _ := ret[0].(*eks.ListUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUpdates, reflect.TypeOf((*MockEksClient)(nil).ListUpdates), varargs...)
}
