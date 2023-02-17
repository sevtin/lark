package utils

import "encoding/base64"

func DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
