// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/chantapat.t/GolandProjects/mini-redis-go/internal/service/server/core/redis_service.go

// Package mock_core is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIRedis is a mock of IRedis interface.
type MockIRedis struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisMockRecorder
}

// MockIRedisMockRecorder is the mock recorder for MockIRedis.
type MockIRedisMockRecorder struct {
	mock *MockIRedis
}

// NewMockIRedis creates a new mock instance.
func NewMockIRedis(ctrl *gomock.Controller) *MockIRedis {
	mock := &MockIRedis{ctrl: ctrl}
	mock.recorder = &MockIRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedis) EXPECT() *MockIRedisMockRecorder {
	return m.recorder
}

// Db mocks base method.
func (m *MockIRedis) Db() map[string][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Db")
	ret0, _ := ret[0].(map[string][]byte)
	return ret0
}

// Db indicates an expected call of Db.
func (mr *MockIRedisMockRecorder) Db() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Db", reflect.TypeOf((*MockIRedis)(nil).Db))
}

// ExistsByKey mocks base method.
func (m *MockIRedis) ExistsByKey(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByKey", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsByKey indicates an expected call of ExistsByKey.
func (mr *MockIRedisMockRecorder) ExistsByKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByKey", reflect.TypeOf((*MockIRedis)(nil).ExistsByKey), key)
}

// Get mocks base method.
func (m *MockIRedis) Get(key string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIRedisMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRedis)(nil).Get), key)
}

// ReadCache mocks base method.
func (m *MockIRedis) ReadCache(cacheFolder string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReadCache", cacheFolder)
}

// ReadCache indicates an expected call of ReadCache.
func (mr *MockIRedisMockRecorder) ReadCache(cacheFolder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCache", reflect.TypeOf((*MockIRedis)(nil).ReadCache), cacheFolder)
}

// SetByteArray mocks base method.
func (m *MockIRedis) SetByteArray(key string, value []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetByteArray", key, value)
}

// SetByteArray indicates an expected call of SetByteArray.
func (mr *MockIRedisMockRecorder) SetByteArray(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByteArray", reflect.TypeOf((*MockIRedis)(nil).SetByteArray), key, value)
}

// SetInt mocks base method.
func (m *MockIRedis) SetInt(key string, value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInt", key, value)
}

// SetInt indicates an expected call of SetInt.
func (mr *MockIRedisMockRecorder) SetInt(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInt", reflect.TypeOf((*MockIRedis)(nil).SetInt), key, value)
}

// SetString mocks base method.
func (m *MockIRedis) SetString(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetString", key, value)
}

// SetString indicates an expected call of SetString.
func (mr *MockIRedisMockRecorder) SetString(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetString", reflect.TypeOf((*MockIRedis)(nil).SetString), key, value)
}

// WriteCache mocks base method.
func (m *MockIRedis) WriteCache(cacheFolder, k string, v []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteCache", cacheFolder, k, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteCache indicates an expected call of WriteCache.
func (mr *MockIRedisMockRecorder) WriteCache(cacheFolder, k, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCache", reflect.TypeOf((*MockIRedis)(nil).WriteCache), cacheFolder, k, v)
}