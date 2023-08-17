package cache

import (
	"os"
	"path/filepath"
)

// CacheWriterService is a service to write cache to files
type CacheWriterService interface {
	WriteToFile(cacheFolder string, k string, v []byte) error
}

type cacheWriterService struct {
}

func NewCacheWriterService() *cacheWriterService {
	return &cacheWriterService{}
}

func (*cacheWriterService) WriteToFile(cacheFolder string, k string, v []byte) error {
	file, err := os.OpenFile(filepath.Join(cacheFolder, k), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	return writeToFile(file, v)
}

func writeToFile(file *os.File, v []byte) error {
	_, err := file.Write(v)
	if err != nil {
		return err
	}
	return nil
}
