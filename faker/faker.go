package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

type SomeStruct struct {
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
	String   string
	Bool     bool
	SString  []string
	SInt     []int
	SInt8    []int8
	SInt16   []int16
	SInt32   []int32
	SInt64   []int64
	SFloat32 []float32
	SFloat64 []float64
	SBool    []bool
	Struct   AStruct
}
type AStruct struct {
	Number        int64
	Height        int64
	AnotherStruct BStruct
}

type BStruct struct {
	Image string
}

func main() {

	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
