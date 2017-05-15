package main
import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "158.37.63.180:8009")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Møte Fr 5.5 14:45 Flåklypa")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}