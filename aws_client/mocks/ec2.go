package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

type MockEc2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockEc2ClientMockRecorder
}

type MockEc2ClientMockRecorder struct {
	mock *MockEc2Client
}

func NewMockEc2Client(ctrl *gomock.Controller) *MockEc2Client {
	mock := &MockEc2Client{ctrl: ctrl}
	mock.recorder = &MockEc2ClientMockRecorder{mock}
	return mock
}

func (m *MockEc2Client) EXPECT() *MockEc2ClientMockRecorder {
	return m.recorder
}

func (m *MockEc2Client) DescribeAccountAttributes(arg0 context.Context, arg1 *ec2.DescribeAccountAttributesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAttributes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAccountAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAttributes, reflect.TypeOf((*MockEc2Client)(nil).DescribeAccountAttributes), varargs...)
}

func (m *MockEc2Client) DescribeAddressTransfers(arg0 context.Context, arg1 *ec2.DescribeAddressTransfersInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAddressTransfersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAddressTransfers, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAddressTransfersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAddressTransfers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAddressTransfers, reflect.TypeOf((*MockEc2Client)(nil).DescribeAddressTransfers), varargs...)
}

func (m *MockEc2Client) DescribeAddresses(arg0 context.Context, arg1 *ec2.DescribeAddressesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAddresses, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAddresses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAddresses, reflect.TypeOf((*MockEc2Client)(nil).DescribeAddresses), varargs...)
}

func (m *MockEc2Client) DescribeAddressesAttribute(arg0 context.Context, arg1 *ec2.DescribeAddressesAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAddressesAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAddressesAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAddressesAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAddressesAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAddressesAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeAddressesAttribute), varargs...)
}

func (m *MockEc2Client) DescribeAggregateIdFormat(arg0 context.Context, arg1 *ec2.DescribeAggregateIdFormatInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAggregateIdFormatOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAggregateIdFormat, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAggregateIdFormatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAggregateIdFormat(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAggregateIdFormat, reflect.TypeOf((*MockEc2Client)(nil).DescribeAggregateIdFormat), varargs...)
}

func (m *MockEc2Client) DescribeAvailabilityZones(arg0 context.Context, arg1 *ec2.DescribeAvailabilityZonesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAvailabilityZones, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAvailabilityZones(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAvailabilityZones, reflect.TypeOf((*MockEc2Client)(nil).DescribeAvailabilityZones), varargs...)
}

func (m *MockEc2Client) DescribeAwsNetworkPerformanceMetricSubscriptions(arg0 context.Context, arg1 *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAwsNetworkPerformanceMetricSubscriptions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAwsNetworkPerformanceMetricSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAwsNetworkPerformanceMetricSubscriptions, reflect.TypeOf((*MockEc2Client)(nil).DescribeAwsNetworkPerformanceMetricSubscriptions), varargs...)
}

func (m *MockEc2Client) DescribeBundleTasks(arg0 context.Context, arg1 *ec2.DescribeBundleTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeBundleTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBundleTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeBundleTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeBundleTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBundleTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeBundleTasks), varargs...)
}

func (m *MockEc2Client) DescribeByoipCidrs(arg0 context.Context, arg1 *ec2.DescribeByoipCidrsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeByoipCidrs, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeByoipCidrsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeByoipCidrs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeByoipCidrs, reflect.TypeOf((*MockEc2Client)(nil).DescribeByoipCidrs), varargs...)
}

func (m *MockEc2Client) DescribeCapacityReservationFleets(arg0 context.Context, arg1 *ec2.DescribeCapacityReservationFleetsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCapacityReservationFleets, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeCapacityReservationFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeCapacityReservationFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCapacityReservationFleets, reflect.TypeOf((*MockEc2Client)(nil).DescribeCapacityReservationFleets), varargs...)
}

func (m *MockEc2Client) DescribeCapacityReservations(arg0 context.Context, arg1 *ec2.DescribeCapacityReservationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCapacityReservations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeCapacityReservationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeCapacityReservations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCapacityReservations, reflect.TypeOf((*MockEc2Client)(nil).DescribeCapacityReservations), varargs...)
}

func (m *MockEc2Client) DescribeCarrierGateways(arg0 context.Context, arg1 *ec2.DescribeCarrierGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeCarrierGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCarrierGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeCarrierGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeCarrierGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCarrierGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeCarrierGateways), varargs...)
}

func (m *MockEc2Client) DescribeClassicLinkInstances(arg0 context.Context, arg1 *ec2.DescribeClassicLinkInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClassicLinkInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeClassicLinkInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeClassicLinkInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClassicLinkInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeClassicLinkInstances), varargs...)
}

func (m *MockEc2Client) DescribeClientVpnAuthorizationRules(arg0 context.Context, arg1 *ec2.DescribeClientVpnAuthorizationRulesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientVpnAuthorizationRules, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeClientVpnAuthorizationRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeClientVpnAuthorizationRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientVpnAuthorizationRules, reflect.TypeOf((*MockEc2Client)(nil).DescribeClientVpnAuthorizationRules), varargs...)
}

func (m *MockEc2Client) DescribeClientVpnConnections(arg0 context.Context, arg1 *ec2.DescribeClientVpnConnectionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientVpnConnections, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeClientVpnConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeClientVpnConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientVpnConnections, reflect.TypeOf((*MockEc2Client)(nil).DescribeClientVpnConnections), varargs...)
}

func (m *MockEc2Client) DescribeClientVpnEndpoints(arg0 context.Context, arg1 *ec2.DescribeClientVpnEndpointsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientVpnEndpoints, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeClientVpnEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeClientVpnEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientVpnEndpoints, reflect.TypeOf((*MockEc2Client)(nil).DescribeClientVpnEndpoints), varargs...)
}

func (m *MockEc2Client) DescribeClientVpnRoutes(arg0 context.Context, arg1 *ec2.DescribeClientVpnRoutesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeClientVpnRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientVpnRoutes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeClientVpnRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeClientVpnRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientVpnRoutes, reflect.TypeOf((*MockEc2Client)(nil).DescribeClientVpnRoutes), varargs...)
}

func (m *MockEc2Client) DescribeClientVpnTargetNetworks(arg0 context.Context, arg1 *ec2.DescribeClientVpnTargetNetworksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClientVpnTargetNetworks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeClientVpnTargetNetworksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeClientVpnTargetNetworks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClientVpnTargetNetworks, reflect.TypeOf((*MockEc2Client)(nil).DescribeClientVpnTargetNetworks), varargs...)
}

func (m *MockEc2Client) DescribeCoipPools(arg0 context.Context, arg1 *ec2.DescribeCoipPoolsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeCoipPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCoipPools, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeCoipPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeCoipPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCoipPools, reflect.TypeOf((*MockEc2Client)(nil).DescribeCoipPools), varargs...)
}

func (m *MockEc2Client) DescribeConversionTasks(arg0 context.Context, arg1 *ec2.DescribeConversionTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeConversionTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConversionTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeConversionTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeConversionTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConversionTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeConversionTasks), varargs...)
}

func (m *MockEc2Client) DescribeCustomerGateways(arg0 context.Context, arg1 *ec2.DescribeCustomerGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCustomerGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeCustomerGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeCustomerGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCustomerGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeCustomerGateways), varargs...)
}

func (m *MockEc2Client) DescribeDhcpOptions(arg0 context.Context, arg1 *ec2.DescribeDhcpOptionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeDhcpOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDhcpOptions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeDhcpOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeDhcpOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDhcpOptions, reflect.TypeOf((*MockEc2Client)(nil).DescribeDhcpOptions), varargs...)
}

func (m *MockEc2Client) DescribeEgressOnlyInternetGateways(arg0 context.Context, arg1 *ec2.DescribeEgressOnlyInternetGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEgressOnlyInternetGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeEgressOnlyInternetGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeEgressOnlyInternetGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEgressOnlyInternetGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeEgressOnlyInternetGateways), varargs...)
}

