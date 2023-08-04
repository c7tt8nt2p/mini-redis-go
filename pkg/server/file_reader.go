package server

import (
	"bufio"
	"fmt"
	"mini-redis-go/pkg/core"
	"os"
	"path/filepath"
)

func readCache(cacheFolder, cacheFileName string) {
	fmt.Println("Reading cache... from", filepath.Join(cacheFolder, cacheFileName))
	readFile, err := os.Open(filepath.Join(cacheFolder, cacheFileName))
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
