package plugins

import "github.com/JREAMLU/study/interface/server"

type m1Magic struct {
	name  string
	level string
	opts  server.Options
}

func init() {
}

func (m *m1Magic) Init(opts ...server.Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	return nil
}

func (m *m1Magic) Start() (string, error) {
	return m.name, nil
}

func (m *m1Magic) Close() (string, error) {
	return m.name, nil
}

func (m *m1Magic) Options() server.Options {
	return m.opts
}

func NewMagic(opts ...server.Option) server.Magic {
	options := server.Options{}
	for _, o := range opts {
		o(&options)
	}

	return &m1Magic{
		name:  "m1",
		level: "1",
		opts:  options,
	}
}