func (m *MockEc2Client) DescribeElasticGpus(arg0 context.Context, arg1 *ec2.DescribeElasticGpusInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeElasticGpusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeElasticGpus, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeElasticGpusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeElasticGpus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeElasticGpus, reflect.TypeOf((*MockEc2Client)(nil).DescribeElasticGpus), varargs...)
}

func (m *MockEc2Client) DescribeExportImageTasks(arg0 context.Context, arg1 *ec2.DescribeExportImageTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeExportImageTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExportImageTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeExportImageTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeExportImageTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExportImageTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeExportImageTasks), varargs...)
}

func (m *MockEc2Client) DescribeExportTasks(arg0 context.Context, arg1 *ec2.DescribeExportTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeExportTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExportTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeExportTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeExportTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExportTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeExportTasks), varargs...)
}

func (m *MockEc2Client) DescribeFastLaunchImages(arg0 context.Context, arg1 *ec2.DescribeFastLaunchImagesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFastLaunchImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFastLaunchImages, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFastLaunchImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFastLaunchImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFastLaunchImages, reflect.TypeOf((*MockEc2Client)(nil).DescribeFastLaunchImages), varargs...)
}

func (m *MockEc2Client) DescribeFastSnapshotRestores(arg0 context.Context, arg1 *ec2.DescribeFastSnapshotRestoresInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFastSnapshotRestores, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFastSnapshotRestoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFastSnapshotRestores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFastSnapshotRestores, reflect.TypeOf((*MockEc2Client)(nil).DescribeFastSnapshotRestores), varargs...)
}

func (m *MockEc2Client) DescribeFleetHistory(arg0 context.Context, arg1 *ec2.DescribeFleetHistoryInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFleetHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetHistory, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFleetHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFleetHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetHistory, reflect.TypeOf((*MockEc2Client)(nil).DescribeFleetHistory), varargs...)
}

func (m *MockEc2Client) DescribeFleetInstances(arg0 context.Context, arg1 *ec2.DescribeFleetInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFleetInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFleetInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFleetInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeFleetInstances), varargs...)
}

func (m *MockEc2Client) DescribeFleets(arg0 context.Context, arg1 *ec2.DescribeFleetsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleets, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleets, reflect.TypeOf((*MockEc2Client)(nil).DescribeFleets), varargs...)
}

func (m *MockEc2Client) DescribeFlowLogs(arg0 context.Context, arg1 *ec2.DescribeFlowLogsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFlowLogs, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFlowLogsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFlowLogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFlowLogs, reflect.TypeOf((*MockEc2Client)(nil).DescribeFlowLogs), varargs...)
}

func (m *MockEc2Client) DescribeFpgaImageAttribute(arg0 context.Context, arg1 *ec2.DescribeFpgaImageAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFpgaImageAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFpgaImageAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFpgaImageAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFpgaImageAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeFpgaImageAttribute), varargs...)
}

func (m *MockEc2Client) DescribeFpgaImages(arg0 context.Context, arg1 *ec2.DescribeFpgaImagesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFpgaImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFpgaImages, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFpgaImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFpgaImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFpgaImages, reflect.TypeOf((*MockEc2Client)(nil).DescribeFpgaImages), varargs...)
}

func (m *MockEc2Client) DescribeHostReservationOfferings(arg0 context.Context, arg1 *ec2.DescribeHostReservationOfferingsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHostReservationOfferings, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeHostReservationOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeHostReservationOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHostReservationOfferings, reflect.TypeOf((*MockEc2Client)(nil).DescribeHostReservationOfferings), varargs...)
}

func (m *MockEc2Client) DescribeHostReservations(arg0 context.Context, arg1 *ec2.DescribeHostReservationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeHostReservationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHostReservations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeHostReservationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeHostReservations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHostReservations, reflect.TypeOf((*MockEc2Client)(nil).DescribeHostReservations), varargs...)
}

func (m *MockEc2Client) DescribeHosts(arg0 context.Context, arg1 *ec2.DescribeHostsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHosts, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeHostsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeHosts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHosts, reflect.TypeOf((*MockEc2Client)(nil).DescribeHosts), varargs...)
}

func (m *MockEc2Client) DescribeIamInstanceProfileAssociations(arg0 context.Context, arg1 *ec2.DescribeIamInstanceProfileAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIamInstanceProfileAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIamInstanceProfileAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIamInstanceProfileAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIamInstanceProfileAssociations, reflect.TypeOf((*MockEc2Client)(nil).DescribeIamInstanceProfileAssociations), varargs...)
}

func (m *MockEc2Client) DescribeIdFormat(arg0 context.Context, arg1 *ec2.DescribeIdFormatInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIdFormatOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIdFormat, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIdFormatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIdFormat(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIdFormat, reflect.TypeOf((*MockEc2Client)(nil).DescribeIdFormat), varargs...)
}

func (m *MockEc2Client) DescribeIdentityIdFormat(arg0 context.Context, arg1 *ec2.DescribeIdentityIdFormatInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIdentityIdFormatOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIdentityIdFormat, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIdentityIdFormatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIdentityIdFormat(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIdentityIdFormat, reflect.TypeOf((*MockEc2Client)(nil).DescribeIdentityIdFormat), varargs...)
}

func (m *MockEc2Client) DescribeImageAttribute(arg0 context.Context, arg1 *ec2.DescribeImageAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImageAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeImageAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeImageAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImageAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeImageAttribute), varargs...)
}

func (m *MockEc2Client) DescribeImages(arg0 context.Context, arg1 *ec2.DescribeImagesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImages, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImages, reflect.TypeOf((*MockEc2Client)(nil).DescribeImages), varargs...)
}

func (m *MockEc2Client) DescribeImportImageTasks(arg0 context.Context, arg1 *ec2.DescribeImportImageTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeImportImageTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImportImageTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeImportImageTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeImportImageTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImportImageTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeImportImageTasks), varargs...)
}

func (m *MockEc2Client) DescribeImportSnapshotTasks(arg0 context.Context, arg1 *ec2.DescribeImportSnapshotTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImportSnapshotTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeImportSnapshotTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeImportSnapshotTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImportSnapshotTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeImportSnapshotTasks), varargs...)
}

func (m *MockEc2Client) DescribeInstanceAttribute(arg0 context.Context, arg1 *ec2.DescribeInstanceAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceAttribute), varargs...)
}

func (m *MockEc2Client) DescribeInstanceCreditSpecifications(arg0 context.Context, arg1 *ec2.DescribeInstanceCreditSpecificationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceCreditSpecifications, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceCreditSpecificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceCreditSpecifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceCreditSpecifications, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceCreditSpecifications), varargs...)
}

func (m *MockEc2Client) DescribeInstanceEventNotificationAttributes(arg0 context.Context, arg1 *ec2.DescribeInstanceEventNotificationAttributesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceEventNotificationAttributes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceEventNotificationAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceEventNotificationAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceEventNotificationAttributes, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceEventNotificationAttributes), varargs...)
}

func (m *MockEc2Client) DescribeInstanceEventWindows(arg0 context.Context, arg1 *ec2.DescribeInstanceEventWindowsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceEventWindows, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceEventWindowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceEventWindows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceEventWindows, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceEventWindows), varargs...)
}

func (m *MockEc2Client) DescribeInstanceStatus(arg0 context.Context, arg1 *ec2.DescribeInstanceStatusInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceStatus, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceStatus, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceStatus), varargs...)
}

func (m *MockEc2Client) DescribeInstanceTypeOfferings(arg0 context.Context, arg1 *ec2.DescribeInstanceTypeOfferingsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceTypeOfferings, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypeOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceTypeOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceTypeOfferings, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceTypeOfferings), varargs...)
}

