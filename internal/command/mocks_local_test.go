// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces_local.go

// Package command is a generated GoMock package.
package command

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockexecCmd is a mock of execCmd interface.
type MockexecCmd struct {
	ctrl     *gomock.Controller
	recorder *MockexecCmdMockRecorder
}

// MockexecCmdMockRecorder is the mock recorder for MockexecCmd.
type MockexecCmdMockRecorder struct {
	mock *MockexecCmd
}

// NewMockexecCmd creates a new mock instance.
func NewMockexecCmd(ctrl *gomock.Controller) *MockexecCmd {
	mock := &MockexecCmd{ctrl: ctrl}
	mock.recorder = &MockexecCmdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockexecCmd) EXPECT() *MockexecCmdMockRecorder {
	return m.recorder
}

// CombinedOutput mocks base method.
func (m *MockexecCmd) CombinedOutput() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CombinedOutput")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CombinedOutput indicates an expected call of CombinedOutput.
func (mr *MockexecCmdMockRecorder) CombinedOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CombinedOutput", reflect.TypeOf((*MockexecCmd)(nil).CombinedOutput))
}

// Start mocks base method.
func (m *MockexecCmd) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockexecCmdMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockexecCmd)(nil).Start))
}

// StderrPipe mocks base method.
func (m *MockexecCmd) StderrPipe() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StderrPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StderrPipe indicates an expected call of StderrPipe.
func (mr *MockexecCmdMockRecorder) StderrPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StderrPipe", reflect.TypeOf((*MockexecCmd)(nil).StderrPipe))
}

// StdoutPipe mocks base method.
func (m *MockexecCmd) StdoutPipe() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdoutPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StdoutPipe indicates an expected call of StdoutPipe.
func (mr *MockexecCmdMockRecorder) StdoutPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdoutPipe", reflect.TypeOf((*MockexecCmd)(nil).StdoutPipe))
}

// Wait mocks base method.
func (m *MockexecCmd) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockexecCmdMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockexecCmd)(nil).Wait))
}
