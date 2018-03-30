package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	websocketAddr := "riveng28-sht.gw.riven.panda.tv:8080"
	u := url.URL{Scheme: "wss", Host: websocketAddr, Path: "/"}

	origin := "https://www.panda.tv"
	wsHeaders := http.Header{
		"Origin": {origin},
	}

	wd := websocket.Dialer{
		HandshakeTimeout: time.Second * 5,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Proxy: http.ProxyFromEnvironment,
	}

	c, _, err := wd.Dial(u.String(), wsHeaders)
	if err != nil {
		log.Fatal("-dial:", err)
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

	/*
		websocketAddr := "riveng28-sht.gw.riven.panda.tv:8080"
		u := url.URL{Scheme: "wss", Host: websocketAddr, Path: "/"}

		netConn, err := net.Dial("tcp", u.Host)
		if err != nil {
			log.Fatal("net dial err: ", err)
		}

		wsHeaders := http.Header{
			"Origin": {u.Host},
		}

		wd := websocket.Dialer{
			HandshakeTimeout: time.Second * 5,
			// TLSClientConfig: &tls.Config{
			// 	InsecureSkipVerify: true,
			// },
			Proxy: http.ProxyFromEnvironment,
			NetDial: func(net, addr string) (net.Conn, error) {
				return netConn, nil
			},
		}

		c, _, err := wd.Dial(u.String(), wsHeaders)
		if err != nil {
			log.Fatal("-dial:", err)
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
	*/

	/*
	   var origin = "http://127.0.0.1:8080/"
	   var url = "wss://riveng28-sht.gw.riven.panda.tv:8080/"

	   ws, err := websocket.Dial(url, "", origin)
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Println("++++++++++++: ")
	   message := []byte("hello, world!你好")
	   _, err = ws.Write(message)
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Send: %s\n", message)

	   var msg = make([]byte, 512)
	   m, err := ws.Read(msg)
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Receive: %s\n", msg[:m])

	   ws.Close() //关闭连接
	*/
}
