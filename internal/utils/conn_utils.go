package utils

import (
	"crypto/tls"
	"log"
)

func WriteToServer(conn *tls.Conn, message string) {
	_, err := (*conn).Write([]byte(message))
	if err != nil {
		log.Panic("Error sending message to server: ", err)
	}
}
