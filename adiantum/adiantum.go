package main

import (
	"fmt"

	"lukechampine.com/adiantum"
)

func main() {
	key := make([]byte, 64) // in practice, read this from crypto/rand
	cipher := adiantum.New20(key)
	tweak := make([]byte, 16) // can be any length
	plaintext := []byte("Hello, world! LUj")
	ciphertext := cipher.Encrypt(plaintext, tweak)
	fmt.Println("++++++++++++: ", string(ciphertext), len(ciphertext))
	recovered := cipher.Decrypt(ciphertext, tweak)
	fmt.Println("++++++++++++: ", string(recovered))
}
