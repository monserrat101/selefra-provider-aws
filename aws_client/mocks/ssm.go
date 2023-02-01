package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go-v2/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

type MockSsmClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSsmClientMockRecorder
}

type MockSsmClientMockRecorder struct {
	mock *MockSsmClient
}

func NewMockSsmClient(ctrl *gomock.Controller) *MockSsmClient {
	mock := &MockSsmClient{ctrl: ctrl}
	mock.recorder = &MockSsmClientMockRecorder{mock}
	return mock
}

func (m *MockSsmClient) EXPECT() *MockSsmClientMockRecorder {
	return m.recorder
}

func (m *MockSsmClient) DescribeActivations(arg0 context.Context, arg1 *ssm.DescribeActivationsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeActivationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeActivations, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeActivationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeActivations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeActivations, reflect.TypeOf((*MockSsmClient)(nil).DescribeActivations), varargs...)
}

func (m *MockSsmClient) DescribeAssociation(arg0 context.Context, arg1 *ssm.DescribeAssociationInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeAssociationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAssociation, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAssociation, reflect.TypeOf((*MockSsmClient)(nil).DescribeAssociation), varargs...)
}

func (m *MockSsmClient) DescribeAssociationExecutionTargets(arg0 context.Context, arg1 *ssm.DescribeAssociationExecutionTargetsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAssociationExecutionTargets, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeAssociationExecutionTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeAssociationExecutionTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAssociationExecutionTargets, reflect.TypeOf((*MockSsmClient)(nil).DescribeAssociationExecutionTargets), varargs...)
}

func (m *MockSsmClient) DescribeAssociationExecutions(arg0 context.Context, arg1 *ssm.DescribeAssociationExecutionsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeAssociationExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAssociationExecutions, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeAssociationExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeAssociationExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAssociationExecutions, reflect.TypeOf((*MockSsmClient)(nil).DescribeAssociationExecutions), varargs...)
}

func (m *MockSsmClient) DescribeAutomationExecutions(arg0 context.Context, arg1 *ssm.DescribeAutomationExecutionsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeAutomationExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutomationExecutions, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeAutomationExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeAutomationExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutomationExecutions, reflect.TypeOf((*MockSsmClient)(nil).DescribeAutomationExecutions), varargs...)
}

func (m *MockSsmClient) DescribeAutomationStepExecutions(arg0 context.Context, arg1 *ssm.DescribeAutomationStepExecutionsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAutomationStepExecutions, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeAutomationStepExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeAutomationStepExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAutomationStepExecutions, reflect.TypeOf((*MockSsmClient)(nil).DescribeAutomationStepExecutions), varargs...)
}

func (m *MockSsmClient) DescribeAvailablePatches(arg0 context.Context, arg1 *ssm.DescribeAvailablePatchesInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeAvailablePatchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeAvailablePatches, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeAvailablePatchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeAvailablePatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeAvailablePatches, reflect.TypeOf((*MockSsmClient)(nil).DescribeAvailablePatches), varargs...)
}

func (m *MockSsmClient) DescribeDocument(arg0 context.Context, arg1 *ssm.DescribeDocumentInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeDocumentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDocument, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeDocumentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeDocument(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDocument, reflect.TypeOf((*MockSsmClient)(nil).DescribeDocument), varargs...)
}

func (m *MockSsmClient) DescribeDocumentPermission(arg0 context.Context, arg1 *ssm.DescribeDocumentPermissionInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeDocumentPermissionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDocumentPermission, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeDocumentPermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeDocumentPermission(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDocumentPermission, reflect.TypeOf((*MockSsmClient)(nil).DescribeDocumentPermission), varargs...)
}

func (m *MockSsmClient) DescribeEffectiveInstanceAssociations(arg0 context.Context, arg1 *ssm.DescribeEffectiveInstanceAssociationsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEffectiveInstanceAssociations, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeEffectiveInstanceAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeEffectiveInstanceAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEffectiveInstanceAssociations, reflect.TypeOf((*MockSsmClient)(nil).DescribeEffectiveInstanceAssociations), varargs...)
}

