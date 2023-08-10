package core

import (
	"mini-redis-go/internal/service/server/core/cache"
	"mini-redis-go/internal/utils"
	"sync"
)

var redisServiceInstance *RedisService
var redisServiceMutex = &sync.Mutex{}

type IRedis interface {
	Get(key string) []byte
	SetByteArray(key string, value []byte)
	SetString(key string, value string)
	SetInt(key string, value int)
	//SetStruct(key string, value struct{})
	Db() map[string][]byte
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
	mutex sync.Mutex
	cache map[string][]byte
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
	bytes := r.db.cache[key]
	return bytes
}

func (r *RedisService) SetByteArray(key string, value []byte) {
	r.db.mutex.Lock()
	defer r.db.mutex.Unlock()

	r.db.cache[key] = value
}

func (r *RedisService) SetString(key string, value string) {
	r.db.mutex.Lock()
	defer r.db.mutex.Unlock()

	byteArray, _ := utils.ToByteArray(value)
	r.db.cache[key] = byteArray
}

func (r *RedisService) SetInt(key string, value int) {
	r.db.mutex.Lock()
	defer r.db.mutex.Unlock()

	byteArray, _ := utils.ToByteArray(value)
	r.db.cache[key] = byteArray
}

func (r *RedisService) Db() map[string][]byte {
	return r.db.cache
}

func (r *RedisService) ExistsByKey(key string) bool {
	_, exists := r.db.cache[key]
	return exists
}

func (r *RedisService) ReadCache(cacheFolder string) {
	foundCache := r.cacheReaderService.ReadFromFile(cacheFolder)
	for k, v := range foundCache {
		r.SetByteArray(k, v)
	}
}

func (r *RedisService) WriteCache(cacheFolder string, k string, v []byte) error {
	return r.cacheWriterService.WriteToFile(cacheFolder, k, v)
}