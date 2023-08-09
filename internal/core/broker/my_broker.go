package broker

import (
	"net"
	"sync"
)

type IBroker interface {
	IsSubscriptionConnection(*net.Conn) bool
	Subscribe(conn *net.Conn, topic string)
	Unsubscribe(conn *net.Conn)
	GetTopicFromConnection(conn *net.Conn) (bool, string)
	Publish(conn *net.Conn, topic string, message string)
}

var instance *MyBroker
var mutex = &sync.Mutex{}

type MyBroker struct {
	mutex sync.Mutex
	// map of connection and topic
	clients map[*net.Conn]string
	// map of topic and connections
	subscribers map[string][]*net.Conn
}

func InitMyBroker() {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &MyBroker{
				clients:     map[*net.Conn]string{},
				subscribers: map[string][]*net.Conn{},
			}
		}
	}
}

func GetMyBroker() *MyBroker {
	return instance
}

func (m *MyBroker) IsSubscriptionConnection(conn *net.Conn) bool {
	_, exists := m.clients[conn]
	return exists
}

func (m *MyBroker) Subscribe(conn *net.Conn, topic string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, keyExists := m.subscribers[topic]
	if keyExists {
		updatedConns := append(v, conn)
		m.subscribers[topic] = updatedConns
	} else {
		m.subscribers[topic] = []*net.Conn{conn}
	}
	m.clients[conn] = topic
}

func (m *MyBroker) Unsubscribe(conn *net.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	topic, hasPreviouslySubscribed := m.clients[conn]
	if hasPreviouslySubscribed {
		updatedSubscriber := removeConnection(m.subscribers[topic], conn)
		m.subscribers[topic] = updatedSubscriber
		delete(m.clients, conn)
	}

	_ = (*conn).Close()
}

func removeConnection(conns []*net.Conn, conn *net.Conn) []*net.Conn {
	for i, v := range conns {
		if v == conn {
			updatedList := append(conns[:i], conns[i+1:]...)
			return updatedList
		}
	}
	return conns
}

func (m *MyBroker) GetTopicFromConnection(conn *net.Conn) (string, bool) {
	topic, exists := m.clients[conn]
	return topic, exists
}

func (m *MyBroker) Publish(conn *net.Conn, topic string, message string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	subscribers := m.subscribers[topic]
	for _, aSubscriber := range subscribers {
		// exclude the sender
		if aSubscriber == conn {
			continue
		}
		_, _ = (*aSubscriber).Write([]byte(message))
	}
}
