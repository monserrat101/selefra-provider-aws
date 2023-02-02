package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	redshift "github.com/aws/aws-sdk-go-v2/service/redshift"
	gomock "github.com/golang/mock/gomock"
)

type MockRedshiftClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRedshiftClientMockRecorder
}

type MockRedshiftClientMockRecorder struct {
	mock *MockRedshiftClient
}

func NewMockRedshiftClient(ctrl *gomock.Controller) *MockRedshiftClient {
	mock := &MockRedshiftClient{ctrl: ctrl}
	mock.recorder = &MockRedshiftClientMockRecorder{mock}
	return mock
}

func (m *MockRedshiftClient) EXPECT() *MockRedshiftClientMockRecorder {
	return m.recorder
}

func (m *MockRedshiftClient) DescribeAccountAttributes(arg0 context.Context, arg1 *redshift.DescribeAccountAttributesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAttributes, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeAccountAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAttributes, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeAccountAttributes), varargs...)
}

func (m *MockRedshiftClient) DescribeAuthenticationProfiles(arg0 context.Context, arg1 *redshift.DescribeAuthenticationProfilesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeAuthenticationProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAuthenticationProfiles, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeAuthenticationProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeAuthenticationProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAuthenticationProfiles, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeAuthenticationProfiles), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterDbRevisions(arg0 context.Context, arg1 *redshift.DescribeClusterDbRevisionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterDbRevisionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterDbRevisions, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterDbRevisionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterDbRevisions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterDbRevisions, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterDbRevisions), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterParameterGroups(arg0 context.Context, arg1 *redshift.DescribeClusterParameterGroupsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterParameterGroups, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterParameterGroups, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterParameterGroups), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterParameters(arg0 context.Context, arg1 *redshift.DescribeClusterParametersInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterParameters, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterParameters, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterParameters), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterSecurityGroups(arg0 context.Context, arg1 *redshift.DescribeClusterSecurityGroupsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterSecurityGroups, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterSecurityGroups, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterSecurityGroups), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterSnapshots(arg0 context.Context, arg1 *redshift.DescribeClusterSnapshotsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterSnapshots, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterSnapshots, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterSnapshots), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterSubnetGroups(arg0 context.Context, arg1 *redshift.DescribeClusterSubnetGroupsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterSubnetGroups, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterSubnetGroups, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterSubnetGroups), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterTracks(arg0 context.Context, arg1 *redshift.DescribeClusterTracksInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterTracksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterTracks, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterTracksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterTracks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterTracks, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterTracks), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterVersions(arg0 context.Context, arg1 *redshift.DescribeClusterVersionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusterVersions, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusterVersions, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterVersions), varargs...)
}

func (m *MockRedshiftClient) DescribeClusters(arg0 context.Context, arg1 *redshift.DescribeClustersInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusters, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusters, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusters), varargs...)
}

func (m *MockRedshiftClient) DescribeDataShares(arg0 context.Context, arg1 *redshift.DescribeDataSharesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeDataSharesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataShares, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeDataSharesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeDataShares(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataShares, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeDataShares), varargs...)
}

func (m *MockRedshiftClient) DescribeDataSharesForConsumer(arg0 context.Context, arg1 *redshift.DescribeDataSharesForConsumerInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeDataSharesForConsumerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataSharesForConsumer, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeDataSharesForConsumerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeDataSharesForConsumer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataSharesForConsumer, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeDataSharesForConsumer), varargs...)
}

func (m *MockRedshiftClient) DescribeDataSharesForProducer(arg0 context.Context, arg1 *redshift.DescribeDataSharesForProducerInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeDataSharesForProducerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataSharesForProducer, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeDataSharesForProducerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeDataSharesForProducer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataSharesForProducer, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeDataSharesForProducer), varargs...)
}

func (m *MockRedshiftClient) DescribeDefaultClusterParameters(arg0 context.Context, arg1 *redshift.DescribeDefaultClusterParametersInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeDefaultClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDefaultClusterParameters, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeDefaultClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeDefaultClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDefaultClusterParameters, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeDefaultClusterParameters), varargs...)
}

