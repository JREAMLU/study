package model

import (
	"study/core"
	"testing"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	core.InitGorm("root:root@tcp(localhost:3306)/plucron?charset=utf8&parseTime=True&loc=Local")
}

func TestInsert(t *testing.T) {
	Convey("测试Insert", t, func() {
		cronlist := Cronlist{
			Name: "Iversion",
			Type: 2,
		}
		id, err := Insert(cronlist)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		So(err, ShouldBeNil)
	})
}

func TestUpdate(t *testing.T) {
	Convey("测试Update", t, func() {
		cronlist := Cronlist{
			Name: "jream",
			Type: 3,
		}
		So(Update(cronlist, []uint64{6}), ShouldBeNil)
	})
}

func TestDelete(t *testing.T) {
	Convey("测试Delete", t, func() {
		So(Delete([]uint64{345, 346}), ShouldBeNil)
	})
}

func TestSelect(t *testing.T) {
	Convey("测试Select", t, func() {
		cronlist, err := Select([]uint64{1, 2})
		cl := unsafe.Sizeof(cronlist)
		So(cl, ShouldBeGreaterThan, 2)
		So(err, ShouldBeNil)
	})
}

func TestTransact(t *testing.T) {
	Convey("测试Transact", t, func() {
		So(Transact(), ShouldBeNil)
	})
}

func BenchmarkInsert(b *testing.B) {
	cronlist := Cronlist{
		Name: "Iversion",
		Type: 2,
	}
	for i := 0; i < b.N; i++ {
		Insert(cronlist)
	}
}

func BenchmarkUpdate(b *testing.B) {
	cronlist := Cronlist{
		Name: "jream",
		Type: 12,
	}
	for i := 0; i < b.N; i++ {
		if i > 200 {
			Update(cronlist, []uint64{uint64(i)})
		}
	}
}
