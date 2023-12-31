package controllers

import (
	"challenge-goapi/config"
	"challenge-goapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var request models.Request
	err := c.ShouldBind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
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
	} else {
		c.JSON(http.StatusOK, request)
		return
	}

	billDate, _ := time.Parse("2006-01-02", request.BillDate)
	entryDate, _ := time.Parse("2006-01-02", request.EntryDate)
	finishDate, _ := time.Parse("2006-01-02", request.FinishDate)

	// Create a new transaction in the database using a transaction
	tx, err := db.Begin()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	// Insert the transaction record
	var transactionID string
	err = tx.QueryRow("INSERT INTO bill (bill_date, entry_date, finish_date, employee_id, customer_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		billDate, entryDate, finishDate, request.EmployeeID, request.CustomerID).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Insert the bill details
	for _, billDetail := range request.BillDetails {
		_, err = tx.Exec("INSERT INTO bill_detail (bill_id, product_id, quantity) VALUES ($1, $2, $3)",
			transactionID, billDetail.Product.ID, billDetail.Qty)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to create bill detail"})
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit"})
	}
	request.ID = transactionID
	response := models.CustomResponse{
		Message: "Transaction creatd successfully",
		Data:    request,
	}

	c.JSON(http.StatusOK, response)
}
