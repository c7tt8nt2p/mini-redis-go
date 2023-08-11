package cache

import (
	"os"
	"path/filepath"
	"sync"
)

var cacheWriterServiceInstance *CacheWriterService
var cacheWriterServiceMutex = &sync.Mutex{}

// ICacheWriter is a service to write cache to files
type ICacheWriter interface {
	WriteToFile(cacheFolder string, k string, v []byte) error
}

type CacheWriterService struct {
}

func NewCacheWriterService() *CacheWriterService {
	if cacheWriterServiceInstance == nil {
		cacheWriterServiceMutex.Lock()
		defer cacheWriterServiceMutex.Unlock()
		if cacheWriterServiceInstance == nil {
			cacheWriterServiceInstance = &CacheWriterService{}
		}
	}
	return cacheWriterServiceInstance
}

func (c *CacheWriterService) WriteToFile(cacheFolder string, k string, v []byte) error {
	file, err := os.OpenFile(filepath.Join(cacheFolder, k), os.O_CREATE|os.O_WRONLY, 0644)
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
