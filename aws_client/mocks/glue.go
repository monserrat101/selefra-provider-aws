package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	glue "github.com/aws/aws-sdk-go-v2/service/glue"
	gomock "github.com/golang/mock/gomock"
)

type MockGlueClient struct {
	ctrl		*gomock.Controller
	recorder	*MockGlueClientMockRecorder
}

type MockGlueClientMockRecorder struct {
	mock *MockGlueClient
}

func NewMockGlueClient(ctrl *gomock.Controller) *MockGlueClient {
	mock := &MockGlueClient{ctrl: ctrl}
	mock.recorder = &MockGlueClientMockRecorder{mock}
	return mock
}

func (m *MockGlueClient) EXPECT() *MockGlueClientMockRecorder {
	return m.recorder
}

func (m *MockGlueClient) BatchGetBlueprints(arg0 context.Context, arg1 *glue.BatchGetBlueprintsInput, arg2 ...func(*glue.Options)) (*glue.BatchGetBlueprintsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetBlueprints, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetBlueprintsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetBlueprints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetBlueprints, reflect.TypeOf((*MockGlueClient)(nil).BatchGetBlueprints), varargs...)
}

func (m *MockGlueClient) BatchGetCrawlers(arg0 context.Context, arg1 *glue.BatchGetCrawlersInput, arg2 ...func(*glue.Options)) (*glue.BatchGetCrawlersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetCrawlers, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetCrawlersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetCrawlers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetCrawlers, reflect.TypeOf((*MockGlueClient)(nil).BatchGetCrawlers), varargs...)
}

func (m *MockGlueClient) BatchGetCustomEntityTypes(arg0 context.Context, arg1 *glue.BatchGetCustomEntityTypesInput, arg2 ...func(*glue.Options)) (*glue.BatchGetCustomEntityTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetCustomEntityTypes, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetCustomEntityTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetCustomEntityTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetCustomEntityTypes, reflect.TypeOf((*MockGlueClient)(nil).BatchGetCustomEntityTypes), varargs...)
}

func (m *MockGlueClient) BatchGetDataQualityResult(arg0 context.Context, arg1 *glue.BatchGetDataQualityResultInput, arg2 ...func(*glue.Options)) (*glue.BatchGetDataQualityResultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetDataQualityResult, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetDataQualityResultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetDataQualityResult(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetDataQualityResult, reflect.TypeOf((*MockGlueClient)(nil).BatchGetDataQualityResult), varargs...)
}

func (m *MockGlueClient) BatchGetDevEndpoints(arg0 context.Context, arg1 *glue.BatchGetDevEndpointsInput, arg2 ...func(*glue.Options)) (*glue.BatchGetDevEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetDevEndpoints, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetDevEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetDevEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetDevEndpoints, reflect.TypeOf((*MockGlueClient)(nil).BatchGetDevEndpoints), varargs...)
}

func (m *MockGlueClient) BatchGetJobs(arg0 context.Context, arg1 *glue.BatchGetJobsInput, arg2 ...func(*glue.Options)) (*glue.BatchGetJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetJobs, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetJobs, reflect.TypeOf((*MockGlueClient)(nil).BatchGetJobs), varargs...)
}

func (m *MockGlueClient) BatchGetPartition(arg0 context.Context, arg1 *glue.BatchGetPartitionInput, arg2 ...func(*glue.Options)) (*glue.BatchGetPartitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetPartition, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetPartitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetPartition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetPartition, reflect.TypeOf((*MockGlueClient)(nil).BatchGetPartition), varargs...)
}

func (m *MockGlueClient) BatchGetTriggers(arg0 context.Context, arg1 *glue.BatchGetTriggersInput, arg2 ...func(*glue.Options)) (*glue.BatchGetTriggersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetTriggers, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetTriggersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetTriggers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetTriggers, reflect.TypeOf((*MockGlueClient)(nil).BatchGetTriggers), varargs...)
}

func (m *MockGlueClient) BatchGetWorkflows(arg0 context.Context, arg1 *glue.BatchGetWorkflowsInput, arg2 ...func(*glue.Options)) (*glue.BatchGetWorkflowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetWorkflows, varargs...)
	ret0, _ := ret[0].(*glue.BatchGetWorkflowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) BatchGetWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetWorkflows, reflect.TypeOf((*MockGlueClient)(nil).BatchGetWorkflows), varargs...)
}

func (m *MockGlueClient) GetBlueprint(arg0 context.Context, arg1 *glue.GetBlueprintInput, arg2 ...func(*glue.Options)) (*glue.GetBlueprintOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlueprint, varargs...)
	ret0, _ := ret[0].(*glue.GetBlueprintOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetBlueprint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlueprint, reflect.TypeOf((*MockGlueClient)(nil).GetBlueprint), varargs...)
}

func (m *MockGlueClient) GetBlueprintRun(arg0 context.Context, arg1 *glue.GetBlueprintRunInput, arg2 ...func(*glue.Options)) (*glue.GetBlueprintRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlueprintRun, varargs...)
	ret0, _ := ret[0].(*glue.GetBlueprintRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetBlueprintRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlueprintRun, reflect.TypeOf((*MockGlueClient)(nil).GetBlueprintRun), varargs...)
}

