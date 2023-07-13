package utils

import "encoding/base64"

func EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
