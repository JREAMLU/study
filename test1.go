package main

import (
	"fmt"
)

var Chunk = 5

func main() {
	rid := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
	rids := SliceChunkString(rid, Chunk)
	fmt.Println(len(rids))
	for _, registrationID := range rids {
		fmt.Println(len(registrationID))
		fmt.Println(registrationID)
	}
}

func SliceChunkString(slice []string, size int) (chunkslice [][]string) {
	size1 := len(slice) / size
	if size == 0 || len(slice)%size > 0 {
		size1++
	}

	chunkSize := (len(slice) + size1 - 1) / size1

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunkslice = append(chunkslice, slice[i:end])
	}

	return chunkslice
}
