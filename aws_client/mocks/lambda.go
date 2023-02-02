package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	lambda "github.com/aws/aws-sdk-go-v2/service/lambda"
	gomock "github.com/golang/mock/gomock"
)

type MockLambdaClient struct {
	ctrl		*gomock.Controller
	recorder	*MockLambdaClientMockRecorder
}

type MockLambdaClientMockRecorder struct {
	mock *MockLambdaClient
}

func NewMockLambdaClient(ctrl *gomock.Controller) *MockLambdaClient {
	mock := &MockLambdaClient{ctrl: ctrl}
	mock.recorder = &MockLambdaClientMockRecorder{mock}
	return mock
}

func (m *MockLambdaClient) EXPECT() *MockLambdaClientMockRecorder {
	return m.recorder
}

func (m *MockLambdaClient) GetAccountSettings(arg0 context.Context, arg1 *lambda.GetAccountSettingsInput, arg2 ...func(*lambda.Options)) (*lambda.GetAccountSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccountSettings, varargs...)
	ret0, _ := ret[0].(*lambda.GetAccountSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetAccountSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccountSettings, reflect.TypeOf((*MockLambdaClient)(nil).GetAccountSettings), varargs...)
}

func (m *MockLambdaClient) GetAlias(arg0 context.Context, arg1 *lambda.GetAliasInput, arg2 ...func(*lambda.Options)) (*lambda.GetAliasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAlias, varargs...)
	ret0, _ := ret[0].(*lambda.GetAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetAlias(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAlias, reflect.TypeOf((*MockLambdaClient)(nil).GetAlias), varargs...)
}

func (m *MockLambdaClient) GetCodeSigningConfig(arg0 context.Context, arg1 *lambda.GetCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCodeSigningConfig, varargs...)
	ret0, _ := ret[0].(*lambda.GetCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCodeSigningConfig, reflect.TypeOf((*MockLambdaClient)(nil).GetCodeSigningConfig), varargs...)
}

func (m *MockLambdaClient) GetEventSourceMapping(arg0 context.Context, arg1 *lambda.GetEventSourceMappingInput, arg2 ...func(*lambda.Options)) (*lambda.GetEventSourceMappingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEventSourceMapping, varargs...)
	ret0, _ := ret[0].(*lambda.GetEventSourceMappingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetEventSourceMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEventSourceMapping, reflect.TypeOf((*MockLambdaClient)(nil).GetEventSourceMapping), varargs...)
}

func (m *MockLambdaClient) GetFunction(arg0 context.Context, arg1 *lambda.GetFunctionInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunction, varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunction, reflect.TypeOf((*MockLambdaClient)(nil).GetFunction), varargs...)
}

func (m *MockLambdaClient) GetFunctionCodeSigningConfig(arg0 context.Context, arg1 *lambda.GetFunctionCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunctionCodeSigningConfig, varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunctionCodeSigningConfig, reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionCodeSigningConfig), varargs...)
}

func (m *MockLambdaClient) GetFunctionConcurrency(arg0 context.Context, arg1 *lambda.GetFunctionConcurrencyInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionConcurrencyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunctionConcurrency, varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionConcurrencyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionConcurrency(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunctionConcurrency, reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionConcurrency), varargs...)
}

func (m *MockLambdaClient) GetFunctionConfiguration(arg0 context.Context, arg1 *lambda.GetFunctionConfigurationInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunctionConfiguration, varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunctionConfiguration, reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionConfiguration), varargs...)
}

func (m *MockLambdaClient) GetFunctionEventInvokeConfig(arg0 context.Context, arg1 *lambda.GetFunctionEventInvokeConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunctionEventInvokeConfig, varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionEventInvokeConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionEventInvokeConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunctionEventInvokeConfig, reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionEventInvokeConfig), varargs...)
}

func (m *MockLambdaClient) GetFunctionUrlConfig(arg0 context.Context, arg1 *lambda.GetFunctionUrlConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunctionUrlConfig, varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionUrlConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionUrlConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunctionUrlConfig, reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionUrlConfig), varargs...)
}

func (m *MockLambdaClient) GetLayerVersion(arg0 context.Context, arg1 *lambda.GetLayerVersionInput, arg2 ...func(*lambda.Options)) (*lambda.GetLayerVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLayerVersion, varargs...)
	ret0, _ := ret[0].(*lambda.GetLayerVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetLayerVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLayerVersion, reflect.TypeOf((*MockLambdaClient)(nil).GetLayerVersion), varargs...)
}

func (m *MockLambdaClient) GetLayerVersionByArn(arg0 context.Context, arg1 *lambda.GetLayerVersionByArnInput, arg2 ...func(*lambda.Options)) (*lambda.GetLayerVersionByArnOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLayerVersionByArn, varargs...)
	ret0, _ := ret[0].(*lambda.GetLayerVersionByArnOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetLayerVersionByArn(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLayerVersionByArn, reflect.TypeOf((*MockLambdaClient)(nil).GetLayerVersionByArn), varargs...)
}

func (m *MockLambdaClient) GetLayerVersionPolicy(arg0 context.Context, arg1 *lambda.GetLayerVersionPolicyInput, arg2 ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLayerVersionPolicy, varargs...)
	ret0, _ := ret[0].(*lambda.GetLayerVersionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetLayerVersionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLayerVersionPolicy, reflect.TypeOf((*MockLambdaClient)(nil).GetLayerVersionPolicy), varargs...)
}

