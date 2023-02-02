package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	iot "github.com/aws/aws-sdk-go-v2/service/iot"
	gomock "github.com/golang/mock/gomock"
)

type MockIotClient struct {
	ctrl		*gomock.Controller
	recorder	*MockIotClientMockRecorder
}

type MockIotClientMockRecorder struct {
	mock *MockIotClient
}

func NewMockIotClient(ctrl *gomock.Controller) *MockIotClient {
	mock := &MockIotClient{ctrl: ctrl}
	mock.recorder = &MockIotClientMockRecorder{mock}
	return mock
}

func (m *MockIotClient) EXPECT() *MockIotClientMockRecorder {
	return m.recorder
}

func (m *MockIotClient) DescribeAccountAuditConfiguration(arg0 context.Context, arg1 *iot.DescribeAccountAuditConfigurationInput, arg2 ...func(*iot.Options)) (*iot.DescribeAccountAuditConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAccountAuditConfiguration, varargs...)
	ret0, _ := ret[0].(*iot.DescribeAccountAuditConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeAccountAuditConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAccountAuditConfiguration, reflect.TypeOf((*MockIotClient)(nil).DescribeAccountAuditConfiguration), varargs...)
}

func (m *MockIotClient) DescribeAuditFinding(arg0 context.Context, arg1 *iot.DescribeAuditFindingInput, arg2 ...func(*iot.Options)) (*iot.DescribeAuditFindingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAuditFinding, varargs...)
	ret0, _ := ret[0].(*iot.DescribeAuditFindingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeAuditFinding(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAuditFinding, reflect.TypeOf((*MockIotClient)(nil).DescribeAuditFinding), varargs...)
}

func (m *MockIotClient) DescribeAuditMitigationActionsTask(arg0 context.Context, arg1 *iot.DescribeAuditMitigationActionsTaskInput, arg2 ...func(*iot.Options)) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAuditMitigationActionsTask, varargs...)
	ret0, _ := ret[0].(*iot.DescribeAuditMitigationActionsTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeAuditMitigationActionsTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAuditMitigationActionsTask, reflect.TypeOf((*MockIotClient)(nil).DescribeAuditMitigationActionsTask), varargs...)
}

func (m *MockIotClient) DescribeAuditSuppression(arg0 context.Context, arg1 *iot.DescribeAuditSuppressionInput, arg2 ...func(*iot.Options)) (*iot.DescribeAuditSuppressionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAuditSuppression, varargs...)
	ret0, _ := ret[0].(*iot.DescribeAuditSuppressionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeAuditSuppression(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAuditSuppression, reflect.TypeOf((*MockIotClient)(nil).DescribeAuditSuppression), varargs...)
}

func (m *MockIotClient) DescribeAuditTask(arg0 context.Context, arg1 *iot.DescribeAuditTaskInput, arg2 ...func(*iot.Options)) (*iot.DescribeAuditTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAuditTask, varargs...)
	ret0, _ := ret[0].(*iot.DescribeAuditTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeAuditTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAuditTask, reflect.TypeOf((*MockIotClient)(nil).DescribeAuditTask), varargs...)
}

func (m *MockIotClient) DescribeAuthorizer(arg0 context.Context, arg1 *iot.DescribeAuthorizerInput, arg2 ...func(*iot.Options)) (*iot.DescribeAuthorizerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAuthorizer, varargs...)
	ret0, _ := ret[0].(*iot.DescribeAuthorizerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeAuthorizer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAuthorizer, reflect.TypeOf((*MockIotClient)(nil).DescribeAuthorizer), varargs...)
}

func (m *MockIotClient) DescribeBillingGroup(arg0 context.Context, arg1 *iot.DescribeBillingGroupInput, arg2 ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBillingGroup, varargs...)
	ret0, _ := ret[0].(*iot.DescribeBillingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeBillingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBillingGroup, reflect.TypeOf((*MockIotClient)(nil).DescribeBillingGroup), varargs...)
}

func (m *MockIotClient) DescribeCACertificate(arg0 context.Context, arg1 *iot.DescribeCACertificateInput, arg2 ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCACertificate, varargs...)
	ret0, _ := ret[0].(*iot.DescribeCACertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeCACertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCACertificate, reflect.TypeOf((*MockIotClient)(nil).DescribeCACertificate), varargs...)
}

func (m *MockIotClient) DescribeCertificate(arg0 context.Context, arg1 *iot.DescribeCertificateInput, arg2 ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCertificate, varargs...)
	ret0, _ := ret[0].(*iot.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCertificate, reflect.TypeOf((*MockIotClient)(nil).DescribeCertificate), varargs...)
}

func (m *MockIotClient) DescribeCustomMetric(arg0 context.Context, arg1 *iot.DescribeCustomMetricInput, arg2 ...func(*iot.Options)) (*iot.DescribeCustomMetricOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCustomMetric, varargs...)
	ret0, _ := ret[0].(*iot.DescribeCustomMetricOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeCustomMetric(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCustomMetric, reflect.TypeOf((*MockIotClient)(nil).DescribeCustomMetric), varargs...)
}

func (m *MockIotClient) DescribeDefaultAuthorizer(arg0 context.Context, arg1 *iot.DescribeDefaultAuthorizerInput, arg2 ...func(*iot.Options)) (*iot.DescribeDefaultAuthorizerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDefaultAuthorizer, varargs...)
	ret0, _ := ret[0].(*iot.DescribeDefaultAuthorizerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeDefaultAuthorizer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDefaultAuthorizer, reflect.TypeOf((*MockIotClient)(nil).DescribeDefaultAuthorizer), varargs...)
}

func (m *MockIotClient) DescribeDetectMitigationActionsTask(arg0 context.Context, arg1 *iot.DescribeDetectMitigationActionsTaskInput, arg2 ...func(*iot.Options)) (*iot.DescribeDetectMitigationActionsTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDetectMitigationActionsTask, varargs...)
	ret0, _ := ret[0].(*iot.DescribeDetectMitigationActionsTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeDetectMitigationActionsTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDetectMitigationActionsTask, reflect.TypeOf((*MockIotClient)(nil).DescribeDetectMitigationActionsTask), varargs...)
}

