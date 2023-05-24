package main

import (
	"rest-api/config"
	"rest-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// config.ConnectDatabase()
	config.ConnectDatabase2()

	r.GET("/", controller.RootHandler)
	r.GET("/item/:id", controller.ItemParam)
	r.GET("/item", controller.ItemQuery)
	r.POST("/items", controller.ItemPost)
	r.GET("/user", controller.GetAllUser)
	r.POST("/user/login", controller.GetUserLogin)
	r.GET("/vehicle", controller.GetAllVehicle)
	// r.GET("/vehicle2", controller.GetAllVehicle2)

	r.Run(":8888")
}
