package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	route53 "github.com/aws/aws-sdk-go-v2/service/route53"
	gomock "github.com/golang/mock/gomock"
)

type MockRoute53Client struct {
	ctrl		*gomock.Controller
	recorder	*MockRoute53ClientMockRecorder
}

type MockRoute53ClientMockRecorder struct {
	mock *MockRoute53Client
}

func NewMockRoute53Client(ctrl *gomock.Controller) *MockRoute53Client {
	mock := &MockRoute53Client{ctrl: ctrl}
	mock.recorder = &MockRoute53ClientMockRecorder{mock}
	return mock
}

func (m *MockRoute53Client) EXPECT() *MockRoute53ClientMockRecorder {
	return m.recorder
}

func (m *MockRoute53Client) GetAccountLimit(arg0 context.Context, arg1 *route53.GetAccountLimitInput, arg2 ...func(*route53.Options)) (*route53.GetAccountLimitOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountLimit, varargs...)
	ret0, _ := ret[0].(*route53.GetAccountLimitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetAccountLimit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountLimit, reflect.TypeOf((*MockRoute53Client)(nil).GetAccountLimit), varargs...)
}

func (m *MockRoute53Client) GetChange(arg0 context.Context, arg1 *route53.GetChangeInput, arg2 ...func(*route53.Options)) (*route53.GetChangeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetChange, varargs...)
	ret0, _ := ret[0].(*route53.GetChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetChange(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetChange, reflect.TypeOf((*MockRoute53Client)(nil).GetChange), varargs...)
}

func (m *MockRoute53Client) GetCheckerIpRanges(arg0 context.Context, arg1 *route53.GetCheckerIpRangesInput, arg2 ...func(*route53.Options)) (*route53.GetCheckerIpRangesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCheckerIpRanges, varargs...)
	ret0, _ := ret[0].(*route53.GetCheckerIpRangesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetCheckerIpRanges(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCheckerIpRanges, reflect.TypeOf((*MockRoute53Client)(nil).GetCheckerIpRanges), varargs...)
}

func (m *MockRoute53Client) GetDNSSEC(arg0 context.Context, arg1 *route53.GetDNSSECInput, arg2 ...func(*route53.Options)) (*route53.GetDNSSECOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDNSSEC, varargs...)
	ret0, _ := ret[0].(*route53.GetDNSSECOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetDNSSEC(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDNSSEC, reflect.TypeOf((*MockRoute53Client)(nil).GetDNSSEC), varargs...)
}

func (m *MockRoute53Client) GetGeoLocation(arg0 context.Context, arg1 *route53.GetGeoLocationInput, arg2 ...func(*route53.Options)) (*route53.GetGeoLocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGeoLocation, varargs...)
	ret0, _ := ret[0].(*route53.GetGeoLocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetGeoLocation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGeoLocation, reflect.TypeOf((*MockRoute53Client)(nil).GetGeoLocation), varargs...)
}

func (m *MockRoute53Client) GetHealthCheck(arg0 context.Context, arg1 *route53.GetHealthCheckInput, arg2 ...func(*route53.Options)) (*route53.GetHealthCheckOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHealthCheck, varargs...)
	ret0, _ := ret[0].(*route53.GetHealthCheckOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHealthCheck(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHealthCheck, reflect.TypeOf((*MockRoute53Client)(nil).GetHealthCheck), varargs...)
}

func (m *MockRoute53Client) GetHealthCheckCount(arg0 context.Context, arg1 *route53.GetHealthCheckCountInput, arg2 ...func(*route53.Options)) (*route53.GetHealthCheckCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHealthCheckCount, varargs...)
	ret0, _ := ret[0].(*route53.GetHealthCheckCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHealthCheckCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHealthCheckCount, reflect.TypeOf((*MockRoute53Client)(nil).GetHealthCheckCount), varargs...)
}

func (m *MockRoute53Client) GetHealthCheckLastFailureReason(arg0 context.Context, arg1 *route53.GetHealthCheckLastFailureReasonInput, arg2 ...func(*route53.Options)) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHealthCheckLastFailureReason, varargs...)
	ret0, _ := ret[0].(*route53.GetHealthCheckLastFailureReasonOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHealthCheckLastFailureReason(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHealthCheckLastFailureReason, reflect.TypeOf((*MockRoute53Client)(nil).GetHealthCheckLastFailureReason), varargs...)
}

