package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	servicecatalogappregistry "github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	gomock "github.com/golang/mock/gomock"
)

type MockServicecatalogappregistryClient struct {
	ctrl		*gomock.Controller
	recorder	*MockServicecatalogappregistryClientMockRecorder
}

type MockServicecatalogappregistryClientMockRecorder struct {
	mock *MockServicecatalogappregistryClient
}

func NewMockServicecatalogappregistryClient(ctrl *gomock.Controller) *MockServicecatalogappregistryClient {
	mock := &MockServicecatalogappregistryClient{ctrl: ctrl}
	mock.recorder = &MockServicecatalogappregistryClientMockRecorder{mock}
	return mock
}

func (m *MockServicecatalogappregistryClient) EXPECT() *MockServicecatalogappregistryClientMockRecorder {
	return m.recorder
}

func (m *MockServicecatalogappregistryClient) GetApplication(arg0 context.Context, arg1 *servicecatalogappregistry.GetApplicationInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetApplication, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) GetApplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetApplication, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetApplication), varargs...)
}

func (m *MockServicecatalogappregistryClient) GetAssociatedResource(arg0 context.Context, arg1 *servicecatalogappregistry.GetAssociatedResourceInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAssociatedResource, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetAssociatedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) GetAssociatedResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAssociatedResource, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetAssociatedResource), varargs...)
}

func (m *MockServicecatalogappregistryClient) GetAttributeGroup(arg0 context.Context, arg1 *servicecatalogappregistry.GetAttributeGroupInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAttributeGroup, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetAttributeGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) GetAttributeGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAttributeGroup, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetAttributeGroup), varargs...)
}

func (m *MockServicecatalogappregistryClient) GetConfiguration(arg0 context.Context, arg1 *servicecatalogappregistry.GetConfigurationInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConfiguration, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) GetConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConfiguration, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetConfiguration), varargs...)
}

func (m *MockServicecatalogappregistryClient) ListApplications(arg0 context.Context, arg1 *servicecatalogappregistry.ListApplicationsInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListApplications, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) ListApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListApplications, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListApplications), varargs...)
}

func (m *MockServicecatalogappregistryClient) ListAssociatedAttributeGroups(arg0 context.Context, arg1 *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociatedAttributeGroups, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) ListAssociatedAttributeGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociatedAttributeGroups, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAssociatedAttributeGroups), varargs...)
}

func (m *MockServicecatalogappregistryClient) ListAssociatedResources(arg0 context.Context, arg1 *servicecatalogappregistry.ListAssociatedResourcesInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociatedResources, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAssociatedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) ListAssociatedResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociatedResources, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAssociatedResources), varargs...)
}

func (m *MockServicecatalogappregistryClient) ListAttributeGroups(arg0 context.Context, arg1 *servicecatalogappregistry.ListAttributeGroupsInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttributeGroups, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAttributeGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) ListAttributeGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttributeGroups, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAttributeGroups), varargs...)
}

func (m *MockServicecatalogappregistryClient) ListAttributeGroupsForApplication(arg0 context.Context, arg1 *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttributeGroupsForApplication, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) ListAttributeGroupsForApplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttributeGroupsForApplication, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAttributeGroupsForApplication), varargs...)
}

func (m *MockServicecatalogappregistryClient) ListTagsForResource(arg0 context.Context, arg1 *servicecatalogappregistry.ListTagsForResourceInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogappregistryClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListTagsForResource), varargs...)
}