func (m *MockIotClient) DescribeDimension(arg0 context.Context, arg1 *iot.DescribeDimensionInput, arg2 ...func(*iot.Options)) (*iot.DescribeDimensionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDimension, varargs...)
	ret0, _ := ret[0].(*iot.DescribeDimensionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeDimension(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDimension, reflect.TypeOf((*MockIotClient)(nil).DescribeDimension), varargs...)
}

func (m *MockIotClient) DescribeDomainConfiguration(arg0 context.Context, arg1 *iot.DescribeDomainConfigurationInput, arg2 ...func(*iot.Options)) (*iot.DescribeDomainConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDomainConfiguration, varargs...)
	ret0, _ := ret[0].(*iot.DescribeDomainConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeDomainConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDomainConfiguration, reflect.TypeOf((*MockIotClient)(nil).DescribeDomainConfiguration), varargs...)
}

func (m *MockIotClient) DescribeEndpoint(arg0 context.Context, arg1 *iot.DescribeEndpointInput, arg2 ...func(*iot.Options)) (*iot.DescribeEndpointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpoint, varargs...)
	ret0, _ := ret[0].(*iot.DescribeEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeEndpoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpoint, reflect.TypeOf((*MockIotClient)(nil).DescribeEndpoint), varargs...)
}

func (m *MockIotClient) DescribeEventConfigurations(arg0 context.Context, arg1 *iot.DescribeEventConfigurationsInput, arg2 ...func(*iot.Options)) (*iot.DescribeEventConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEventConfigurations, varargs...)
	ret0, _ := ret[0].(*iot.DescribeEventConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeEventConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEventConfigurations, reflect.TypeOf((*MockIotClient)(nil).DescribeEventConfigurations), varargs...)
}

func (m *MockIotClient) DescribeFleetMetric(arg0 context.Context, arg1 *iot.DescribeFleetMetricInput, arg2 ...func(*iot.Options)) (*iot.DescribeFleetMetricOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFleetMetric, varargs...)
	ret0, _ := ret[0].(*iot.DescribeFleetMetricOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeFleetMetric(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFleetMetric, reflect.TypeOf((*MockIotClient)(nil).DescribeFleetMetric), varargs...)
}

func (m *MockIotClient) DescribeIndex(arg0 context.Context, arg1 *iot.DescribeIndexInput, arg2 ...func(*iot.Options)) (*iot.DescribeIndexOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIndex, varargs...)
	ret0, _ := ret[0].(*iot.DescribeIndexOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeIndex(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIndex, reflect.TypeOf((*MockIotClient)(nil).DescribeIndex), varargs...)
}

func (m *MockIotClient) DescribeJob(arg0 context.Context, arg1 *iot.DescribeJobInput, arg2 ...func(*iot.Options)) (*iot.DescribeJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJob, varargs...)
	ret0, _ := ret[0].(*iot.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJob, reflect.TypeOf((*MockIotClient)(nil).DescribeJob), varargs...)
}

func (m *MockIotClient) DescribeJobExecution(arg0 context.Context, arg1 *iot.DescribeJobExecutionInput, arg2 ...func(*iot.Options)) (*iot.DescribeJobExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJobExecution, varargs...)
	ret0, _ := ret[0].(*iot.DescribeJobExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeJobExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJobExecution, reflect.TypeOf((*MockIotClient)(nil).DescribeJobExecution), varargs...)
}

func (m *MockIotClient) DescribeJobTemplate(arg0 context.Context, arg1 *iot.DescribeJobTemplateInput, arg2 ...func(*iot.Options)) (*iot.DescribeJobTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJobTemplate, varargs...)
	ret0, _ := ret[0].(*iot.DescribeJobTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeJobTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJobTemplate, reflect.TypeOf((*MockIotClient)(nil).DescribeJobTemplate), varargs...)
}

func (m *MockIotClient) DescribeManagedJobTemplate(arg0 context.Context, arg1 *iot.DescribeManagedJobTemplateInput, arg2 ...func(*iot.Options)) (*iot.DescribeManagedJobTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeManagedJobTemplate, varargs...)
	ret0, _ := ret[0].(*iot.DescribeManagedJobTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeManagedJobTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeManagedJobTemplate, reflect.TypeOf((*MockIotClient)(nil).DescribeManagedJobTemplate), varargs...)
}

func (m *MockIotClient) DescribeMitigationAction(arg0 context.Context, arg1 *iot.DescribeMitigationActionInput, arg2 ...func(*iot.Options)) (*iot.DescribeMitigationActionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMitigationAction, varargs...)
	ret0, _ := ret[0].(*iot.DescribeMitigationActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeMitigationAction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMitigationAction, reflect.TypeOf((*MockIotClient)(nil).DescribeMitigationAction), varargs...)
}

func (m *MockIotClient) DescribeProvisioningTemplate(arg0 context.Context, arg1 *iot.DescribeProvisioningTemplateInput, arg2 ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProvisioningTemplate, varargs...)
	ret0, _ := ret[0].(*iot.DescribeProvisioningTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeProvisioningTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProvisioningTemplate, reflect.TypeOf((*MockIotClient)(nil).DescribeProvisioningTemplate), varargs...)
}

func (m *MockIotClient) DescribeProvisioningTemplateVersion(arg0 context.Context, arg1 *iot.DescribeProvisioningTemplateVersionInput, arg2 ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProvisioningTemplateVersion, varargs...)
	ret0, _ := ret[0].(*iot.DescribeProvisioningTemplateVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeProvisioningTemplateVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProvisioningTemplateVersion, reflect.TypeOf((*MockIotClient)(nil).DescribeProvisioningTemplateVersion), varargs...)
}

