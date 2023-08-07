package core

import (
	"sync"
)

type Redis interface {
	Get(key string) []byte
	Set(key string, value []byte)
	Db() map[string][]byte
	ExistsByKey(key string) bool
}

type myRedis struct {
	db *MyDb
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

func InitMyRedis() {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			db := MyDb{
				cache: map[string][]byte{},
			}
			instance = &myRedis{
				db: &db,
			}
		}
	}
}

func GetMyRedis() Redis {
	return instance
}

func (c *myRedis) Get(key string) []byte {
	bytes := c.db.cache[key]
	return bytes
}

func (c *myRedis) Set(key string, value []byte) {
	c.db.mutex.Lock()
	defer c.db.mutex.Unlock()

	c.db.cache[key] = value
}

func (c *myRedis) Db() map[string][]byte {
	return c.db.cache
}

func (c *myRedis) ExistsByKey(key string) bool {
	_, exists := c.db.cache[key]
	return exists
}
