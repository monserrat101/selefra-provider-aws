package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	athena "github.com/aws/aws-sdk-go-v2/service/athena"
	gomock "github.com/golang/mock/gomock"
)

type MockAthenaClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAthenaClientMockRecorder
}

type MockAthenaClientMockRecorder struct {
	mock *MockAthenaClient
}

func NewMockAthenaClient(ctrl *gomock.Controller) *MockAthenaClient {
	mock := &MockAthenaClient{ctrl: ctrl}
	mock.recorder = &MockAthenaClientMockRecorder{mock}
	return mock
}

func (m *MockAthenaClient) EXPECT() *MockAthenaClientMockRecorder {
	return m.recorder
}

func (m *MockAthenaClient) BatchGetNamedQuery(arg0 context.Context, arg1 *athena.BatchGetNamedQueryInput, arg2 ...func(*athena.Options)) (*athena.BatchGetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetNamedQuery, varargs...)
	ret0, _ := ret[0].(*athena.BatchGetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) BatchGetNamedQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetNamedQuery, reflect.TypeOf((*MockAthenaClient)(nil).BatchGetNamedQuery), varargs...)
}

func (m *MockAthenaClient) BatchGetPreparedStatement(arg0 context.Context, arg1 *athena.BatchGetPreparedStatementInput, arg2 ...func(*athena.Options)) (*athena.BatchGetPreparedStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetPreparedStatement, varargs...)
	ret0, _ := ret[0].(*athena.BatchGetPreparedStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) BatchGetPreparedStatement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetPreparedStatement, reflect.TypeOf((*MockAthenaClient)(nil).BatchGetPreparedStatement), varargs...)
}

func (m *MockAthenaClient) BatchGetQueryExecution(arg0 context.Context, arg1 *athena.BatchGetQueryExecutionInput, arg2 ...func(*athena.Options)) (*athena.BatchGetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetQueryExecution, varargs...)
	ret0, _ := ret[0].(*athena.BatchGetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) BatchGetQueryExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetQueryExecution, reflect.TypeOf((*MockAthenaClient)(nil).BatchGetQueryExecution), varargs...)
}

func (m *MockAthenaClient) GetCalculationExecution(arg0 context.Context, arg1 *athena.GetCalculationExecutionInput, arg2 ...func(*athena.Options)) (*athena.GetCalculationExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCalculationExecution, varargs...)
	ret0, _ := ret[0].(*athena.GetCalculationExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetCalculationExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCalculationExecution, reflect.TypeOf((*MockAthenaClient)(nil).GetCalculationExecution), varargs...)
}

func (m *MockAthenaClient) GetCalculationExecutionCode(arg0 context.Context, arg1 *athena.GetCalculationExecutionCodeInput, arg2 ...func(*athena.Options)) (*athena.GetCalculationExecutionCodeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCalculationExecutionCode, varargs...)
	ret0, _ := ret[0].(*athena.GetCalculationExecutionCodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetCalculationExecutionCode(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCalculationExecutionCode, reflect.TypeOf((*MockAthenaClient)(nil).GetCalculationExecutionCode), varargs...)
}

func (m *MockAthenaClient) GetCalculationExecutionStatus(arg0 context.Context, arg1 *athena.GetCalculationExecutionStatusInput, arg2 ...func(*athena.Options)) (*athena.GetCalculationExecutionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCalculationExecutionStatus, varargs...)
	ret0, _ := ret[0].(*athena.GetCalculationExecutionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetCalculationExecutionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCalculationExecutionStatus, reflect.TypeOf((*MockAthenaClient)(nil).GetCalculationExecutionStatus), varargs...)
}

func (m *MockAthenaClient) GetDataCatalog(arg0 context.Context, arg1 *athena.GetDataCatalogInput, arg2 ...func(*athena.Options)) (*athena.GetDataCatalogOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataCatalog, varargs...)
	ret0, _ := ret[0].(*athena.GetDataCatalogOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetDataCatalog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataCatalog, reflect.TypeOf((*MockAthenaClient)(nil).GetDataCatalog), varargs...)
}

func (m *MockAthenaClient) GetDatabase(arg0 context.Context, arg1 *athena.GetDatabaseInput, arg2 ...func(*athena.Options)) (*athena.GetDatabaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDatabase, varargs...)
	ret0, _ := ret[0].(*athena.GetDatabaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDatabase, reflect.TypeOf((*MockAthenaClient)(nil).GetDatabase), varargs...)
}

func (m *MockAthenaClient) GetNamedQuery(arg0 context.Context, arg1 *athena.GetNamedQueryInput, arg2 ...func(*athena.Options)) (*athena.GetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetNamedQuery, varargs...)
	ret0, _ := ret[0].(*athena.GetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetNamedQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetNamedQuery, reflect.TypeOf((*MockAthenaClient)(nil).GetNamedQuery), varargs...)
}

