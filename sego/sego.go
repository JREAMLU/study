package main

import (
	"fmt"
	"os"

	"github.com/huichen/sego"
)

func main() {
	// load dictionary
	goPath := os.Getenv("GOPATH")
	dir := fmt.Sprintf("%v/src/%v", goPath, "github.com/huichen/sego/data/dictionary.txt")

	var segmenter sego.Segmenter
	segmenter.LoadDictionary(dir)

	// segmenter
	text := []byte("中华人民共和国中央人民政府")
	segments := segmenter.Segment(text)

	// handle searchMode
	fmt.Println(sego.SegmentsToString(segments, true))
}
