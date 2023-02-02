package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	rds "github.com/aws/aws-sdk-go-v2/service/rds"
	gomock "github.com/golang/mock/gomock"
)

type MockRdsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRdsClientMockRecorder
}

type MockRdsClientMockRecorder struct {
	mock *MockRdsClient
}

func NewMockRdsClient(ctrl *gomock.Controller) *MockRdsClient {
	mock := &MockRdsClient{ctrl: ctrl}
	mock.recorder = &MockRdsClientMockRecorder{mock}
	return mock
}

func (m *MockRdsClient) EXPECT() *MockRdsClientMockRecorder {
	return m.recorder
}

func (m *MockRdsClient) DescribeAccountAttributes(arg0 context.Context, arg1 *rds.DescribeAccountAttributesInput, arg2 ...func(*rds.Options)) (*rds.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAttributes, varargs...)
	ret0, _ := ret[0].(*rds.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeAccountAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAttributes, reflect.TypeOf((*MockRdsClient)(nil).DescribeAccountAttributes), varargs...)
}

func (m *MockRdsClient) DescribeBlueGreenDeployments(arg0 context.Context, arg1 *rds.DescribeBlueGreenDeploymentsInput, arg2 ...func(*rds.Options)) (*rds.DescribeBlueGreenDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBlueGreenDeployments, varargs...)
	ret0, _ := ret[0].(*rds.DescribeBlueGreenDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeBlueGreenDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBlueGreenDeployments, reflect.TypeOf((*MockRdsClient)(nil).DescribeBlueGreenDeployments), varargs...)
}

func (m *MockRdsClient) DescribeCertificates(arg0 context.Context, arg1 *rds.DescribeCertificatesInput, arg2 ...func(*rds.Options)) (*rds.DescribeCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCertificates, varargs...)
	ret0, _ := ret[0].(*rds.DescribeCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCertificates, reflect.TypeOf((*MockRdsClient)(nil).DescribeCertificates), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterBacktracks(arg0 context.Context, arg1 *rds.DescribeDBClusterBacktracksInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterBacktracksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterBacktracks, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterBacktracksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterBacktracks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterBacktracks, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterBacktracks), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterEndpoints(arg0 context.Context, arg1 *rds.DescribeDBClusterEndpointsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterEndpoints, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterEndpoints, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterEndpoints), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterParameterGroups(arg0 context.Context, arg1 *rds.DescribeDBClusterParameterGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterParameterGroups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterParameterGroups, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterParameterGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterParameters(arg0 context.Context, arg1 *rds.DescribeDBClusterParametersInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterParameters, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterParameters, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterParameters), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterSnapshotAttributes(arg0 context.Context, arg1 *rds.DescribeDBClusterSnapshotAttributesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterSnapshotAttributes, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterSnapshotAttributes, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterSnapshotAttributes), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterSnapshots(arg0 context.Context, arg1 *rds.DescribeDBClusterSnapshotsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterSnapshots, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterSnapshots, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterSnapshots), varargs...)
}

func (m *MockRdsClient) DescribeDBClusters(arg0 context.Context, arg1 *rds.DescribeDBClustersInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusters, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusters, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusters), varargs...)
}

func (m *MockRdsClient) DescribeDBEngineVersions(arg0 context.Context, arg1 *rds.DescribeDBEngineVersionsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBEngineVersions, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBEngineVersions, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBEngineVersions), varargs...)
}

func (m *MockRdsClient) DescribeDBInstanceAutomatedBackups(arg0 context.Context, arg1 *rds.DescribeDBInstanceAutomatedBackupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBInstanceAutomatedBackups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBInstanceAutomatedBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBInstanceAutomatedBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBInstanceAutomatedBackups, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBInstanceAutomatedBackups), varargs...)
}

func (m *MockRdsClient) DescribeDBInstances(arg0 context.Context, arg1 *rds.DescribeDBInstancesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBInstances, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBInstances, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBInstances), varargs...)
}

func (m *MockRdsClient) DescribeDBLogFiles(arg0 context.Context, arg1 *rds.DescribeDBLogFilesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBLogFilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBLogFiles, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBLogFilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBLogFiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBLogFiles, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBLogFiles), varargs...)
}

func (m *MockRdsClient) DescribeDBParameterGroups(arg0 context.Context, arg1 *rds.DescribeDBParameterGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBParameterGroups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBParameterGroups, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBParameterGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBParameters(arg0 context.Context, arg1 *rds.DescribeDBParametersInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBParameters, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBParameters, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBParameters), varargs...)
}

