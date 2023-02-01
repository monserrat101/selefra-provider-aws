package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gomock "github.com/golang/mock/gomock"
)

type MockDynamodbClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDynamodbClientMockRecorder
}

type MockDynamodbClientMockRecorder struct {
	mock *MockDynamodbClient
}

func NewMockDynamodbClient(ctrl *gomock.Controller) *MockDynamodbClient {
	mock := &MockDynamodbClient{ctrl: ctrl}
	mock.recorder = &MockDynamodbClientMockRecorder{mock}
	return mock
}

func (m *MockDynamodbClient) EXPECT() *MockDynamodbClientMockRecorder {
	return m.recorder
}

func (m *MockDynamodbClient) BatchGetItem(arg0 context.Context, arg1 *dynamodb.BatchGetItemInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetItem, varargs...)
	ret0, _ := ret[0].(*dynamodb.BatchGetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) BatchGetItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetItem, reflect.TypeOf((*MockDynamodbClient)(nil).BatchGetItem), varargs...)
}

func (m *MockDynamodbClient) DescribeBackup(arg0 context.Context, arg1 *dynamodb.DescribeBackupInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBackup, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeBackupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeBackup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBackup, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeBackup), varargs...)
}

func (m *MockDynamodbClient) DescribeContinuousBackups(arg0 context.Context, arg1 *dynamodb.DescribeContinuousBackupsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeContinuousBackups, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeContinuousBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeContinuousBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeContinuousBackups, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeContinuousBackups), varargs...)
}

func (m *MockDynamodbClient) DescribeContributorInsights(arg0 context.Context, arg1 *dynamodb.DescribeContributorInsightsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeContributorInsights, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeContributorInsightsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeContributorInsights(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeContributorInsights, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeContributorInsights), varargs...)
}

func (m *MockDynamodbClient) DescribeEndpoints(arg0 context.Context, arg1 *dynamodb.DescribeEndpointsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpoints, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpoints, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeEndpoints), varargs...)
}

func (m *MockDynamodbClient) DescribeExport(arg0 context.Context, arg1 *dynamodb.DescribeExportInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExport, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeExportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeExport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExport, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeExport), varargs...)
}

func (m *MockDynamodbClient) DescribeGlobalTable(arg0 context.Context, arg1 *dynamodb.DescribeGlobalTableInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGlobalTable, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeGlobalTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeGlobalTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGlobalTable, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeGlobalTable), varargs...)
}

func (m *MockDynamodbClient) DescribeGlobalTableSettings(arg0 context.Context, arg1 *dynamodb.DescribeGlobalTableSettingsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGlobalTableSettings, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeGlobalTableSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeGlobalTableSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGlobalTableSettings, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeGlobalTableSettings), varargs...)
}

func (m *MockDynamodbClient) DescribeImport(arg0 context.Context, arg1 *dynamodb.DescribeImportInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeImportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImport, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImport, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeImport), varargs...)
}

func (m *MockDynamodbClient) DescribeKinesisStreamingDestination(arg0 context.Context, arg1 *dynamodb.DescribeKinesisStreamingDestinationInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeKinesisStreamingDestination, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeKinesisStreamingDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeKinesisStreamingDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeKinesisStreamingDestination, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeKinesisStreamingDestination), varargs...)
}

func (m *MockDynamodbClient) DescribeLimits(arg0 context.Context, arg1 *dynamodb.DescribeLimitsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLimits, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLimits, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeLimits), varargs...)
}

func (m *MockDynamodbClient) DescribeTable(arg0 context.Context, arg1 *dynamodb.DescribeTableInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTable, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTable, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeTable), varargs...)
}

func (m *MockDynamodbClient) DescribeTableReplicaAutoScaling(arg0 context.Context, arg1 *dynamodb.DescribeTableReplicaAutoScalingInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTableReplicaAutoScaling, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeTableReplicaAutoScalingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeTableReplicaAutoScaling(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTableReplicaAutoScaling, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeTableReplicaAutoScaling), varargs...)
}

func (m *MockDynamodbClient) DescribeTimeToLive(arg0 context.Context, arg1 *dynamodb.DescribeTimeToLiveInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTimeToLive, varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeTimeToLiveOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) DescribeTimeToLive(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTimeToLive, reflect.TypeOf((*MockDynamodbClient)(nil).DescribeTimeToLive), varargs...)
}

func (m *MockDynamodbClient) GetItem(arg0 context.Context, arg1 *dynamodb.GetItemInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetItem, varargs...)
	ret0, _ := ret[0].(*dynamodb.GetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) GetItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetItem, reflect.TypeOf((*MockDynamodbClient)(nil).GetItem), varargs...)
}

func (m *MockDynamodbClient) ListBackups(arg0 context.Context, arg1 *dynamodb.ListBackupsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBackups, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackups, reflect.TypeOf((*MockDynamodbClient)(nil).ListBackups), varargs...)
}

func (m *MockDynamodbClient) ListContributorInsights(arg0 context.Context, arg1 *dynamodb.ListContributorInsightsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListContributorInsights, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListContributorInsightsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListContributorInsights(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListContributorInsights, reflect.TypeOf((*MockDynamodbClient)(nil).ListContributorInsights), varargs...)
}

func (m *MockDynamodbClient) ListExports(arg0 context.Context, arg1 *dynamodb.ListExportsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExports, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListExportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListExports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExports, reflect.TypeOf((*MockDynamodbClient)(nil).ListExports), varargs...)
}

func (m *MockDynamodbClient) ListGlobalTables(arg0 context.Context, arg1 *dynamodb.ListGlobalTablesInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGlobalTables, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListGlobalTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListGlobalTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGlobalTables, reflect.TypeOf((*MockDynamodbClient)(nil).ListGlobalTables), varargs...)
}

func (m *MockDynamodbClient) ListImports(arg0 context.Context, arg1 *dynamodb.ListImportsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListImportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImports, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListImportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListImports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImports, reflect.TypeOf((*MockDynamodbClient)(nil).ListImports), varargs...)
}

func (m *MockDynamodbClient) ListTables(arg0 context.Context, arg1 *dynamodb.ListTablesInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTables, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTables, reflect.TypeOf((*MockDynamodbClient)(nil).ListTables), varargs...)
}

func (m *MockDynamodbClient) ListTagsOfResource(arg0 context.Context, arg1 *dynamodb.ListTagsOfResourceInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsOfResource, varargs...)
	ret0, _ := ret[0].(*dynamodb.ListTagsOfResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamodbClientMockRecorder) ListTagsOfResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsOfResource, reflect.TypeOf((*MockDynamodbClient)(nil).ListTagsOfResource), varargs...)
}
