package utils

import (
	"math"
	"strconv"
	"strings"
	"time"
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

func GetAge(birthTs int64) int32 {
	return int32(time.Now().Year() - time.Unix(birthTs/1000, 0).Year())
}

func FloatToString(num float64) string {
	// 整数部分
	intPart := int(num)
	// 小数部分
	floatPart := num - float64(intPart)
	var result string
	// 如果有小数
	if floatPart != 0 {
		// 四舍五入到2位小数
		rounded := math.Round(num*100) / 100
		result = strconv.FormatFloat(rounded, 'f', 2, 64)

		// 如果第二位小数为0,只保留1位小数
		if result[len(result)-1] == '0' {
			result = strconv.FormatFloat(rounded, 'f', 1, 64)
		}
	} else {
		// 如果没小数,只返回整数部分
		result = strconv.Itoa(intPart)
	}
	return result
}
