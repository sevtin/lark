package ws

import (
	"go.uber.org/zap"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/middleware"
	"lark/pkg/utils"
)

var (
	wsLog *zap.SugaredLogger
)

type WServer struct {
	port     int
	serverId int
	hub      *Hub
	gin      *xgin.GinServer
}

func NewWServer(wsServer *conf.WsServer, callback MessageCallback) *WServer {
	var (
		ws *WServer
	)
	ws = &WServer{
		port:     wsServer.Port,
		serverId: wsServer.ServerId,
		hub:      NewHub(wsServer.ServerId, callback),
		gin:      xgin.NewGinServer(),
	}
	wsLog = xlog.NewLog(wsServer.Log, wsServer.Name+utils.ToString(wsServer.ServerId))
	ws.gin.Engine.Use(middleware.JwtAuth())
	ws.gin.Engine.GET("/", ws.hub.wsHandler)
	return ws
}

func (ws *WServer) Run() {
	go func() {
		ws.hub.Run()
		ws.gin.Run(ws.port)
	}()
}

func (ws *WServer) SendMessage(uid int64, platform int32, message []byte) (result int32) {
	return ws.hub.SendMessage(uid, platform, message)
}

func (ws *WServer) IsOnline(uid int64, platform int32) (ok bool) {
	return ws.hub.IsOnline(uid, platform)
}

func (ws *WServer) NumberOfOnline() int {
	return ws.hub.NumberOfOnline()
}
