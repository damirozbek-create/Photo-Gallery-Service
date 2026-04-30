package routes

import (
	"photo-gallery/controllers"
	"photo-gallery/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 🔥 ПУБЛИЧНЫЕ РОУТЫ (без токена)
	r.GET("/photos", controllers.GetPhotos)
	r.GET("/photos/:id", controllers.GetPhotoByID)

	// 🔒 ЗАЩИЩЕННЫЕ РОУТЫ
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/photos/upload", controllers.UploadPhoto)
		auth.DELETE("/photos/:id", controllers.DeletePhoto)
		auth.GET("/profile", controllers.GetProfile)
		auth.PUT("/profile", controllers.UpdateProfile)
	}
}
