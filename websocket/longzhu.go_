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
	websocketAddr := "mbgows.plu.cn:8805"
	u := url.URL{Scheme: "ws", Host: websocketAddr, Path: "/", RawQuery: "room_id=1518057&group=0&connType=1"}

	fmt.Println("++++++++++++: ", u.String())
	// cookieStr := "p1u_id=9912bd1fb9cad1fd4382983c38b00c87e8cbe49a643cdf324e191e29003ed601ed2020e04864f1cb"
	websocket.DefaultDialer.HandshakeTimeout = time.Second * 5
	// wsHeader := http.Header{
	// 	"Origin": {websocketAddr},
	// 	"Cookie": {cookieStr},
	// }

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
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