func (m *MockIotClient) DescribeRoleAlias(arg0 context.Context, arg1 *iot.DescribeRoleAliasInput, arg2 ...func(*iot.Options)) (*iot.DescribeRoleAliasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRoleAlias, varargs...)
	ret0, _ := ret[0].(*iot.DescribeRoleAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeRoleAlias(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRoleAlias, reflect.TypeOf((*MockIotClient)(nil).DescribeRoleAlias), varargs...)
}

func (m *MockIotClient) DescribeScheduledAudit(arg0 context.Context, arg1 *iot.DescribeScheduledAuditInput, arg2 ...func(*iot.Options)) (*iot.DescribeScheduledAuditOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeScheduledAudit, varargs...)
	ret0, _ := ret[0].(*iot.DescribeScheduledAuditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeScheduledAudit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeScheduledAudit, reflect.TypeOf((*MockIotClient)(nil).DescribeScheduledAudit), varargs...)
}

func (m *MockIotClient) DescribeSecurityProfile(arg0 context.Context, arg1 *iot.DescribeSecurityProfileInput, arg2 ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSecurityProfile, varargs...)
	ret0, _ := ret[0].(*iot.DescribeSecurityProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeSecurityProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSecurityProfile, reflect.TypeOf((*MockIotClient)(nil).DescribeSecurityProfile), varargs...)
}

func (m *MockIotClient) DescribeStream(arg0 context.Context, arg1 *iot.DescribeStreamInput, arg2 ...func(*iot.Options)) (*iot.DescribeStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStream, varargs...)
	ret0, _ := ret[0].(*iot.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStream, reflect.TypeOf((*MockIotClient)(nil).DescribeStream), varargs...)
}

func (m *MockIotClient) DescribeThing(arg0 context.Context, arg1 *iot.DescribeThingInput, arg2 ...func(*iot.Options)) (*iot.DescribeThingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeThing, varargs...)
	ret0, _ := ret[0].(*iot.DescribeThingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeThing(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeThing, reflect.TypeOf((*MockIotClient)(nil).DescribeThing), varargs...)
}

func (m *MockIotClient) DescribeThingGroup(arg0 context.Context, arg1 *iot.DescribeThingGroupInput, arg2 ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeThingGroup, varargs...)
	ret0, _ := ret[0].(*iot.DescribeThingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeThingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeThingGroup, reflect.TypeOf((*MockIotClient)(nil).DescribeThingGroup), varargs...)
}

func (m *MockIotClient) DescribeThingRegistrationTask(arg0 context.Context, arg1 *iot.DescribeThingRegistrationTaskInput, arg2 ...func(*iot.Options)) (*iot.DescribeThingRegistrationTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeThingRegistrationTask, varargs...)
	ret0, _ := ret[0].(*iot.DescribeThingRegistrationTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeThingRegistrationTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeThingRegistrationTask, reflect.TypeOf((*MockIotClient)(nil).DescribeThingRegistrationTask), varargs...)
}

func (m *MockIotClient) DescribeThingType(arg0 context.Context, arg1 *iot.DescribeThingTypeInput, arg2 ...func(*iot.Options)) (*iot.DescribeThingTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeThingType, varargs...)
	ret0, _ := ret[0].(*iot.DescribeThingTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) DescribeThingType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeThingType, reflect.TypeOf((*MockIotClient)(nil).DescribeThingType), varargs...)
}

func (m *MockIotClient) GetBehaviorModelTrainingSummaries(arg0 context.Context, arg1 *iot.GetBehaviorModelTrainingSummariesInput, arg2 ...func(*iot.Options)) (*iot.GetBehaviorModelTrainingSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBehaviorModelTrainingSummaries, varargs...)
	ret0, _ := ret[0].(*iot.GetBehaviorModelTrainingSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetBehaviorModelTrainingSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBehaviorModelTrainingSummaries, reflect.TypeOf((*MockIotClient)(nil).GetBehaviorModelTrainingSummaries), varargs...)
}

func (m *MockIotClient) GetBucketsAggregation(arg0 context.Context, arg1 *iot.GetBucketsAggregationInput, arg2 ...func(*iot.Options)) (*iot.GetBucketsAggregationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketsAggregation, varargs...)
	ret0, _ := ret[0].(*iot.GetBucketsAggregationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetBucketsAggregation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketsAggregation, reflect.TypeOf((*MockIotClient)(nil).GetBucketsAggregation), varargs...)
}

func (m *MockIotClient) GetCardinality(arg0 context.Context, arg1 *iot.GetCardinalityInput, arg2 ...func(*iot.Options)) (*iot.GetCardinalityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCardinality, varargs...)
	ret0, _ := ret[0].(*iot.GetCardinalityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetCardinality(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCardinality, reflect.TypeOf((*MockIotClient)(nil).GetCardinality), varargs...)
}

func (m *MockIotClient) GetEffectivePolicies(arg0 context.Context, arg1 *iot.GetEffectivePoliciesInput, arg2 ...func(*iot.Options)) (*iot.GetEffectivePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEffectivePolicies, varargs...)
	ret0, _ := ret[0].(*iot.GetEffectivePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetEffectivePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEffectivePolicies, reflect.TypeOf((*MockIotClient)(nil).GetEffectivePolicies), varargs...)
}

func (m *MockIotClient) GetIndexingConfiguration(arg0 context.Context, arg1 *iot.GetIndexingConfigurationInput, arg2 ...func(*iot.Options)) (*iot.GetIndexingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIndexingConfiguration, varargs...)
	ret0, _ := ret[0].(*iot.GetIndexingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetIndexingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIndexingConfiguration, reflect.TypeOf((*MockIotClient)(nil).GetIndexingConfiguration), varargs...)
}

func (m *MockIotClient) GetJobDocument(arg0 context.Context, arg1 *iot.GetJobDocumentInput, arg2 ...func(*iot.Options)) (*iot.GetJobDocumentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobDocument, varargs...)
	ret0, _ := ret[0].(*iot.GetJobDocumentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetJobDocument(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobDocument, reflect.TypeOf((*MockIotClient)(nil).GetJobDocument), varargs...)
}

