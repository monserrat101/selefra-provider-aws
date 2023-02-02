package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	databasemigrationservice "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	gomock "github.com/golang/mock/gomock"
)

type MockDatabasemigrationserviceClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDatabasemigrationserviceClientMockRecorder
}

type MockDatabasemigrationserviceClientMockRecorder struct {
	mock *MockDatabasemigrationserviceClient
}

func NewMockDatabasemigrationserviceClient(ctrl *gomock.Controller) *MockDatabasemigrationserviceClient {
	mock := &MockDatabasemigrationserviceClient{ctrl: ctrl}
	mock.recorder = &MockDatabasemigrationserviceClientMockRecorder{mock}
	return mock
}

func (m *MockDatabasemigrationserviceClient) EXPECT() *MockDatabasemigrationserviceClientMockRecorder {
	return m.recorder
}

func (m *MockDatabasemigrationserviceClient) DescribeAccountAttributes(arg0 context.Context, arg1 *databasemigrationservice.DescribeAccountAttributesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAttributes, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeAccountAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAttributes, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeAccountAttributes), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeApplicableIndividualAssessments(arg0 context.Context, arg1 *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApplicableIndividualAssessments, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeApplicableIndividualAssessments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApplicableIndividualAssessments, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeApplicableIndividualAssessments), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeCertificates(arg0 context.Context, arg1 *databasemigrationservice.DescribeCertificatesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCertificates, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCertificates, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeCertificates), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeConnections(arg0 context.Context, arg1 *databasemigrationservice.DescribeConnectionsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnections, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnections, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeConnections), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeEndpointSettings(arg0 context.Context, arg1 *databasemigrationservice.DescribeEndpointSettingsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpointSettings, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeEndpointSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeEndpointSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpointSettings, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeEndpointSettings), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeEndpointTypes(arg0 context.Context, arg1 *databasemigrationservice.DescribeEndpointTypesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpointTypes, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeEndpointTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeEndpointTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpointTypes, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeEndpointTypes), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeEndpoints(arg0 context.Context, arg1 *databasemigrationservice.DescribeEndpointsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpoints, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpoints, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeEndpoints), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeEventCategories(arg0 context.Context, arg1 *databasemigrationservice.DescribeEventCategoriesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventCategories, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeEventCategoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeEventCategories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventCategories, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeEventCategories), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *databasemigrationservice.DescribeEventSubscriptionsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventSubscriptions, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventSubscriptions, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeEvents(arg0 context.Context, arg1 *databasemigrationservice.DescribeEventsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeEvents), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeFleetAdvisorCollectors(arg0 context.Context, arg1 *databasemigrationservice.DescribeFleetAdvisorCollectorsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetAdvisorCollectors, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeFleetAdvisorCollectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetAdvisorCollectors, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeFleetAdvisorCollectors), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeFleetAdvisorDatabases(arg0 context.Context, arg1 *databasemigrationservice.DescribeFleetAdvisorDatabasesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetAdvisorDatabases, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeFleetAdvisorDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetAdvisorDatabases, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeFleetAdvisorDatabases), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeFleetAdvisorLsaAnalysis(arg0 context.Context, arg1 *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetAdvisorLsaAnalysis, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeFleetAdvisorLsaAnalysis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetAdvisorLsaAnalysis, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeFleetAdvisorLsaAnalysis), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeFleetAdvisorSchemaObjectSummary(arg0 context.Context, arg1 *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetAdvisorSchemaObjectSummary, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeFleetAdvisorSchemaObjectSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetAdvisorSchemaObjectSummary, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeFleetAdvisorSchemaObjectSummary), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeFleetAdvisorSchemas(arg0 context.Context, arg1 *databasemigrationservice.DescribeFleetAdvisorSchemasInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetAdvisorSchemas, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeFleetAdvisorSchemasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeFleetAdvisorSchemas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetAdvisorSchemas, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeFleetAdvisorSchemas), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeOrderableReplicationInstances(arg0 context.Context, arg1 *databasemigrationservice.DescribeOrderableReplicationInstancesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrderableReplicationInstances, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeOrderableReplicationInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrderableReplicationInstances, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeOrderableReplicationInstances), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribePendingMaintenanceActions(arg0 context.Context, arg1 *databasemigrationservice.DescribePendingMaintenanceActionsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePendingMaintenanceActions, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribePendingMaintenanceActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribePendingMaintenanceActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePendingMaintenanceActions, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribePendingMaintenanceActions), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeRefreshSchemasStatus(arg0 context.Context, arg1 *databasemigrationservice.DescribeRefreshSchemasStatusInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRefreshSchemasStatus, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeRefreshSchemasStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeRefreshSchemasStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRefreshSchemasStatus, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeRefreshSchemasStatus), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationInstanceTaskLogs(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationInstanceTaskLogs, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationInstanceTaskLogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationInstanceTaskLogs, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationInstanceTaskLogs), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationInstances(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationInstancesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationInstances, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationInstances, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationInstances), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationSubnetGroups(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationSubnetGroupsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationSubnetGroups, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationSubnetGroups, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationSubnetGroups), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationTaskAssessmentResults(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationTaskAssessmentResults, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationTaskAssessmentResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationTaskAssessmentResults, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationTaskAssessmentResults), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationTaskAssessmentRuns(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationTaskAssessmentRuns, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationTaskAssessmentRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationTaskAssessmentRuns, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationTaskAssessmentRuns), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationTaskIndividualAssessments(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationTaskIndividualAssessments, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationTaskIndividualAssessments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationTaskIndividualAssessments, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationTaskIndividualAssessments), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationTasks(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationTasksInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationTasks, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationTasks, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationTasks), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeSchemas(arg0 context.Context, arg1 *databasemigrationservice.DescribeSchemasInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeSchemasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSchemas, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeSchemasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeSchemas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSchemas, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeSchemas), varargs...)
}

func (m *MockDatabasemigrationserviceClient) DescribeTableStatistics(arg0 context.Context, arg1 *databasemigrationservice.DescribeTableStatisticsInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTableStatistics, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeTableStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeTableStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTableStatistics, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeTableStatistics), varargs...)
}

func (m *MockDatabasemigrationserviceClient) ListTagsForResource(arg0 context.Context, arg1 *databasemigrationservice.ListTagsForResourceInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).ListTagsForResource), varargs...)
}
