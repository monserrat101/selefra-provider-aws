package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	qldb "github.com/aws/aws-sdk-go-v2/service/qldb"
	gomock "github.com/golang/mock/gomock"
)

type MockQldbClient struct {
	ctrl		*gomock.Controller
	recorder	*MockQldbClientMockRecorder
}

type MockQldbClientMockRecorder struct {
	mock *MockQldbClient
}

func NewMockQldbClient(ctrl *gomock.Controller) *MockQldbClient {
	mock := &MockQldbClient{ctrl: ctrl}
	mock.recorder = &MockQldbClientMockRecorder{mock}
	return mock
}

func (m *MockQldbClient) EXPECT() *MockQldbClientMockRecorder {
	return m.recorder
}

func (m *MockQldbClient) DescribeJournalKinesisStream(arg0 context.Context, arg1 *qldb.DescribeJournalKinesisStreamInput, arg2 ...func(*qldb.Options)) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJournalKinesisStream, varargs...)
	ret0, _ := ret[0].(*qldb.DescribeJournalKinesisStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) DescribeJournalKinesisStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJournalKinesisStream, reflect.TypeOf((*MockQldbClient)(nil).DescribeJournalKinesisStream), varargs...)
}

func (m *MockQldbClient) DescribeJournalS3Export(arg0 context.Context, arg1 *qldb.DescribeJournalS3ExportInput, arg2 ...func(*qldb.Options)) (*qldb.DescribeJournalS3ExportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeJournalSExport, varargs...)
	ret0, _ := ret[0].(*qldb.DescribeJournalS3ExportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) DescribeJournalS3Export(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeJournalSExport, reflect.TypeOf((*MockQldbClient)(nil).DescribeJournalS3Export), varargs...)
}

func (m *MockQldbClient) DescribeLedger(arg0 context.Context, arg1 *qldb.DescribeLedgerInput, arg2 ...func(*qldb.Options)) (*qldb.DescribeLedgerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeLedger, varargs...)
	ret0, _ := ret[0].(*qldb.DescribeLedgerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) DescribeLedger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeLedger, reflect.TypeOf((*MockQldbClient)(nil).DescribeLedger), varargs...)
}

func (m *MockQldbClient) GetBlock(arg0 context.Context, arg1 *qldb.GetBlockInput, arg2 ...func(*qldb.Options)) (*qldb.GetBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBlock, varargs...)
	ret0, _ := ret[0].(*qldb.GetBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) GetBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBlock, reflect.TypeOf((*MockQldbClient)(nil).GetBlock), varargs...)
}

func (m *MockQldbClient) GetDigest(arg0 context.Context, arg1 *qldb.GetDigestInput, arg2 ...func(*qldb.Options)) (*qldb.GetDigestOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDigest, varargs...)
	ret0, _ := ret[0].(*qldb.GetDigestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) GetDigest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDigest, reflect.TypeOf((*MockQldbClient)(nil).GetDigest), varargs...)
}

func (m *MockQldbClient) GetRevision(arg0 context.Context, arg1 *qldb.GetRevisionInput, arg2 ...func(*qldb.Options)) (*qldb.GetRevisionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetRevision, varargs...)
	ret0, _ := ret[0].(*qldb.GetRevisionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) GetRevision(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetRevision, reflect.TypeOf((*MockQldbClient)(nil).GetRevision), varargs...)
}

func (m *MockQldbClient) ListJournalKinesisStreamsForLedger(arg0 context.Context, arg1 *qldb.ListJournalKinesisStreamsForLedgerInput, arg2 ...func(*qldb.Options)) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJournalKinesisStreamsForLedger, varargs...)
	ret0, _ := ret[0].(*qldb.ListJournalKinesisStreamsForLedgerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) ListJournalKinesisStreamsForLedger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJournalKinesisStreamsForLedger, reflect.TypeOf((*MockQldbClient)(nil).ListJournalKinesisStreamsForLedger), varargs...)
}

func (m *MockQldbClient) ListJournalS3Exports(arg0 context.Context, arg1 *qldb.ListJournalS3ExportsInput, arg2 ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJournalSExports, varargs...)
	ret0, _ := ret[0].(*qldb.ListJournalS3ExportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) ListJournalS3Exports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJournalSExports, reflect.TypeOf((*MockQldbClient)(nil).ListJournalS3Exports), varargs...)
}

func (m *MockQldbClient) ListJournalS3ExportsForLedger(arg0 context.Context, arg1 *qldb.ListJournalS3ExportsForLedgerInput, arg2 ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListJournalSExportsForLedger, varargs...)
	ret0, _ := ret[0].(*qldb.ListJournalS3ExportsForLedgerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) ListJournalS3ExportsForLedger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListJournalSExportsForLedger, reflect.TypeOf((*MockQldbClient)(nil).ListJournalS3ExportsForLedger), varargs...)
}

func (m *MockQldbClient) ListLedgers(arg0 context.Context, arg1 *qldb.ListLedgersInput, arg2 ...func(*qldb.Options)) (*qldb.ListLedgersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLedgers, varargs...)
	ret0, _ := ret[0].(*qldb.ListLedgersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) ListLedgers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLedgers, reflect.TypeOf((*MockQldbClient)(nil).ListLedgers), varargs...)
}

func (m *MockQldbClient) ListTagsForResource(arg0 context.Context, arg1 *qldb.ListTagsForResourceInput, arg2 ...func(*qldb.Options)) (*qldb.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*qldb.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQldbClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockQldbClient)(nil).ListTagsForResource), varargs...)
}