func (m *MockGlueClient) GetBlueprintRuns(arg0 context.Context, arg1 *glue.GetBlueprintRunsInput, arg2 ...func(*glue.Options)) (*glue.GetBlueprintRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlueprintRuns, varargs...)
	ret0, _ := ret[0].(*glue.GetBlueprintRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetBlueprintRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlueprintRuns, reflect.TypeOf((*MockGlueClient)(nil).GetBlueprintRuns), varargs...)
}

func (m *MockGlueClient) GetCatalogImportStatus(arg0 context.Context, arg1 *glue.GetCatalogImportStatusInput, arg2 ...func(*glue.Options)) (*glue.GetCatalogImportStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCatalogImportStatus, varargs...)
	ret0, _ := ret[0].(*glue.GetCatalogImportStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetCatalogImportStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCatalogImportStatus, reflect.TypeOf((*MockGlueClient)(nil).GetCatalogImportStatus), varargs...)
}

func (m *MockGlueClient) GetClassifier(arg0 context.Context, arg1 *glue.GetClassifierInput, arg2 ...func(*glue.Options)) (*glue.GetClassifierOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClassifier, varargs...)
	ret0, _ := ret[0].(*glue.GetClassifierOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetClassifier(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClassifier, reflect.TypeOf((*MockGlueClient)(nil).GetClassifier), varargs...)
}

func (m *MockGlueClient) GetClassifiers(arg0 context.Context, arg1 *glue.GetClassifiersInput, arg2 ...func(*glue.Options)) (*glue.GetClassifiersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClassifiers, varargs...)
	ret0, _ := ret[0].(*glue.GetClassifiersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetClassifiers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClassifiers, reflect.TypeOf((*MockGlueClient)(nil).GetClassifiers), varargs...)
}

func (m *MockGlueClient) GetColumnStatisticsForPartition(arg0 context.Context, arg1 *glue.GetColumnStatisticsForPartitionInput, arg2 ...func(*glue.Options)) (*glue.GetColumnStatisticsForPartitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetColumnStatisticsForPartition, varargs...)
	ret0, _ := ret[0].(*glue.GetColumnStatisticsForPartitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetColumnStatisticsForPartition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetColumnStatisticsForPartition, reflect.TypeOf((*MockGlueClient)(nil).GetColumnStatisticsForPartition), varargs...)
}

func (m *MockGlueClient) GetColumnStatisticsForTable(arg0 context.Context, arg1 *glue.GetColumnStatisticsForTableInput, arg2 ...func(*glue.Options)) (*glue.GetColumnStatisticsForTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetColumnStatisticsForTable, varargs...)
	ret0, _ := ret[0].(*glue.GetColumnStatisticsForTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetColumnStatisticsForTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetColumnStatisticsForTable, reflect.TypeOf((*MockGlueClient)(nil).GetColumnStatisticsForTable), varargs...)
}

func (m *MockGlueClient) GetConnection(arg0 context.Context, arg1 *glue.GetConnectionInput, arg2 ...func(*glue.Options)) (*glue.GetConnectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConnection, varargs...)
	ret0, _ := ret[0].(*glue.GetConnectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetConnection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConnection, reflect.TypeOf((*MockGlueClient)(nil).GetConnection), varargs...)
}

func (m *MockGlueClient) GetConnections(arg0 context.Context, arg1 *glue.GetConnectionsInput, arg2 ...func(*glue.Options)) (*glue.GetConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConnections, varargs...)
	ret0, _ := ret[0].(*glue.GetConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConnections, reflect.TypeOf((*MockGlueClient)(nil).GetConnections), varargs...)
}

func (m *MockGlueClient) GetCrawler(arg0 context.Context, arg1 *glue.GetCrawlerInput, arg2 ...func(*glue.Options)) (*glue.GetCrawlerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCrawler, varargs...)
	ret0, _ := ret[0].(*glue.GetCrawlerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetCrawler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCrawler, reflect.TypeOf((*MockGlueClient)(nil).GetCrawler), varargs...)
}

func (m *MockGlueClient) GetCrawlerMetrics(arg0 context.Context, arg1 *glue.GetCrawlerMetricsInput, arg2 ...func(*glue.Options)) (*glue.GetCrawlerMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCrawlerMetrics, varargs...)
	ret0, _ := ret[0].(*glue.GetCrawlerMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetCrawlerMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCrawlerMetrics, reflect.TypeOf((*MockGlueClient)(nil).GetCrawlerMetrics), varargs...)
}

func (m *MockGlueClient) GetCrawlers(arg0 context.Context, arg1 *glue.GetCrawlersInput, arg2 ...func(*glue.Options)) (*glue.GetCrawlersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCrawlers, varargs...)
	ret0, _ := ret[0].(*glue.GetCrawlersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetCrawlers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCrawlers, reflect.TypeOf((*MockGlueClient)(nil).GetCrawlers), varargs...)
}

func (m *MockGlueClient) GetCustomEntityType(arg0 context.Context, arg1 *glue.GetCustomEntityTypeInput, arg2 ...func(*glue.Options)) (*glue.GetCustomEntityTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCustomEntityType, varargs...)
	ret0, _ := ret[0].(*glue.GetCustomEntityTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetCustomEntityType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCustomEntityType, reflect.TypeOf((*MockGlueClient)(nil).GetCustomEntityType), varargs...)
}