func (m *MockRedshiftClient) DescribeEndpointAccess(arg0 context.Context, arg1 *redshift.DescribeEndpointAccessInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeEndpointAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpointAccess, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeEndpointAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeEndpointAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpointAccess, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeEndpointAccess), varargs...)
}

func (m *MockRedshiftClient) DescribeEndpointAuthorization(arg0 context.Context, arg1 *redshift.DescribeEndpointAuthorizationInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeEndpointAuthorizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpointAuthorization, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeEndpointAuthorizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeEndpointAuthorization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpointAuthorization, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeEndpointAuthorization), varargs...)
}

func (m *MockRedshiftClient) DescribeEventCategories(arg0 context.Context, arg1 *redshift.DescribeEventCategoriesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeEventCategoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventCategories, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeEventCategoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeEventCategories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventCategories, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeEventCategories), varargs...)
}

func (m *MockRedshiftClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *redshift.DescribeEventSubscriptionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventSubscriptions, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventSubscriptions, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockRedshiftClient) DescribeEvents(arg0 context.Context, arg1 *redshift.DescribeEventsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEvents, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEvents, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeEvents), varargs...)
}

func (m *MockRedshiftClient) DescribeHsmClientCertificates(arg0 context.Context, arg1 *redshift.DescribeHsmClientCertificatesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeHsmClientCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHsmClientCertificates, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeHsmClientCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeHsmClientCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHsmClientCertificates, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeHsmClientCertificates), varargs...)
}

func (m *MockRedshiftClient) DescribeHsmConfigurations(arg0 context.Context, arg1 *redshift.DescribeHsmConfigurationsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeHsmConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHsmConfigurations, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeHsmConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeHsmConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHsmConfigurations, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeHsmConfigurations), varargs...)
}

func (m *MockRedshiftClient) DescribeLoggingStatus(arg0 context.Context, arg1 *redshift.DescribeLoggingStatusInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeLoggingStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoggingStatus, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeLoggingStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeLoggingStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoggingStatus, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeLoggingStatus), varargs...)
}

func (m *MockRedshiftClient) DescribeNodeConfigurationOptions(arg0 context.Context, arg1 *redshift.DescribeNodeConfigurationOptionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNodeConfigurationOptions, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeNodeConfigurationOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeNodeConfigurationOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNodeConfigurationOptions, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeNodeConfigurationOptions), varargs...)
}

func (m *MockRedshiftClient) DescribeOrderableClusterOptions(arg0 context.Context, arg1 *redshift.DescribeOrderableClusterOptionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOrderableClusterOptions, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeOrderableClusterOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeOrderableClusterOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOrderableClusterOptions, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeOrderableClusterOptions), varargs...)
}

func (m *MockRedshiftClient) DescribePartners(arg0 context.Context, arg1 *redshift.DescribePartnersInput, arg2 ...func(*redshift.Options)) (*redshift.DescribePartnersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePartners, varargs...)
	ret0, _ := ret[0].(*redshift.DescribePartnersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribePartners(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePartners, reflect.TypeOf((*MockRedshiftClient)(nil).DescribePartners), varargs...)
}

func (m *MockRedshiftClient) DescribeReservedNodeExchangeStatus(arg0 context.Context, arg1 *redshift.DescribeReservedNodeExchangeStatusInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeReservedNodeExchangeStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedNodeExchangeStatus, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeReservedNodeExchangeStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeReservedNodeExchangeStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedNodeExchangeStatus, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeReservedNodeExchangeStatus), varargs...)
}

func (m *MockRedshiftClient) DescribeReservedNodeOfferings(arg0 context.Context, arg1 *redshift.DescribeReservedNodeOfferingsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedNodeOfferings, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeReservedNodeOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeReservedNodeOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedNodeOfferings, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeReservedNodeOfferings), varargs...)
}

func (m *MockRedshiftClient) DescribeReservedNodes(arg0 context.Context, arg1 *redshift.DescribeReservedNodesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeReservedNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedNodes, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeReservedNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeReservedNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedNodes, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeReservedNodes), varargs...)
}

func (m *MockRedshiftClient) DescribeResize(arg0 context.Context, arg1 *redshift.DescribeResizeInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeResizeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeResize, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeResizeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeResize(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeResize, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeResize), varargs...)
}

