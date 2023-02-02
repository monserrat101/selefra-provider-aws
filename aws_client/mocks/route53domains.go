package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	route53domains "github.com/aws/aws-sdk-go-v2/service/route53domains"
	gomock "github.com/golang/mock/gomock"
)

type MockRoute53domainsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRoute53domainsClientMockRecorder
}

type MockRoute53domainsClientMockRecorder struct {
	mock *MockRoute53domainsClient
}

func NewMockRoute53domainsClient(ctrl *gomock.Controller) *MockRoute53domainsClient {
	mock := &MockRoute53domainsClient{ctrl: ctrl}
	mock.recorder = &MockRoute53domainsClientMockRecorder{mock}
	return mock
}

func (m *MockRoute53domainsClient) EXPECT() *MockRoute53domainsClientMockRecorder {
	return m.recorder
}

func (m *MockRoute53domainsClient) GetContactReachabilityStatus(arg0 context.Context, arg1 *route53domains.GetContactReachabilityStatusInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetContactReachabilityStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContactReachabilityStatus, varargs...)
	ret0, _ := ret[0].(*route53domains.GetContactReachabilityStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) GetContactReachabilityStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContactReachabilityStatus, reflect.TypeOf((*MockRoute53domainsClient)(nil).GetContactReachabilityStatus), varargs...)
}

func (m *MockRoute53domainsClient) GetDomainDetail(arg0 context.Context, arg1 *route53domains.GetDomainDetailInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainDetail, varargs...)
	ret0, _ := ret[0].(*route53domains.GetDomainDetailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) GetDomainDetail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainDetail, reflect.TypeOf((*MockRoute53domainsClient)(nil).GetDomainDetail), varargs...)
}

func (m *MockRoute53domainsClient) GetDomainSuggestions(arg0 context.Context, arg1 *route53domains.GetDomainSuggestionsInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetDomainSuggestionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainSuggestions, varargs...)
	ret0, _ := ret[0].(*route53domains.GetDomainSuggestionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) GetDomainSuggestions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainSuggestions, reflect.TypeOf((*MockRoute53domainsClient)(nil).GetDomainSuggestions), varargs...)
}

func (m *MockRoute53domainsClient) GetOperationDetail(arg0 context.Context, arg1 *route53domains.GetOperationDetailInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetOperationDetailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOperationDetail, varargs...)
	ret0, _ := ret[0].(*route53domains.GetOperationDetailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) GetOperationDetail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOperationDetail, reflect.TypeOf((*MockRoute53domainsClient)(nil).GetOperationDetail), varargs...)
}

func (m *MockRoute53domainsClient) ListDomains(arg0 context.Context, arg1 *route53domains.ListDomainsInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomains, varargs...)
	ret0, _ := ret[0].(*route53domains.ListDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) ListDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomains, reflect.TypeOf((*MockRoute53domainsClient)(nil).ListDomains), varargs...)
}

func (m *MockRoute53domainsClient) ListOperations(arg0 context.Context, arg1 *route53domains.ListOperationsInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOperations, varargs...)
	ret0, _ := ret[0].(*route53domains.ListOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOperations, reflect.TypeOf((*MockRoute53domainsClient)(nil).ListOperations), varargs...)
}

func (m *MockRoute53domainsClient) ListPrices(arg0 context.Context, arg1 *route53domains.ListPricesInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListPricesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPrices, varargs...)
	ret0, _ := ret[0].(*route53domains.ListPricesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) ListPrices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPrices, reflect.TypeOf((*MockRoute53domainsClient)(nil).ListPrices), varargs...)
}

func (m *MockRoute53domainsClient) ListTagsForDomain(arg0 context.Context, arg1 *route53domains.ListTagsForDomainInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForDomain, varargs...)
	ret0, _ := ret[0].(*route53domains.ListTagsForDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53domainsClientMockRecorder) ListTagsForDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForDomain, reflect.TypeOf((*MockRoute53domainsClient)(nil).ListTagsForDomain), varargs...)
}
