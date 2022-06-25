package main

import (
	"GoChatServer/GoChat/Client"
	"fmt"
	"log"
	"net"
	"os"
)

func main(){
	var Name string
	fmt.Println("Please enter your chat name: ")
	fmt.Scanln(&Name)
	Name = Name + "\n"

	conn, err := net.Dial("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte(Name))

	go Client.TerminalOutput(conn) // output terminal
	Client.MustCopy(conn,os.Stdin) //在这儿卡住一直等stdin输入

	conn.Close()
	Client.WaitDone() //就算我们按下ctrl+c也会在这儿等done同步,除非conn.close()执行否则就会一直等着
}