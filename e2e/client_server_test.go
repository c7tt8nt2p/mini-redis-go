package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"mini-redis-go/e2e/internal"
	"mini-redis-go/internal/model"
	"mini-redis-go/internal/test_utils"
	"os"
	"testing"
)

func TestPingPong(t *testing.T) {
	tempFolder := test_utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := internal.StartServer(tempFolder)
	defer s.Stop()

	c := internal.ConnectToServer(internal.GetClientConfigTest())

	internal.Write(t, c, []byte("PING\n"))
	response := internal.Read(t, c)
	assert.Equal(t, "PONG\n", response)
}

func TestSetAndGet(t *testing.T) {
	tempFolder := test_utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := internal.StartServer(tempFolder)
	defer s.Stop()

	c := internal.ConnectToServer(internal.GetClientConfigTest())

	internal.Set(t, c, "hello", "world")
	response1 := internal.Read(t, c)
	assert.Equal(t, "Set ok\n", response1)

	response2 := internal.Get(t, c, "hello")
	assert.Equal(t, "world\n", response2)
}

func TestCache(t *testing.T) {
	tempFolder := test_utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	test_utils.CreateFileWithData(tempFolder, "testKey", append([]byte{byte(model.StringByteType)}, []byte("tesValue")...))
	s := internal.StartServer(tempFolder)
	defer s.Stop()
	c := internal.ConnectToServer(internal.GetClientConfigTest())
	response := internal.Get(t, c, "testKey")
	assert.Equal(t, "tesValue\n", response)
}
