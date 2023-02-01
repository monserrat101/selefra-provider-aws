package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	emr "github.com/aws/aws-sdk-go-v2/service/emr"
	gomock "github.com/golang/mock/gomock"
)

type MockEmrClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEmrClientMockRecorder
}

type MockEmrClientMockRecorder struct {
	mock *MockEmrClient
}

func NewMockEmrClient(ctrl *gomock.Controller) *MockEmrClient {
	mock := &MockEmrClient{ctrl: ctrl}
	mock.recorder = &MockEmrClientMockRecorder{mock}
	return mock
}

func (m *MockEmrClient) EXPECT() *MockEmrClientMockRecorder {
	return m.recorder
}

func (m *MockEmrClient) DescribeCluster(arg0 context.Context, arg1 *emr.DescribeClusterInput, arg2 ...func(*emr.Options)) (*emr.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCluster, varargs...)
	ret0, _ := ret[0].(*emr.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCluster, reflect.TypeOf((*MockEmrClient)(nil).DescribeCluster), varargs...)
}

func (m *MockEmrClient) DescribeJobFlows(arg0 context.Context, arg1 *emr.DescribeJobFlowsInput, arg2 ...func(*emr.Options)) (*emr.DescribeJobFlowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJobFlows, varargs...)
	ret0, _ := ret[0].(*emr.DescribeJobFlowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeJobFlows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJobFlows, reflect.TypeOf((*MockEmrClient)(nil).DescribeJobFlows), varargs...)
}

func (m *MockEmrClient) DescribeNotebookExecution(arg0 context.Context, arg1 *emr.DescribeNotebookExecutionInput, arg2 ...func(*emr.Options)) (*emr.DescribeNotebookExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNotebookExecution, varargs...)
	ret0, _ := ret[0].(*emr.DescribeNotebookExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeNotebookExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNotebookExecution, reflect.TypeOf((*MockEmrClient)(nil).DescribeNotebookExecution), varargs...)
}

func (m *MockEmrClient) DescribeReleaseLabel(arg0 context.Context, arg1 *emr.DescribeReleaseLabelInput, arg2 ...func(*emr.Options)) (*emr.DescribeReleaseLabelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReleaseLabel, varargs...)
	ret0, _ := ret[0].(*emr.DescribeReleaseLabelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeReleaseLabel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReleaseLabel, reflect.TypeOf((*MockEmrClient)(nil).DescribeReleaseLabel), varargs...)
}

func (m *MockEmrClient) DescribeSecurityConfiguration(arg0 context.Context, arg1 *emr.DescribeSecurityConfigurationInput, arg2 ...func(*emr.Options)) (*emr.DescribeSecurityConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecurityConfiguration, varargs...)
	ret0, _ := ret[0].(*emr.DescribeSecurityConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeSecurityConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecurityConfiguration, reflect.TypeOf((*MockEmrClient)(nil).DescribeSecurityConfiguration), varargs...)
}

func (m *MockEmrClient) DescribeStep(arg0 context.Context, arg1 *emr.DescribeStepInput, arg2 ...func(*emr.Options)) (*emr.DescribeStepOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStep, varargs...)
	ret0, _ := ret[0].(*emr.DescribeStepOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeStep(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStep, reflect.TypeOf((*MockEmrClient)(nil).DescribeStep), varargs...)
}

func (m *MockEmrClient) DescribeStudio(arg0 context.Context, arg1 *emr.DescribeStudioInput, arg2 ...func(*emr.Options)) (*emr.DescribeStudioOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStudio, varargs...)
	ret0, _ := ret[0].(*emr.DescribeStudioOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeStudio(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStudio, reflect.TypeOf((*MockEmrClient)(nil).DescribeStudio), varargs...)
}

func (m *MockEmrClient) GetAutoTerminationPolicy(arg0 context.Context, arg1 *emr.GetAutoTerminationPolicyInput, arg2 ...func(*emr.Options)) (*emr.GetAutoTerminationPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAutoTerminationPolicy, varargs...)
	ret0, _ := ret[0].(*emr.GetAutoTerminationPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) GetAutoTerminationPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAutoTerminationPolicy, reflect.TypeOf((*MockEmrClient)(nil).GetAutoTerminationPolicy), varargs...)
}

func (m *MockEmrClient) GetBlockPublicAccessConfiguration(arg0 context.Context, arg1 *emr.GetBlockPublicAccessConfigurationInput, arg2 ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlockPublicAccessConfiguration, varargs...)
	ret0, _ := ret[0].(*emr.GetBlockPublicAccessConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) GetBlockPublicAccessConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlockPublicAccessConfiguration, reflect.TypeOf((*MockEmrClient)(nil).GetBlockPublicAccessConfiguration), varargs...)
}

func (m *MockEmrClient) GetClusterSessionCredentials(arg0 context.Context, arg1 *emr.GetClusterSessionCredentialsInput, arg2 ...func(*emr.Options)) (*emr.GetClusterSessionCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClusterSessionCredentials, varargs...)
	ret0, _ := ret[0].(*emr.GetClusterSessionCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) GetClusterSessionCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClusterSessionCredentials, reflect.TypeOf((*MockEmrClient)(nil).GetClusterSessionCredentials), varargs...)
}

