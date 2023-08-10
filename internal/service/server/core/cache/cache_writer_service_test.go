package cache

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewCacheWriterService(t *testing.T) {
	service := NewCacheWriterService()

	assert.NotNil(t, service)
}

func TestCacheWriterService_WriteToFile(t *testing.T) {
	tempFolder, _ := os.MkdirTemp("", "mini-redis")
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)

	k := "keyA"
	v := []byte("ValueA")
	service := NewCacheWriterService()
	if err := service.WriteToFile(tempFolder, k, v); err != nil {
		t.Error("Error WriteToFile", err)
	}

	data, err := os.ReadFile(filepath.Join(tempFolder, k))
	if err != nil {
		t.Error("Error ReadFile", err)
	}
	assert.Equal(t, v, data)
}
