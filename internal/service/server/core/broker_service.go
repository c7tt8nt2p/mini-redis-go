package core

import (
	"net"
	"sync"
)

var brokerServiceInstance *BrokerService
var brokerServiceMutex = &sync.Mutex{}

// IBroker is a service that handles pub/sub
type IBroker interface {
	IsSubscriptionConnection(net.Conn) bool
	Subscribe(conn net.Conn, topic string)
	Unsubscribe(conn net.Conn)
	GetTopicFromConnection(conn net.Conn) (string, bool)
	Publish(conn net.Conn, topic string, message string)
}

type BrokerService struct {
	mutex sync.Mutex
	// clients map of connection and topic
	clients map[net.Conn]string
	// subscribers map of topic and connections
	subscribers map[string][]net.Conn
}

func NewBrokerService() *BrokerService {
	if brokerServiceInstance == nil {
		brokerServiceMutex.Lock()
		defer brokerServiceMutex.Unlock()
		if brokerServiceInstance == nil {
			brokerServiceInstance = &BrokerService{
				clients:     map[net.Conn]string{},
				subscribers: map[string][]net.Conn{},
			}
		}
	}
	return brokerServiceInstance
}

func (m *BrokerService) IsSubscriptionConnection(conn net.Conn) bool {
	_, exists := m.clients[conn]
	return exists
}

func (m *BrokerService) Subscribe(conn net.Conn, topic string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, keyExists := m.subscribers[topic]
	if keyExists {
		updatedConns := append(v, conn)
		m.subscribers[topic] = updatedConns
	} else {
		m.subscribers[topic] = []net.Conn{conn}
	}
	m.clients[conn] = topic
}

func (m *BrokerService) Unsubscribe(conn net.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	topic, hasPreviouslySubscribed := m.clients[conn]
	if hasPreviouslySubscribed {
		updatedSubscriber := removeConnection(m.subscribers[topic], conn)
		m.subscribers[topic] = updatedSubscriber
		delete(m.clients, conn)
	}

	_ = conn.Close()
}

func removeConnection(conns []net.Conn, conn net.Conn) []net.Conn {
	for i, v := range conns {
		if v == conn {
			updatedList := append(conns[:i], conns[i+1:]...)
			return updatedList
		}
	}
	return conns
}

func (m *BrokerService) GetTopicFromConnection(conn net.Conn) (string, bool) {
	topic, exists := m.clients[conn]
	return topic, exists
}

func (m *BrokerService) Publish(conn net.Conn, topic string, message string) {
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
