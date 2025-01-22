// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/gosettings/reader (interfaces: Source)

// Package settings is a generated GoMock package.
package settings

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSource is a mock of Source interface.
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource.
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance.
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSource) Get(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSourceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSource)(nil).Get), arg0)
}

// KeyTransform mocks base method.
func (m *MockSource) KeyTransform(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyTransform", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// KeyTransform indicates an expected call of KeyTransform.
func (mr *MockSourceMockRecorder) KeyTransform(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyTransform", reflect.TypeOf((*MockSource)(nil).KeyTransform), arg0)
}

// String mocks base method.
func (m *MockSource) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockSourceMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockSource)(nil).String))
}
