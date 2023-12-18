package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SixteenMD5(str string) string {
	return MD5(str)[8:24]
}

func AesEncrypt(key, plaintext []byte) (buf []byte, err error) {
	var (
		block cipher.Block
	)
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	buf = ciphertext
	return
}

func AesDecrypt(key, ciphertext []byte) (buf []byte, err error) {
	var (
		block cipher.Block
	)
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(ciphertext) < aes.BlockSize {
		return
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	buf = ciphertext
	return
}

func RC4Encrypt(str string, key []byte) string {
	plaintext := []byte(str)
	cipher, _ := rc4.NewCipher(key)
	out := make([]byte, len(plaintext))
	cipher.XORKeyStream(out, plaintext)
	return hex.EncodeToString(out)
}

func RC4Decrypt(str string, key []byte) string {
	ciphertext, _ := hex.DecodeString(str)
	if len(ciphertext) == 0 {
		return ""
	}
	cipher, _ := rc4.NewCipher(key)
	out := make([]byte, len(ciphertext))
	cipher.XORKeyStream(out, ciphertext)
	return string(out)
}
