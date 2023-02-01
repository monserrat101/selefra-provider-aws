package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	backup "github.com/aws/aws-sdk-go-v2/service/backup"
	gomock "github.com/golang/mock/gomock"
)

type MockBackupClient struct {
	ctrl		*gomock.Controller
	recorder	*MockBackupClientMockRecorder
}

type MockBackupClientMockRecorder struct {
	mock *MockBackupClient
}

func NewMockBackupClient(ctrl *gomock.Controller) *MockBackupClient {
	mock := &MockBackupClient{ctrl: ctrl}
	mock.recorder = &MockBackupClientMockRecorder{mock}
	return mock
}

func (m *MockBackupClient) EXPECT() *MockBackupClientMockRecorder {
	return m.recorder
}

func (m *MockBackupClient) DescribeBackupJob(arg0 context.Context, arg1 *backup.DescribeBackupJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeBackupJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBackupJob, varargs...)
	ret0, _ := ret[0].(*backup.DescribeBackupJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeBackupJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBackupJob, reflect.TypeOf((*MockBackupClient)(nil).DescribeBackupJob), varargs...)
}

func (m *MockBackupClient) DescribeBackupVault(arg0 context.Context, arg1 *backup.DescribeBackupVaultInput, arg2 ...func(*backup.Options)) (*backup.DescribeBackupVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBackupVault, varargs...)
	ret0, _ := ret[0].(*backup.DescribeBackupVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeBackupVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBackupVault, reflect.TypeOf((*MockBackupClient)(nil).DescribeBackupVault), varargs...)
}

func (m *MockBackupClient) DescribeCopyJob(arg0 context.Context, arg1 *backup.DescribeCopyJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeCopyJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCopyJob, varargs...)
	ret0, _ := ret[0].(*backup.DescribeCopyJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeCopyJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCopyJob, reflect.TypeOf((*MockBackupClient)(nil).DescribeCopyJob), varargs...)
}

func (m *MockBackupClient) DescribeFramework(arg0 context.Context, arg1 *backup.DescribeFrameworkInput, arg2 ...func(*backup.Options)) (*backup.DescribeFrameworkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFramework, varargs...)
	ret0, _ := ret[0].(*backup.DescribeFrameworkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeFramework(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFramework, reflect.TypeOf((*MockBackupClient)(nil).DescribeFramework), varargs...)
}

func (m *MockBackupClient) DescribeGlobalSettings(arg0 context.Context, arg1 *backup.DescribeGlobalSettingsInput, arg2 ...func(*backup.Options)) (*backup.DescribeGlobalSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGlobalSettings, varargs...)
	ret0, _ := ret[0].(*backup.DescribeGlobalSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeGlobalSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGlobalSettings, reflect.TypeOf((*MockBackupClient)(nil).DescribeGlobalSettings), varargs...)
}

func (m *MockBackupClient) DescribeProtectedResource(arg0 context.Context, arg1 *backup.DescribeProtectedResourceInput, arg2 ...func(*backup.Options)) (*backup.DescribeProtectedResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProtectedResource, varargs...)
	ret0, _ := ret[0].(*backup.DescribeProtectedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeProtectedResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProtectedResource, reflect.TypeOf((*MockBackupClient)(nil).DescribeProtectedResource), varargs...)
}

func (m *MockBackupClient) DescribeRecoveryPoint(arg0 context.Context, arg1 *backup.DescribeRecoveryPointInput, arg2 ...func(*backup.Options)) (*backup.DescribeRecoveryPointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRecoveryPoint, varargs...)
	ret0, _ := ret[0].(*backup.DescribeRecoveryPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeRecoveryPoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRecoveryPoint, reflect.TypeOf((*MockBackupClient)(nil).DescribeRecoveryPoint), varargs...)
}

