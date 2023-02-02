package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cloudtrail "github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudtrailClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudtrailClientMockRecorder
}

type MockCloudtrailClientMockRecorder struct {
	mock *MockCloudtrailClient
}

func NewMockCloudtrailClient(ctrl *gomock.Controller) *MockCloudtrailClient {
	mock := &MockCloudtrailClient{ctrl: ctrl}
	mock.recorder = &MockCloudtrailClientMockRecorder{mock}
	return mock
}

func (m *MockCloudtrailClient) EXPECT() *MockCloudtrailClientMockRecorder {
	return m.recorder
}

func (m *MockCloudtrailClient) DescribeQuery(arg0 context.Context, arg1 *cloudtrail.DescribeQueryInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.DescribeQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeQuery, varargs...)
	ret0, _ := ret[0].(*cloudtrail.DescribeQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) DescribeQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeQuery, reflect.TypeOf((*MockCloudtrailClient)(nil).DescribeQuery), varargs...)
}

func (m *MockCloudtrailClient) DescribeTrails(arg0 context.Context, arg1 *cloudtrail.DescribeTrailsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrails, varargs...)
	ret0, _ := ret[0].(*cloudtrail.DescribeTrailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) DescribeTrails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrails, reflect.TypeOf((*MockCloudtrailClient)(nil).DescribeTrails), varargs...)
}

func (m *MockCloudtrailClient) GetChannel(arg0 context.Context, arg1 *cloudtrail.GetChannelInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetChannelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetChannel, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetChannelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetChannel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetChannel, reflect.TypeOf((*MockCloudtrailClient)(nil).GetChannel), varargs...)
}

func (m *MockCloudtrailClient) GetEventDataStore(arg0 context.Context, arg1 *cloudtrail.GetEventDataStoreInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetEventDataStoreOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEventDataStore, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetEventDataStoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetEventDataStore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEventDataStore, reflect.TypeOf((*MockCloudtrailClient)(nil).GetEventDataStore), varargs...)
}

func (m *MockCloudtrailClient) GetEventSelectors(arg0 context.Context, arg1 *cloudtrail.GetEventSelectorsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEventSelectors, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetEventSelectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetEventSelectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEventSelectors, reflect.TypeOf((*MockCloudtrailClient)(nil).GetEventSelectors), varargs...)
}

func (m *MockCloudtrailClient) GetImport(arg0 context.Context, arg1 *cloudtrail.GetImportInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetImportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetImport, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetImport, reflect.TypeOf((*MockCloudtrailClient)(nil).GetImport), varargs...)
}

func (m *MockCloudtrailClient) GetInsightSelectors(arg0 context.Context, arg1 *cloudtrail.GetInsightSelectorsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetInsightSelectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInsightSelectors, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetInsightSelectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetInsightSelectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInsightSelectors, reflect.TypeOf((*MockCloudtrailClient)(nil).GetInsightSelectors), varargs...)
}

func (m *MockCloudtrailClient) GetQueryResults(arg0 context.Context, arg1 *cloudtrail.GetQueryResultsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetQueryResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetQueryResults, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetQueryResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetQueryResults, reflect.TypeOf((*MockCloudtrailClient)(nil).GetQueryResults), varargs...)
}

func (m *MockCloudtrailClient) GetTrail(arg0 context.Context, arg1 *cloudtrail.GetTrailInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTrail, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetTrailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetTrail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTrail, reflect.TypeOf((*MockCloudtrailClient)(nil).GetTrail), varargs...)
}

func (m *MockCloudtrailClient) GetTrailStatus(arg0 context.Context, arg1 *cloudtrail.GetTrailStatusInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTrailStatus, varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetTrailStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetTrailStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTrailStatus, reflect.TypeOf((*MockCloudtrailClient)(nil).GetTrailStatus), varargs...)
}

func (m *MockCloudtrailClient) ListChannels(arg0 context.Context, arg1 *cloudtrail.ListChannelsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListChannelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListChannels, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListChannelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListChannels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListChannels, reflect.TypeOf((*MockCloudtrailClient)(nil).ListChannels), varargs...)
}

func (m *MockCloudtrailClient) ListEventDataStores(arg0 context.Context, arg1 *cloudtrail.ListEventDataStoresInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListEventDataStoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEventDataStores, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListEventDataStoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListEventDataStores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEventDataStores, reflect.TypeOf((*MockCloudtrailClient)(nil).ListEventDataStores), varargs...)
}

func (m *MockCloudtrailClient) ListImportFailures(arg0 context.Context, arg1 *cloudtrail.ListImportFailuresInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListImportFailuresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImportFailures, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListImportFailuresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListImportFailures(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImportFailures, reflect.TypeOf((*MockCloudtrailClient)(nil).ListImportFailures), varargs...)
}

func (m *MockCloudtrailClient) ListImports(arg0 context.Context, arg1 *cloudtrail.ListImportsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListImportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImports, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListImportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListImports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImports, reflect.TypeOf((*MockCloudtrailClient)(nil).ListImports), varargs...)
}

func (m *MockCloudtrailClient) ListPublicKeys(arg0 context.Context, arg1 *cloudtrail.ListPublicKeysInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListPublicKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPublicKeys, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListPublicKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListPublicKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPublicKeys, reflect.TypeOf((*MockCloudtrailClient)(nil).ListPublicKeys), varargs...)
}

func (m *MockCloudtrailClient) ListQueries(arg0 context.Context, arg1 *cloudtrail.ListQueriesInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListQueries, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListQueries, reflect.TypeOf((*MockCloudtrailClient)(nil).ListQueries), varargs...)
}

func (m *MockCloudtrailClient) ListTags(arg0 context.Context, arg1 *cloudtrail.ListTagsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockCloudtrailClient)(nil).ListTags), varargs...)
}

func (m *MockCloudtrailClient) ListTrails(arg0 context.Context, arg1 *cloudtrail.ListTrailsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListTrailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrails, varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListTrailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListTrails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrails, reflect.TypeOf((*MockCloudtrailClient)(nil).ListTrails), varargs...)
}