func (m *MockEc2Client) DescribeInstanceTypes(arg0 context.Context, arg1 *ec2.DescribeInstanceTypesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceTypes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceTypes, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceTypes), varargs...)
}

func (m *MockEc2Client) DescribeInstances(arg0 context.Context, arg1 *ec2.DescribeInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeInstances), varargs...)
}

func (m *MockEc2Client) DescribeInternetGateways(arg0 context.Context, arg1 *ec2.DescribeInternetGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInternetGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInternetGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInternetGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInternetGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeInternetGateways), varargs...)
}

func (m *MockEc2Client) DescribeIpamPools(arg0 context.Context, arg1 *ec2.DescribeIpamPoolsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIpamPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIpamPools, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIpamPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIpamPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIpamPools, reflect.TypeOf((*MockEc2Client)(nil).DescribeIpamPools), varargs...)
}

func (m *MockEc2Client) DescribeIpamScopes(arg0 context.Context, arg1 *ec2.DescribeIpamScopesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIpamScopesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIpamScopes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIpamScopesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIpamScopes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIpamScopes, reflect.TypeOf((*MockEc2Client)(nil).DescribeIpamScopes), varargs...)
}

func (m *MockEc2Client) DescribeIpams(arg0 context.Context, arg1 *ec2.DescribeIpamsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIpamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIpams, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIpamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIpams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIpams, reflect.TypeOf((*MockEc2Client)(nil).DescribeIpams), varargs...)
}

func (m *MockEc2Client) DescribeIpv6Pools(arg0 context.Context, arg1 *ec2.DescribeIpv6PoolsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeIpv6PoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIpvPools, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeIpv6PoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeIpv6Pools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIpvPools, reflect.TypeOf((*MockEc2Client)(nil).DescribeIpv6Pools), varargs...)
}

func (m *MockEc2Client) DescribeKeyPairs(arg0 context.Context, arg1 *ec2.DescribeKeyPairsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeKeyPairs, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeKeyPairsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeKeyPairs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeKeyPairs, reflect.TypeOf((*MockEc2Client)(nil).DescribeKeyPairs), varargs...)
}

func (m *MockEc2Client) DescribeLaunchTemplateVersions(arg0 context.Context, arg1 *ec2.DescribeLaunchTemplateVersionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLaunchTemplateVersions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLaunchTemplateVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLaunchTemplateVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLaunchTemplateVersions, reflect.TypeOf((*MockEc2Client)(nil).DescribeLaunchTemplateVersions), varargs...)
}

func (m *MockEc2Client) DescribeLaunchTemplates(arg0 context.Context, arg1 *ec2.DescribeLaunchTemplatesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLaunchTemplates, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLaunchTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLaunchTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLaunchTemplates, reflect.TypeOf((*MockEc2Client)(nil).DescribeLaunchTemplates), varargs...)
}

func (m *MockEc2Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(arg0 context.Context, arg1 *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations, reflect.TypeOf((*MockEc2Client)(nil).DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations), varargs...)
}

func (m *MockEc2Client) DescribeLocalGatewayRouteTableVpcAssociations(arg0 context.Context, arg1 *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocalGatewayRouteTableVpcAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLocalGatewayRouteTableVpcAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocalGatewayRouteTableVpcAssociations, reflect.TypeOf((*MockEc2Client)(nil).DescribeLocalGatewayRouteTableVpcAssociations), varargs...)
}

func (m *MockEc2Client) DescribeLocalGatewayRouteTables(arg0 context.Context, arg1 *ec2.DescribeLocalGatewayRouteTablesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocalGatewayRouteTables, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLocalGatewayRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLocalGatewayRouteTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocalGatewayRouteTables, reflect.TypeOf((*MockEc2Client)(nil).DescribeLocalGatewayRouteTables), varargs...)
}

func (m *MockEc2Client) DescribeLocalGatewayVirtualInterfaceGroups(arg0 context.Context, arg1 *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocalGatewayVirtualInterfaceGroups, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLocalGatewayVirtualInterfaceGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocalGatewayVirtualInterfaceGroups, reflect.TypeOf((*MockEc2Client)(nil).DescribeLocalGatewayVirtualInterfaceGroups), varargs...)
}

func (m *MockEc2Client) DescribeLocalGatewayVirtualInterfaces(arg0 context.Context, arg1 *ec2.DescribeLocalGatewayVirtualInterfacesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocalGatewayVirtualInterfaces, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLocalGatewayVirtualInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLocalGatewayVirtualInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocalGatewayVirtualInterfaces, reflect.TypeOf((*MockEc2Client)(nil).DescribeLocalGatewayVirtualInterfaces), varargs...)
}

func (m *MockEc2Client) DescribeLocalGateways(arg0 context.Context, arg1 *ec2.DescribeLocalGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocalGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeLocalGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeLocalGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocalGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeLocalGateways), varargs...)
}

func (m *MockEc2Client) DescribeManagedPrefixLists(arg0 context.Context, arg1 *ec2.DescribeManagedPrefixListsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeManagedPrefixListsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeManagedPrefixLists, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeManagedPrefixListsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeManagedPrefixLists(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeManagedPrefixLists, reflect.TypeOf((*MockEc2Client)(nil).DescribeManagedPrefixLists), varargs...)
}

func (m *MockEc2Client) DescribeMovingAddresses(arg0 context.Context, arg1 *ec2.DescribeMovingAddressesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeMovingAddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMovingAddresses, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeMovingAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeMovingAddresses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMovingAddresses, reflect.TypeOf((*MockEc2Client)(nil).DescribeMovingAddresses), varargs...)
}

func (m *MockEc2Client) DescribeNatGateways(arg0 context.Context, arg1 *ec2.DescribeNatGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNatGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNatGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNatGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNatGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeNatGateways), varargs...)
}

func (m *MockEc2Client) DescribeNetworkAcls(arg0 context.Context, arg1 *ec2.DescribeNetworkAclsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkAcls, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkAclsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkAcls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkAcls, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkAcls), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInsightsAccessScopeAnalyses(arg0 context.Context, arg1 *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInsightsAccessScopeAnalyses, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInsightsAccessScopeAnalyses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInsightsAccessScopeAnalyses, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInsightsAccessScopeAnalyses), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInsightsAccessScopes(arg0 context.Context, arg1 *ec2.DescribeNetworkInsightsAccessScopesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInsightsAccessScopes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInsightsAccessScopesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInsightsAccessScopes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInsightsAccessScopes, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInsightsAccessScopes), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInsightsAnalyses(arg0 context.Context, arg1 *ec2.DescribeNetworkInsightsAnalysesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInsightsAnalyses, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInsightsAnalysesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInsightsAnalyses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInsightsAnalyses, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInsightsAnalyses), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInsightsPaths(arg0 context.Context, arg1 *ec2.DescribeNetworkInsightsPathsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInsightsPaths, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInsightsPathsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInsightsPaths(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInsightsPaths, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInsightsPaths), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInterfaceAttribute(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfaceAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInterfaceAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInterfaceAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInterfaceAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInterfaceAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInterfaceAttribute), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInterfacePermissions(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfacePermissionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInterfacePermissions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInterfacePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInterfacePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInterfacePermissions, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInterfacePermissions), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInterfaces(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfacesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNetworkInterfaces, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNetworkInterfaces, reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInterfaces), varargs...)
}

func (m *MockEc2Client) DescribePlacementGroups(arg0 context.Context, arg1 *ec2.DescribePlacementGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribePlacementGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePlacementGroups, varargs...)
	ret0, _ := ret[0].(*ec2.DescribePlacementGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribePlacementGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePlacementGroups, reflect.TypeOf((*MockEc2Client)(nil).DescribePlacementGroups), varargs...)
}

