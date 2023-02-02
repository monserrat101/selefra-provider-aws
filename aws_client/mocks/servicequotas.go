package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	servicequotas "github.com/aws/aws-sdk-go-v2/service/servicequotas"
	gomock "github.com/golang/mock/gomock"
)

type MockServicequotasClient struct {
	ctrl		*gomock.Controller
	recorder	*MockServicequotasClientMockRecorder
}

type MockServicequotasClientMockRecorder struct {
	mock *MockServicequotasClient
}

func NewMockServicequotasClient(ctrl *gomock.Controller) *MockServicequotasClient {
	mock := &MockServicequotasClient{ctrl: ctrl}
	mock.recorder = &MockServicequotasClientMockRecorder{mock}
	return mock
}

func (m *MockServicequotasClient) EXPECT() *MockServicequotasClientMockRecorder {
	return m.recorder
}

func (m *MockServicequotasClient) GetAWSDefaultServiceQuota(arg0 context.Context, arg1 *servicequotas.GetAWSDefaultServiceQuotaInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAWSDefaultServiceQuota, varargs...)
	ret0, _ := ret[0].(*servicequotas.GetAWSDefaultServiceQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) GetAWSDefaultServiceQuota(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAWSDefaultServiceQuota, reflect.TypeOf((*MockServicequotasClient)(nil).GetAWSDefaultServiceQuota), varargs...)
}

func (m *MockServicequotasClient) GetAssociationForServiceQuotaTemplate(arg0 context.Context, arg1 *servicequotas.GetAssociationForServiceQuotaTemplateInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAssociationForServiceQuotaTemplate, varargs...)
	ret0, _ := ret[0].(*servicequotas.GetAssociationForServiceQuotaTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) GetAssociationForServiceQuotaTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAssociationForServiceQuotaTemplate, reflect.TypeOf((*MockServicequotasClient)(nil).GetAssociationForServiceQuotaTemplate), varargs...)
}

func (m *MockServicequotasClient) GetRequestedServiceQuotaChange(arg0 context.Context, arg1 *servicequotas.GetRequestedServiceQuotaChangeInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRequestedServiceQuotaChange, varargs...)
	ret0, _ := ret[0].(*servicequotas.GetRequestedServiceQuotaChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) GetRequestedServiceQuotaChange(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRequestedServiceQuotaChange, reflect.TypeOf((*MockServicequotasClient)(nil).GetRequestedServiceQuotaChange), varargs...)
}

func (m *MockServicequotasClient) GetServiceQuota(arg0 context.Context, arg1 *servicequotas.GetServiceQuotaInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.GetServiceQuotaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceQuota, varargs...)
	ret0, _ := ret[0].(*servicequotas.GetServiceQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) GetServiceQuota(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceQuota, reflect.TypeOf((*MockServicequotasClient)(nil).GetServiceQuota), varargs...)
}

func (m *MockServicequotasClient) GetServiceQuotaIncreaseRequestFromTemplate(arg0 context.Context, arg1 *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceQuotaIncreaseRequestFromTemplate, varargs...)
	ret0, _ := ret[0].(*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) GetServiceQuotaIncreaseRequestFromTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceQuotaIncreaseRequestFromTemplate, reflect.TypeOf((*MockServicequotasClient)(nil).GetServiceQuotaIncreaseRequestFromTemplate), varargs...)
}

func (m *MockServicequotasClient) ListAWSDefaultServiceQuotas(arg0 context.Context, arg1 *servicequotas.ListAWSDefaultServiceQuotasInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAWSDefaultServiceQuotas, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListAWSDefaultServiceQuotasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListAWSDefaultServiceQuotas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAWSDefaultServiceQuotas, reflect.TypeOf((*MockServicequotasClient)(nil).ListAWSDefaultServiceQuotas), varargs...)
}

func (m *MockServicequotasClient) ListRequestedServiceQuotaChangeHistory(arg0 context.Context, arg1 *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRequestedServiceQuotaChangeHistory, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListRequestedServiceQuotaChangeHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRequestedServiceQuotaChangeHistory, reflect.TypeOf((*MockServicequotasClient)(nil).ListRequestedServiceQuotaChangeHistory), varargs...)
}

func (m *MockServicequotasClient) ListRequestedServiceQuotaChangeHistoryByQuota(arg0 context.Context, arg1 *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRequestedServiceQuotaChangeHistoryByQuota, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListRequestedServiceQuotaChangeHistoryByQuota(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRequestedServiceQuotaChangeHistoryByQuota, reflect.TypeOf((*MockServicequotasClient)(nil).ListRequestedServiceQuotaChangeHistoryByQuota), varargs...)
}

func (m *MockServicequotasClient) ListServiceQuotaIncreaseRequestsInTemplate(arg0 context.Context, arg1 *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServiceQuotaIncreaseRequestsInTemplate, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListServiceQuotaIncreaseRequestsInTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServiceQuotaIncreaseRequestsInTemplate, reflect.TypeOf((*MockServicequotasClient)(nil).ListServiceQuotaIncreaseRequestsInTemplate), varargs...)
}

func (m *MockServicequotasClient) ListServiceQuotas(arg0 context.Context, arg1 *servicequotas.ListServiceQuotasInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServiceQuotas, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListServiceQuotasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListServiceQuotas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServiceQuotas, reflect.TypeOf((*MockServicequotasClient)(nil).ListServiceQuotas), varargs...)
}

func (m *MockServicequotasClient) ListServices(arg0 context.Context, arg1 *servicequotas.ListServicesInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServices, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServices, reflect.TypeOf((*MockServicequotasClient)(nil).ListServices), varargs...)
}

func (m *MockServicequotasClient) ListTagsForResource(arg0 context.Context, arg1 *servicequotas.ListTagsForResourceInput, arg2 ...func(*servicequotas.Options)) (*servicequotas.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*servicequotas.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicequotasClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockServicequotasClient)(nil).ListTagsForResource), varargs...)
}
