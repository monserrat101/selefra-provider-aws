package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	apigateway "github.com/aws/aws-sdk-go-v2/service/apigateway"
	gomock "github.com/golang/mock/gomock"
)

type MockApigatewayClient struct {
	ctrl		*gomock.Controller
	recorder	*MockApigatewayClientMockRecorder
}

type MockApigatewayClientMockRecorder struct {
	mock *MockApigatewayClient
}

func NewMockApigatewayClient(ctrl *gomock.Controller) *MockApigatewayClient {
	mock := &MockApigatewayClient{ctrl: ctrl}
	mock.recorder = &MockApigatewayClientMockRecorder{mock}
	return mock
}

func (m *MockApigatewayClient) EXPECT() *MockApigatewayClientMockRecorder {
	return m.recorder
}

func (m *MockApigatewayClient) GetAccount(arg0 context.Context, arg1 *apigateway.GetAccountInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccount, varargs...)
	ret0, _ := ret[0].(*apigateway.GetAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccount, reflect.TypeOf((*MockApigatewayClient)(nil).GetAccount), varargs...)
}

func (m *MockApigatewayClient) GetApiKey(arg0 context.Context, arg1 *apigateway.GetApiKeyInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetApiKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApiKey, varargs...)
	ret0, _ := ret[0].(*apigateway.GetApiKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetApiKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApiKey, reflect.TypeOf((*MockApigatewayClient)(nil).GetApiKey), varargs...)
}

func (m *MockApigatewayClient) GetApiKeys(arg0 context.Context, arg1 *apigateway.GetApiKeysInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApiKeys, varargs...)
	ret0, _ := ret[0].(*apigateway.GetApiKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetApiKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApiKeys, reflect.TypeOf((*MockApigatewayClient)(nil).GetApiKeys), varargs...)
}

func (m *MockApigatewayClient) GetAuthorizer(arg0 context.Context, arg1 *apigateway.GetAuthorizerInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetAuthorizerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAuthorizer, varargs...)
	ret0, _ := ret[0].(*apigateway.GetAuthorizerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetAuthorizer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAuthorizer, reflect.TypeOf((*MockApigatewayClient)(nil).GetAuthorizer), varargs...)
}

func (m *MockApigatewayClient) GetAuthorizers(arg0 context.Context, arg1 *apigateway.GetAuthorizersInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAuthorizers, varargs...)
	ret0, _ := ret[0].(*apigateway.GetAuthorizersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetAuthorizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAuthorizers, reflect.TypeOf((*MockApigatewayClient)(nil).GetAuthorizers), varargs...)
}

func (m *MockApigatewayClient) GetBasePathMapping(arg0 context.Context, arg1 *apigateway.GetBasePathMappingInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBasePathMapping, varargs...)
	ret0, _ := ret[0].(*apigateway.GetBasePathMappingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetBasePathMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBasePathMapping, reflect.TypeOf((*MockApigatewayClient)(nil).GetBasePathMapping), varargs...)
}

func (m *MockApigatewayClient) GetBasePathMappings(arg0 context.Context, arg1 *apigateway.GetBasePathMappingsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBasePathMappings, varargs...)
	ret0, _ := ret[0].(*apigateway.GetBasePathMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetBasePathMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBasePathMappings, reflect.TypeOf((*MockApigatewayClient)(nil).GetBasePathMappings), varargs...)
}

func (m *MockApigatewayClient) GetClientCertificate(arg0 context.Context, arg1 *apigateway.GetClientCertificateInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetClientCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClientCertificate, varargs...)
	ret0, _ := ret[0].(*apigateway.GetClientCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetClientCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClientCertificate, reflect.TypeOf((*MockApigatewayClient)(nil).GetClientCertificate), varargs...)
}

func (m *MockApigatewayClient) GetClientCertificates(arg0 context.Context, arg1 *apigateway.GetClientCertificatesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetClientCertificates, varargs...)
	ret0, _ := ret[0].(*apigateway.GetClientCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetClientCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetClientCertificates, reflect.TypeOf((*MockApigatewayClient)(nil).GetClientCertificates), varargs...)
}

