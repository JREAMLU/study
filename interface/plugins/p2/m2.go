package p2

import (
	"fmt"

	"github.com/JREAMLU/study/interface/server"
)

type m2Magic struct {
	name  string
	level string
	opts  server.Options
}

func init() {
	server.MG = NewMagic()
}

func (m *m2Magic) Init(opts ...server.Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	return nil
}

func (m *m2Magic) Start(name string) (string, error) {
	fmt.Printf("m2Magic")
	return name, nil
}

func (m *m2Magic) Close(name string) (string, error) {
	return name, nil
}

func (m *m2Magic) Options() server.Options {
	return m.opts
}

// NewMagic new magic
func NewMagic(opts ...server.Option) server.Magic {
	options := server.Options{}
	for _, o := range opts {
		o(&options)
	}

	return &m2Magic{
		opts: options,
	}
}