func (m *MockIotClient) GetLoggingOptions(arg0 context.Context, arg1 *iot.GetLoggingOptionsInput, arg2 ...func(*iot.Options)) (*iot.GetLoggingOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoggingOptions, varargs...)
	ret0, _ := ret[0].(*iot.GetLoggingOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetLoggingOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoggingOptions, reflect.TypeOf((*MockIotClient)(nil).GetLoggingOptions), varargs...)
}

func (m *MockIotClient) GetOTAUpdate(arg0 context.Context, arg1 *iot.GetOTAUpdateInput, arg2 ...func(*iot.Options)) (*iot.GetOTAUpdateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOTAUpdate, varargs...)
	ret0, _ := ret[0].(*iot.GetOTAUpdateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetOTAUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOTAUpdate, reflect.TypeOf((*MockIotClient)(nil).GetOTAUpdate), varargs...)
}

func (m *MockIotClient) GetPercentiles(arg0 context.Context, arg1 *iot.GetPercentilesInput, arg2 ...func(*iot.Options)) (*iot.GetPercentilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPercentiles, varargs...)
	ret0, _ := ret[0].(*iot.GetPercentilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetPercentiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPercentiles, reflect.TypeOf((*MockIotClient)(nil).GetPercentiles), varargs...)
}

func (m *MockIotClient) GetPolicy(arg0 context.Context, arg1 *iot.GetPolicyInput, arg2 ...func(*iot.Options)) (*iot.GetPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPolicy, varargs...)
	ret0, _ := ret[0].(*iot.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPolicy, reflect.TypeOf((*MockIotClient)(nil).GetPolicy), varargs...)
}

func (m *MockIotClient) GetPolicyVersion(arg0 context.Context, arg1 *iot.GetPolicyVersionInput, arg2 ...func(*iot.Options)) (*iot.GetPolicyVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPolicyVersion, varargs...)
	ret0, _ := ret[0].(*iot.GetPolicyVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetPolicyVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPolicyVersion, reflect.TypeOf((*MockIotClient)(nil).GetPolicyVersion), varargs...)
}

func (m *MockIotClient) GetRegistrationCode(arg0 context.Context, arg1 *iot.GetRegistrationCodeInput, arg2 ...func(*iot.Options)) (*iot.GetRegistrationCodeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegistrationCode, varargs...)
	ret0, _ := ret[0].(*iot.GetRegistrationCodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetRegistrationCode(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegistrationCode, reflect.TypeOf((*MockIotClient)(nil).GetRegistrationCode), varargs...)
}

func (m *MockIotClient) GetStatistics(arg0 context.Context, arg1 *iot.GetStatisticsInput, arg2 ...func(*iot.Options)) (*iot.GetStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStatistics, varargs...)
	ret0, _ := ret[0].(*iot.GetStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStatistics, reflect.TypeOf((*MockIotClient)(nil).GetStatistics), varargs...)
}

func (m *MockIotClient) GetTopicRule(arg0 context.Context, arg1 *iot.GetTopicRuleInput, arg2 ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTopicRule, varargs...)
	ret0, _ := ret[0].(*iot.GetTopicRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetTopicRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTopicRule, reflect.TypeOf((*MockIotClient)(nil).GetTopicRule), varargs...)
}

func (m *MockIotClient) GetTopicRuleDestination(arg0 context.Context, arg1 *iot.GetTopicRuleDestinationInput, arg2 ...func(*iot.Options)) (*iot.GetTopicRuleDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTopicRuleDestination, varargs...)
	ret0, _ := ret[0].(*iot.GetTopicRuleDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetTopicRuleDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTopicRuleDestination, reflect.TypeOf((*MockIotClient)(nil).GetTopicRuleDestination), varargs...)
}

func (m *MockIotClient) GetV2LoggingOptions(arg0 context.Context, arg1 *iot.GetV2LoggingOptionsInput, arg2 ...func(*iot.Options)) (*iot.GetV2LoggingOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVLoggingOptions, varargs...)
	ret0, _ := ret[0].(*iot.GetV2LoggingOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) GetV2LoggingOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVLoggingOptions, reflect.TypeOf((*MockIotClient)(nil).GetV2LoggingOptions), varargs...)
}

func (m *MockIotClient) ListActiveViolations(arg0 context.Context, arg1 *iot.ListActiveViolationsInput, arg2 ...func(*iot.Options)) (*iot.ListActiveViolationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActiveViolations, varargs...)
	ret0, _ := ret[0].(*iot.ListActiveViolationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListActiveViolations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActiveViolations, reflect.TypeOf((*MockIotClient)(nil).ListActiveViolations), varargs...)
}

func (m *MockIotClient) ListAttachedPolicies(arg0 context.Context, arg1 *iot.ListAttachedPoliciesInput, arg2 ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttachedPolicies, varargs...)
	ret0, _ := ret[0].(*iot.ListAttachedPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAttachedPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttachedPolicies, reflect.TypeOf((*MockIotClient)(nil).ListAttachedPolicies), varargs...)
}

func (m *MockIotClient) ListAuditFindings(arg0 context.Context, arg1 *iot.ListAuditFindingsInput, arg2 ...func(*iot.Options)) (*iot.ListAuditFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAuditFindings, varargs...)
	ret0, _ := ret[0].(*iot.ListAuditFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAuditFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAuditFindings, reflect.TypeOf((*MockIotClient)(nil).ListAuditFindings), varargs...)
}

func (m *MockIotClient) ListAuditMitigationActionsExecutions(arg0 context.Context, arg1 *iot.ListAuditMitigationActionsExecutionsInput, arg2 ...func(*iot.Options)) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAuditMitigationActionsExecutions, varargs...)
	ret0, _ := ret[0].(*iot.ListAuditMitigationActionsExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAuditMitigationActionsExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAuditMitigationActionsExecutions, reflect.TypeOf((*MockIotClient)(nil).ListAuditMitigationActionsExecutions), varargs...)
}

