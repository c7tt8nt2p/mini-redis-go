package core

import "sync"

type MyDb struct {
	mutex sync.Mutex
	cache map[string][]byte
}
