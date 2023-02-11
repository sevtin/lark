package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SixteenMD5(str string) string {
	return MD5(str)[8:24]
}
