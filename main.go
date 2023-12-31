package main

import (
	"challenge-goapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterCustomerRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterTransactionRoutes(router)
	router.Run(":8080")
}
