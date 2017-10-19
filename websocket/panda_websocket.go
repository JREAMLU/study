package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	websocketAddr := "riveng28-sht.gw.riven.panda.tv:8080"
	u := url.URL{Scheme: "wss", Host: websocketAddr, Path: "/"}

	fmt.Println("++++++++++++: ", u.String())

	wd := websocket.Dialer{
		HandshakeTimeout: time.Second * 5,
		TLSClientConfig: &tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: true,
		},
	}

	c, _, err := wd.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

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