func (m *MockRedshiftClient) DescribeScheduledActions(arg0 context.Context, arg1 *redshift.DescribeScheduledActionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeScheduledActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScheduledActions, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeScheduledActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScheduledActions, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeScheduledActions), varargs...)
}

func (m *MockRedshiftClient) DescribeSnapshotCopyGrants(arg0 context.Context, arg1 *redshift.DescribeSnapshotCopyGrantsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshotCopyGrants, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeSnapshotCopyGrantsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeSnapshotCopyGrants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshotCopyGrants, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeSnapshotCopyGrants), varargs...)
}

func (m *MockRedshiftClient) DescribeSnapshotSchedules(arg0 context.Context, arg1 *redshift.DescribeSnapshotSchedulesInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeSnapshotSchedulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshotSchedules, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeSnapshotSchedulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeSnapshotSchedules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshotSchedules, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeSnapshotSchedules), varargs...)
}

func (m *MockRedshiftClient) DescribeStorage(arg0 context.Context, arg1 *redshift.DescribeStorageInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeStorageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStorage, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeStorageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeStorage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStorage, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeStorage), varargs...)
}

func (m *MockRedshiftClient) DescribeTableRestoreStatus(arg0 context.Context, arg1 *redshift.DescribeTableRestoreStatusInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeTableRestoreStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTableRestoreStatus, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeTableRestoreStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeTableRestoreStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTableRestoreStatus, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeTableRestoreStatus), varargs...)
}

func (m *MockRedshiftClient) DescribeTags(arg0 context.Context, arg1 *redshift.DescribeTagsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeTags), varargs...)
}

func (m *MockRedshiftClient) DescribeUsageLimits(arg0 context.Context, arg1 *redshift.DescribeUsageLimitsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeUsageLimitsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUsageLimits, varargs...)
	ret0, _ := ret[0].(*redshift.DescribeUsageLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeUsageLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUsageLimits, reflect.TypeOf((*MockRedshiftClient)(nil).DescribeUsageLimits), varargs...)
}

func (m *MockRedshiftClient) GetClusterCredentials(arg0 context.Context, arg1 *redshift.GetClusterCredentialsInput, arg2 ...func(*redshift.Options)) (*redshift.GetClusterCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClusterCredentials, varargs...)
	ret0, _ := ret[0].(*redshift.GetClusterCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) GetClusterCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClusterCredentials, reflect.TypeOf((*MockRedshiftClient)(nil).GetClusterCredentials), varargs...)
}

func (m *MockRedshiftClient) GetClusterCredentialsWithIAM(arg0 context.Context, arg1 *redshift.GetClusterCredentialsWithIAMInput, arg2 ...func(*redshift.Options)) (*redshift.GetClusterCredentialsWithIAMOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClusterCredentialsWithIAM, varargs...)
	ret0, _ := ret[0].(*redshift.GetClusterCredentialsWithIAMOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) GetClusterCredentialsWithIAM(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClusterCredentialsWithIAM, reflect.TypeOf((*MockRedshiftClient)(nil).GetClusterCredentialsWithIAM), varargs...)
}

func (m *MockRedshiftClient) GetReservedNodeExchangeConfigurationOptions(arg0 context.Context, arg1 *redshift.GetReservedNodeExchangeConfigurationOptionsInput, arg2 ...func(*redshift.Options)) (*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetReservedNodeExchangeConfigurationOptions, varargs...)
	ret0, _ := ret[0].(*redshift.GetReservedNodeExchangeConfigurationOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) GetReservedNodeExchangeConfigurationOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetReservedNodeExchangeConfigurationOptions, reflect.TypeOf((*MockRedshiftClient)(nil).GetReservedNodeExchangeConfigurationOptions), varargs...)
}

func (m *MockRedshiftClient) GetReservedNodeExchangeOfferings(arg0 context.Context, arg1 *redshift.GetReservedNodeExchangeOfferingsInput, arg2 ...func(*redshift.Options)) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetReservedNodeExchangeOfferings, varargs...)
	ret0, _ := ret[0].(*redshift.GetReservedNodeExchangeOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) GetReservedNodeExchangeOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetReservedNodeExchangeOfferings, reflect.TypeOf((*MockRedshiftClient)(nil).GetReservedNodeExchangeOfferings), varargs...)
}