func (m *MockGlueClient) GetDataCatalogEncryptionSettings(arg0 context.Context, arg1 *glue.GetDataCatalogEncryptionSettingsInput, arg2 ...func(*glue.Options)) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataCatalogEncryptionSettings, varargs...)
	ret0, _ := ret[0].(*glue.GetDataCatalogEncryptionSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataCatalogEncryptionSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataCatalogEncryptionSettings, reflect.TypeOf((*MockGlueClient)(nil).GetDataCatalogEncryptionSettings), varargs...)
}

func (m *MockGlueClient) GetDataQualityResult(arg0 context.Context, arg1 *glue.GetDataQualityResultInput, arg2 ...func(*glue.Options)) (*glue.GetDataQualityResultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataQualityResult, varargs...)
	ret0, _ := ret[0].(*glue.GetDataQualityResultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataQualityResult(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataQualityResult, reflect.TypeOf((*MockGlueClient)(nil).GetDataQualityResult), varargs...)
}

func (m *MockGlueClient) GetDataQualityRuleRecommendationRun(arg0 context.Context, arg1 *glue.GetDataQualityRuleRecommendationRunInput, arg2 ...func(*glue.Options)) (*glue.GetDataQualityRuleRecommendationRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataQualityRuleRecommendationRun, varargs...)
	ret0, _ := ret[0].(*glue.GetDataQualityRuleRecommendationRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataQualityRuleRecommendationRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataQualityRuleRecommendationRun, reflect.TypeOf((*MockGlueClient)(nil).GetDataQualityRuleRecommendationRun), varargs...)
}

func (m *MockGlueClient) GetDataQualityRuleset(arg0 context.Context, arg1 *glue.GetDataQualityRulesetInput, arg2 ...func(*glue.Options)) (*glue.GetDataQualityRulesetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataQualityRuleset, varargs...)
	ret0, _ := ret[0].(*glue.GetDataQualityRulesetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataQualityRuleset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataQualityRuleset, reflect.TypeOf((*MockGlueClient)(nil).GetDataQualityRuleset), varargs...)
}

func (m *MockGlueClient) GetDataQualityRulesetEvaluationRun(arg0 context.Context, arg1 *glue.GetDataQualityRulesetEvaluationRunInput, arg2 ...func(*glue.Options)) (*glue.GetDataQualityRulesetEvaluationRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataQualityRulesetEvaluationRun, varargs...)
	ret0, _ := ret[0].(*glue.GetDataQualityRulesetEvaluationRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataQualityRulesetEvaluationRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataQualityRulesetEvaluationRun, reflect.TypeOf((*MockGlueClient)(nil).GetDataQualityRulesetEvaluationRun), varargs...)
}

func (m *MockGlueClient) GetDatabase(arg0 context.Context, arg1 *glue.GetDatabaseInput, arg2 ...func(*glue.Options)) (*glue.GetDatabaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDatabase, varargs...)
	ret0, _ := ret[0].(*glue.GetDatabaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDatabase, reflect.TypeOf((*MockGlueClient)(nil).GetDatabase), varargs...)
}

func (m *MockGlueClient) GetDatabases(arg0 context.Context, arg1 *glue.GetDatabasesInput, arg2 ...func(*glue.Options)) (*glue.GetDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDatabases, varargs...)
	ret0, _ := ret[0].(*glue.GetDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDatabases, reflect.TypeOf((*MockGlueClient)(nil).GetDatabases), varargs...)
}

func (m *MockGlueClient) GetDataflowGraph(arg0 context.Context, arg1 *glue.GetDataflowGraphInput, arg2 ...func(*glue.Options)) (*glue.GetDataflowGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataflowGraph, varargs...)
	ret0, _ := ret[0].(*glue.GetDataflowGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataflowGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataflowGraph, reflect.TypeOf((*MockGlueClient)(nil).GetDataflowGraph), varargs...)
}

func (m *MockGlueClient) GetDevEndpoint(arg0 context.Context, arg1 *glue.GetDevEndpointInput, arg2 ...func(*glue.Options)) (*glue.GetDevEndpointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDevEndpoint, varargs...)
	ret0, _ := ret[0].(*glue.GetDevEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDevEndpoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDevEndpoint, reflect.TypeOf((*MockGlueClient)(nil).GetDevEndpoint), varargs...)
}

func (m *MockGlueClient) GetDevEndpoints(arg0 context.Context, arg1 *glue.GetDevEndpointsInput, arg2 ...func(*glue.Options)) (*glue.GetDevEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDevEndpoints, varargs...)
	ret0, _ := ret[0].(*glue.GetDevEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDevEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDevEndpoints, reflect.TypeOf((*MockGlueClient)(nil).GetDevEndpoints), varargs...)
}

func (m *MockGlueClient) GetJob(arg0 context.Context, arg1 *glue.GetJobInput, arg2 ...func(*glue.Options)) (*glue.GetJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJob, varargs...)
	ret0, _ := ret[0].(*glue.GetJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJob, reflect.TypeOf((*MockGlueClient)(nil).GetJob), varargs...)
}

