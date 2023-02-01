package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	frauddetector "github.com/aws/aws-sdk-go-v2/service/frauddetector"
	gomock "github.com/golang/mock/gomock"
)

type MockFrauddetectorClient struct {
	ctrl		*gomock.Controller
	recorder	*MockFrauddetectorClientMockRecorder
}

type MockFrauddetectorClientMockRecorder struct {
	mock *MockFrauddetectorClient
}

func NewMockFrauddetectorClient(ctrl *gomock.Controller) *MockFrauddetectorClient {
	mock := &MockFrauddetectorClient{ctrl: ctrl}
	mock.recorder = &MockFrauddetectorClientMockRecorder{mock}
	return mock
}

func (m *MockFrauddetectorClient) EXPECT() *MockFrauddetectorClientMockRecorder {
	return m.recorder
}

func (m *MockFrauddetectorClient) BatchGetVariable(arg0 context.Context, arg1 *frauddetector.BatchGetVariableInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.BatchGetVariableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.BatchGetVariable, varargs...)
	ret0, _ := ret[0].(*frauddetector.BatchGetVariableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) BatchGetVariable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchGetVariable, reflect.TypeOf((*MockFrauddetectorClient)(nil).BatchGetVariable), varargs...)
}

func (m *MockFrauddetectorClient) DescribeDetector(arg0 context.Context, arg1 *frauddetector.DescribeDetectorInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.DescribeDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDetector, varargs...)
	ret0, _ := ret[0].(*frauddetector.DescribeDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) DescribeDetector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDetector, reflect.TypeOf((*MockFrauddetectorClient)(nil).DescribeDetector), varargs...)
}

func (m *MockFrauddetectorClient) DescribeModelVersions(arg0 context.Context, arg1 *frauddetector.DescribeModelVersionsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.DescribeModelVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeModelVersions, varargs...)
	ret0, _ := ret[0].(*frauddetector.DescribeModelVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) DescribeModelVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeModelVersions, reflect.TypeOf((*MockFrauddetectorClient)(nil).DescribeModelVersions), varargs...)
}

func (m *MockFrauddetectorClient) GetBatchImportJobs(arg0 context.Context, arg1 *frauddetector.GetBatchImportJobsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetBatchImportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBatchImportJobs, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetBatchImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetBatchImportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBatchImportJobs, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetBatchImportJobs), varargs...)
}

func (m *MockFrauddetectorClient) GetBatchPredictionJobs(arg0 context.Context, arg1 *frauddetector.GetBatchPredictionJobsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetBatchPredictionJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBatchPredictionJobs, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetBatchPredictionJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetBatchPredictionJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBatchPredictionJobs, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetBatchPredictionJobs), varargs...)
}

func (m *MockFrauddetectorClient) GetDeleteEventsByEventTypeStatus(arg0 context.Context, arg1 *frauddetector.GetDeleteEventsByEventTypeStatusInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeleteEventsByEventTypeStatus, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetDeleteEventsByEventTypeStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetDeleteEventsByEventTypeStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeleteEventsByEventTypeStatus, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetDeleteEventsByEventTypeStatus), varargs...)
}

func (m *MockFrauddetectorClient) GetDetectorVersion(arg0 context.Context, arg1 *frauddetector.GetDetectorVersionInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetDetectorVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDetectorVersion, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetDetectorVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetDetectorVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDetectorVersion, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetDetectorVersion), varargs...)
}

func (m *MockFrauddetectorClient) GetDetectors(arg0 context.Context, arg1 *frauddetector.GetDetectorsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDetectors, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDetectors, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetDetectors), varargs...)
}

func (m *MockFrauddetectorClient) GetEntityTypes(arg0 context.Context, arg1 *frauddetector.GetEntityTypesInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetEntityTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEntityTypes, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetEntityTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetEntityTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEntityTypes, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetEntityTypes), varargs...)
}

func (m *MockFrauddetectorClient) GetEvent(arg0 context.Context, arg1 *frauddetector.GetEventInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetEventOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEvent, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetEventOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEvent, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetEvent), varargs...)
}

func (m *MockFrauddetectorClient) GetEventPrediction(arg0 context.Context, arg1 *frauddetector.GetEventPredictionInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetEventPredictionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEventPrediction, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetEventPredictionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetEventPrediction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEventPrediction, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetEventPrediction), varargs...)
}

func (m *MockFrauddetectorClient) GetEventPredictionMetadata(arg0 context.Context, arg1 *frauddetector.GetEventPredictionMetadataInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetEventPredictionMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEventPredictionMetadata, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetEventPredictionMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetEventPredictionMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEventPredictionMetadata, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetEventPredictionMetadata), varargs...)
}

func (m *MockFrauddetectorClient) GetEventTypes(arg0 context.Context, arg1 *frauddetector.GetEventTypesInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetEventTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEventTypes, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetEventTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetEventTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEventTypes, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetEventTypes), varargs...)
}

func (m *MockFrauddetectorClient) GetExternalModels(arg0 context.Context, arg1 *frauddetector.GetExternalModelsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetExternalModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetExternalModels, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetExternalModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetExternalModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetExternalModels, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetExternalModels), varargs...)
}

func (m *MockFrauddetectorClient) GetKMSEncryptionKey(arg0 context.Context, arg1 *frauddetector.GetKMSEncryptionKeyInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKMSEncryptionKey, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetKMSEncryptionKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetKMSEncryptionKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKMSEncryptionKey, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetKMSEncryptionKey), varargs...)
}

func (m *MockFrauddetectorClient) GetLabels(arg0 context.Context, arg1 *frauddetector.GetLabelsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetLabelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLabels, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetLabelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetLabels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLabels, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetLabels), varargs...)
}

func (m *MockFrauddetectorClient) GetModelVersion(arg0 context.Context, arg1 *frauddetector.GetModelVersionInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetModelVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModelVersion, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetModelVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetModelVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModelVersion, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetModelVersion), varargs...)
}

func (m *MockFrauddetectorClient) GetModels(arg0 context.Context, arg1 *frauddetector.GetModelsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetModels, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetModels, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetModels), varargs...)
}

func (m *MockFrauddetectorClient) GetOutcomes(arg0 context.Context, arg1 *frauddetector.GetOutcomesInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetOutcomesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOutcomes, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetOutcomesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetOutcomes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOutcomes, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetOutcomes), varargs...)
}

func (m *MockFrauddetectorClient) GetRules(arg0 context.Context, arg1 *frauddetector.GetRulesInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRules, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRules, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetRules), varargs...)
}

func (m *MockFrauddetectorClient) GetVariables(arg0 context.Context, arg1 *frauddetector.GetVariablesInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.GetVariablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetVariables, varargs...)
	ret0, _ := ret[0].(*frauddetector.GetVariablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) GetVariables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetVariables, reflect.TypeOf((*MockFrauddetectorClient)(nil).GetVariables), varargs...)
}

func (m *MockFrauddetectorClient) ListEventPredictions(arg0 context.Context, arg1 *frauddetector.ListEventPredictionsInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.ListEventPredictionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEventPredictions, varargs...)
	ret0, _ := ret[0].(*frauddetector.ListEventPredictionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) ListEventPredictions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEventPredictions, reflect.TypeOf((*MockFrauddetectorClient)(nil).ListEventPredictions), varargs...)
}

func (m *MockFrauddetectorClient) ListTagsForResource(arg0 context.Context, arg1 *frauddetector.ListTagsForResourceInput, arg2 ...func(*frauddetector.Options)) (*frauddetector.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*frauddetector.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrauddetectorClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockFrauddetectorClient)(nil).ListTagsForResource), varargs...)
}
