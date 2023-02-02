package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	appsync "github.com/aws/aws-sdk-go-v2/service/appsync"
	gomock "github.com/golang/mock/gomock"
)

type MockAppsyncClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAppsyncClientMockRecorder
}

type MockAppsyncClientMockRecorder struct {
	mock *MockAppsyncClient
}

func NewMockAppsyncClient(ctrl *gomock.Controller) *MockAppsyncClient {
	mock := &MockAppsyncClient{ctrl: ctrl}
	mock.recorder = &MockAppsyncClientMockRecorder{mock}
	return mock
}

func (m *MockAppsyncClient) EXPECT() *MockAppsyncClientMockRecorder {
	return m.recorder
}

func (m *MockAppsyncClient) GetApiAssociation(arg0 context.Context, arg1 *appsync.GetApiAssociationInput, arg2 ...func(*appsync.Options)) (*appsync.GetApiAssociationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApiAssociation, varargs...)
	ret0, _ := ret[0].(*appsync.GetApiAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetApiAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApiAssociation, reflect.TypeOf((*MockAppsyncClient)(nil).GetApiAssociation), varargs...)
}

func (m *MockAppsyncClient) GetApiCache(arg0 context.Context, arg1 *appsync.GetApiCacheInput, arg2 ...func(*appsync.Options)) (*appsync.GetApiCacheOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApiCache, varargs...)
	ret0, _ := ret[0].(*appsync.GetApiCacheOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetApiCache(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApiCache, reflect.TypeOf((*MockAppsyncClient)(nil).GetApiCache), varargs...)
}

func (m *MockAppsyncClient) GetDataSource(arg0 context.Context, arg1 *appsync.GetDataSourceInput, arg2 ...func(*appsync.Options)) (*appsync.GetDataSourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataSource, varargs...)
	ret0, _ := ret[0].(*appsync.GetDataSourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetDataSource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataSource, reflect.TypeOf((*MockAppsyncClient)(nil).GetDataSource), varargs...)
}

func (m *MockAppsyncClient) GetDomainName(arg0 context.Context, arg1 *appsync.GetDomainNameInput, arg2 ...func(*appsync.Options)) (*appsync.GetDomainNameOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomainName, varargs...)
	ret0, _ := ret[0].(*appsync.GetDomainNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetDomainName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomainName, reflect.TypeOf((*MockAppsyncClient)(nil).GetDomainName), varargs...)
}

func (m *MockAppsyncClient) GetFunction(arg0 context.Context, arg1 *appsync.GetFunctionInput, arg2 ...func(*appsync.Options)) (*appsync.GetFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetFunction, varargs...)
	ret0, _ := ret[0].(*appsync.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFunction, reflect.TypeOf((*MockAppsyncClient)(nil).GetFunction), varargs...)
}

func (m *MockAppsyncClient) GetGraphqlApi(arg0 context.Context, arg1 *appsync.GetGraphqlApiInput, arg2 ...func(*appsync.Options)) (*appsync.GetGraphqlApiOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetGraphqlApi, varargs...)
	ret0, _ := ret[0].(*appsync.GetGraphqlApiOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetGraphqlApi(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetGraphqlApi, reflect.TypeOf((*MockAppsyncClient)(nil).GetGraphqlApi), varargs...)
}

func (m *MockAppsyncClient) GetIntrospectionSchema(arg0 context.Context, arg1 *appsync.GetIntrospectionSchemaInput, arg2 ...func(*appsync.Options)) (*appsync.GetIntrospectionSchemaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIntrospectionSchema, varargs...)
	ret0, _ := ret[0].(*appsync.GetIntrospectionSchemaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetIntrospectionSchema(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIntrospectionSchema, reflect.TypeOf((*MockAppsyncClient)(nil).GetIntrospectionSchema), varargs...)
}

func (m *MockAppsyncClient) GetResolver(arg0 context.Context, arg1 *appsync.GetResolverInput, arg2 ...func(*appsync.Options)) (*appsync.GetResolverOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResolver, varargs...)
	ret0, _ := ret[0].(*appsync.GetResolverOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetResolver(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResolver, reflect.TypeOf((*MockAppsyncClient)(nil).GetResolver), varargs...)
}