func (m *MockBackupClient) DescribeRegionSettings(arg0 context.Context, arg1 *backup.DescribeRegionSettingsInput, arg2 ...func(*backup.Options)) (*backup.DescribeRegionSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRegionSettings, varargs...)
	ret0, _ := ret[0].(*backup.DescribeRegionSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeRegionSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRegionSettings, reflect.TypeOf((*MockBackupClient)(nil).DescribeRegionSettings), varargs...)
}

func (m *MockBackupClient) DescribeReportJob(arg0 context.Context, arg1 *backup.DescribeReportJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeReportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReportJob, varargs...)
	ret0, _ := ret[0].(*backup.DescribeReportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeReportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReportJob, reflect.TypeOf((*MockBackupClient)(nil).DescribeReportJob), varargs...)
}

func (m *MockBackupClient) DescribeReportPlan(arg0 context.Context, arg1 *backup.DescribeReportPlanInput, arg2 ...func(*backup.Options)) (*backup.DescribeReportPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReportPlan, varargs...)
	ret0, _ := ret[0].(*backup.DescribeReportPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeReportPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReportPlan, reflect.TypeOf((*MockBackupClient)(nil).DescribeReportPlan), varargs...)
}

func (m *MockBackupClient) DescribeRestoreJob(arg0 context.Context, arg1 *backup.DescribeRestoreJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeRestoreJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRestoreJob, varargs...)
	ret0, _ := ret[0].(*backup.DescribeRestoreJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeRestoreJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRestoreJob, reflect.TypeOf((*MockBackupClient)(nil).DescribeRestoreJob), varargs...)
}

func (m *MockBackupClient) GetBackupPlan(arg0 context.Context, arg1 *backup.GetBackupPlanInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBackupPlan, varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBackupPlan, reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlan), varargs...)
}

func (m *MockBackupClient) GetBackupPlanFromJSON(arg0 context.Context, arg1 *backup.GetBackupPlanFromJSONInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanFromJSONOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBackupPlanFromJSON, varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanFromJSONOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupPlanFromJSON(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBackupPlanFromJSON, reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlanFromJSON), varargs...)
}

func (m *MockBackupClient) GetBackupPlanFromTemplate(arg0 context.Context, arg1 *backup.GetBackupPlanFromTemplateInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanFromTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBackupPlanFromTemplate, varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanFromTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupPlanFromTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBackupPlanFromTemplate, reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlanFromTemplate), varargs...)
}

func (m *MockBackupClient) GetBackupSelection(arg0 context.Context, arg1 *backup.GetBackupSelectionInput, arg2 ...func(*backup.Options)) (*backup.GetBackupSelectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBackupSelection, varargs...)
	ret0, _ := ret[0].(*backup.GetBackupSelectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupSelection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBackupSelection, reflect.TypeOf((*MockBackupClient)(nil).GetBackupSelection), varargs...)
}

func (m *MockBackupClient) GetBackupVaultAccessPolicy(arg0 context.Context, arg1 *backup.GetBackupVaultAccessPolicyInput, arg2 ...func(*backup.Options)) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBackupVaultAccessPolicy, varargs...)
	ret0, _ := ret[0].(*backup.GetBackupVaultAccessPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupVaultAccessPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBackupVaultAccessPolicy, reflect.TypeOf((*MockBackupClient)(nil).GetBackupVaultAccessPolicy), varargs...)
}

func (m *MockBackupClient) GetBackupVaultNotifications(arg0 context.Context, arg1 *backup.GetBackupVaultNotificationsInput, arg2 ...func(*backup.Options)) (*backup.GetBackupVaultNotificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBackupVaultNotifications, varargs...)
	ret0, _ := ret[0].(*backup.GetBackupVaultNotificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupVaultNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBackupVaultNotifications, reflect.TypeOf((*MockBackupClient)(nil).GetBackupVaultNotifications), varargs...)
}

func (m *MockBackupClient) GetLegalHold(arg0 context.Context, arg1 *backup.GetLegalHoldInput, arg2 ...func(*backup.Options)) (*backup.GetLegalHoldOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLegalHold, varargs...)
	ret0, _ := ret[0].(*backup.GetLegalHoldOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetLegalHold(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLegalHold, reflect.TypeOf((*MockBackupClient)(nil).GetLegalHold), varargs...)
}

