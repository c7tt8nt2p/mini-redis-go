package server

import (
	"bufio"
	"fmt"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/core"
	"os"
)

func readCache() {
	fmt.Println("Reading cache...")
	readFile, err := os.Open(config.CacheFileName)
	if err != nil {
		fmt.Println("Error reading cache", err)
	}
	defer func(readFile *os.File) {
		_ = readFile.Close()
	}(readFile)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	myRedis := core.GetMyRedis()
	for fileScanner.Scan() {
		ok, k, v := extractKeyValueCache(fileScanner.Text())
		if ok {
			fmt.Println("	Found cache", k, v)
			myRedis.Set(k, v)
		}
	}
}
