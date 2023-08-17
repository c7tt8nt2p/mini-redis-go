package main

import (
	"flag"
	"mini-redis-go/internal/config"
)

const (
	defaultPublicKeyFile  = "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/server/server.pem"
	defaultPrivateKeyFile = "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/server/server.key"
	defaultHost           = "localhost"
	defaultPort           = "6973"
	defaultCacheFolder    = "/Users/chantapat.t/GolandProjects/mini-redis-go/cache"
)

func GetServerConfig() *config.ServerConfig {
	publicKeyFile := flag.String("publickey", defaultPublicKeyFile, "a server public key file")
	privateKeyFile := flag.String("privatekey", defaultPrivateKeyFile, "a server private key file")
	hostFlag := flag.String("host", defaultHost, "a connection host to listen on")
	portFlag := flag.String("port", defaultPort, "a connection port to listen on")
	cacheFolder := flag.String("cachefolder", defaultCacheFolder, "a cache folder to store states")
	flag.Parse()

	return &config.ServerConfig{
		PublicKeyFile:  *publicKeyFile,
		PrivateKeyFile: *privateKeyFile,
		Host:           *hostFlag,
		Port:           *portFlag,
		CacheFolder:    *cacheFolder,
	}
}

func main() {
	myApp := InitializeServer()
	myApp.StartServer()
}
