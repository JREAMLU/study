package hello

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSay(t *testing.T) {
	Convey("say", t, func() {
		s := Say()
		So(s, ShouldEqual, "abc")
		t.Log(s)
	})
}
