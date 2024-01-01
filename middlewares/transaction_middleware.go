package middlewares

import (
	"challenge-goapi/config"
	"challenge-goapi/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TransactionMiddelware(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var request models.Transaction
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(request)

	// Validate the request data
	if request.BillDate == "" || request.EntryDate == "" || request.FinishDate == "" || request.EmployeeID == "" || request.CustomerID == "" || len(request.BillDetails) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Validate date received from the POST request is in the format "dd-mm-yyyy"
	_, err = time.Parse("02-01-2006", request.BillDate)
	if err != nil {
		// Handle the error if the date string is not in the expected format
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	_, err = time.Parse("02-01-2006", request.EntryDate)
	if err != nil {
		// Handle the error if the date string is not in the expected format
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	_, err = time.Parse("02-01-2006", request.FinishDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	c.Set("data", request)

	c.Next()
}