func (m *MockRdsClient) DescribeDBProxies(arg0 context.Context, arg1 *rds.DescribeDBProxiesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBProxiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBProxies, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBProxiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBProxies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBProxies, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBProxies), varargs...)
}

func (m *MockRdsClient) DescribeDBProxyEndpoints(arg0 context.Context, arg1 *rds.DescribeDBProxyEndpointsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBProxyEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBProxyEndpoints, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBProxyEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBProxyEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBProxyEndpoints, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBProxyEndpoints), varargs...)
}

func (m *MockRdsClient) DescribeDBProxyTargetGroups(arg0 context.Context, arg1 *rds.DescribeDBProxyTargetGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBProxyTargetGroups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBProxyTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBProxyTargetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBProxyTargetGroups, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBProxyTargetGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBProxyTargets(arg0 context.Context, arg1 *rds.DescribeDBProxyTargetsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBProxyTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBProxyTargets, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBProxyTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBProxyTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBProxyTargets, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBProxyTargets), varargs...)
}

func (m *MockRdsClient) DescribeDBSecurityGroups(arg0 context.Context, arg1 *rds.DescribeDBSecurityGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBSecurityGroups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBSecurityGroups, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSecurityGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBSnapshotAttributes(arg0 context.Context, arg1 *rds.DescribeDBSnapshotAttributesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBSnapshotAttributes, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBSnapshotAttributes, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSnapshotAttributes), varargs...)
}

func (m *MockRdsClient) DescribeDBSnapshots(arg0 context.Context, arg1 *rds.DescribeDBSnapshotsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBSnapshots, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBSnapshots, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSnapshots), varargs...)
}

func (m *MockRdsClient) DescribeDBSubnetGroups(arg0 context.Context, arg1 *rds.DescribeDBSubnetGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBSubnetGroups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBSubnetGroups, reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSubnetGroups), varargs...)
}

func (m *MockRdsClient) DescribeEngineDefaultClusterParameters(arg0 context.Context, arg1 *rds.DescribeEngineDefaultClusterParametersInput, arg2 ...func(*rds.Options)) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEngineDefaultClusterParameters, varargs...)
	ret0, _ := ret[0].(*rds.DescribeEngineDefaultClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeEngineDefaultClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEngineDefaultClusterParameters, reflect.TypeOf((*MockRdsClient)(nil).DescribeEngineDefaultClusterParameters), varargs...)
}

func (m *MockRdsClient) DescribeEngineDefaultParameters(arg0 context.Context, arg1 *rds.DescribeEngineDefaultParametersInput, arg2 ...func(*rds.Options)) (*rds.DescribeEngineDefaultParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEngineDefaultParameters, varargs...)
	ret0, _ := ret[0].(*rds.DescribeEngineDefaultParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeEngineDefaultParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEngineDefaultParameters, reflect.TypeOf((*MockRdsClient)(nil).DescribeEngineDefaultParameters), varargs...)
}

func (m *MockRdsClient) DescribeEventCategories(arg0 context.Context, arg1 *rds.DescribeEventCategoriesInput, arg2 ...func(*rds.Options)) (*rds.DescribeEventCategoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventCategories, varargs...)
	ret0, _ := ret[0].(*rds.DescribeEventCategoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeEventCategories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventCategories, reflect.TypeOf((*MockRdsClient)(nil).DescribeEventCategories), varargs...)
}

func (m *MockRdsClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *rds.DescribeEventSubscriptionsInput, arg2 ...func(*rds.Options)) (*rds.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventSubscriptions, varargs...)
	ret0, _ := ret[0].(*rds.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventSubscriptions, reflect.TypeOf((*MockRdsClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockRdsClient) DescribeEvents(arg0 context.Context, arg1 *rds.DescribeEventsInput, arg2 ...func(*rds.Options)) (*rds.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*rds.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockRdsClient)(nil).DescribeEvents), varargs...)
}

func (m *MockRdsClient) DescribeExportTasks(arg0 context.Context, arg1 *rds.DescribeExportTasksInput, arg2 ...func(*rds.Options)) (*rds.DescribeExportTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExportTasks, varargs...)
	ret0, _ := ret[0].(*rds.DescribeExportTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeExportTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExportTasks, reflect.TypeOf((*MockRdsClient)(nil).DescribeExportTasks), varargs...)
}

