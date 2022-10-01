package router

import (
	"backend/internal/controller"
	"backend/internal/defines"
	"backend/internal/service"
	"backend/internal/wsserver"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// WebSockets Server
	wsConfig := &wsserver.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wss := wsserver.New(wsConfig)

	// Services
	crSrv := service.NewChatRoomsService(wss)

	// Controllers
	crCtrl := controller.NewChatRoomsController(crSrv)

	// Endpoints
	r.GET(defines.EndpointChatRoomConnect, crCtrl.Connect)
}
