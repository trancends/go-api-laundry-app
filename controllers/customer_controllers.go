package controllers

import (
	"challenge-goapi/config"
	"challenge-goapi/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = config.ConnectDB()

func GetAllCustomer(c *gin.Context) {
	query := "SELECT id,name,phone_number,address FROM customer"

	var rows *sql.Rows
	var err error

	rows, err = db.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	defer rows.Close()

	var Customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, customer.Address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		Customers = append(Customers, customer)
	}

	if len(Customers) > 0 {
		c.JSON(http.StatusOK, Customers)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer Not Found"})
	}
}

func CreateCustomer(c *gin.Context) {
	var newCustomer models.Customer
	err := c.ShouldBind(&newCustomer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO customer (name,phone_number,address) VALUES ($1,$2,$3) RETURNING id"

	var customerId string
	err = db.QueryRow(query, newCustomer.Name, newCustomer.PhoneNumber, newCustomer.Address).Scan(&customerId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newCustomer.Id = customerId
	c.JSON(http.StatusCreated, newCustomer)
}
