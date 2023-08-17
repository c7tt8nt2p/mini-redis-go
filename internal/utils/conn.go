package utils

import (
	"io"
	"log"
)

func WriteToServer(w io.Writer, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Panic("Error sending message to server: ", err)
	}
}
