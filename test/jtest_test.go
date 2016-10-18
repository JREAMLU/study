package jtest

import (
	"fmt"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoutin(t *testing.T) {
	Convey("TestRoutin", t, func() {
		go Add(1)
		runtime.Gosched()
		Update(1)
		fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
		So(runtime.NumGoroutine(), ShouldBeGreaterThanOrEqualTo, 2)
	})
}

func TestAdd(t *testing.T) {
	Convey("TestAdd", t, func() {
		m := Add(1)
		So(m, ShouldEqual, 1)
		n := Add(2)
		So(n, ShouldEqual, 2)
	})
}

func TestUpdate(t *testing.T) {
	Convey("TestUpdate", t, func() {
		m := Update(1)
		So(m, ShouldEqual, 1)
		n := Update(2)
		So(n, ShouldEqual, 2)
	})
}
