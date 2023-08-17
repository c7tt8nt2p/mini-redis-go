// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/chantapat.t/GolandProjects/mini-redis-go/internal/service/server/core/cache/reader_service.go

// Package mock_cache is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCacheReaderService is a mock of CacheReaderService interface.
type MockCacheReaderService struct {
	ctrl     *gomock.Controller
	recorder *MockCacheReaderServiceMockRecorder
}

// MockCacheReaderServiceMockRecorder is the mock recorder for MockCacheReaderService.
type MockCacheReaderServiceMockRecorder struct {
	mock *MockCacheReaderService
}

// NewMockCacheReaderService creates a new mock instance.
func NewMockCacheReaderService(ctrl *gomock.Controller) *MockCacheReaderService {
	mock := &MockCacheReaderService{ctrl: ctrl}
	mock.recorder = &MockCacheReaderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheReaderService) EXPECT() *MockCacheReaderServiceMockRecorder {
	return m.recorder
}

// ReadFromFile mocks base method.
func (m *MockCacheReaderService) ReadFromFile(cacheFolder string) map[string][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromFile", cacheFolder)
	ret0, _ := ret[0].(map[string][]byte)
	return ret0
}

// ReadFromFile indicates an expected call of ReadFromFile.
func (mr *MockCacheReaderServiceMockRecorder) ReadFromFile(cacheFolder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromFile", reflect.TypeOf((*MockCacheReaderService)(nil).ReadFromFile), cacheFolder)
}