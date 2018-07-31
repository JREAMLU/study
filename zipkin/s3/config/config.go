package config

import (
	"github.com/JREAMLU/j-kit/go-micro/util"
)

const (
	serverName    = "s3"
	configVersion = "v1"
)

// S3Config s3 config
type S3Config struct {
	*util.Config

	// s3 custom define config
	s3 struct {
		Secret string
	}
}

// Load load config
func Load() (*S3Config, error) {
	// load redis mysql elastic client

	// load parent config
	s3Config := &S3Config{}
	err := util.LoadCustomConfig("10.200.202.35:8500", serverName, configVersion, s3Config)

	return s3Config, err
}
