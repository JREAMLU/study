package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	a := "1a"
	aint, err := AssignGetInt64(a)
	if err == nil {
		fmt.Println("int64: ", aint)
	}

	b := "1.23"
	bfloat, err := AssignGetFloat64(b)
	if err == nil {
		fmt.Println("float64: ", bfloat)
	}
}

func AssignGetInt64(str string) (int64, error) {
	stri, _ := strconv.ParseInt(str, 10, 64)
	strs := strconv.FormatInt(stri, 10)
	if str != strs {
		return stri, errors.New("not equals")
	}
	return stri, nil
}

func AssignGetFloat64(str string) (float64, error) {
	strf, _ := strconv.ParseFloat(str, 64)
	strs := strconv.FormatFloat(strf, 'f', -1, 64)
	if str != strs {
		return strf, errors.New("not equals")
	}
	return strf, nil
}
