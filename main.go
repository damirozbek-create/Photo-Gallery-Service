package main

import (
	"photo-gallery/database"
	"photo-gallery/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 🔥 CORS ДО ВСЕХ РОУТОВ
	r.Use(cors.Default())

	database.Connect()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