func (m *MockIotClient) ListAuditMitigationActionsTasks(arg0 context.Context, arg1 *iot.ListAuditMitigationActionsTasksInput, arg2 ...func(*iot.Options)) (*iot.ListAuditMitigationActionsTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAuditMitigationActionsTasks, varargs...)
	ret0, _ := ret[0].(*iot.ListAuditMitigationActionsTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAuditMitigationActionsTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAuditMitigationActionsTasks, reflect.TypeOf((*MockIotClient)(nil).ListAuditMitigationActionsTasks), varargs...)
}

func (m *MockIotClient) ListAuditSuppressions(arg0 context.Context, arg1 *iot.ListAuditSuppressionsInput, arg2 ...func(*iot.Options)) (*iot.ListAuditSuppressionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAuditSuppressions, varargs...)
	ret0, _ := ret[0].(*iot.ListAuditSuppressionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAuditSuppressions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAuditSuppressions, reflect.TypeOf((*MockIotClient)(nil).ListAuditSuppressions), varargs...)
}

func (m *MockIotClient) ListAuditTasks(arg0 context.Context, arg1 *iot.ListAuditTasksInput, arg2 ...func(*iot.Options)) (*iot.ListAuditTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAuditTasks, varargs...)
	ret0, _ := ret[0].(*iot.ListAuditTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAuditTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAuditTasks, reflect.TypeOf((*MockIotClient)(nil).ListAuditTasks), varargs...)
}

func (m *MockIotClient) ListAuthorizers(arg0 context.Context, arg1 *iot.ListAuthorizersInput, arg2 ...func(*iot.Options)) (*iot.ListAuthorizersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAuthorizers, varargs...)
	ret0, _ := ret[0].(*iot.ListAuthorizersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListAuthorizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAuthorizers, reflect.TypeOf((*MockIotClient)(nil).ListAuthorizers), varargs...)
}

func (m *MockIotClient) ListBillingGroups(arg0 context.Context, arg1 *iot.ListBillingGroupsInput, arg2 ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBillingGroups, varargs...)
	ret0, _ := ret[0].(*iot.ListBillingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListBillingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBillingGroups, reflect.TypeOf((*MockIotClient)(nil).ListBillingGroups), varargs...)
}

func (m *MockIotClient) ListCACertificates(arg0 context.Context, arg1 *iot.ListCACertificatesInput, arg2 ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCACertificates, varargs...)
	ret0, _ := ret[0].(*iot.ListCACertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListCACertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCACertificates, reflect.TypeOf((*MockIotClient)(nil).ListCACertificates), varargs...)
}

func (m *MockIotClient) ListCertificates(arg0 context.Context, arg1 *iot.ListCertificatesInput, arg2 ...func(*iot.Options)) (*iot.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCertificates, varargs...)
	ret0, _ := ret[0].(*iot.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCertificates, reflect.TypeOf((*MockIotClient)(nil).ListCertificates), varargs...)
}

func (m *MockIotClient) ListCertificatesByCA(arg0 context.Context, arg1 *iot.ListCertificatesByCAInput, arg2 ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCertificatesByCA, varargs...)
	ret0, _ := ret[0].(*iot.ListCertificatesByCAOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListCertificatesByCA(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCertificatesByCA, reflect.TypeOf((*MockIotClient)(nil).ListCertificatesByCA), varargs...)
}

func (m *MockIotClient) ListCustomMetrics(arg0 context.Context, arg1 *iot.ListCustomMetricsInput, arg2 ...func(*iot.Options)) (*iot.ListCustomMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCustomMetrics, varargs...)
	ret0, _ := ret[0].(*iot.ListCustomMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListCustomMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCustomMetrics, reflect.TypeOf((*MockIotClient)(nil).ListCustomMetrics), varargs...)
}

func (m *MockIotClient) ListDetectMitigationActionsExecutions(arg0 context.Context, arg1 *iot.ListDetectMitigationActionsExecutionsInput, arg2 ...func(*iot.Options)) (*iot.ListDetectMitigationActionsExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDetectMitigationActionsExecutions, varargs...)
	ret0, _ := ret[0].(*iot.ListDetectMitigationActionsExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListDetectMitigationActionsExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDetectMitigationActionsExecutions, reflect.TypeOf((*MockIotClient)(nil).ListDetectMitigationActionsExecutions), varargs...)
}

func (m *MockIotClient) ListDetectMitigationActionsTasks(arg0 context.Context, arg1 *iot.ListDetectMitigationActionsTasksInput, arg2 ...func(*iot.Options)) (*iot.ListDetectMitigationActionsTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDetectMitigationActionsTasks, varargs...)
	ret0, _ := ret[0].(*iot.ListDetectMitigationActionsTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListDetectMitigationActionsTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDetectMitigationActionsTasks, reflect.TypeOf((*MockIotClient)(nil).ListDetectMitigationActionsTasks), varargs...)
}

func (m *MockIotClient) ListDimensions(arg0 context.Context, arg1 *iot.ListDimensionsInput, arg2 ...func(*iot.Options)) (*iot.ListDimensionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDimensions, varargs...)
	ret0, _ := ret[0].(*iot.ListDimensionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListDimensions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDimensions, reflect.TypeOf((*MockIotClient)(nil).ListDimensions), varargs...)
}

func (m *MockIotClient) ListDomainConfigurations(arg0 context.Context, arg1 *iot.ListDomainConfigurationsInput, arg2 ...func(*iot.Options)) (*iot.ListDomainConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomainConfigurations, varargs...)
	ret0, _ := ret[0].(*iot.ListDomainConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListDomainConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomainConfigurations, reflect.TypeOf((*MockIotClient)(nil).ListDomainConfigurations), varargs...)
}