func (m *MockGlueClient) GetJobBookmark(arg0 context.Context, arg1 *glue.GetJobBookmarkInput, arg2 ...func(*glue.Options)) (*glue.GetJobBookmarkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobBookmark, varargs...)
	ret0, _ := ret[0].(*glue.GetJobBookmarkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJobBookmark(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobBookmark, reflect.TypeOf((*MockGlueClient)(nil).GetJobBookmark), varargs...)
}

func (m *MockGlueClient) GetJobRun(arg0 context.Context, arg1 *glue.GetJobRunInput, arg2 ...func(*glue.Options)) (*glue.GetJobRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobRun, varargs...)
	ret0, _ := ret[0].(*glue.GetJobRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJobRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobRun, reflect.TypeOf((*MockGlueClient)(nil).GetJobRun), varargs...)
}

func (m *MockGlueClient) GetJobRuns(arg0 context.Context, arg1 *glue.GetJobRunsInput, arg2 ...func(*glue.Options)) (*glue.GetJobRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobRuns, varargs...)
	ret0, _ := ret[0].(*glue.GetJobRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJobRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobRuns, reflect.TypeOf((*MockGlueClient)(nil).GetJobRuns), varargs...)
}

func (m *MockGlueClient) GetJobs(arg0 context.Context, arg1 *glue.GetJobsInput, arg2 ...func(*glue.Options)) (*glue.GetJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobs, varargs...)
	ret0, _ := ret[0].(*glue.GetJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobs, reflect.TypeOf((*MockGlueClient)(nil).GetJobs), varargs...)
}

func (m *MockGlueClient) GetMLTaskRun(arg0 context.Context, arg1 *glue.GetMLTaskRunInput, arg2 ...func(*glue.Options)) (*glue.GetMLTaskRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMLTaskRun, varargs...)
	ret0, _ := ret[0].(*glue.GetMLTaskRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMLTaskRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMLTaskRun, reflect.TypeOf((*MockGlueClient)(nil).GetMLTaskRun), varargs...)
}

func (m *MockGlueClient) GetMLTaskRuns(arg0 context.Context, arg1 *glue.GetMLTaskRunsInput, arg2 ...func(*glue.Options)) (*glue.GetMLTaskRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMLTaskRuns, varargs...)
	ret0, _ := ret[0].(*glue.GetMLTaskRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMLTaskRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMLTaskRuns, reflect.TypeOf((*MockGlueClient)(nil).GetMLTaskRuns), varargs...)
}

func (m *MockGlueClient) GetMLTransform(arg0 context.Context, arg1 *glue.GetMLTransformInput, arg2 ...func(*glue.Options)) (*glue.GetMLTransformOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMLTransform, varargs...)
	ret0, _ := ret[0].(*glue.GetMLTransformOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMLTransform(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMLTransform, reflect.TypeOf((*MockGlueClient)(nil).GetMLTransform), varargs...)
}

func (m *MockGlueClient) GetMLTransforms(arg0 context.Context, arg1 *glue.GetMLTransformsInput, arg2 ...func(*glue.Options)) (*glue.GetMLTransformsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMLTransforms, varargs...)
	ret0, _ := ret[0].(*glue.GetMLTransformsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMLTransforms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMLTransforms, reflect.TypeOf((*MockGlueClient)(nil).GetMLTransforms), varargs...)
}

func (m *MockGlueClient) GetMapping(arg0 context.Context, arg1 *glue.GetMappingInput, arg2 ...func(*glue.Options)) (*glue.GetMappingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMapping, varargs...)
	ret0, _ := ret[0].(*glue.GetMappingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMapping, reflect.TypeOf((*MockGlueClient)(nil).GetMapping), varargs...)
}

func (m *MockGlueClient) GetPartition(arg0 context.Context, arg1 *glue.GetPartitionInput, arg2 ...func(*glue.Options)) (*glue.GetPartitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPartition, varargs...)
	ret0, _ := ret[0].(*glue.GetPartitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetPartition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPartition, reflect.TypeOf((*MockGlueClient)(nil).GetPartition), varargs...)
}

func (m *MockGlueClient) GetPartitionIndexes(arg0 context.Context, arg1 *glue.GetPartitionIndexesInput, arg2 ...func(*glue.Options)) (*glue.GetPartitionIndexesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPartitionIndexes, varargs...)
	ret0, _ := ret[0].(*glue.GetPartitionIndexesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetPartitionIndexes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPartitionIndexes, reflect.TypeOf((*MockGlueClient)(nil).GetPartitionIndexes), varargs...)
}

func (m *MockGlueClient) GetPartitions(arg0 context.Context, arg1 *glue.GetPartitionsInput, arg2 ...func(*glue.Options)) (*glue.GetPartitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPartitions, varargs...)
	ret0, _ := ret[0].(*glue.GetPartitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetPartitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPartitions, reflect.TypeOf((*MockGlueClient)(nil).GetPartitions), varargs...)
}

func (m *MockGlueClient) GetPlan(arg0 context.Context, arg1 *glue.GetPlanInput, arg2 ...func(*glue.Options)) (*glue.GetPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPlan, varargs...)
	ret0, _ := ret[0].(*glue.GetPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPlan, reflect.TypeOf((*MockGlueClient)(nil).GetPlan), varargs...)
}

