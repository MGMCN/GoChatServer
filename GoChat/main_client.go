package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main(){
	var Name string
	fmt.Println("Please enter your chat name: ")
	fmt.Scanln(&Name)
	Name = Name + "\n"
	//fmt.Println(Name)

	conn, err := net.Dial("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte(Name))
	done := make(chan struct{})

	// output terminal
	go func() {
		io.Copy(os.Stdout,conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn,os.Stdin) //在这儿卡住一直等stdin输入

	conn.Close()
	<-done //就算我们按下ctrl+c也会在这儿等done同步,除非conn.close()执行否则就会一直等着
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}