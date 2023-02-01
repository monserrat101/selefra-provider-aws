package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	sns "github.com/aws/aws-sdk-go-v2/service/sns"
	gomock "github.com/golang/mock/gomock"
)

type MockSnsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSnsClientMockRecorder
}

type MockSnsClientMockRecorder struct {
	mock *MockSnsClient
}

func NewMockSnsClient(ctrl *gomock.Controller) *MockSnsClient {
	mock := &MockSnsClient{ctrl: ctrl}
	mock.recorder = &MockSnsClientMockRecorder{mock}
	return mock
}

func (m *MockSnsClient) EXPECT() *MockSnsClientMockRecorder {
	return m.recorder
}

func (m *MockSnsClient) GetDataProtectionPolicy(arg0 context.Context, arg1 *sns.GetDataProtectionPolicyInput, arg2 ...func(*sns.Options)) (*sns.GetDataProtectionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDataProtectionPolicy, varargs...)
	ret0, _ := ret[0].(*sns.GetDataProtectionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetDataProtectionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDataProtectionPolicy, reflect.TypeOf((*MockSnsClient)(nil).GetDataProtectionPolicy), varargs...)
}

func (m *MockSnsClient) GetEndpointAttributes(arg0 context.Context, arg1 *sns.GetEndpointAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetEndpointAttributes, varargs...)
	ret0, _ := ret[0].(*sns.GetEndpointAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetEndpointAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetEndpointAttributes, reflect.TypeOf((*MockSnsClient)(nil).GetEndpointAttributes), varargs...)
}

func (m *MockSnsClient) GetPlatformApplicationAttributes(arg0 context.Context, arg1 *sns.GetPlatformApplicationAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPlatformApplicationAttributes, varargs...)
	ret0, _ := ret[0].(*sns.GetPlatformApplicationAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetPlatformApplicationAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPlatformApplicationAttributes, reflect.TypeOf((*MockSnsClient)(nil).GetPlatformApplicationAttributes), varargs...)
}

func (m *MockSnsClient) GetSMSAttributes(arg0 context.Context, arg1 *sns.GetSMSAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSMSAttributes, varargs...)
	ret0, _ := ret[0].(*sns.GetSMSAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetSMSAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSMSAttributes, reflect.TypeOf((*MockSnsClient)(nil).GetSMSAttributes), varargs...)
}

func (m *MockSnsClient) GetSMSSandboxAccountStatus(arg0 context.Context, arg1 *sns.GetSMSSandboxAccountStatusInput, arg2 ...func(*sns.Options)) (*sns.GetSMSSandboxAccountStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSMSSandboxAccountStatus, varargs...)
	ret0, _ := ret[0].(*sns.GetSMSSandboxAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetSMSSandboxAccountStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSMSSandboxAccountStatus, reflect.TypeOf((*MockSnsClient)(nil).GetSMSSandboxAccountStatus), varargs...)
}

func (m *MockSnsClient) GetSubscriptionAttributes(arg0 context.Context, arg1 *sns.GetSubscriptionAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetSubscriptionAttributes, varargs...)
	ret0, _ := ret[0].(*sns.GetSubscriptionAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetSubscriptionAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetSubscriptionAttributes, reflect.TypeOf((*MockSnsClient)(nil).GetSubscriptionAttributes), varargs...)
}

func (m *MockSnsClient) GetTopicAttributes(arg0 context.Context, arg1 *sns.GetTopicAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetTopicAttributes, varargs...)
	ret0, _ := ret[0].(*sns.GetTopicAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetTopicAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTopicAttributes, reflect.TypeOf((*MockSnsClient)(nil).GetTopicAttributes), varargs...)
}

func (m *MockSnsClient) ListEndpointsByPlatformApplication(arg0 context.Context, arg1 *sns.ListEndpointsByPlatformApplicationInput, arg2 ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListEndpointsByPlatformApplication, varargs...)
	ret0, _ := ret[0].(*sns.ListEndpointsByPlatformApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListEndpointsByPlatformApplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListEndpointsByPlatformApplication, reflect.TypeOf((*MockSnsClient)(nil).ListEndpointsByPlatformApplication), varargs...)
}

func (m *MockSnsClient) ListOriginationNumbers(arg0 context.Context, arg1 *sns.ListOriginationNumbersInput, arg2 ...func(*sns.Options)) (*sns.ListOriginationNumbersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOriginationNumbers, varargs...)
	ret0, _ := ret[0].(*sns.ListOriginationNumbersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListOriginationNumbers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOriginationNumbers, reflect.TypeOf((*MockSnsClient)(nil).ListOriginationNumbers), varargs...)
}

func (m *MockSnsClient) ListPhoneNumbersOptedOut(arg0 context.Context, arg1 *sns.ListPhoneNumbersOptedOutInput, arg2 ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPhoneNumbersOptedOut, varargs...)
	ret0, _ := ret[0].(*sns.ListPhoneNumbersOptedOutOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListPhoneNumbersOptedOut(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPhoneNumbersOptedOut, reflect.TypeOf((*MockSnsClient)(nil).ListPhoneNumbersOptedOut), varargs...)
}

func (m *MockSnsClient) ListPlatformApplications(arg0 context.Context, arg1 *sns.ListPlatformApplicationsInput, arg2 ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPlatformApplications, varargs...)
	ret0, _ := ret[0].(*sns.ListPlatformApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListPlatformApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPlatformApplications, reflect.TypeOf((*MockSnsClient)(nil).ListPlatformApplications), varargs...)
}

func (m *MockSnsClient) ListSMSSandboxPhoneNumbers(arg0 context.Context, arg1 *sns.ListSMSSandboxPhoneNumbersInput, arg2 ...func(*sns.Options)) (*sns.ListSMSSandboxPhoneNumbersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSMSSandboxPhoneNumbers, varargs...)
	ret0, _ := ret[0].(*sns.ListSMSSandboxPhoneNumbersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListSMSSandboxPhoneNumbers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSMSSandboxPhoneNumbers, reflect.TypeOf((*MockSnsClient)(nil).ListSMSSandboxPhoneNumbers), varargs...)
}

func (m *MockSnsClient) ListSubscriptions(arg0 context.Context, arg1 *sns.ListSubscriptionsInput, arg2 ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSubscriptions, varargs...)
	ret0, _ := ret[0].(*sns.ListSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSubscriptions, reflect.TypeOf((*MockSnsClient)(nil).ListSubscriptions), varargs...)
}

func (m *MockSnsClient) ListSubscriptionsByTopic(arg0 context.Context, arg1 *sns.ListSubscriptionsByTopicInput, arg2 ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListSubscriptionsByTopic, varargs...)
	ret0, _ := ret[0].(*sns.ListSubscriptionsByTopicOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListSubscriptionsByTopic(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListSubscriptionsByTopic, reflect.TypeOf((*MockSnsClient)(nil).ListSubscriptionsByTopic), varargs...)
}

func (m *MockSnsClient) ListTagsForResource(arg0 context.Context, arg1 *sns.ListTagsForResourceInput, arg2 ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*sns.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSnsClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockSnsClient) ListTopics(arg0 context.Context, arg1 *sns.ListTopicsInput, arg2 ...func(*sns.Options)) (*sns.ListTopicsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTopics, varargs...)
	ret0, _ := ret[0].(*sns.ListTopicsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListTopics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTopics, reflect.TypeOf((*MockSnsClient)(nil).ListTopics), varargs...)
}
