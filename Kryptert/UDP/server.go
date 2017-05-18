// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main
import (
	"fmt"
	"net"
	"./Crypt"
)



func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

const key = "HeiIS105"

func main() {
	listenAndReceive()
}

func listenAndReceive(){
	p := make([]byte, 2048)

	addr := net.UDPAddr{
		Port: 8009,
		IP: net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		var msg, _ = Crypt.AesDecrypt([]byte(p), []byte(key))
		fmt.Printf("Read a message from %v %s \n", remoteaddr, string(msg))

		if err !=  nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}
