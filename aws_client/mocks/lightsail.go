package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	lightsail "github.com/aws/aws-sdk-go-v2/service/lightsail"
	gomock "github.com/golang/mock/gomock"
)

type MockLightsailClient struct {
	ctrl		*gomock.Controller
	recorder	*MockLightsailClientMockRecorder
}

type MockLightsailClientMockRecorder struct {
	mock *MockLightsailClient
}

func NewMockLightsailClient(ctrl *gomock.Controller) *MockLightsailClient {
	mock := &MockLightsailClient{ctrl: ctrl}
	mock.recorder = &MockLightsailClientMockRecorder{mock}
	return mock
}

func (m *MockLightsailClient) EXPECT() *MockLightsailClientMockRecorder {
	return m.recorder
}

func (m *MockLightsailClient) GetActiveNames(arg0 context.Context, arg1 *lightsail.GetActiveNamesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetActiveNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetActiveNames, varargs...)
	ret0, _ := ret[0].(*lightsail.GetActiveNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetActiveNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetActiveNames, reflect.TypeOf((*MockLightsailClient)(nil).GetActiveNames), varargs...)
}

func (m *MockLightsailClient) GetAlarms(arg0 context.Context, arg1 *lightsail.GetAlarmsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetAlarmsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAlarms, varargs...)
	ret0, _ := ret[0].(*lightsail.GetAlarmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetAlarms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAlarms, reflect.TypeOf((*MockLightsailClient)(nil).GetAlarms), varargs...)
}

func (m *MockLightsailClient) GetAutoSnapshots(arg0 context.Context, arg1 *lightsail.GetAutoSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetAutoSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAutoSnapshots, varargs...)
	ret0, _ := ret[0].(*lightsail.GetAutoSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetAutoSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAutoSnapshots, reflect.TypeOf((*MockLightsailClient)(nil).GetAutoSnapshots), varargs...)
}

func (m *MockLightsailClient) GetBlueprints(arg0 context.Context, arg1 *lightsail.GetBlueprintsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBlueprintsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlueprints, varargs...)
	ret0, _ := ret[0].(*lightsail.GetBlueprintsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBlueprints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlueprints, reflect.TypeOf((*MockLightsailClient)(nil).GetBlueprints), varargs...)
}

func (m *MockLightsailClient) GetBucketAccessKeys(arg0 context.Context, arg1 *lightsail.GetBucketAccessKeysInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketAccessKeys, varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBucketAccessKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketAccessKeys, reflect.TypeOf((*MockLightsailClient)(nil).GetBucketAccessKeys), varargs...)
}

func (m *MockLightsailClient) GetBucketBundles(arg0 context.Context, arg1 *lightsail.GetBucketBundlesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketBundles, varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBucketBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketBundles, reflect.TypeOf((*MockLightsailClient)(nil).GetBucketBundles), varargs...)
}

func (m *MockLightsailClient) GetBucketMetricData(arg0 context.Context, arg1 *lightsail.GetBucketMetricDataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketMetricData, varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBucketMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketMetricData, reflect.TypeOf((*MockLightsailClient)(nil).GetBucketMetricData), varargs...)
}

func (m *MockLightsailClient) GetBuckets(arg0 context.Context, arg1 *lightsail.GetBucketsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBuckets, varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBuckets, reflect.TypeOf((*MockLightsailClient)(nil).GetBuckets), varargs...)
}

func (m *MockLightsailClient) GetBundles(arg0 context.Context, arg1 *lightsail.GetBundlesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBundles, varargs...)
	ret0, _ := ret[0].(*lightsail.GetBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBundles, reflect.TypeOf((*MockLightsailClient)(nil).GetBundles), varargs...)
}

func (m *MockLightsailClient) GetCertificates(arg0 context.Context, arg1 *lightsail.GetCertificatesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCertificates, varargs...)
	ret0, _ := ret[0].(*lightsail.GetCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCertificates, reflect.TypeOf((*MockLightsailClient)(nil).GetCertificates), varargs...)
}

func (m *MockLightsailClient) GetCloudFormationStackRecords(arg0 context.Context, arg1 *lightsail.GetCloudFormationStackRecordsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCloudFormationStackRecords, varargs...)
	ret0, _ := ret[0].(*lightsail.GetCloudFormationStackRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetCloudFormationStackRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCloudFormationStackRecords, reflect.TypeOf((*MockLightsailClient)(nil).GetCloudFormationStackRecords), varargs...)
}

