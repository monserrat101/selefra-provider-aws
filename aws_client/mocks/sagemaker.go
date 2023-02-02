package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	sagemaker "github.com/aws/aws-sdk-go-v2/service/sagemaker"
	gomock "github.com/golang/mock/gomock"
)

type MockSagemakerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSagemakerClientMockRecorder
}

type MockSagemakerClientMockRecorder struct {
	mock *MockSagemakerClient
}

func NewMockSagemakerClient(ctrl *gomock.Controller) *MockSagemakerClient {
	mock := &MockSagemakerClient{ctrl: ctrl}
	mock.recorder = &MockSagemakerClientMockRecorder{mock}
	return mock
}

func (m *MockSagemakerClient) EXPECT() *MockSagemakerClientMockRecorder {
	return m.recorder
}

func (m *MockSagemakerClient) DescribeAction(arg0 context.Context, arg1 *sagemaker.DescribeActionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeActionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAction, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeAction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAction, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeAction), varargs...)
}

func (m *MockSagemakerClient) DescribeAlgorithm(arg0 context.Context, arg1 *sagemaker.DescribeAlgorithmInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeAlgorithmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAlgorithm, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeAlgorithmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeAlgorithm(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAlgorithm, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeAlgorithm), varargs...)
}

func (m *MockSagemakerClient) DescribeApp(arg0 context.Context, arg1 *sagemaker.DescribeAppInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeAppOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeApp, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeAppOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeApp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeApp, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeApp), varargs...)
}

func (m *MockSagemakerClient) DescribeAppImageConfig(arg0 context.Context, arg1 *sagemaker.DescribeAppImageConfigInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeAppImageConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAppImageConfig, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeAppImageConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeAppImageConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAppImageConfig, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeAppImageConfig), varargs...)
}

func (m *MockSagemakerClient) DescribeArtifact(arg0 context.Context, arg1 *sagemaker.DescribeArtifactInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeArtifactOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeArtifact, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeArtifactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeArtifact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeArtifact, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeArtifact), varargs...)
}

func (m *MockSagemakerClient) DescribeAutoMLJob(arg0 context.Context, arg1 *sagemaker.DescribeAutoMLJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeAutoMLJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutoMLJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeAutoMLJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeAutoMLJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutoMLJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeAutoMLJob), varargs...)
}

func (m *MockSagemakerClient) DescribeCodeRepository(arg0 context.Context, arg1 *sagemaker.DescribeCodeRepositoryInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeCodeRepositoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCodeRepository, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeCodeRepositoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeCodeRepository(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCodeRepository, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeCodeRepository), varargs...)
}

func (m *MockSagemakerClient) DescribeCompilationJob(arg0 context.Context, arg1 *sagemaker.DescribeCompilationJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeCompilationJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCompilationJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeCompilationJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeCompilationJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCompilationJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeCompilationJob), varargs...)
}

func (m *MockSagemakerClient) DescribeContext(arg0 context.Context, arg1 *sagemaker.DescribeContextInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeContextOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeContext, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeContextOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeContext, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeContext), varargs...)
}

func (m *MockSagemakerClient) DescribeDataQualityJobDefinition(arg0 context.Context, arg1 *sagemaker.DescribeDataQualityJobDefinitionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeDataQualityJobDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDataQualityJobDefinition, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeDataQualityJobDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeDataQualityJobDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDataQualityJobDefinition, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeDataQualityJobDefinition), varargs...)
}

func (m *MockSagemakerClient) DescribeDevice(arg0 context.Context, arg1 *sagemaker.DescribeDeviceInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeDeviceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDevice, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeDeviceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeDevice(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDevice, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeDevice), varargs...)
}

func (m *MockSagemakerClient) DescribeDeviceFleet(arg0 context.Context, arg1 *sagemaker.DescribeDeviceFleetInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeDeviceFleetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDeviceFleet, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeDeviceFleetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeDeviceFleet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDeviceFleet, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeDeviceFleet), varargs...)
}

func (m *MockSagemakerClient) DescribeDomain(arg0 context.Context, arg1 *sagemaker.DescribeDomainInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDomain, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDomain, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeDomain), varargs...)
}

func (m *MockSagemakerClient) DescribeEdgeDeploymentPlan(arg0 context.Context, arg1 *sagemaker.DescribeEdgeDeploymentPlanInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeEdgeDeploymentPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEdgeDeploymentPlan, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeEdgeDeploymentPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeEdgeDeploymentPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEdgeDeploymentPlan, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeEdgeDeploymentPlan), varargs...)
}

func (m *MockSagemakerClient) DescribeEdgePackagingJob(arg0 context.Context, arg1 *sagemaker.DescribeEdgePackagingJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeEdgePackagingJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEdgePackagingJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeEdgePackagingJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeEdgePackagingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEdgePackagingJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeEdgePackagingJob), varargs...)
}

func (m *MockSagemakerClient) DescribeEndpoint(arg0 context.Context, arg1 *sagemaker.DescribeEndpointInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpoint, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeEndpoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpoint, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeEndpoint), varargs...)
}

func (m *MockSagemakerClient) DescribeEndpointConfig(arg0 context.Context, arg1 *sagemaker.DescribeEndpointConfigInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEndpointConfig, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeEndpointConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeEndpointConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEndpointConfig, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeEndpointConfig), varargs...)
}

