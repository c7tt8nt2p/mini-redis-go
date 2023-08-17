package client

import (
	"fmt"
)

// ISubscriber is a client API interface contains functions that user can interact with Redis
// after a client has subscribed to a topic
type ISubscriber interface {
	NextMessage() (string, error)
	Publish(message string) error
	Unsubscribe() error
}

type Subscriber struct {
	clientService *clientService
}

func (s *Subscriber) NextMessage() (string, error) {
	return s.clientService.Read()
}

func (s *Subscriber) Publish(message string) error {
	msg := fmt.Sprintf("%s\n", message)
	return s.clientService.Write([]byte(msg))
}

func (s *Subscriber) Unsubscribe() error {
	msg := "unsubscribe\n"
	return s.clientService.Write([]byte(msg))
}
