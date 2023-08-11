package cache

import (
	"log"
	"mini-redis-go/internal/model"
	"os"
	"path/filepath"
	"sync"
)

var cacheReaderServiceInstance *CacheReaderService
var cacheReaderServiceMutex = &sync.Mutex{}

// ICacheReader is a service to read cache from files
type ICacheReader interface {
	ReadFromFile(cacheFolder string) map[string][]byte
}

type CacheReaderService struct {
}

func NewCacheReaderService() *CacheReaderService {
	if cacheReaderServiceInstance == nil {
		cacheReaderServiceMutex.Lock()
		defer cacheReaderServiceMutex.Unlock()
		if cacheReaderServiceInstance == nil {
			cacheReaderServiceInstance = &CacheReaderService{}
		}
	}
	return cacheReaderServiceInstance
}

func (*CacheReaderService) ReadFromFile(cacheFolder string) map[string][]byte {
	log.Println("reading cache... from", cacheFolder)

	cache := map[string][]byte{}
	err := filepath.Walk(cacheFolder, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			v, readErr := readCacheFromFile(path)
			if readErr == nil {
				cache[fileInfo.Name()] = v
			}
		}
		return nil
	})
	if err != nil {
		log.Panic("error reading cache: ", err)
	}
	log.Println("reading cache... done")
	return cache
}

func readCacheFromFile(cacheFilePath string) ([]byte, error) {
	log.Println("	uncache:", cacheFilePath)

	data, err := os.ReadFile(cacheFilePath)
	if err != nil {
		return nil, err
	}
	_, extractedDate := model.ExtractByteTypeAndValue(data)
	return extractedDate, nil
}
