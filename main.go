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
	router.POST("/api/subs/create", controllers.CreateSub)
	router.PUT("/api/subs/:subId", controllers.EditSub)
	router.DELETE("/api/subs/:subId", controllers.DeleteSub)
	router.GET("/api/subs/:subId", controllers.SubById)
	router.GET("/api/subs", controllers.GetSubs)
	router.POST("/api/signup", controllers.SignUp)

	router.Run("localhost: 3010")
}
