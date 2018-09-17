package config

import (
	"github.com/JREAMLU/j-kit/go-micro/util"
)

const (
	serverName    = "s2"
	configVersion = "v1"
)

// S2Config s2 config
type S2Config struct {
	*util.Config

	// s2 custom define config
	s2 struct {
		Secret string
	}
}

// Load load config
func Load() (*S2Config, error) {
	// load redis mysql elastic client

	// load parent config
	s2Config := &S2Config{}
	err := util.LoadCustomConfig("127.0.0.1:8500", serverName, configVersion, s2Config)

	return s2Config, err
}
