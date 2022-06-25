package Server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client chan<- string

var (
	login         = make(chan client)
	logout        = make(chan client)
	messages      = make(chan string)
	userNameTable = make(map[client]string)
)
func Broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <- messages:
			for cli := range clients {
				msgs := strings.Split(msg,":")
				if userNameTable[cli] != msgs[0] {
					cli <- msg
				}
			}
		case cli := <- login:
			clients[cli] = true // 有用 ? hahaha
		case cli := <- logout:
			delete(clients,cli)
			close(cli)
		}
	}
}

func HandleConn(conn net.Conn) {
	ch := make(chan string)
	// write
	go clientWriter(conn,ch) // 将ch与conn的输出联系起来了...

	//incomeip := conn.RemoteAddr().String()
	var incomeUser string
	input := bufio.NewScanner(conn)
	if input.Scan() {
		incomeUser = input.Text()
		//fmt.Println(incomeUser)
	}
	userNameTable[ch] = incomeUser

	ch <- incomeUser + " welcome!\n" // 欢迎加入方
	messages <- incomeUser + "　has joined!\n" // 告知其他人有新人加入了
	login <- ch // 再将新人登录下

	var text string
	for input.Scan() {
		text = input.Text()
		//fmt.Println(text)
		messages <- incomeUser + ": " + text + "\n"
	}

	logout <- ch
	messages <- incomeUser + " has left!\n"
	conn.Close()
}

func clientWriter(conn net.Conn,ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn,msg)
	}
}