func (m *MockSsmClient) DescribeEffectivePatchesForPatchBaseline(arg0 context.Context, arg1 *ssm.DescribeEffectivePatchesForPatchBaselineInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeEffectivePatchesForPatchBaseline, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeEffectivePatchesForPatchBaselineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeEffectivePatchesForPatchBaseline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeEffectivePatchesForPatchBaseline, reflect.TypeOf((*MockSsmClient)(nil).DescribeEffectivePatchesForPatchBaseline), varargs...)
}

func (m *MockSsmClient) DescribeInstanceAssociationsStatus(arg0 context.Context, arg1 *ssm.DescribeInstanceAssociationsStatusInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceAssociationsStatus, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInstanceAssociationsStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeInstanceAssociationsStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceAssociationsStatus, reflect.TypeOf((*MockSsmClient)(nil).DescribeInstanceAssociationsStatus), varargs...)
}

func (m *MockSsmClient) DescribeInstanceInformation(arg0 context.Context, arg1 *ssm.DescribeInstanceInformationInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstanceInformation, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInstanceInformationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeInstanceInformation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstanceInformation, reflect.TypeOf((*MockSsmClient)(nil).DescribeInstanceInformation), varargs...)
}

func (m *MockSsmClient) DescribeInstancePatchStates(arg0 context.Context, arg1 *ssm.DescribeInstancePatchStatesInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInstancePatchStatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstancePatchStates, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInstancePatchStatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeInstancePatchStates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstancePatchStates, reflect.TypeOf((*MockSsmClient)(nil).DescribeInstancePatchStates), varargs...)
}

func (m *MockSsmClient) DescribeInstancePatchStatesForPatchGroup(arg0 context.Context, arg1 *ssm.DescribeInstancePatchStatesForPatchGroupInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstancePatchStatesForPatchGroup, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInstancePatchStatesForPatchGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeInstancePatchStatesForPatchGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstancePatchStatesForPatchGroup, reflect.TypeOf((*MockSsmClient)(nil).DescribeInstancePatchStatesForPatchGroup), varargs...)
}

func (m *MockSsmClient) DescribeInstancePatches(arg0 context.Context, arg1 *ssm.DescribeInstancePatchesInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInstancePatchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInstancePatches, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInstancePatchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeInstancePatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInstancePatches, reflect.TypeOf((*MockSsmClient)(nil).DescribeInstancePatches), varargs...)
}

func (m *MockSsmClient) DescribeInventoryDeletions(arg0 context.Context, arg1 *ssm.DescribeInventoryDeletionsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInventoryDeletionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeInventoryDeletions, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInventoryDeletionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeInventoryDeletions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeInventoryDeletions, reflect.TypeOf((*MockSsmClient)(nil).DescribeInventoryDeletions), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowExecutionTaskInvocations(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowExecutionTaskInvocations, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowExecutionTaskInvocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowExecutionTaskInvocations, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowExecutionTaskInvocations), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowExecutionTasks(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowExecutionTasksInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowExecutionTasks, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowExecutionTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowExecutionTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowExecutionTasks, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowExecutionTasks), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowExecutions(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowExecutionsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowExecutions, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowExecutions, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowExecutions), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowSchedule(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowScheduleInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowSchedule, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowSchedule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowSchedule, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowSchedule), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowTargets(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowTargetsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowTargets, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowTargets, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowTargets), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowTasks(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowTasksInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowTasks, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowTasks, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowTasks), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindows(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindows, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindows, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindows), varargs...)
}

func (m *MockSsmClient) DescribeMaintenanceWindowsForTarget(arg0 context.Context, arg1 *ssm.DescribeMaintenanceWindowsForTargetInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeMaintenanceWindowsForTarget, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeMaintenanceWindowsForTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeMaintenanceWindowsForTarget(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeMaintenanceWindowsForTarget, reflect.TypeOf((*MockSsmClient)(nil).DescribeMaintenanceWindowsForTarget), varargs...)
}

func (m *MockSsmClient) DescribeOpsItems(arg0 context.Context, arg1 *ssm.DescribeOpsItemsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeOpsItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeOpsItems, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeOpsItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeOpsItems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeOpsItems, reflect.TypeOf((*MockSsmClient)(nil).DescribeOpsItems), varargs...)
}

