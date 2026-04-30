package controllers

import (
	"net/http"

	"photo-gallery/database"
	"photo-gallery/models"

	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	path := "uploads/" + file.Filename
	c.SaveUploadedFile(file, path)

	userID := c.MustGet("user_id").(uint) // 🔥 правильно

	photo := models.Photo{
		UserID: userID,
		URL:    path,
	}

	database.DB.Create(&photo)

	c.JSON(200, photo)
}

func GetPhotos(c *gin.Context) {

	var photos []models.Photo

	result := database.DB.Find(&photos)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if len(photos) == 0 {
		c.JSON(200, []models.Photo{}) // 🔥 пустой массив, а не null
		return
	}

	c.JSON(200, photos)
}
func GetPhotoByID(c *gin.Context) {

	id := c.Param("id")

	var photo models.Photo

	result := database.DB.First(&photo, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func DeletePhoto(c *gin.Context) {

	id := c.Param("id")

	var photo models.Photo

	result := database.DB.First(&photo, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	database.DB.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