func (m *MockApigatewayClient) GetDeployment(arg0 context.Context, arg1 *apigateway.GetDeploymentInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDeploymentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeployment, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDeploymentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeployment, reflect.TypeOf((*MockApigatewayClient)(nil).GetDeployment), varargs...)
}

func (m *MockApigatewayClient) GetDeployments(arg0 context.Context, arg1 *apigateway.GetDeploymentsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeployments, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeployments, reflect.TypeOf((*MockApigatewayClient)(nil).GetDeployments), varargs...)
}

func (m *MockApigatewayClient) GetDocumentationPart(arg0 context.Context, arg1 *apigateway.GetDocumentationPartInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDocumentationPart, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDocumentationPartOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDocumentationPart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDocumentationPart, reflect.TypeOf((*MockApigatewayClient)(nil).GetDocumentationPart), varargs...)
}

func (m *MockApigatewayClient) GetDocumentationParts(arg0 context.Context, arg1 *apigateway.GetDocumentationPartsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDocumentationParts, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDocumentationPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDocumentationParts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDocumentationParts, reflect.TypeOf((*MockApigatewayClient)(nil).GetDocumentationParts), varargs...)
}

func (m *MockApigatewayClient) GetDocumentationVersion(arg0 context.Context, arg1 *apigateway.GetDocumentationVersionInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDocumentationVersion, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDocumentationVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDocumentationVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDocumentationVersion, reflect.TypeOf((*MockApigatewayClient)(nil).GetDocumentationVersion), varargs...)
}

func (m *MockApigatewayClient) GetDocumentationVersions(arg0 context.Context, arg1 *apigateway.GetDocumentationVersionsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDocumentationVersions, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDocumentationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDocumentationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDocumentationVersions, reflect.TypeOf((*MockApigatewayClient)(nil).GetDocumentationVersions), varargs...)
}

func (m *MockApigatewayClient) GetDomainName(arg0 context.Context, arg1 *apigateway.GetDomainNameInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDomainNameOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainName, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDomainNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDomainName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainName, reflect.TypeOf((*MockApigatewayClient)(nil).GetDomainName), varargs...)
}

func (m *MockApigatewayClient) GetDomainNames(arg0 context.Context, arg1 *apigateway.GetDomainNamesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainNames, varargs...)
	ret0, _ := ret[0].(*apigateway.GetDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainNames, reflect.TypeOf((*MockApigatewayClient)(nil).GetDomainNames), varargs...)
}

func (m *MockApigatewayClient) GetExport(arg0 context.Context, arg1 *apigateway.GetExportInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetExportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetExport, varargs...)
	ret0, _ := ret[0].(*apigateway.GetExportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetExport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetExport, reflect.TypeOf((*MockApigatewayClient)(nil).GetExport), varargs...)
}

func (m *MockApigatewayClient) GetGatewayResponse(arg0 context.Context, arg1 *apigateway.GetGatewayResponseInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGatewayResponse, varargs...)
	ret0, _ := ret[0].(*apigateway.GetGatewayResponseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetGatewayResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGatewayResponse, reflect.TypeOf((*MockApigatewayClient)(nil).GetGatewayResponse), varargs...)
}

func (m *MockApigatewayClient) GetGatewayResponses(arg0 context.Context, arg1 *apigateway.GetGatewayResponsesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGatewayResponses, varargs...)
	ret0, _ := ret[0].(*apigateway.GetGatewayResponsesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetGatewayResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGatewayResponses, reflect.TypeOf((*MockApigatewayClient)(nil).GetGatewayResponses), varargs...)
}

func (m *MockApigatewayClient) GetIntegration(arg0 context.Context, arg1 *apigateway.GetIntegrationInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetIntegrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntegration, varargs...)
	ret0, _ := ret[0].(*apigateway.GetIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetIntegration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntegration, reflect.TypeOf((*MockApigatewayClient)(nil).GetIntegration), varargs...)
}

func (m *MockApigatewayClient) GetIntegrationResponse(arg0 context.Context, arg1 *apigateway.GetIntegrationResponseInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetIntegrationResponseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntegrationResponse, varargs...)
	ret0, _ := ret[0].(*apigateway.GetIntegrationResponseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetIntegrationResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntegrationResponse, reflect.TypeOf((*MockApigatewayClient)(nil).GetIntegrationResponse), varargs...)
}

