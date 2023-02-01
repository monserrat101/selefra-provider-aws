package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cloudfront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudfrontClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudfrontClientMockRecorder
}

type MockCloudfrontClientMockRecorder struct {
	mock *MockCloudfrontClient
}

func NewMockCloudfrontClient(ctrl *gomock.Controller) *MockCloudfrontClient {
	mock := &MockCloudfrontClient{ctrl: ctrl}
	mock.recorder = &MockCloudfrontClientMockRecorder{mock}
	return mock
}

func (m *MockCloudfrontClient) EXPECT() *MockCloudfrontClientMockRecorder {
	return m.recorder
}

func (m *MockCloudfrontClient) DescribeFunction(arg0 context.Context, arg1 *cloudfront.DescribeFunctionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.DescribeFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFunction, varargs...)
	ret0, _ := ret[0].(*cloudfront.DescribeFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) DescribeFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFunction, reflect.TypeOf((*MockCloudfrontClient)(nil).DescribeFunction), varargs...)
}

func (m *MockCloudfrontClient) GetCachePolicy(arg0 context.Context, arg1 *cloudfront.GetCachePolicyInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCachePolicy, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetCachePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetCachePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCachePolicy, reflect.TypeOf((*MockCloudfrontClient)(nil).GetCachePolicy), varargs...)
}

func (m *MockCloudfrontClient) GetCachePolicyConfig(arg0 context.Context, arg1 *cloudfront.GetCachePolicyConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCachePolicyConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetCachePolicyConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetCachePolicyConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCachePolicyConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetCachePolicyConfig), varargs...)
}

func (m *MockCloudfrontClient) GetCloudFrontOriginAccessIdentity(arg0 context.Context, arg1 *cloudfront.GetCloudFrontOriginAccessIdentityInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCloudFrontOriginAccessIdentity, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetCloudFrontOriginAccessIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetCloudFrontOriginAccessIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCloudFrontOriginAccessIdentity, reflect.TypeOf((*MockCloudfrontClient)(nil).GetCloudFrontOriginAccessIdentity), varargs...)
}

func (m *MockCloudfrontClient) GetCloudFrontOriginAccessIdentityConfig(arg0 context.Context, arg1 *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCloudFrontOriginAccessIdentityConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetCloudFrontOriginAccessIdentityConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCloudFrontOriginAccessIdentityConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetCloudFrontOriginAccessIdentityConfig), varargs...)
}

func (m *MockCloudfrontClient) GetContinuousDeploymentPolicy(arg0 context.Context, arg1 *cloudfront.GetContinuousDeploymentPolicyInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetContinuousDeploymentPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContinuousDeploymentPolicy, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetContinuousDeploymentPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetContinuousDeploymentPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContinuousDeploymentPolicy, reflect.TypeOf((*MockCloudfrontClient)(nil).GetContinuousDeploymentPolicy), varargs...)
}

func (m *MockCloudfrontClient) GetContinuousDeploymentPolicyConfig(arg0 context.Context, arg1 *cloudfront.GetContinuousDeploymentPolicyConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetContinuousDeploymentPolicyConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContinuousDeploymentPolicyConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetContinuousDeploymentPolicyConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetContinuousDeploymentPolicyConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContinuousDeploymentPolicyConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetContinuousDeploymentPolicyConfig), varargs...)
}

func (m *MockCloudfrontClient) GetDistribution(arg0 context.Context, arg1 *cloudfront.GetDistributionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetDistributionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDistribution, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetDistributionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetDistribution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDistribution, reflect.TypeOf((*MockCloudfrontClient)(nil).GetDistribution), varargs...)
}

func (m *MockCloudfrontClient) GetDistributionConfig(arg0 context.Context, arg1 *cloudfront.GetDistributionConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetDistributionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDistributionConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetDistributionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetDistributionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDistributionConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetDistributionConfig), varargs...)
}

func (m *MockCloudfrontClient) GetFieldLevelEncryption(arg0 context.Context, arg1 *cloudfront.GetFieldLevelEncryptionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFieldLevelEncryption, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetFieldLevelEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetFieldLevelEncryption(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFieldLevelEncryption, reflect.TypeOf((*MockCloudfrontClient)(nil).GetFieldLevelEncryption), varargs...)
}

func (m *MockCloudfrontClient) GetFieldLevelEncryptionConfig(arg0 context.Context, arg1 *cloudfront.GetFieldLevelEncryptionConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFieldLevelEncryptionConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetFieldLevelEncryptionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetFieldLevelEncryptionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFieldLevelEncryptionConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetFieldLevelEncryptionConfig), varargs...)
}