func (m *MockSsmClient) DescribeParameters(arg0 context.Context, arg1 *ssm.DescribeParametersInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeParameters, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeParameters, reflect.TypeOf((*MockSsmClient)(nil).DescribeParameters), varargs...)
}

func (m *MockSsmClient) DescribePatchBaselines(arg0 context.Context, arg1 *ssm.DescribePatchBaselinesInput, arg2 ...func(*ssm.Options)) (*ssm.DescribePatchBaselinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePatchBaselines, varargs...)
	ret0, _ := ret[0].(*ssm.DescribePatchBaselinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribePatchBaselines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePatchBaselines, reflect.TypeOf((*MockSsmClient)(nil).DescribePatchBaselines), varargs...)
}

func (m *MockSsmClient) DescribePatchGroupState(arg0 context.Context, arg1 *ssm.DescribePatchGroupStateInput, arg2 ...func(*ssm.Options)) (*ssm.DescribePatchGroupStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePatchGroupState, varargs...)
	ret0, _ := ret[0].(*ssm.DescribePatchGroupStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribePatchGroupState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePatchGroupState, reflect.TypeOf((*MockSsmClient)(nil).DescribePatchGroupState), varargs...)
}

func (m *MockSsmClient) DescribePatchGroups(arg0 context.Context, arg1 *ssm.DescribePatchGroupsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribePatchGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePatchGroups, varargs...)
	ret0, _ := ret[0].(*ssm.DescribePatchGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribePatchGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePatchGroups, reflect.TypeOf((*MockSsmClient)(nil).DescribePatchGroups), varargs...)
}

func (m *MockSsmClient) DescribePatchProperties(arg0 context.Context, arg1 *ssm.DescribePatchPropertiesInput, arg2 ...func(*ssm.Options)) (*ssm.DescribePatchPropertiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePatchProperties, varargs...)
	ret0, _ := ret[0].(*ssm.DescribePatchPropertiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribePatchProperties(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePatchProperties, reflect.TypeOf((*MockSsmClient)(nil).DescribePatchProperties), varargs...)
}

func (m *MockSsmClient) DescribeSessions(arg0 context.Context, arg1 *ssm.DescribeSessionsInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeSessions, varargs...)
	ret0, _ := ret[0].(*ssm.DescribeSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) DescribeSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeSessions, reflect.TypeOf((*MockSsmClient)(nil).DescribeSessions), varargs...)
}

func (m *MockSsmClient) GetAutomationExecution(arg0 context.Context, arg1 *ssm.GetAutomationExecutionInput, arg2 ...func(*ssm.Options)) (*ssm.GetAutomationExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAutomationExecution, varargs...)
	ret0, _ := ret[0].(*ssm.GetAutomationExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetAutomationExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAutomationExecution, reflect.TypeOf((*MockSsmClient)(nil).GetAutomationExecution), varargs...)
}

func (m *MockSsmClient) GetCalendarState(arg0 context.Context, arg1 *ssm.GetCalendarStateInput, arg2 ...func(*ssm.Options)) (*ssm.GetCalendarStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCalendarState, varargs...)
	ret0, _ := ret[0].(*ssm.GetCalendarStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetCalendarState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCalendarState, reflect.TypeOf((*MockSsmClient)(nil).GetCalendarState), varargs...)
}

func (m *MockSsmClient) GetCommandInvocation(arg0 context.Context, arg1 *ssm.GetCommandInvocationInput, arg2 ...func(*ssm.Options)) (*ssm.GetCommandInvocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetCommandInvocation, varargs...)
	ret0, _ := ret[0].(*ssm.GetCommandInvocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetCommandInvocation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetCommandInvocation, reflect.TypeOf((*MockSsmClient)(nil).GetCommandInvocation), varargs...)
}

func (m *MockSsmClient) GetConnectionStatus(arg0 context.Context, arg1 *ssm.GetConnectionStatusInput, arg2 ...func(*ssm.Options)) (*ssm.GetConnectionStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetConnectionStatus, varargs...)
	ret0, _ := ret[0].(*ssm.GetConnectionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetConnectionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetConnectionStatus, reflect.TypeOf((*MockSsmClient)(nil).GetConnectionStatus), varargs...)
}

