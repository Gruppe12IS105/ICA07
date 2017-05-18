// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main
import (
	"fmt"
	"net"
	"./Crypt"
	"bufio"
)

const key = "HeiIS105"

func main() {
	var msg = Crypt.AesEncrypt([]byte("Møte fr 5.5 14:45 Flåklypa"), key)
}

func connectAndSend(key, message []byte){

	p :=  make([]byte, 2048)

	conn, err := net.Dial("udp", "158.37.63.180:8009")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	fmt.Fprintf(conn, string(message))
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