func (m *MockGlueClient) GetRegistry(arg0 context.Context, arg1 *glue.GetRegistryInput, arg2 ...func(*glue.Options)) (*glue.GetRegistryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegistry, varargs...)
	ret0, _ := ret[0].(*glue.GetRegistryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetRegistry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegistry, reflect.TypeOf((*MockGlueClient)(nil).GetRegistry), varargs...)
}

func (m *MockGlueClient) GetResourcePolicies(arg0 context.Context, arg1 *glue.GetResourcePoliciesInput, arg2 ...func(*glue.Options)) (*glue.GetResourcePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourcePolicies, varargs...)
	ret0, _ := ret[0].(*glue.GetResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourcePolicies, reflect.TypeOf((*MockGlueClient)(nil).GetResourcePolicies), varargs...)
}

func (m *MockGlueClient) GetResourcePolicy(arg0 context.Context, arg1 *glue.GetResourcePolicyInput, arg2 ...func(*glue.Options)) (*glue.GetResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourcePolicy, varargs...)
	ret0, _ := ret[0].(*glue.GetResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourcePolicy, reflect.TypeOf((*MockGlueClient)(nil).GetResourcePolicy), varargs...)
}

func (m *MockGlueClient) GetSchema(arg0 context.Context, arg1 *glue.GetSchemaInput, arg2 ...func(*glue.Options)) (*glue.GetSchemaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSchema, varargs...)
	ret0, _ := ret[0].(*glue.GetSchemaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSchema(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSchema, reflect.TypeOf((*MockGlueClient)(nil).GetSchema), varargs...)
}

func (m *MockGlueClient) GetSchemaByDefinition(arg0 context.Context, arg1 *glue.GetSchemaByDefinitionInput, arg2 ...func(*glue.Options)) (*glue.GetSchemaByDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSchemaByDefinition, varargs...)
	ret0, _ := ret[0].(*glue.GetSchemaByDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSchemaByDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSchemaByDefinition, reflect.TypeOf((*MockGlueClient)(nil).GetSchemaByDefinition), varargs...)
}

func (m *MockGlueClient) GetSchemaVersion(arg0 context.Context, arg1 *glue.GetSchemaVersionInput, arg2 ...func(*glue.Options)) (*glue.GetSchemaVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSchemaVersion, varargs...)
	ret0, _ := ret[0].(*glue.GetSchemaVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSchemaVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSchemaVersion, reflect.TypeOf((*MockGlueClient)(nil).GetSchemaVersion), varargs...)
}

func (m *MockGlueClient) GetSchemaVersionsDiff(arg0 context.Context, arg1 *glue.GetSchemaVersionsDiffInput, arg2 ...func(*glue.Options)) (*glue.GetSchemaVersionsDiffOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSchemaVersionsDiff, varargs...)
	ret0, _ := ret[0].(*glue.GetSchemaVersionsDiffOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSchemaVersionsDiff(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSchemaVersionsDiff, reflect.TypeOf((*MockGlueClient)(nil).GetSchemaVersionsDiff), varargs...)
}

func (m *MockGlueClient) GetSecurityConfiguration(arg0 context.Context, arg1 *glue.GetSecurityConfigurationInput, arg2 ...func(*glue.Options)) (*glue.GetSecurityConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSecurityConfiguration, varargs...)
	ret0, _ := ret[0].(*glue.GetSecurityConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSecurityConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSecurityConfiguration, reflect.TypeOf((*MockGlueClient)(nil).GetSecurityConfiguration), varargs...)
}

func (m *MockGlueClient) GetSecurityConfigurations(arg0 context.Context, arg1 *glue.GetSecurityConfigurationsInput, arg2 ...func(*glue.Options)) (*glue.GetSecurityConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSecurityConfigurations, varargs...)
	ret0, _ := ret[0].(*glue.GetSecurityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSecurityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSecurityConfigurations, reflect.TypeOf((*MockGlueClient)(nil).GetSecurityConfigurations), varargs...)
}

func (m *MockGlueClient) GetSession(arg0 context.Context, arg1 *glue.GetSessionInput, arg2 ...func(*glue.Options)) (*glue.GetSessionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSession, varargs...)
	ret0, _ := ret[0].(*glue.GetSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSession(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSession, reflect.TypeOf((*MockGlueClient)(nil).GetSession), varargs...)
}

func (m *MockGlueClient) GetStatement(arg0 context.Context, arg1 *glue.GetStatementInput, arg2 ...func(*glue.Options)) (*glue.GetStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStatement, varargs...)
	ret0, _ := ret[0].(*glue.GetStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetStatement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStatement, reflect.TypeOf((*MockGlueClient)(nil).GetStatement), varargs...)
}

func (m *MockGlueClient) GetTable(arg0 context.Context, arg1 *glue.GetTableInput, arg2 ...func(*glue.Options)) (*glue.GetTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTable, varargs...)
	ret0, _ := ret[0].(*glue.GetTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTable, reflect.TypeOf((*MockGlueClient)(nil).GetTable), varargs...)
}