func (m *MockRoute53Client) GetHealthCheckStatus(arg0 context.Context, arg1 *route53.GetHealthCheckStatusInput, arg2 ...func(*route53.Options)) (*route53.GetHealthCheckStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHealthCheckStatus, varargs...)
	ret0, _ := ret[0].(*route53.GetHealthCheckStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHealthCheckStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHealthCheckStatus, reflect.TypeOf((*MockRoute53Client)(nil).GetHealthCheckStatus), varargs...)
}

func (m *MockRoute53Client) GetHostedZone(arg0 context.Context, arg1 *route53.GetHostedZoneInput, arg2 ...func(*route53.Options)) (*route53.GetHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHostedZone, varargs...)
	ret0, _ := ret[0].(*route53.GetHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHostedZone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHostedZone, reflect.TypeOf((*MockRoute53Client)(nil).GetHostedZone), varargs...)
}

func (m *MockRoute53Client) GetHostedZoneCount(arg0 context.Context, arg1 *route53.GetHostedZoneCountInput, arg2 ...func(*route53.Options)) (*route53.GetHostedZoneCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHostedZoneCount, varargs...)
	ret0, _ := ret[0].(*route53.GetHostedZoneCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHostedZoneCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHostedZoneCount, reflect.TypeOf((*MockRoute53Client)(nil).GetHostedZoneCount), varargs...)
}

func (m *MockRoute53Client) GetHostedZoneLimit(arg0 context.Context, arg1 *route53.GetHostedZoneLimitInput, arg2 ...func(*route53.Options)) (*route53.GetHostedZoneLimitOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHostedZoneLimit, varargs...)
	ret0, _ := ret[0].(*route53.GetHostedZoneLimitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetHostedZoneLimit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHostedZoneLimit, reflect.TypeOf((*MockRoute53Client)(nil).GetHostedZoneLimit), varargs...)
}

func (m *MockRoute53Client) GetQueryLoggingConfig(arg0 context.Context, arg1 *route53.GetQueryLoggingConfigInput, arg2 ...func(*route53.Options)) (*route53.GetQueryLoggingConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueryLoggingConfig, varargs...)
	ret0, _ := ret[0].(*route53.GetQueryLoggingConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetQueryLoggingConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueryLoggingConfig, reflect.TypeOf((*MockRoute53Client)(nil).GetQueryLoggingConfig), varargs...)
}

func (m *MockRoute53Client) GetReusableDelegationSet(arg0 context.Context, arg1 *route53.GetReusableDelegationSetInput, arg2 ...func(*route53.Options)) (*route53.GetReusableDelegationSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetReusableDelegationSet, varargs...)
	ret0, _ := ret[0].(*route53.GetReusableDelegationSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetReusableDelegationSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetReusableDelegationSet, reflect.TypeOf((*MockRoute53Client)(nil).GetReusableDelegationSet), varargs...)
}

func (m *MockRoute53Client) GetReusableDelegationSetLimit(arg0 context.Context, arg1 *route53.GetReusableDelegationSetLimitInput, arg2 ...func(*route53.Options)) (*route53.GetReusableDelegationSetLimitOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetReusableDelegationSetLimit, varargs...)
	ret0, _ := ret[0].(*route53.GetReusableDelegationSetLimitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetReusableDelegationSetLimit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetReusableDelegationSetLimit, reflect.TypeOf((*MockRoute53Client)(nil).GetReusableDelegationSetLimit), varargs...)
}

func (m *MockRoute53Client) GetTrafficPolicy(arg0 context.Context, arg1 *route53.GetTrafficPolicyInput, arg2 ...func(*route53.Options)) (*route53.GetTrafficPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTrafficPolicy, varargs...)
	ret0, _ := ret[0].(*route53.GetTrafficPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetTrafficPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTrafficPolicy, reflect.TypeOf((*MockRoute53Client)(nil).GetTrafficPolicy), varargs...)
}

func (m *MockRoute53Client) GetTrafficPolicyInstance(arg0 context.Context, arg1 *route53.GetTrafficPolicyInstanceInput, arg2 ...func(*route53.Options)) (*route53.GetTrafficPolicyInstanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTrafficPolicyInstance, varargs...)
	ret0, _ := ret[0].(*route53.GetTrafficPolicyInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetTrafficPolicyInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTrafficPolicyInstance, reflect.TypeOf((*MockRoute53Client)(nil).GetTrafficPolicyInstance), varargs...)
}

