package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	acm "github.com/aws/aws-sdk-go-v2/service/acm"
	gomock "github.com/golang/mock/gomock"
)

type MockAcmClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAcmClientMockRecorder
}

type MockAcmClientMockRecorder struct {
	mock *MockAcmClient
}

func NewMockAcmClient(ctrl *gomock.Controller) *MockAcmClient {
	mock := &MockAcmClient{ctrl: ctrl}
	mock.recorder = &MockAcmClientMockRecorder{mock}
	return mock
}

func (m *MockAcmClient) EXPECT() *MockAcmClientMockRecorder {
	return m.recorder
}

func (m *MockAcmClient) DescribeCertificate(arg0 context.Context, arg1 *acm.DescribeCertificateInput, arg2 ...func(*acm.Options)) (*acm.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCertificate, varargs...)
	ret0, _ := ret[0].(*acm.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAcmClientMockRecorder) DescribeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCertificate, reflect.TypeOf((*MockAcmClient)(nil).DescribeCertificate), varargs...)
}

func (m *MockAcmClient) GetAccountConfiguration(arg0 context.Context, arg1 *acm.GetAccountConfigurationInput, arg2 ...func(*acm.Options)) (*acm.GetAccountConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountConfiguration, varargs...)
	ret0, _ := ret[0].(*acm.GetAccountConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAcmClientMockRecorder) GetAccountConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountConfiguration, reflect.TypeOf((*MockAcmClient)(nil).GetAccountConfiguration), varargs...)
}

func (m *MockAcmClient) GetCertificate(arg0 context.Context, arg1 *acm.GetCertificateInput, arg2 ...func(*acm.Options)) (*acm.GetCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCertificate, varargs...)
	ret0, _ := ret[0].(*acm.GetCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAcmClientMockRecorder) GetCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCertificate, reflect.TypeOf((*MockAcmClient)(nil).GetCertificate), varargs...)
}

func (m *MockAcmClient) ListCertificates(arg0 context.Context, arg1 *acm.ListCertificatesInput, arg2 ...func(*acm.Options)) (*acm.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCertificates, varargs...)
	ret0, _ := ret[0].(*acm.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAcmClientMockRecorder) ListCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCertificates, reflect.TypeOf((*MockAcmClient)(nil).ListCertificates), varargs...)
}

func (m *MockAcmClient) ListTagsForCertificate(arg0 context.Context, arg1 *acm.ListTagsForCertificateInput, arg2 ...func(*acm.Options)) (*acm.ListTagsForCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForCertificate, varargs...)
	ret0, _ := ret[0].(*acm.ListTagsForCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAcmClientMockRecorder) ListTagsForCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForCertificate, reflect.TypeOf((*MockAcmClient)(nil).ListTagsForCertificate), varargs...)
}
