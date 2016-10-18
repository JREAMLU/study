package main

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

const base64GifPixel = "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="

func respHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("=======", r.Form)
	w.Header().Set("Content-Type", "image/gif")
	output, _ := base64.StdEncoding.DecodeString(base64GifPixel)
	io.WriteString(w, string(output))
}

func main() {
	http.HandleFunc("/u.gif", respHandler)
	http.ListenAndServe(":8086", nil)
}
