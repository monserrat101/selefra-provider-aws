package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	codebuild "github.com/aws/aws-sdk-go-v2/service/codebuild"
	gomock "github.com/golang/mock/gomock"
)

type MockCodebuildClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCodebuildClientMockRecorder
}

type MockCodebuildClientMockRecorder struct {
	mock *MockCodebuildClient
}

func NewMockCodebuildClient(ctrl *gomock.Controller) *MockCodebuildClient {
	mock := &MockCodebuildClient{ctrl: ctrl}
	mock.recorder = &MockCodebuildClientMockRecorder{mock}
	return mock
}

func (m *MockCodebuildClient) EXPECT() *MockCodebuildClientMockRecorder {
	return m.recorder
}

func (m *MockCodebuildClient) BatchGetBuildBatches(arg0 context.Context, arg1 *codebuild.BatchGetBuildBatchesInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetBuildBatchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetBuildBatches, varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetBuildBatchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) BatchGetBuildBatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetBuildBatches, reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetBuildBatches), varargs...)
}

func (m *MockCodebuildClient) BatchGetBuilds(arg0 context.Context, arg1 *codebuild.BatchGetBuildsInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetBuildsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetBuilds, varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetBuildsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) BatchGetBuilds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetBuilds, reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetBuilds), varargs...)
}

func (m *MockCodebuildClient) BatchGetProjects(arg0 context.Context, arg1 *codebuild.BatchGetProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetProjects, varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) BatchGetProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetProjects, reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetProjects), varargs...)
}

func (m *MockCodebuildClient) BatchGetReportGroups(arg0 context.Context, arg1 *codebuild.BatchGetReportGroupsInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetReportGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetReportGroups, varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetReportGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) BatchGetReportGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetReportGroups, reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetReportGroups), varargs...)
}

func (m *MockCodebuildClient) BatchGetReports(arg0 context.Context, arg1 *codebuild.BatchGetReportsInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetReportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetReports, varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) BatchGetReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetReports, reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetReports), varargs...)
}

func (m *MockCodebuildClient) DescribeCodeCoverages(arg0 context.Context, arg1 *codebuild.DescribeCodeCoveragesInput, arg2 ...func(*codebuild.Options)) (*codebuild.DescribeCodeCoveragesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCodeCoverages, varargs...)
	ret0, _ := ret[0].(*codebuild.DescribeCodeCoveragesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) DescribeCodeCoverages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCodeCoverages, reflect.TypeOf((*MockCodebuildClient)(nil).DescribeCodeCoverages), varargs...)
}

func (m *MockCodebuildClient) DescribeTestCases(arg0 context.Context, arg1 *codebuild.DescribeTestCasesInput, arg2 ...func(*codebuild.Options)) (*codebuild.DescribeTestCasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTestCases, varargs...)
	ret0, _ := ret[0].(*codebuild.DescribeTestCasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) DescribeTestCases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTestCases, reflect.TypeOf((*MockCodebuildClient)(nil).DescribeTestCases), varargs...)
}

func (m *MockCodebuildClient) GetReportGroupTrend(arg0 context.Context, arg1 *codebuild.GetReportGroupTrendInput, arg2 ...func(*codebuild.Options)) (*codebuild.GetReportGroupTrendOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetReportGroupTrend, varargs...)
	ret0, _ := ret[0].(*codebuild.GetReportGroupTrendOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) GetReportGroupTrend(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetReportGroupTrend, reflect.TypeOf((*MockCodebuildClient)(nil).GetReportGroupTrend), varargs...)
}

func (m *MockCodebuildClient) GetResourcePolicy(arg0 context.Context, arg1 *codebuild.GetResourcePolicyInput, arg2 ...func(*codebuild.Options)) (*codebuild.GetResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourcePolicy, varargs...)
	ret0, _ := ret[0].(*codebuild.GetResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) GetResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourcePolicy, reflect.TypeOf((*MockCodebuildClient)(nil).GetResourcePolicy), varargs...)
}

func (m *MockCodebuildClient) ListBuildBatches(arg0 context.Context, arg1 *codebuild.ListBuildBatchesInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildBatchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBuildBatches, varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildBatchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListBuildBatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBuildBatches, reflect.TypeOf((*MockCodebuildClient)(nil).ListBuildBatches), varargs...)
}

