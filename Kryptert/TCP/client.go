package main
import (
	"net"
	"os"
	"./Crypt"
)

var key := "TestkeyHei"

func main() {
	str := "Møte fr 5.5 14:45 Flåklypa"
	servAddr := "158.37.63.180:8009"
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
	}
	strEcho, _ := Crypt.AesEncrypt(str, key)
	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	println("write to server = ", str)

	reply := make([]byte, 1024)
	rawMsg , err := conn.Read(reply)
	msg := Crypt.AesDecrypt([]byte(rawMsg, key))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(msg))

	conn.Close()
}