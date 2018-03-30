package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var tpcConn *net.TCPConn

func main() {
	websocketAddr := "mbgows.plu.cn:8805"
	u := url.URL{Scheme: "ws", Host: websocketAddr, Path: "/", RawQuery: "room_id=154275&group=0&connType=1"}

	fmt.Println("++++++++++++: ", u.String())
	websocket.DefaultDialer.HandshakeTimeout = time.Second * 5
	websocket.DefaultDialer.Proxy = http.ProxyURL(&url.URL{
		Scheme: "http",
		Host:   "27.46.74.27:9999",
		Path:   "/",
	})

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	fmt.Println("++++++++++++: ", c.LocalAddr())
	fmt.Println("++++++++++++: ", c.RemoteAddr())
	defer func() {
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "88"))
		c.Close()
	}()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}

}
