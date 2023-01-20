package controllers

import (
	"net/http"
	"online_store/models"

	"github.com/gin-gonic/gin"
)

type CartInput struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Amount    int `json:"amount"`
}

func CreateCart(c *gin.Context) {
	var input CartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart := models.Cart{}
	cart.ProductID = input.ProductID
	cart.Quantity = input.Quantity
	cart.Amount = input.Amount

	err := models.DB.Create(&cart).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Message: "success",
		Data:    cart,
	})
}

func GetAllCart(c *gin.Context) {
	var cart []models.Cart
	models.DB.Find(&cart)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    cart,
	})
}

func GetCartById(c *gin.Context) {
	var cart models.Cart
	id := c.Param("id")
	if err := models.DB.First(&cart, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "Cart id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    cart,
	})
}

func UpdateCart(c *gin.Context) {
	var cart models.Cart
	var input CartInput
	var updateinput models.Cart
	id := c.Param("id")

	if err := models.DB.First(&cart, id).Error; err != nil {
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

	updateinput.ProductID = input.ProductID
	updateinput.Quantity = input.Quantity
	updateinput.Amount = input.Amount

	models.DB.Model(&cart).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    cart,
	})
}

func DeleteCart(c *gin.Context) {
	var cart models.Cart
	id := c.Param("id")
	if err := models.DB.First(&cart, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	models.DB.Delete(&cart)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
