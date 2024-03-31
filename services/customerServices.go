package services

import (
	"github.com/gin-gonic/gin"
	"go-service-aws/models"
	"go-service-aws/configs"
	"net/http"
	"time"
)

func GetAllCustomers(c *gin.Context) {
	var customers []models.Customers

	showCustomers := configs.GetDB().Find(&customers)

	if showCustomers.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": showCustomers.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Get All Customers Success",
		"data": customers,
	})
}

func AddCustomer(c *gin.Context) {
	var customer models.Customers

	customer.CustomerCreatedAt = time.Now().Format(time.RFC3339)

	bindDataToCustomerModels := c.ShouldBindJSON(&customer)
	if bindDataToCustomerModels != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": bindDataToCustomerModels.Error,
		})
		return
	}

	addDataToDatabase := configs.GetDB().Create(&customer)
	if addDataToDatabase.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": addDataToDatabase.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"message": "Add Customer Success",
		"data": customer,
	})
}

func DeleteCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")

	var customer models.Customers

	deleteDataFromDatabase := configs.GetDB().Where("customer_id = ?", customerID).Delete(&customer)
	if deleteDataFromDatabase.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": deleteDataFromDatabase.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Delete Customer Success",
	})
}