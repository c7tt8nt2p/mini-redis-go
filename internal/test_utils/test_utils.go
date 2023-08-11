package test_utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func ShouldPanicWithError(t *testing.T, f func(), expectedErr string) {
	t.Helper()
	defer func() {
		err := recover()
		assert.Equal(t, expectedErr, err)
	}()
	f()
	t.Errorf("should have panicked but did not")
}

func CreateTempFolder() string {
	folder, err := os.MkdirTemp("", "mini-redis")
	if err != nil {
		log.Fatal("error creating temp folder", err)
	}
	return folder
}

func CreateFileWithData(folder string, k string, v []byte) {
	file, _ := os.OpenFile(filepath.Join(folder, k), os.O_CREATE|os.O_WRONLY, 0600)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, _ = file.Write(v)
}
