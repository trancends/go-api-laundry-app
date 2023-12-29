package controllers

import (
	"challenge-goapi/config"
	"challenge-goapi/models"
	"database/sql"
	"fmt"
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

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	query := "SELECT id,name,phone_number,address FROM customer WHERE id = $1"

	fmt.Println(customer)
	err := db.QueryRow(query, id).Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer Not Found"})
		return
	}
	response := models.CustomResponse{
		Message: "successfully retreived customer",
		Data:    customer,
	}
	c.JSON(http.StatusOK, response)
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
	fmt.Println(newCustomer)
	err = db.QueryRow(query, newCustomer.Name, newCustomer.PhoneNumber, newCustomer.Address).Scan(&customerId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newCustomer.Id = customerId
	c.JSON(http.StatusCreated, newCustomer)
}

func UpdateCustomerById(c *gin.Context) {
	customerId := c.Param("id")

	var customer models.Customer

	if err := c.ShouldBind(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.Id = customerId

	updateQuery := `UPDATE customer SET name = $2, phone_number = $3, address = $4 WHERE id = $1`
	_, err := db.Exec(updateQuery, customerId, customer.Name, customer.PhoneNumber, customer.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Customer"})
		return
	}

	response := models.CustomResponse{
		Message: "successfully updated customer",
		Data:    customer,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteCustomer(c *gin.Context) {
	customerId := c.Param("id")

	deleteQuery := `DELETE FROM customer  WHERE id = $1`
	_, err := db.Exec(deleteQuery, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.CustomResponse{
		Message: "successfully deleted customer",
		Data:    "OK",
	}

	c.JSON(http.StatusOK, response)
}
