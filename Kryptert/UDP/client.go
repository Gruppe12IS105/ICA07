package main
import (
	"fmt"
	"net"
	"bufio"
	"crypto/sha256"
)

const melding = "Møte fr 5.5 14:45 Flåklypa"

func main() {
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "158.37.63.180:8009")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	cryptMelding := sha256.New()
	cryptMelding.Write([]byte(melding))

	fmt.Fprintf(conn, string(cryptMelding.Sum(nil)))
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}