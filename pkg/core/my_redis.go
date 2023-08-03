package core

import (
	"sync"
)

type Redis interface {
	Get(key string) string
	Set(key string, value string)
	Db() map[string][]byte
	Exists(key string) bool
}

type myRedis struct {
	db *MyDb
}

var instance *myRedis
var once sync.Once

func NewMyRedis() Redis {
	once.Do(func() {
		db := MyDb{
			cache: map[string][]byte{},
		}
		instance = &myRedis{
			db: &db,
		}
	})
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

func (c *myRedis) Exists(key string) bool {
	_, exists := c.db.cache[key]
	return exists
}
