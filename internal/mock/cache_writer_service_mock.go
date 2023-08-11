// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/chantapat.t/GolandProjects/mini-redis-go/internal/service/server/core/cache/cache_writer_service.go

// Package mock_cache is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockICacheWriter is a mock of ICacheWriter interface.
type MockICacheWriter struct {
	ctrl     *gomock.Controller
	recorder *MockICacheWriterMockRecorder
}

// MockICacheWriterMockRecorder is the mock recorder for MockICacheWriter.
type MockICacheWriterMockRecorder struct {
	mock *MockICacheWriter
}

// NewMockICacheWriter creates a new mock instance.
func NewMockICacheWriter(ctrl *gomock.Controller) *MockICacheWriter {
	mock := &MockICacheWriter{ctrl: ctrl}
	mock.recorder = &MockICacheWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICacheWriter) EXPECT() *MockICacheWriterMockRecorder {
	return m.recorder
}

// WriteToFile mocks base method.
func (m *MockICacheWriter) WriteToFile(cacheFolder, k string, v []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToFile", cacheFolder, k, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToFile indicates an expected call of WriteToFile.
func (mr *MockICacheWriterMockRecorder) WriteToFile(cacheFolder, k, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockICacheWriter)(nil).WriteToFile), cacheFolder, k, v)
}
