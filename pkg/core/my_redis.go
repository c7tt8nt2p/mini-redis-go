package core

import (
	"path/filepath"
	"sync"
)

type Redis interface {
	Get(key string) string
	Set(key string, value string)
	Db() map[string][]byte
	ExistsByKey(key string) bool
}

type myRedis struct {
	db            *MyDb
	cacheFileName string
}

var instance *myRedis
var mutex = &sync.Mutex{}

//var once sync.Once
//
//func GetMyRedis(cacheFolder, cacheFileName string) Redis {
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

func InitMyRedis(cacheFolder, cacheFileName string) {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			db := MyDb{
				cache: map[string][]byte{},
			}
			instance = &myRedis{
				db:            &db,
				cacheFileName: filepath.Join(cacheFolder, cacheFileName),
			}
		}
	}
}

func GetMyRedis() Redis {
	return instance
}

func (c *myRedis) Get(key string) string {
	bytes := c.db.cache[key]
	return string(bytes)
}

func (c *myRedis) Set(key, value string) {
	c.db.mutex.Lock()
	defer c.db.mutex.Unlock()

	c.db.cache[key] = []byte(value)
}

func (c *myRedis) Db() map[string][]byte {
	return c.db.cache
}

func (c *myRedis) ExistsByKey(key string) bool {
	_, exists := c.db.cache[key]
	return exists
}
