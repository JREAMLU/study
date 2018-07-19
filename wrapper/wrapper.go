package main

import "fmt"

func main() {
	opts := newOpts(
		wrappName("lu"),
		wrappAge(18),
	)

	fmt.Println("++++++++++++: ", opts)
}

type option func(*options)

type options struct {
	Name string
	Age  int
}

func newOpts(opts ...option) options {
	opt := options{}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func wrappName(name string) option {
	return func(opts *options) {
		opts.Name = name
	}
}

func wrappAge(age int) option {
	return func(opts *options) {
		opts.Age = age
	}
}
