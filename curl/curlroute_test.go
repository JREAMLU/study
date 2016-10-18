package main

import "testing"

func BenchmarkA(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rou()
	}
}
