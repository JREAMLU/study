package main

import (
	"context"
	"fmt"

	"github.com/JREAMLU/j-kit/http"
)

// frist row
const url = `https://api-v2.soundcloud.com/stream/users/35228969?client_id=5IHUoTCYwQmJR7RbijX9OigWp2zCoiyC&limit=15&offset=0&linked_partitioning=1&app_version=1557315749&app_locale=en`

// next_href + tail
const tail = `&client_id=5IHUoTCYwQmJR7RbijX9OigWp2zCoiyC&app_version=1557315749&app_locale=en`

var httpClient *http.Requests

func main() {
	initHTTPClient()
	run()
}

func run() {
	// loop pageIndex pageSize
}

func initHTTPClient() {
	req := http.NewRequests(nil)
	req.SetTimeout(30)
}

func getSoundClouds(url string) ([]string, error) {
	resp, err := httpClient.RequestCURL(
		context.Background(),
		"GET",
		url,
		map[string]string{},
		``,
		nil,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("++++++++++++: ", resp)

	return nil, nil
}

func writeFile(filename string) bool {
	return false
}
