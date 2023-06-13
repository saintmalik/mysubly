package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saintmalik/mysubly/configs"
	"github.com/saintmalik/mysubly/controllers"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()
	router.POST("/api/subs/create", controllers.CreateSubs)
	router.PUT("/api/subs/:subId", controllers.EditSubs)
	router.DELETE("/api/subs/:subId", controllers.DeleteSubs)
	router.GET("/api/subs/:subId", controllers.GetASub)
	router.GET("/api/subs", controllers.GetSubs)

	router.Run("localhost: 3010")
}