func (m *MockAthenaClient) GetNotebookMetadata(arg0 context.Context, arg1 *athena.GetNotebookMetadataInput, arg2 ...func(*athena.Options)) (*athena.GetNotebookMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetNotebookMetadata, varargs...)
	ret0, _ := ret[0].(*athena.GetNotebookMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetNotebookMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetNotebookMetadata, reflect.TypeOf((*MockAthenaClient)(nil).GetNotebookMetadata), varargs...)
}

func (m *MockAthenaClient) GetPreparedStatement(arg0 context.Context, arg1 *athena.GetPreparedStatementInput, arg2 ...func(*athena.Options)) (*athena.GetPreparedStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPreparedStatement, varargs...)
	ret0, _ := ret[0].(*athena.GetPreparedStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetPreparedStatement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPreparedStatement, reflect.TypeOf((*MockAthenaClient)(nil).GetPreparedStatement), varargs...)
}

func (m *MockAthenaClient) GetQueryExecution(arg0 context.Context, arg1 *athena.GetQueryExecutionInput, arg2 ...func(*athena.Options)) (*athena.GetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueryExecution, varargs...)
	ret0, _ := ret[0].(*athena.GetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetQueryExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueryExecution, reflect.TypeOf((*MockAthenaClient)(nil).GetQueryExecution), varargs...)
}

func (m *MockAthenaClient) GetQueryResults(arg0 context.Context, arg1 *athena.GetQueryResultsInput, arg2 ...func(*athena.Options)) (*athena.GetQueryResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueryResults, varargs...)
	ret0, _ := ret[0].(*athena.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetQueryResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueryResults, reflect.TypeOf((*MockAthenaClient)(nil).GetQueryResults), varargs...)
}

func (m *MockAthenaClient) GetQueryRuntimeStatistics(arg0 context.Context, arg1 *athena.GetQueryRuntimeStatisticsInput, arg2 ...func(*athena.Options)) (*athena.GetQueryRuntimeStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueryRuntimeStatistics, varargs...)
	ret0, _ := ret[0].(*athena.GetQueryRuntimeStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetQueryRuntimeStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueryRuntimeStatistics, reflect.TypeOf((*MockAthenaClient)(nil).GetQueryRuntimeStatistics), varargs...)
}

func (m *MockAthenaClient) GetSession(arg0 context.Context, arg1 *athena.GetSessionInput, arg2 ...func(*athena.Options)) (*athena.GetSessionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSession, varargs...)
	ret0, _ := ret[0].(*athena.GetSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetSession(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSession, reflect.TypeOf((*MockAthenaClient)(nil).GetSession), varargs...)
}

func (m *MockAthenaClient) GetSessionStatus(arg0 context.Context, arg1 *athena.GetSessionStatusInput, arg2 ...func(*athena.Options)) (*athena.GetSessionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSessionStatus, varargs...)
	ret0, _ := ret[0].(*athena.GetSessionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetSessionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSessionStatus, reflect.TypeOf((*MockAthenaClient)(nil).GetSessionStatus), varargs...)
}

func (m *MockAthenaClient) GetTableMetadata(arg0 context.Context, arg1 *athena.GetTableMetadataInput, arg2 ...func(*athena.Options)) (*athena.GetTableMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTableMetadata, varargs...)
	ret0, _ := ret[0].(*athena.GetTableMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetTableMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTableMetadata, reflect.TypeOf((*MockAthenaClient)(nil).GetTableMetadata), varargs...)
}

func (m *MockAthenaClient) GetWorkGroup(arg0 context.Context, arg1 *athena.GetWorkGroupInput, arg2 ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetWorkGroup, varargs...)
	ret0, _ := ret[0].(*athena.GetWorkGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetWorkGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetWorkGroup, reflect.TypeOf((*MockAthenaClient)(nil).GetWorkGroup), varargs...)
}

func (m *MockAthenaClient) ListApplicationDPUSizes(arg0 context.Context, arg1 *athena.ListApplicationDPUSizesInput, arg2 ...func(*athena.Options)) (*athena.ListApplicationDPUSizesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListApplicationDPUSizes, varargs...)
	ret0, _ := ret[0].(*athena.ListApplicationDPUSizesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListApplicationDPUSizes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListApplicationDPUSizes, reflect.TypeOf((*MockAthenaClient)(nil).ListApplicationDPUSizes), varargs...)
}

func (m *MockAthenaClient) ListCalculationExecutions(arg0 context.Context, arg1 *athena.ListCalculationExecutionsInput, arg2 ...func(*athena.Options)) (*athena.ListCalculationExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCalculationExecutions, varargs...)
	ret0, _ := ret[0].(*athena.ListCalculationExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListCalculationExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCalculationExecutions, reflect.TypeOf((*MockAthenaClient)(nil).ListCalculationExecutions), varargs...)
}

func (m *MockAthenaClient) ListDataCatalogs(arg0 context.Context, arg1 *athena.ListDataCatalogsInput, arg2 ...func(*athena.Options)) (*athena.ListDataCatalogsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataCatalogs, varargs...)
	ret0, _ := ret[0].(*athena.ListDataCatalogsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListDataCatalogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataCatalogs, reflect.TypeOf((*MockAthenaClient)(nil).ListDataCatalogs), varargs...)
}