func (m *MockSagemakerClient) DescribeExperiment(arg0 context.Context, arg1 *sagemaker.DescribeExperimentInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeExperimentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeExperiment, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeExperimentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeExperiment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeExperiment, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeExperiment), varargs...)
}

func (m *MockSagemakerClient) DescribeFeatureGroup(arg0 context.Context, arg1 *sagemaker.DescribeFeatureGroupInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeFeatureGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFeatureGroup, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeFeatureGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeFeatureGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFeatureGroup, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeFeatureGroup), varargs...)
}

func (m *MockSagemakerClient) DescribeFeatureMetadata(arg0 context.Context, arg1 *sagemaker.DescribeFeatureMetadataInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeFeatureMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFeatureMetadata, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeFeatureMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeFeatureMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFeatureMetadata, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeFeatureMetadata), varargs...)
}

func (m *MockSagemakerClient) DescribeFlowDefinition(arg0 context.Context, arg1 *sagemaker.DescribeFlowDefinitionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeFlowDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeFlowDefinition, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeFlowDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeFlowDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeFlowDefinition, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeFlowDefinition), varargs...)
}

func (m *MockSagemakerClient) DescribeHub(arg0 context.Context, arg1 *sagemaker.DescribeHubInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeHubOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHub, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeHubOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeHub(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHub, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeHub), varargs...)
}

func (m *MockSagemakerClient) DescribeHubContent(arg0 context.Context, arg1 *sagemaker.DescribeHubContentInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeHubContentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHubContent, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeHubContentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeHubContent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHubContent, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeHubContent), varargs...)
}

func (m *MockSagemakerClient) DescribeHumanTaskUi(arg0 context.Context, arg1 *sagemaker.DescribeHumanTaskUiInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeHumanTaskUiOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHumanTaskUi, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeHumanTaskUiOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeHumanTaskUi(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHumanTaskUi, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeHumanTaskUi), varargs...)
}

func (m *MockSagemakerClient) DescribeHyperParameterTuningJob(arg0 context.Context, arg1 *sagemaker.DescribeHyperParameterTuningJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeHyperParameterTuningJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeHyperParameterTuningJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeHyperParameterTuningJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeHyperParameterTuningJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeHyperParameterTuningJob), varargs...)
}

func (m *MockSagemakerClient) DescribeImage(arg0 context.Context, arg1 *sagemaker.DescribeImageInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeImageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImage, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImage, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeImage), varargs...)
}

func (m *MockSagemakerClient) DescribeImageVersion(arg0 context.Context, arg1 *sagemaker.DescribeImageVersionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeImageVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeImageVersion, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeImageVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeImageVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeImageVersion, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeImageVersion), varargs...)
}

func (m *MockSagemakerClient) DescribeInferenceExperiment(arg0 context.Context, arg1 *sagemaker.DescribeInferenceExperimentInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeInferenceExperimentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInferenceExperiment, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeInferenceExperimentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeInferenceExperiment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInferenceExperiment, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeInferenceExperiment), varargs...)
}

func (m *MockSagemakerClient) DescribeInferenceRecommendationsJob(arg0 context.Context, arg1 *sagemaker.DescribeInferenceRecommendationsJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeInferenceRecommendationsJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInferenceRecommendationsJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeInferenceRecommendationsJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeInferenceRecommendationsJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInferenceRecommendationsJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeInferenceRecommendationsJob), varargs...)
}

func (m *MockSagemakerClient) DescribeLabelingJob(arg0 context.Context, arg1 *sagemaker.DescribeLabelingJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeLabelingJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLabelingJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeLabelingJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeLabelingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLabelingJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeLabelingJob), varargs...)
}

func (m *MockSagemakerClient) DescribeLineageGroup(arg0 context.Context, arg1 *sagemaker.DescribeLineageGroupInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeLineageGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLineageGroup, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeLineageGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeLineageGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLineageGroup, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeLineageGroup), varargs...)
}

func (m *MockSagemakerClient) DescribeModel(arg0 context.Context, arg1 *sagemaker.DescribeModelInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModel, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModel, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModel), varargs...)
}

func (m *MockSagemakerClient) DescribeModelBiasJobDefinition(arg0 context.Context, arg1 *sagemaker.DescribeModelBiasJobDefinitionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelBiasJobDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelBiasJobDefinition, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelBiasJobDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelBiasJobDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelBiasJobDefinition, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelBiasJobDefinition), varargs...)
}

func (m *MockSagemakerClient) DescribeModelCard(arg0 context.Context, arg1 *sagemaker.DescribeModelCardInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelCardOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelCard, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelCardOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelCard(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelCard, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelCard), varargs...)
}

func (m *MockSagemakerClient) DescribeModelCardExportJob(arg0 context.Context, arg1 *sagemaker.DescribeModelCardExportJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelCardExportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelCardExportJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelCardExportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelCardExportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelCardExportJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelCardExportJob), varargs...)
}

func (m *MockSagemakerClient) DescribeModelExplainabilityJobDefinition(arg0 context.Context, arg1 *sagemaker.DescribeModelExplainabilityJobDefinitionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelExplainabilityJobDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelExplainabilityJobDefinition, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelExplainabilityJobDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelExplainabilityJobDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelExplainabilityJobDefinition, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelExplainabilityJobDefinition), varargs...)
}

