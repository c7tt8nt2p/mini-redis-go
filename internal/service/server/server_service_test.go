package server

import (
	"bufio"
	"crypto/tls"
	"net"
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		host        string
		port        string
		cacheFolder string
	}
	tests := []struct {
		name string
		args args
		want IServer
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.host, tt.args.port, tt.args.cacheFolder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	type fields struct {
		listener    *net.Listener
		Addr        string
		CacheFolder string
		stopSignal  chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				listener:    tt.fields.listener,
				Addr:        tt.fields.Addr,
				CacheFolder: tt.fields.CacheFolder,
				stopSignal:  tt.fields.stopSignal,
			}
			s.Start()
		})
	}
}

func TestServer_Stop(t *testing.T) {
	type fields struct {
		listener    *net.Listener
		Addr        string
		CacheFolder string
		stopSignal  chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				listener:    tt.fields.listener,
				Addr:        tt.fields.Addr,
				CacheFolder: tt.fields.CacheFolder,
				stopSignal:  tt.fields.stopSignal,
			}
			s.Stop()
		})
	}
}

func TestServer_listen(t *testing.T) {
	type fields struct {
		listener    *net.Listener
		Addr        string
		CacheFolder string
		stopSignal  chan bool
	}
	type args struct {
		listener net.Listener
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				listener:    tt.fields.listener,
				Addr:        tt.fields.Addr,
				CacheFolder: tt.fields.CacheFolder,
				stopSignal:  tt.fields.stopSignal,
			}
			s.listen(tt.args.listener)
		})
	}
}

func Test_acceptANewConnection(t *testing.T) {
	type args struct {
		listener *net.Listener
	}
	tests := []struct {
		name    string
		args    args
		want    *net.Conn
		wantErr bool
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := acceptANewConnection(tt.args.listener)
			if (err != nil) != tt.wantErr {
				t.Errorf("acceptANewConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("acceptANewConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleConnection(t *testing.T) {
	type args struct {
		server *Server
		conn   *net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleConnection(tt.args.server, tt.args.conn)
		})
	}
}

func Test_handleNonSubscriptionConnection(t *testing.T) {
	type args struct {
		server  *Server
		conn    *net.Conn
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleNonSubscriptionConnection(tt.args.server, tt.args.conn, tt.args.message)
		})
	}
}

func Test_handleSubscriptionConnection(t *testing.T) {
	type args struct {
		conn    *net.Conn
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleSubscriptionConnection(tt.args.conn, tt.args.message)
		})
	}
}

func Test_loadCert(t *testing.T) {
	tests := []struct {
		name string
		want *tls.Certificate
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadCert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadCert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readMessage(t *testing.T) {
	type args struct {
		reader *bufio.Reader
		conn   *net.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test_utils cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readMessage(tt.args.reader, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("readMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
