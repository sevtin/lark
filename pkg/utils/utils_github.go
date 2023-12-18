package utils

import "strings"

func GetRepository(url string) (fullName string, owner string, repo string) {
	var (
		arr []string
	)
	arr = strings.Split(url, "https://github.com/")
	if len(arr) != 2 {
		return
	}
	fullName = arr[1]
	fullName = strings.ReplaceAll(fullName, " ", "")
	if strings.HasSuffix(arr[1], "/") {
		fullName = fullName[0 : len(fullName)-1]
	}
	// 1、获取仓库信息
	arr = strings.Split(fullName, "/")
	if len(arr) != 2 {
		return
	}
	owner = arr[0]
	repo = arr[1]
	return
}

func UnpackFullName(fullName string) (owner string, repo string) {
	var (
		arr []string
	)
	arr = strings.Split(fullName, "/")
	if len(arr) != 2 {
		return
	}
	owner = arr[0]
	repo = arr[1]
	return
}
