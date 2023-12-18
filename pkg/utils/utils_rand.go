package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// 随机数范围为0-9的数字
const numSet = "0123456789"

func Rand(len int) string {
	// 产生一个以当前时间为种子的随机数生成器
	rand.Seed(time.Now().UnixNano())
	// 容器,用于存放结果
	b := make([]byte, len)
	for i := range b {
		// 产生0-9的随机索引
		r := rand.Intn(10)
		// 选择随机数字并存入结果容器
		b[i] = numSet[r]
	}
	// 转换为字符串并输出
	return string(b)
}

func GetVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(900000) + 100000)
}