func (m *MockAthenaClient) ListDatabases(arg0 context.Context, arg1 *athena.ListDatabasesInput, arg2 ...func(*athena.Options)) (*athena.ListDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDatabases, varargs...)
	ret0, _ := ret[0].(*athena.ListDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDatabases, reflect.TypeOf((*MockAthenaClient)(nil).ListDatabases), varargs...)
}

func (m *MockAthenaClient) ListEngineVersions(arg0 context.Context, arg1 *athena.ListEngineVersionsInput, arg2 ...func(*athena.Options)) (*athena.ListEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEngineVersions, varargs...)
	ret0, _ := ret[0].(*athena.ListEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEngineVersions, reflect.TypeOf((*MockAthenaClient)(nil).ListEngineVersions), varargs...)
}

func (m *MockAthenaClient) ListExecutors(arg0 context.Context, arg1 *athena.ListExecutorsInput, arg2 ...func(*athena.Options)) (*athena.ListExecutorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExecutors, varargs...)
	ret0, _ := ret[0].(*athena.ListExecutorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListExecutors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExecutors, reflect.TypeOf((*MockAthenaClient)(nil).ListExecutors), varargs...)
}

func (m *MockAthenaClient) ListNamedQueries(arg0 context.Context, arg1 *athena.ListNamedQueriesInput, arg2 ...func(*athena.Options)) (*athena.ListNamedQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNamedQueries, varargs...)
	ret0, _ := ret[0].(*athena.ListNamedQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListNamedQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNamedQueries, reflect.TypeOf((*MockAthenaClient)(nil).ListNamedQueries), varargs...)
}

func (m *MockAthenaClient) ListNotebookMetadata(arg0 context.Context, arg1 *athena.ListNotebookMetadataInput, arg2 ...func(*athena.Options)) (*athena.ListNotebookMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNotebookMetadata, varargs...)
	ret0, _ := ret[0].(*athena.ListNotebookMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListNotebookMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNotebookMetadata, reflect.TypeOf((*MockAthenaClient)(nil).ListNotebookMetadata), varargs...)
}

func (m *MockAthenaClient) ListNotebookSessions(arg0 context.Context, arg1 *athena.ListNotebookSessionsInput, arg2 ...func(*athena.Options)) (*athena.ListNotebookSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNotebookSessions, varargs...)
	ret0, _ := ret[0].(*athena.ListNotebookSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListNotebookSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNotebookSessions, reflect.TypeOf((*MockAthenaClient)(nil).ListNotebookSessions), varargs...)
}

func (m *MockAthenaClient) ListPreparedStatements(arg0 context.Context, arg1 *athena.ListPreparedStatementsInput, arg2 ...func(*athena.Options)) (*athena.ListPreparedStatementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPreparedStatements, varargs...)
	ret0, _ := ret[0].(*athena.ListPreparedStatementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListPreparedStatements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPreparedStatements, reflect.TypeOf((*MockAthenaClient)(nil).ListPreparedStatements), varargs...)
}

func (m *MockAthenaClient) ListQueryExecutions(arg0 context.Context, arg1 *athena.ListQueryExecutionsInput, arg2 ...func(*athena.Options)) (*athena.ListQueryExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListQueryExecutions, varargs...)
	ret0, _ := ret[0].(*athena.ListQueryExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListQueryExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListQueryExecutions, reflect.TypeOf((*MockAthenaClient)(nil).ListQueryExecutions), varargs...)
}

func (m *MockAthenaClient) ListSessions(arg0 context.Context, arg1 *athena.ListSessionsInput, arg2 ...func(*athena.Options)) (*athena.ListSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSessions, varargs...)
	ret0, _ := ret[0].(*athena.ListSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSessions, reflect.TypeOf((*MockAthenaClient)(nil).ListSessions), varargs...)
}

func (m *MockAthenaClient) ListTableMetadata(arg0 context.Context, arg1 *athena.ListTableMetadataInput, arg2 ...func(*athena.Options)) (*athena.ListTableMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTableMetadata, varargs...)
	ret0, _ := ret[0].(*athena.ListTableMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListTableMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTableMetadata, reflect.TypeOf((*MockAthenaClient)(nil).ListTableMetadata), varargs...)
}

func (m *MockAthenaClient) ListTagsForResource(arg0 context.Context, arg1 *athena.ListTagsForResourceInput, arg2 ...func(*athena.Options)) (*athena.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*athena.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockAthenaClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockAthenaClient) ListWorkGroups(arg0 context.Context, arg1 *athena.ListWorkGroupsInput, arg2 ...func(*athena.Options)) (*athena.ListWorkGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWorkGroups, varargs...)
	ret0, _ := ret[0].(*athena.ListWorkGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListWorkGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWorkGroups, reflect.TypeOf((*MockAthenaClient)(nil).ListWorkGroups), varargs...)
}
