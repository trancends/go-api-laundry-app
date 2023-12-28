package routes

import (
	"challenge-goapi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine) {
	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("/", controllers.GetAllCustomer)
		customerGroup.POST("/", controllers.CreateCustomer)
		customerGroup.PUT("/:id", controllers.UpdateCustomerById)
	}
}
