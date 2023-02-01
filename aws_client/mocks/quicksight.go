package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	quicksight "github.com/aws/aws-sdk-go-v2/service/quicksight"
	gomock "github.com/golang/mock/gomock"
)

type MockQuicksightClient struct {
	ctrl		*gomock.Controller
	recorder	*MockQuicksightClientMockRecorder
}

type MockQuicksightClientMockRecorder struct {
	mock *MockQuicksightClient
}

func NewMockQuicksightClient(ctrl *gomock.Controller) *MockQuicksightClient {
	mock := &MockQuicksightClient{ctrl: ctrl}
	mock.recorder = &MockQuicksightClientMockRecorder{mock}
	return mock
}

func (m *MockQuicksightClient) EXPECT() *MockQuicksightClientMockRecorder {
	return m.recorder
}

func (m *MockQuicksightClient) DescribeAccountCustomization(arg0 context.Context, arg1 *quicksight.DescribeAccountCustomizationInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeAccountCustomizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountCustomization, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeAccountCustomizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeAccountCustomization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountCustomization, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeAccountCustomization), varargs...)
}

func (m *MockQuicksightClient) DescribeAccountSettings(arg0 context.Context, arg1 *quicksight.DescribeAccountSettingsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeAccountSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountSettings, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeAccountSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeAccountSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountSettings, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeAccountSettings), varargs...)
}

func (m *MockQuicksightClient) DescribeAccountSubscription(arg0 context.Context, arg1 *quicksight.DescribeAccountSubscriptionInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeAccountSubscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountSubscription, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeAccountSubscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeAccountSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountSubscription, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeAccountSubscription), varargs...)
}

func (m *MockQuicksightClient) DescribeAnalysis(arg0 context.Context, arg1 *quicksight.DescribeAnalysisInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAnalysis, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeAnalysisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeAnalysis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAnalysis, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeAnalysis), varargs...)
}

func (m *MockQuicksightClient) DescribeAnalysisDefinition(arg0 context.Context, arg1 *quicksight.DescribeAnalysisDefinitionInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAnalysisDefinition, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeAnalysisDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeAnalysisDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAnalysisDefinition, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeAnalysisDefinition), varargs...)
}

func (m *MockQuicksightClient) DescribeAnalysisPermissions(arg0 context.Context, arg1 *quicksight.DescribeAnalysisPermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAnalysisPermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeAnalysisPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeAnalysisPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAnalysisPermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeAnalysisPermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeDashboard(arg0 context.Context, arg1 *quicksight.DescribeDashboardInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDashboardOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDashboard, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDashboardOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDashboard(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDashboard, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDashboard), varargs...)
}

func (m *MockQuicksightClient) DescribeDashboardDefinition(arg0 context.Context, arg1 *quicksight.DescribeDashboardDefinitionInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDashboardDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDashboardDefinition, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDashboardDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDashboardDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDashboardDefinition, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDashboardDefinition), varargs...)
}

func (m *MockQuicksightClient) DescribeDashboardPermissions(arg0 context.Context, arg1 *quicksight.DescribeDashboardPermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDashboardPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDashboardPermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDashboardPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDashboardPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDashboardPermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDashboardPermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeDataSet(arg0 context.Context, arg1 *quicksight.DescribeDataSetInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDataSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataSet, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDataSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDataSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataSet, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDataSet), varargs...)
}

func (m *MockQuicksightClient) DescribeDataSetPermissions(arg0 context.Context, arg1 *quicksight.DescribeDataSetPermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDataSetPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataSetPermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDataSetPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDataSetPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataSetPermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDataSetPermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeDataSource(arg0 context.Context, arg1 *quicksight.DescribeDataSourceInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDataSourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataSource, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDataSourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDataSource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataSource, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDataSource), varargs...)
}

func (m *MockQuicksightClient) DescribeDataSourcePermissions(arg0 context.Context, arg1 *quicksight.DescribeDataSourcePermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataSourcePermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeDataSourcePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeDataSourcePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataSourcePermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeDataSourcePermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeFolder(arg0 context.Context, arg1 *quicksight.DescribeFolderInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeFolderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFolder, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeFolderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeFolder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFolder, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeFolder), varargs...)
}

