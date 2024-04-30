package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/PageList/routers"
)

func main() {

	router := gin.Default()

	routers.SetupRouter(router)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