func (m *MockSsmClient) GetDefaultPatchBaseline(arg0 context.Context, arg1 *ssm.GetDefaultPatchBaselineInput, arg2 ...func(*ssm.Options)) (*ssm.GetDefaultPatchBaselineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDefaultPatchBaseline, varargs...)
	ret0, _ := ret[0].(*ssm.GetDefaultPatchBaselineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetDefaultPatchBaseline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDefaultPatchBaseline, reflect.TypeOf((*MockSsmClient)(nil).GetDefaultPatchBaseline), varargs...)
}

func (m *MockSsmClient) GetDeployablePatchSnapshotForInstance(arg0 context.Context, arg1 *ssm.GetDeployablePatchSnapshotForInstanceInput, arg2 ...func(*ssm.Options)) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDeployablePatchSnapshotForInstance, varargs...)
	ret0, _ := ret[0].(*ssm.GetDeployablePatchSnapshotForInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetDeployablePatchSnapshotForInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDeployablePatchSnapshotForInstance, reflect.TypeOf((*MockSsmClient)(nil).GetDeployablePatchSnapshotForInstance), varargs...)
}

func (m *MockSsmClient) GetDocument(arg0 context.Context, arg1 *ssm.GetDocumentInput, arg2 ...func(*ssm.Options)) (*ssm.GetDocumentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetDocument, varargs...)
	ret0, _ := ret[0].(*ssm.GetDocumentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetDocument(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetDocument, reflect.TypeOf((*MockSsmClient)(nil).GetDocument), varargs...)
}

func (m *MockSsmClient) GetInventory(arg0 context.Context, arg1 *ssm.GetInventoryInput, arg2 ...func(*ssm.Options)) (*ssm.GetInventoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInventory, varargs...)
	ret0, _ := ret[0].(*ssm.GetInventoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetInventory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInventory, reflect.TypeOf((*MockSsmClient)(nil).GetInventory), varargs...)
}

func (m *MockSsmClient) GetInventorySchema(arg0 context.Context, arg1 *ssm.GetInventorySchemaInput, arg2 ...func(*ssm.Options)) (*ssm.GetInventorySchemaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetInventorySchema, varargs...)
	ret0, _ := ret[0].(*ssm.GetInventorySchemaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetInventorySchema(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetInventorySchema, reflect.TypeOf((*MockSsmClient)(nil).GetInventorySchema), varargs...)
}

func (m *MockSsmClient) GetMaintenanceWindow(arg0 context.Context, arg1 *ssm.GetMaintenanceWindowInput, arg2 ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMaintenanceWindow, varargs...)
	ret0, _ := ret[0].(*ssm.GetMaintenanceWindowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetMaintenanceWindow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMaintenanceWindow, reflect.TypeOf((*MockSsmClient)(nil).GetMaintenanceWindow), varargs...)
}

func (m *MockSsmClient) GetMaintenanceWindowExecution(arg0 context.Context, arg1 *ssm.GetMaintenanceWindowExecutionInput, arg2 ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMaintenanceWindowExecution, varargs...)
	ret0, _ := ret[0].(*ssm.GetMaintenanceWindowExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetMaintenanceWindowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMaintenanceWindowExecution, reflect.TypeOf((*MockSsmClient)(nil).GetMaintenanceWindowExecution), varargs...)
}

func (m *MockSsmClient) GetMaintenanceWindowExecutionTask(arg0 context.Context, arg1 *ssm.GetMaintenanceWindowExecutionTaskInput, arg2 ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMaintenanceWindowExecutionTask, varargs...)
	ret0, _ := ret[0].(*ssm.GetMaintenanceWindowExecutionTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetMaintenanceWindowExecutionTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMaintenanceWindowExecutionTask, reflect.TypeOf((*MockSsmClient)(nil).GetMaintenanceWindowExecutionTask), varargs...)
}

func (m *MockSsmClient) GetMaintenanceWindowExecutionTaskInvocation(arg0 context.Context, arg1 *ssm.GetMaintenanceWindowExecutionTaskInvocationInput, arg2 ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMaintenanceWindowExecutionTaskInvocation, varargs...)
	ret0, _ := ret[0].(*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetMaintenanceWindowExecutionTaskInvocation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMaintenanceWindowExecutionTaskInvocation, reflect.TypeOf((*MockSsmClient)(nil).GetMaintenanceWindowExecutionTaskInvocation), varargs...)
}