func (m *MockQuicksightClient) DescribeFolderPermissions(arg0 context.Context, arg1 *quicksight.DescribeFolderPermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeFolderPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFolderPermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeFolderPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeFolderPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFolderPermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeFolderPermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeFolderResolvedPermissions(arg0 context.Context, arg1 *quicksight.DescribeFolderResolvedPermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeFolderResolvedPermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFolderResolvedPermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeFolderResolvedPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeFolderResolvedPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFolderResolvedPermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeFolderResolvedPermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeGroup(arg0 context.Context, arg1 *quicksight.DescribeGroupInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGroup, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGroup, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeGroup), varargs...)
}

func (m *MockQuicksightClient) DescribeGroupMembership(arg0 context.Context, arg1 *quicksight.DescribeGroupMembershipInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeGroupMembershipOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGroupMembership, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeGroupMembershipOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeGroupMembership(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGroupMembership, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeGroupMembership), varargs...)
}

func (m *MockQuicksightClient) DescribeIAMPolicyAssignment(arg0 context.Context, arg1 *quicksight.DescribeIAMPolicyAssignmentInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIAMPolicyAssignment, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeIAMPolicyAssignmentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeIAMPolicyAssignment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIAMPolicyAssignment, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeIAMPolicyAssignment), varargs...)
}

func (m *MockQuicksightClient) DescribeIngestion(arg0 context.Context, arg1 *quicksight.DescribeIngestionInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeIngestionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIngestion, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeIngestionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeIngestion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIngestion, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeIngestion), varargs...)
}

func (m *MockQuicksightClient) DescribeIpRestriction(arg0 context.Context, arg1 *quicksight.DescribeIpRestrictionInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeIpRestrictionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIpRestriction, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeIpRestrictionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeIpRestriction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIpRestriction, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeIpRestriction), varargs...)
}

func (m *MockQuicksightClient) DescribeNamespace(arg0 context.Context, arg1 *quicksight.DescribeNamespaceInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeNamespaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNamespace, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeNamespaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeNamespace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNamespace, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeNamespace), varargs...)
}

func (m *MockQuicksightClient) DescribeTemplate(arg0 context.Context, arg1 *quicksight.DescribeTemplateInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTemplate, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTemplate, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeTemplate), varargs...)
}

func (m *MockQuicksightClient) DescribeTemplateAlias(arg0 context.Context, arg1 *quicksight.DescribeTemplateAliasInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeTemplateAliasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTemplateAlias, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeTemplateAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeTemplateAlias(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTemplateAlias, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeTemplateAlias), varargs...)
}

func (m *MockQuicksightClient) DescribeTemplateDefinition(arg0 context.Context, arg1 *quicksight.DescribeTemplateDefinitionInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeTemplateDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTemplateDefinition, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeTemplateDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeTemplateDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTemplateDefinition, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeTemplateDefinition), varargs...)
}

func (m *MockQuicksightClient) DescribeTemplatePermissions(arg0 context.Context, arg1 *quicksight.DescribeTemplatePermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeTemplatePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTemplatePermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeTemplatePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeTemplatePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTemplatePermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeTemplatePermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeTheme(arg0 context.Context, arg1 *quicksight.DescribeThemeInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeThemeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTheme, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeThemeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeTheme(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTheme, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeTheme), varargs...)
}

func (m *MockQuicksightClient) DescribeThemeAlias(arg0 context.Context, arg1 *quicksight.DescribeThemeAliasInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeThemeAliasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeThemeAlias, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeThemeAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeThemeAlias(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeThemeAlias, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeThemeAlias), varargs...)
}

func (m *MockQuicksightClient) DescribeThemePermissions(arg0 context.Context, arg1 *quicksight.DescribeThemePermissionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeThemePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeThemePermissions, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeThemePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeThemePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeThemePermissions, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeThemePermissions), varargs...)
}

