// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// "Ringer" serveren.
	conn, _ := net.Dial("tcp", "158.39.77.29:8011")
	for {
		//Les fra standard-in
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Tekst til server: ")
		text, _ := reader.ReadString('\n')
		// Beskjeden som sendes
		fmt.Fprintf(conn, text+"\n")
		//Bekreftelsen
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server: " + message)
	}
}
