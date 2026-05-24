package main

import (
	"fmt"
	"net"
	"strings"
)

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

		if value.typ != "array" {
			fmt.Println("Invalid request, array expected")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, array > 0 expected")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
