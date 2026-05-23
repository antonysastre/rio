package main

import (
	"fmt"
	"net"
)

func print(s string) {
	fmt.Println(s)
}

func main(){
	fmt.Println("Listening on port :6379")

	srvr, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := srvr.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// input := "$5\r\nMando\r\n"
	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)
		conn.Write([]byte("+OK\r\n"))
	}
}
