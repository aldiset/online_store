package controllers

import (
	"net/http"
	"online_store/models"

	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	CartID            uint   `json:"name"`
	PaymentMethodCode string `json:"code"`
}

func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction := models.Transaction{}
	transaction.CartID = input.CartID
	transaction.PaymentMethodCode = input.PaymentMethodCode

	err := models.DB.Create(&transaction).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Message: "success",
		Data:    transaction,
	})
}

func GetAllTransaction(c *gin.Context) {
	var transaction []models.Transaction
	models.DB.Find(&transaction)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    transaction,
	})
}

func GetTransactionById(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")
	if err := models.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "transaction id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    transaction,
	})
}

func UpdateTransaction(c *gin.Context) {
	var transaction models.Transaction
	var input TransactionInput
	var updateinput models.Transaction
	id := c.Param("id")

	if err := models.DB.First(&transaction, id).Error; err != nil {
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

	updateinput.CartID = input.CartID
	updateinput.PaymentMethodCode = input.PaymentMethodCode

	models.DB.Model(&transaction).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    transaction,
	})
}

func DeleteTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")
	if err := models.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	models.DB.Delete(&transaction)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
