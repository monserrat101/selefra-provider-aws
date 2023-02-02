package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cognitoidentityprovider "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	gomock "github.com/golang/mock/gomock"
)

type MockCognitoidentityproviderClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCognitoidentityproviderClientMockRecorder
}

type MockCognitoidentityproviderClientMockRecorder struct {
	mock *MockCognitoidentityproviderClient
}

func NewMockCognitoidentityproviderClient(ctrl *gomock.Controller) *MockCognitoidentityproviderClient {
	mock := &MockCognitoidentityproviderClient{ctrl: ctrl}
	mock.recorder = &MockCognitoidentityproviderClientMockRecorder{mock}
	return mock
}

func (m *MockCognitoidentityproviderClient) EXPECT() *MockCognitoidentityproviderClientMockRecorder {
	return m.recorder
}

func (m *MockCognitoidentityproviderClient) DescribeIdentityProvider(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeIdentityProviderInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIdentityProvider, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeIdentityProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeIdentityProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIdentityProvider, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeIdentityProvider), varargs...)
}

func (m *MockCognitoidentityproviderClient) DescribeResourceServer(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeResourceServerInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeResourceServer, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeResourceServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeResourceServer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeResourceServer, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeResourceServer), varargs...)
}

func (m *MockCognitoidentityproviderClient) DescribeRiskConfiguration(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeRiskConfigurationInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRiskConfiguration, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeRiskConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeRiskConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRiskConfiguration, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeRiskConfiguration), varargs...)
}

func (m *MockCognitoidentityproviderClient) DescribeUserImportJob(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserImportJobInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserImportJob, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserImportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserImportJob, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserImportJob), varargs...)
}

func (m *MockCognitoidentityproviderClient) DescribeUserPool(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserPool, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserPool, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserPool), varargs...)
}

func (m *MockCognitoidentityproviderClient) DescribeUserPoolClient(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolClientInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserPoolClient, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolClientOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserPoolClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserPoolClient, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserPoolClient), varargs...)
}

func (m *MockCognitoidentityproviderClient) DescribeUserPoolDomain(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolDomainInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserPoolDomain, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserPoolDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserPoolDomain, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserPoolDomain), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetCSVHeader(arg0 context.Context, arg1 *cognitoidentityprovider.GetCSVHeaderInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCSVHeader, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetCSVHeaderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetCSVHeader(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCSVHeader, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetCSVHeader), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetDevice(arg0 context.Context, arg1 *cognitoidentityprovider.GetDeviceInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDevice, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetDeviceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetDevice(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDevice, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetDevice), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetGroup(arg0 context.Context, arg1 *cognitoidentityprovider.GetGroupInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGroup, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGroup, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetGroup), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetIdentityProviderByIdentifier(arg0 context.Context, arg1 *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityProviderByIdentifier, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetIdentityProviderByIdentifier(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityProviderByIdentifier, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetIdentityProviderByIdentifier), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetSigningCertificate(arg0 context.Context, arg1 *cognitoidentityprovider.GetSigningCertificateInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSigningCertificate, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetSigningCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetSigningCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSigningCertificate, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetSigningCertificate), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetUICustomization(arg0 context.Context, arg1 *cognitoidentityprovider.GetUICustomizationInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUICustomization, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUICustomizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetUICustomization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUICustomization, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUICustomization), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetUser(arg0 context.Context, arg1 *cognitoidentityprovider.GetUserInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUser, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUser, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUser), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetUserAttributeVerificationCode(arg0 context.Context, arg1 *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUserAttributeVerificationCode, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetUserAttributeVerificationCode(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUserAttributeVerificationCode, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUserAttributeVerificationCode), varargs...)
}

func (m *MockCognitoidentityproviderClient) GetUserPoolMfaConfig(arg0 context.Context, arg1 *cognitoidentityprovider.GetUserPoolMfaConfigInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUserPoolMfaConfig, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUserPoolMfaConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) GetUserPoolMfaConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUserPoolMfaConfig, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUserPoolMfaConfig), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListDevices(arg0 context.Context, arg1 *cognitoidentityprovider.ListDevicesInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDevices, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListDevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListDevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDevices, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListDevices), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListGroups(arg0 context.Context, arg1 *cognitoidentityprovider.ListGroupsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGroups, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGroups, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListGroups), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListIdentityProviders(arg0 context.Context, arg1 *cognitoidentityprovider.ListIdentityProvidersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIdentityProviders, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListIdentityProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListIdentityProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIdentityProviders, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListIdentityProviders), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListResourceServers(arg0 context.Context, arg1 *cognitoidentityprovider.ListResourceServersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceServers, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListResourceServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListResourceServers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceServers, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListResourceServers), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListTagsForResource(arg0 context.Context, arg1 *cognitoidentityprovider.ListTagsForResourceInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListUserImportJobs(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserImportJobsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserImportJobs, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListUserImportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserImportJobs, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUserImportJobs), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListUserPoolClients(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserPoolClientsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserPoolClients, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserPoolClientsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListUserPoolClients(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserPoolClients, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUserPoolClients), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListUserPools(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserPoolsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserPools, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListUserPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserPools, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUserPools), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListUsers(arg0 context.Context, arg1 *cognitoidentityprovider.ListUsersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsers, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsers, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUsers), varargs...)
}

func (m *MockCognitoidentityproviderClient) ListUsersInGroup(arg0 context.Context, arg1 *cognitoidentityprovider.ListUsersInGroupInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUsersInGroup, varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUsersInGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityproviderClientMockRecorder) ListUsersInGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUsersInGroup, reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUsersInGroup), varargs...)
}
