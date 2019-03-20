package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// conns save WebSocket *Conn
// other goroutine send message
var conns map[string]*websocket.Conn
var rwMux sync.RWMutex

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	fmt.Println("++++++++++++: ", c.LocalAddr())
	fmt.Println("++++++++++++: ", c.RemoteAddr())
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			delete(conns, string(message))
			break
		}

		sendMessage := fmt.Sprintf("%s %s", "server", string(message))
		log.Printf("recv: %d, %s", mt, message)
		err = c.WriteMessage(mt, []byte(sendMessage))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}

func main() {
	conns = make(map[string]*websocket.Conn, 2)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe("127.0.0.1:8811", nil))
}
