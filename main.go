package main

import (
	"io"
	"os"
	"tourism/controllers"
	"tourism/db"
	"tourism/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDb()
}
func setupLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.Auth())
	server.POST("/details", controllers.DetailsCreate)
	server.GET("/details", controllers.GetAllDetails)
	server.GET("/details/:id", controllers.FindOneDetail)
	server.PATCH("/details/:id", controllers.UpdateDetails)
	server.DELETE("/details/:id", controllers.DetailsDelete)
	server.Run()
}
