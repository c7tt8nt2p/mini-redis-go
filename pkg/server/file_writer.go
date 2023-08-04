package server

import (
	"fmt"
	"log"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/core"
	"os"
	"path/filepath"
	"strings"
)

func cacheRewrite(myRedis *core.Redis, cacheFolder string) {
	f, err := os.OpenFile(filepath.Join(cacheFolder, config.CacheFileName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	rewriteAllKeyValues(myRedis, f)
}

func cacheAppend(cacheFolder, k, v string) {
	f, err := os.OpenFile(filepath.Join(cacheFolder, config.CacheFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	appendKeyValue(k, v, f)
}

func rewriteAllKeyValues(myRedis *core.Redis, f *os.File) {
	var builder strings.Builder

	for k, v := range (*myRedis).Db() {
		s := fmt.Sprintf("%s=%s\n", k, v)
		builder.WriteString(s)
	}

	if _, err := f.WriteString(builder.String()); err != nil {
		log.Println(err)
	}
}

func appendKeyValue(key string, value string, f *os.File) {
	text := fmt.Sprintf("%s=%s\n", key, value)

	if _, err := f.WriteString(text); err != nil {
		log.Println(err)
	}
}