func (m *MockRoute53Client) GetTrafficPolicyInstanceCount(arg0 context.Context, arg1 *route53.GetTrafficPolicyInstanceCountInput, arg2 ...func(*route53.Options)) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTrafficPolicyInstanceCount, varargs...)
	ret0, _ := ret[0].(*route53.GetTrafficPolicyInstanceCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) GetTrafficPolicyInstanceCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTrafficPolicyInstanceCount, reflect.TypeOf((*MockRoute53Client)(nil).GetTrafficPolicyInstanceCount), varargs...)
}

func (m *MockRoute53Client) ListCidrBlocks(arg0 context.Context, arg1 *route53.ListCidrBlocksInput, arg2 ...func(*route53.Options)) (*route53.ListCidrBlocksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCidrBlocks, varargs...)
	ret0, _ := ret[0].(*route53.ListCidrBlocksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListCidrBlocks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCidrBlocks, reflect.TypeOf((*MockRoute53Client)(nil).ListCidrBlocks), varargs...)
}

func (m *MockRoute53Client) ListCidrCollections(arg0 context.Context, arg1 *route53.ListCidrCollectionsInput, arg2 ...func(*route53.Options)) (*route53.ListCidrCollectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCidrCollections, varargs...)
	ret0, _ := ret[0].(*route53.ListCidrCollectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListCidrCollections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCidrCollections, reflect.TypeOf((*MockRoute53Client)(nil).ListCidrCollections), varargs...)
}

func (m *MockRoute53Client) ListCidrLocations(arg0 context.Context, arg1 *route53.ListCidrLocationsInput, arg2 ...func(*route53.Options)) (*route53.ListCidrLocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCidrLocations, varargs...)
	ret0, _ := ret[0].(*route53.ListCidrLocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListCidrLocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCidrLocations, reflect.TypeOf((*MockRoute53Client)(nil).ListCidrLocations), varargs...)
}

func (m *MockRoute53Client) ListGeoLocations(arg0 context.Context, arg1 *route53.ListGeoLocationsInput, arg2 ...func(*route53.Options)) (*route53.ListGeoLocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGeoLocations, varargs...)
	ret0, _ := ret[0].(*route53.ListGeoLocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListGeoLocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGeoLocations, reflect.TypeOf((*MockRoute53Client)(nil).ListGeoLocations), varargs...)
}

func (m *MockRoute53Client) ListHealthChecks(arg0 context.Context, arg1 *route53.ListHealthChecksInput, arg2 ...func(*route53.Options)) (*route53.ListHealthChecksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHealthChecks, varargs...)
	ret0, _ := ret[0].(*route53.ListHealthChecksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListHealthChecks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHealthChecks, reflect.TypeOf((*MockRoute53Client)(nil).ListHealthChecks), varargs...)
}

func (m *MockRoute53Client) ListHostedZones(arg0 context.Context, arg1 *route53.ListHostedZonesInput, arg2 ...func(*route53.Options)) (*route53.ListHostedZonesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHostedZones, varargs...)
	ret0, _ := ret[0].(*route53.ListHostedZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListHostedZones(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHostedZones, reflect.TypeOf((*MockRoute53Client)(nil).ListHostedZones), varargs...)
}

func (m *MockRoute53Client) ListHostedZonesByName(arg0 context.Context, arg1 *route53.ListHostedZonesByNameInput, arg2 ...func(*route53.Options)) (*route53.ListHostedZonesByNameOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHostedZonesByName, varargs...)
	ret0, _ := ret[0].(*route53.ListHostedZonesByNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListHostedZonesByName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHostedZonesByName, reflect.TypeOf((*MockRoute53Client)(nil).ListHostedZonesByName), varargs...)
}

func (m *MockRoute53Client) ListHostedZonesByVPC(arg0 context.Context, arg1 *route53.ListHostedZonesByVPCInput, arg2 ...func(*route53.Options)) (*route53.ListHostedZonesByVPCOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHostedZonesByVPC, varargs...)
	ret0, _ := ret[0].(*route53.ListHostedZonesByVPCOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListHostedZonesByVPC(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHostedZonesByVPC, reflect.TypeOf((*MockRoute53Client)(nil).ListHostedZonesByVPC), varargs...)
}

