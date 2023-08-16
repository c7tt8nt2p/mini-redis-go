package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"mini-redis-go/e2e/utils"
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
	s := utils.StartServer(tempFolder)
	defer s.Stop()

	c := utils.ConnectToServer(utils.GetClientConfigTest())

	utils.Write(t, c, []byte("PING\n"))
	response := utils.Read(t, c)
	assert.Equal(t, "PONG\n", response)
}

func TestSetAndGet(t *testing.T) {
	tempFolder := test_utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := utils.StartServer(tempFolder)
	defer s.Stop()

	c := utils.ConnectToServer(utils.GetClientConfigTest())

	utils.Set(t, c, "hello", "world")
	response1 := utils.Read(t, c)
	assert.Equal(t, "Set ok\n", response1)

	response2 := utils.Get(t, c, "hello")
	assert.Equal(t, "world\n", response2)
}

func TestCache(t *testing.T) {
	tempFolder := test_utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	test_utils.CreateFileWithData(tempFolder, "testKey", append([]byte{byte(model.StringByteType)}, []byte("tesValue")...))
	s := utils.StartServer(tempFolder)
	defer s.Stop()
	c := utils.ConnectToServer(utils.GetClientConfigTest())
	response := utils.Get(t, c, "testKey")
	assert.Equal(t, "tesValue\n", response)
}
