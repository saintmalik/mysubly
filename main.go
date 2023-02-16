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
	router.POST("/api/subs", controllers.CreateSubs)

	router.Run("localhost: 3010")
}