func (m *MockLambdaClient) GetPolicy(arg0 context.Context, arg1 *lambda.GetPolicyInput, arg2 ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPolicy, varargs...)
	ret0, _ := ret[0].(*lambda.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPolicy, reflect.TypeOf((*MockLambdaClient)(nil).GetPolicy), varargs...)
}

func (m *MockLambdaClient) GetProvisionedConcurrencyConfig(arg0 context.Context, arg1 *lambda.GetProvisionedConcurrencyConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetProvisionedConcurrencyConfig, varargs...)
	ret0, _ := ret[0].(*lambda.GetProvisionedConcurrencyConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetProvisionedConcurrencyConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetProvisionedConcurrencyConfig, reflect.TypeOf((*MockLambdaClient)(nil).GetProvisionedConcurrencyConfig), varargs...)
}

func (m *MockLambdaClient) ListAliases(arg0 context.Context, arg1 *lambda.ListAliasesInput, arg2 ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAliases, varargs...)
	ret0, _ := ret[0].(*lambda.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAliases, reflect.TypeOf((*MockLambdaClient)(nil).ListAliases), varargs...)
}

func (m *MockLambdaClient) ListCodeSigningConfigs(arg0 context.Context, arg1 *lambda.ListCodeSigningConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListCodeSigningConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCodeSigningConfigs, varargs...)
	ret0, _ := ret[0].(*lambda.ListCodeSigningConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListCodeSigningConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCodeSigningConfigs, reflect.TypeOf((*MockLambdaClient)(nil).ListCodeSigningConfigs), varargs...)
}

func (m *MockLambdaClient) ListEventSourceMappings(arg0 context.Context, arg1 *lambda.ListEventSourceMappingsInput, arg2 ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEventSourceMappings, varargs...)
	ret0, _ := ret[0].(*lambda.ListEventSourceMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListEventSourceMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEventSourceMappings, reflect.TypeOf((*MockLambdaClient)(nil).ListEventSourceMappings), varargs...)
}

func (m *MockLambdaClient) ListFunctionEventInvokeConfigs(arg0 context.Context, arg1 *lambda.ListFunctionEventInvokeConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFunctionEventInvokeConfigs, varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionEventInvokeConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListFunctionEventInvokeConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFunctionEventInvokeConfigs, reflect.TypeOf((*MockLambdaClient)(nil).ListFunctionEventInvokeConfigs), varargs...)
}

func (m *MockLambdaClient) ListFunctionUrlConfigs(arg0 context.Context, arg1 *lambda.ListFunctionUrlConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionUrlConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFunctionUrlConfigs, varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionUrlConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListFunctionUrlConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFunctionUrlConfigs, reflect.TypeOf((*MockLambdaClient)(nil).ListFunctionUrlConfigs), varargs...)
}

func (m *MockLambdaClient) ListFunctions(arg0 context.Context, arg1 *lambda.ListFunctionsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFunctions, varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFunctions, reflect.TypeOf((*MockLambdaClient)(nil).ListFunctions), varargs...)
}

func (m *MockLambdaClient) ListFunctionsByCodeSigningConfig(arg0 context.Context, arg1 *lambda.ListFunctionsByCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionsByCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFunctionsByCodeSigningConfig, varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionsByCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListFunctionsByCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFunctionsByCodeSigningConfig, reflect.TypeOf((*MockLambdaClient)(nil).ListFunctionsByCodeSigningConfig), varargs...)
}

func (m *MockLambdaClient) ListLayerVersions(arg0 context.Context, arg1 *lambda.ListLayerVersionsInput, arg2 ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLayerVersions, varargs...)
	ret0, _ := ret[0].(*lambda.ListLayerVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListLayerVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLayerVersions, reflect.TypeOf((*MockLambdaClient)(nil).ListLayerVersions), varargs...)
}

func (m *MockLambdaClient) ListLayers(arg0 context.Context, arg1 *lambda.ListLayersInput, arg2 ...func(*lambda.Options)) (*lambda.ListLayersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLayers, varargs...)
	ret0, _ := ret[0].(*lambda.ListLayersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListLayers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLayers, reflect.TypeOf((*MockLambdaClient)(nil).ListLayers), varargs...)
}

func (m *MockLambdaClient) ListProvisionedConcurrencyConfigs(arg0 context.Context, arg1 *lambda.ListProvisionedConcurrencyConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisionedConcurrencyConfigs, varargs...)
	ret0, _ := ret[0].(*lambda.ListProvisionedConcurrencyConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListProvisionedConcurrencyConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisionedConcurrencyConfigs, reflect.TypeOf((*MockLambdaClient)(nil).ListProvisionedConcurrencyConfigs), varargs...)
}

func (m *MockLambdaClient) ListTags(arg0 context.Context, arg1 *lambda.ListTagsInput, arg2 ...func(*lambda.Options)) (*lambda.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*lambda.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockLambdaClient)(nil).ListTags), varargs...)
}

func (m *MockLambdaClient) ListVersionsByFunction(arg0 context.Context, arg1 *lambda.ListVersionsByFunctionInput, arg2 ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListVersionsByFunction, varargs...)
	ret0, _ := ret[0].(*lambda.ListVersionsByFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListVersionsByFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVersionsByFunction, reflect.TypeOf((*MockLambdaClient)(nil).ListVersionsByFunction), varargs...)
}
