package utils

import (
	"fmt"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	var (
		key        = []byte("1234567890123456")
		plaintext  = []byte("1234567890123456")
		ciphertext []byte
		buf        []byte
		err        error
	)
	ciphertext, err = AesEncrypt(key, plaintext)
	if err != nil {
		return
	}
	buf, err = AesDecrypt(key, ciphertext)
	if err != nil {
		return
	}
	fmt.Println(string(buf))
}