func (m *MockSsmClient) GetMaintenanceWindowTask(arg0 context.Context, arg1 *ssm.GetMaintenanceWindowTaskInput, arg2 ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetMaintenanceWindowTask, varargs...)
	ret0, _ := ret[0].(*ssm.GetMaintenanceWindowTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetMaintenanceWindowTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetMaintenanceWindowTask, reflect.TypeOf((*MockSsmClient)(nil).GetMaintenanceWindowTask), varargs...)
}

func (m *MockSsmClient) GetOpsItem(arg0 context.Context, arg1 *ssm.GetOpsItemInput, arg2 ...func(*ssm.Options)) (*ssm.GetOpsItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOpsItem, varargs...)
	ret0, _ := ret[0].(*ssm.GetOpsItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetOpsItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOpsItem, reflect.TypeOf((*MockSsmClient)(nil).GetOpsItem), varargs...)
}

func (m *MockSsmClient) GetOpsMetadata(arg0 context.Context, arg1 *ssm.GetOpsMetadataInput, arg2 ...func(*ssm.Options)) (*ssm.GetOpsMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOpsMetadata, varargs...)
	ret0, _ := ret[0].(*ssm.GetOpsMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetOpsMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOpsMetadata, reflect.TypeOf((*MockSsmClient)(nil).GetOpsMetadata), varargs...)
}

func (m *MockSsmClient) GetOpsSummary(arg0 context.Context, arg1 *ssm.GetOpsSummaryInput, arg2 ...func(*ssm.Options)) (*ssm.GetOpsSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetOpsSummary, varargs...)
	ret0, _ := ret[0].(*ssm.GetOpsSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetOpsSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOpsSummary, reflect.TypeOf((*MockSsmClient)(nil).GetOpsSummary), varargs...)
}

func (m *MockSsmClient) GetParameter(arg0 context.Context, arg1 *ssm.GetParameterInput, arg2 ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetParameter, varargs...)
	ret0, _ := ret[0].(*ssm.GetParameterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetParameter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetParameter, reflect.TypeOf((*MockSsmClient)(nil).GetParameter), varargs...)
}

func (m *MockSsmClient) GetParameterHistory(arg0 context.Context, arg1 *ssm.GetParameterHistoryInput, arg2 ...func(*ssm.Options)) (*ssm.GetParameterHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetParameterHistory, varargs...)
	ret0, _ := ret[0].(*ssm.GetParameterHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetParameterHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetParameterHistory, reflect.TypeOf((*MockSsmClient)(nil).GetParameterHistory), varargs...)
}

func (m *MockSsmClient) GetParameters(arg0 context.Context, arg1 *ssm.GetParametersInput, arg2 ...func(*ssm.Options)) (*ssm.GetParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetParameters, varargs...)
	ret0, _ := ret[0].(*ssm.GetParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetParameters, reflect.TypeOf((*MockSsmClient)(nil).GetParameters), varargs...)
}

func (m *MockSsmClient) GetParametersByPath(arg0 context.Context, arg1 *ssm.GetParametersByPathInput, arg2 ...func(*ssm.Options)) (*ssm.GetParametersByPathOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetParametersByPath, varargs...)
	ret0, _ := ret[0].(*ssm.GetParametersByPathOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetParametersByPath(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetParametersByPath, reflect.TypeOf((*MockSsmClient)(nil).GetParametersByPath), varargs...)
}

func (m *MockSsmClient) GetPatchBaseline(arg0 context.Context, arg1 *ssm.GetPatchBaselineInput, arg2 ...func(*ssm.Options)) (*ssm.GetPatchBaselineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPatchBaseline, varargs...)
	ret0, _ := ret[0].(*ssm.GetPatchBaselineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetPatchBaseline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPatchBaseline, reflect.TypeOf((*MockSsmClient)(nil).GetPatchBaseline), varargs...)
}

func (m *MockSsmClient) GetPatchBaselineForPatchGroup(arg0 context.Context, arg1 *ssm.GetPatchBaselineForPatchGroupInput, arg2 ...func(*ssm.Options)) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetPatchBaselineForPatchGroup, varargs...)
	ret0, _ := ret[0].(*ssm.GetPatchBaselineForPatchGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetPatchBaselineForPatchGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPatchBaselineForPatchGroup, reflect.TypeOf((*MockSsmClient)(nil).GetPatchBaselineForPatchGroup), varargs...)
}

