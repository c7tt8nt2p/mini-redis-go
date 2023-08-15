package config

type ServerConfig struct {
	ServerPublicKeyFile  string
	ServerPrivateKeyFile string
	ConnectionHost       string
	ConnectionPort       string
	CacheFolder          string
}
