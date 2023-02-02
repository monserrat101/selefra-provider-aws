package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

type MockS3Client struct {
	ctrl		*gomock.Controller
	recorder	*MockS3ClientMockRecorder
}

type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

func (m *MockS3Client) GetBucketAccelerateConfiguration(arg0 context.Context, arg1 *s3.GetBucketAccelerateConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketAccelerateConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAccelerateConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketAccelerateConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketAccelerateConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketAccelerateConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketAcl(arg0 context.Context, arg1 *s3.GetBucketAclInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketAcl, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAclOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketAcl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketAcl, reflect.TypeOf((*MockS3Client)(nil).GetBucketAcl), varargs...)
}

func (m *MockS3Client) GetBucketAnalyticsConfiguration(arg0 context.Context, arg1 *s3.GetBucketAnalyticsConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketAnalyticsConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAnalyticsConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketAnalyticsConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketAnalyticsConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketAnalyticsConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketCors(arg0 context.Context, arg1 *s3.GetBucketCorsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketCors, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketCorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketCors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketCors, reflect.TypeOf((*MockS3Client)(nil).GetBucketCors), varargs...)
}

func (m *MockS3Client) GetBucketEncryption(arg0 context.Context, arg1 *s3.GetBucketEncryptionInput, arg2 ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketEncryption, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketEncryption(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketEncryption, reflect.TypeOf((*MockS3Client)(nil).GetBucketEncryption), varargs...)
}

func (m *MockS3Client) GetBucketIntelligentTieringConfiguration(arg0 context.Context, arg1 *s3.GetBucketIntelligentTieringConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketIntelligentTieringConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketIntelligentTieringConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketIntelligentTieringConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketIntelligentTieringConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketIntelligentTieringConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketInventoryConfiguration(arg0 context.Context, arg1 *s3.GetBucketInventoryConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketInventoryConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketInventoryConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketInventoryConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketInventoryConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketInventoryConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketLifecycleConfiguration(arg0 context.Context, arg1 *s3.GetBucketLifecycleConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketLifecycleConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketLifecycleConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketLifecycleConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketLocation(arg0 context.Context, arg1 *s3.GetBucketLocationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketLocation, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketLocation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketLocation, reflect.TypeOf((*MockS3Client)(nil).GetBucketLocation), varargs...)
}

func (m *MockS3Client) GetBucketLogging(arg0 context.Context, arg1 *s3.GetBucketLoggingInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketLogging, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketLogging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketLogging, reflect.TypeOf((*MockS3Client)(nil).GetBucketLogging), varargs...)
}

func (m *MockS3Client) GetBucketMetricsConfiguration(arg0 context.Context, arg1 *s3.GetBucketMetricsConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketMetricsConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketMetricsConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketMetricsConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketMetricsConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketMetricsConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketNotificationConfiguration(arg0 context.Context, arg1 *s3.GetBucketNotificationConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketNotificationConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketNotificationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketNotificationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketNotificationConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetBucketNotificationConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketOwnershipControls(arg0 context.Context, arg1 *s3.GetBucketOwnershipControlsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketOwnershipControls, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketOwnershipControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketOwnershipControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketOwnershipControls, reflect.TypeOf((*MockS3Client)(nil).GetBucketOwnershipControls), varargs...)
}

func (m *MockS3Client) GetBucketPolicy(arg0 context.Context, arg1 *s3.GetBucketPolicyInput, arg2 ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketPolicy, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketPolicy, reflect.TypeOf((*MockS3Client)(nil).GetBucketPolicy), varargs...)
}

func (m *MockS3Client) GetBucketPolicyStatus(arg0 context.Context, arg1 *s3.GetBucketPolicyStatusInput, arg2 ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketPolicyStatus, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketPolicyStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketPolicyStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketPolicyStatus, reflect.TypeOf((*MockS3Client)(nil).GetBucketPolicyStatus), varargs...)
}