func (m *MockRdsClient) DescribeGlobalClusters(arg0 context.Context, arg1 *rds.DescribeGlobalClustersInput, arg2 ...func(*rds.Options)) (*rds.DescribeGlobalClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGlobalClusters, varargs...)
	ret0, _ := ret[0].(*rds.DescribeGlobalClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeGlobalClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGlobalClusters, reflect.TypeOf((*MockRdsClient)(nil).DescribeGlobalClusters), varargs...)
}

func (m *MockRdsClient) DescribeOptionGroupOptions(arg0 context.Context, arg1 *rds.DescribeOptionGroupOptionsInput, arg2 ...func(*rds.Options)) (*rds.DescribeOptionGroupOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOptionGroupOptions, varargs...)
	ret0, _ := ret[0].(*rds.DescribeOptionGroupOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeOptionGroupOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOptionGroupOptions, reflect.TypeOf((*MockRdsClient)(nil).DescribeOptionGroupOptions), varargs...)
}

func (m *MockRdsClient) DescribeOptionGroups(arg0 context.Context, arg1 *rds.DescribeOptionGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeOptionGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOptionGroups, varargs...)
	ret0, _ := ret[0].(*rds.DescribeOptionGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeOptionGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOptionGroups, reflect.TypeOf((*MockRdsClient)(nil).DescribeOptionGroups), varargs...)
}

func (m *MockRdsClient) DescribeOrderableDBInstanceOptions(arg0 context.Context, arg1 *rds.DescribeOrderableDBInstanceOptionsInput, arg2 ...func(*rds.Options)) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrderableDBInstanceOptions, varargs...)
	ret0, _ := ret[0].(*rds.DescribeOrderableDBInstanceOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeOrderableDBInstanceOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrderableDBInstanceOptions, reflect.TypeOf((*MockRdsClient)(nil).DescribeOrderableDBInstanceOptions), varargs...)
}

func (m *MockRdsClient) DescribePendingMaintenanceActions(arg0 context.Context, arg1 *rds.DescribePendingMaintenanceActionsInput, arg2 ...func(*rds.Options)) (*rds.DescribePendingMaintenanceActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePendingMaintenanceActions, varargs...)
	ret0, _ := ret[0].(*rds.DescribePendingMaintenanceActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribePendingMaintenanceActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePendingMaintenanceActions, reflect.TypeOf((*MockRdsClient)(nil).DescribePendingMaintenanceActions), varargs...)
}

func (m *MockRdsClient) DescribeReservedDBInstances(arg0 context.Context, arg1 *rds.DescribeReservedDBInstancesInput, arg2 ...func(*rds.Options)) (*rds.DescribeReservedDBInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedDBInstances, varargs...)
	ret0, _ := ret[0].(*rds.DescribeReservedDBInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeReservedDBInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedDBInstances, reflect.TypeOf((*MockRdsClient)(nil).DescribeReservedDBInstances), varargs...)
}

func (m *MockRdsClient) DescribeReservedDBInstancesOfferings(arg0 context.Context, arg1 *rds.DescribeReservedDBInstancesOfferingsInput, arg2 ...func(*rds.Options)) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedDBInstancesOfferings, varargs...)
	ret0, _ := ret[0].(*rds.DescribeReservedDBInstancesOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeReservedDBInstancesOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedDBInstancesOfferings, reflect.TypeOf((*MockRdsClient)(nil).DescribeReservedDBInstancesOfferings), varargs...)
}

func (m *MockRdsClient) DescribeSourceRegions(arg0 context.Context, arg1 *rds.DescribeSourceRegionsInput, arg2 ...func(*rds.Options)) (*rds.DescribeSourceRegionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSourceRegions, varargs...)
	ret0, _ := ret[0].(*rds.DescribeSourceRegionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeSourceRegions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSourceRegions, reflect.TypeOf((*MockRdsClient)(nil).DescribeSourceRegions), varargs...)
}

func (m *MockRdsClient) DescribeValidDBInstanceModifications(arg0 context.Context, arg1 *rds.DescribeValidDBInstanceModificationsInput, arg2 ...func(*rds.Options)) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeValidDBInstanceModifications, varargs...)
	ret0, _ := ret[0].(*rds.DescribeValidDBInstanceModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeValidDBInstanceModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeValidDBInstanceModifications, reflect.TypeOf((*MockRdsClient)(nil).DescribeValidDBInstanceModifications), varargs...)
}

func (m *MockRdsClient) ListTagsForResource(arg0 context.Context, arg1 *rds.ListTagsForResourceInput, arg2 ...func(*rds.Options)) (*rds.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*rds.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockRdsClient)(nil).ListTagsForResource), varargs...)
}