func (m *MockBackupClient) GetRecoveryPointRestoreMetadata(arg0 context.Context, arg1 *backup.GetRecoveryPointRestoreMetadataInput, arg2 ...func(*backup.Options)) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRecoveryPointRestoreMetadata, varargs...)
	ret0, _ := ret[0].(*backup.GetRecoveryPointRestoreMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetRecoveryPointRestoreMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRecoveryPointRestoreMetadata, reflect.TypeOf((*MockBackupClient)(nil).GetRecoveryPointRestoreMetadata), varargs...)
}

func (m *MockBackupClient) GetSupportedResourceTypes(arg0 context.Context, arg1 *backup.GetSupportedResourceTypesInput, arg2 ...func(*backup.Options)) (*backup.GetSupportedResourceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSupportedResourceTypes, varargs...)
	ret0, _ := ret[0].(*backup.GetSupportedResourceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetSupportedResourceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSupportedResourceTypes, reflect.TypeOf((*MockBackupClient)(nil).GetSupportedResourceTypes), varargs...)
}

func (m *MockBackupClient) ListBackupJobs(arg0 context.Context, arg1 *backup.ListBackupJobsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackupJobs, varargs...)
	ret0, _ := ret[0].(*backup.ListBackupJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackupJobs, reflect.TypeOf((*MockBackupClient)(nil).ListBackupJobs), varargs...)
}

func (m *MockBackupClient) ListBackupPlanTemplates(arg0 context.Context, arg1 *backup.ListBackupPlanTemplatesInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlanTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackupPlanTemplates, varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlanTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupPlanTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackupPlanTemplates, reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlanTemplates), varargs...)
}

func (m *MockBackupClient) ListBackupPlanVersions(arg0 context.Context, arg1 *backup.ListBackupPlanVersionsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlanVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackupPlanVersions, varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlanVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupPlanVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackupPlanVersions, reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlanVersions), varargs...)
}

func (m *MockBackupClient) ListBackupPlans(arg0 context.Context, arg1 *backup.ListBackupPlansInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackupPlans, varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackupPlans, reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlans), varargs...)
}

func (m *MockBackupClient) ListBackupSelections(arg0 context.Context, arg1 *backup.ListBackupSelectionsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupSelectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackupSelections, varargs...)
	ret0, _ := ret[0].(*backup.ListBackupSelectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupSelections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackupSelections, reflect.TypeOf((*MockBackupClient)(nil).ListBackupSelections), varargs...)
}

func (m *MockBackupClient) ListBackupVaults(arg0 context.Context, arg1 *backup.ListBackupVaultsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupVaultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackupVaults, varargs...)
	ret0, _ := ret[0].(*backup.ListBackupVaultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupVaults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackupVaults, reflect.TypeOf((*MockBackupClient)(nil).ListBackupVaults), varargs...)
}

func (m *MockBackupClient) ListCopyJobs(arg0 context.Context, arg1 *backup.ListCopyJobsInput, arg2 ...func(*backup.Options)) (*backup.ListCopyJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCopyJobs, varargs...)
	ret0, _ := ret[0].(*backup.ListCopyJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListCopyJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCopyJobs, reflect.TypeOf((*MockBackupClient)(nil).ListCopyJobs), varargs...)
}

func (m *MockBackupClient) ListFrameworks(arg0 context.Context, arg1 *backup.ListFrameworksInput, arg2 ...func(*backup.Options)) (*backup.ListFrameworksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFrameworks, varargs...)
	ret0, _ := ret[0].(*backup.ListFrameworksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListFrameworks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFrameworks, reflect.TypeOf((*MockBackupClient)(nil).ListFrameworks), varargs...)
}