func (m *MockIotClient) ListFleetMetrics(arg0 context.Context, arg1 *iot.ListFleetMetricsInput, arg2 ...func(*iot.Options)) (*iot.ListFleetMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFleetMetrics, varargs...)
	ret0, _ := ret[0].(*iot.ListFleetMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListFleetMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFleetMetrics, reflect.TypeOf((*MockIotClient)(nil).ListFleetMetrics), varargs...)
}

func (m *MockIotClient) ListIndices(arg0 context.Context, arg1 *iot.ListIndicesInput, arg2 ...func(*iot.Options)) (*iot.ListIndicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIndices, varargs...)
	ret0, _ := ret[0].(*iot.ListIndicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListIndices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIndices, reflect.TypeOf((*MockIotClient)(nil).ListIndices), varargs...)
}

func (m *MockIotClient) ListJobExecutionsForJob(arg0 context.Context, arg1 *iot.ListJobExecutionsForJobInput, arg2 ...func(*iot.Options)) (*iot.ListJobExecutionsForJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobExecutionsForJob, varargs...)
	ret0, _ := ret[0].(*iot.ListJobExecutionsForJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListJobExecutionsForJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobExecutionsForJob, reflect.TypeOf((*MockIotClient)(nil).ListJobExecutionsForJob), varargs...)
}

func (m *MockIotClient) ListJobExecutionsForThing(arg0 context.Context, arg1 *iot.ListJobExecutionsForThingInput, arg2 ...func(*iot.Options)) (*iot.ListJobExecutionsForThingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobExecutionsForThing, varargs...)
	ret0, _ := ret[0].(*iot.ListJobExecutionsForThingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListJobExecutionsForThing(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobExecutionsForThing, reflect.TypeOf((*MockIotClient)(nil).ListJobExecutionsForThing), varargs...)
}

func (m *MockIotClient) ListJobTemplates(arg0 context.Context, arg1 *iot.ListJobTemplatesInput, arg2 ...func(*iot.Options)) (*iot.ListJobTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobTemplates, varargs...)
	ret0, _ := ret[0].(*iot.ListJobTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListJobTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobTemplates, reflect.TypeOf((*MockIotClient)(nil).ListJobTemplates), varargs...)
}

func (m *MockIotClient) ListJobs(arg0 context.Context, arg1 *iot.ListJobsInput, arg2 ...func(*iot.Options)) (*iot.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobs, varargs...)
	ret0, _ := ret[0].(*iot.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobs, reflect.TypeOf((*MockIotClient)(nil).ListJobs), varargs...)
}

func (m *MockIotClient) ListManagedJobTemplates(arg0 context.Context, arg1 *iot.ListManagedJobTemplatesInput, arg2 ...func(*iot.Options)) (*iot.ListManagedJobTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListManagedJobTemplates, varargs...)
	ret0, _ := ret[0].(*iot.ListManagedJobTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListManagedJobTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListManagedJobTemplates, reflect.TypeOf((*MockIotClient)(nil).ListManagedJobTemplates), varargs...)
}

func (m *MockIotClient) ListMetricValues(arg0 context.Context, arg1 *iot.ListMetricValuesInput, arg2 ...func(*iot.Options)) (*iot.ListMetricValuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMetricValues, varargs...)
	ret0, _ := ret[0].(*iot.ListMetricValuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListMetricValues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMetricValues, reflect.TypeOf((*MockIotClient)(nil).ListMetricValues), varargs...)
}

func (m *MockIotClient) ListMitigationActions(arg0 context.Context, arg1 *iot.ListMitigationActionsInput, arg2 ...func(*iot.Options)) (*iot.ListMitigationActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMitigationActions, varargs...)
	ret0, _ := ret[0].(*iot.ListMitigationActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListMitigationActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMitigationActions, reflect.TypeOf((*MockIotClient)(nil).ListMitigationActions), varargs...)
}

func (m *MockIotClient) ListOTAUpdates(arg0 context.Context, arg1 *iot.ListOTAUpdatesInput, arg2 ...func(*iot.Options)) (*iot.ListOTAUpdatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOTAUpdates, varargs...)
	ret0, _ := ret[0].(*iot.ListOTAUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListOTAUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOTAUpdates, reflect.TypeOf((*MockIotClient)(nil).ListOTAUpdates), varargs...)
}

func (m *MockIotClient) ListOutgoingCertificates(arg0 context.Context, arg1 *iot.ListOutgoingCertificatesInput, arg2 ...func(*iot.Options)) (*iot.ListOutgoingCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOutgoingCertificates, varargs...)
	ret0, _ := ret[0].(*iot.ListOutgoingCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListOutgoingCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOutgoingCertificates, reflect.TypeOf((*MockIotClient)(nil).ListOutgoingCertificates), varargs...)
}

func (m *MockIotClient) ListPolicies(arg0 context.Context, arg1 *iot.ListPoliciesInput, arg2 ...func(*iot.Options)) (*iot.ListPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicies, varargs...)
	ret0, _ := ret[0].(*iot.ListPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicies, reflect.TypeOf((*MockIotClient)(nil).ListPolicies), varargs...)
}

func (m *MockIotClient) ListPolicyPrincipals(arg0 context.Context, arg1 *iot.ListPolicyPrincipalsInput, arg2 ...func(*iot.Options)) (*iot.ListPolicyPrincipalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicyPrincipals, varargs...)
	ret0, _ := ret[0].(*iot.ListPolicyPrincipalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListPolicyPrincipals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicyPrincipals, reflect.TypeOf((*MockIotClient)(nil).ListPolicyPrincipals), varargs...)
}

func (m *MockIotClient) ListPolicyVersions(arg0 context.Context, arg1 *iot.ListPolicyVersionsInput, arg2 ...func(*iot.Options)) (*iot.ListPolicyVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPolicyVersions, varargs...)
	ret0, _ := ret[0].(*iot.ListPolicyVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListPolicyVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPolicyVersions, reflect.TypeOf((*MockIotClient)(nil).ListPolicyVersions), varargs...)
}

