package main

import (
	"net"
	"os"
	"fmt"
	//"./Crypt"
)

var key = []byte("TestkeyHei")

func main() {
	str := "Møte fr 5.5 14:45 Flåklypa"
	servAddr := "158.37.63.180:8009"
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

conn, _ := net.DialTCP("tcp", nil, tcpAddr)

//strEcho, _ := Crypt.AesEncrypt([]byte(str), key)
_, _ = conn.Write([]byte(str))

println("write to server = ", str)

//reply
reply := make([]byte, 1024)
conn.Read(reply)
//msg, _ := Crypt.AesDecrypt([]byte(reply), key)
//fmt.Println("reply from server=", (reply))
//fmt.Printf("%s", msg)
fmt.Printf("%s", reply)
conn.Close()
}