func (m *MockGlueClient) GetTableVersion(arg0 context.Context, arg1 *glue.GetTableVersionInput, arg2 ...func(*glue.Options)) (*glue.GetTableVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTableVersion, varargs...)
	ret0, _ := ret[0].(*glue.GetTableVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTableVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTableVersion, reflect.TypeOf((*MockGlueClient)(nil).GetTableVersion), varargs...)
}

func (m *MockGlueClient) GetTableVersions(arg0 context.Context, arg1 *glue.GetTableVersionsInput, arg2 ...func(*glue.Options)) (*glue.GetTableVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTableVersions, varargs...)
	ret0, _ := ret[0].(*glue.GetTableVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTableVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTableVersions, reflect.TypeOf((*MockGlueClient)(nil).GetTableVersions), varargs...)
}

func (m *MockGlueClient) GetTables(arg0 context.Context, arg1 *glue.GetTablesInput, arg2 ...func(*glue.Options)) (*glue.GetTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTables, varargs...)
	ret0, _ := ret[0].(*glue.GetTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTables, reflect.TypeOf((*MockGlueClient)(nil).GetTables), varargs...)
}

func (m *MockGlueClient) GetTags(arg0 context.Context, arg1 *glue.GetTagsInput, arg2 ...func(*glue.Options)) (*glue.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTags, varargs...)
	ret0, _ := ret[0].(*glue.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTags, reflect.TypeOf((*MockGlueClient)(nil).GetTags), varargs...)
}

func (m *MockGlueClient) GetTrigger(arg0 context.Context, arg1 *glue.GetTriggerInput, arg2 ...func(*glue.Options)) (*glue.GetTriggerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTrigger, varargs...)
	ret0, _ := ret[0].(*glue.GetTriggerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTrigger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTrigger, reflect.TypeOf((*MockGlueClient)(nil).GetTrigger), varargs...)
}

func (m *MockGlueClient) GetTriggers(arg0 context.Context, arg1 *glue.GetTriggersInput, arg2 ...func(*glue.Options)) (*glue.GetTriggersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTriggers, varargs...)
	ret0, _ := ret[0].(*glue.GetTriggersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTriggers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTriggers, reflect.TypeOf((*MockGlueClient)(nil).GetTriggers), varargs...)
}

func (m *MockGlueClient) GetUnfilteredPartitionMetadata(arg0 context.Context, arg1 *glue.GetUnfilteredPartitionMetadataInput, arg2 ...func(*glue.Options)) (*glue.GetUnfilteredPartitionMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUnfilteredPartitionMetadata, varargs...)
	ret0, _ := ret[0].(*glue.GetUnfilteredPartitionMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetUnfilteredPartitionMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUnfilteredPartitionMetadata, reflect.TypeOf((*MockGlueClient)(nil).GetUnfilteredPartitionMetadata), varargs...)
}

func (m *MockGlueClient) GetUnfilteredPartitionsMetadata(arg0 context.Context, arg1 *glue.GetUnfilteredPartitionsMetadataInput, arg2 ...func(*glue.Options)) (*glue.GetUnfilteredPartitionsMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUnfilteredPartitionsMetadata, varargs...)
	ret0, _ := ret[0].(*glue.GetUnfilteredPartitionsMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetUnfilteredPartitionsMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUnfilteredPartitionsMetadata, reflect.TypeOf((*MockGlueClient)(nil).GetUnfilteredPartitionsMetadata), varargs...)
}

func (m *MockGlueClient) GetUnfilteredTableMetadata(arg0 context.Context, arg1 *glue.GetUnfilteredTableMetadataInput, arg2 ...func(*glue.Options)) (*glue.GetUnfilteredTableMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUnfilteredTableMetadata, varargs...)
	ret0, _ := ret[0].(*glue.GetUnfilteredTableMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetUnfilteredTableMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUnfilteredTableMetadata, reflect.TypeOf((*MockGlueClient)(nil).GetUnfilteredTableMetadata), varargs...)
}

func (m *MockGlueClient) GetUserDefinedFunction(arg0 context.Context, arg1 *glue.GetUserDefinedFunctionInput, arg2 ...func(*glue.Options)) (*glue.GetUserDefinedFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUserDefinedFunction, varargs...)
	ret0, _ := ret[0].(*glue.GetUserDefinedFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetUserDefinedFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUserDefinedFunction, reflect.TypeOf((*MockGlueClient)(nil).GetUserDefinedFunction), varargs...)
}

func (m *MockGlueClient) GetUserDefinedFunctions(arg0 context.Context, arg1 *glue.GetUserDefinedFunctionsInput, arg2 ...func(*glue.Options)) (*glue.GetUserDefinedFunctionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUserDefinedFunctions, varargs...)
	ret0, _ := ret[0].(*glue.GetUserDefinedFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetUserDefinedFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUserDefinedFunctions, reflect.TypeOf((*MockGlueClient)(nil).GetUserDefinedFunctions), varargs...)
}

func (m *MockGlueClient) GetWorkflow(arg0 context.Context, arg1 *glue.GetWorkflowInput, arg2 ...func(*glue.Options)) (*glue.GetWorkflowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWorkflow, varargs...)
	ret0, _ := ret[0].(*glue.GetWorkflowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWorkflow, reflect.TypeOf((*MockGlueClient)(nil).GetWorkflow), varargs...)
}