func (m *MockEmrClient) GetManagedScalingPolicy(arg0 context.Context, arg1 *emr.GetManagedScalingPolicyInput, arg2 ...func(*emr.Options)) (*emr.GetManagedScalingPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetManagedScalingPolicy, varargs...)
	ret0, _ := ret[0].(*emr.GetManagedScalingPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) GetManagedScalingPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetManagedScalingPolicy, reflect.TypeOf((*MockEmrClient)(nil).GetManagedScalingPolicy), varargs...)
}

func (m *MockEmrClient) GetStudioSessionMapping(arg0 context.Context, arg1 *emr.GetStudioSessionMappingInput, arg2 ...func(*emr.Options)) (*emr.GetStudioSessionMappingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStudioSessionMapping, varargs...)
	ret0, _ := ret[0].(*emr.GetStudioSessionMappingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) GetStudioSessionMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStudioSessionMapping, reflect.TypeOf((*MockEmrClient)(nil).GetStudioSessionMapping), varargs...)
}

func (m *MockEmrClient) ListBootstrapActions(arg0 context.Context, arg1 *emr.ListBootstrapActionsInput, arg2 ...func(*emr.Options)) (*emr.ListBootstrapActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBootstrapActions, varargs...)
	ret0, _ := ret[0].(*emr.ListBootstrapActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListBootstrapActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBootstrapActions, reflect.TypeOf((*MockEmrClient)(nil).ListBootstrapActions), varargs...)
}

func (m *MockEmrClient) ListClusters(arg0 context.Context, arg1 *emr.ListClustersInput, arg2 ...func(*emr.Options)) (*emr.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListClusters, varargs...)
	ret0, _ := ret[0].(*emr.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListClusters, reflect.TypeOf((*MockEmrClient)(nil).ListClusters), varargs...)
}

func (m *MockEmrClient) ListInstanceFleets(arg0 context.Context, arg1 *emr.ListInstanceFleetsInput, arg2 ...func(*emr.Options)) (*emr.ListInstanceFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstanceFleets, varargs...)
	ret0, _ := ret[0].(*emr.ListInstanceFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListInstanceFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstanceFleets, reflect.TypeOf((*MockEmrClient)(nil).ListInstanceFleets), varargs...)
}

func (m *MockEmrClient) ListInstanceGroups(arg0 context.Context, arg1 *emr.ListInstanceGroupsInput, arg2 ...func(*emr.Options)) (*emr.ListInstanceGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstanceGroups, varargs...)
	ret0, _ := ret[0].(*emr.ListInstanceGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListInstanceGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstanceGroups, reflect.TypeOf((*MockEmrClient)(nil).ListInstanceGroups), varargs...)
}

func (m *MockEmrClient) ListInstances(arg0 context.Context, arg1 *emr.ListInstancesInput, arg2 ...func(*emr.Options)) (*emr.ListInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInstances, varargs...)
	ret0, _ := ret[0].(*emr.ListInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstances, reflect.TypeOf((*MockEmrClient)(nil).ListInstances), varargs...)
}

func (m *MockEmrClient) ListNotebookExecutions(arg0 context.Context, arg1 *emr.ListNotebookExecutionsInput, arg2 ...func(*emr.Options)) (*emr.ListNotebookExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNotebookExecutions, varargs...)
	ret0, _ := ret[0].(*emr.ListNotebookExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListNotebookExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNotebookExecutions, reflect.TypeOf((*MockEmrClient)(nil).ListNotebookExecutions), varargs...)
}

func (m *MockEmrClient) ListReleaseLabels(arg0 context.Context, arg1 *emr.ListReleaseLabelsInput, arg2 ...func(*emr.Options)) (*emr.ListReleaseLabelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReleaseLabels, varargs...)
	ret0, _ := ret[0].(*emr.ListReleaseLabelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListReleaseLabels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReleaseLabels, reflect.TypeOf((*MockEmrClient)(nil).ListReleaseLabels), varargs...)
}

func (m *MockEmrClient) ListSecurityConfigurations(arg0 context.Context, arg1 *emr.ListSecurityConfigurationsInput, arg2 ...func(*emr.Options)) (*emr.ListSecurityConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSecurityConfigurations, varargs...)
	ret0, _ := ret[0].(*emr.ListSecurityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListSecurityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSecurityConfigurations, reflect.TypeOf((*MockEmrClient)(nil).ListSecurityConfigurations), varargs...)
}

func (m *MockEmrClient) ListSteps(arg0 context.Context, arg1 *emr.ListStepsInput, arg2 ...func(*emr.Options)) (*emr.ListStepsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSteps, varargs...)
	ret0, _ := ret[0].(*emr.ListStepsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListSteps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSteps, reflect.TypeOf((*MockEmrClient)(nil).ListSteps), varargs...)
}

func (m *MockEmrClient) ListStudioSessionMappings(arg0 context.Context, arg1 *emr.ListStudioSessionMappingsInput, arg2 ...func(*emr.Options)) (*emr.ListStudioSessionMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStudioSessionMappings, varargs...)
	ret0, _ := ret[0].(*emr.ListStudioSessionMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListStudioSessionMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStudioSessionMappings, reflect.TypeOf((*MockEmrClient)(nil).ListStudioSessionMappings), varargs...)
}

func (m *MockEmrClient) ListStudios(arg0 context.Context, arg1 *emr.ListStudiosInput, arg2 ...func(*emr.Options)) (*emr.ListStudiosOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStudios, varargs...)
	ret0, _ := ret[0].(*emr.ListStudiosOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListStudios(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStudios, reflect.TypeOf((*MockEmrClient)(nil).ListStudios), varargs...)
}
