package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"mini-redis-go/internal/model"
	"mini-redis-go/internal/test_utils"
	"os"
	"testing"
)

func TestNewCacheReaderService(t *testing.T) {
	service := NewCacheReaderService()

	assert.NotNil(t, service)
}

func TestCacheReaderService_ReadFromFile(t *testing.T) {
	tempFolder, _ := os.MkdirTemp("", "mini-redis")
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	test_utils.CreateFileWithData(tempFolder, "key1", append([]byte{byte(model.StringByteType)}, []byte("hello world")...))
	test_utils.CreateFileWithData(tempFolder, "key2", append([]byte{byte(model.StringByteType)}, []byte("hello there")...))

	service := NewCacheReaderService()
	cache := service.ReadFromFile(tempFolder)

	assert.Equal(t, 2, len(cache))
	assert.Equal(t, []byte("hello world"), cache["key1"])
	assert.Equal(t, []byte("hello there"), cache["key2"])
}

func TestCacheReaderService_ReadFromFile_ShouldError(t *testing.T) {
	tempFolder := "./invalidFolder"
	test_utils.CreateFileWithData(tempFolder, "key1", append([]byte{byte(model.StringByteType)}, []byte("hello world")...))

	service := NewCacheReaderService()

	expectedErr := fmt.Sprintf("error reading cache: lstat %s: no such file or directory", tempFolder)
	test_utils.ShouldPanicWithError(t, func() { service.ReadFromFile(tempFolder) }, expectedErr)

}