func (m *MockS3Client) GetBucketReplication(arg0 context.Context, arg1 *s3.GetBucketReplicationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketReplication, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketReplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketReplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketReplication, reflect.TypeOf((*MockS3Client)(nil).GetBucketReplication), varargs...)
}

func (m *MockS3Client) GetBucketRequestPayment(arg0 context.Context, arg1 *s3.GetBucketRequestPaymentInput, arg2 ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketRequestPayment, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketRequestPaymentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketRequestPayment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketRequestPayment, reflect.TypeOf((*MockS3Client)(nil).GetBucketRequestPayment), varargs...)
}

func (m *MockS3Client) GetBucketTagging(arg0 context.Context, arg1 *s3.GetBucketTaggingInput, arg2 ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketTagging, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketTagging, reflect.TypeOf((*MockS3Client)(nil).GetBucketTagging), varargs...)
}

func (m *MockS3Client) GetBucketVersioning(arg0 context.Context, arg1 *s3.GetBucketVersioningInput, arg2 ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketVersioning, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketVersioningOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketVersioning(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketVersioning, reflect.TypeOf((*MockS3Client)(nil).GetBucketVersioning), varargs...)
}

func (m *MockS3Client) GetBucketWebsite(arg0 context.Context, arg1 *s3.GetBucketWebsiteInput, arg2 ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketWebsite, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketWebsiteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketWebsite(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketWebsite, reflect.TypeOf((*MockS3Client)(nil).GetBucketWebsite), varargs...)
}

func (m *MockS3Client) GetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObject, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObject, reflect.TypeOf((*MockS3Client)(nil).GetObject), varargs...)
}

func (m *MockS3Client) GetObjectAcl(arg0 context.Context, arg1 *s3.GetObjectAclInput, arg2 ...func(*s3.Options)) (*s3.GetObjectAclOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectAcl, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectAclOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectAcl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectAcl, reflect.TypeOf((*MockS3Client)(nil).GetObjectAcl), varargs...)
}

func (m *MockS3Client) GetObjectAttributes(arg0 context.Context, arg1 *s3.GetObjectAttributesInput, arg2 ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectAttributes, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectAttributes, reflect.TypeOf((*MockS3Client)(nil).GetObjectAttributes), varargs...)
}

func (m *MockS3Client) GetObjectLegalHold(arg0 context.Context, arg1 *s3.GetObjectLegalHoldInput, arg2 ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectLegalHold, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectLegalHoldOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectLegalHold(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectLegalHold, reflect.TypeOf((*MockS3Client)(nil).GetObjectLegalHold), varargs...)
}

func (m *MockS3Client) GetObjectLockConfiguration(arg0 context.Context, arg1 *s3.GetObjectLockConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectLockConfiguration, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectLockConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectLockConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectLockConfiguration, reflect.TypeOf((*MockS3Client)(nil).GetObjectLockConfiguration), varargs...)
}

func (m *MockS3Client) GetObjectRetention(arg0 context.Context, arg1 *s3.GetObjectRetentionInput, arg2 ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectRetention, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectRetentionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectRetention(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectRetention, reflect.TypeOf((*MockS3Client)(nil).GetObjectRetention), varargs...)
}

func (m *MockS3Client) GetObjectTagging(arg0 context.Context, arg1 *s3.GetObjectTaggingInput, arg2 ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectTagging, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectTagging, reflect.TypeOf((*MockS3Client)(nil).GetObjectTagging), varargs...)
}

func (m *MockS3Client) GetObjectTorrent(arg0 context.Context, arg1 *s3.GetObjectTorrentInput, arg2 ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetObjectTorrent, varargs...)
	ret0, _ := ret[0].(*s3.GetObjectTorrentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetObjectTorrent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetObjectTorrent, reflect.TypeOf((*MockS3Client)(nil).GetObjectTorrent), varargs...)
}

