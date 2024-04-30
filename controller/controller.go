package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})
}
