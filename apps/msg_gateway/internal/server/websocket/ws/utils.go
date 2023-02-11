package ws

import (
	"strconv"
)

func int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func int32ToStr(val int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(val)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func clientKey(uid int64, platform int32) (key string) {
	return int64ToStr(uid) + "-" + int32ToStr(platform)
}