func (m *MockIotClient) ListPrincipalPolicies(arg0 context.Context, arg1 *iot.ListPrincipalPoliciesInput, arg2 ...func(*iot.Options)) (*iot.ListPrincipalPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPrincipalPolicies, varargs...)
	ret0, _ := ret[0].(*iot.ListPrincipalPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListPrincipalPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPrincipalPolicies, reflect.TypeOf((*MockIotClient)(nil).ListPrincipalPolicies), varargs...)
}

func (m *MockIotClient) ListPrincipalThings(arg0 context.Context, arg1 *iot.ListPrincipalThingsInput, arg2 ...func(*iot.Options)) (*iot.ListPrincipalThingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPrincipalThings, varargs...)
	ret0, _ := ret[0].(*iot.ListPrincipalThingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListPrincipalThings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPrincipalThings, reflect.TypeOf((*MockIotClient)(nil).ListPrincipalThings), varargs...)
}

func (m *MockIotClient) ListProvisioningTemplateVersions(arg0 context.Context, arg1 *iot.ListProvisioningTemplateVersionsInput, arg2 ...func(*iot.Options)) (*iot.ListProvisioningTemplateVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisioningTemplateVersions, varargs...)
	ret0, _ := ret[0].(*iot.ListProvisioningTemplateVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListProvisioningTemplateVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisioningTemplateVersions, reflect.TypeOf((*MockIotClient)(nil).ListProvisioningTemplateVersions), varargs...)
}

func (m *MockIotClient) ListProvisioningTemplates(arg0 context.Context, arg1 *iot.ListProvisioningTemplatesInput, arg2 ...func(*iot.Options)) (*iot.ListProvisioningTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisioningTemplates, varargs...)
	ret0, _ := ret[0].(*iot.ListProvisioningTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListProvisioningTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisioningTemplates, reflect.TypeOf((*MockIotClient)(nil).ListProvisioningTemplates), varargs...)
}

func (m *MockIotClient) ListRelatedResourcesForAuditFinding(arg0 context.Context, arg1 *iot.ListRelatedResourcesForAuditFindingInput, arg2 ...func(*iot.Options)) (*iot.ListRelatedResourcesForAuditFindingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRelatedResourcesForAuditFinding, varargs...)
	ret0, _ := ret[0].(*iot.ListRelatedResourcesForAuditFindingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListRelatedResourcesForAuditFinding(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRelatedResourcesForAuditFinding, reflect.TypeOf((*MockIotClient)(nil).ListRelatedResourcesForAuditFinding), varargs...)
}

func (m *MockIotClient) ListRoleAliases(arg0 context.Context, arg1 *iot.ListRoleAliasesInput, arg2 ...func(*iot.Options)) (*iot.ListRoleAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRoleAliases, varargs...)
	ret0, _ := ret[0].(*iot.ListRoleAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListRoleAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRoleAliases, reflect.TypeOf((*MockIotClient)(nil).ListRoleAliases), varargs...)
}

func (m *MockIotClient) ListScheduledAudits(arg0 context.Context, arg1 *iot.ListScheduledAuditsInput, arg2 ...func(*iot.Options)) (*iot.ListScheduledAuditsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListScheduledAudits, varargs...)
	ret0, _ := ret[0].(*iot.ListScheduledAuditsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListScheduledAudits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListScheduledAudits, reflect.TypeOf((*MockIotClient)(nil).ListScheduledAudits), varargs...)
}

func (m *MockIotClient) ListSecurityProfiles(arg0 context.Context, arg1 *iot.ListSecurityProfilesInput, arg2 ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSecurityProfiles, varargs...)
	ret0, _ := ret[0].(*iot.ListSecurityProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListSecurityProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSecurityProfiles, reflect.TypeOf((*MockIotClient)(nil).ListSecurityProfiles), varargs...)
}

func (m *MockIotClient) ListSecurityProfilesForTarget(arg0 context.Context, arg1 *iot.ListSecurityProfilesForTargetInput, arg2 ...func(*iot.Options)) (*iot.ListSecurityProfilesForTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSecurityProfilesForTarget, varargs...)
	ret0, _ := ret[0].(*iot.ListSecurityProfilesForTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListSecurityProfilesForTarget(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSecurityProfilesForTarget, reflect.TypeOf((*MockIotClient)(nil).ListSecurityProfilesForTarget), varargs...)
}

func (m *MockIotClient) ListStreams(arg0 context.Context, arg1 *iot.ListStreamsInput, arg2 ...func(*iot.Options)) (*iot.ListStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStreams, varargs...)
	ret0, _ := ret[0].(*iot.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStreams, reflect.TypeOf((*MockIotClient)(nil).ListStreams), varargs...)
}

func (m *MockIotClient) ListTagsForResource(arg0 context.Context, arg1 *iot.ListTagsForResourceInput, arg2 ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*iot.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockIotClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockIotClient) ListTargetsForPolicy(arg0 context.Context, arg1 *iot.ListTargetsForPolicyInput, arg2 ...func(*iot.Options)) (*iot.ListTargetsForPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTargetsForPolicy, varargs...)
	ret0, _ := ret[0].(*iot.ListTargetsForPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListTargetsForPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTargetsForPolicy, reflect.TypeOf((*MockIotClient)(nil).ListTargetsForPolicy), varargs...)
}

func (m *MockIotClient) ListTargetsForSecurityProfile(arg0 context.Context, arg1 *iot.ListTargetsForSecurityProfileInput, arg2 ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTargetsForSecurityProfile, varargs...)
	ret0, _ := ret[0].(*iot.ListTargetsForSecurityProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListTargetsForSecurityProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTargetsForSecurityProfile, reflect.TypeOf((*MockIotClient)(nil).ListTargetsForSecurityProfile), varargs...)
}

