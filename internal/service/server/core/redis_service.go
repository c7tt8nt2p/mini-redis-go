package core

import (
	"mini-redis-go/internal/service/server/core/cache"
	"sync"
)

var redisServiceInstance *RedisService
var redisServiceMutex = &sync.Mutex{}

// IRedis is a Redis service itself
type IRedis interface {
	Get(key string) []byte
	Set(key string, value []byte)
	ExistsByKey(key string) bool
	ReadCache(cacheFolder string)
	WriteCache(cacheFolder string, k string, v []byte) error
}

type RedisService struct {
	cacheReaderService cache.ICacheReader
	cacheWriterService cache.ICacheWriter
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

func NewRedisService() *RedisService {
	if redisServiceInstance == nil {
		redisServiceMutex.Lock()
		defer redisServiceMutex.Unlock()
		if redisServiceInstance == nil {
			db := MyDb{
				cache: map[string][]byte{},
			}
			redisServiceInstance = &RedisService{
				cacheReaderService: cache.NewCacheReaderService(),
				cacheWriterService: cache.NewCacheWriterService(),
				db:                 &db,
			}
		}
	}
	return redisServiceInstance
}

func (r *RedisService) Get(key string) []byte {
	r.db.rwMutex.RLock()
	defer r.db.rwMutex.RUnlock()
	bytes := r.db.cache[key]
	return bytes
}

func (r *RedisService) Set(key string, value []byte) {
	r.db.rwMutex.Lock()
	defer r.db.rwMutex.Unlock()

	r.db.cache[key] = value
}

func (r *RedisService) ExistsByKey(key string) bool {
	r.db.rwMutex.RLock()
	defer r.db.rwMutex.RUnlock()
	_, exists := r.db.cache[key]
	return exists
}

func (r *RedisService) ReadCache(cacheFolder string) {
	foundCache := r.cacheReaderService.ReadFromFile(cacheFolder)
	for k, v := range foundCache {
		r.Set(k, v)
	}
}

func (r *RedisService) WriteCache(cacheFolder string, k string, v []byte) error {
	return r.cacheWriterService.WriteToFile(cacheFolder, k, v)
}