func (m *MockLightsailClient) GetContactMethods(arg0 context.Context, arg1 *lightsail.GetContactMethodsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContactMethodsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContactMethods, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContactMethodsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContactMethods(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContactMethods, reflect.TypeOf((*MockLightsailClient)(nil).GetContactMethods), varargs...)
}

func (m *MockLightsailClient) GetContainerAPIMetadata(arg0 context.Context, arg1 *lightsail.GetContainerAPIMetadataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerAPIMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerAPIMetadata, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerAPIMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerAPIMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerAPIMetadata, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerAPIMetadata), varargs...)
}

func (m *MockLightsailClient) GetContainerImages(arg0 context.Context, arg1 *lightsail.GetContainerImagesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerImages, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerImages, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerImages), varargs...)
}

func (m *MockLightsailClient) GetContainerLog(arg0 context.Context, arg1 *lightsail.GetContainerLogInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerLogOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerLog, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerLogOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerLog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerLog, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerLog), varargs...)
}

func (m *MockLightsailClient) GetContainerServiceDeployments(arg0 context.Context, arg1 *lightsail.GetContainerServiceDeploymentsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerServiceDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerServiceDeployments, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerServiceDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerServiceDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerServiceDeployments, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerServiceDeployments), varargs...)
}

func (m *MockLightsailClient) GetContainerServiceMetricData(arg0 context.Context, arg1 *lightsail.GetContainerServiceMetricDataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerServiceMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerServiceMetricData, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerServiceMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerServiceMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerServiceMetricData, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerServiceMetricData), varargs...)
}

func (m *MockLightsailClient) GetContainerServicePowers(arg0 context.Context, arg1 *lightsail.GetContainerServicePowersInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerServicePowersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerServicePowers, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerServicePowersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerServicePowers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerServicePowers, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerServicePowers), varargs...)
}

func (m *MockLightsailClient) GetContainerServices(arg0 context.Context, arg1 *lightsail.GetContainerServicesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetContainerServices, varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetContainerServices, reflect.TypeOf((*MockLightsailClient)(nil).GetContainerServices), varargs...)
}

func (m *MockLightsailClient) GetDisk(arg0 context.Context, arg1 *lightsail.GetDiskInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDiskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDisk, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDiskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDisk(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDisk, reflect.TypeOf((*MockLightsailClient)(nil).GetDisk), varargs...)
}

func (m *MockLightsailClient) GetDiskSnapshot(arg0 context.Context, arg1 *lightsail.GetDiskSnapshotInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDiskSnapshot, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDiskSnapshotOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDiskSnapshot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDiskSnapshot, reflect.TypeOf((*MockLightsailClient)(nil).GetDiskSnapshot), varargs...)
}

func (m *MockLightsailClient) GetDiskSnapshots(arg0 context.Context, arg1 *lightsail.GetDiskSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDiskSnapshots, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDiskSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDiskSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDiskSnapshots, reflect.TypeOf((*MockLightsailClient)(nil).GetDiskSnapshots), varargs...)
}

func (m *MockLightsailClient) GetDisks(arg0 context.Context, arg1 *lightsail.GetDisksInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDisksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDisks, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDisksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDisks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDisks, reflect.TypeOf((*MockLightsailClient)(nil).GetDisks), varargs...)
}

func (m *MockLightsailClient) GetDistributionBundles(arg0 context.Context, arg1 *lightsail.GetDistributionBundlesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDistributionBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDistributionBundles, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDistributionBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDistributionBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDistributionBundles, reflect.TypeOf((*MockLightsailClient)(nil).GetDistributionBundles), varargs...)
}

func (m *MockLightsailClient) GetDistributionLatestCacheReset(arg0 context.Context, arg1 *lightsail.GetDistributionLatestCacheResetInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDistributionLatestCacheReset, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDistributionLatestCacheResetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDistributionLatestCacheReset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDistributionLatestCacheReset, reflect.TypeOf((*MockLightsailClient)(nil).GetDistributionLatestCacheReset), varargs...)
}

