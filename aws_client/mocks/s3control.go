package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	s3control "github.com/aws/aws-sdk-go-v2/service/s3control"
	gomock "github.com/golang/mock/gomock"
)

type MockS3controlClient struct {
	ctrl		*gomock.Controller
	recorder	*MockS3controlClientMockRecorder
}

type MockS3controlClientMockRecorder struct {
	mock *MockS3controlClient
}

func NewMockS3controlClient(ctrl *gomock.Controller) *MockS3controlClient {
	mock := &MockS3controlClient{ctrl: ctrl}
	mock.recorder = &MockS3controlClientMockRecorder{mock}
	return mock
}

func (m *MockS3controlClient) EXPECT() *MockS3controlClientMockRecorder {
	return m.recorder
}

func (m *MockS3controlClient) DescribeJob(arg0 context.Context, arg1 *s3control.DescribeJobInput, arg2 ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJob, varargs...)
	ret0, _ := ret[0].(*s3control.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) DescribeJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJob, reflect.TypeOf((*MockS3controlClient)(nil).DescribeJob), varargs...)
}

func (m *MockS3controlClient) DescribeMultiRegionAccessPointOperation(arg0 context.Context, arg1 *s3control.DescribeMultiRegionAccessPointOperationInput, arg2 ...func(*s3control.Options)) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMultiRegionAccessPointOperation, varargs...)
	ret0, _ := ret[0].(*s3control.DescribeMultiRegionAccessPointOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) DescribeMultiRegionAccessPointOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMultiRegionAccessPointOperation, reflect.TypeOf((*MockS3controlClient)(nil).DescribeMultiRegionAccessPointOperation), varargs...)
}

func (m *MockS3controlClient) GetAccessPoint(arg0 context.Context, arg1 *s3control.GetAccessPointInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPoint, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPoint, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPoint), varargs...)
}

func (m *MockS3controlClient) GetAccessPointConfigurationForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointConfigurationForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPointConfigurationForObjectLambda, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointConfigurationForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPointConfigurationForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPointConfigurationForObjectLambda, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointConfigurationForObjectLambda), varargs...)
}

func (m *MockS3controlClient) GetAccessPointForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPointForObjectLambda, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPointForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPointForObjectLambda, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointForObjectLambda), varargs...)
}

func (m *MockS3controlClient) GetAccessPointPolicy(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPointPolicy, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPointPolicy, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicy), varargs...)
}

func (m *MockS3controlClient) GetAccessPointPolicyForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPointPolicyForObjectLambda, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicyForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPointPolicyForObjectLambda, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicyForObjectLambda), varargs...)
}

func (m *MockS3controlClient) GetAccessPointPolicyStatus(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyStatusInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPointPolicyStatus, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicyStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPointPolicyStatus, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicyStatus), varargs...)
}

func (m *MockS3controlClient) GetAccessPointPolicyStatusForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAccessPointPolicyStatusForObjectLambda, varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicyStatusForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAccessPointPolicyStatusForObjectLambda, reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicyStatusForObjectLambda), varargs...)
}

func (m *MockS3controlClient) GetBucket(arg0 context.Context, arg1 *s3control.GetBucketInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucket, varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetBucket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucket, reflect.TypeOf((*MockS3controlClient)(nil).GetBucket), varargs...)
}

func (m *MockS3controlClient) GetBucketLifecycleConfiguration(arg0 context.Context, arg1 *s3control.GetBucketLifecycleConfigurationInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketLifecycleConfiguration, varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetBucketLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketLifecycleConfiguration, reflect.TypeOf((*MockS3controlClient)(nil).GetBucketLifecycleConfiguration), varargs...)
}

func (m *MockS3controlClient) GetBucketPolicy(arg0 context.Context, arg1 *s3control.GetBucketPolicyInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketPolicy, varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetBucketPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketPolicy, reflect.TypeOf((*MockS3controlClient)(nil).GetBucketPolicy), varargs...)
}

func (m *MockS3controlClient) GetBucketTagging(arg0 context.Context, arg1 *s3control.GetBucketTaggingInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketTagging, varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetBucketTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketTagging, reflect.TypeOf((*MockS3controlClient)(nil).GetBucketTagging), varargs...)
}

func (m *MockS3controlClient) GetBucketVersioning(arg0 context.Context, arg1 *s3control.GetBucketVersioningInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketVersioningOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketVersioning, varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketVersioningOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetBucketVersioning(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketVersioning, reflect.TypeOf((*MockS3controlClient)(nil).GetBucketVersioning), varargs...)
}

func (m *MockS3controlClient) GetJobTagging(arg0 context.Context, arg1 *s3control.GetJobTaggingInput, arg2 ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetJobTagging, varargs...)
	ret0, _ := ret[0].(*s3control.GetJobTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetJobTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetJobTagging, reflect.TypeOf((*MockS3controlClient)(nil).GetJobTagging), varargs...)
}

