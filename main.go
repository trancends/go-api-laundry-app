package main

import (
	"challenge-goapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterCustomerRoutes(router)
	router.Run(":8080")
}
