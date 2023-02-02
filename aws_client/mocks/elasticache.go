package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	elasticache "github.com/aws/aws-sdk-go-v2/service/elasticache"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticacheClient struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticacheClientMockRecorder
}

type MockElasticacheClientMockRecorder struct {
	mock *MockElasticacheClient
}

func NewMockElasticacheClient(ctrl *gomock.Controller) *MockElasticacheClient {
	mock := &MockElasticacheClient{ctrl: ctrl}
	mock.recorder = &MockElasticacheClientMockRecorder{mock}
	return mock
}

func (m *MockElasticacheClient) EXPECT() *MockElasticacheClientMockRecorder {
	return m.recorder
}

func (m *MockElasticacheClient) DescribeCacheClusters(arg0 context.Context, arg1 *elasticache.DescribeCacheClustersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCacheClusters, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeCacheClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCacheClusters, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheClusters), varargs...)
}

func (m *MockElasticacheClient) DescribeCacheEngineVersions(arg0 context.Context, arg1 *elasticache.DescribeCacheEngineVersionsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCacheEngineVersions, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeCacheEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCacheEngineVersions, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheEngineVersions), varargs...)
}

func (m *MockElasticacheClient) DescribeCacheParameterGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheParameterGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCacheParameterGroups, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeCacheParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCacheParameterGroups, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheParameterGroups), varargs...)
}

func (m *MockElasticacheClient) DescribeCacheParameters(arg0 context.Context, arg1 *elasticache.DescribeCacheParametersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCacheParameters, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeCacheParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCacheParameters, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheParameters), varargs...)
}

func (m *MockElasticacheClient) DescribeCacheSecurityGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheSecurityGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCacheSecurityGroups, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeCacheSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCacheSecurityGroups, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheSecurityGroups), varargs...)
}

func (m *MockElasticacheClient) DescribeCacheSubnetGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheSubnetGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCacheSubnetGroups, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeCacheSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCacheSubnetGroups, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheSubnetGroups), varargs...)
}

func (m *MockElasticacheClient) DescribeEngineDefaultParameters(arg0 context.Context, arg1 *elasticache.DescribeEngineDefaultParametersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEngineDefaultParameters, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeEngineDefaultParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeEngineDefaultParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEngineDefaultParameters, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeEngineDefaultParameters), varargs...)
}

func (m *MockElasticacheClient) DescribeEvents(arg0 context.Context, arg1 *elasticache.DescribeEventsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeEvents), varargs...)
}

func (m *MockElasticacheClient) DescribeGlobalReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeGlobalReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeGlobalReplicationGroups, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeGlobalReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeGlobalReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeGlobalReplicationGroups, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeGlobalReplicationGroups), varargs...)
}

func (m *MockElasticacheClient) DescribeReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplicationGroups, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplicationGroups, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeReplicationGroups), varargs...)
}

func (m *MockElasticacheClient) DescribeReservedCacheNodes(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedCacheNodes, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeReservedCacheNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedCacheNodes, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeReservedCacheNodes), varargs...)
}

func (m *MockElasticacheClient) DescribeReservedCacheNodesOfferings(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesOfferingsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedCacheNodesOfferings, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeReservedCacheNodesOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedCacheNodesOfferings, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeReservedCacheNodesOfferings), varargs...)
}

func (m *MockElasticacheClient) DescribeServiceUpdates(arg0 context.Context, arg1 *elasticache.DescribeServiceUpdatesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeServiceUpdates, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeServiceUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeServiceUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeServiceUpdates, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeServiceUpdates), varargs...)
}

func (m *MockElasticacheClient) DescribeSnapshots(arg0 context.Context, arg1 *elasticache.DescribeSnapshotsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshots, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshots, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeSnapshots), varargs...)
}

func (m *MockElasticacheClient) DescribeUpdateActions(arg0 context.Context, arg1 *elasticache.DescribeUpdateActionsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUpdateActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUpdateActions, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUpdateActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeUpdateActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUpdateActions, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeUpdateActions), varargs...)
}

func (m *MockElasticacheClient) DescribeUserGroups(arg0 context.Context, arg1 *elasticache.DescribeUserGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserGroups, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUserGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeUserGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserGroups, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeUserGroups), varargs...)
}

func (m *MockElasticacheClient) DescribeUsers(arg0 context.Context, arg1 *elasticache.DescribeUsersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUsers, varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) DescribeUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUsers, reflect.TypeOf((*MockElasticacheClient)(nil).DescribeUsers), varargs...)
}

func (m *MockElasticacheClient) ListAllowedNodeTypeModifications(arg0 context.Context, arg1 *elasticache.ListAllowedNodeTypeModificationsInput, arg2 ...func(*elasticache.Options)) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAllowedNodeTypeModifications, varargs...)
	ret0, _ := ret[0].(*elasticache.ListAllowedNodeTypeModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) ListAllowedNodeTypeModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAllowedNodeTypeModifications, reflect.TypeOf((*MockElasticacheClient)(nil).ListAllowedNodeTypeModifications), varargs...)
}

func (m *MockElasticacheClient) ListTagsForResource(arg0 context.Context, arg1 *elasticache.ListTagsForResourceInput, arg2 ...func(*elasticache.Options)) (*elasticache.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*elasticache.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticacheClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockElasticacheClient)(nil).ListTagsForResource), varargs...)
}
