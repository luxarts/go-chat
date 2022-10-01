package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTMLController interface {
	Index(ctx *gin.Context)
}

type htmlController struct {
}

func NewHTMLController() HTMLController {
	return &htmlController{}
}

func (ctrl *htmlController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
