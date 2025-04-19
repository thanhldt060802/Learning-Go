package main

import (
	"go-api-with-gin/config"
	"go-api-with-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDB()
	r := gin.Default()
	routes.RegisterRoutes(r, db)
	r.Run(":8080")

}
