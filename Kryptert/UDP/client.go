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
// Kode fra https://gist.github.com/iwanbk/2295233, https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go -
// - https://systembash.com/a-simple-go-tcp-server-and-tcp-client/ og http://stackoverflow.com/questions/27176523/udp-in-golang-listen-not-a-blocking-call delvis gjennbrukt her.
