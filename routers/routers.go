package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/PageList/controller"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/", controller.ShowIndexPage)
	router.GET("/ping", controller.Ping)
	router.GET("/article/view/:id", controller.ShowArticlePage)
}
