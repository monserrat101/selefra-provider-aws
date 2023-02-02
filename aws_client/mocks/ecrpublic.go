package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ecrpublic "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	gomock "github.com/golang/mock/gomock"
)

type MockEcrpublicClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEcrpublicClientMockRecorder
}

type MockEcrpublicClientMockRecorder struct {
	mock *MockEcrpublicClient
}

func NewMockEcrpublicClient(ctrl *gomock.Controller) *MockEcrpublicClient {
	mock := &MockEcrpublicClient{ctrl: ctrl}
	mock.recorder = &MockEcrpublicClientMockRecorder{mock}
	return mock
}

func (m *MockEcrpublicClient) EXPECT() *MockEcrpublicClientMockRecorder {
	return m.recorder
}

func (m *MockEcrpublicClient) DescribeImageTags(arg0 context.Context, arg1 *ecrpublic.DescribeImageTagsInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImageTags, varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeImageTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) DescribeImageTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImageTags, reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeImageTags), varargs...)
}

func (m *MockEcrpublicClient) DescribeImages(arg0 context.Context, arg1 *ecrpublic.DescribeImagesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImages, varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImages, reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeImages), varargs...)
}

func (m *MockEcrpublicClient) DescribeRegistries(arg0 context.Context, arg1 *ecrpublic.DescribeRegistriesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRegistries, varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeRegistriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) DescribeRegistries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRegistries, reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeRegistries), varargs...)
}

func (m *MockEcrpublicClient) DescribeRepositories(arg0 context.Context, arg1 *ecrpublic.DescribeRepositoriesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRepositories, varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) DescribeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRepositories, reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeRepositories), varargs...)
}

func (m *MockEcrpublicClient) GetAuthorizationToken(arg0 context.Context, arg1 *ecrpublic.GetAuthorizationTokenInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAuthorizationToken, varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) GetAuthorizationToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAuthorizationToken, reflect.TypeOf((*MockEcrpublicClient)(nil).GetAuthorizationToken), varargs...)
}

func (m *MockEcrpublicClient) GetRegistryCatalogData(arg0 context.Context, arg1 *ecrpublic.GetRegistryCatalogDataInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRegistryCatalogDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegistryCatalogData, varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRegistryCatalogDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) GetRegistryCatalogData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegistryCatalogData, reflect.TypeOf((*MockEcrpublicClient)(nil).GetRegistryCatalogData), varargs...)
}

func (m *MockEcrpublicClient) GetRepositoryCatalogData(arg0 context.Context, arg1 *ecrpublic.GetRepositoryCatalogDataInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryCatalogDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRepositoryCatalogData, varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRepositoryCatalogDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) GetRepositoryCatalogData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRepositoryCatalogData, reflect.TypeOf((*MockEcrpublicClient)(nil).GetRepositoryCatalogData), varargs...)
}

func (m *MockEcrpublicClient) GetRepositoryPolicy(arg0 context.Context, arg1 *ecrpublic.GetRepositoryPolicyInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRepositoryPolicy, varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) GetRepositoryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRepositoryPolicy, reflect.TypeOf((*MockEcrpublicClient)(nil).GetRepositoryPolicy), varargs...)
}

func (m *MockEcrpublicClient) ListTagsForResource(arg0 context.Context, arg1 *ecrpublic.ListTagsForResourceInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*ecrpublic.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrpublicClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockEcrpublicClient)(nil).ListTagsForResource), varargs...)
}
