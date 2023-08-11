package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"mini-redis-go/internal/mock"
	"net"
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

	fmt.Println(service.clients)
}

func TestBrokerService_Subscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn1 := mock.NewMockConn(ctrl)
	conn2 := mock.NewMockConn(ctrl)
	topic := "topicA"

	service := NewBrokerService()
	service.Subscribe(conn1, topic)

	assert.Equal(t, 1, len(service.subscribers[topic]))
	assert.ElementsMatch(t, []net.Conn{conn1}, service.subscribers[topic])

	service.Subscribe(conn2, topic)
	assert.Equal(t, 2, len(service.subscribers[topic]))
	assert.ElementsMatch(t, []net.Conn{conn1, conn2}, service.subscribers[topic])
}

func TestBrokerService_Unsubscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := mock.NewMockConn(ctrl)
	conn.EXPECT().Close().Times(1)
	topic := "topicB"

	service := NewBrokerService()
	service.Subscribe(conn, topic)
	service.Unsubscribe(conn)

	assert.Equal(t, 0, len(service.subscribers[topic]))
	_, exists := service.clients[conn]
	assert.False(t, exists)
}

func TestBrokerService_GetTopicFromConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := mock.NewMockConn(ctrl)
	topic := "topicC"

	service := NewBrokerService()
	service.clients[conn] = topic

	response, _ := service.GetTopicFromConnection(conn)
	assert.Equal(t, topic, response)
}

func TestBrokerService_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn1 := mock.NewMockConn(ctrl)
	conn2 := mock.NewMockConn(ctrl)
	topic := "topicC"
	message := "hello"
	conn2.EXPECT().Write([]byte(message))

	service := NewBrokerService()
	service.Subscribe(conn1, topic)
	service.Subscribe(conn2, topic)

	service.Publish(conn1, topic, message)
}