func (m *MockCloudfrontClient) GetFieldLevelEncryptionProfile(arg0 context.Context, arg1 *cloudfront.GetFieldLevelEncryptionProfileInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFieldLevelEncryptionProfile, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetFieldLevelEncryptionProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetFieldLevelEncryptionProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFieldLevelEncryptionProfile, reflect.TypeOf((*MockCloudfrontClient)(nil).GetFieldLevelEncryptionProfile), varargs...)
}

func (m *MockCloudfrontClient) GetFieldLevelEncryptionProfileConfig(arg0 context.Context, arg1 *cloudfront.GetFieldLevelEncryptionProfileConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFieldLevelEncryptionProfileConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetFieldLevelEncryptionProfileConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetFieldLevelEncryptionProfileConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFieldLevelEncryptionProfileConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetFieldLevelEncryptionProfileConfig), varargs...)
}

func (m *MockCloudfrontClient) GetFunction(arg0 context.Context, arg1 *cloudfront.GetFunctionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunction, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunction, reflect.TypeOf((*MockCloudfrontClient)(nil).GetFunction), varargs...)
}

func (m *MockCloudfrontClient) GetInvalidation(arg0 context.Context, arg1 *cloudfront.GetInvalidationInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetInvalidationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInvalidation, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetInvalidationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetInvalidation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInvalidation, reflect.TypeOf((*MockCloudfrontClient)(nil).GetInvalidation), varargs...)
}

func (m *MockCloudfrontClient) GetKeyGroup(arg0 context.Context, arg1 *cloudfront.GetKeyGroupInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKeyGroup, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetKeyGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetKeyGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKeyGroup, reflect.TypeOf((*MockCloudfrontClient)(nil).GetKeyGroup), varargs...)
}

func (m *MockCloudfrontClient) GetKeyGroupConfig(arg0 context.Context, arg1 *cloudfront.GetKeyGroupConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKeyGroupConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetKeyGroupConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetKeyGroupConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKeyGroupConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetKeyGroupConfig), varargs...)
}

func (m *MockCloudfrontClient) GetMonitoringSubscription(arg0 context.Context, arg1 *cloudfront.GetMonitoringSubscriptionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMonitoringSubscription, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetMonitoringSubscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetMonitoringSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMonitoringSubscription, reflect.TypeOf((*MockCloudfrontClient)(nil).GetMonitoringSubscription), varargs...)
}

func (m *MockCloudfrontClient) GetOriginAccessControl(arg0 context.Context, arg1 *cloudfront.GetOriginAccessControlInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOriginAccessControl, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetOriginAccessControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetOriginAccessControl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOriginAccessControl, reflect.TypeOf((*MockCloudfrontClient)(nil).GetOriginAccessControl), varargs...)
}

func (m *MockCloudfrontClient) GetOriginAccessControlConfig(arg0 context.Context, arg1 *cloudfront.GetOriginAccessControlConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOriginAccessControlConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetOriginAccessControlConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetOriginAccessControlConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOriginAccessControlConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetOriginAccessControlConfig), varargs...)
}

func (m *MockCloudfrontClient) GetOriginRequestPolicy(arg0 context.Context, arg1 *cloudfront.GetOriginRequestPolicyInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOriginRequestPolicy, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetOriginRequestPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetOriginRequestPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOriginRequestPolicy, reflect.TypeOf((*MockCloudfrontClient)(nil).GetOriginRequestPolicy), varargs...)
}

func (m *MockCloudfrontClient) GetOriginRequestPolicyConfig(arg0 context.Context, arg1 *cloudfront.GetOriginRequestPolicyConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOriginRequestPolicyConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetOriginRequestPolicyConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetOriginRequestPolicyConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOriginRequestPolicyConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetOriginRequestPolicyConfig), varargs...)
}

func (m *MockCloudfrontClient) GetPublicKey(arg0 context.Context, arg1 *cloudfront.GetPublicKeyInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPublicKey, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetPublicKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPublicKey, reflect.TypeOf((*MockCloudfrontClient)(nil).GetPublicKey), varargs...)
}

func (m *MockCloudfrontClient) GetPublicKeyConfig(arg0 context.Context, arg1 *cloudfront.GetPublicKeyConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPublicKeyConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetPublicKeyConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetPublicKeyConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPublicKeyConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetPublicKeyConfig), varargs...)
}

func (m *MockCloudfrontClient) GetRealtimeLogConfig(arg0 context.Context, arg1 *cloudfront.GetRealtimeLogConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetRealtimeLogConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRealtimeLogConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetRealtimeLogConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetRealtimeLogConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRealtimeLogConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetRealtimeLogConfig), varargs...)
}

