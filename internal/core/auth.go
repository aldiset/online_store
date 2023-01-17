package core

import (
	"net/http"

	"online_store/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterDataUser struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var data RegisterDataUser

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	user := models.User{}

	user.Fullname = data.Fullname
	user.Username = data.Username
	user.Password = string(password)
	user.Email = data.Email

	_, err := user.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}
