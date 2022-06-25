package main

import (
	"GoChatServer/GoChat/Server"
	"log"
	"net"
)

func main() {
	listener,err := net.Listen("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go Server.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go Server.HandleConn(conn)
	}
}