func (m *MockLightsailClient) GetDistributionMetricData(arg0 context.Context, arg1 *lightsail.GetDistributionMetricDataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDistributionMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDistributionMetricData, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDistributionMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDistributionMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDistributionMetricData, reflect.TypeOf((*MockLightsailClient)(nil).GetDistributionMetricData), varargs...)
}

func (m *MockLightsailClient) GetDistributions(arg0 context.Context, arg1 *lightsail.GetDistributionsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDistributionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDistributions, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDistributionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDistributions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDistributions, reflect.TypeOf((*MockLightsailClient)(nil).GetDistributions), varargs...)
}

func (m *MockLightsailClient) GetDomain(arg0 context.Context, arg1 *lightsail.GetDomainInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomain, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomain, reflect.TypeOf((*MockLightsailClient)(nil).GetDomain), varargs...)
}

func (m *MockLightsailClient) GetDomains(arg0 context.Context, arg1 *lightsail.GetDomainsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDomains, varargs...)
	ret0, _ := ret[0].(*lightsail.GetDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDomains, reflect.TypeOf((*MockLightsailClient)(nil).GetDomains), varargs...)
}

func (m *MockLightsailClient) GetExportSnapshotRecords(arg0 context.Context, arg1 *lightsail.GetExportSnapshotRecordsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetExportSnapshotRecords, varargs...)
	ret0, _ := ret[0].(*lightsail.GetExportSnapshotRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetExportSnapshotRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetExportSnapshotRecords, reflect.TypeOf((*MockLightsailClient)(nil).GetExportSnapshotRecords), varargs...)
}

func (m *MockLightsailClient) GetInstance(arg0 context.Context, arg1 *lightsail.GetInstanceInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstance, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstance, reflect.TypeOf((*MockLightsailClient)(nil).GetInstance), varargs...)
}

func (m *MockLightsailClient) GetInstanceAccessDetails(arg0 context.Context, arg1 *lightsail.GetInstanceAccessDetailsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceAccessDetails, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceAccessDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceAccessDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceAccessDetails, reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceAccessDetails), varargs...)
}

func (m *MockLightsailClient) GetInstanceMetricData(arg0 context.Context, arg1 *lightsail.GetInstanceMetricDataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceMetricData, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceMetricData, reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceMetricData), varargs...)
}

func (m *MockLightsailClient) GetInstancePortStates(arg0 context.Context, arg1 *lightsail.GetInstancePortStatesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstancePortStatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstancePortStates, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstancePortStatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstancePortStates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstancePortStates, reflect.TypeOf((*MockLightsailClient)(nil).GetInstancePortStates), varargs...)
}

func (m *MockLightsailClient) GetInstanceSnapshot(arg0 context.Context, arg1 *lightsail.GetInstanceSnapshotInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceSnapshot, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceSnapshotOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceSnapshot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceSnapshot, reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceSnapshot), varargs...)
}

func (m *MockLightsailClient) GetInstanceSnapshots(arg0 context.Context, arg1 *lightsail.GetInstanceSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceSnapshots, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceSnapshots, reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceSnapshots), varargs...)
}

func (m *MockLightsailClient) GetInstanceState(arg0 context.Context, arg1 *lightsail.GetInstanceStateInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstanceState, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstanceState, reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceState), varargs...)
}

func (m *MockLightsailClient) GetInstances(arg0 context.Context, arg1 *lightsail.GetInstancesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInstances, varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInstances, reflect.TypeOf((*MockLightsailClient)(nil).GetInstances), varargs...)
}

func (m *MockLightsailClient) GetKeyPair(arg0 context.Context, arg1 *lightsail.GetKeyPairInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetKeyPairOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKeyPair, varargs...)
	ret0, _ := ret[0].(*lightsail.GetKeyPairOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetKeyPair(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKeyPair, reflect.TypeOf((*MockLightsailClient)(nil).GetKeyPair), varargs...)
}

