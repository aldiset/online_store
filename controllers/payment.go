package controllers

import (
	"net/http"
	"online_store/models"

	"github.com/gin-gonic/gin"
)

type PaymentInput struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func CreatePayment(c *gin.Context) {
	var input PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment := models.Category{}
	payment.Name = input.Name
	payment.Code = input.Code

	err := models.DB.Create(&payment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "Success",
		Message: "Success",
		Data:    payment,
	})
}

func GetAllPayment(c *gin.Context) {
	var payment []models.PaymentMethod
	models.DB.Find(&payment)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    payment,
	})
}

func GetPaymentById(c *gin.Context) {
	var payment models.PaymentMethod
	id := c.Param("id")
	if err := models.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "payment id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    payment,
	})
}

func UpdatePayment(c *gin.Context) {
	var payment models.PaymentMethod
	var input PaymentInput
	var updateinput models.PaymentMethod
	id := c.Param("id")

	if err := models.DB.First(&payment, id).Error; err != nil {
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

	updateinput.Name = input.Name
	updateinput.Code = input.Code

	models.DB.Model(&payment).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    payment,
	})
}

func DeletePayment(c *gin.Context) {
	var payment models.PaymentMethod
	id := c.Param("id")
	if err := models.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	models.DB.Delete(&payment)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
