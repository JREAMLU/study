package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	resp := `{"status":"success"}`
	fmt.Fprintf(w, resp)
}

func main() {
	http.HandleFunc("/nihao", sayhelloName)  //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