func (m *MockS3Client) GetPublicAccessBlock(arg0 context.Context, arg1 *s3.GetPublicAccessBlockInput, arg2 ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPublicAccessBlock, varargs...)
	ret0, _ := ret[0].(*s3.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetPublicAccessBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPublicAccessBlock, reflect.TypeOf((*MockS3Client)(nil).GetPublicAccessBlock), varargs...)
}

func (m *MockS3Client) ListBucketAnalyticsConfigurations(arg0 context.Context, arg1 *s3.ListBucketAnalyticsConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBucketAnalyticsConfigurations, varargs...)
	ret0, _ := ret[0].(*s3.ListBucketAnalyticsConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListBucketAnalyticsConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBucketAnalyticsConfigurations, reflect.TypeOf((*MockS3Client)(nil).ListBucketAnalyticsConfigurations), varargs...)
}

func (m *MockS3Client) ListBucketIntelligentTieringConfigurations(arg0 context.Context, arg1 *s3.ListBucketIntelligentTieringConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBucketIntelligentTieringConfigurations, varargs...)
	ret0, _ := ret[0].(*s3.ListBucketIntelligentTieringConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListBucketIntelligentTieringConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBucketIntelligentTieringConfigurations, reflect.TypeOf((*MockS3Client)(nil).ListBucketIntelligentTieringConfigurations), varargs...)
}

func (m *MockS3Client) ListBucketInventoryConfigurations(arg0 context.Context, arg1 *s3.ListBucketInventoryConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBucketInventoryConfigurations, varargs...)
	ret0, _ := ret[0].(*s3.ListBucketInventoryConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListBucketInventoryConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBucketInventoryConfigurations, reflect.TypeOf((*MockS3Client)(nil).ListBucketInventoryConfigurations), varargs...)
}

func (m *MockS3Client) ListBucketMetricsConfigurations(arg0 context.Context, arg1 *s3.ListBucketMetricsConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBucketMetricsConfigurations, varargs...)
	ret0, _ := ret[0].(*s3.ListBucketMetricsConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListBucketMetricsConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBucketMetricsConfigurations, reflect.TypeOf((*MockS3Client)(nil).ListBucketMetricsConfigurations), varargs...)
}

func (m *MockS3Client) ListBuckets(arg0 context.Context, arg1 *s3.ListBucketsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBuckets, varargs...)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBuckets, reflect.TypeOf((*MockS3Client)(nil).ListBuckets), varargs...)
}

func (m *MockS3Client) ListMultipartUploads(arg0 context.Context, arg1 *s3.ListMultipartUploadsInput, arg2 ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMultipartUploads, varargs...)
	ret0, _ := ret[0].(*s3.ListMultipartUploadsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListMultipartUploads(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMultipartUploads, reflect.TypeOf((*MockS3Client)(nil).ListMultipartUploads), varargs...)
}

func (m *MockS3Client) ListObjectVersions(arg0 context.Context, arg1 *s3.ListObjectVersionsInput, arg2 ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListObjectVersions, varargs...)
	ret0, _ := ret[0].(*s3.ListObjectVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListObjectVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListObjectVersions, reflect.TypeOf((*MockS3Client)(nil).ListObjectVersions), varargs...)
}

func (m *MockS3Client) ListObjects(arg0 context.Context, arg1 *s3.ListObjectsInput, arg2 ...func(*s3.Options)) (*s3.ListObjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListObjects, varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListObjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListObjects, reflect.TypeOf((*MockS3Client)(nil).ListObjects), varargs...)
}

func (m *MockS3Client) ListObjectsV2(arg0 context.Context, arg1 *s3.ListObjectsV2Input, arg2 ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListObjectsV, varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListObjectsV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListObjectsV, reflect.TypeOf((*MockS3Client)(nil).ListObjectsV2), varargs...)
}

func (m *MockS3Client) ListParts(arg0 context.Context, arg1 *s3.ListPartsInput, arg2 ...func(*s3.Options)) (*s3.ListPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListParts, varargs...)
	ret0, _ := ret[0].(*s3.ListPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListParts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListParts, reflect.TypeOf((*MockS3Client)(nil).ListParts), varargs...)
}
