package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	elasticsearchservice "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticsearchserviceClient struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticsearchserviceClientMockRecorder
}

type MockElasticsearchserviceClientMockRecorder struct {
	mock *MockElasticsearchserviceClient
}

func NewMockElasticsearchserviceClient(ctrl *gomock.Controller) *MockElasticsearchserviceClient {
	mock := &MockElasticsearchserviceClient{ctrl: ctrl}
	mock.recorder = &MockElasticsearchserviceClientMockRecorder{mock}
	return mock
}

func (m *MockElasticsearchserviceClient) EXPECT() *MockElasticsearchserviceClientMockRecorder {
	return m.recorder
}

func (m *MockElasticsearchserviceClient) DescribeDomainAutoTunes(arg0 context.Context, arg1 *elasticsearchservice.DescribeDomainAutoTunesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeDomainAutoTunesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDomainAutoTunes, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeDomainAutoTunesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeDomainAutoTunes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDomainAutoTunes, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeDomainAutoTunes), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeDomainChangeProgress(arg0 context.Context, arg1 *elasticsearchservice.DescribeDomainChangeProgressInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeDomainChangeProgressOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDomainChangeProgress, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeDomainChangeProgressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeDomainChangeProgress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDomainChangeProgress, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeDomainChangeProgress), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeElasticsearchDomain(arg0 context.Context, arg1 *elasticsearchservice.DescribeElasticsearchDomainInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeElasticsearchDomain, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeElasticsearchDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeElasticsearchDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeElasticsearchDomain, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeElasticsearchDomain), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeElasticsearchDomainConfig(arg0 context.Context, arg1 *elasticsearchservice.DescribeElasticsearchDomainConfigInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeElasticsearchDomainConfig, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeElasticsearchDomainConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeElasticsearchDomainConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeElasticsearchDomainConfig, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeElasticsearchDomainConfig), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeElasticsearchDomains(arg0 context.Context, arg1 *elasticsearchservice.DescribeElasticsearchDomainsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeElasticsearchDomains, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeElasticsearchDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeElasticsearchDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeElasticsearchDomains, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeElasticsearchDomains), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeElasticsearchInstanceTypeLimits(arg0 context.Context, arg1 *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeElasticsearchInstanceTypeLimits, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeElasticsearchInstanceTypeLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeElasticsearchInstanceTypeLimits, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeElasticsearchInstanceTypeLimits), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeInboundCrossClusterSearchConnections(arg0 context.Context, arg1 *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInboundCrossClusterSearchConnections, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeInboundCrossClusterSearchConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInboundCrossClusterSearchConnections, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeInboundCrossClusterSearchConnections), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeOutboundCrossClusterSearchConnections(arg0 context.Context, arg1 *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOutboundCrossClusterSearchConnections, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeOutboundCrossClusterSearchConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOutboundCrossClusterSearchConnections, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeOutboundCrossClusterSearchConnections), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribePackages(arg0 context.Context, arg1 *elasticsearchservice.DescribePackagesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribePackagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePackages, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribePackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribePackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePackages, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribePackages), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeReservedElasticsearchInstanceOfferings(arg0 context.Context, arg1 *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedElasticsearchInstanceOfferings, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeReservedElasticsearchInstanceOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedElasticsearchInstanceOfferings, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeReservedElasticsearchInstanceOfferings), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeReservedElasticsearchInstances(arg0 context.Context, arg1 *elasticsearchservice.DescribeReservedElasticsearchInstancesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedElasticsearchInstances, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeReservedElasticsearchInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedElasticsearchInstances, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeReservedElasticsearchInstances), varargs...)
}

func (m *MockElasticsearchserviceClient) DescribeVpcEndpoints(arg0 context.Context, arg1 *elasticsearchservice.DescribeVpcEndpointsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeVpcEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpoints, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeVpcEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) DescribeVpcEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpoints, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).DescribeVpcEndpoints), varargs...)
}

