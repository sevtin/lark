package utils

import (
	"strconv"
	"strings"
)

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int32ToStr(val int32) string {
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

func IntToStr(val int) string {
	return strconv.Itoa(val)
}

func StrToInt64(str string) (val int64) {
	val, _ = strconv.ParseInt(str, 10, 64)
	return
}

func StrListToInt64List(ins []string) (outs []int64) {
	outs = make([]int64, len(ins))
	if len(ins) == 0 {
		return
	}
	var (
		index int
		v     string
	)
	for index, v = range ins {
		outs[index] = StrToInt64(v)
	}
	return
}

func SplitToInt64List(s string, sep string) (values []int64) {
	arr := strings.Split(s, sep)
	return StrListToInt64List(arr)
}

func Int64ListToStrList(ins []int64) (outs []string) {
	outs = make([]string, len(ins))
	if len(ins) == 0 {
		return
	}
	var (
		index int
		v     int64
	)
	for index, v = range ins {
		outs[index] = Int64ToStr(v)
	}
	return
}
