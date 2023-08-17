package core

import (
	"net"
	"sync"
)

// BrokerService is a service that handles pub/sub
type BrokerService interface {
	IsSubscriptionConnection(net.Conn) bool
	Subscribe(conn net.Conn, topic string)
	Unsubscribe(conn net.Conn)
	GetTopicFromConnection(conn net.Conn) (string, bool)
	Publish(conn net.Conn, topic string, message string)
}

type brokerService struct {
	mutex sync.Mutex
	// clients map of connection and topic
	clients map[net.Conn]string
	// subscribers map of topic and connections
	subscribers map[string][]net.Conn
}

func NewBrokerService() *brokerService {
	return &brokerService{
		clients:     map[net.Conn]string{},
		subscribers: map[string][]net.Conn{},
	}
}

func (m *brokerService) IsSubscriptionConnection(conn net.Conn) bool {
	_, exists := m.clients[conn]
	return exists
}

func (m *brokerService) Subscribe(conn net.Conn, topic string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, exists := m.subscribers[topic]
	if exists {
		v := append(v, conn)
		m.subscribers[topic] = v
	} else {
		m.subscribers[topic] = []net.Conn{conn}
	}
	m.clients[conn] = topic
}

func (m *brokerService) Unsubscribe(conn net.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	topic, exists := m.clients[conn]
	if exists {
		updatedSubscriber := removeConnection(m.subscribers[topic], conn)
		m.subscribers[topic] = updatedSubscriber
		delete(m.clients, conn)
	}

	_ = conn.Close()
}

func removeConnection(conns []net.Conn, conn net.Conn) []net.Conn {
	for i, v := range conns {
		if v == conn {
			return append(conns[:i], conns[i+1:]...)
		}
	}
	return conns
}

func (m *brokerService) GetTopicFromConnection(conn net.Conn) (string, bool) {
	topic, exists := m.clients[conn]
	return topic, exists
}

func (m *brokerService) Publish(conn net.Conn, topic string, message string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	subscribers := m.subscribers[topic]
	for _, aSubscriber := range subscribers {
		// exclude the sender
		if aSubscriber == conn {
			continue
		}
		_, _ = aSubscriber.Write([]byte(message))
	}
}
