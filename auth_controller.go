package controllers

import (
	"net/http"

	"photo-gallery/database"
	"photo-gallery/models"
	"photo-gallery/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var body models.User

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot hash password"})
		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

func Login(c *gin.Context) {

	var body models.User
	var user models.User

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	// find user
	database.DB.Where("email = ?", body.Email).First(&user)

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong email or password"})
		return
	}

	// generate JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetProfile(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
func UpdateProfile(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid data"})
		return
	}

	user.Username = body.Username
	user.Email = body.Email

	database.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "profile updated",
		"user":    user,
	})
}
