package routes

import (
	"challenge-goapi/controllers"
	"challenge-goapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine) {
	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("", controllers.GetAllCustomer)
		customerGroup.GET("/:id", controllers.GetCustomerById)
		customerGroup.POST("", controllers.CreateCustomer)
		customerGroup.PUT("/:id", controllers.UpdateCustomerById)
		customerGroup.DELETE("/:id", controllers.DeleteCustomer)
	}
}

func RegisterProductRoutes(r *gin.Engine) {
	// r.GET("/products", controllers.GetAllProduct)
	productGroup := r.Group("/products")
	{
		productGroup.GET("", controllers.GetAllProduct)
		productGroup.GET("/:id", controllers.GetProductById)
		productGroup.POST("", controllers.CreateProduct)
		productGroup.PUT("/:id", controllers.UpdateProductById)
		productGroup.DELETE("/:id", controllers.DeleteProduct)
	}
}

func RegisterTransactionRoutes(r *gin.Engine) {
	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.GET("/:id", controllers.GetTransactionByID)
		transactionGroup.GET("", controllers.GetTransaction)
		transactionGroup.Use(middlewares.TransactionMiddelware)
		transactionGroup.POST("", controllers.CreateTransaction)
	}
}
