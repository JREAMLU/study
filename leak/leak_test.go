package main

import (
	"context"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
)

// Default "Check" will poll for 5 seconds to check that all
// goroutines are cleaned up
func TestPool(t *testing.T) {
	defer leaktest.Check(t)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

// Helper function to timeout after X duration
func TestPoolTimeout(t *testing.T) {
	defer leaktest.CheckTimeout(t, time.Second)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

// Use Go 1.7+ context.Context for cancellation
func TestPoolContext(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	defer leaktest.CheckContext(ctx, t)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}
