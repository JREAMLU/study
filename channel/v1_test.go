package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkV1(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	Convey("func V1()", b, func() {
		for i := 0; i < b.N; i++ {
			// serverV2()
		}
	})
	b.StopTimer()
}
