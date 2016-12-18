package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("content: ", string(buf[0:]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