func (m *MockEc2Client) DescribePrefixLists(arg0 context.Context, arg1 *ec2.DescribePrefixListsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribePrefixListsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePrefixLists, varargs...)
	ret0, _ := ret[0].(*ec2.DescribePrefixListsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribePrefixLists(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePrefixLists, reflect.TypeOf((*MockEc2Client)(nil).DescribePrefixLists), varargs...)
}

func (m *MockEc2Client) DescribePrincipalIdFormat(arg0 context.Context, arg1 *ec2.DescribePrincipalIdFormatInput, arg2 ...func(*ec2.Options)) (*ec2.DescribePrincipalIdFormatOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePrincipalIdFormat, varargs...)
	ret0, _ := ret[0].(*ec2.DescribePrincipalIdFormatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribePrincipalIdFormat(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePrincipalIdFormat, reflect.TypeOf((*MockEc2Client)(nil).DescribePrincipalIdFormat), varargs...)
}

func (m *MockEc2Client) DescribePublicIpv4Pools(arg0 context.Context, arg1 *ec2.DescribePublicIpv4PoolsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePublicIpvPools, varargs...)
	ret0, _ := ret[0].(*ec2.DescribePublicIpv4PoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribePublicIpv4Pools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePublicIpvPools, reflect.TypeOf((*MockEc2Client)(nil).DescribePublicIpv4Pools), varargs...)
}

func (m *MockEc2Client) DescribeRegions(arg0 context.Context, arg1 *ec2.DescribeRegionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRegions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeRegionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeRegions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRegions, reflect.TypeOf((*MockEc2Client)(nil).DescribeRegions), varargs...)
}

func (m *MockEc2Client) DescribeReplaceRootVolumeTasks(arg0 context.Context, arg1 *ec2.DescribeReplaceRootVolumeTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReplaceRootVolumeTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeReplaceRootVolumeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeReplaceRootVolumeTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReplaceRootVolumeTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeReplaceRootVolumeTasks), varargs...)
}

func (m *MockEc2Client) DescribeReservedInstances(arg0 context.Context, arg1 *ec2.DescribeReservedInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeReservedInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeReservedInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeReservedInstances), varargs...)
}

func (m *MockEc2Client) DescribeReservedInstancesListings(arg0 context.Context, arg1 *ec2.DescribeReservedInstancesListingsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedInstancesListings, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeReservedInstancesListingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeReservedInstancesListings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedInstancesListings, reflect.TypeOf((*MockEc2Client)(nil).DescribeReservedInstancesListings), varargs...)
}

func (m *MockEc2Client) DescribeReservedInstancesModifications(arg0 context.Context, arg1 *ec2.DescribeReservedInstancesModificationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedInstancesModifications, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeReservedInstancesModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeReservedInstancesModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedInstancesModifications, reflect.TypeOf((*MockEc2Client)(nil).DescribeReservedInstancesModifications), varargs...)
}

func (m *MockEc2Client) DescribeReservedInstancesOfferings(arg0 context.Context, arg1 *ec2.DescribeReservedInstancesOfferingsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeReservedInstancesOfferings, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeReservedInstancesOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeReservedInstancesOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeReservedInstancesOfferings, reflect.TypeOf((*MockEc2Client)(nil).DescribeReservedInstancesOfferings), varargs...)
}

func (m *MockEc2Client) DescribeRouteTables(arg0 context.Context, arg1 *ec2.DescribeRouteTablesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRouteTables, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeRouteTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRouteTables, reflect.TypeOf((*MockEc2Client)(nil).DescribeRouteTables), varargs...)
}

func (m *MockEc2Client) DescribeScheduledInstanceAvailability(arg0 context.Context, arg1 *ec2.DescribeScheduledInstanceAvailabilityInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScheduledInstanceAvailability, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeScheduledInstanceAvailabilityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeScheduledInstanceAvailability(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScheduledInstanceAvailability, reflect.TypeOf((*MockEc2Client)(nil).DescribeScheduledInstanceAvailability), varargs...)
}

func (m *MockEc2Client) DescribeScheduledInstances(arg0 context.Context, arg1 *ec2.DescribeScheduledInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScheduledInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeScheduledInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeScheduledInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScheduledInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeScheduledInstances), varargs...)
}

func (m *MockEc2Client) DescribeSecurityGroupReferences(arg0 context.Context, arg1 *ec2.DescribeSecurityGroupReferencesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecurityGroupReferences, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupReferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSecurityGroupReferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecurityGroupReferences, reflect.TypeOf((*MockEc2Client)(nil).DescribeSecurityGroupReferences), varargs...)
}

func (m *MockEc2Client) DescribeSecurityGroupRules(arg0 context.Context, arg1 *ec2.DescribeSecurityGroupRulesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecurityGroupRules, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSecurityGroupRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecurityGroupRules, reflect.TypeOf((*MockEc2Client)(nil).DescribeSecurityGroupRules), varargs...)
}

func (m *MockEc2Client) DescribeSecurityGroups(arg0 context.Context, arg1 *ec2.DescribeSecurityGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecurityGroups, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecurityGroups, reflect.TypeOf((*MockEc2Client)(nil).DescribeSecurityGroups), varargs...)
}

func (m *MockEc2Client) DescribeSnapshotAttribute(arg0 context.Context, arg1 *ec2.DescribeSnapshotAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshotAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSnapshotAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshotAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeSnapshotAttribute), varargs...)
}

func (m *MockEc2Client) DescribeSnapshotTierStatus(arg0 context.Context, arg1 *ec2.DescribeSnapshotTierStatusInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshotTierStatus, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotTierStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSnapshotTierStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshotTierStatus, reflect.TypeOf((*MockEc2Client)(nil).DescribeSnapshotTierStatus), varargs...)
}

func (m *MockEc2Client) DescribeSnapshots(arg0 context.Context, arg1 *ec2.DescribeSnapshotsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSnapshots, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSnapshots, reflect.TypeOf((*MockEc2Client)(nil).DescribeSnapshots), varargs...)
}

func (m *MockEc2Client) DescribeSpotDatafeedSubscription(arg0 context.Context, arg1 *ec2.DescribeSpotDatafeedSubscriptionInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpotDatafeedSubscription, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSpotDatafeedSubscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSpotDatafeedSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpotDatafeedSubscription, reflect.TypeOf((*MockEc2Client)(nil).DescribeSpotDatafeedSubscription), varargs...)
}

func (m *MockEc2Client) DescribeSpotFleetInstances(arg0 context.Context, arg1 *ec2.DescribeSpotFleetInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpotFleetInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSpotFleetInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSpotFleetInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpotFleetInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeSpotFleetInstances), varargs...)
}

func (m *MockEc2Client) DescribeSpotFleetRequestHistory(arg0 context.Context, arg1 *ec2.DescribeSpotFleetRequestHistoryInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpotFleetRequestHistory, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSpotFleetRequestHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSpotFleetRequestHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpotFleetRequestHistory, reflect.TypeOf((*MockEc2Client)(nil).DescribeSpotFleetRequestHistory), varargs...)
}

func (m *MockEc2Client) DescribeSpotFleetRequests(arg0 context.Context, arg1 *ec2.DescribeSpotFleetRequestsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpotFleetRequests, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSpotFleetRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSpotFleetRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpotFleetRequests, reflect.TypeOf((*MockEc2Client)(nil).DescribeSpotFleetRequests), varargs...)
}

func (m *MockEc2Client) DescribeSpotInstanceRequests(arg0 context.Context, arg1 *ec2.DescribeSpotInstanceRequestsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpotInstanceRequests, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSpotInstanceRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSpotInstanceRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpotInstanceRequests, reflect.TypeOf((*MockEc2Client)(nil).DescribeSpotInstanceRequests), varargs...)
}

func (m *MockEc2Client) DescribeSpotPriceHistory(arg0 context.Context, arg1 *ec2.DescribeSpotPriceHistoryInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpotPriceHistory, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSpotPriceHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSpotPriceHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpotPriceHistory, reflect.TypeOf((*MockEc2Client)(nil).DescribeSpotPriceHistory), varargs...)
}