func (m *MockSagemakerClient) DescribeModelPackage(arg0 context.Context, arg1 *sagemaker.DescribeModelPackageInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelPackageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelPackage, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelPackageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelPackage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelPackage, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelPackage), varargs...)
}

func (m *MockSagemakerClient) DescribeModelPackageGroup(arg0 context.Context, arg1 *sagemaker.DescribeModelPackageGroupInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelPackageGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelPackageGroup, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelPackageGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelPackageGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelPackageGroup, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelPackageGroup), varargs...)
}

func (m *MockSagemakerClient) DescribeModelQualityJobDefinition(arg0 context.Context, arg1 *sagemaker.DescribeModelQualityJobDefinitionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelQualityJobDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelQualityJobDefinition, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelQualityJobDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeModelQualityJobDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelQualityJobDefinition, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeModelQualityJobDefinition), varargs...)
}

func (m *MockSagemakerClient) DescribeMonitoringSchedule(arg0 context.Context, arg1 *sagemaker.DescribeMonitoringScheduleInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMonitoringSchedule, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeMonitoringScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeMonitoringSchedule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMonitoringSchedule, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeMonitoringSchedule), varargs...)
}

func (m *MockSagemakerClient) DescribeNotebookInstance(arg0 context.Context, arg1 *sagemaker.DescribeNotebookInstanceInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNotebookInstance, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeNotebookInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeNotebookInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNotebookInstance, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeNotebookInstance), varargs...)
}

func (m *MockSagemakerClient) DescribeNotebookInstanceLifecycleConfig(arg0 context.Context, arg1 *sagemaker.DescribeNotebookInstanceLifecycleConfigInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeNotebookInstanceLifecycleConfig, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeNotebookInstanceLifecycleConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeNotebookInstanceLifecycleConfig, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeNotebookInstanceLifecycleConfig), varargs...)
}

func (m *MockSagemakerClient) DescribePipeline(arg0 context.Context, arg1 *sagemaker.DescribePipelineInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePipeline, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribePipeline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePipeline, reflect.TypeOf((*MockSagemakerClient)(nil).DescribePipeline), varargs...)
}

func (m *MockSagemakerClient) DescribePipelineDefinitionForExecution(arg0 context.Context, arg1 *sagemaker.DescribePipelineDefinitionForExecutionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineDefinitionForExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePipelineDefinitionForExecution, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribePipelineDefinitionForExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribePipelineDefinitionForExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePipelineDefinitionForExecution, reflect.TypeOf((*MockSagemakerClient)(nil).DescribePipelineDefinitionForExecution), varargs...)
}

func (m *MockSagemakerClient) DescribePipelineExecution(arg0 context.Context, arg1 *sagemaker.DescribePipelineExecutionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePipelineExecution, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribePipelineExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribePipelineExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePipelineExecution, reflect.TypeOf((*MockSagemakerClient)(nil).DescribePipelineExecution), varargs...)
}

func (m *MockSagemakerClient) DescribeProcessingJob(arg0 context.Context, arg1 *sagemaker.DescribeProcessingJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeProcessingJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProcessingJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeProcessingJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeProcessingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProcessingJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeProcessingJob), varargs...)
}

func (m *MockSagemakerClient) DescribeProject(arg0 context.Context, arg1 *sagemaker.DescribeProjectInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProject, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeProject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProject, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeProject), varargs...)
}

func (m *MockSagemakerClient) DescribeSpace(arg0 context.Context, arg1 *sagemaker.DescribeSpaceInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeSpaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSpace, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeSpaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeSpace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSpace, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeSpace), varargs...)
}

func (m *MockSagemakerClient) DescribeStudioLifecycleConfig(arg0 context.Context, arg1 *sagemaker.DescribeStudioLifecycleConfigInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeStudioLifecycleConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeStudioLifecycleConfig, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeStudioLifecycleConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeStudioLifecycleConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeStudioLifecycleConfig, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeStudioLifecycleConfig), varargs...)
}

func (m *MockSagemakerClient) DescribeSubscribedWorkteam(arg0 context.Context, arg1 *sagemaker.DescribeSubscribedWorkteamInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSubscribedWorkteam, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeSubscribedWorkteamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeSubscribedWorkteam(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSubscribedWorkteam, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeSubscribedWorkteam), varargs...)
}

func (m *MockSagemakerClient) DescribeTrainingJob(arg0 context.Context, arg1 *sagemaker.DescribeTrainingJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeTrainingJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrainingJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeTrainingJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeTrainingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrainingJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeTrainingJob), varargs...)
}

func (m *MockSagemakerClient) DescribeTransformJob(arg0 context.Context, arg1 *sagemaker.DescribeTransformJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeTransformJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTransformJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeTransformJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeTransformJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTransformJob, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeTransformJob), varargs...)
}

func (m *MockSagemakerClient) DescribeTrial(arg0 context.Context, arg1 *sagemaker.DescribeTrialInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeTrialOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrial, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeTrialOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeTrial(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrial, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeTrial), varargs...)
}

func (m *MockSagemakerClient) DescribeTrialComponent(arg0 context.Context, arg1 *sagemaker.DescribeTrialComponentInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeTrialComponentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTrialComponent, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeTrialComponentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeTrialComponent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTrialComponent, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeTrialComponent), varargs...)
}

