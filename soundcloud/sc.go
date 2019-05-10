package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/JREAMLU/j-kit/http"
)

const (
	// frist row
	url = `https://api-v2.soundcloud.com/stream/users/35228969?client_id=5IHUoTCYwQmJR7RbijX9OigWp2zCoiyC&limit=15&offset=0&linked_partitioning=1&app_version=1557315749&app_locale=en`
	// next_href + tail
	tail = `&client_id=5IHUoTCYwQmJR7RbijX9OigWp2zCoiyC&app_version=1557315749&app_locale=en`
	// local proxy ss
	locProxy = `http://127.0.0.1:1087`
	// filename
	filename = `urls`
	// max size
	max = 2
	// interval ms
	interval = 1000
)

var httpClient *http.Requests

func main() {
	initHTTPClient()
	run()
}

func run() {
	u := url
	// loop pageIndex pageSize
	for i := 0; i < max; i++ {
		d, err := getSoundClouds(u)
		if err != nil {
			fmt.Println("++++++++++++: url", url)
			panic(err)
			// continue
		}

		ts, href, err := parseData(d)
		if err != nil {
			fmt.Println("++++++++++++: url", url)
			panic(err)
		}

		for _, v := range ts {
			err := writeFile(filename, v)
			if err != nil {
				panic(err)
			}
			fmt.Println("++++++++++++: 完成", i)
		}

		u = fmt.Sprintf("%s%s", href, tail)
		time.Sleep(interval * time.Millisecond)
	}
}

func initHTTPClient() {
	httpClient = http.NewRequests(nil)
	httpClient.SetTimeout(30)
	httpClient.SetProxy(locProxy)
}

func getSoundClouds(url string) (string, error) {
	resp, err := httpClient.RequestCURLNotTrace(
		context.Background(),
		"GET",
		url,
		map[string]string{},
		``,
		nil,
	)

	if err != nil {
		return "", err
	}

	return resp.Body, nil
}

func parseData(d string) ([]string, string, error) {
	var tracks Tracks
	err := json.Unmarshal([]byte(d), &tracks)
	if err != nil {
		return nil, "", err
	}

	var ts []string

	for _, v := range tracks.Collection {
		switch v.Type {
		case "playlist", "playlist-repost":
			ts = append(ts, v.Playlist.PermalinkURL)
			for _, vv := range v.Playlist.Tracks {
				if vv.PermalinkURL != "" {
					ts = append(ts, vv.PermalinkURL)
				}
			}
		case "track", "track-repost":
			ts = append(ts, v.Track.PermalinkURL)
		}
	}

	return ts, tracks.NextHref, nil
}

func writeFile(filename, content string) error {
	err := checkFile(filename)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(fmt.Sprintf("%s\n", content))

	return nil
}

// check
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}

	return nil
}
