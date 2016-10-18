package ip

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExternalIP(t *testing.T) {
	Convey("getip", t, func() {
		ip, err := ExternalIP()
		So(err, ShouldEqual, nil)
		So(ip, ShouldEqual, "172.16.202.10")
	})
}

func Benchmark_getAdmin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExternalIP()
	}
}

func Benchmark_add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1)
	}
}
