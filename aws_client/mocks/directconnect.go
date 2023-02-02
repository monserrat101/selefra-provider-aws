package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	directconnect "github.com/aws/aws-sdk-go-v2/service/directconnect"
	gomock "github.com/golang/mock/gomock"
)

type MockDirectconnectClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDirectconnectClientMockRecorder
}

type MockDirectconnectClientMockRecorder struct {
	mock *MockDirectconnectClient
}

func NewMockDirectconnectClient(ctrl *gomock.Controller) *MockDirectconnectClient {
	mock := &MockDirectconnectClient{ctrl: ctrl}
	mock.recorder = &MockDirectconnectClientMockRecorder{mock}
	return mock
}

func (m *MockDirectconnectClient) EXPECT() *MockDirectconnectClientMockRecorder {
	return m.recorder
}

func (m *MockDirectconnectClient) DescribeConnectionLoa(arg0 context.Context, arg1 *directconnect.DescribeConnectionLoaInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionLoaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnectionLoa, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionLoaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeConnectionLoa(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnectionLoa, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnectionLoa), varargs...)
}

func (m *MockDirectconnectClient) DescribeConnections(arg0 context.Context, arg1 *directconnect.DescribeConnectionsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnections, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnections, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnections), varargs...)
}

func (m *MockDirectconnectClient) DescribeConnectionsOnInterconnect(arg0 context.Context, arg1 *directconnect.DescribeConnectionsOnInterconnectInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOnInterconnectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnectionsOnInterconnect, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionsOnInterconnectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeConnectionsOnInterconnect(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnectionsOnInterconnect, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnectionsOnInterconnect), varargs...)
}

func (m *MockDirectconnectClient) DescribeCustomerMetadata(arg0 context.Context, arg1 *directconnect.DescribeCustomerMetadataInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeCustomerMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCustomerMetadata, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeCustomerMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeCustomerMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCustomerMetadata, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeCustomerMetadata), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAssociationProposals(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDirectConnectGatewayAssociationProposals, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAssociationProposals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDirectConnectGatewayAssociationProposals, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAssociationProposals), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAssociations(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAssociationsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDirectConnectGatewayAssociations, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDirectConnectGatewayAssociations, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAssociations), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAttachments(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAttachmentsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDirectConnectGatewayAttachments, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDirectConnectGatewayAttachments, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAttachments), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGateways(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDirectConnectGateways, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDirectConnectGateways, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGateways), varargs...)
}

func (m *MockDirectconnectClient) DescribeHostedConnections(arg0 context.Context, arg1 *directconnect.DescribeHostedConnectionsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeHostedConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHostedConnections, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeHostedConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeHostedConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHostedConnections, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeHostedConnections), varargs...)
}

func (m *MockDirectconnectClient) DescribeInterconnectLoa(arg0 context.Context, arg1 *directconnect.DescribeInterconnectLoaInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectLoaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInterconnectLoa, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeInterconnectLoaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeInterconnectLoa(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInterconnectLoa, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeInterconnectLoa), varargs...)
}

func (m *MockDirectconnectClient) DescribeInterconnects(arg0 context.Context, arg1 *directconnect.DescribeInterconnectsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInterconnects, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeInterconnectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeInterconnects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInterconnects, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeInterconnects), varargs...)
}

func (m *MockDirectconnectClient) DescribeLags(arg0 context.Context, arg1 *directconnect.DescribeLagsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLags, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeLags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLags, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLags), varargs...)
}

func (m *MockDirectconnectClient) DescribeLoa(arg0 context.Context, arg1 *directconnect.DescribeLoaInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLoaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLoa, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLoaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeLoa(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLoa, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLoa), varargs...)
}

func (m *MockDirectconnectClient) DescribeLocations(arg0 context.Context, arg1 *directconnect.DescribeLocationsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLocations, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeLocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLocations, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLocations), varargs...)
}

func (m *MockDirectconnectClient) DescribeRouterConfiguration(arg0 context.Context, arg1 *directconnect.DescribeRouterConfigurationInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeRouterConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRouterConfiguration, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeRouterConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeRouterConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRouterConfiguration, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeRouterConfiguration), varargs...)
}

func (m *MockDirectconnectClient) DescribeTags(arg0 context.Context, arg1 *directconnect.DescribeTagsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTags, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTags, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeTags), varargs...)
}

func (m *MockDirectconnectClient) DescribeVirtualGateways(arg0 context.Context, arg1 *directconnect.DescribeVirtualGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVirtualGateways, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVirtualGateways, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualGateways), varargs...)
}

func (m *MockDirectconnectClient) DescribeVirtualInterfaces(arg0 context.Context, arg1 *directconnect.DescribeVirtualInterfacesInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeVirtualInterfaces, varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeVirtualInterfaces, reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualInterfaces), varargs...)
}

func (m *MockDirectconnectClient) ListVirtualInterfaceTestHistory(arg0 context.Context, arg1 *directconnect.ListVirtualInterfaceTestHistoryInput, arg2 ...func(*directconnect.Options)) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVirtualInterfaceTestHistory, varargs...)
	ret0, _ := ret[0].(*directconnect.ListVirtualInterfaceTestHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) ListVirtualInterfaceTestHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVirtualInterfaceTestHistory, reflect.TypeOf((*MockDirectconnectClient)(nil).ListVirtualInterfaceTestHistory), varargs...)
}
