// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/chantapat.t/GolandProjects/mini-redis-go/internal/service/server/core/cache/cache_reader_service.go

// Package mock_cache is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockICacheReader is a mock of ICacheReader interface.
type MockICacheReader struct {
	ctrl     *gomock.Controller
	recorder *MockICacheReaderMockRecorder
}

// MockICacheReaderMockRecorder is the mock recorder for MockICacheReader.
type MockICacheReaderMockRecorder struct {
	mock *MockICacheReader
}

// NewMockICacheReader creates a new mock instance.
func NewMockICacheReader(ctrl *gomock.Controller) *MockICacheReader {
	mock := &MockICacheReader{ctrl: ctrl}
	mock.recorder = &MockICacheReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICacheReader) EXPECT() *MockICacheReaderMockRecorder {
	return m.recorder
}

// ReadFromFile mocks base method.
func (m *MockICacheReader) ReadFromFile(cacheFolder string) map[string][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromFile", cacheFolder)
	ret0, _ := ret[0].(map[string][]byte)
	return ret0
}

// ReadFromFile indicates an expected call of ReadFromFile.
func (mr *MockICacheReaderMockRecorder) ReadFromFile(cacheFolder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromFile", reflect.TypeOf((*MockICacheReader)(nil).ReadFromFile), cacheFolder)
}
