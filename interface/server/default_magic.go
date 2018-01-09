package server

import (
	"fmt"
)

type defaultMagic struct {
	name  string
	level string
	opts  Options
}

func init() {
}

func (m *defaultMagic) Init(opts ...Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	return nil
}

func (m *defaultMagic) Start(name string) (string, error) {
	fmt.Printf("defaultMagic")
	return name, nil
}

func (m *defaultMagic) Close(name string) (string, error) {
	return name, nil
}

func (m *defaultMagic) Options() Options {
	return m.opts
}

func newDefaultMagic(opts ...Option) Magic {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	return &defaultMagic{
		opts: options,
	}
}
