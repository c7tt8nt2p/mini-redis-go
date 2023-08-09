package utils

import (
	"crypto/tls"
	"testing"
)

func TestWriteToServer(t *testing.T) {
	type args struct {
		conn    *tls.Conn
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteToServer(tt.args.conn, tt.args.message)
		})
	}
}