func (m *MockIotClient) ListThingGroups(arg0 context.Context, arg1 *iot.ListThingGroupsInput, arg2 ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingGroups, varargs...)
	ret0, _ := ret[0].(*iot.ListThingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingGroups, reflect.TypeOf((*MockIotClient)(nil).ListThingGroups), varargs...)
}

func (m *MockIotClient) ListThingGroupsForThing(arg0 context.Context, arg1 *iot.ListThingGroupsForThingInput, arg2 ...func(*iot.Options)) (*iot.ListThingGroupsForThingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingGroupsForThing, varargs...)
	ret0, _ := ret[0].(*iot.ListThingGroupsForThingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingGroupsForThing(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingGroupsForThing, reflect.TypeOf((*MockIotClient)(nil).ListThingGroupsForThing), varargs...)
}

func (m *MockIotClient) ListThingPrincipals(arg0 context.Context, arg1 *iot.ListThingPrincipalsInput, arg2 ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingPrincipals, varargs...)
	ret0, _ := ret[0].(*iot.ListThingPrincipalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingPrincipals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingPrincipals, reflect.TypeOf((*MockIotClient)(nil).ListThingPrincipals), varargs...)
}

func (m *MockIotClient) ListThingRegistrationTaskReports(arg0 context.Context, arg1 *iot.ListThingRegistrationTaskReportsInput, arg2 ...func(*iot.Options)) (*iot.ListThingRegistrationTaskReportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingRegistrationTaskReports, varargs...)
	ret0, _ := ret[0].(*iot.ListThingRegistrationTaskReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingRegistrationTaskReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingRegistrationTaskReports, reflect.TypeOf((*MockIotClient)(nil).ListThingRegistrationTaskReports), varargs...)
}

func (m *MockIotClient) ListThingRegistrationTasks(arg0 context.Context, arg1 *iot.ListThingRegistrationTasksInput, arg2 ...func(*iot.Options)) (*iot.ListThingRegistrationTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingRegistrationTasks, varargs...)
	ret0, _ := ret[0].(*iot.ListThingRegistrationTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingRegistrationTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingRegistrationTasks, reflect.TypeOf((*MockIotClient)(nil).ListThingRegistrationTasks), varargs...)
}

func (m *MockIotClient) ListThingTypes(arg0 context.Context, arg1 *iot.ListThingTypesInput, arg2 ...func(*iot.Options)) (*iot.ListThingTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingTypes, varargs...)
	ret0, _ := ret[0].(*iot.ListThingTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingTypes, reflect.TypeOf((*MockIotClient)(nil).ListThingTypes), varargs...)
}

func (m *MockIotClient) ListThings(arg0 context.Context, arg1 *iot.ListThingsInput, arg2 ...func(*iot.Options)) (*iot.ListThingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThings, varargs...)
	ret0, _ := ret[0].(*iot.ListThingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThings, reflect.TypeOf((*MockIotClient)(nil).ListThings), varargs...)
}

func (m *MockIotClient) ListThingsInBillingGroup(arg0 context.Context, arg1 *iot.ListThingsInBillingGroupInput, arg2 ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingsInBillingGroup, varargs...)
	ret0, _ := ret[0].(*iot.ListThingsInBillingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingsInBillingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingsInBillingGroup, reflect.TypeOf((*MockIotClient)(nil).ListThingsInBillingGroup), varargs...)
}

func (m *MockIotClient) ListThingsInThingGroup(arg0 context.Context, arg1 *iot.ListThingsInThingGroupInput, arg2 ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListThingsInThingGroup, varargs...)
	ret0, _ := ret[0].(*iot.ListThingsInThingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListThingsInThingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListThingsInThingGroup, reflect.TypeOf((*MockIotClient)(nil).ListThingsInThingGroup), varargs...)
}

func (m *MockIotClient) ListTopicRuleDestinations(arg0 context.Context, arg1 *iot.ListTopicRuleDestinationsInput, arg2 ...func(*iot.Options)) (*iot.ListTopicRuleDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTopicRuleDestinations, varargs...)
	ret0, _ := ret[0].(*iot.ListTopicRuleDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListTopicRuleDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTopicRuleDestinations, reflect.TypeOf((*MockIotClient)(nil).ListTopicRuleDestinations), varargs...)
}

func (m *MockIotClient) ListTopicRules(arg0 context.Context, arg1 *iot.ListTopicRulesInput, arg2 ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTopicRules, varargs...)
	ret0, _ := ret[0].(*iot.ListTopicRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListTopicRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTopicRules, reflect.TypeOf((*MockIotClient)(nil).ListTopicRules), varargs...)
}

func (m *MockIotClient) ListV2LoggingLevels(arg0 context.Context, arg1 *iot.ListV2LoggingLevelsInput, arg2 ...func(*iot.Options)) (*iot.ListV2LoggingLevelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVLoggingLevels, varargs...)
	ret0, _ := ret[0].(*iot.ListV2LoggingLevelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListV2LoggingLevels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVLoggingLevels, reflect.TypeOf((*MockIotClient)(nil).ListV2LoggingLevels), varargs...)
}

func (m *MockIotClient) ListViolationEvents(arg0 context.Context, arg1 *iot.ListViolationEventsInput, arg2 ...func(*iot.Options)) (*iot.ListViolationEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListViolationEvents, varargs...)
	ret0, _ := ret[0].(*iot.ListViolationEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) ListViolationEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListViolationEvents, reflect.TypeOf((*MockIotClient)(nil).ListViolationEvents), varargs...)
}

func (m *MockIotClient) SearchIndex(arg0 context.Context, arg1 *iot.SearchIndexInput, arg2 ...func(*iot.Options)) (*iot.SearchIndexOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchIndex, varargs...)
	ret0, _ := ret[0].(*iot.SearchIndexOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotClientMockRecorder) SearchIndex(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchIndex, reflect.TypeOf((*MockIotClient)(nil).SearchIndex), varargs...)
}
