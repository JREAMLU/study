package main

import "fmt"

func main() {
	o := newOptions(client("nnnn"))
	fmt.Println("++++++++++++: ", o)
}

type Option func(*Options)

type Options struct {
	Name string
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Name: "defaultName",
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func client(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}
