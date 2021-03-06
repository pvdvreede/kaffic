// Code generated by MockGen. DO NOT EDIT.
// Source: store_manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockStoreManager is a mock of StoreManager interface
type MockStoreManager struct {
	ctrl     *gomock.Controller
	recorder *MockStoreManagerMockRecorder
}

// MockStoreManagerMockRecorder is the mock recorder for MockStoreManager
type MockStoreManagerMockRecorder struct {
	mock *MockStoreManager
}

// NewMockStoreManager creates a new mock instance
func NewMockStoreManager(ctrl *gomock.Controller) *MockStoreManager {
	mock := &MockStoreManager{ctrl: ctrl}
	mock.recorder = &MockStoreManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreManager) EXPECT() *MockStoreManagerMockRecorder {
	return m.recorder
}

// GetMessageWriter mocks base method
func (m *MockStoreManager) GetMessageWriter() (io.Writer, error) {
	ret := m.ctrl.Call(m, "GetMessageWriter")
	ret0, _ := ret[0].(io.Writer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageWriter indicates an expected call of GetMessageWriter
func (mr *MockStoreManagerMockRecorder) GetMessageWriter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageWriter", reflect.TypeOf((*MockStoreManager)(nil).GetMessageWriter))
}

// GetIndexWriter mocks base method
func (m *MockStoreManager) GetIndexWriter() (io.Writer, error) {
	ret := m.ctrl.Call(m, "GetIndexWriter")
	ret0, _ := ret[0].(io.Writer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexWriter indicates an expected call of GetIndexWriter
func (mr *MockStoreManagerMockRecorder) GetIndexWriter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexWriter", reflect.TypeOf((*MockStoreManager)(nil).GetIndexWriter))
}

// GetMetaDataReadWriter mocks base method
func (m *MockStoreManager) GetMetaDataReadWriter() (io.ReadWriter, error) {
	ret := m.ctrl.Call(m, "GetMetaDataReadWriter")
	ret0, _ := ret[0].(io.ReadWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaDataReadWriter indicates an expected call of GetMetaDataReadWriter
func (mr *MockStoreManagerMockRecorder) GetMetaDataReadWriter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaDataReadWriter", reflect.TypeOf((*MockStoreManager)(nil).GetMetaDataReadWriter))
}
