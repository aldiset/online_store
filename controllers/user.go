package controllers

import (
	"errors"
	"net/http"
	"online_store/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Id       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func CreateUser(user *models.User) (*models.User, error) {
	var err error
	err = models.DB.Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		return user, errors.New("User not found!")
	}
	return user, nil
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    user,
	})
}

func GetAllUser(c *gin.Context) {
	var user []models.User
	models.DB.Find(&user)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var input UserCreate
	var updateinput models.User
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	updateinput.Fullname = input.Fullname
	updateinput.Username = input.Username
	updateinput.Password = string(password)
	updateinput.Email = input.Email

	models.DB.Model(&user).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    user,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	models.DB.Delete(&user)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