func (m *MockEc2Client) DescribeStaleSecurityGroups(arg0 context.Context, arg1 *ec2.DescribeStaleSecurityGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStaleSecurityGroups, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeStaleSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeStaleSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStaleSecurityGroups, reflect.TypeOf((*MockEc2Client)(nil).DescribeStaleSecurityGroups), varargs...)
}

func (m *MockEc2Client) DescribeStoreImageTasks(arg0 context.Context, arg1 *ec2.DescribeStoreImageTasksInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeStoreImageTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStoreImageTasks, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeStoreImageTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeStoreImageTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStoreImageTasks, reflect.TypeOf((*MockEc2Client)(nil).DescribeStoreImageTasks), varargs...)
}

func (m *MockEc2Client) DescribeSubnets(arg0 context.Context, arg1 *ec2.DescribeSubnetsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSubnets, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSubnets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSubnets, reflect.TypeOf((*MockEc2Client)(nil).DescribeSubnets), varargs...)
}

func (m *MockEc2Client) DescribeTags(arg0 context.Context, arg1 *ec2.DescribeTagsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockEc2Client)(nil).DescribeTags), varargs...)
}

func (m *MockEc2Client) DescribeTrafficMirrorFilters(arg0 context.Context, arg1 *ec2.DescribeTrafficMirrorFiltersInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrafficMirrorFilters, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTrafficMirrorFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTrafficMirrorFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrafficMirrorFilters, reflect.TypeOf((*MockEc2Client)(nil).DescribeTrafficMirrorFilters), varargs...)
}

func (m *MockEc2Client) DescribeTrafficMirrorSessions(arg0 context.Context, arg1 *ec2.DescribeTrafficMirrorSessionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrafficMirrorSessions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTrafficMirrorSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTrafficMirrorSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrafficMirrorSessions, reflect.TypeOf((*MockEc2Client)(nil).DescribeTrafficMirrorSessions), varargs...)
}

func (m *MockEc2Client) DescribeTrafficMirrorTargets(arg0 context.Context, arg1 *ec2.DescribeTrafficMirrorTargetsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrafficMirrorTargets, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTrafficMirrorTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTrafficMirrorTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrafficMirrorTargets, reflect.TypeOf((*MockEc2Client)(nil).DescribeTrafficMirrorTargets), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayAttachments(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayAttachmentsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayAttachments, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayAttachments, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayAttachments), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayConnectPeers(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayConnectPeersInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayConnectPeers, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayConnectPeersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayConnectPeers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayConnectPeers, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayConnectPeers), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayConnects(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayConnectsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayConnects, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayConnectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayConnects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayConnects, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayConnects), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayMulticastDomains(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayMulticastDomainsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayMulticastDomains, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayMulticastDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayMulticastDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayMulticastDomains, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayMulticastDomains), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayPeeringAttachments(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayPeeringAttachmentsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayPeeringAttachments, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayPeeringAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayPeeringAttachments, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayPeeringAttachments), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayPolicyTables(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayPolicyTablesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayPolicyTables, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayPolicyTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayPolicyTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayPolicyTables, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayPolicyTables), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayRouteTableAnnouncements(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayRouteTableAnnouncements, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayRouteTableAnnouncements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayRouteTableAnnouncements, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayRouteTableAnnouncements), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayRouteTables(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayRouteTablesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayRouteTables, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayRouteTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayRouteTables, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayRouteTables), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayVpcAttachments(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayVpcAttachmentsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGatewayVpcAttachments, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayVpcAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayVpcAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGatewayVpcAttachments, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayVpcAttachments), varargs...)
}

func (m *MockEc2Client) DescribeTransitGateways(arg0 context.Context, arg1 *ec2.DescribeTransitGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransitGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransitGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGateways), varargs...)
}

func (m *MockEc2Client) DescribeTrunkInterfaceAssociations(arg0 context.Context, arg1 *ec2.DescribeTrunkInterfaceAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrunkInterfaceAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTrunkInterfaceAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTrunkInterfaceAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrunkInterfaceAssociations, reflect.TypeOf((*MockEc2Client)(nil).DescribeTrunkInterfaceAssociations), varargs...)
}

func (m *MockEc2Client) DescribeVerifiedAccessEndpoints(arg0 context.Context, arg1 *ec2.DescribeVerifiedAccessEndpointsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVerifiedAccessEndpoints, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVerifiedAccessEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVerifiedAccessEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVerifiedAccessEndpoints, reflect.TypeOf((*MockEc2Client)(nil).DescribeVerifiedAccessEndpoints), varargs...)
}

func (m *MockEc2Client) DescribeVerifiedAccessGroups(arg0 context.Context, arg1 *ec2.DescribeVerifiedAccessGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVerifiedAccessGroups, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVerifiedAccessGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVerifiedAccessGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVerifiedAccessGroups, reflect.TypeOf((*MockEc2Client)(nil).DescribeVerifiedAccessGroups), varargs...)
}

func (m *MockEc2Client) DescribeVerifiedAccessInstanceLoggingConfigurations(arg0 context.Context, arg1 *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVerifiedAccessInstanceLoggingConfigurations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVerifiedAccessInstanceLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVerifiedAccessInstanceLoggingConfigurations, reflect.TypeOf((*MockEc2Client)(nil).DescribeVerifiedAccessInstanceLoggingConfigurations), varargs...)
}

func (m *MockEc2Client) DescribeVerifiedAccessInstances(arg0 context.Context, arg1 *ec2.DescribeVerifiedAccessInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVerifiedAccessInstances, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVerifiedAccessInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVerifiedAccessInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVerifiedAccessInstances, reflect.TypeOf((*MockEc2Client)(nil).DescribeVerifiedAccessInstances), varargs...)
}

func (m *MockEc2Client) DescribeVerifiedAccessTrustProviders(arg0 context.Context, arg1 *ec2.DescribeVerifiedAccessTrustProvidersInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVerifiedAccessTrustProviders, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVerifiedAccessTrustProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVerifiedAccessTrustProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVerifiedAccessTrustProviders, reflect.TypeOf((*MockEc2Client)(nil).DescribeVerifiedAccessTrustProviders), varargs...)
}

func (m *MockEc2Client) DescribeVolumeAttribute(arg0 context.Context, arg1 *ec2.DescribeVolumeAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVolumeAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVolumeAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumeAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVolumeAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVolumeAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeVolumeAttribute), varargs...)
}

func (m *MockEc2Client) DescribeVolumeStatus(arg0 context.Context, arg1 *ec2.DescribeVolumeStatusInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVolumeStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVolumeStatus, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumeStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVolumeStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVolumeStatus, reflect.TypeOf((*MockEc2Client)(nil).DescribeVolumeStatus), varargs...)
}

func (m *MockEc2Client) DescribeVolumes(arg0 context.Context, arg1 *ec2.DescribeVolumesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVolumes, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVolumes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVolumes, reflect.TypeOf((*MockEc2Client)(nil).DescribeVolumes), varargs...)
}

func (m *MockEc2Client) DescribeVolumesModifications(arg0 context.Context, arg1 *ec2.DescribeVolumesModificationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVolumesModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVolumesModifications, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumesModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVolumesModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVolumesModifications, reflect.TypeOf((*MockEc2Client)(nil).DescribeVolumesModifications), varargs...)
}

func (m *MockEc2Client) DescribeVpcAttribute(arg0 context.Context, arg1 *ec2.DescribeVpcAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcAttribute, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcAttribute, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcAttribute), varargs...)
}