func (m *MockAppsyncClient) GetSchemaCreationStatus(arg0 context.Context, arg1 *appsync.GetSchemaCreationStatusInput, arg2 ...func(*appsync.Options)) (*appsync.GetSchemaCreationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSchemaCreationStatus, varargs...)
	ret0, _ := ret[0].(*appsync.GetSchemaCreationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetSchemaCreationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSchemaCreationStatus, reflect.TypeOf((*MockAppsyncClient)(nil).GetSchemaCreationStatus), varargs...)
}

func (m *MockAppsyncClient) GetType(arg0 context.Context, arg1 *appsync.GetTypeInput, arg2 ...func(*appsync.Options)) (*appsync.GetTypeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetType, varargs...)
	ret0, _ := ret[0].(*appsync.GetTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) GetType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetType, reflect.TypeOf((*MockAppsyncClient)(nil).GetType), varargs...)
}

func (m *MockAppsyncClient) ListApiKeys(arg0 context.Context, arg1 *appsync.ListApiKeysInput, arg2 ...func(*appsync.Options)) (*appsync.ListApiKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListApiKeys, varargs...)
	ret0, _ := ret[0].(*appsync.ListApiKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListApiKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListApiKeys, reflect.TypeOf((*MockAppsyncClient)(nil).ListApiKeys), varargs...)
}

func (m *MockAppsyncClient) ListDataSources(arg0 context.Context, arg1 *appsync.ListDataSourcesInput, arg2 ...func(*appsync.Options)) (*appsync.ListDataSourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataSources, varargs...)
	ret0, _ := ret[0].(*appsync.ListDataSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListDataSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataSources, reflect.TypeOf((*MockAppsyncClient)(nil).ListDataSources), varargs...)
}

func (m *MockAppsyncClient) ListDomainNames(arg0 context.Context, arg1 *appsync.ListDomainNamesInput, arg2 ...func(*appsync.Options)) (*appsync.ListDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomainNames, varargs...)
	ret0, _ := ret[0].(*appsync.ListDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomainNames, reflect.TypeOf((*MockAppsyncClient)(nil).ListDomainNames), varargs...)
}

func (m *MockAppsyncClient) ListFunctions(arg0 context.Context, arg1 *appsync.ListFunctionsInput, arg2 ...func(*appsync.Options)) (*appsync.ListFunctionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFunctions, varargs...)
	ret0, _ := ret[0].(*appsync.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFunctions, reflect.TypeOf((*MockAppsyncClient)(nil).ListFunctions), varargs...)
}

func (m *MockAppsyncClient) ListGraphqlApis(arg0 context.Context, arg1 *appsync.ListGraphqlApisInput, arg2 ...func(*appsync.Options)) (*appsync.ListGraphqlApisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListGraphqlApis, varargs...)
	ret0, _ := ret[0].(*appsync.ListGraphqlApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListGraphqlApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListGraphqlApis, reflect.TypeOf((*MockAppsyncClient)(nil).ListGraphqlApis), varargs...)
}

func (m *MockAppsyncClient) ListResolvers(arg0 context.Context, arg1 *appsync.ListResolversInput, arg2 ...func(*appsync.Options)) (*appsync.ListResolversOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResolvers, varargs...)
	ret0, _ := ret[0].(*appsync.ListResolversOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListResolvers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResolvers, reflect.TypeOf((*MockAppsyncClient)(nil).ListResolvers), varargs...)
}

func (m *MockAppsyncClient) ListResolversByFunction(arg0 context.Context, arg1 *appsync.ListResolversByFunctionInput, arg2 ...func(*appsync.Options)) (*appsync.ListResolversByFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResolversByFunction, varargs...)
	ret0, _ := ret[0].(*appsync.ListResolversByFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListResolversByFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResolversByFunction, reflect.TypeOf((*MockAppsyncClient)(nil).ListResolversByFunction), varargs...)
}

func (m *MockAppsyncClient) ListTagsForResource(arg0 context.Context, arg1 *appsync.ListTagsForResourceInput, arg2 ...func(*appsync.Options)) (*appsync.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*appsync.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockAppsyncClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockAppsyncClient) ListTypes(arg0 context.Context, arg1 *appsync.ListTypesInput, arg2 ...func(*appsync.Options)) (*appsync.ListTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTypes, varargs...)
	ret0, _ := ret[0].(*appsync.ListTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppsyncClientMockRecorder) ListTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTypes, reflect.TypeOf((*MockAppsyncClient)(nil).ListTypes), varargs...)
}
