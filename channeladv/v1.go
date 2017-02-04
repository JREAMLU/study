package main

import (
	jcontext "context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JREAMLU/core/guid"
)

var queue chan int
var workNumber int

func main() {
	queue = make(chan int, 10000)
	workNumber = 5

	go router()

	gid := getID()
	jctx := jcontext.WithValue(jcontext.Background(), "requestID", gid)
	go processor(jctx)

	select {}
}

func router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9999", mux)
	if err != nil {
		log.Fatal("启动错误: ", err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "golang http")
}

// v1
// func upload(w http.ResponseWriter, r *http.Request) {
// 	for i := 0; i < 5; i++ {
// 		uploadToS3()
// 	}
// 	fmt.Fprintf(w, "done")
// 	fmt.Println(getID(), " done ")
// }

// v2
// func upload(w http.ResponseWriter, r *http.Request) {
// 	for i := 0; i < 5; i++ {
// 		go uploadToS3()
// 	}
// 	fmt.Fprintf(w, "done")
// 	fmt.Println(getID(), " done ")
// }

// v3

func upload(w http.ResponseWriter, r *http.Request) {

	gid := getID()
	// jctx := jcontext.WithValue(jcontext.Background(), "requestID", gid)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			queue <- i
		}
	}()

	// go processor(jctx)

	fmt.Fprintf(w, "done")
	fmt.Println(gid, " done ")
}

func processor(jctx jcontext.Context) {
	fmt.Println("gid: ", jctx.Value("requestID"))

	for i := 0; i < workNumber; i++ {
		go func() {
			for {
				select {
				case <-queue:
					fmt.Println("upload:", jctx.Value("requestID"))
					uploadToS3(jctx)
				}
			}
		}()
	}
}

func uploadToS3(jctx jcontext.Context) int {
	num := 3
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Println(jctx.Value("requestID"), " sleep ", num)

	return num
}

func getID() string {
	return guid.NewObjectID().Hex()
}