func (m *MockApigatewayClient) GetMethod(arg0 context.Context, arg1 *apigateway.GetMethodInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetMethodOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMethod, varargs...)
	ret0, _ := ret[0].(*apigateway.GetMethodOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetMethod(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMethod, reflect.TypeOf((*MockApigatewayClient)(nil).GetMethod), varargs...)
}

func (m *MockApigatewayClient) GetMethodResponse(arg0 context.Context, arg1 *apigateway.GetMethodResponseInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetMethodResponseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMethodResponse, varargs...)
	ret0, _ := ret[0].(*apigateway.GetMethodResponseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetMethodResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMethodResponse, reflect.TypeOf((*MockApigatewayClient)(nil).GetMethodResponse), varargs...)
}

func (m *MockApigatewayClient) GetModel(arg0 context.Context, arg1 *apigateway.GetModelInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetModelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModel, varargs...)
	ret0, _ := ret[0].(*apigateway.GetModelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModel, reflect.TypeOf((*MockApigatewayClient)(nil).GetModel), varargs...)
}

func (m *MockApigatewayClient) GetModelTemplate(arg0 context.Context, arg1 *apigateway.GetModelTemplateInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModelTemplate, varargs...)
	ret0, _ := ret[0].(*apigateway.GetModelTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetModelTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModelTemplate, reflect.TypeOf((*MockApigatewayClient)(nil).GetModelTemplate), varargs...)
}

func (m *MockApigatewayClient) GetModels(arg0 context.Context, arg1 *apigateway.GetModelsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModels, varargs...)
	ret0, _ := ret[0].(*apigateway.GetModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModels, reflect.TypeOf((*MockApigatewayClient)(nil).GetModels), varargs...)
}

func (m *MockApigatewayClient) GetRequestValidator(arg0 context.Context, arg1 *apigateway.GetRequestValidatorInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRequestValidator, varargs...)
	ret0, _ := ret[0].(*apigateway.GetRequestValidatorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetRequestValidator(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRequestValidator, reflect.TypeOf((*MockApigatewayClient)(nil).GetRequestValidator), varargs...)
}

func (m *MockApigatewayClient) GetRequestValidators(arg0 context.Context, arg1 *apigateway.GetRequestValidatorsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRequestValidators, varargs...)
	ret0, _ := ret[0].(*apigateway.GetRequestValidatorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetRequestValidators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRequestValidators, reflect.TypeOf((*MockApigatewayClient)(nil).GetRequestValidators), varargs...)
}

func (m *MockApigatewayClient) GetResource(arg0 context.Context, arg1 *apigateway.GetResourceInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResource, varargs...)
	ret0, _ := ret[0].(*apigateway.GetResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResource, reflect.TypeOf((*MockApigatewayClient)(nil).GetResource), varargs...)
}

func (m *MockApigatewayClient) GetResources(arg0 context.Context, arg1 *apigateway.GetResourcesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResources, varargs...)
	ret0, _ := ret[0].(*apigateway.GetResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResources, reflect.TypeOf((*MockApigatewayClient)(nil).GetResources), varargs...)
}

func (m *MockApigatewayClient) GetRestApi(arg0 context.Context, arg1 *apigateway.GetRestApiInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetRestApiOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRestApi, varargs...)
	ret0, _ := ret[0].(*apigateway.GetRestApiOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetRestApi(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRestApi, reflect.TypeOf((*MockApigatewayClient)(nil).GetRestApi), varargs...)
}

func (m *MockApigatewayClient) GetRestApis(arg0 context.Context, arg1 *apigateway.GetRestApisInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRestApis, varargs...)
	ret0, _ := ret[0].(*apigateway.GetRestApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetRestApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRestApis, reflect.TypeOf((*MockApigatewayClient)(nil).GetRestApis), varargs...)
}

func (m *MockApigatewayClient) GetSdk(arg0 context.Context, arg1 *apigateway.GetSdkInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetSdkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSdk, varargs...)
	ret0, _ := ret[0].(*apigateway.GetSdkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetSdk(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSdk, reflect.TypeOf((*MockApigatewayClient)(nil).GetSdk), varargs...)
}

