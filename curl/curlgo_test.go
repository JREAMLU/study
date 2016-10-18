package main

import "testing"

func BenchmarkG(b *testing.B) {
	go say("world")
	say("hello")
}
