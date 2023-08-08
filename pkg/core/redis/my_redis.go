package redis

import (
	"mini-redis-go/pkg/server/conversion"
	"sync"
)

type ByteType uint8

const (
	Unknown ByteType = iota
	StringByteType
	IntByteType
	StructByteType
)

type Redis interface {
	Get(key string) []byte
	SetByteArray(key string, value []byte)
	SetString(key string, value string)
	SetInt(key string, value int)
	//SetStruct(key string, value struct{})
	Db() map[string][]byte
	ExistsByKey(key string) bool
}

type MyRedis struct {
	db *MyDb
}

var instance *MyRedis
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
			instance = &MyRedis{
				db: &db,
			}
		}
	}
}

func GetMyRedis() *MyRedis {
	return instance
}

func (c *MyRedis) Get(key string) []byte {
	bytes := c.db.cache[key]
	return bytes
}

func (c *MyRedis) SetByteArray(key string, value []byte) {
	c.db.mutex.Lock()
	defer c.db.mutex.Unlock()

	c.db.cache[key] = value
}

func (c *MyRedis) SetString(key string, value string) {
	c.db.mutex.Lock()
	defer c.db.mutex.Unlock()

	byteArray, _ := conversion.ToByteArray(value)
	c.db.cache[key] = byteArray
}

func (c *MyRedis) SetInt(key string, value int) {
	c.db.mutex.Lock()
	defer c.db.mutex.Unlock()

	byteArray, _ := conversion.ToByteArray(value)
	c.db.cache[key] = byteArray
}

func (c *MyRedis) Db() map[string][]byte {
	return c.db.cache
}

func (c *MyRedis) ExistsByKey(key string) bool {
	_, exists := c.db.cache[key]
	return exists
}
