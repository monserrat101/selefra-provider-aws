package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	docdb "github.com/aws/aws-sdk-go-v2/service/docdb"
	gomock "github.com/golang/mock/gomock"
)

type MockDocdbClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDocdbClientMockRecorder
}

type MockDocdbClientMockRecorder struct {
	mock *MockDocdbClient
}

func NewMockDocdbClient(ctrl *gomock.Controller) *MockDocdbClient {
	mock := &MockDocdbClient{ctrl: ctrl}
	mock.recorder = &MockDocdbClientMockRecorder{mock}
	return mock
}

func (m *MockDocdbClient) EXPECT() *MockDocdbClientMockRecorder {
	return m.recorder
}

func (m *MockDocdbClient) DescribeCertificates(arg0 context.Context, arg1 *docdb.DescribeCertificatesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCertificates, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCertificates, reflect.TypeOf((*MockDocdbClient)(nil).DescribeCertificates), varargs...)
}

func (m *MockDocdbClient) DescribeDBClusterParameterGroups(arg0 context.Context, arg1 *docdb.DescribeDBClusterParameterGroupsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterParameterGroups, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBClusterParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterParameterGroups, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterParameterGroups), varargs...)
}

func (m *MockDocdbClient) DescribeDBClusterParameters(arg0 context.Context, arg1 *docdb.DescribeDBClusterParametersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterParameters, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterParameters, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterParameters), varargs...)
}

func (m *MockDocdbClient) DescribeDBClusterSnapshotAttributes(arg0 context.Context, arg1 *docdb.DescribeDBClusterSnapshotAttributesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterSnapshotAttributes, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBClusterSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterSnapshotAttributes, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterSnapshotAttributes), varargs...)
}

func (m *MockDocdbClient) DescribeDBClusterSnapshots(arg0 context.Context, arg1 *docdb.DescribeDBClusterSnapshotsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusterSnapshots, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusterSnapshots, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterSnapshots), varargs...)
}

func (m *MockDocdbClient) DescribeDBClusters(arg0 context.Context, arg1 *docdb.DescribeDBClustersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBClusters, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBClusters, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusters), varargs...)
}

func (m *MockDocdbClient) DescribeDBEngineVersions(arg0 context.Context, arg1 *docdb.DescribeDBEngineVersionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBEngineVersions, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBEngineVersions, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBEngineVersions), varargs...)
}

func (m *MockDocdbClient) DescribeDBInstances(arg0 context.Context, arg1 *docdb.DescribeDBInstancesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBInstances, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBInstances, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBInstances), varargs...)
}

func (m *MockDocdbClient) DescribeDBSubnetGroups(arg0 context.Context, arg1 *docdb.DescribeDBSubnetGroupsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDBSubnetGroups, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeDBSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDBSubnetGroups, reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBSubnetGroups), varargs...)
}

func (m *MockDocdbClient) DescribeEngineDefaultClusterParameters(arg0 context.Context, arg1 *docdb.DescribeEngineDefaultClusterParametersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEngineDefaultClusterParameters, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEngineDefaultClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeEngineDefaultClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEngineDefaultClusterParameters, reflect.TypeOf((*MockDocdbClient)(nil).DescribeEngineDefaultClusterParameters), varargs...)
}

func (m *MockDocdbClient) DescribeEventCategories(arg0 context.Context, arg1 *docdb.DescribeEventCategoriesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEventCategoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventCategories, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEventCategoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeEventCategories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventCategories, reflect.TypeOf((*MockDocdbClient)(nil).DescribeEventCategories), varargs...)
}

func (m *MockDocdbClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *docdb.DescribeEventSubscriptionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventSubscriptions, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventSubscriptions, reflect.TypeOf((*MockDocdbClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockDocdbClient) DescribeEvents(arg0 context.Context, arg1 *docdb.DescribeEventsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockDocdbClient)(nil).DescribeEvents), varargs...)
}

func (m *MockDocdbClient) DescribeGlobalClusters(arg0 context.Context, arg1 *docdb.DescribeGlobalClustersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeGlobalClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGlobalClusters, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeGlobalClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeGlobalClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGlobalClusters, reflect.TypeOf((*MockDocdbClient)(nil).DescribeGlobalClusters), varargs...)
}

func (m *MockDocdbClient) DescribeOrderableDBInstanceOptions(arg0 context.Context, arg1 *docdb.DescribeOrderableDBInstanceOptionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrderableDBInstanceOptions, varargs...)
	ret0, _ := ret[0].(*docdb.DescribeOrderableDBInstanceOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribeOrderableDBInstanceOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrderableDBInstanceOptions, reflect.TypeOf((*MockDocdbClient)(nil).DescribeOrderableDBInstanceOptions), varargs...)
}

func (m *MockDocdbClient) DescribePendingMaintenanceActions(arg0 context.Context, arg1 *docdb.DescribePendingMaintenanceActionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePendingMaintenanceActions, varargs...)
	ret0, _ := ret[0].(*docdb.DescribePendingMaintenanceActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) DescribePendingMaintenanceActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePendingMaintenanceActions, reflect.TypeOf((*MockDocdbClient)(nil).DescribePendingMaintenanceActions), varargs...)
}

func (m *MockDocdbClient) ListTagsForResource(arg0 context.Context, arg1 *docdb.ListTagsForResourceInput, arg2 ...func(*docdb.Options)) (*docdb.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*docdb.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocdbClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockDocdbClient)(nil).ListTagsForResource), varargs...)
}
