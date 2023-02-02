package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	cognitoidentity "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	gomock "github.com/golang/mock/gomock"
)

type MockCognitoidentityClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCognitoidentityClientMockRecorder
}

type MockCognitoidentityClientMockRecorder struct {
	mock *MockCognitoidentityClient
}

func NewMockCognitoidentityClient(ctrl *gomock.Controller) *MockCognitoidentityClient {
	mock := &MockCognitoidentityClient{ctrl: ctrl}
	mock.recorder = &MockCognitoidentityClientMockRecorder{mock}
	return mock
}

func (m *MockCognitoidentityClient) EXPECT() *MockCognitoidentityClientMockRecorder {
	return m.recorder
}

func (m *MockCognitoidentityClient) DescribeIdentity(arg0 context.Context, arg1 *cognitoidentity.DescribeIdentityInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIdentity, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.DescribeIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) DescribeIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIdentity, reflect.TypeOf((*MockCognitoidentityClient)(nil).DescribeIdentity), varargs...)
}

func (m *MockCognitoidentityClient) DescribeIdentityPool(arg0 context.Context, arg1 *cognitoidentity.DescribeIdentityPoolInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeIdentityPool, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.DescribeIdentityPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) DescribeIdentityPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeIdentityPool, reflect.TypeOf((*MockCognitoidentityClient)(nil).DescribeIdentityPool), varargs...)
}

func (m *MockCognitoidentityClient) GetCredentialsForIdentity(arg0 context.Context, arg1 *cognitoidentity.GetCredentialsForIdentityInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCredentialsForIdentity, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetCredentialsForIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) GetCredentialsForIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCredentialsForIdentity, reflect.TypeOf((*MockCognitoidentityClient)(nil).GetCredentialsForIdentity), varargs...)
}

func (m *MockCognitoidentityClient) GetId(arg0 context.Context, arg1 *cognitoidentity.GetIdInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetId, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) GetId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetId, reflect.TypeOf((*MockCognitoidentityClient)(nil).GetId), varargs...)
}

func (m *MockCognitoidentityClient) GetIdentityPoolRoles(arg0 context.Context, arg1 *cognitoidentity.GetIdentityPoolRolesInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetIdentityPoolRoles, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetIdentityPoolRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) GetIdentityPoolRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetIdentityPoolRoles, reflect.TypeOf((*MockCognitoidentityClient)(nil).GetIdentityPoolRoles), varargs...)
}

func (m *MockCognitoidentityClient) GetOpenIdToken(arg0 context.Context, arg1 *cognitoidentity.GetOpenIdTokenInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOpenIdToken, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetOpenIdTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) GetOpenIdToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOpenIdToken, reflect.TypeOf((*MockCognitoidentityClient)(nil).GetOpenIdToken), varargs...)
}

func (m *MockCognitoidentityClient) GetOpenIdTokenForDeveloperIdentity(arg0 context.Context, arg1 *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOpenIdTokenForDeveloperIdentity, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) GetOpenIdTokenForDeveloperIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOpenIdTokenForDeveloperIdentity, reflect.TypeOf((*MockCognitoidentityClient)(nil).GetOpenIdTokenForDeveloperIdentity), varargs...)
}

func (m *MockCognitoidentityClient) GetPrincipalTagAttributeMap(arg0 context.Context, arg1 *cognitoidentity.GetPrincipalTagAttributeMapInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPrincipalTagAttributeMap, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetPrincipalTagAttributeMapOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) GetPrincipalTagAttributeMap(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPrincipalTagAttributeMap, reflect.TypeOf((*MockCognitoidentityClient)(nil).GetPrincipalTagAttributeMap), varargs...)
}

func (m *MockCognitoidentityClient) ListIdentities(arg0 context.Context, arg1 *cognitoidentity.ListIdentitiesInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIdentities, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.ListIdentitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) ListIdentities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIdentities, reflect.TypeOf((*MockCognitoidentityClient)(nil).ListIdentities), varargs...)
}

func (m *MockCognitoidentityClient) ListIdentityPools(arg0 context.Context, arg1 *cognitoidentity.ListIdentityPoolsInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListIdentityPools, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.ListIdentityPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) ListIdentityPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListIdentityPools, reflect.TypeOf((*MockCognitoidentityClient)(nil).ListIdentityPools), varargs...)
}

func (m *MockCognitoidentityClient) ListTagsForResource(arg0 context.Context, arg1 *cognitoidentity.ListTagsForResourceInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*cognitoidentity.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoidentityClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockCognitoidentityClient)(nil).ListTagsForResource), varargs...)
}
