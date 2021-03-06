// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/helm/plan.go

// Package mock_helm is a generated GoMock package.
package helm

import (
	gomock "github.com/golang/mock/gomock"
	run "github.com/pelotech/drone-helm3/internal/run"
	reflect "reflect"
)

// MockStep is a mock of Step interface
type MockStep struct {
	ctrl     *gomock.Controller
	recorder *MockStepMockRecorder
}

// MockStepMockRecorder is the mock recorder for MockStep
type MockStepMockRecorder struct {
	mock *MockStep
}

// NewMockStep creates a new mock instance
func NewMockStep(ctrl *gomock.Controller) *MockStep {
	mock := &MockStep{ctrl: ctrl}
	mock.recorder = &MockStepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStep) EXPECT() *MockStepMockRecorder {
	return m.recorder
}

// Prepare mocks base method
func (m *MockStep) Prepare(arg0 run.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare
func (mr *MockStepMockRecorder) Prepare(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockStep)(nil).Prepare), arg0)
}

// Execute mocks base method
func (m *MockStep) Execute(arg0 run.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockStepMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStep)(nil).Execute), arg0)
}