func (m *MockCodebuildClient) ListBuildBatchesForProject(arg0 context.Context, arg1 *codebuild.ListBuildBatchesForProjectInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBuildBatchesForProject, varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildBatchesForProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListBuildBatchesForProject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBuildBatchesForProject, reflect.TypeOf((*MockCodebuildClient)(nil).ListBuildBatchesForProject), varargs...)
}

func (m *MockCodebuildClient) ListBuilds(arg0 context.Context, arg1 *codebuild.ListBuildsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBuilds, varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListBuilds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBuilds, reflect.TypeOf((*MockCodebuildClient)(nil).ListBuilds), varargs...)
}

func (m *MockCodebuildClient) ListBuildsForProject(arg0 context.Context, arg1 *codebuild.ListBuildsForProjectInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildsForProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBuildsForProject, varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildsForProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListBuildsForProject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBuildsForProject, reflect.TypeOf((*MockCodebuildClient)(nil).ListBuildsForProject), varargs...)
}

func (m *MockCodebuildClient) ListCuratedEnvironmentImages(arg0 context.Context, arg1 *codebuild.ListCuratedEnvironmentImagesInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCuratedEnvironmentImages, varargs...)
	ret0, _ := ret[0].(*codebuild.ListCuratedEnvironmentImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListCuratedEnvironmentImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCuratedEnvironmentImages, reflect.TypeOf((*MockCodebuildClient)(nil).ListCuratedEnvironmentImages), varargs...)
}

func (m *MockCodebuildClient) ListProjects(arg0 context.Context, arg1 *codebuild.ListProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProjects, varargs...)
	ret0, _ := ret[0].(*codebuild.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProjects, reflect.TypeOf((*MockCodebuildClient)(nil).ListProjects), varargs...)
}

func (m *MockCodebuildClient) ListReportGroups(arg0 context.Context, arg1 *codebuild.ListReportGroupsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListReportGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReportGroups, varargs...)
	ret0, _ := ret[0].(*codebuild.ListReportGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListReportGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReportGroups, reflect.TypeOf((*MockCodebuildClient)(nil).ListReportGroups), varargs...)
}

func (m *MockCodebuildClient) ListReports(arg0 context.Context, arg1 *codebuild.ListReportsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListReportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReports, varargs...)
	ret0, _ := ret[0].(*codebuild.ListReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReports, reflect.TypeOf((*MockCodebuildClient)(nil).ListReports), varargs...)
}

func (m *MockCodebuildClient) ListReportsForReportGroup(arg0 context.Context, arg1 *codebuild.ListReportsForReportGroupInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListReportsForReportGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReportsForReportGroup, varargs...)
	ret0, _ := ret[0].(*codebuild.ListReportsForReportGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListReportsForReportGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReportsForReportGroup, reflect.TypeOf((*MockCodebuildClient)(nil).ListReportsForReportGroup), varargs...)
}

func (m *MockCodebuildClient) ListSharedProjects(arg0 context.Context, arg1 *codebuild.ListSharedProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListSharedProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSharedProjects, varargs...)
	ret0, _ := ret[0].(*codebuild.ListSharedProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListSharedProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSharedProjects, reflect.TypeOf((*MockCodebuildClient)(nil).ListSharedProjects), varargs...)
}

func (m *MockCodebuildClient) ListSharedReportGroups(arg0 context.Context, arg1 *codebuild.ListSharedReportGroupsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListSharedReportGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSharedReportGroups, varargs...)
	ret0, _ := ret[0].(*codebuild.ListSharedReportGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListSharedReportGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSharedReportGroups, reflect.TypeOf((*MockCodebuildClient)(nil).ListSharedReportGroups), varargs...)
}

func (m *MockCodebuildClient) ListSourceCredentials(arg0 context.Context, arg1 *codebuild.ListSourceCredentialsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListSourceCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSourceCredentials, varargs...)
	ret0, _ := ret[0].(*codebuild.ListSourceCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListSourceCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSourceCredentials, reflect.TypeOf((*MockCodebuildClient)(nil).ListSourceCredentials), varargs...)
}
