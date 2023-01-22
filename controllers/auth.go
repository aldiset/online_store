package controllers

import (
	"fmt"
	"net/http"
	"online_store/middleware"
	"online_store/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var data UserCreate
	Logger(c)
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
	user.UpdateAt = time.Now()

	_, err := CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	Logger(c)
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Username = data.Username
	user.Password = data.Password

	token, err := LoginCheck(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password isincorrect!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(c *gin.Context) {
	user_id, err := middleware.ExtractTokenID(c)
	Logger(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user_id)
	user, err := GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": UserResponse{
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email}})
}

func VerifyPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func SaveToken(id uint, token string) {
	models.DB.Model(models.User{}).Where("id = ?", id).Update(models.User{Token: token, UpdateAt: time.Now()})
}

func LoginCheck(username, password string) (string, error) {
	var err error

	user := models.User{}

	err = models.DB.Model(models.User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := middleware.GenerateToken(user.Id)
	SaveToken(user.Id, token)

	if err != nil {
		return "", err
	}
	return token, nil
}

func Logout(c *gin.Context) {
	var user models.User
	token := middleware.ExtractToken(c)
	if err := models.DB.First(&user, "token = ?", token).Error; err != nil {
		c.JSON(http.StatusBadGateway, Response{
			Status:  "Failed",
			Message: "Failed",
			Data:    "Wrong Request",
		})
		return
	}
	fmt.Println(token)
	models.DB.Model(models.User{}).Where("token = ?", token).Update(models.User{Token: "-", UpdateAt: time.Now()})
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    "User Logout",
	})
}
