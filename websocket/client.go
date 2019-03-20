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
	go conn(1)
	conn(2)
}

func conn(id int64) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8811", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func() {
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "88"))
		c.Close()
	}()

	// guid, _ := uuid.Generate()
	// content := fmt.Sprintf("%s-%s", guid, "client")
	content := fmt.Sprintf("%s-%d", "client", id)
	err = c.WriteMessage(websocket.TextMessage, []byte(content))
	fmt.Printf("err: %v \n", err)

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read err:", err)
				return
			}
			log.Printf("uid: %d, recv: %s", id, message)
		}
	}()

	time.Sleep(time.Second * 5)
}
