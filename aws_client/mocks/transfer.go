package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	transfer "github.com/aws/aws-sdk-go-v2/service/transfer"
	gomock "github.com/golang/mock/gomock"
)

type MockTransferClient struct {
	ctrl		*gomock.Controller
	recorder	*MockTransferClientMockRecorder
}

type MockTransferClientMockRecorder struct {
	mock *MockTransferClient
}

func NewMockTransferClient(ctrl *gomock.Controller) *MockTransferClient {
	mock := &MockTransferClient{ctrl: ctrl}
	mock.recorder = &MockTransferClientMockRecorder{mock}
	return mock
}

func (m *MockTransferClient) EXPECT() *MockTransferClientMockRecorder {
	return m.recorder
}

func (m *MockTransferClient) DescribeAccess(arg0 context.Context, arg1 *transfer.DescribeAccessInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccess, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccess, reflect.TypeOf((*MockTransferClient)(nil).DescribeAccess), varargs...)
}

func (m *MockTransferClient) DescribeAgreement(arg0 context.Context, arg1 *transfer.DescribeAgreementInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeAgreementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAgreement, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeAgreementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeAgreement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAgreement, reflect.TypeOf((*MockTransferClient)(nil).DescribeAgreement), varargs...)
}

func (m *MockTransferClient) DescribeCertificate(arg0 context.Context, arg1 *transfer.DescribeCertificateInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCertificate, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCertificate, reflect.TypeOf((*MockTransferClient)(nil).DescribeCertificate), varargs...)
}

func (m *MockTransferClient) DescribeConnector(arg0 context.Context, arg1 *transfer.DescribeConnectorInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeConnectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConnector, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeConnectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConnector, reflect.TypeOf((*MockTransferClient)(nil).DescribeConnector), varargs...)
}

func (m *MockTransferClient) DescribeExecution(arg0 context.Context, arg1 *transfer.DescribeExecutionInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExecution, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExecution, reflect.TypeOf((*MockTransferClient)(nil).DescribeExecution), varargs...)
}

func (m *MockTransferClient) DescribeHostKey(arg0 context.Context, arg1 *transfer.DescribeHostKeyInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeHostKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHostKey, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeHostKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeHostKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHostKey, reflect.TypeOf((*MockTransferClient)(nil).DescribeHostKey), varargs...)
}

func (m *MockTransferClient) DescribeProfile(arg0 context.Context, arg1 *transfer.DescribeProfileInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProfile, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProfile, reflect.TypeOf((*MockTransferClient)(nil).DescribeProfile), varargs...)
}

func (m *MockTransferClient) DescribeSecurityPolicy(arg0 context.Context, arg1 *transfer.DescribeSecurityPolicyInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeSecurityPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecurityPolicy, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeSecurityPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeSecurityPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecurityPolicy, reflect.TypeOf((*MockTransferClient)(nil).DescribeSecurityPolicy), varargs...)
}

func (m *MockTransferClient) DescribeServer(arg0 context.Context, arg1 *transfer.DescribeServerInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeServer, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeServer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeServer, reflect.TypeOf((*MockTransferClient)(nil).DescribeServer), varargs...)
}

func (m *MockTransferClient) DescribeUser(arg0 context.Context, arg1 *transfer.DescribeUserInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUser, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUser, reflect.TypeOf((*MockTransferClient)(nil).DescribeUser), varargs...)
}

func (m *MockTransferClient) DescribeWorkflow(arg0 context.Context, arg1 *transfer.DescribeWorkflowInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeWorkflowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkflow, varargs...)
	ret0, _ := ret[0].(*transfer.DescribeWorkflowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkflow, reflect.TypeOf((*MockTransferClient)(nil).DescribeWorkflow), varargs...)
}

func (m *MockTransferClient) ListAccesses(arg0 context.Context, arg1 *transfer.ListAccessesInput, arg2 ...func(*transfer.Options)) (*transfer.ListAccessesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccesses, varargs...)
	ret0, _ := ret[0].(*transfer.ListAccessesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListAccesses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccesses, reflect.TypeOf((*MockTransferClient)(nil).ListAccesses), varargs...)
}