func (m *MockLightsailClient) GetKeyPairs(arg0 context.Context, arg1 *lightsail.GetKeyPairsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetKeyPairsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetKeyPairs, varargs...)
	ret0, _ := ret[0].(*lightsail.GetKeyPairsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetKeyPairs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetKeyPairs, reflect.TypeOf((*MockLightsailClient)(nil).GetKeyPairs), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancer(arg0 context.Context, arg1 *lightsail.GetLoadBalancerInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoadBalancer, varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoadBalancer, reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancer), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancerMetricData(arg0 context.Context, arg1 *lightsail.GetLoadBalancerMetricDataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoadBalancerMetricData, varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancerMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancerMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoadBalancerMetricData, reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancerMetricData), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancerTlsCertificates(arg0 context.Context, arg1 *lightsail.GetLoadBalancerTlsCertificatesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoadBalancerTlsCertificates, varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancerTlsCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancerTlsCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoadBalancerTlsCertificates, reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancerTlsCertificates), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancerTlsPolicies(arg0 context.Context, arg1 *lightsail.GetLoadBalancerTlsPoliciesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoadBalancerTlsPolicies, varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancerTlsPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancerTlsPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoadBalancerTlsPolicies, reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancerTlsPolicies), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancers(arg0 context.Context, arg1 *lightsail.GetLoadBalancersInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetLoadBalancers, varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLoadBalancers, reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancers), varargs...)
}

func (m *MockLightsailClient) GetOperation(arg0 context.Context, arg1 *lightsail.GetOperationInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetOperationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOperation, varargs...)
	ret0, _ := ret[0].(*lightsail.GetOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOperation, reflect.TypeOf((*MockLightsailClient)(nil).GetOperation), varargs...)
}

func (m *MockLightsailClient) GetOperations(arg0 context.Context, arg1 *lightsail.GetOperationsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOperations, varargs...)
	ret0, _ := ret[0].(*lightsail.GetOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOperations, reflect.TypeOf((*MockLightsailClient)(nil).GetOperations), varargs...)
}

func (m *MockLightsailClient) GetOperationsForResource(arg0 context.Context, arg1 *lightsail.GetOperationsForResourceInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetOperationsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOperationsForResource, varargs...)
	ret0, _ := ret[0].(*lightsail.GetOperationsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetOperationsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOperationsForResource, reflect.TypeOf((*MockLightsailClient)(nil).GetOperationsForResource), varargs...)
}

func (m *MockLightsailClient) GetRegions(arg0 context.Context, arg1 *lightsail.GetRegionsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRegionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRegions, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRegionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRegions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRegions, reflect.TypeOf((*MockLightsailClient)(nil).GetRegions), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabase(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabase, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabase, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabase), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseBlueprints(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseBlueprintsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseBlueprints, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseBlueprintsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseBlueprints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseBlueprints, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseBlueprints), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseBundles(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseBundlesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseBundles, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseBundles, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseBundles), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseEvents(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseEventsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseEvents, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseEvents, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseEvents), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseLogEvents(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseLogEventsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseLogEvents, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseLogEvents, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseLogEvents), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseLogStreams(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseLogStreamsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseLogStreams, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseLogStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseLogStreams, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseLogStreams), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseMasterUserPassword(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseMasterUserPasswordInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseMasterUserPassword, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseMasterUserPasswordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseMasterUserPassword(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseMasterUserPassword, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseMasterUserPassword), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseMetricData(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseMetricDataInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseMetricData, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseMetricData, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseMetricData), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseParameters(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseParametersInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseParameters, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseParameters, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseParameters), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseSnapshot(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseSnapshotInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseSnapshot, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseSnapshotOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseSnapshot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseSnapshot, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseSnapshot), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseSnapshots(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabaseSnapshots, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabaseSnapshots, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseSnapshots), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabases(arg0 context.Context, arg1 *lightsail.GetRelationalDatabasesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRelationalDatabases, varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRelationalDatabases, reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabases), varargs...)
}

func (m *MockLightsailClient) GetStaticIp(arg0 context.Context, arg1 *lightsail.GetStaticIpInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetStaticIpOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStaticIp, varargs...)
	ret0, _ := ret[0].(*lightsail.GetStaticIpOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetStaticIp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStaticIp, reflect.TypeOf((*MockLightsailClient)(nil).GetStaticIp), varargs...)
}

func (m *MockLightsailClient) GetStaticIps(arg0 context.Context, arg1 *lightsail.GetStaticIpsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetStaticIpsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStaticIps, varargs...)
	ret0, _ := ret[0].(*lightsail.GetStaticIpsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetStaticIps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStaticIps, reflect.TypeOf((*MockLightsailClient)(nil).GetStaticIps), varargs...)
}
