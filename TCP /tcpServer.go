// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {

	fmt.Println("Launcher server...")

	ln, _ := net.Listen("tcp", ":8011")

	conn, _ := ln.Accept()

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Beskjed mottatt:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
// Kode fra https://gist.github.com/iwanbk/2295233, https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go -
// - https://systembash.com/a-simple-go-tcp-server-and-tcp-client/ og http://stackoverflow.com/questions/27176523/udp-in-golang-listen-not-a-blocking-call delvis gjennbrukt her.
