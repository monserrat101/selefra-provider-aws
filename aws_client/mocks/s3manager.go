package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

type MockS3managerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockS3managerClientMockRecorder
}

type MockS3managerClientMockRecorder struct {
	mock *MockS3managerClient
}

func NewMockS3managerClient(ctrl *gomock.Controller) *MockS3managerClient {
	mock := &MockS3managerClient{ctrl: ctrl}
	mock.recorder = &MockS3managerClientMockRecorder{mock}
	return mock
}

func (m *MockS3managerClient) EXPECT() *MockS3managerClientMockRecorder {
	return m.recorder
}

func (m *MockS3managerClient) GetBucketRegion(arg0 context.Context, arg1 string, arg2 ...func(*s3.Options)) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketRegion, varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3managerClientMockRecorder) GetBucketRegion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketRegion, reflect.TypeOf((*MockS3managerClient)(nil).GetBucketRegion), varargs...)
}