func (m *MockCloudfrontClient) GetResponseHeadersPolicy(arg0 context.Context, arg1 *cloudfront.GetResponseHeadersPolicyInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResponseHeadersPolicy, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetResponseHeadersPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetResponseHeadersPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResponseHeadersPolicy, reflect.TypeOf((*MockCloudfrontClient)(nil).GetResponseHeadersPolicy), varargs...)
}

func (m *MockCloudfrontClient) GetResponseHeadersPolicyConfig(arg0 context.Context, arg1 *cloudfront.GetResponseHeadersPolicyConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResponseHeadersPolicyConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetResponseHeadersPolicyConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetResponseHeadersPolicyConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResponseHeadersPolicyConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetResponseHeadersPolicyConfig), varargs...)
}

func (m *MockCloudfrontClient) GetStreamingDistribution(arg0 context.Context, arg1 *cloudfront.GetStreamingDistributionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStreamingDistribution, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetStreamingDistributionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetStreamingDistribution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStreamingDistribution, reflect.TypeOf((*MockCloudfrontClient)(nil).GetStreamingDistribution), varargs...)
}

func (m *MockCloudfrontClient) GetStreamingDistributionConfig(arg0 context.Context, arg1 *cloudfront.GetStreamingDistributionConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStreamingDistributionConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.GetStreamingDistributionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetStreamingDistributionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStreamingDistributionConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).GetStreamingDistributionConfig), varargs...)
}

func (m *MockCloudfrontClient) ListCachePolicies(arg0 context.Context, arg1 *cloudfront.ListCachePoliciesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCachePolicies, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListCachePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListCachePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCachePolicies, reflect.TypeOf((*MockCloudfrontClient)(nil).ListCachePolicies), varargs...)
}

func (m *MockCloudfrontClient) ListCloudFrontOriginAccessIdentities(arg0 context.Context, arg1 *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCloudFrontOriginAccessIdentities, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListCloudFrontOriginAccessIdentities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCloudFrontOriginAccessIdentities, reflect.TypeOf((*MockCloudfrontClient)(nil).ListCloudFrontOriginAccessIdentities), varargs...)
}

func (m *MockCloudfrontClient) ListConflictingAliases(arg0 context.Context, arg1 *cloudfront.ListConflictingAliasesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListConflictingAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConflictingAliases, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListConflictingAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListConflictingAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConflictingAliases, reflect.TypeOf((*MockCloudfrontClient)(nil).ListConflictingAliases), varargs...)
}

func (m *MockCloudfrontClient) ListContinuousDeploymentPolicies(arg0 context.Context, arg1 *cloudfront.ListContinuousDeploymentPoliciesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListContinuousDeploymentPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListContinuousDeploymentPolicies, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListContinuousDeploymentPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListContinuousDeploymentPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListContinuousDeploymentPolicies, reflect.TypeOf((*MockCloudfrontClient)(nil).ListContinuousDeploymentPolicies), varargs...)
}

func (m *MockCloudfrontClient) ListDistributions(arg0 context.Context, arg1 *cloudfront.ListDistributionsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributions, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributions, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributions), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByCachePolicyId(arg0 context.Context, arg1 *cloudfront.ListDistributionsByCachePolicyIdInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributionsByCachePolicyId, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByCachePolicyIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByCachePolicyId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributionsByCachePolicyId, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByCachePolicyId), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByKeyGroup(arg0 context.Context, arg1 *cloudfront.ListDistributionsByKeyGroupInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByKeyGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributionsByKeyGroup, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByKeyGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByKeyGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributionsByKeyGroup, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByKeyGroup), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByOriginRequestPolicyId(arg0 context.Context, arg1 *cloudfront.ListDistributionsByOriginRequestPolicyIdInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributionsByOriginRequestPolicyId, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByOriginRequestPolicyId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributionsByOriginRequestPolicyId, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByOriginRequestPolicyId), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByRealtimeLogConfig(arg0 context.Context, arg1 *cloudfront.ListDistributionsByRealtimeLogConfigInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributionsByRealtimeLogConfig, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByRealtimeLogConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByRealtimeLogConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributionsByRealtimeLogConfig, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByRealtimeLogConfig), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByResponseHeadersPolicyId(arg0 context.Context, arg1 *cloudfront.ListDistributionsByResponseHeadersPolicyIdInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributionsByResponseHeadersPolicyId, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByResponseHeadersPolicyId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributionsByResponseHeadersPolicyId, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByResponseHeadersPolicyId), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByWebACLId(arg0 context.Context, arg1 *cloudfront.ListDistributionsByWebACLIdInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDistributionsByWebACLId, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByWebACLIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByWebACLId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDistributionsByWebACLId, reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByWebACLId), varargs...)
}

