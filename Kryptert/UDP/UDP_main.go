package main

import (
	"./client"
  "fmt"
)

func main() {
	//message := "Møte Fr 5.5 14:45 Flåklypa"
  message := []byte("exampleplaintext")
  key := []byte("example key 1234")

  //encode := "f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded"

fmt.Println("-----------ENCRYPTING-----------")
// Encodes the message
	var encode = client.CBCEncrypter(message, key)
  	fmt.Printf("%x\n", encode)

fmt.Println("\n")

fmt.Println("-----------DECRYPTING-----------")
// Decodes the encoded message
  var decode = client.CBCDecrypter(string(encode), key)
    fmt.Printf("%s\n", decode)
}
