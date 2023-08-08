package broker

import (
	"fmt"
	"net"
	"sync"
)

type MiniRedisBroker interface {
	Subscribe(conn *net.Conn, topic string)
	Unsubscribe(conn *net.Conn, topic string)
	Disconnect(conn *net.Conn)
}

var instance *MyBroker
var mutex = &sync.Mutex{}

type MyBroker struct {
	mutex   sync.Mutex
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

func (m *MyBroker) Unsubscribe(conn *net.Conn, topic string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_ = (*conn).Close()
}

func (m *MyBroker) Disconnect(conn *net.Conn) {
	fmt.Println("Disconnect", conn)
	topic, hasPreviouslySubscribed := m.clients[conn]
	if hasPreviouslySubscribed {
		updatedSubscriber := removeConnection(m.subscribers[topic], conn)
		m.subscribers[topic] = updatedSubscriber
	}
}

func removeConnection(conns []*net.Conn, conn *net.Conn) []*net.Conn {
	for i, v := range conns {
		if v == conn {
			i2 := append(conns[:i], conns[i+1:]...)
			return i2
		}
	}
	return conns
}
