package utils

import (
	"lark/pkg/proto/pb_enum"
	"strings"
)

func GetMsgGatewayServer(in string) (server string, serverId int64, port int) {
	defer func() {
		if server == "" {
			server = "lark_msg_gateway_server"
			port = 7301
			serverId = 1
		}
	}()
	arr := strings.Split(in, ":")
	if len(arr) != 3 {
		return
	}
	server = arr[0]
	serverId, _ = ToInt64(arr[1])
	port, _ = ToInt(arr[2])
	return
}

func GetServer(key string) (name string, port string) {
	var (
		index int
		str   string
	)
	index = strings.LastIndex(key, "///")
	str = key[index+len("///"):]
	index = strings.Index(str, "/")
	name = str[:index]
	index = strings.LastIndex(key, ":")
	port = key[index+len(":"):]
	return
}

func GetServerId(val int64) (ios, android, mac, windows, web int64) {
	ios = val / 10000
	android = val / 1000 % 10
	mac = val / 100 % 10
	windows = val / 10 % 10
	web = val % 10
	return
}

func NewServerId(old int64, serverId int64, platform pb_enum.PLATFORM_TYPE) int64 {
	ios := old / 10000
	android := old / 1000 % 10
	mac := old / 100 % 10
	windows := old / 10 % 10
	web := old % 10
	switch platform {
	case pb_enum.PLATFORM_TYPE_IOS:
		ios = serverId
	case pb_enum.PLATFORM_TYPE_ANDROID:
		android = serverId
	case pb_enum.PLATFORM_TYPE_MAC:
		mac = serverId
	case pb_enum.PLATFORM_TYPE_WINDOWS:
		windows = serverId
	case pb_enum.PLATFORM_TYPE_WEB:
		web = serverId
	}
	return ios*10000 + android*1000 + mac*100 + windows*10 + web
}
