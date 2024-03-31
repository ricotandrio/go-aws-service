package routers

import (
	"github.com/gin-gonic/gin"
	"go-service-aws/services"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound, 
			"message": "Page not found",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Go is running ...",
		})
	})

	router.GET("/customers", services.GetAllCustomers)
	router.POST("/customers", services.AddCustomer)
	router.DELETE("/customers/:customer_id", services.DeleteCustomer)
	
	return router
}
