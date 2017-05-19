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