func (m *MockSagemakerClient) DescribeUserProfile(arg0 context.Context, arg1 *sagemaker.DescribeUserProfileInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeUserProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeUserProfile, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeUserProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeUserProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeUserProfile, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeUserProfile), varargs...)
}

func (m *MockSagemakerClient) DescribeWorkforce(arg0 context.Context, arg1 *sagemaker.DescribeWorkforceInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeWorkforceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkforce, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeWorkforceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeWorkforce(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkforce, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeWorkforce), varargs...)
}

func (m *MockSagemakerClient) DescribeWorkteam(arg0 context.Context, arg1 *sagemaker.DescribeWorkteamInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeWorkteamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeWorkteam, varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeWorkteamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) DescribeWorkteam(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeWorkteam, reflect.TypeOf((*MockSagemakerClient)(nil).DescribeWorkteam), varargs...)
}

func (m *MockSagemakerClient) GetDeviceFleetReport(arg0 context.Context, arg1 *sagemaker.GetDeviceFleetReportInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.GetDeviceFleetReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeviceFleetReport, varargs...)
	ret0, _ := ret[0].(*sagemaker.GetDeviceFleetReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) GetDeviceFleetReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeviceFleetReport, reflect.TypeOf((*MockSagemakerClient)(nil).GetDeviceFleetReport), varargs...)
}

func (m *MockSagemakerClient) GetLineageGroupPolicy(arg0 context.Context, arg1 *sagemaker.GetLineageGroupPolicyInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.GetLineageGroupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLineageGroupPolicy, varargs...)
	ret0, _ := ret[0].(*sagemaker.GetLineageGroupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) GetLineageGroupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLineageGroupPolicy, reflect.TypeOf((*MockSagemakerClient)(nil).GetLineageGroupPolicy), varargs...)
}

func (m *MockSagemakerClient) GetModelPackageGroupPolicy(arg0 context.Context, arg1 *sagemaker.GetModelPackageGroupPolicyInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.GetModelPackageGroupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModelPackageGroupPolicy, varargs...)
	ret0, _ := ret[0].(*sagemaker.GetModelPackageGroupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) GetModelPackageGroupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModelPackageGroupPolicy, reflect.TypeOf((*MockSagemakerClient)(nil).GetModelPackageGroupPolicy), varargs...)
}

func (m *MockSagemakerClient) GetSagemakerServicecatalogPortfolioStatus(arg0 context.Context, arg1 *sagemaker.GetSagemakerServicecatalogPortfolioStatusInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSagemakerServicecatalogPortfolioStatus, varargs...)
	ret0, _ := ret[0].(*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) GetSagemakerServicecatalogPortfolioStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSagemakerServicecatalogPortfolioStatus, reflect.TypeOf((*MockSagemakerClient)(nil).GetSagemakerServicecatalogPortfolioStatus), varargs...)
}

func (m *MockSagemakerClient) GetSearchSuggestions(arg0 context.Context, arg1 *sagemaker.GetSearchSuggestionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.GetSearchSuggestionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSearchSuggestions, varargs...)
	ret0, _ := ret[0].(*sagemaker.GetSearchSuggestionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) GetSearchSuggestions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSearchSuggestions, reflect.TypeOf((*MockSagemakerClient)(nil).GetSearchSuggestions), varargs...)
}

func (m *MockSagemakerClient) ListActions(arg0 context.Context, arg1 *sagemaker.ListActionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListActions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListActions, reflect.TypeOf((*MockSagemakerClient)(nil).ListActions), varargs...)
}

func (m *MockSagemakerClient) ListAlgorithms(arg0 context.Context, arg1 *sagemaker.ListAlgorithmsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListAlgorithmsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAlgorithms, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListAlgorithmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListAlgorithms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAlgorithms, reflect.TypeOf((*MockSagemakerClient)(nil).ListAlgorithms), varargs...)
}

func (m *MockSagemakerClient) ListAliases(arg0 context.Context, arg1 *sagemaker.ListAliasesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAliases, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAliases, reflect.TypeOf((*MockSagemakerClient)(nil).ListAliases), varargs...)
}

func (m *MockSagemakerClient) ListAppImageConfigs(arg0 context.Context, arg1 *sagemaker.ListAppImageConfigsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListAppImageConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAppImageConfigs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListAppImageConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListAppImageConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAppImageConfigs, reflect.TypeOf((*MockSagemakerClient)(nil).ListAppImageConfigs), varargs...)
}

func (m *MockSagemakerClient) ListApps(arg0 context.Context, arg1 *sagemaker.ListAppsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListAppsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListApps, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListAppsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListApps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListApps, reflect.TypeOf((*MockSagemakerClient)(nil).ListApps), varargs...)
}

func (m *MockSagemakerClient) ListArtifacts(arg0 context.Context, arg1 *sagemaker.ListArtifactsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListArtifactsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListArtifacts, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListArtifactsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListArtifacts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListArtifacts, reflect.TypeOf((*MockSagemakerClient)(nil).ListArtifacts), varargs...)
}

func (m *MockSagemakerClient) ListAssociations(arg0 context.Context, arg1 *sagemaker.ListAssociationsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociations, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociations, reflect.TypeOf((*MockSagemakerClient)(nil).ListAssociations), varargs...)
}

