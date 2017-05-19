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

// Kode fra https://gist.github.com/iwanbk/2295233, https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go -
// - https://systembash.com/a-simple-go-tcp-server-and-tcp-client/ og http://stackoverflow.com/questions/27176523/udp-in-golang-listen-not-a-blocking-call delvis gjennbrukt her.