func (m *MockQuicksightClient) DescribeUser(arg0 context.Context, arg1 *quicksight.DescribeUserInput, arg2 ...func(*quicksight.Options)) (*quicksight.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUser, varargs...)
	ret0, _ := ret[0].(*quicksight.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) DescribeUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUser, reflect.TypeOf((*MockQuicksightClient)(nil).DescribeUser), varargs...)
}

func (m *MockQuicksightClient) GetDashboardEmbedUrl(arg0 context.Context, arg1 *quicksight.GetDashboardEmbedUrlInput, arg2 ...func(*quicksight.Options)) (*quicksight.GetDashboardEmbedUrlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDashboardEmbedUrl, varargs...)
	ret0, _ := ret[0].(*quicksight.GetDashboardEmbedUrlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) GetDashboardEmbedUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDashboardEmbedUrl, reflect.TypeOf((*MockQuicksightClient)(nil).GetDashboardEmbedUrl), varargs...)
}

func (m *MockQuicksightClient) GetSessionEmbedUrl(arg0 context.Context, arg1 *quicksight.GetSessionEmbedUrlInput, arg2 ...func(*quicksight.Options)) (*quicksight.GetSessionEmbedUrlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSessionEmbedUrl, varargs...)
	ret0, _ := ret[0].(*quicksight.GetSessionEmbedUrlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) GetSessionEmbedUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSessionEmbedUrl, reflect.TypeOf((*MockQuicksightClient)(nil).GetSessionEmbedUrl), varargs...)
}

func (m *MockQuicksightClient) ListAnalyses(arg0 context.Context, arg1 *quicksight.ListAnalysesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListAnalysesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAnalyses, varargs...)
	ret0, _ := ret[0].(*quicksight.ListAnalysesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListAnalyses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAnalyses, reflect.TypeOf((*MockQuicksightClient)(nil).ListAnalyses), varargs...)
}

func (m *MockQuicksightClient) ListDashboardVersions(arg0 context.Context, arg1 *quicksight.ListDashboardVersionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListDashboardVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDashboardVersions, varargs...)
	ret0, _ := ret[0].(*quicksight.ListDashboardVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListDashboardVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDashboardVersions, reflect.TypeOf((*MockQuicksightClient)(nil).ListDashboardVersions), varargs...)
}

func (m *MockQuicksightClient) ListDashboards(arg0 context.Context, arg1 *quicksight.ListDashboardsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListDashboardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDashboards, varargs...)
	ret0, _ := ret[0].(*quicksight.ListDashboardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListDashboards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDashboards, reflect.TypeOf((*MockQuicksightClient)(nil).ListDashboards), varargs...)
}

func (m *MockQuicksightClient) ListDataSets(arg0 context.Context, arg1 *quicksight.ListDataSetsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListDataSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataSets, varargs...)
	ret0, _ := ret[0].(*quicksight.ListDataSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListDataSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataSets, reflect.TypeOf((*MockQuicksightClient)(nil).ListDataSets), varargs...)
}

func (m *MockQuicksightClient) ListDataSources(arg0 context.Context, arg1 *quicksight.ListDataSourcesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListDataSourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataSources, varargs...)
	ret0, _ := ret[0].(*quicksight.ListDataSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListDataSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataSources, reflect.TypeOf((*MockQuicksightClient)(nil).ListDataSources), varargs...)
}

func (m *MockQuicksightClient) ListFolderMembers(arg0 context.Context, arg1 *quicksight.ListFolderMembersInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListFolderMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFolderMembers, varargs...)
	ret0, _ := ret[0].(*quicksight.ListFolderMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListFolderMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFolderMembers, reflect.TypeOf((*MockQuicksightClient)(nil).ListFolderMembers), varargs...)
}

func (m *MockQuicksightClient) ListFolders(arg0 context.Context, arg1 *quicksight.ListFoldersInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListFoldersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFolders, varargs...)
	ret0, _ := ret[0].(*quicksight.ListFoldersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListFolders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFolders, reflect.TypeOf((*MockQuicksightClient)(nil).ListFolders), varargs...)
}

