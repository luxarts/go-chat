package router

import (
	"frontend/internal/controller"
	"frontend/internal/defines"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/fonts", "./static/fonts")
	r.LoadHTMLGlob("./static/*.html")

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	htmlCtrl := controller.NewHTMLController()

	r.GET(defines.EndpointIndex, htmlCtrl.Index)
}
