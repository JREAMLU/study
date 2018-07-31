package config

import (
	"github.com/JREAMLU/j-kit/go-micro/util"
)

const (
	serverName    = "s1"
	configVersion = "v1"
)

// S1Config s1 config
type S1Config struct {
	*util.Config

	// S1 custom define config
	S1 struct {
		Secret string
	}
}

// Load load config
func Load() (*S1Config, error) {
	// load redis mysql elastic client

	// load parent config
	s1config := &S1Config{}
	err := util.LoadCustomConfig("10.200.202.35:8500", serverName, configVersion, s1config)

	return s1config, err
}
