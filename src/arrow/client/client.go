package client

import (
	"fmt"
	"net"
	"bufio"
)

func Main(){
	fmt.Println("client Main func")

}

func Clinet(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}