func (m *MockApigatewayClient) GetSdkType(arg0 context.Context, arg1 *apigateway.GetSdkTypeInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetSdkTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSdkType, varargs...)
	ret0, _ := ret[0].(*apigateway.GetSdkTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetSdkType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSdkType, reflect.TypeOf((*MockApigatewayClient)(nil).GetSdkType), varargs...)
}

func (m *MockApigatewayClient) GetSdkTypes(arg0 context.Context, arg1 *apigateway.GetSdkTypesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetSdkTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSdkTypes, varargs...)
	ret0, _ := ret[0].(*apigateway.GetSdkTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetSdkTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSdkTypes, reflect.TypeOf((*MockApigatewayClient)(nil).GetSdkTypes), varargs...)
}

func (m *MockApigatewayClient) GetStage(arg0 context.Context, arg1 *apigateway.GetStageInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetStageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStage, varargs...)
	ret0, _ := ret[0].(*apigateway.GetStageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetStage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStage, reflect.TypeOf((*MockApigatewayClient)(nil).GetStage), varargs...)
}

func (m *MockApigatewayClient) GetStages(arg0 context.Context, arg1 *apigateway.GetStagesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStages, varargs...)
	ret0, _ := ret[0].(*apigateway.GetStagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetStages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStages, reflect.TypeOf((*MockApigatewayClient)(nil).GetStages), varargs...)
}

func (m *MockApigatewayClient) GetTags(arg0 context.Context, arg1 *apigateway.GetTagsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTags, varargs...)
	ret0, _ := ret[0].(*apigateway.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTags, reflect.TypeOf((*MockApigatewayClient)(nil).GetTags), varargs...)
}

func (m *MockApigatewayClient) GetUsage(arg0 context.Context, arg1 *apigateway.GetUsageInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUsage, varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetUsage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUsage, reflect.TypeOf((*MockApigatewayClient)(nil).GetUsage), varargs...)
}

func (m *MockApigatewayClient) GetUsagePlan(arg0 context.Context, arg1 *apigateway.GetUsagePlanInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUsagePlan, varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsagePlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetUsagePlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUsagePlan, reflect.TypeOf((*MockApigatewayClient)(nil).GetUsagePlan), varargs...)
}

func (m *MockApigatewayClient) GetUsagePlanKey(arg0 context.Context, arg1 *apigateway.GetUsagePlanKeyInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUsagePlanKey, varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsagePlanKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetUsagePlanKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUsagePlanKey, reflect.TypeOf((*MockApigatewayClient)(nil).GetUsagePlanKey), varargs...)
}

func (m *MockApigatewayClient) GetUsagePlanKeys(arg0 context.Context, arg1 *apigateway.GetUsagePlanKeysInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUsagePlanKeys, varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsagePlanKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetUsagePlanKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUsagePlanKeys, reflect.TypeOf((*MockApigatewayClient)(nil).GetUsagePlanKeys), varargs...)
}

func (m *MockApigatewayClient) GetUsagePlans(arg0 context.Context, arg1 *apigateway.GetUsagePlansInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetUsagePlans, varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsagePlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetUsagePlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetUsagePlans, reflect.TypeOf((*MockApigatewayClient)(nil).GetUsagePlans), varargs...)
}

func (m *MockApigatewayClient) GetVpcLink(arg0 context.Context, arg1 *apigateway.GetVpcLinkInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetVpcLinkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVpcLink, varargs...)
	ret0, _ := ret[0].(*apigateway.GetVpcLinkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetVpcLink(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVpcLink, reflect.TypeOf((*MockApigatewayClient)(nil).GetVpcLink), varargs...)
}

func (m *MockApigatewayClient) GetVpcLinks(arg0 context.Context, arg1 *apigateway.GetVpcLinksInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVpcLinks, varargs...)
	ret0, _ := ret[0].(*apigateway.GetVpcLinksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayClientMockRecorder) GetVpcLinks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVpcLinks, reflect.TypeOf((*MockApigatewayClient)(nil).GetVpcLinks), varargs...)
}