func (m *MockQuicksightClient) ListGroupMemberships(arg0 context.Context, arg1 *quicksight.ListGroupMembershipsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListGroupMembershipsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroupMemberships, varargs...)
	ret0, _ := ret[0].(*quicksight.ListGroupMembershipsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListGroupMemberships(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroupMemberships, reflect.TypeOf((*MockQuicksightClient)(nil).ListGroupMemberships), varargs...)
}

func (m *MockQuicksightClient) ListGroups(arg0 context.Context, arg1 *quicksight.ListGroupsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroups, varargs...)
	ret0, _ := ret[0].(*quicksight.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroups, reflect.TypeOf((*MockQuicksightClient)(nil).ListGroups), varargs...)
}

func (m *MockQuicksightClient) ListIAMPolicyAssignments(arg0 context.Context, arg1 *quicksight.ListIAMPolicyAssignmentsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIAMPolicyAssignments, varargs...)
	ret0, _ := ret[0].(*quicksight.ListIAMPolicyAssignmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListIAMPolicyAssignments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIAMPolicyAssignments, reflect.TypeOf((*MockQuicksightClient)(nil).ListIAMPolicyAssignments), varargs...)
}

func (m *MockQuicksightClient) ListIAMPolicyAssignmentsForUser(arg0 context.Context, arg1 *quicksight.ListIAMPolicyAssignmentsForUserInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIAMPolicyAssignmentsForUser, varargs...)
	ret0, _ := ret[0].(*quicksight.ListIAMPolicyAssignmentsForUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListIAMPolicyAssignmentsForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIAMPolicyAssignmentsForUser, reflect.TypeOf((*MockQuicksightClient)(nil).ListIAMPolicyAssignmentsForUser), varargs...)
}

func (m *MockQuicksightClient) ListIngestions(arg0 context.Context, arg1 *quicksight.ListIngestionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListIngestionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIngestions, varargs...)
	ret0, _ := ret[0].(*quicksight.ListIngestionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListIngestions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIngestions, reflect.TypeOf((*MockQuicksightClient)(nil).ListIngestions), varargs...)
}

func (m *MockQuicksightClient) ListNamespaces(arg0 context.Context, arg1 *quicksight.ListNamespacesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListNamespacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNamespaces, varargs...)
	ret0, _ := ret[0].(*quicksight.ListNamespacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListNamespaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNamespaces, reflect.TypeOf((*MockQuicksightClient)(nil).ListNamespaces), varargs...)
}

func (m *MockQuicksightClient) ListTagsForResource(arg0 context.Context, arg1 *quicksight.ListTagsForResourceInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*quicksight.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockQuicksightClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockQuicksightClient) ListTemplateAliases(arg0 context.Context, arg1 *quicksight.ListTemplateAliasesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListTemplateAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTemplateAliases, varargs...)
	ret0, _ := ret[0].(*quicksight.ListTemplateAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListTemplateAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTemplateAliases, reflect.TypeOf((*MockQuicksightClient)(nil).ListTemplateAliases), varargs...)
}

func (m *MockQuicksightClient) ListTemplateVersions(arg0 context.Context, arg1 *quicksight.ListTemplateVersionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListTemplateVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTemplateVersions, varargs...)
	ret0, _ := ret[0].(*quicksight.ListTemplateVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListTemplateVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTemplateVersions, reflect.TypeOf((*MockQuicksightClient)(nil).ListTemplateVersions), varargs...)
}

func (m *MockQuicksightClient) ListTemplates(arg0 context.Context, arg1 *quicksight.ListTemplatesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTemplates, varargs...)
	ret0, _ := ret[0].(*quicksight.ListTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTemplates, reflect.TypeOf((*MockQuicksightClient)(nil).ListTemplates), varargs...)
}

func (m *MockQuicksightClient) ListThemeAliases(arg0 context.Context, arg1 *quicksight.ListThemeAliasesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListThemeAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThemeAliases, varargs...)
	ret0, _ := ret[0].(*quicksight.ListThemeAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListThemeAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThemeAliases, reflect.TypeOf((*MockQuicksightClient)(nil).ListThemeAliases), varargs...)
}