func (m *MockSagemakerClient) ListAutoMLJobs(arg0 context.Context, arg1 *sagemaker.ListAutoMLJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListAutoMLJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAutoMLJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListAutoMLJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListAutoMLJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAutoMLJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListAutoMLJobs), varargs...)
}

func (m *MockSagemakerClient) ListCandidatesForAutoMLJob(arg0 context.Context, arg1 *sagemaker.ListCandidatesForAutoMLJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCandidatesForAutoMLJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListCandidatesForAutoMLJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListCandidatesForAutoMLJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCandidatesForAutoMLJob, reflect.TypeOf((*MockSagemakerClient)(nil).ListCandidatesForAutoMLJob), varargs...)
}

func (m *MockSagemakerClient) ListCodeRepositories(arg0 context.Context, arg1 *sagemaker.ListCodeRepositoriesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListCodeRepositoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCodeRepositories, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListCodeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListCodeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCodeRepositories, reflect.TypeOf((*MockSagemakerClient)(nil).ListCodeRepositories), varargs...)
}

func (m *MockSagemakerClient) ListCompilationJobs(arg0 context.Context, arg1 *sagemaker.ListCompilationJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListCompilationJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCompilationJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListCompilationJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListCompilationJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCompilationJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListCompilationJobs), varargs...)
}

func (m *MockSagemakerClient) ListContexts(arg0 context.Context, arg1 *sagemaker.ListContextsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListContextsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListContexts, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListContextsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListContexts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListContexts, reflect.TypeOf((*MockSagemakerClient)(nil).ListContexts), varargs...)
}

func (m *MockSagemakerClient) ListDataQualityJobDefinitions(arg0 context.Context, arg1 *sagemaker.ListDataQualityJobDefinitionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListDataQualityJobDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDataQualityJobDefinitions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListDataQualityJobDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListDataQualityJobDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDataQualityJobDefinitions, reflect.TypeOf((*MockSagemakerClient)(nil).ListDataQualityJobDefinitions), varargs...)
}

func (m *MockSagemakerClient) ListDeviceFleets(arg0 context.Context, arg1 *sagemaker.ListDeviceFleetsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListDeviceFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDeviceFleets, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListDeviceFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListDeviceFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDeviceFleets, reflect.TypeOf((*MockSagemakerClient)(nil).ListDeviceFleets), varargs...)
}

func (m *MockSagemakerClient) ListDevices(arg0 context.Context, arg1 *sagemaker.ListDevicesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListDevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDevices, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListDevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListDevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDevices, reflect.TypeOf((*MockSagemakerClient)(nil).ListDevices), varargs...)
}

func (m *MockSagemakerClient) ListDomains(arg0 context.Context, arg1 *sagemaker.ListDomainsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDomains, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDomains, reflect.TypeOf((*MockSagemakerClient)(nil).ListDomains), varargs...)
}

func (m *MockSagemakerClient) ListEdgeDeploymentPlans(arg0 context.Context, arg1 *sagemaker.ListEdgeDeploymentPlansInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListEdgeDeploymentPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEdgeDeploymentPlans, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListEdgeDeploymentPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListEdgeDeploymentPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEdgeDeploymentPlans, reflect.TypeOf((*MockSagemakerClient)(nil).ListEdgeDeploymentPlans), varargs...)
}

func (m *MockSagemakerClient) ListEdgePackagingJobs(arg0 context.Context, arg1 *sagemaker.ListEdgePackagingJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListEdgePackagingJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEdgePackagingJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListEdgePackagingJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListEdgePackagingJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEdgePackagingJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListEdgePackagingJobs), varargs...)
}

func (m *MockSagemakerClient) ListEndpointConfigs(arg0 context.Context, arg1 *sagemaker.ListEndpointConfigsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListEndpointConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEndpointConfigs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListEndpointConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListEndpointConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEndpointConfigs, reflect.TypeOf((*MockSagemakerClient)(nil).ListEndpointConfigs), varargs...)
}

func (m *MockSagemakerClient) ListEndpoints(arg0 context.Context, arg1 *sagemaker.ListEndpointsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEndpoints, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEndpoints, reflect.TypeOf((*MockSagemakerClient)(nil).ListEndpoints), varargs...)
}

func (m *MockSagemakerClient) ListExperiments(arg0 context.Context, arg1 *sagemaker.ListExperimentsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListExperimentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListExperiments, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListExperimentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListExperiments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExperiments, reflect.TypeOf((*MockSagemakerClient)(nil).ListExperiments), varargs...)
}

func (m *MockSagemakerClient) ListFeatureGroups(arg0 context.Context, arg1 *sagemaker.ListFeatureGroupsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListFeatureGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFeatureGroups, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListFeatureGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListFeatureGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFeatureGroups, reflect.TypeOf((*MockSagemakerClient)(nil).ListFeatureGroups), varargs...)
}

func (m *MockSagemakerClient) ListFlowDefinitions(arg0 context.Context, arg1 *sagemaker.ListFlowDefinitionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListFlowDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListFlowDefinitions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListFlowDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListFlowDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListFlowDefinitions, reflect.TypeOf((*MockSagemakerClient)(nil).ListFlowDefinitions), varargs...)
}