func (m *MockElasticsearchserviceClient) GetCompatibleElasticsearchVersions(arg0 context.Context, arg1 *elasticsearchservice.GetCompatibleElasticsearchVersionsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCompatibleElasticsearchVersions, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) GetCompatibleElasticsearchVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCompatibleElasticsearchVersions, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).GetCompatibleElasticsearchVersions), varargs...)
}

func (m *MockElasticsearchserviceClient) GetPackageVersionHistory(arg0 context.Context, arg1 *elasticsearchservice.GetPackageVersionHistoryInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetPackageVersionHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPackageVersionHistory, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.GetPackageVersionHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) GetPackageVersionHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPackageVersionHistory, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).GetPackageVersionHistory), varargs...)
}

func (m *MockElasticsearchserviceClient) GetUpgradeHistory(arg0 context.Context, arg1 *elasticsearchservice.GetUpgradeHistoryInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUpgradeHistory, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.GetUpgradeHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) GetUpgradeHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUpgradeHistory, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).GetUpgradeHistory), varargs...)
}

func (m *MockElasticsearchserviceClient) GetUpgradeStatus(arg0 context.Context, arg1 *elasticsearchservice.GetUpgradeStatusInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUpgradeStatus, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.GetUpgradeStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) GetUpgradeStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUpgradeStatus, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).GetUpgradeStatus), varargs...)
}

func (m *MockElasticsearchserviceClient) ListDomainNames(arg0 context.Context, arg1 *elasticsearchservice.ListDomainNamesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomainNames, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomainNames, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListDomainNames), varargs...)
}

func (m *MockElasticsearchserviceClient) ListDomainsForPackage(arg0 context.Context, arg1 *elasticsearchservice.ListDomainsForPackageInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomainsForPackage, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListDomainsForPackageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListDomainsForPackage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomainsForPackage, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListDomainsForPackage), varargs...)
}

func (m *MockElasticsearchserviceClient) ListElasticsearchInstanceTypes(arg0 context.Context, arg1 *elasticsearchservice.ListElasticsearchInstanceTypesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListElasticsearchInstanceTypes, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListElasticsearchInstanceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListElasticsearchInstanceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListElasticsearchInstanceTypes, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListElasticsearchInstanceTypes), varargs...)
}

func (m *MockElasticsearchserviceClient) ListElasticsearchVersions(arg0 context.Context, arg1 *elasticsearchservice.ListElasticsearchVersionsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListElasticsearchVersions, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListElasticsearchVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListElasticsearchVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListElasticsearchVersions, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListElasticsearchVersions), varargs...)
}

func (m *MockElasticsearchserviceClient) ListPackagesForDomain(arg0 context.Context, arg1 *elasticsearchservice.ListPackagesForDomainInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPackagesForDomain, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListPackagesForDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListPackagesForDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPackagesForDomain, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListPackagesForDomain), varargs...)
}

func (m *MockElasticsearchserviceClient) ListTags(arg0 context.Context, arg1 *elasticsearchservice.ListTagsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListTags), varargs...)
}

func (m *MockElasticsearchserviceClient) ListVpcEndpointAccess(arg0 context.Context, arg1 *elasticsearchservice.ListVpcEndpointAccessInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListVpcEndpointAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVpcEndpointAccess, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListVpcEndpointAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListVpcEndpointAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVpcEndpointAccess, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListVpcEndpointAccess), varargs...)
}

func (m *MockElasticsearchserviceClient) ListVpcEndpoints(arg0 context.Context, arg1 *elasticsearchservice.ListVpcEndpointsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListVpcEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVpcEndpoints, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListVpcEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListVpcEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVpcEndpoints, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListVpcEndpoints), varargs...)
}

func (m *MockElasticsearchserviceClient) ListVpcEndpointsForDomain(arg0 context.Context, arg1 *elasticsearchservice.ListVpcEndpointsForDomainInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListVpcEndpointsForDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVpcEndpointsForDomain, varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListVpcEndpointsForDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticsearchserviceClientMockRecorder) ListVpcEndpointsForDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVpcEndpointsForDomain, reflect.TypeOf((*MockElasticsearchserviceClient)(nil).ListVpcEndpointsForDomain), varargs...)
}
