package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ecr "github.com/aws/aws-sdk-go-v2/service/ecr"
	gomock "github.com/golang/mock/gomock"
)

type MockEcrClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEcrClientMockRecorder
}

type MockEcrClientMockRecorder struct {
	mock *MockEcrClient
}

func NewMockEcrClient(ctrl *gomock.Controller) *MockEcrClient {
	mock := &MockEcrClient{ctrl: ctrl}
	mock.recorder = &MockEcrClientMockRecorder{mock}
	return mock
}

func (m *MockEcrClient) EXPECT() *MockEcrClientMockRecorder {
	return m.recorder
}

func (m *MockEcrClient) BatchGetImage(arg0 context.Context, arg1 *ecr.BatchGetImageInput, arg2 ...func(*ecr.Options)) (*ecr.BatchGetImageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetImage, varargs...)
	ret0, _ := ret[0].(*ecr.BatchGetImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) BatchGetImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetImage, reflect.TypeOf((*MockEcrClient)(nil).BatchGetImage), varargs...)
}

func (m *MockEcrClient) BatchGetRepositoryScanningConfiguration(arg0 context.Context, arg1 *ecr.BatchGetRepositoryScanningConfigurationInput, arg2 ...func(*ecr.Options)) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetRepositoryScanningConfiguration, varargs...)
	ret0, _ := ret[0].(*ecr.BatchGetRepositoryScanningConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) BatchGetRepositoryScanningConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetRepositoryScanningConfiguration, reflect.TypeOf((*MockEcrClient)(nil).BatchGetRepositoryScanningConfiguration), varargs...)
}

func (m *MockEcrClient) DescribeImageReplicationStatus(arg0 context.Context, arg1 *ecr.DescribeImageReplicationStatusInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImageReplicationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImageReplicationStatus, varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImageReplicationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeImageReplicationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImageReplicationStatus, reflect.TypeOf((*MockEcrClient)(nil).DescribeImageReplicationStatus), varargs...)
}

func (m *MockEcrClient) DescribeImageScanFindings(arg0 context.Context, arg1 *ecr.DescribeImageScanFindingsInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImageScanFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImageScanFindings, varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImageScanFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeImageScanFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImageScanFindings, reflect.TypeOf((*MockEcrClient)(nil).DescribeImageScanFindings), varargs...)
}

func (m *MockEcrClient) DescribeImages(arg0 context.Context, arg1 *ecr.DescribeImagesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImages, varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImages, reflect.TypeOf((*MockEcrClient)(nil).DescribeImages), varargs...)
}

func (m *MockEcrClient) DescribePullThroughCacheRules(arg0 context.Context, arg1 *ecr.DescribePullThroughCacheRulesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribePullThroughCacheRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePullThroughCacheRules, varargs...)
	ret0, _ := ret[0].(*ecr.DescribePullThroughCacheRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribePullThroughCacheRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePullThroughCacheRules, reflect.TypeOf((*MockEcrClient)(nil).DescribePullThroughCacheRules), varargs...)
}

func (m *MockEcrClient) DescribeRegistry(arg0 context.Context, arg1 *ecr.DescribeRegistryInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeRegistryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRegistry, varargs...)
	ret0, _ := ret[0].(*ecr.DescribeRegistryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeRegistry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRegistry, reflect.TypeOf((*MockEcrClient)(nil).DescribeRegistry), varargs...)
}

func (m *MockEcrClient) DescribeRepositories(arg0 context.Context, arg1 *ecr.DescribeRepositoriesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRepositories, varargs...)
	ret0, _ := ret[0].(*ecr.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRepositories, reflect.TypeOf((*MockEcrClient)(nil).DescribeRepositories), varargs...)
}

func (m *MockEcrClient) GetAuthorizationToken(arg0 context.Context, arg1 *ecr.GetAuthorizationTokenInput, arg2 ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAuthorizationToken, varargs...)
	ret0, _ := ret[0].(*ecr.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetAuthorizationToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAuthorizationToken, reflect.TypeOf((*MockEcrClient)(nil).GetAuthorizationToken), varargs...)
}

func (m *MockEcrClient) GetDownloadUrlForLayer(arg0 context.Context, arg1 *ecr.GetDownloadUrlForLayerInput, arg2 ...func(*ecr.Options)) (*ecr.GetDownloadUrlForLayerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDownloadUrlForLayer, varargs...)
	ret0, _ := ret[0].(*ecr.GetDownloadUrlForLayerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetDownloadUrlForLayer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDownloadUrlForLayer, reflect.TypeOf((*MockEcrClient)(nil).GetDownloadUrlForLayer), varargs...)
}

func (m *MockEcrClient) GetLifecyclePolicy(arg0 context.Context, arg1 *ecr.GetLifecyclePolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLifecyclePolicy, varargs...)
	ret0, _ := ret[0].(*ecr.GetLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetLifecyclePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLifecyclePolicy, reflect.TypeOf((*MockEcrClient)(nil).GetLifecyclePolicy), varargs...)
}

func (m *MockEcrClient) GetLifecyclePolicyPreview(arg0 context.Context, arg1 *ecr.GetLifecyclePolicyPreviewInput, arg2 ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLifecyclePolicyPreview, varargs...)
	ret0, _ := ret[0].(*ecr.GetLifecyclePolicyPreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetLifecyclePolicyPreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLifecyclePolicyPreview, reflect.TypeOf((*MockEcrClient)(nil).GetLifecyclePolicyPreview), varargs...)
}

func (m *MockEcrClient) GetRegistryPolicy(arg0 context.Context, arg1 *ecr.GetRegistryPolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetRegistryPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegistryPolicy, varargs...)
	ret0, _ := ret[0].(*ecr.GetRegistryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetRegistryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegistryPolicy, reflect.TypeOf((*MockEcrClient)(nil).GetRegistryPolicy), varargs...)
}

func (m *MockEcrClient) GetRegistryScanningConfiguration(arg0 context.Context, arg1 *ecr.GetRegistryScanningConfigurationInput, arg2 ...func(*ecr.Options)) (*ecr.GetRegistryScanningConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegistryScanningConfiguration, varargs...)
	ret0, _ := ret[0].(*ecr.GetRegistryScanningConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetRegistryScanningConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegistryScanningConfiguration, reflect.TypeOf((*MockEcrClient)(nil).GetRegistryScanningConfiguration), varargs...)
}

func (m *MockEcrClient) GetRepositoryPolicy(arg0 context.Context, arg1 *ecr.GetRepositoryPolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetRepositoryPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRepositoryPolicy, varargs...)
	ret0, _ := ret[0].(*ecr.GetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetRepositoryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRepositoryPolicy, reflect.TypeOf((*MockEcrClient)(nil).GetRepositoryPolicy), varargs...)
}

func (m *MockEcrClient) ListImages(arg0 context.Context, arg1 *ecr.ListImagesInput, arg2 ...func(*ecr.Options)) (*ecr.ListImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImages, varargs...)
	ret0, _ := ret[0].(*ecr.ListImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) ListImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImages, reflect.TypeOf((*MockEcrClient)(nil).ListImages), varargs...)
}

func (m *MockEcrClient) ListTagsForResource(arg0 context.Context, arg1 *ecr.ListTagsForResourceInput, arg2 ...func(*ecr.Options)) (*ecr.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*ecr.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockEcrClient)(nil).ListTagsForResource), varargs...)
}