func (m *MockGlueClient) GetWorkflowRun(arg0 context.Context, arg1 *glue.GetWorkflowRunInput, arg2 ...func(*glue.Options)) (*glue.GetWorkflowRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWorkflowRun, varargs...)
	ret0, _ := ret[0].(*glue.GetWorkflowRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetWorkflowRun(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWorkflowRun, reflect.TypeOf((*MockGlueClient)(nil).GetWorkflowRun), varargs...)
}

func (m *MockGlueClient) GetWorkflowRunProperties(arg0 context.Context, arg1 *glue.GetWorkflowRunPropertiesInput, arg2 ...func(*glue.Options)) (*glue.GetWorkflowRunPropertiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWorkflowRunProperties, varargs...)
	ret0, _ := ret[0].(*glue.GetWorkflowRunPropertiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetWorkflowRunProperties(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWorkflowRunProperties, reflect.TypeOf((*MockGlueClient)(nil).GetWorkflowRunProperties), varargs...)
}

func (m *MockGlueClient) GetWorkflowRuns(arg0 context.Context, arg1 *glue.GetWorkflowRunsInput, arg2 ...func(*glue.Options)) (*glue.GetWorkflowRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWorkflowRuns, varargs...)
	ret0, _ := ret[0].(*glue.GetWorkflowRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetWorkflowRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWorkflowRuns, reflect.TypeOf((*MockGlueClient)(nil).GetWorkflowRuns), varargs...)
}

func (m *MockGlueClient) ListBlueprints(arg0 context.Context, arg1 *glue.ListBlueprintsInput, arg2 ...func(*glue.Options)) (*glue.ListBlueprintsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBlueprints, varargs...)
	ret0, _ := ret[0].(*glue.ListBlueprintsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListBlueprints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBlueprints, reflect.TypeOf((*MockGlueClient)(nil).ListBlueprints), varargs...)
}

func (m *MockGlueClient) ListCrawlers(arg0 context.Context, arg1 *glue.ListCrawlersInput, arg2 ...func(*glue.Options)) (*glue.ListCrawlersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCrawlers, varargs...)
	ret0, _ := ret[0].(*glue.ListCrawlersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListCrawlers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCrawlers, reflect.TypeOf((*MockGlueClient)(nil).ListCrawlers), varargs...)
}

func (m *MockGlueClient) ListCrawls(arg0 context.Context, arg1 *glue.ListCrawlsInput, arg2 ...func(*glue.Options)) (*glue.ListCrawlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCrawls, varargs...)
	ret0, _ := ret[0].(*glue.ListCrawlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListCrawls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCrawls, reflect.TypeOf((*MockGlueClient)(nil).ListCrawls), varargs...)
}

func (m *MockGlueClient) ListCustomEntityTypes(arg0 context.Context, arg1 *glue.ListCustomEntityTypesInput, arg2 ...func(*glue.Options)) (*glue.ListCustomEntityTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCustomEntityTypes, varargs...)
	ret0, _ := ret[0].(*glue.ListCustomEntityTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListCustomEntityTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCustomEntityTypes, reflect.TypeOf((*MockGlueClient)(nil).ListCustomEntityTypes), varargs...)
}

func (m *MockGlueClient) ListDataQualityResults(arg0 context.Context, arg1 *glue.ListDataQualityResultsInput, arg2 ...func(*glue.Options)) (*glue.ListDataQualityResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataQualityResults, varargs...)
	ret0, _ := ret[0].(*glue.ListDataQualityResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListDataQualityResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataQualityResults, reflect.TypeOf((*MockGlueClient)(nil).ListDataQualityResults), varargs...)
}

func (m *MockGlueClient) ListDataQualityRuleRecommendationRuns(arg0 context.Context, arg1 *glue.ListDataQualityRuleRecommendationRunsInput, arg2 ...func(*glue.Options)) (*glue.ListDataQualityRuleRecommendationRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataQualityRuleRecommendationRuns, varargs...)
	ret0, _ := ret[0].(*glue.ListDataQualityRuleRecommendationRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListDataQualityRuleRecommendationRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataQualityRuleRecommendationRuns, reflect.TypeOf((*MockGlueClient)(nil).ListDataQualityRuleRecommendationRuns), varargs...)
}

func (m *MockGlueClient) ListDataQualityRulesetEvaluationRuns(arg0 context.Context, arg1 *glue.ListDataQualityRulesetEvaluationRunsInput, arg2 ...func(*glue.Options)) (*glue.ListDataQualityRulesetEvaluationRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataQualityRulesetEvaluationRuns, varargs...)
	ret0, _ := ret[0].(*glue.ListDataQualityRulesetEvaluationRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListDataQualityRulesetEvaluationRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataQualityRulesetEvaluationRuns, reflect.TypeOf((*MockGlueClient)(nil).ListDataQualityRulesetEvaluationRuns), varargs...)
}

