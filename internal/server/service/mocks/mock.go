// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dragun-igor/img-strg/internal/server/service (interfaces: Storage)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// GetBirthTimeFile mocks base method.
func (m *MockStorage) GetBirthTimeFile(arg0 string) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBirthTimeFile", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBirthTimeFile indicates an expected call of GetBirthTimeFile.
func (mr *MockStorageMockRecorder) GetBirthTimeFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBirthTimeFile", reflect.TypeOf((*MockStorage)(nil).GetBirthTimeFile), arg0)
}

// SetBirthTimeFile mocks base method.
func (m *MockStorage) SetBirthTimeFile(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBirthTimeFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBirthTimeFile indicates an expected call of SetBirthTimeFile.
func (mr *MockStorageMockRecorder) SetBirthTimeFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBirthTimeFile", reflect.TypeOf((*MockStorage)(nil).SetBirthTimeFile), arg0, arg1)
}
