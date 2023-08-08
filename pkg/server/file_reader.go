package server

import (
	"log"
	"mini-redis-go/pkg/core/redis"
	"os"
	"path/filepath"
)

func readCache(myRedis redis.Redis, cacheFolder string) {
	log.Println("reading cache... from", cacheFolder)

	err := filepath.Walk(cacheFolder, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			return readCacheFile(myRedis, fileInfo.Name(), path)
		}
		return nil
	})
	if err != nil {
		log.Panic("error reading cache: ", err)
	}
	log.Println("reading cache... done")
}

func readCacheFile(myRedis redis.Redis, k, cacheFilePath string) error {
	log.Println("	uncache:", cacheFilePath)
	file, err := os.Open(cacheFilePath)
	if err != nil {
		return err
	}
	defer func(readFile *os.File) {
		_ = readFile.Close()
	}(file)

	data, err := os.ReadFile(cacheFilePath)
	if err != nil {
		return err
	}
	_, realData := getByteTypeAndValue(data)
	myRedis.SetByteArray(k, realData)
	return nil
}

func getByteTypeAndValue(originalByteArray []byte) (redis.ByteType, []byte) {
	firstByte := originalByteArray[0]
	if firstByte == uint8(redis.StringByteType) {
		return redis.StringByteType, originalByteArray[1:]
	} else if firstByte == uint8(redis.IntByteType) {
		return redis.IntByteType, originalByteArray[1:]
	} else {
		return redis.Unknown, nil
	}
}