func (m *MockSagemakerClient) ListHubContentVersions(arg0 context.Context, arg1 *sagemaker.ListHubContentVersionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListHubContentVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHubContentVersions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListHubContentVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListHubContentVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHubContentVersions, reflect.TypeOf((*MockSagemakerClient)(nil).ListHubContentVersions), varargs...)
}

func (m *MockSagemakerClient) ListHubContents(arg0 context.Context, arg1 *sagemaker.ListHubContentsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListHubContentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHubContents, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListHubContentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListHubContents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHubContents, reflect.TypeOf((*MockSagemakerClient)(nil).ListHubContents), varargs...)
}

func (m *MockSagemakerClient) ListHubs(arg0 context.Context, arg1 *sagemaker.ListHubsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListHubsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHubs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListHubsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListHubs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHubs, reflect.TypeOf((*MockSagemakerClient)(nil).ListHubs), varargs...)
}

func (m *MockSagemakerClient) ListHumanTaskUis(arg0 context.Context, arg1 *sagemaker.ListHumanTaskUisInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListHumanTaskUisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHumanTaskUis, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListHumanTaskUisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListHumanTaskUis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHumanTaskUis, reflect.TypeOf((*MockSagemakerClient)(nil).ListHumanTaskUis), varargs...)
}

func (m *MockSagemakerClient) ListHyperParameterTuningJobs(arg0 context.Context, arg1 *sagemaker.ListHyperParameterTuningJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListHyperParameterTuningJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListHyperParameterTuningJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListHyperParameterTuningJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHyperParameterTuningJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListHyperParameterTuningJobs), varargs...)
}

func (m *MockSagemakerClient) ListImageVersions(arg0 context.Context, arg1 *sagemaker.ListImageVersionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListImageVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImageVersions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListImageVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListImageVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImageVersions, reflect.TypeOf((*MockSagemakerClient)(nil).ListImageVersions), varargs...)
}

func (m *MockSagemakerClient) ListImages(arg0 context.Context, arg1 *sagemaker.ListImagesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListImages, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListImages, reflect.TypeOf((*MockSagemakerClient)(nil).ListImages), varargs...)
}

func (m *MockSagemakerClient) ListInferenceExperiments(arg0 context.Context, arg1 *sagemaker.ListInferenceExperimentsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListInferenceExperimentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInferenceExperiments, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListInferenceExperimentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListInferenceExperiments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInferenceExperiments, reflect.TypeOf((*MockSagemakerClient)(nil).ListInferenceExperiments), varargs...)
}

func (m *MockSagemakerClient) ListInferenceRecommendationsJobSteps(arg0 context.Context, arg1 *sagemaker.ListInferenceRecommendationsJobStepsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListInferenceRecommendationsJobStepsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInferenceRecommendationsJobSteps, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListInferenceRecommendationsJobStepsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListInferenceRecommendationsJobSteps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInferenceRecommendationsJobSteps, reflect.TypeOf((*MockSagemakerClient)(nil).ListInferenceRecommendationsJobSteps), varargs...)
}

func (m *MockSagemakerClient) ListInferenceRecommendationsJobs(arg0 context.Context, arg1 *sagemaker.ListInferenceRecommendationsJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListInferenceRecommendationsJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInferenceRecommendationsJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListInferenceRecommendationsJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListInferenceRecommendationsJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInferenceRecommendationsJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListInferenceRecommendationsJobs), varargs...)
}

func (m *MockSagemakerClient) ListLabelingJobs(arg0 context.Context, arg1 *sagemaker.ListLabelingJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListLabelingJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLabelingJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListLabelingJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListLabelingJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLabelingJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListLabelingJobs), varargs...)
}

func (m *MockSagemakerClient) ListLabelingJobsForWorkteam(arg0 context.Context, arg1 *sagemaker.ListLabelingJobsForWorkteamInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLabelingJobsForWorkteam, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListLabelingJobsForWorkteamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListLabelingJobsForWorkteam(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLabelingJobsForWorkteam, reflect.TypeOf((*MockSagemakerClient)(nil).ListLabelingJobsForWorkteam), varargs...)
}

func (m *MockSagemakerClient) ListLineageGroups(arg0 context.Context, arg1 *sagemaker.ListLineageGroupsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListLineageGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLineageGroups, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListLineageGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListLineageGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLineageGroups, reflect.TypeOf((*MockSagemakerClient)(nil).ListLineageGroups), varargs...)
}

func (m *MockSagemakerClient) ListModelBiasJobDefinitions(arg0 context.Context, arg1 *sagemaker.ListModelBiasJobDefinitionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelBiasJobDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelBiasJobDefinitions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelBiasJobDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelBiasJobDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelBiasJobDefinitions, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelBiasJobDefinitions), varargs...)
}

func (m *MockSagemakerClient) ListModelCardExportJobs(arg0 context.Context, arg1 *sagemaker.ListModelCardExportJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelCardExportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelCardExportJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelCardExportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelCardExportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelCardExportJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelCardExportJobs), varargs...)
}

func (m *MockSagemakerClient) ListModelCardVersions(arg0 context.Context, arg1 *sagemaker.ListModelCardVersionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelCardVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelCardVersions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelCardVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelCardVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelCardVersions, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelCardVersions), varargs...)
}

func (m *MockSagemakerClient) ListModelCards(arg0 context.Context, arg1 *sagemaker.ListModelCardsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelCardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelCards, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelCardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelCards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelCards, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelCards), varargs...)
}

