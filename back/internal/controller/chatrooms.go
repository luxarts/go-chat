package controller

import (
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ChatRoomsController interface {
	Connect(ctx *gin.Context)
}

type chatRoomsController struct {
	srv service.ChatRoomsService
}

func NewChatRoomsController(srv service.ChatRoomsService) ChatRoomsController {
	return &chatRoomsController{
		srv: srv,
	}
}

func (ctrl *chatRoomsController) Connect(ctx *gin.Context) {
	// Get writer and reader from context
	w := ctx.Writer
	r := ctx.Request

	// Service call
	ctrl.srv.Connect(w, r)
}
