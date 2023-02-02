package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	apigatewayv2 "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	gomock "github.com/golang/mock/gomock"
)

type MockApigatewayv2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockApigatewayv2ClientMockRecorder
}

type MockApigatewayv2ClientMockRecorder struct {
	mock *MockApigatewayv2Client
}

func NewMockApigatewayv2Client(ctrl *gomock.Controller) *MockApigatewayv2Client {
	mock := &MockApigatewayv2Client{ctrl: ctrl}
	mock.recorder = &MockApigatewayv2ClientMockRecorder{mock}
	return mock
}

func (m *MockApigatewayv2Client) EXPECT() *MockApigatewayv2ClientMockRecorder {
	return m.recorder
}

func (m *MockApigatewayv2Client) GetApi(arg0 context.Context, arg1 *apigatewayv2.GetApiInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApi, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetApiOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetApi(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApi, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetApi), varargs...)
}

func (m *MockApigatewayv2Client) GetApiMapping(arg0 context.Context, arg1 *apigatewayv2.GetApiMappingInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApiMapping, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetApiMappingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetApiMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApiMapping, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetApiMapping), varargs...)
}

func (m *MockApigatewayv2Client) GetApiMappings(arg0 context.Context, arg1 *apigatewayv2.GetApiMappingsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApiMappings, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetApiMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetApiMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApiMappings, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetApiMappings), varargs...)
}

func (m *MockApigatewayv2Client) GetApis(arg0 context.Context, arg1 *apigatewayv2.GetApisInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApis, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApis, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetApis), varargs...)
}

func (m *MockApigatewayv2Client) GetAuthorizer(arg0 context.Context, arg1 *apigatewayv2.GetAuthorizerInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAuthorizer, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetAuthorizerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetAuthorizer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAuthorizer, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetAuthorizer), varargs...)
}

func (m *MockApigatewayv2Client) GetAuthorizers(arg0 context.Context, arg1 *apigatewayv2.GetAuthorizersInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAuthorizers, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetAuthorizersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetAuthorizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAuthorizers, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetAuthorizers), varargs...)
}

func (m *MockApigatewayv2Client) GetDeployment(arg0 context.Context, arg1 *apigatewayv2.GetDeploymentInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeployment, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetDeploymentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeployment, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetDeployment), varargs...)
}

func (m *MockApigatewayv2Client) GetDeployments(arg0 context.Context, arg1 *apigatewayv2.GetDeploymentsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeployments, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeployments, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetDeployments), varargs...)
}

func (m *MockApigatewayv2Client) GetDomainName(arg0 context.Context, arg1 *apigatewayv2.GetDomainNameInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNameOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainName, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetDomainNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetDomainName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainName, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetDomainName), varargs...)
}

func (m *MockApigatewayv2Client) GetDomainNames(arg0 context.Context, arg1 *apigatewayv2.GetDomainNamesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainNames, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainNames, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetDomainNames), varargs...)
}

func (m *MockApigatewayv2Client) GetIntegration(arg0 context.Context, arg1 *apigatewayv2.GetIntegrationInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntegration, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetIntegration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntegration, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetIntegration), varargs...)
}

func (m *MockApigatewayv2Client) GetIntegrationResponse(arg0 context.Context, arg1 *apigatewayv2.GetIntegrationResponseInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntegrationResponse, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetIntegrationResponseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetIntegrationResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntegrationResponse, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetIntegrationResponse), varargs...)
}

func (m *MockApigatewayv2Client) GetIntegrationResponses(arg0 context.Context, arg1 *apigatewayv2.GetIntegrationResponsesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntegrationResponses, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetIntegrationResponsesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetIntegrationResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntegrationResponses, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetIntegrationResponses), varargs...)
}

func (m *MockApigatewayv2Client) GetIntegrations(arg0 context.Context, arg1 *apigatewayv2.GetIntegrationsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntegrations, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetIntegrationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetIntegrations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntegrations, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetIntegrations), varargs...)
}

func (m *MockApigatewayv2Client) GetModel(arg0 context.Context, arg1 *apigatewayv2.GetModelInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModel, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetModelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModel, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetModel), varargs...)
}

func (m *MockApigatewayv2Client) GetModelTemplate(arg0 context.Context, arg1 *apigatewayv2.GetModelTemplateInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModelTemplate, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetModelTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetModelTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModelTemplate, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetModelTemplate), varargs...)
}

func (m *MockApigatewayv2Client) GetModels(arg0 context.Context, arg1 *apigatewayv2.GetModelsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModels, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModels, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetModels), varargs...)
}

func (m *MockApigatewayv2Client) GetRoute(arg0 context.Context, arg1 *apigatewayv2.GetRouteInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRoute, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetRouteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetRoute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRoute, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetRoute), varargs...)
}

func (m *MockApigatewayv2Client) GetRouteResponse(arg0 context.Context, arg1 *apigatewayv2.GetRouteResponseInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRouteResponse, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetRouteResponseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetRouteResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRouteResponse, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetRouteResponse), varargs...)
}

func (m *MockApigatewayv2Client) GetRouteResponses(arg0 context.Context, arg1 *apigatewayv2.GetRouteResponsesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponsesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRouteResponses, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetRouteResponsesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetRouteResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRouteResponses, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetRouteResponses), varargs...)
}

func (m *MockApigatewayv2Client) GetRoutes(arg0 context.Context, arg1 *apigatewayv2.GetRoutesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRoutes, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRoutes, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetRoutes), varargs...)
}

func (m *MockApigatewayv2Client) GetStage(arg0 context.Context, arg1 *apigatewayv2.GetStageInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStage, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetStageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetStage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStage, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetStage), varargs...)
}

func (m *MockApigatewayv2Client) GetStages(arg0 context.Context, arg1 *apigatewayv2.GetStagesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStages, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetStagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetStages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStages, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetStages), varargs...)
}

func (m *MockApigatewayv2Client) GetTags(arg0 context.Context, arg1 *apigatewayv2.GetTagsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTags, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTags, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetTags), varargs...)
}

func (m *MockApigatewayv2Client) GetVpcLink(arg0 context.Context, arg1 *apigatewayv2.GetVpcLinkInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVpcLink, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetVpcLinkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetVpcLink(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVpcLink, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetVpcLink), varargs...)
}

func (m *MockApigatewayv2Client) GetVpcLinks(arg0 context.Context, arg1 *apigatewayv2.GetVpcLinksInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVpcLinks, varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetVpcLinksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetVpcLinks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVpcLinks, reflect.TypeOf((*MockApigatewayv2Client)(nil).GetVpcLinks), varargs...)
}
