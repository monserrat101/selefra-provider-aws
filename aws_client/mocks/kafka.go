package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	kafka "github.com/aws/aws-sdk-go-v2/service/kafka"
	gomock "github.com/golang/mock/gomock"
)

type MockKafkaClient struct {
	ctrl		*gomock.Controller
	recorder	*MockKafkaClientMockRecorder
}

type MockKafkaClientMockRecorder struct {
	mock *MockKafkaClient
}

func NewMockKafkaClient(ctrl *gomock.Controller) *MockKafkaClient {
	mock := &MockKafkaClient{ctrl: ctrl}
	mock.recorder = &MockKafkaClientMockRecorder{mock}
	return mock
}

func (m *MockKafkaClient) EXPECT() *MockKafkaClientMockRecorder {
	return m.recorder
}

func (m *MockKafkaClient) DescribeCluster(arg0 context.Context, arg1 *kafka.DescribeClusterInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCluster, varargs...)
	ret0, _ := ret[0].(*kafka.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCluster, reflect.TypeOf((*MockKafkaClient)(nil).DescribeCluster), varargs...)
}

func (m *MockKafkaClient) DescribeClusterOperation(arg0 context.Context, arg1 *kafka.DescribeClusterOperationInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeClusterOperationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterOperation, varargs...)
	ret0, _ := ret[0].(*kafka.DescribeClusterOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) DescribeClusterOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterOperation, reflect.TypeOf((*MockKafkaClient)(nil).DescribeClusterOperation), varargs...)
}

func (m *MockKafkaClient) DescribeClusterV2(arg0 context.Context, arg1 *kafka.DescribeClusterV2Input, arg2 ...func(*kafka.Options)) (*kafka.DescribeClusterV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterV, varargs...)
	ret0, _ := ret[0].(*kafka.DescribeClusterV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) DescribeClusterV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterV, reflect.TypeOf((*MockKafkaClient)(nil).DescribeClusterV2), varargs...)
}

func (m *MockKafkaClient) DescribeConfiguration(arg0 context.Context, arg1 *kafka.DescribeConfigurationInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfiguration, varargs...)
	ret0, _ := ret[0].(*kafka.DescribeConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) DescribeConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfiguration, reflect.TypeOf((*MockKafkaClient)(nil).DescribeConfiguration), varargs...)
}

func (m *MockKafkaClient) DescribeConfigurationRevision(arg0 context.Context, arg1 *kafka.DescribeConfigurationRevisionInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeConfigurationRevisionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConfigurationRevision, varargs...)
	ret0, _ := ret[0].(*kafka.DescribeConfigurationRevisionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) DescribeConfigurationRevision(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConfigurationRevision, reflect.TypeOf((*MockKafkaClient)(nil).DescribeConfigurationRevision), varargs...)
}

func (m *MockKafkaClient) GetBootstrapBrokers(arg0 context.Context, arg1 *kafka.GetBootstrapBrokersInput, arg2 ...func(*kafka.Options)) (*kafka.GetBootstrapBrokersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBootstrapBrokers, varargs...)
	ret0, _ := ret[0].(*kafka.GetBootstrapBrokersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) GetBootstrapBrokers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBootstrapBrokers, reflect.TypeOf((*MockKafkaClient)(nil).GetBootstrapBrokers), varargs...)
}

func (m *MockKafkaClient) GetCompatibleKafkaVersions(arg0 context.Context, arg1 *kafka.GetCompatibleKafkaVersionsInput, arg2 ...func(*kafka.Options)) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCompatibleKafkaVersions, varargs...)
	ret0, _ := ret[0].(*kafka.GetCompatibleKafkaVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) GetCompatibleKafkaVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCompatibleKafkaVersions, reflect.TypeOf((*MockKafkaClient)(nil).GetCompatibleKafkaVersions), varargs...)
}

func (m *MockKafkaClient) ListClusterOperations(arg0 context.Context, arg1 *kafka.ListClusterOperationsInput, arg2 ...func(*kafka.Options)) (*kafka.ListClusterOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListClusterOperations, varargs...)
	ret0, _ := ret[0].(*kafka.ListClusterOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListClusterOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListClusterOperations, reflect.TypeOf((*MockKafkaClient)(nil).ListClusterOperations), varargs...)
}

func (m *MockKafkaClient) ListClusters(arg0 context.Context, arg1 *kafka.ListClustersInput, arg2 ...func(*kafka.Options)) (*kafka.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListClusters, varargs...)
	ret0, _ := ret[0].(*kafka.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListClusters, reflect.TypeOf((*MockKafkaClient)(nil).ListClusters), varargs...)
}

func (m *MockKafkaClient) ListClustersV2(arg0 context.Context, arg1 *kafka.ListClustersV2Input, arg2 ...func(*kafka.Options)) (*kafka.ListClustersV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListClustersV, varargs...)
	ret0, _ := ret[0].(*kafka.ListClustersV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListClustersV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListClustersV, reflect.TypeOf((*MockKafkaClient)(nil).ListClustersV2), varargs...)
}

func (m *MockKafkaClient) ListConfigurationRevisions(arg0 context.Context, arg1 *kafka.ListConfigurationRevisionsInput, arg2 ...func(*kafka.Options)) (*kafka.ListConfigurationRevisionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConfigurationRevisions, varargs...)
	ret0, _ := ret[0].(*kafka.ListConfigurationRevisionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListConfigurationRevisions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConfigurationRevisions, reflect.TypeOf((*MockKafkaClient)(nil).ListConfigurationRevisions), varargs...)
}

func (m *MockKafkaClient) ListConfigurations(arg0 context.Context, arg1 *kafka.ListConfigurationsInput, arg2 ...func(*kafka.Options)) (*kafka.ListConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConfigurations, varargs...)
	ret0, _ := ret[0].(*kafka.ListConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConfigurations, reflect.TypeOf((*MockKafkaClient)(nil).ListConfigurations), varargs...)
}

func (m *MockKafkaClient) ListKafkaVersions(arg0 context.Context, arg1 *kafka.ListKafkaVersionsInput, arg2 ...func(*kafka.Options)) (*kafka.ListKafkaVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListKafkaVersions, varargs...)
	ret0, _ := ret[0].(*kafka.ListKafkaVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListKafkaVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListKafkaVersions, reflect.TypeOf((*MockKafkaClient)(nil).ListKafkaVersions), varargs...)
}

func (m *MockKafkaClient) ListNodes(arg0 context.Context, arg1 *kafka.ListNodesInput, arg2 ...func(*kafka.Options)) (*kafka.ListNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNodes, varargs...)
	ret0, _ := ret[0].(*kafka.ListNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNodes, reflect.TypeOf((*MockKafkaClient)(nil).ListNodes), varargs...)
}

func (m *MockKafkaClient) ListScramSecrets(arg0 context.Context, arg1 *kafka.ListScramSecretsInput, arg2 ...func(*kafka.Options)) (*kafka.ListScramSecretsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListScramSecrets, varargs...)
	ret0, _ := ret[0].(*kafka.ListScramSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListScramSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListScramSecrets, reflect.TypeOf((*MockKafkaClient)(nil).ListScramSecrets), varargs...)
}

func (m *MockKafkaClient) ListTagsForResource(arg0 context.Context, arg1 *kafka.ListTagsForResourceInput, arg2 ...func(*kafka.Options)) (*kafka.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*kafka.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKafkaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockKafkaClient)(nil).ListTagsForResource), varargs...)
}
