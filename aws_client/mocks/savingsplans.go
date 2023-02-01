package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	savingsplans "github.com/aws/aws-sdk-go-v2/service/savingsplans"
	gomock "github.com/golang/mock/gomock"
)

type MockSavingsplansClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSavingsplansClientMockRecorder
}

type MockSavingsplansClientMockRecorder struct {
	mock *MockSavingsplansClient
}

func NewMockSavingsplansClient(ctrl *gomock.Controller) *MockSavingsplansClient {
	mock := &MockSavingsplansClient{ctrl: ctrl}
	mock.recorder = &MockSavingsplansClientMockRecorder{mock}
	return mock
}

func (m *MockSavingsplansClient) EXPECT() *MockSavingsplansClientMockRecorder {
	return m.recorder
}

func (m *MockSavingsplansClient) DescribeSavingsPlanRates(arg0 context.Context, arg1 *savingsplans.DescribeSavingsPlanRatesInput, arg2 ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSavingsPlanRates, varargs...)
	ret0, _ := ret[0].(*savingsplans.DescribeSavingsPlanRatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSavingsplansClientMockRecorder) DescribeSavingsPlanRates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSavingsPlanRates, reflect.TypeOf((*MockSavingsplansClient)(nil).DescribeSavingsPlanRates), varargs...)
}

func (m *MockSavingsplansClient) DescribeSavingsPlans(arg0 context.Context, arg1 *savingsplans.DescribeSavingsPlansInput, arg2 ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSavingsPlans, varargs...)
	ret0, _ := ret[0].(*savingsplans.DescribeSavingsPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSavingsplansClientMockRecorder) DescribeSavingsPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSavingsPlans, reflect.TypeOf((*MockSavingsplansClient)(nil).DescribeSavingsPlans), varargs...)
}

func (m *MockSavingsplansClient) DescribeSavingsPlansOfferingRates(arg0 context.Context, arg1 *savingsplans.DescribeSavingsPlansOfferingRatesInput, arg2 ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSavingsPlansOfferingRates, varargs...)
	ret0, _ := ret[0].(*savingsplans.DescribeSavingsPlansOfferingRatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSavingsplansClientMockRecorder) DescribeSavingsPlansOfferingRates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSavingsPlansOfferingRates, reflect.TypeOf((*MockSavingsplansClient)(nil).DescribeSavingsPlansOfferingRates), varargs...)
}

func (m *MockSavingsplansClient) DescribeSavingsPlansOfferings(arg0 context.Context, arg1 *savingsplans.DescribeSavingsPlansOfferingsInput, arg2 ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSavingsPlansOfferings, varargs...)
	ret0, _ := ret[0].(*savingsplans.DescribeSavingsPlansOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSavingsplansClientMockRecorder) DescribeSavingsPlansOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSavingsPlansOfferings, reflect.TypeOf((*MockSavingsplansClient)(nil).DescribeSavingsPlansOfferings), varargs...)
}

func (m *MockSavingsplansClient) ListTagsForResource(arg0 context.Context, arg1 *savingsplans.ListTagsForResourceInput, arg2 ...func(*savingsplans.Options)) (*savingsplans.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*savingsplans.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSavingsplansClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSavingsplansClient)(nil).ListTagsForResource), varargs...)
}
