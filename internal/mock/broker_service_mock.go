// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/chantapat.t/GolandProjects/mini-redis-go/internal/service/server/core/broker_service.go

// Package mock_core is a generated GoMock package.
package mock

import (
	net "net"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIBroker is a mock of IBroker interface.
type MockIBroker struct {
	ctrl     *gomock.Controller
	recorder *MockIBrokerMockRecorder
}

// MockIBrokerMockRecorder is the mock recorder for MockIBroker.
type MockIBrokerMockRecorder struct {
	mock *MockIBroker
}

// NewMockIBroker creates a new mock instance.
func NewMockIBroker(ctrl *gomock.Controller) *MockIBroker {
	mock := &MockIBroker{ctrl: ctrl}
	mock.recorder = &MockIBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBroker) EXPECT() *MockIBrokerMockRecorder {
	return m.recorder
}

// GetTopicFromConnection mocks base method.
func (m *MockIBroker) GetTopicFromConnection(conn net.Conn) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicFromConnection", conn)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTopicFromConnection indicates an expected call of GetTopicFromConnection.
func (mr *MockIBrokerMockRecorder) GetTopicFromConnection(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicFromConnection", reflect.TypeOf((*MockIBroker)(nil).GetTopicFromConnection), conn)
}

// IsSubscriptionConnection mocks base method.
func (m *MockIBroker) IsSubscriptionConnection(arg0 net.Conn) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSubscriptionConnection", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSubscriptionConnection indicates an expected call of IsSubscriptionConnection.
func (mr *MockIBrokerMockRecorder) IsSubscriptionConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubscriptionConnection", reflect.TypeOf((*MockIBroker)(nil).IsSubscriptionConnection), arg0)
}

// Publish mocks base method.
func (m *MockIBroker) Publish(conn net.Conn, topic, message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publish", conn, topic, message)
}

// Publish indicates an expected call of Publish.
func (mr *MockIBrokerMockRecorder) Publish(conn, topic, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockIBroker)(nil).Publish), conn, topic, message)
}

// Subscribe mocks base method.
func (m *MockIBroker) Subscribe(conn net.Conn, topic string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", conn, topic)
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockIBrokerMockRecorder) Subscribe(conn, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockIBroker)(nil).Subscribe), conn, topic)
}

// Unsubscribe mocks base method.
func (m *MockIBroker) Unsubscribe(conn net.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", conn)
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockIBrokerMockRecorder) Unsubscribe(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockIBroker)(nil).Unsubscribe), conn)
}