package guid

import "testing"

func Benchmark_guid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewObjectId().Hex()
	}
}

func Benchmark_tunix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tunix()
	}
}
