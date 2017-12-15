package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tt", test)

	log.Println("启动http服务, 端口:7890")
	err := http.ListenAndServe(":7890", nil)
	if err != nil {
		log.Fatal("ListenAndServe live: ", err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	str := "hello docker"
	fmt.Println("++++++++++++: ", str)
	fmt.Fprintf(w, str)
}
