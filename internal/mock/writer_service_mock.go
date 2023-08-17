// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/chantapat.t/GolandProjects/mini-redis-go/internal/service/server/core/cache/writer_service.go

// Package mock_cache is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCacheWriterService is a mock of CacheWriterService interface.
type MockCacheWriterService struct {
	ctrl     *gomock.Controller
	recorder *MockCacheWriterServiceMockRecorder
}

// MockCacheWriterServiceMockRecorder is the mock recorder for MockCacheWriterService.
type MockCacheWriterServiceMockRecorder struct {
	mock *MockCacheWriterService
}

// NewMockCacheWriterService creates a new mock instance.
func NewMockCacheWriterService(ctrl *gomock.Controller) *MockCacheWriterService {
	mock := &MockCacheWriterService{ctrl: ctrl}
	mock.recorder = &MockCacheWriterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheWriterService) EXPECT() *MockCacheWriterServiceMockRecorder {
	return m.recorder
}

// WriteToFile mocks base method.
func (m *MockCacheWriterService) WriteToFile(cacheFolder, k string, v []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToFile", cacheFolder, k, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToFile indicates an expected call of WriteToFile.
func (mr *MockCacheWriterServiceMockRecorder) WriteToFile(cacheFolder, k, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockCacheWriterService)(nil).WriteToFile), cacheFolder, k, v)
}