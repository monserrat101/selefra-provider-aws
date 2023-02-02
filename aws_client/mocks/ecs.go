package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

type MockEcsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEcsClientMockRecorder
}

type MockEcsClientMockRecorder struct {
	mock *MockEcsClient
}

func NewMockEcsClient(ctrl *gomock.Controller) *MockEcsClient {
	mock := &MockEcsClient{ctrl: ctrl}
	mock.recorder = &MockEcsClientMockRecorder{mock}
	return mock
}

func (m *MockEcsClient) EXPECT() *MockEcsClientMockRecorder {
	return m.recorder
}

func (m *MockEcsClient) DescribeCapacityProviders(arg0 context.Context, arg1 *ecs.DescribeCapacityProvidersInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeCapacityProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCapacityProviders, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeCapacityProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeCapacityProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCapacityProviders, reflect.TypeOf((*MockEcsClient)(nil).DescribeCapacityProviders), varargs...)
}

func (m *MockEcsClient) DescribeClusters(arg0 context.Context, arg1 *ecs.DescribeClustersInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusters, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusters, reflect.TypeOf((*MockEcsClient)(nil).DescribeClusters), varargs...)
}

func (m *MockEcsClient) DescribeContainerInstances(arg0 context.Context, arg1 *ecs.DescribeContainerInstancesInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeContainerInstances, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeContainerInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeContainerInstances, reflect.TypeOf((*MockEcsClient)(nil).DescribeContainerInstances), varargs...)
}

func (m *MockEcsClient) DescribeServices(arg0 context.Context, arg1 *ecs.DescribeServicesInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeServices, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeServices, reflect.TypeOf((*MockEcsClient)(nil).DescribeServices), varargs...)
}

func (m *MockEcsClient) DescribeTaskDefinition(arg0 context.Context, arg1 *ecs.DescribeTaskDefinitionInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTaskDefinition, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeTaskDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTaskDefinition, reflect.TypeOf((*MockEcsClient)(nil).DescribeTaskDefinition), varargs...)
}

func (m *MockEcsClient) DescribeTaskSets(arg0 context.Context, arg1 *ecs.DescribeTaskSetsInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTaskSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTaskSets, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTaskSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeTaskSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTaskSets, reflect.TypeOf((*MockEcsClient)(nil).DescribeTaskSets), varargs...)
}

func (m *MockEcsClient) DescribeTasks(arg0 context.Context, arg1 *ecs.DescribeTasksInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTasks, varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTasks, reflect.TypeOf((*MockEcsClient)(nil).DescribeTasks), varargs...)
}

func (m *MockEcsClient) GetTaskProtection(arg0 context.Context, arg1 *ecs.GetTaskProtectionInput, arg2 ...func(*ecs.Options)) (*ecs.GetTaskProtectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTaskProtection, varargs...)
	ret0, _ := ret[0].(*ecs.GetTaskProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) GetTaskProtection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTaskProtection, reflect.TypeOf((*MockEcsClient)(nil).GetTaskProtection), varargs...)
}

func (m *MockEcsClient) ListAccountSettings(arg0 context.Context, arg1 *ecs.ListAccountSettingsInput, arg2 ...func(*ecs.Options)) (*ecs.ListAccountSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccountSettings, varargs...)
	ret0, _ := ret[0].(*ecs.ListAccountSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListAccountSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccountSettings, reflect.TypeOf((*MockEcsClient)(nil).ListAccountSettings), varargs...)
}

func (m *MockEcsClient) ListAttributes(arg0 context.Context, arg1 *ecs.ListAttributesInput, arg2 ...func(*ecs.Options)) (*ecs.ListAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAttributes, varargs...)
	ret0, _ := ret[0].(*ecs.ListAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAttributes, reflect.TypeOf((*MockEcsClient)(nil).ListAttributes), varargs...)
}

func (m *MockEcsClient) ListClusters(arg0 context.Context, arg1 *ecs.ListClustersInput, arg2 ...func(*ecs.Options)) (*ecs.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListClusters, varargs...)
	ret0, _ := ret[0].(*ecs.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListClusters, reflect.TypeOf((*MockEcsClient)(nil).ListClusters), varargs...)
}

func (m *MockEcsClient) ListContainerInstances(arg0 context.Context, arg1 *ecs.ListContainerInstancesInput, arg2 ...func(*ecs.Options)) (*ecs.ListContainerInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListContainerInstances, varargs...)
	ret0, _ := ret[0].(*ecs.ListContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListContainerInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListContainerInstances, reflect.TypeOf((*MockEcsClient)(nil).ListContainerInstances), varargs...)
}

func (m *MockEcsClient) ListServices(arg0 context.Context, arg1 *ecs.ListServicesInput, arg2 ...func(*ecs.Options)) (*ecs.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServices, varargs...)
	ret0, _ := ret[0].(*ecs.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServices, reflect.TypeOf((*MockEcsClient)(nil).ListServices), varargs...)
}

func (m *MockEcsClient) ListServicesByNamespace(arg0 context.Context, arg1 *ecs.ListServicesByNamespaceInput, arg2 ...func(*ecs.Options)) (*ecs.ListServicesByNamespaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServicesByNamespace, varargs...)
	ret0, _ := ret[0].(*ecs.ListServicesByNamespaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListServicesByNamespace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServicesByNamespace, reflect.TypeOf((*MockEcsClient)(nil).ListServicesByNamespace), varargs...)
}

func (m *MockEcsClient) ListTagsForResource(arg0 context.Context, arg1 *ecs.ListTagsForResourceInput, arg2 ...func(*ecs.Options)) (*ecs.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*ecs.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockEcsClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockEcsClient) ListTaskDefinitionFamilies(arg0 context.Context, arg1 *ecs.ListTaskDefinitionFamiliesInput, arg2 ...func(*ecs.Options)) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTaskDefinitionFamilies, varargs...)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionFamiliesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListTaskDefinitionFamilies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTaskDefinitionFamilies, reflect.TypeOf((*MockEcsClient)(nil).ListTaskDefinitionFamilies), varargs...)
}

func (m *MockEcsClient) ListTaskDefinitions(arg0 context.Context, arg1 *ecs.ListTaskDefinitionsInput, arg2 ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTaskDefinitions, varargs...)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListTaskDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTaskDefinitions, reflect.TypeOf((*MockEcsClient)(nil).ListTaskDefinitions), varargs...)
}

func (m *MockEcsClient) ListTasks(arg0 context.Context, arg1 *ecs.ListTasksInput, arg2 ...func(*ecs.Options)) (*ecs.ListTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTasks, varargs...)
	ret0, _ := ret[0].(*ecs.ListTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTasks, reflect.TypeOf((*MockEcsClient)(nil).ListTasks), varargs...)
}