func (m *MockSsmClient) GetResourcePolicies(arg0 context.Context, arg1 *ssm.GetResourcePoliciesInput, arg2 ...func(*ssm.Options)) (*ssm.GetResourcePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetResourcePolicies, varargs...)
	ret0, _ := ret[0].(*ssm.GetResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetResourcePolicies, reflect.TypeOf((*MockSsmClient)(nil).GetResourcePolicies), varargs...)
}

func (m *MockSsmClient) GetServiceSetting(arg0 context.Context, arg1 *ssm.GetServiceSettingInput, arg2 ...func(*ssm.Options)) (*ssm.GetServiceSettingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetServiceSetting, varargs...)
	ret0, _ := ret[0].(*ssm.GetServiceSettingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) GetServiceSetting(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetServiceSetting, reflect.TypeOf((*MockSsmClient)(nil).GetServiceSetting), varargs...)
}

func (m *MockSsmClient) ListAssociationVersions(arg0 context.Context, arg1 *ssm.ListAssociationVersionsInput, arg2 ...func(*ssm.Options)) (*ssm.ListAssociationVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociationVersions, varargs...)
	ret0, _ := ret[0].(*ssm.ListAssociationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListAssociationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociationVersions, reflect.TypeOf((*MockSsmClient)(nil).ListAssociationVersions), varargs...)
}

func (m *MockSsmClient) ListAssociations(arg0 context.Context, arg1 *ssm.ListAssociationsInput, arg2 ...func(*ssm.Options)) (*ssm.ListAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAssociations, varargs...)
	ret0, _ := ret[0].(*ssm.ListAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAssociations, reflect.TypeOf((*MockSsmClient)(nil).ListAssociations), varargs...)
}

func (m *MockSsmClient) ListCommandInvocations(arg0 context.Context, arg1 *ssm.ListCommandInvocationsInput, arg2 ...func(*ssm.Options)) (*ssm.ListCommandInvocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCommandInvocations, varargs...)
	ret0, _ := ret[0].(*ssm.ListCommandInvocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListCommandInvocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCommandInvocations, reflect.TypeOf((*MockSsmClient)(nil).ListCommandInvocations), varargs...)
}

func (m *MockSsmClient) ListCommands(arg0 context.Context, arg1 *ssm.ListCommandsInput, arg2 ...func(*ssm.Options)) (*ssm.ListCommandsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListCommands, varargs...)
	ret0, _ := ret[0].(*ssm.ListCommandsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListCommands(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListCommands, reflect.TypeOf((*MockSsmClient)(nil).ListCommands), varargs...)
}

func (m *MockSsmClient) ListComplianceItems(arg0 context.Context, arg1 *ssm.ListComplianceItemsInput, arg2 ...func(*ssm.Options)) (*ssm.ListComplianceItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListComplianceItems, varargs...)
	ret0, _ := ret[0].(*ssm.ListComplianceItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListComplianceItems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListComplianceItems, reflect.TypeOf((*MockSsmClient)(nil).ListComplianceItems), varargs...)
}

func (m *MockSsmClient) ListComplianceSummaries(arg0 context.Context, arg1 *ssm.ListComplianceSummariesInput, arg2 ...func(*ssm.Options)) (*ssm.ListComplianceSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListComplianceSummaries, varargs...)
	ret0, _ := ret[0].(*ssm.ListComplianceSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListComplianceSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListComplianceSummaries, reflect.TypeOf((*MockSsmClient)(nil).ListComplianceSummaries), varargs...)
}

func (m *MockSsmClient) ListDocumentMetadataHistory(arg0 context.Context, arg1 *ssm.ListDocumentMetadataHistoryInput, arg2 ...func(*ssm.Options)) (*ssm.ListDocumentMetadataHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDocumentMetadataHistory, varargs...)
	ret0, _ := ret[0].(*ssm.ListDocumentMetadataHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListDocumentMetadataHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDocumentMetadataHistory, reflect.TypeOf((*MockSsmClient)(nil).ListDocumentMetadataHistory), varargs...)
}

