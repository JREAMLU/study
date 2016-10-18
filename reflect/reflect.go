package main

import (
	"encoding/binary"
	"math"
	"reflect"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var i interface{}

	i = 123456789.1469690450

	kind := reflect.TypeOf(i).Kind()

	a := reflect.ValueOf(i).Float()

	j := strconv.FormatFloat(a, 'f', -1, 64)

	spew.Dump(kind, i, a, j)

}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}
