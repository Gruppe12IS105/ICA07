package client

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func CBCDecrypter(encode string, keyString []byte) []byte {
	key := []byte(keyString)
	ciphertext, _ := hex.DecodeString(encode)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext
}
