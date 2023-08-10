package core

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"mini-redis-go/internal/mock"
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestNewBrokerService(t *testing.T) {
	service := NewBrokerService()

	assert.NotNil(t, service)
}

func TestBrokerService_IsSubscriptionConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn1 := mock.NewMockConn(ctrl)
	conn2 := mock.NewMockConn(ctrl)

	service := NewBrokerService()
	service.clients[conn1] = "topic1"

	assert.True(t, service.IsSubscriptionConnection(conn1))
	assert.False(t, service.IsSubscriptionConnection(conn2))
}

func TestBrokerService_GetTopicFromConnection(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		clients     map[net.Conn]string
		subscribers map[string][]net.Conn
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BrokerService{
				mutex:       tt.fields.mutex,
				clients:     tt.fields.clients,
				subscribers: tt.fields.subscribers,
			}
			got, got1 := m.GetTopicFromConnection(tt.args.conn)
			if got != tt.want {
				t.Errorf("GetTopicFromConnection() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTopicFromConnection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBrokerService_Publish(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		clients     map[net.Conn]string
		subscribers map[string][]net.Conn
	}
	type args struct {
		conn    net.Conn
		topic   string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BrokerService{
				mutex:       tt.fields.mutex,
				clients:     tt.fields.clients,
				subscribers: tt.fields.subscribers,
			}
			m.Publish(tt.args.conn, tt.args.topic, tt.args.message)
		})
	}
}

func TestBrokerService_Subscribe(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		clients     map[net.Conn]string
		subscribers map[string][]net.Conn
	}
	type args struct {
		conn  net.Conn
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BrokerService{
				mutex:       tt.fields.mutex,
				clients:     tt.fields.clients,
				subscribers: tt.fields.subscribers,
			}
			m.Subscribe(tt.args.conn, tt.args.topic)
		})
	}
}

func TestBrokerService_Unsubscribe(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		clients     map[net.Conn]string
		subscribers map[string][]net.Conn
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BrokerService{
				mutex:       tt.fields.mutex,
				clients:     tt.fields.clients,
				subscribers: tt.fields.subscribers,
			}
			m.Unsubscribe(tt.args.conn)
		})
	}
}

func Test_removeConnection(t *testing.T) {
	type args struct {
		conns []net.Conn
		conn  net.Conn
	}
	tests := []struct {
		name string
		args args
		want []net.Conn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeConnection(tt.args.conns, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
