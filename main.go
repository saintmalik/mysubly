package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/saintmalik/mysubly/configs"
	"github.com/saintmalik/mysubly/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/api/subs/:subid", controllers.SubById)
	router.DELETE("/api/subs/:subid", controllers.DeleteSub)
	router.GET("/api/subs", controllers.Sub)
	router.POST("/api/subs", controllers.CreateSub)

	router.Run(":3010")
}