func (m *MockQuicksightClient) ListThemeVersions(arg0 context.Context, arg1 *quicksight.ListThemeVersionsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListThemeVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThemeVersions, varargs...)
	ret0, _ := ret[0].(*quicksight.ListThemeVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListThemeVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThemeVersions, reflect.TypeOf((*MockQuicksightClient)(nil).ListThemeVersions), varargs...)
}

func (m *MockQuicksightClient) ListThemes(arg0 context.Context, arg1 *quicksight.ListThemesInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListThemesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThemes, varargs...)
	ret0, _ := ret[0].(*quicksight.ListThemesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListThemes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThemes, reflect.TypeOf((*MockQuicksightClient)(nil).ListThemes), varargs...)
}

func (m *MockQuicksightClient) ListUserGroups(arg0 context.Context, arg1 *quicksight.ListUserGroupsInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListUserGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserGroups, varargs...)
	ret0, _ := ret[0].(*quicksight.ListUserGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListUserGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserGroups, reflect.TypeOf((*MockQuicksightClient)(nil).ListUserGroups), varargs...)
}

func (m *MockQuicksightClient) ListUsers(arg0 context.Context, arg1 *quicksight.ListUsersInput, arg2 ...func(*quicksight.Options)) (*quicksight.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsers, varargs...)
	ret0, _ := ret[0].(*quicksight.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsers, reflect.TypeOf((*MockQuicksightClient)(nil).ListUsers), varargs...)
}

func (m *MockQuicksightClient) SearchAnalyses(arg0 context.Context, arg1 *quicksight.SearchAnalysesInput, arg2 ...func(*quicksight.Options)) (*quicksight.SearchAnalysesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchAnalyses, varargs...)
	ret0, _ := ret[0].(*quicksight.SearchAnalysesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) SearchAnalyses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchAnalyses, reflect.TypeOf((*MockQuicksightClient)(nil).SearchAnalyses), varargs...)
}

func (m *MockQuicksightClient) SearchDashboards(arg0 context.Context, arg1 *quicksight.SearchDashboardsInput, arg2 ...func(*quicksight.Options)) (*quicksight.SearchDashboardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchDashboards, varargs...)
	ret0, _ := ret[0].(*quicksight.SearchDashboardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) SearchDashboards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchDashboards, reflect.TypeOf((*MockQuicksightClient)(nil).SearchDashboards), varargs...)
}

func (m *MockQuicksightClient) SearchDataSets(arg0 context.Context, arg1 *quicksight.SearchDataSetsInput, arg2 ...func(*quicksight.Options)) (*quicksight.SearchDataSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchDataSets, varargs...)
	ret0, _ := ret[0].(*quicksight.SearchDataSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) SearchDataSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchDataSets, reflect.TypeOf((*MockQuicksightClient)(nil).SearchDataSets), varargs...)
}

func (m *MockQuicksightClient) SearchDataSources(arg0 context.Context, arg1 *quicksight.SearchDataSourcesInput, arg2 ...func(*quicksight.Options)) (*quicksight.SearchDataSourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchDataSources, varargs...)
	ret0, _ := ret[0].(*quicksight.SearchDataSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) SearchDataSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchDataSources, reflect.TypeOf((*MockQuicksightClient)(nil).SearchDataSources), varargs...)
}

func (m *MockQuicksightClient) SearchFolders(arg0 context.Context, arg1 *quicksight.SearchFoldersInput, arg2 ...func(*quicksight.Options)) (*quicksight.SearchFoldersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchFolders, varargs...)
	ret0, _ := ret[0].(*quicksight.SearchFoldersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) SearchFolders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchFolders, reflect.TypeOf((*MockQuicksightClient)(nil).SearchFolders), varargs...)
}

func (m *MockQuicksightClient) SearchGroups(arg0 context.Context, arg1 *quicksight.SearchGroupsInput, arg2 ...func(*quicksight.Options)) (*quicksight.SearchGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchGroups, varargs...)
	ret0, _ := ret[0].(*quicksight.SearchGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQuicksightClientMockRecorder) SearchGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchGroups, reflect.TypeOf((*MockQuicksightClient)(nil).SearchGroups), varargs...)
}
