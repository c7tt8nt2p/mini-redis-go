package integration_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/integration_test/utils"
	"mini-redis-go/internal/model"
	"os"
	"path/filepath"
	"testing"
)

func TestPingPong(t *testing.T) {
	tempFolder := utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := utils.StartServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	c := utils.ConnectToServer(config.ConnectionHost, config.ConnectionPort)

	utils.Write(t, c, []byte("PING\n"))
	response := utils.Read(t, c)
	assert.Equal(t, "PONG\n", response)

	s.Stop()
}

func TestSetAndGet(t *testing.T) {
	tempFolder := utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := utils.StartServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	c := utils.ConnectToServer(config.ConnectionHost, config.ConnectionPort)

	utils.Set(t, c, "hello", "world")
	response1 := utils.Read(t, c)
	assert.Equal(t, "Set ok\n", response1)

	response2 := utils.Get(t, c, "hello")
	assert.Equal(t, "world\n", response2)

	s.Stop()
}

func TestCache(t *testing.T) {
	tempFolder := utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	fmt.Println("tempFolder", tempFolder)
	createCacheFileWithData(tempFolder, "testKey", append([]byte{byte(model.StringByteType)}, []byte("tesValue")...))
	s := utils.StartServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	c := utils.ConnectToServer(config.ConnectionHost, config.ConnectionPort)

	response := utils.Get(t, c, "testKey")
	assert.Equal(t, "tesValue\n", response)

	s.Stop()
}

func createCacheFileWithData(folder string, k string, v []byte) {
	file, _ := os.OpenFile(filepath.Join(folder, k), os.O_CREATE|os.O_WRONLY, 0644)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, _ = file.Write(v)
}