func (m *MockEc2Client) DescribeVpcClassicLink(arg0 context.Context, arg1 *ec2.DescribeVpcClassicLinkInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcClassicLink, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcClassicLinkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcClassicLink(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcClassicLink, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcClassicLink), varargs...)
}

func (m *MockEc2Client) DescribeVpcClassicLinkDnsSupport(arg0 context.Context, arg1 *ec2.DescribeVpcClassicLinkDnsSupportInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcClassicLinkDnsSupport, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcClassicLinkDnsSupportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcClassicLinkDnsSupport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcClassicLinkDnsSupport, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcClassicLinkDnsSupport), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointConnectionNotifications(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointConnectionNotificationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpointConnectionNotifications, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointConnectionNotificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointConnectionNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpointConnectionNotifications, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointConnectionNotifications), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointConnections(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointConnectionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpointConnections, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpointConnections, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointConnections), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointServiceConfigurations(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointServiceConfigurationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpointServiceConfigurations, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServiceConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointServiceConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpointServiceConfigurations, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointServiceConfigurations), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointServicePermissions(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointServicePermissionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpointServicePermissions, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServicePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointServicePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpointServicePermissions, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointServicePermissions), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointServices(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointServicesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpointServices, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpointServices, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointServices), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpoints(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcEndpoints, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcEndpoints, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpoints), varargs...)
}

func (m *MockEc2Client) DescribeVpcPeeringConnections(arg0 context.Context, arg1 *ec2.DescribeVpcPeeringConnectionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcPeeringConnections, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcPeeringConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcPeeringConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcPeeringConnections, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcPeeringConnections), varargs...)
}

func (m *MockEc2Client) DescribeVpcs(arg0 context.Context, arg1 *ec2.DescribeVpcsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpcs, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpcs, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcs), varargs...)
}

func (m *MockEc2Client) DescribeVpnConnections(arg0 context.Context, arg1 *ec2.DescribeVpnConnectionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpnConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpnConnections, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpnConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpnConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpnConnections, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpnConnections), varargs...)
}

func (m *MockEc2Client) DescribeVpnGateways(arg0 context.Context, arg1 *ec2.DescribeVpnGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVpnGateways, varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpnGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpnGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVpnGateways, reflect.TypeOf((*MockEc2Client)(nil).DescribeVpnGateways), varargs...)
}

func (m *MockEc2Client) GetAssociatedEnclaveCertificateIamRoles(arg0 context.Context, arg1 *ec2.GetAssociatedEnclaveCertificateIamRolesInput, arg2 ...func(*ec2.Options)) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAssociatedEnclaveCertificateIamRoles, varargs...)
	ret0, _ := ret[0].(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetAssociatedEnclaveCertificateIamRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAssociatedEnclaveCertificateIamRoles, reflect.TypeOf((*MockEc2Client)(nil).GetAssociatedEnclaveCertificateIamRoles), varargs...)
}

func (m *MockEc2Client) GetAssociatedIpv6PoolCidrs(arg0 context.Context, arg1 *ec2.GetAssociatedIpv6PoolCidrsInput, arg2 ...func(*ec2.Options)) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAssociatedIpvPoolCidrs, varargs...)
	ret0, _ := ret[0].(*ec2.GetAssociatedIpv6PoolCidrsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetAssociatedIpv6PoolCidrs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAssociatedIpvPoolCidrs, reflect.TypeOf((*MockEc2Client)(nil).GetAssociatedIpv6PoolCidrs), varargs...)
}

func (m *MockEc2Client) GetAwsNetworkPerformanceData(arg0 context.Context, arg1 *ec2.GetAwsNetworkPerformanceDataInput, arg2 ...func(*ec2.Options)) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAwsNetworkPerformanceData, varargs...)
	ret0, _ := ret[0].(*ec2.GetAwsNetworkPerformanceDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetAwsNetworkPerformanceData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAwsNetworkPerformanceData, reflect.TypeOf((*MockEc2Client)(nil).GetAwsNetworkPerformanceData), varargs...)
}

func (m *MockEc2Client) GetCapacityReservationUsage(arg0 context.Context, arg1 *ec2.GetCapacityReservationUsageInput, arg2 ...func(*ec2.Options)) (*ec2.GetCapacityReservationUsageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCapacityReservationUsage, varargs...)
	ret0, _ := ret[0].(*ec2.GetCapacityReservationUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetCapacityReservationUsage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCapacityReservationUsage, reflect.TypeOf((*MockEc2Client)(nil).GetCapacityReservationUsage), varargs...)
}

func (m *MockEc2Client) GetCoipPoolUsage(arg0 context.Context, arg1 *ec2.GetCoipPoolUsageInput, arg2 ...func(*ec2.Options)) (*ec2.GetCoipPoolUsageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCoipPoolUsage, varargs...)
	ret0, _ := ret[0].(*ec2.GetCoipPoolUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetCoipPoolUsage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCoipPoolUsage, reflect.TypeOf((*MockEc2Client)(nil).GetCoipPoolUsage), varargs...)
}

func (m *MockEc2Client) GetConsoleOutput(arg0 context.Context, arg1 *ec2.GetConsoleOutputInput, arg2 ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConsoleOutput, varargs...)
	ret0, _ := ret[0].(*ec2.GetConsoleOutputOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetConsoleOutput(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConsoleOutput, reflect.TypeOf((*MockEc2Client)(nil).GetConsoleOutput), varargs...)
}

func (m *MockEc2Client) GetConsoleScreenshot(arg0 context.Context, arg1 *ec2.GetConsoleScreenshotInput, arg2 ...func(*ec2.Options)) (*ec2.GetConsoleScreenshotOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConsoleScreenshot, varargs...)
	ret0, _ := ret[0].(*ec2.GetConsoleScreenshotOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetConsoleScreenshot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConsoleScreenshot, reflect.TypeOf((*MockEc2Client)(nil).GetConsoleScreenshot), varargs...)
}

func (m *MockEc2Client) GetDefaultCreditSpecification(arg0 context.Context, arg1 *ec2.GetDefaultCreditSpecificationInput, arg2 ...func(*ec2.Options)) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDefaultCreditSpecification, varargs...)
	ret0, _ := ret[0].(*ec2.GetDefaultCreditSpecificationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetDefaultCreditSpecification(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDefaultCreditSpecification, reflect.TypeOf((*MockEc2Client)(nil).GetDefaultCreditSpecification), varargs...)
}

func (m *MockEc2Client) GetEbsDefaultKmsKeyId(arg0 context.Context, arg1 *ec2.GetEbsDefaultKmsKeyIdInput, arg2 ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEbsDefaultKmsKeyId, varargs...)
	ret0, _ := ret[0].(*ec2.GetEbsDefaultKmsKeyIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetEbsDefaultKmsKeyId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEbsDefaultKmsKeyId, reflect.TypeOf((*MockEc2Client)(nil).GetEbsDefaultKmsKeyId), varargs...)
}

func (m *MockEc2Client) GetEbsEncryptionByDefault(arg0 context.Context, arg1 *ec2.GetEbsEncryptionByDefaultInput, arg2 ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEbsEncryptionByDefault, varargs...)
	ret0, _ := ret[0].(*ec2.GetEbsEncryptionByDefaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetEbsEncryptionByDefault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEbsEncryptionByDefault, reflect.TypeOf((*MockEc2Client)(nil).GetEbsEncryptionByDefault), varargs...)
}

func (m *MockEc2Client) GetFlowLogsIntegrationTemplate(arg0 context.Context, arg1 *ec2.GetFlowLogsIntegrationTemplateInput, arg2 ...func(*ec2.Options)) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFlowLogsIntegrationTemplate, varargs...)
	ret0, _ := ret[0].(*ec2.GetFlowLogsIntegrationTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetFlowLogsIntegrationTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFlowLogsIntegrationTemplate, reflect.TypeOf((*MockEc2Client)(nil).GetFlowLogsIntegrationTemplate), varargs...)
}