func (m *MockBackupClient) ListLegalHolds(arg0 context.Context, arg1 *backup.ListLegalHoldsInput, arg2 ...func(*backup.Options)) (*backup.ListLegalHoldsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLegalHolds, varargs...)
	ret0, _ := ret[0].(*backup.ListLegalHoldsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListLegalHolds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLegalHolds, reflect.TypeOf((*MockBackupClient)(nil).ListLegalHolds), varargs...)
}

func (m *MockBackupClient) ListProtectedResources(arg0 context.Context, arg1 *backup.ListProtectedResourcesInput, arg2 ...func(*backup.Options)) (*backup.ListProtectedResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProtectedResources, varargs...)
	ret0, _ := ret[0].(*backup.ListProtectedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListProtectedResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProtectedResources, reflect.TypeOf((*MockBackupClient)(nil).ListProtectedResources), varargs...)
}

func (m *MockBackupClient) ListRecoveryPointsByBackupVault(arg0 context.Context, arg1 *backup.ListRecoveryPointsByBackupVaultInput, arg2 ...func(*backup.Options)) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRecoveryPointsByBackupVault, varargs...)
	ret0, _ := ret[0].(*backup.ListRecoveryPointsByBackupVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListRecoveryPointsByBackupVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRecoveryPointsByBackupVault, reflect.TypeOf((*MockBackupClient)(nil).ListRecoveryPointsByBackupVault), varargs...)
}

func (m *MockBackupClient) ListRecoveryPointsByLegalHold(arg0 context.Context, arg1 *backup.ListRecoveryPointsByLegalHoldInput, arg2 ...func(*backup.Options)) (*backup.ListRecoveryPointsByLegalHoldOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRecoveryPointsByLegalHold, varargs...)
	ret0, _ := ret[0].(*backup.ListRecoveryPointsByLegalHoldOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListRecoveryPointsByLegalHold(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRecoveryPointsByLegalHold, reflect.TypeOf((*MockBackupClient)(nil).ListRecoveryPointsByLegalHold), varargs...)
}

func (m *MockBackupClient) ListRecoveryPointsByResource(arg0 context.Context, arg1 *backup.ListRecoveryPointsByResourceInput, arg2 ...func(*backup.Options)) (*backup.ListRecoveryPointsByResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRecoveryPointsByResource, varargs...)
	ret0, _ := ret[0].(*backup.ListRecoveryPointsByResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListRecoveryPointsByResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRecoveryPointsByResource, reflect.TypeOf((*MockBackupClient)(nil).ListRecoveryPointsByResource), varargs...)
}

func (m *MockBackupClient) ListReportJobs(arg0 context.Context, arg1 *backup.ListReportJobsInput, arg2 ...func(*backup.Options)) (*backup.ListReportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReportJobs, varargs...)
	ret0, _ := ret[0].(*backup.ListReportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListReportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReportJobs, reflect.TypeOf((*MockBackupClient)(nil).ListReportJobs), varargs...)
}

func (m *MockBackupClient) ListReportPlans(arg0 context.Context, arg1 *backup.ListReportPlansInput, arg2 ...func(*backup.Options)) (*backup.ListReportPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReportPlans, varargs...)
	ret0, _ := ret[0].(*backup.ListReportPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListReportPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReportPlans, reflect.TypeOf((*MockBackupClient)(nil).ListReportPlans), varargs...)
}

func (m *MockBackupClient) ListRestoreJobs(arg0 context.Context, arg1 *backup.ListRestoreJobsInput, arg2 ...func(*backup.Options)) (*backup.ListRestoreJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRestoreJobs, varargs...)
	ret0, _ := ret[0].(*backup.ListRestoreJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListRestoreJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRestoreJobs, reflect.TypeOf((*MockBackupClient)(nil).ListRestoreJobs), varargs...)
}

func (m *MockBackupClient) ListTags(arg0 context.Context, arg1 *backup.ListTagsInput, arg2 ...func(*backup.Options)) (*backup.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*backup.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockBackupClient)(nil).ListTags), varargs...)
}