func (m *MockTransferClient) ListAgreements(arg0 context.Context, arg1 *transfer.ListAgreementsInput, arg2 ...func(*transfer.Options)) (*transfer.ListAgreementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAgreements, varargs...)
	ret0, _ := ret[0].(*transfer.ListAgreementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListAgreements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAgreements, reflect.TypeOf((*MockTransferClient)(nil).ListAgreements), varargs...)
}

func (m *MockTransferClient) ListCertificates(arg0 context.Context, arg1 *transfer.ListCertificatesInput, arg2 ...func(*transfer.Options)) (*transfer.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCertificates, varargs...)
	ret0, _ := ret[0].(*transfer.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCertificates, reflect.TypeOf((*MockTransferClient)(nil).ListCertificates), varargs...)
}

func (m *MockTransferClient) ListConnectors(arg0 context.Context, arg1 *transfer.ListConnectorsInput, arg2 ...func(*transfer.Options)) (*transfer.ListConnectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConnectors, varargs...)
	ret0, _ := ret[0].(*transfer.ListConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConnectors, reflect.TypeOf((*MockTransferClient)(nil).ListConnectors), varargs...)
}

func (m *MockTransferClient) ListExecutions(arg0 context.Context, arg1 *transfer.ListExecutionsInput, arg2 ...func(*transfer.Options)) (*transfer.ListExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExecutions, varargs...)
	ret0, _ := ret[0].(*transfer.ListExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExecutions, reflect.TypeOf((*MockTransferClient)(nil).ListExecutions), varargs...)
}

func (m *MockTransferClient) ListHostKeys(arg0 context.Context, arg1 *transfer.ListHostKeysInput, arg2 ...func(*transfer.Options)) (*transfer.ListHostKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHostKeys, varargs...)
	ret0, _ := ret[0].(*transfer.ListHostKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListHostKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHostKeys, reflect.TypeOf((*MockTransferClient)(nil).ListHostKeys), varargs...)
}

func (m *MockTransferClient) ListProfiles(arg0 context.Context, arg1 *transfer.ListProfilesInput, arg2 ...func(*transfer.Options)) (*transfer.ListProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProfiles, varargs...)
	ret0, _ := ret[0].(*transfer.ListProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProfiles, reflect.TypeOf((*MockTransferClient)(nil).ListProfiles), varargs...)
}

func (m *MockTransferClient) ListSecurityPolicies(arg0 context.Context, arg1 *transfer.ListSecurityPoliciesInput, arg2 ...func(*transfer.Options)) (*transfer.ListSecurityPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSecurityPolicies, varargs...)
	ret0, _ := ret[0].(*transfer.ListSecurityPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListSecurityPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSecurityPolicies, reflect.TypeOf((*MockTransferClient)(nil).ListSecurityPolicies), varargs...)
}

func (m *MockTransferClient) ListServers(arg0 context.Context, arg1 *transfer.ListServersInput, arg2 ...func(*transfer.Options)) (*transfer.ListServersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServers, varargs...)
	ret0, _ := ret[0].(*transfer.ListServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListServers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServers, reflect.TypeOf((*MockTransferClient)(nil).ListServers), varargs...)
}

func (m *MockTransferClient) ListTagsForResource(arg0 context.Context, arg1 *transfer.ListTagsForResourceInput, arg2 ...func(*transfer.Options)) (*transfer.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*transfer.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockTransferClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockTransferClient) ListUsers(arg0 context.Context, arg1 *transfer.ListUsersInput, arg2 ...func(*transfer.Options)) (*transfer.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsers, varargs...)
	ret0, _ := ret[0].(*transfer.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsers, reflect.TypeOf((*MockTransferClient)(nil).ListUsers), varargs...)
}

func (m *MockTransferClient) ListWorkflows(arg0 context.Context, arg1 *transfer.ListWorkflowsInput, arg2 ...func(*transfer.Options)) (*transfer.ListWorkflowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWorkflows, varargs...)
	ret0, _ := ret[0].(*transfer.ListWorkflowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWorkflows, reflect.TypeOf((*MockTransferClient)(nil).ListWorkflows), varargs...)
}