func (m *MockSsmClient) ListDocumentVersions(arg0 context.Context, arg1 *ssm.ListDocumentVersionsInput, arg2 ...func(*ssm.Options)) (*ssm.ListDocumentVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDocumentVersions, varargs...)
	ret0, _ := ret[0].(*ssm.ListDocumentVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListDocumentVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDocumentVersions, reflect.TypeOf((*MockSsmClient)(nil).ListDocumentVersions), varargs...)
}

func (m *MockSsmClient) ListDocuments(arg0 context.Context, arg1 *ssm.ListDocumentsInput, arg2 ...func(*ssm.Options)) (*ssm.ListDocumentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDocuments, varargs...)
	ret0, _ := ret[0].(*ssm.ListDocumentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListDocuments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDocuments, reflect.TypeOf((*MockSsmClient)(nil).ListDocuments), varargs...)
}

func (m *MockSsmClient) ListInventoryEntries(arg0 context.Context, arg1 *ssm.ListInventoryEntriesInput, arg2 ...func(*ssm.Options)) (*ssm.ListInventoryEntriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListInventoryEntries, varargs...)
	ret0, _ := ret[0].(*ssm.ListInventoryEntriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListInventoryEntries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInventoryEntries, reflect.TypeOf((*MockSsmClient)(nil).ListInventoryEntries), varargs...)
}

func (m *MockSsmClient) ListOpsItemEvents(arg0 context.Context, arg1 *ssm.ListOpsItemEventsInput, arg2 ...func(*ssm.Options)) (*ssm.ListOpsItemEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOpsItemEvents, varargs...)
	ret0, _ := ret[0].(*ssm.ListOpsItemEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListOpsItemEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOpsItemEvents, reflect.TypeOf((*MockSsmClient)(nil).ListOpsItemEvents), varargs...)
}

func (m *MockSsmClient) ListOpsItemRelatedItems(arg0 context.Context, arg1 *ssm.ListOpsItemRelatedItemsInput, arg2 ...func(*ssm.Options)) (*ssm.ListOpsItemRelatedItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOpsItemRelatedItems, varargs...)
	ret0, _ := ret[0].(*ssm.ListOpsItemRelatedItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListOpsItemRelatedItems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOpsItemRelatedItems, reflect.TypeOf((*MockSsmClient)(nil).ListOpsItemRelatedItems), varargs...)
}

func (m *MockSsmClient) ListOpsMetadata(arg0 context.Context, arg1 *ssm.ListOpsMetadataInput, arg2 ...func(*ssm.Options)) (*ssm.ListOpsMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOpsMetadata, varargs...)
	ret0, _ := ret[0].(*ssm.ListOpsMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListOpsMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOpsMetadata, reflect.TypeOf((*MockSsmClient)(nil).ListOpsMetadata), varargs...)
}

func (m *MockSsmClient) ListResourceComplianceSummaries(arg0 context.Context, arg1 *ssm.ListResourceComplianceSummariesInput, arg2 ...func(*ssm.Options)) (*ssm.ListResourceComplianceSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceComplianceSummaries, varargs...)
	ret0, _ := ret[0].(*ssm.ListResourceComplianceSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListResourceComplianceSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceComplianceSummaries, reflect.TypeOf((*MockSsmClient)(nil).ListResourceComplianceSummaries), varargs...)
}

func (m *MockSsmClient) ListResourceDataSync(arg0 context.Context, arg1 *ssm.ListResourceDataSyncInput, arg2 ...func(*ssm.Options)) (*ssm.ListResourceDataSyncOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourceDataSync, varargs...)
	ret0, _ := ret[0].(*ssm.ListResourceDataSyncOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListResourceDataSync(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourceDataSync, reflect.TypeOf((*MockSsmClient)(nil).ListResourceDataSync), varargs...)
}

func (m *MockSsmClient) ListTagsForResource(arg0 context.Context, arg1 *ssm.ListTagsForResourceInput, arg2 ...func(*ssm.Options)) (*ssm.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForResource, varargs...)
	ret0, _ := ret[0].(*ssm.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSsmClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForResource, reflect.TypeOf((*MockSsmClient)(nil).ListTagsForResource), varargs...)
}
