package core

import (
	"mini-redis-go/internal/service/server/core/cache"
	"sync"
)

// RedisService is a Redis service itself
type RedisService interface {
	Get(key string) []byte
	Set(key string, value []byte)
	ExistsByKey(key string) bool
	ReadCache(cacheFolder string)
	WriteCache(cacheFolder string, k string, v []byte) error
}

type redisService struct {
	cacheReaderService cache.CacheReaderService
	cacheWriterService cache.CacheWriterService
	db                 *MyDb
}

type MyDb struct {
	rwMutex sync.RWMutex
	cache   map[string][]byte
}

//var once sync.Once
//
//func GetMyRedis(cacheFolder, cacheFileName string) IRedis {
//	once.Do(func() {
//		db := MyDb{
//			cache: map[string][]byte{},
//		}
//		instance = &myRedis{
//			db:            &db,
//			cacheFileName: filepath.Join(cacheFolder, cacheFileName),
//		}
//	})
//	return instance
//}

func NewRedisService() *redisService {
	db := MyDb{
		cache: map[string][]byte{},
	}
	return &redisService{
		cacheReaderService: cache.NewCacheReaderService(),
		cacheWriterService: cache.NewCacheWriterService(),
		db:                 &db,
	}
}

func (r *redisService) Get(key string) []byte {
	r.db.rwMutex.RLock()
	defer r.db.rwMutex.RUnlock()
	bytes := r.db.cache[key]
	return bytes
}

func (r *redisService) Set(key string, value []byte) {
	r.db.rwMutex.Lock()
	defer r.db.rwMutex.Unlock()

	r.db.cache[key] = value
}

func (r *redisService) ExistsByKey(key string) bool {
	r.db.rwMutex.RLock()
	defer r.db.rwMutex.RUnlock()
	_, exists := r.db.cache[key]
	return exists
}

func (r *redisService) ReadCache(cacheFolder string) {
	foundCache := r.cacheReaderService.ReadFromFile(cacheFolder)
	for k, v := range foundCache {
		r.Set(k, v)
	}
}

func (r *redisService) WriteCache(cacheFolder string, k string, v []byte) error {
	return r.cacheWriterService.WriteToFile(cacheFolder, k, v)
}