func (m *MockCloudfrontClient) ListFieldLevelEncryptionConfigs(arg0 context.Context, arg1 *cloudfront.ListFieldLevelEncryptionConfigsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFieldLevelEncryptionConfigs, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListFieldLevelEncryptionConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListFieldLevelEncryptionConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFieldLevelEncryptionConfigs, reflect.TypeOf((*MockCloudfrontClient)(nil).ListFieldLevelEncryptionConfigs), varargs...)
}

func (m *MockCloudfrontClient) ListFieldLevelEncryptionProfiles(arg0 context.Context, arg1 *cloudfront.ListFieldLevelEncryptionProfilesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFieldLevelEncryptionProfiles, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListFieldLevelEncryptionProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListFieldLevelEncryptionProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFieldLevelEncryptionProfiles, reflect.TypeOf((*MockCloudfrontClient)(nil).ListFieldLevelEncryptionProfiles), varargs...)
}

func (m *MockCloudfrontClient) ListFunctions(arg0 context.Context, arg1 *cloudfront.ListFunctionsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListFunctionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFunctions, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFunctions, reflect.TypeOf((*MockCloudfrontClient)(nil).ListFunctions), varargs...)
}

func (m *MockCloudfrontClient) ListInvalidations(arg0 context.Context, arg1 *cloudfront.ListInvalidationsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListInvalidationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInvalidations, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListInvalidationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListInvalidations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInvalidations, reflect.TypeOf((*MockCloudfrontClient)(nil).ListInvalidations), varargs...)
}

func (m *MockCloudfrontClient) ListKeyGroups(arg0 context.Context, arg1 *cloudfront.ListKeyGroupsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListKeyGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListKeyGroups, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListKeyGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListKeyGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListKeyGroups, reflect.TypeOf((*MockCloudfrontClient)(nil).ListKeyGroups), varargs...)
}

func (m *MockCloudfrontClient) ListOriginAccessControls(arg0 context.Context, arg1 *cloudfront.ListOriginAccessControlsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListOriginAccessControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOriginAccessControls, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListOriginAccessControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListOriginAccessControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOriginAccessControls, reflect.TypeOf((*MockCloudfrontClient)(nil).ListOriginAccessControls), varargs...)
}

func (m *MockCloudfrontClient) ListOriginRequestPolicies(arg0 context.Context, arg1 *cloudfront.ListOriginRequestPoliciesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOriginRequestPolicies, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListOriginRequestPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListOriginRequestPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOriginRequestPolicies, reflect.TypeOf((*MockCloudfrontClient)(nil).ListOriginRequestPolicies), varargs...)
}

func (m *MockCloudfrontClient) ListPublicKeys(arg0 context.Context, arg1 *cloudfront.ListPublicKeysInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListPublicKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPublicKeys, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListPublicKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListPublicKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPublicKeys, reflect.TypeOf((*MockCloudfrontClient)(nil).ListPublicKeys), varargs...)
}

func (m *MockCloudfrontClient) ListRealtimeLogConfigs(arg0 context.Context, arg1 *cloudfront.ListRealtimeLogConfigsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListRealtimeLogConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRealtimeLogConfigs, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListRealtimeLogConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListRealtimeLogConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRealtimeLogConfigs, reflect.TypeOf((*MockCloudfrontClient)(nil).ListRealtimeLogConfigs), varargs...)
}

func (m *MockCloudfrontClient) ListResponseHeadersPolicies(arg0 context.Context, arg1 *cloudfront.ListResponseHeadersPoliciesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListResponseHeadersPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResponseHeadersPolicies, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListResponseHeadersPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListResponseHeadersPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResponseHeadersPolicies, reflect.TypeOf((*MockCloudfrontClient)(nil).ListResponseHeadersPolicies), varargs...)
}

func (m *MockCloudfrontClient) ListStreamingDistributions(arg0 context.Context, arg1 *cloudfront.ListStreamingDistributionsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListStreamingDistributionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStreamingDistributions, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListStreamingDistributionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListStreamingDistributions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStreamingDistributions, reflect.TypeOf((*MockCloudfrontClient)(nil).ListStreamingDistributions), varargs...)
}

func (m *MockCloudfrontClient) ListTagsForResource(arg0 context.Context, arg1 *cloudfront.ListTagsForResourceInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*cloudfront.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockCloudfrontClient)(nil).ListTagsForResource), varargs...)
}
