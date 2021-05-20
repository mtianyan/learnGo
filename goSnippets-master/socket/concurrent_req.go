package main

import (
	"fmt"
	"net"
	"time"
)

func Client(ipnport string) {
	addr, _ := net.ResolveTCPAddr("tcp", ipnport)
	conn, _ := net.DialTCP("tcp", nil, addr)

	data := make([]byte, 256)
	conn.Read(data)
	fmt.Println(string(data))
}

func Server(port string) {
	addr, _ := net.ResolveTCPAddr("tcp", port)
	listener, _ := net.ListenTCP("tcp", addr)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handle(conn)
		}
	}()
}

func handle(conn net.Conn) {
	defer conn.Close()
	timestamp := time.Now().String()
	conn.Write([]byte(timestamp))
	conn.Close()
}

func main() {
	Server(":9090")
	for i := 0; i < 10; i++ {
		go func() {
			Client("127.0.0.1:9090")
		}()
	}
	select {}
}