func (m *MockS3controlClient) GetMultiRegionAccessPoint(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMultiRegionAccessPoint, varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMultiRegionAccessPoint, reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPoint), varargs...)
}

func (m *MockS3controlClient) GetMultiRegionAccessPointPolicy(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointPolicyInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMultiRegionAccessPointPolicy, varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPointPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMultiRegionAccessPointPolicy, reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPointPolicy), varargs...)
}

func (m *MockS3controlClient) GetMultiRegionAccessPointPolicyStatus(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointPolicyStatusInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMultiRegionAccessPointPolicyStatus, varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointPolicyStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPointPolicyStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMultiRegionAccessPointPolicyStatus, reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPointPolicyStatus), varargs...)
}

func (m *MockS3controlClient) GetMultiRegionAccessPointRoutes(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointRoutesInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMultiRegionAccessPointRoutes, varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPointRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMultiRegionAccessPointRoutes, reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPointRoutes), varargs...)
}

func (m *MockS3controlClient) GetPublicAccessBlock(arg0 context.Context, arg1 *s3control.GetPublicAccessBlockInput, arg2 ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPublicAccessBlock, varargs...)
	ret0, _ := ret[0].(*s3control.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetPublicAccessBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPublicAccessBlock, reflect.TypeOf((*MockS3controlClient)(nil).GetPublicAccessBlock), varargs...)
}

func (m *MockS3controlClient) GetStorageLensConfiguration(arg0 context.Context, arg1 *s3control.GetStorageLensConfigurationInput, arg2 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStorageLensConfiguration, varargs...)
	ret0, _ := ret[0].(*s3control.GetStorageLensConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetStorageLensConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStorageLensConfiguration, reflect.TypeOf((*MockS3controlClient)(nil).GetStorageLensConfiguration), varargs...)
}

func (m *MockS3controlClient) GetStorageLensConfigurationTagging(arg0 context.Context, arg1 *s3control.GetStorageLensConfigurationTaggingInput, arg2 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetStorageLensConfigurationTagging, varargs...)
	ret0, _ := ret[0].(*s3control.GetStorageLensConfigurationTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) GetStorageLensConfigurationTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStorageLensConfigurationTagging, reflect.TypeOf((*MockS3controlClient)(nil).GetStorageLensConfigurationTagging), varargs...)
}

func (m *MockS3controlClient) ListAccessPoints(arg0 context.Context, arg1 *s3control.ListAccessPointsInput, arg2 ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccessPoints, varargs...)
	ret0, _ := ret[0].(*s3control.ListAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) ListAccessPoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccessPoints, reflect.TypeOf((*MockS3controlClient)(nil).ListAccessPoints), varargs...)
}

func (m *MockS3controlClient) ListAccessPointsForObjectLambda(arg0 context.Context, arg1 *s3control.ListAccessPointsForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAccessPointsForObjectLambda, varargs...)
	ret0, _ := ret[0].(*s3control.ListAccessPointsForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) ListAccessPointsForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAccessPointsForObjectLambda, reflect.TypeOf((*MockS3controlClient)(nil).ListAccessPointsForObjectLambda), varargs...)
}

func (m *MockS3controlClient) ListJobs(arg0 context.Context, arg1 *s3control.ListJobsInput, arg2 ...func(*s3control.Options)) (*s3control.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJobs, varargs...)
	ret0, _ := ret[0].(*s3control.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJobs, reflect.TypeOf((*MockS3controlClient)(nil).ListJobs), varargs...)
}

func (m *MockS3controlClient) ListMultiRegionAccessPoints(arg0 context.Context, arg1 *s3control.ListMultiRegionAccessPointsInput, arg2 ...func(*s3control.Options)) (*s3control.ListMultiRegionAccessPointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListMultiRegionAccessPoints, varargs...)
	ret0, _ := ret[0].(*s3control.ListMultiRegionAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) ListMultiRegionAccessPoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMultiRegionAccessPoints, reflect.TypeOf((*MockS3controlClient)(nil).ListMultiRegionAccessPoints), varargs...)
}

func (m *MockS3controlClient) ListRegionalBuckets(arg0 context.Context, arg1 *s3control.ListRegionalBucketsInput, arg2 ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRegionalBuckets, varargs...)
	ret0, _ := ret[0].(*s3control.ListRegionalBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) ListRegionalBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRegionalBuckets, reflect.TypeOf((*MockS3controlClient)(nil).ListRegionalBuckets), varargs...)
}

func (m *MockS3controlClient) ListStorageLensConfigurations(arg0 context.Context, arg1 *s3control.ListStorageLensConfigurationsInput, arg2 ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStorageLensConfigurations, varargs...)
	ret0, _ := ret[0].(*s3control.ListStorageLensConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3controlClientMockRecorder) ListStorageLensConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStorageLensConfigurations, reflect.TypeOf((*MockS3controlClient)(nil).ListStorageLensConfigurations), varargs...)
}
