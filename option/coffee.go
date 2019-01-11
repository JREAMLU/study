package main

import "fmt"

type CoffeeOption func(Opts *CoffeeOptions)

type CoffeeOptions struct {
	sugar        int
	milk         int
	coffeePowder int
}

type Coffee struct {
	opts *CoffeeOptions
}

func CoffeeSugar(sugar int) CoffeeOption {
	return func(opts *CoffeeOptions) {
		opts.sugar = sugar
	}
}

func CoffeeMilk(milk int) CoffeeOption {
	return func(opts *CoffeeOptions) {
		opts.milk = milk
	}
}

func newDefaultCoffeeOptions() *CoffeeOptions {
	return &CoffeeOptions{
		sugar:        2,
		milk:         5,
		coffeePowder: 100,
	}
}

func NewCoffee(opts ...CoffeeOption) *Coffee {
	defaultOptions := newDefaultCoffeeOptions()
	for _, opt := range opts {
		opt(defaultOptions)
	}

	return &Coffee{
		opts: defaultOptions,
	}
}

func main() {
	coffee := NewCoffee(CoffeeMilk(10), CoffeeSugar(99))
	fmt.Println("++++++++++++: ", coffee.opts)
}
