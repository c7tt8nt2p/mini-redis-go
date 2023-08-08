package server

import (
	"fmt"
	"log"
	"mini-redis-go/pkg/core"
	"os"
	"path/filepath"
)

func readCache(myRedis core.Redis, cacheFolder string) {
	fmt.Println("Reading cache... from", cacheFolder)

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
	fmt.Println("Reading cache... done")
}

func readCacheFile(myRedis core.Redis, k, cacheFilePath string) error {
	fmt.Println("	Uncache:", cacheFilePath)
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

func getByteTypeAndValue(originalByteArray []byte) (core.ByteType, []byte) {
	firstByte := originalByteArray[0]
	if firstByte == uint8(core.StringByteType) {
		return core.StringByteType, originalByteArray[1:]
	} else if firstByte == uint8(core.IntByteType) {
		return core.IntByteType, originalByteArray[1:]
	} else {
		return core.Unknown, nil
	}
}