func (m *MockSagemakerClient) ListModelExplainabilityJobDefinitions(arg0 context.Context, arg1 *sagemaker.ListModelExplainabilityJobDefinitionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelExplainabilityJobDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelExplainabilityJobDefinitions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelExplainabilityJobDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelExplainabilityJobDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelExplainabilityJobDefinitions, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelExplainabilityJobDefinitions), varargs...)
}

func (m *MockSagemakerClient) ListModelMetadata(arg0 context.Context, arg1 *sagemaker.ListModelMetadataInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelMetadata, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelMetadata, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelMetadata), varargs...)
}

func (m *MockSagemakerClient) ListModelPackageGroups(arg0 context.Context, arg1 *sagemaker.ListModelPackageGroupsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelPackageGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelPackageGroups, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelPackageGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelPackageGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelPackageGroups, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelPackageGroups), varargs...)
}

func (m *MockSagemakerClient) ListModelPackages(arg0 context.Context, arg1 *sagemaker.ListModelPackagesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelPackagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelPackages, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelPackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelPackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelPackages, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelPackages), varargs...)
}

func (m *MockSagemakerClient) ListModelQualityJobDefinitions(arg0 context.Context, arg1 *sagemaker.ListModelQualityJobDefinitionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelQualityJobDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModelQualityJobDefinitions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelQualityJobDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModelQualityJobDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModelQualityJobDefinitions, reflect.TypeOf((*MockSagemakerClient)(nil).ListModelQualityJobDefinitions), varargs...)
}

func (m *MockSagemakerClient) ListModels(arg0 context.Context, arg1 *sagemaker.ListModelsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListModels, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListModels, reflect.TypeOf((*MockSagemakerClient)(nil).ListModels), varargs...)
}

func (m *MockSagemakerClient) ListMonitoringAlertHistory(arg0 context.Context, arg1 *sagemaker.ListMonitoringAlertHistoryInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringAlertHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMonitoringAlertHistory, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListMonitoringAlertHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListMonitoringAlertHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMonitoringAlertHistory, reflect.TypeOf((*MockSagemakerClient)(nil).ListMonitoringAlertHistory), varargs...)
}

func (m *MockSagemakerClient) ListMonitoringAlerts(arg0 context.Context, arg1 *sagemaker.ListMonitoringAlertsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringAlertsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMonitoringAlerts, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListMonitoringAlertsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListMonitoringAlerts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMonitoringAlerts, reflect.TypeOf((*MockSagemakerClient)(nil).ListMonitoringAlerts), varargs...)
}

func (m *MockSagemakerClient) ListMonitoringExecutions(arg0 context.Context, arg1 *sagemaker.ListMonitoringExecutionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMonitoringExecutions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListMonitoringExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListMonitoringExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMonitoringExecutions, reflect.TypeOf((*MockSagemakerClient)(nil).ListMonitoringExecutions), varargs...)
}

func (m *MockSagemakerClient) ListMonitoringSchedules(arg0 context.Context, arg1 *sagemaker.ListMonitoringSchedulesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringSchedulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMonitoringSchedules, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListMonitoringSchedulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListMonitoringSchedules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMonitoringSchedules, reflect.TypeOf((*MockSagemakerClient)(nil).ListMonitoringSchedules), varargs...)
}

func (m *MockSagemakerClient) ListNotebookInstanceLifecycleConfigs(arg0 context.Context, arg1 *sagemaker.ListNotebookInstanceLifecycleConfigsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNotebookInstanceLifecycleConfigs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListNotebookInstanceLifecycleConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNotebookInstanceLifecycleConfigs, reflect.TypeOf((*MockSagemakerClient)(nil).ListNotebookInstanceLifecycleConfigs), varargs...)
}

func (m *MockSagemakerClient) ListNotebookInstances(arg0 context.Context, arg1 *sagemaker.ListNotebookInstancesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListNotebookInstances, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListNotebookInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListNotebookInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListNotebookInstances, reflect.TypeOf((*MockSagemakerClient)(nil).ListNotebookInstances), varargs...)
}

func (m *MockSagemakerClient) ListPipelineExecutionSteps(arg0 context.Context, arg1 *sagemaker.ListPipelineExecutionStepsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListPipelineExecutionStepsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelineExecutionSteps, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListPipelineExecutionStepsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListPipelineExecutionSteps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelineExecutionSteps, reflect.TypeOf((*MockSagemakerClient)(nil).ListPipelineExecutionSteps), varargs...)
}

func (m *MockSagemakerClient) ListPipelineExecutions(arg0 context.Context, arg1 *sagemaker.ListPipelineExecutionsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListPipelineExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelineExecutions, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListPipelineExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListPipelineExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelineExecutions, reflect.TypeOf((*MockSagemakerClient)(nil).ListPipelineExecutions), varargs...)
}

func (m *MockSagemakerClient) ListPipelineParametersForExecution(arg0 context.Context, arg1 *sagemaker.ListPipelineParametersForExecutionInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListPipelineParametersForExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelineParametersForExecution, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListPipelineParametersForExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListPipelineParametersForExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelineParametersForExecution, reflect.TypeOf((*MockSagemakerClient)(nil).ListPipelineParametersForExecution), varargs...)
}

