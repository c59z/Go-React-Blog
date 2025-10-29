package main

import (
	"blog-go/utils"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func main() {
	// keyLength := 32
	// randomKey, err := GenerateRandomKey(keyLength)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Randomly generated key:", randomKey)
	pass := utils.BcryptHash("123qwe")
	fmt.Println(pass)
}