func (m *MockGlueClient) ListDataQualityRulesets(arg0 context.Context, arg1 *glue.ListDataQualityRulesetsInput, arg2 ...func(*glue.Options)) (*glue.ListDataQualityRulesetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataQualityRulesets, varargs...)
	ret0, _ := ret[0].(*glue.ListDataQualityRulesetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListDataQualityRulesets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataQualityRulesets, reflect.TypeOf((*MockGlueClient)(nil).ListDataQualityRulesets), varargs...)
}

func (m *MockGlueClient) ListDevEndpoints(arg0 context.Context, arg1 *glue.ListDevEndpointsInput, arg2 ...func(*glue.Options)) (*glue.ListDevEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDevEndpoints, varargs...)
	ret0, _ := ret[0].(*glue.ListDevEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListDevEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDevEndpoints, reflect.TypeOf((*MockGlueClient)(nil).ListDevEndpoints), varargs...)
}

func (m *MockGlueClient) ListJobs(arg0 context.Context, arg1 *glue.ListJobsInput, arg2 ...func(*glue.Options)) (*glue.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobs, varargs...)
	ret0, _ := ret[0].(*glue.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobs, reflect.TypeOf((*MockGlueClient)(nil).ListJobs), varargs...)
}

func (m *MockGlueClient) ListMLTransforms(arg0 context.Context, arg1 *glue.ListMLTransformsInput, arg2 ...func(*glue.Options)) (*glue.ListMLTransformsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMLTransforms, varargs...)
	ret0, _ := ret[0].(*glue.ListMLTransformsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListMLTransforms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMLTransforms, reflect.TypeOf((*MockGlueClient)(nil).ListMLTransforms), varargs...)
}

func (m *MockGlueClient) ListRegistries(arg0 context.Context, arg1 *glue.ListRegistriesInput, arg2 ...func(*glue.Options)) (*glue.ListRegistriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegistries, varargs...)
	ret0, _ := ret[0].(*glue.ListRegistriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListRegistries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegistries, reflect.TypeOf((*MockGlueClient)(nil).ListRegistries), varargs...)
}

func (m *MockGlueClient) ListSchemaVersions(arg0 context.Context, arg1 *glue.ListSchemaVersionsInput, arg2 ...func(*glue.Options)) (*glue.ListSchemaVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSchemaVersions, varargs...)
	ret0, _ := ret[0].(*glue.ListSchemaVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListSchemaVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSchemaVersions, reflect.TypeOf((*MockGlueClient)(nil).ListSchemaVersions), varargs...)
}

func (m *MockGlueClient) ListSchemas(arg0 context.Context, arg1 *glue.ListSchemasInput, arg2 ...func(*glue.Options)) (*glue.ListSchemasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSchemas, varargs...)
	ret0, _ := ret[0].(*glue.ListSchemasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListSchemas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSchemas, reflect.TypeOf((*MockGlueClient)(nil).ListSchemas), varargs...)
}

func (m *MockGlueClient) ListSessions(arg0 context.Context, arg1 *glue.ListSessionsInput, arg2 ...func(*glue.Options)) (*glue.ListSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSessions, varargs...)
	ret0, _ := ret[0].(*glue.ListSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSessions, reflect.TypeOf((*MockGlueClient)(nil).ListSessions), varargs...)
}

func (m *MockGlueClient) ListStatements(arg0 context.Context, arg1 *glue.ListStatementsInput, arg2 ...func(*glue.Options)) (*glue.ListStatementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStatements, varargs...)
	ret0, _ := ret[0].(*glue.ListStatementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListStatements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStatements, reflect.TypeOf((*MockGlueClient)(nil).ListStatements), varargs...)
}

func (m *MockGlueClient) ListTriggers(arg0 context.Context, arg1 *glue.ListTriggersInput, arg2 ...func(*glue.Options)) (*glue.ListTriggersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTriggers, varargs...)
	ret0, _ := ret[0].(*glue.ListTriggersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListTriggers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTriggers, reflect.TypeOf((*MockGlueClient)(nil).ListTriggers), varargs...)
}

func (m *MockGlueClient) ListWorkflows(arg0 context.Context, arg1 *glue.ListWorkflowsInput, arg2 ...func(*glue.Options)) (*glue.ListWorkflowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWorkflows, varargs...)
	ret0, _ := ret[0].(*glue.ListWorkflowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWorkflows, reflect.TypeOf((*MockGlueClient)(nil).ListWorkflows), varargs...)
}

func (m *MockGlueClient) QuerySchemaVersionMetadata(arg0 context.Context, arg1 *glue.QuerySchemaVersionMetadataInput, arg2 ...func(*glue.Options)) (*glue.QuerySchemaVersionMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.QuerySchemaVersionMetadata, varargs...)
	ret0, _ := ret[0].(*glue.QuerySchemaVersionMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) QuerySchemaVersionMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.QuerySchemaVersionMetadata, reflect.TypeOf((*MockGlueClient)(nil).QuerySchemaVersionMetadata), varargs...)
}

func (m *MockGlueClient) SearchTables(arg0 context.Context, arg1 *glue.SearchTablesInput, arg2 ...func(*glue.Options)) (*glue.SearchTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchTables, varargs...)
	ret0, _ := ret[0].(*glue.SearchTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) SearchTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchTables, reflect.TypeOf((*MockGlueClient)(nil).SearchTables), varargs...)
}
