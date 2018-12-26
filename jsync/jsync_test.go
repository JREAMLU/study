package jsync

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	Convey("Get()", t, func() {
		c := Cache{
			data: []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99, 200},
		}

		d := c.Get()
		fmt.Println("++++++++++++: ", d)
	})
}

func BenchmarkGet(b *testing.B) {
	d := make([]int, 100)
	for i := 0; i < 100; i++ {
		d[i] = 101 + i
	}

	c := Cache{
		data: d,
	}

	Convey("bench Get()", b, func() {
		for i := 0; i < b.N; i++ {
			c.Get()
		}
	})
}

func TestGetC(t *testing.T) {
	Convey("GetC()", t, func() {
		c := Cache{
			data: []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99, 200},
		}
		data := make([]int, 0, 10)

		next := make(chan struct{})
		for i := range c.GetC(next) {
			data = append(data, i)
			if len(data) >= 10 {
				close(next)
				break
			}

			next <- struct{}{}
		}

		fmt.Println("++++++++++++: ", data)
	})
}

func BenchmarkGetC(b *testing.B) {
	d := make([]int, 100)
	for i := 0; i < 100; i++ {
		d[i] = 101 + i
	}

	c := Cache{
		data: d,
	}

	data := make([]int, 0, 10)

	Convey("bench GetC()", b, func() {
		for i := 0; i < b.N; i++ {
			next := make(chan struct{})
			for i := range c.GetC(next) {
				data = append(data, i)
				if len(data) >= 100 {
					close(next)
					break
				}

				next <- struct{}{}
			}
		}
	})
}
