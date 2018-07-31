package config

import "github.com/JREAMLU/j-kit/go-micro/util"

const (
	serverName    = "h1"
	configVersion = "v1"
)

// HelloConfig hello config
type HelloConfig struct {
	*util.Config

	Hello struct {
		Secret string
	}
}

// Load load config
func Load() (*HelloConfig, error) {
	// load redis mysql elastic client

	// load parent config
	config := &HelloConfig{}
	err := util.LoadCustomConfig("10.200.202.35:8500", serverName, configVersion, config)

	return config, err
}
