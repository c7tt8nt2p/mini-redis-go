package cache

import (
	"os"
	"path/filepath"
)

func WriteCache(cacheFolder string, k string, v []byte) error {
	file, err := os.OpenFile(filepath.Join(cacheFolder, k), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	return writeToFile(file, v)
}

func writeToFile(file *os.File, v []byte) error {
	_, err := file.Write(v)
	if err != nil {
		return err
	}
	return nil
}
