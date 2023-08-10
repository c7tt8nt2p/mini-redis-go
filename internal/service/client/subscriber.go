package client

import (
	"fmt"
)

type ISubscriber interface {
	NextMessage() (string, error)
	Publish(message string) error
	Unsubscribe() error
}

type Subscriber struct {
	c *ClientService
}

func (s *Subscriber) NextMessage() (string, error) {
	return s.c.Read()
}

func (s *Subscriber) Publish(message string) error {
	msg := fmt.Sprintf("%s\n", message)
	return s.c.Write([]byte(msg))
}

func (s *Subscriber) Unsubscribe() error {
	msg := "unsubscribe\n"
	return s.c.Write([]byte(msg))
}
