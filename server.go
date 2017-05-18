
package main
import (
	"fmt"
	"net"
	"os"
	"./Crypt"
)

var key = []byte("TestkeyHei")

func main() {
	addr := net.TCPAddr{
		Port: 8009,
		IP:   net.ParseIP("0.0.0.0"),
	}

	ser, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		conn, err := ser.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest (conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	conn.Read(buf)
	//msg, _ := Crypt.AesDecrypt(buf, key)
	// Send a response back to person contacting us.
	//fmt.Println(msg)
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}