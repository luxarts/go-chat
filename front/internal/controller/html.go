package controller

import (
	"frontend/internal/defines"
	"frontend/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	ctx.HTML(http.StatusOK, "index.html", domain.HTMLData{
		BackendURL: os.Getenv(defines.EnvBackendURL),
	})
}