func (m *MockRoute53Client) ListQueryLoggingConfigs(arg0 context.Context, arg1 *route53.ListQueryLoggingConfigsInput, arg2 ...func(*route53.Options)) (*route53.ListQueryLoggingConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListQueryLoggingConfigs, varargs...)
	ret0, _ := ret[0].(*route53.ListQueryLoggingConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListQueryLoggingConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListQueryLoggingConfigs, reflect.TypeOf((*MockRoute53Client)(nil).ListQueryLoggingConfigs), varargs...)
}

func (m *MockRoute53Client) ListResourceRecordSets(arg0 context.Context, arg1 *route53.ListResourceRecordSetsInput, arg2 ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceRecordSets, varargs...)
	ret0, _ := ret[0].(*route53.ListResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListResourceRecordSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceRecordSets, reflect.TypeOf((*MockRoute53Client)(nil).ListResourceRecordSets), varargs...)
}

func (m *MockRoute53Client) ListReusableDelegationSets(arg0 context.Context, arg1 *route53.ListReusableDelegationSetsInput, arg2 ...func(*route53.Options)) (*route53.ListReusableDelegationSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListReusableDelegationSets, varargs...)
	ret0, _ := ret[0].(*route53.ListReusableDelegationSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListReusableDelegationSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReusableDelegationSets, reflect.TypeOf((*MockRoute53Client)(nil).ListReusableDelegationSets), varargs...)
}

func (m *MockRoute53Client) ListTagsForResource(arg0 context.Context, arg1 *route53.ListTagsForResourceInput, arg2 ...func(*route53.Options)) (*route53.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*route53.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockRoute53Client)(nil).ListTagsForResource), varargs...)
}

func (m *MockRoute53Client) ListTagsForResources(arg0 context.Context, arg1 *route53.ListTagsForResourcesInput, arg2 ...func(*route53.Options)) (*route53.ListTagsForResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResources, varargs...)
	ret0, _ := ret[0].(*route53.ListTagsForResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTagsForResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResources, reflect.TypeOf((*MockRoute53Client)(nil).ListTagsForResources), varargs...)
}

func (m *MockRoute53Client) ListTrafficPolicies(arg0 context.Context, arg1 *route53.ListTrafficPoliciesInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrafficPolicies, varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrafficPolicies, reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicies), varargs...)
}

func (m *MockRoute53Client) ListTrafficPolicyInstances(arg0 context.Context, arg1 *route53.ListTrafficPolicyInstancesInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrafficPolicyInstances, varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPolicyInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicyInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrafficPolicyInstances, reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicyInstances), varargs...)
}

func (m *MockRoute53Client) ListTrafficPolicyInstancesByHostedZone(arg0 context.Context, arg1 *route53.ListTrafficPolicyInstancesByHostedZoneInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrafficPolicyInstancesByHostedZone, varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPolicyInstancesByHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicyInstancesByHostedZone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrafficPolicyInstancesByHostedZone, reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicyInstancesByHostedZone), varargs...)
}

func (m *MockRoute53Client) ListTrafficPolicyInstancesByPolicy(arg0 context.Context, arg1 *route53.ListTrafficPolicyInstancesByPolicyInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrafficPolicyInstancesByPolicy, varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPolicyInstancesByPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicyInstancesByPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrafficPolicyInstancesByPolicy, reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicyInstancesByPolicy), varargs...)
}

func (m *MockRoute53Client) ListTrafficPolicyVersions(arg0 context.Context, arg1 *route53.ListTrafficPolicyVersionsInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPolicyVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrafficPolicyVersions, varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPolicyVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicyVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrafficPolicyVersions, reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicyVersions), varargs...)
}

func (m *MockRoute53Client) ListVPCAssociationAuthorizations(arg0 context.Context, arg1 *route53.ListVPCAssociationAuthorizationsInput, arg2 ...func(*route53.Options)) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVPCAssociationAuthorizations, varargs...)
	ret0, _ := ret[0].(*route53.ListVPCAssociationAuthorizationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53ClientMockRecorder) ListVPCAssociationAuthorizations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVPCAssociationAuthorizations, reflect.TypeOf((*MockRoute53Client)(nil).ListVPCAssociationAuthorizations), varargs...)
}