func (m *MockSagemakerClient) ListPipelines(arg0 context.Context, arg1 *sagemaker.ListPipelinesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListPipelinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPipelines, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListPipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListPipelines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPipelines, reflect.TypeOf((*MockSagemakerClient)(nil).ListPipelines), varargs...)
}

func (m *MockSagemakerClient) ListProcessingJobs(arg0 context.Context, arg1 *sagemaker.ListProcessingJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListProcessingJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProcessingJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListProcessingJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListProcessingJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProcessingJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListProcessingJobs), varargs...)
}

func (m *MockSagemakerClient) ListProjects(arg0 context.Context, arg1 *sagemaker.ListProjectsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProjects, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProjects, reflect.TypeOf((*MockSagemakerClient)(nil).ListProjects), varargs...)
}

func (m *MockSagemakerClient) ListSpaces(arg0 context.Context, arg1 *sagemaker.ListSpacesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListSpacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSpaces, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListSpacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListSpaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSpaces, reflect.TypeOf((*MockSagemakerClient)(nil).ListSpaces), varargs...)
}

func (m *MockSagemakerClient) ListStageDevices(arg0 context.Context, arg1 *sagemaker.ListStageDevicesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListStageDevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStageDevices, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListStageDevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListStageDevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStageDevices, reflect.TypeOf((*MockSagemakerClient)(nil).ListStageDevices), varargs...)
}

func (m *MockSagemakerClient) ListStudioLifecycleConfigs(arg0 context.Context, arg1 *sagemaker.ListStudioLifecycleConfigsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListStudioLifecycleConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStudioLifecycleConfigs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListStudioLifecycleConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListStudioLifecycleConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStudioLifecycleConfigs, reflect.TypeOf((*MockSagemakerClient)(nil).ListStudioLifecycleConfigs), varargs...)
}

func (m *MockSagemakerClient) ListSubscribedWorkteams(arg0 context.Context, arg1 *sagemaker.ListSubscribedWorkteamsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSubscribedWorkteams, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListSubscribedWorkteamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListSubscribedWorkteams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSubscribedWorkteams, reflect.TypeOf((*MockSagemakerClient)(nil).ListSubscribedWorkteams), varargs...)
}

func (m *MockSagemakerClient) ListTags(arg0 context.Context, arg1 *sagemaker.ListTagsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockSagemakerClient)(nil).ListTags), varargs...)
}

func (m *MockSagemakerClient) ListTrainingJobs(arg0 context.Context, arg1 *sagemaker.ListTrainingJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrainingJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTrainingJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListTrainingJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrainingJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListTrainingJobs), varargs...)
}

func (m *MockSagemakerClient) ListTrainingJobsForHyperParameterTuningJob(arg0 context.Context, arg1 *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrainingJobsForHyperParameterTuningJob, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListTrainingJobsForHyperParameterTuningJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrainingJobsForHyperParameterTuningJob, reflect.TypeOf((*MockSagemakerClient)(nil).ListTrainingJobsForHyperParameterTuningJob), varargs...)
}

func (m *MockSagemakerClient) ListTransformJobs(arg0 context.Context, arg1 *sagemaker.ListTransformJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTransformJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTransformJobs, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTransformJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListTransformJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTransformJobs, reflect.TypeOf((*MockSagemakerClient)(nil).ListTransformJobs), varargs...)
}

func (m *MockSagemakerClient) ListTrialComponents(arg0 context.Context, arg1 *sagemaker.ListTrialComponentsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTrialComponentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrialComponents, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTrialComponentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListTrialComponents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrialComponents, reflect.TypeOf((*MockSagemakerClient)(nil).ListTrialComponents), varargs...)
}

func (m *MockSagemakerClient) ListTrials(arg0 context.Context, arg1 *sagemaker.ListTrialsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTrialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTrials, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTrialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListTrials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTrials, reflect.TypeOf((*MockSagemakerClient)(nil).ListTrials), varargs...)
}

func (m *MockSagemakerClient) ListUserProfiles(arg0 context.Context, arg1 *sagemaker.ListUserProfilesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListUserProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListUserProfiles, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListUserProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListUserProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListUserProfiles, reflect.TypeOf((*MockSagemakerClient)(nil).ListUserProfiles), varargs...)
}

func (m *MockSagemakerClient) ListWorkforces(arg0 context.Context, arg1 *sagemaker.ListWorkforcesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListWorkforcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWorkforces, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListWorkforcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListWorkforces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWorkforces, reflect.TypeOf((*MockSagemakerClient)(nil).ListWorkforces), varargs...)
}

func (m *MockSagemakerClient) ListWorkteams(arg0 context.Context, arg1 *sagemaker.ListWorkteamsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListWorkteamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListWorkteams, varargs...)
	ret0, _ := ret[0].(*sagemaker.ListWorkteamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) ListWorkteams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListWorkteams, reflect.TypeOf((*MockSagemakerClient)(nil).ListWorkteams), varargs...)
}

func (m *MockSagemakerClient) Search(arg0 context.Context, arg1 *sagemaker.SearchInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.SearchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Search, varargs...)
	ret0, _ := ret[0].(*sagemaker.SearchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSagemakerClientMockRecorder) Search(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Search, reflect.TypeOf((*MockSagemakerClient)(nil).Search), varargs...)
}