func (m *MockEc2Client) GetGroupsForCapacityReservation(arg0 context.Context, arg1 *ec2.GetGroupsForCapacityReservationInput, arg2 ...func(*ec2.Options)) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroupsForCapacityReservation, varargs...)
	ret0, _ := ret[0].(*ec2.GetGroupsForCapacityReservationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetGroupsForCapacityReservation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroupsForCapacityReservation, reflect.TypeOf((*MockEc2Client)(nil).GetGroupsForCapacityReservation), varargs...)
}

func (m *MockEc2Client) GetHostReservationPurchasePreview(arg0 context.Context, arg1 *ec2.GetHostReservationPurchasePreviewInput, arg2 ...func(*ec2.Options)) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetHostReservationPurchasePreview, varargs...)
	ret0, _ := ret[0].(*ec2.GetHostReservationPurchasePreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetHostReservationPurchasePreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetHostReservationPurchasePreview, reflect.TypeOf((*MockEc2Client)(nil).GetHostReservationPurchasePreview), varargs...)
}

func (m *MockEc2Client) GetInstanceTypesFromInstanceRequirements(arg0 context.Context, arg1 *ec2.GetInstanceTypesFromInstanceRequirementsInput, arg2 ...func(*ec2.Options)) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceTypesFromInstanceRequirements, varargs...)
	ret0, _ := ret[0].(*ec2.GetInstanceTypesFromInstanceRequirementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetInstanceTypesFromInstanceRequirements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceTypesFromInstanceRequirements, reflect.TypeOf((*MockEc2Client)(nil).GetInstanceTypesFromInstanceRequirements), varargs...)
}

func (m *MockEc2Client) GetInstanceUefiData(arg0 context.Context, arg1 *ec2.GetInstanceUefiDataInput, arg2 ...func(*ec2.Options)) (*ec2.GetInstanceUefiDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceUefiData, varargs...)
	ret0, _ := ret[0].(*ec2.GetInstanceUefiDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetInstanceUefiData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceUefiData, reflect.TypeOf((*MockEc2Client)(nil).GetInstanceUefiData), varargs...)
}

func (m *MockEc2Client) GetIpamAddressHistory(arg0 context.Context, arg1 *ec2.GetIpamAddressHistoryInput, arg2 ...func(*ec2.Options)) (*ec2.GetIpamAddressHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIpamAddressHistory, varargs...)
	ret0, _ := ret[0].(*ec2.GetIpamAddressHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetIpamAddressHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIpamAddressHistory, reflect.TypeOf((*MockEc2Client)(nil).GetIpamAddressHistory), varargs...)
}

func (m *MockEc2Client) GetIpamPoolAllocations(arg0 context.Context, arg1 *ec2.GetIpamPoolAllocationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetIpamPoolAllocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIpamPoolAllocations, varargs...)
	ret0, _ := ret[0].(*ec2.GetIpamPoolAllocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetIpamPoolAllocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIpamPoolAllocations, reflect.TypeOf((*MockEc2Client)(nil).GetIpamPoolAllocations), varargs...)
}

func (m *MockEc2Client) GetIpamPoolCidrs(arg0 context.Context, arg1 *ec2.GetIpamPoolCidrsInput, arg2 ...func(*ec2.Options)) (*ec2.GetIpamPoolCidrsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIpamPoolCidrs, varargs...)
	ret0, _ := ret[0].(*ec2.GetIpamPoolCidrsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetIpamPoolCidrs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIpamPoolCidrs, reflect.TypeOf((*MockEc2Client)(nil).GetIpamPoolCidrs), varargs...)
}

func (m *MockEc2Client) GetIpamResourceCidrs(arg0 context.Context, arg1 *ec2.GetIpamResourceCidrsInput, arg2 ...func(*ec2.Options)) (*ec2.GetIpamResourceCidrsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIpamResourceCidrs, varargs...)
	ret0, _ := ret[0].(*ec2.GetIpamResourceCidrsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetIpamResourceCidrs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIpamResourceCidrs, reflect.TypeOf((*MockEc2Client)(nil).GetIpamResourceCidrs), varargs...)
}

func (m *MockEc2Client) GetLaunchTemplateData(arg0 context.Context, arg1 *ec2.GetLaunchTemplateDataInput, arg2 ...func(*ec2.Options)) (*ec2.GetLaunchTemplateDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLaunchTemplateData, varargs...)
	ret0, _ := ret[0].(*ec2.GetLaunchTemplateDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetLaunchTemplateData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLaunchTemplateData, reflect.TypeOf((*MockEc2Client)(nil).GetLaunchTemplateData), varargs...)
}

func (m *MockEc2Client) GetManagedPrefixListAssociations(arg0 context.Context, arg1 *ec2.GetManagedPrefixListAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetManagedPrefixListAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.GetManagedPrefixListAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetManagedPrefixListAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetManagedPrefixListAssociations, reflect.TypeOf((*MockEc2Client)(nil).GetManagedPrefixListAssociations), varargs...)
}

func (m *MockEc2Client) GetManagedPrefixListEntries(arg0 context.Context, arg1 *ec2.GetManagedPrefixListEntriesInput, arg2 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetManagedPrefixListEntries, varargs...)
	ret0, _ := ret[0].(*ec2.GetManagedPrefixListEntriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetManagedPrefixListEntries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetManagedPrefixListEntries, reflect.TypeOf((*MockEc2Client)(nil).GetManagedPrefixListEntries), varargs...)
}

func (m *MockEc2Client) GetNetworkInsightsAccessScopeAnalysisFindings(arg0 context.Context, arg1 *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, arg2 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetNetworkInsightsAccessScopeAnalysisFindings, varargs...)
	ret0, _ := ret[0].(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetNetworkInsightsAccessScopeAnalysisFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetNetworkInsightsAccessScopeAnalysisFindings, reflect.TypeOf((*MockEc2Client)(nil).GetNetworkInsightsAccessScopeAnalysisFindings), varargs...)
}

func (m *MockEc2Client) GetNetworkInsightsAccessScopeContent(arg0 context.Context, arg1 *ec2.GetNetworkInsightsAccessScopeContentInput, arg2 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetNetworkInsightsAccessScopeContent, varargs...)
	ret0, _ := ret[0].(*ec2.GetNetworkInsightsAccessScopeContentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetNetworkInsightsAccessScopeContent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetNetworkInsightsAccessScopeContent, reflect.TypeOf((*MockEc2Client)(nil).GetNetworkInsightsAccessScopeContent), varargs...)
}

func (m *MockEc2Client) GetPasswordData(arg0 context.Context, arg1 *ec2.GetPasswordDataInput, arg2 ...func(*ec2.Options)) (*ec2.GetPasswordDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPasswordData, varargs...)
	ret0, _ := ret[0].(*ec2.GetPasswordDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetPasswordData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPasswordData, reflect.TypeOf((*MockEc2Client)(nil).GetPasswordData), varargs...)
}

func (m *MockEc2Client) GetReservedInstancesExchangeQuote(arg0 context.Context, arg1 *ec2.GetReservedInstancesExchangeQuoteInput, arg2 ...func(*ec2.Options)) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetReservedInstancesExchangeQuote, varargs...)
	ret0, _ := ret[0].(*ec2.GetReservedInstancesExchangeQuoteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetReservedInstancesExchangeQuote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetReservedInstancesExchangeQuote, reflect.TypeOf((*MockEc2Client)(nil).GetReservedInstancesExchangeQuote), varargs...)
}

func (m *MockEc2Client) GetSerialConsoleAccessStatus(arg0 context.Context, arg1 *ec2.GetSerialConsoleAccessStatusInput, arg2 ...func(*ec2.Options)) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSerialConsoleAccessStatus, varargs...)
	ret0, _ := ret[0].(*ec2.GetSerialConsoleAccessStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetSerialConsoleAccessStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSerialConsoleAccessStatus, reflect.TypeOf((*MockEc2Client)(nil).GetSerialConsoleAccessStatus), varargs...)
}

