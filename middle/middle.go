package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Adder Adder
type Adder interface {
	Add(x, y int) int
}

// AdderFunc AdderFunc
type AdderFunc func(x, y int) int

// Add Add
func (a AdderFunc) Add(x, y int) int {
	return a(x, y)
}

// AdderMiddleware function, this function takes in a `Adder` and returns a new `Adder`.
type AdderMiddleware func(Adder) Adder

// WrapLogger WrapLogger
func WrapLogger(logger *log.Logger) AdderMiddleware {
	return func(a Adder) Adder {
		// Using `AdderFunc` to implement the `Adder` interface.
		fn := func(x, y int) (result int) {
			defer func(t time.Time) {
				logger.Printf("took=%v, x=%v, y=%v, result=%v", time.Since(t), x, y, result)
			}(time.Now())
			return a.Add(x, y)
		}

		return AdderFunc(fn)
	}
}

// WrapCache WrapCache
func WrapCache(cache *sync.Map) AdderMiddleware {
	return func(a Adder) Adder {
		fn := func(x, y int) int {
			key := fmt.Sprintf("x=%dy=%d", x, y)
			val, ok := cache.Load(key)
			if ok {
				return val.(int)
			}
			result := a.Add(x, y)
			cache.Store(key, result)
			return result
		}
		return AdderFunc(fn)
	}
}

// WrapJream WrapJream
func WrapJream(name string) AdderMiddleware {
	return func(a Adder) Adder {
		fn := func(x, y int) int {
			fmt.Println("++++++++++++: wrap hook:", name)
			result := a.Add(x, y)
			return result
		}
		return AdderFunc(fn)
	}
}

// Chain Chain
func Chain(outer AdderMiddleware, middleware ...AdderMiddleware) AdderMiddleware {
	return func(a Adder) Adder {
		topIndex := len(middleware) - 1
		for i := range middleware {
			fmt.Println("++++++++++++: 执行")
			fmt.Println("++++++++++++: ", topIndex-i)
			a = middleware[topIndex-i](a)
		}
		return outer(a)
	}
}

func main() {
	/*
		logger := log.New(os.Stdout, "test ", 1)

		var a Adder = AdderFunc(
			func(x, y int) int {
				return x + y
			},
		)

		w := WrapLogger(logger)
		w(a).Add(10, 20)
	*/

	logger := log.New(os.Stdout, "test ", 1)

	var a Adder = AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)
	a = Chain(
		WrapLogger(logger),
		WrapCache(&sync.Map{}),
		WrapJream("abc"),
	)(a)

	a.Add(10, 20)
}
