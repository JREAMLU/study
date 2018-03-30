package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var tpcConn *net.TCPConn

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8811", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func() {
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "88"))
		c.Close()
	}()

	err = c.WriteMessage(websocket.TextMessage, []byte("123"))
	fmt.Printf("err: %v \n", err)

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read err:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	time.Sleep(time.Second * 1)
}