func (m *MockEc2Client) GetSpotPlacementScores(arg0 context.Context, arg1 *ec2.GetSpotPlacementScoresInput, arg2 ...func(*ec2.Options)) (*ec2.GetSpotPlacementScoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSpotPlacementScores, varargs...)
	ret0, _ := ret[0].(*ec2.GetSpotPlacementScoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetSpotPlacementScores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSpotPlacementScores, reflect.TypeOf((*MockEc2Client)(nil).GetSpotPlacementScores), varargs...)
}

func (m *MockEc2Client) GetSubnetCidrReservations(arg0 context.Context, arg1 *ec2.GetSubnetCidrReservationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetSubnetCidrReservationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSubnetCidrReservations, varargs...)
	ret0, _ := ret[0].(*ec2.GetSubnetCidrReservationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetSubnetCidrReservations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSubnetCidrReservations, reflect.TypeOf((*MockEc2Client)(nil).GetSubnetCidrReservations), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayAttachmentPropagations(arg0 context.Context, arg1 *ec2.GetTransitGatewayAttachmentPropagationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayAttachmentPropagations, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayAttachmentPropagationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayAttachmentPropagations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayAttachmentPropagations, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayAttachmentPropagations), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayMulticastDomainAssociations(arg0 context.Context, arg1 *ec2.GetTransitGatewayMulticastDomainAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayMulticastDomainAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayMulticastDomainAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayMulticastDomainAssociations, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayMulticastDomainAssociations), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayPolicyTableAssociations(arg0 context.Context, arg1 *ec2.GetTransitGatewayPolicyTableAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayPolicyTableAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayPolicyTableAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayPolicyTableAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayPolicyTableAssociations, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayPolicyTableAssociations), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayPolicyTableEntries(arg0 context.Context, arg1 *ec2.GetTransitGatewayPolicyTableEntriesInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayPolicyTableEntries, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayPolicyTableEntriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayPolicyTableEntries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayPolicyTableEntries, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayPolicyTableEntries), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayPrefixListReferences(arg0 context.Context, arg1 *ec2.GetTransitGatewayPrefixListReferencesInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayPrefixListReferences, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayPrefixListReferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayPrefixListReferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayPrefixListReferences, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayPrefixListReferences), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayRouteTableAssociations(arg0 context.Context, arg1 *ec2.GetTransitGatewayRouteTableAssociationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayRouteTableAssociations, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayRouteTableAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayRouteTableAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayRouteTableAssociations, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayRouteTableAssociations), varargs...)
}

func (m *MockEc2Client) GetTransitGatewayRouteTablePropagations(arg0 context.Context, arg1 *ec2.GetTransitGatewayRouteTablePropagationsInput, arg2 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTransitGatewayRouteTablePropagations, varargs...)
	ret0, _ := ret[0].(*ec2.GetTransitGatewayRouteTablePropagationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetTransitGatewayRouteTablePropagations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTransitGatewayRouteTablePropagations, reflect.TypeOf((*MockEc2Client)(nil).GetTransitGatewayRouteTablePropagations), varargs...)
}

func (m *MockEc2Client) GetVerifiedAccessEndpointPolicy(arg0 context.Context, arg1 *ec2.GetVerifiedAccessEndpointPolicyInput, arg2 ...func(*ec2.Options)) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVerifiedAccessEndpointPolicy, varargs...)
	ret0, _ := ret[0].(*ec2.GetVerifiedAccessEndpointPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetVerifiedAccessEndpointPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVerifiedAccessEndpointPolicy, reflect.TypeOf((*MockEc2Client)(nil).GetVerifiedAccessEndpointPolicy), varargs...)
}

func (m *MockEc2Client) GetVerifiedAccessGroupPolicy(arg0 context.Context, arg1 *ec2.GetVerifiedAccessGroupPolicyInput, arg2 ...func(*ec2.Options)) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVerifiedAccessGroupPolicy, varargs...)
	ret0, _ := ret[0].(*ec2.GetVerifiedAccessGroupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetVerifiedAccessGroupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVerifiedAccessGroupPolicy, reflect.TypeOf((*MockEc2Client)(nil).GetVerifiedAccessGroupPolicy), varargs...)
}

func (m *MockEc2Client) GetVpnConnectionDeviceSampleConfiguration(arg0 context.Context, arg1 *ec2.GetVpnConnectionDeviceSampleConfigurationInput, arg2 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVpnConnectionDeviceSampleConfiguration, varargs...)
	ret0, _ := ret[0].(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetVpnConnectionDeviceSampleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVpnConnectionDeviceSampleConfiguration, reflect.TypeOf((*MockEc2Client)(nil).GetVpnConnectionDeviceSampleConfiguration), varargs...)
}

func (m *MockEc2Client) GetVpnConnectionDeviceTypes(arg0 context.Context, arg1 *ec2.GetVpnConnectionDeviceTypesInput, arg2 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVpnConnectionDeviceTypes, varargs...)
	ret0, _ := ret[0].(*ec2.GetVpnConnectionDeviceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetVpnConnectionDeviceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVpnConnectionDeviceTypes, reflect.TypeOf((*MockEc2Client)(nil).GetVpnConnectionDeviceTypes), varargs...)
}

func (m *MockEc2Client) ListImagesInRecycleBin(arg0 context.Context, arg1 *ec2.ListImagesInRecycleBinInput, arg2 ...func(*ec2.Options)) (*ec2.ListImagesInRecycleBinOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImagesInRecycleBin, varargs...)
	ret0, _ := ret[0].(*ec2.ListImagesInRecycleBinOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) ListImagesInRecycleBin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImagesInRecycleBin, reflect.TypeOf((*MockEc2Client)(nil).ListImagesInRecycleBin), varargs...)
}

func (m *MockEc2Client) ListSnapshotsInRecycleBin(arg0 context.Context, arg1 *ec2.ListSnapshotsInRecycleBinInput, arg2 ...func(*ec2.Options)) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSnapshotsInRecycleBin, varargs...)
	ret0, _ := ret[0].(*ec2.ListSnapshotsInRecycleBinOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) ListSnapshotsInRecycleBin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSnapshotsInRecycleBin, reflect.TypeOf((*MockEc2Client)(nil).ListSnapshotsInRecycleBin), varargs...)
}

func (m *MockEc2Client) SearchLocalGatewayRoutes(arg0 context.Context, arg1 *ec2.SearchLocalGatewayRoutesInput, arg2 ...func(*ec2.Options)) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchLocalGatewayRoutes, varargs...)
	ret0, _ := ret[0].(*ec2.SearchLocalGatewayRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) SearchLocalGatewayRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchLocalGatewayRoutes, reflect.TypeOf((*MockEc2Client)(nil).SearchLocalGatewayRoutes), varargs...)
}

func (m *MockEc2Client) SearchTransitGatewayMulticastGroups(arg0 context.Context, arg1 *ec2.SearchTransitGatewayMulticastGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchTransitGatewayMulticastGroups, varargs...)
	ret0, _ := ret[0].(*ec2.SearchTransitGatewayMulticastGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) SearchTransitGatewayMulticastGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchTransitGatewayMulticastGroups, reflect.TypeOf((*MockEc2Client)(nil).SearchTransitGatewayMulticastGroups), varargs...)
}

func (m *MockEc2Client) SearchTransitGatewayRoutes(arg0 context.Context, arg1 *ec2.SearchTransitGatewayRoutesInput, arg2 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchTransitGatewayRoutes, varargs...)
	ret0, _ := ret[0].(*ec2.SearchTransitGatewayRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) SearchTransitGatewayRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchTransitGatewayRoutes, reflect.TypeOf((*MockEc2Client)(nil).SearchTransitGatewayRoutes), varargs...)
}
