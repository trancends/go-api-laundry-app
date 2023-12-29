package controllers

import (
	"challenge-goapi/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	query := "SELECT id,name,price,int FROM product"

	var rows *sql.Rows
	var err error

	rows, err = db.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, product.Unit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		products = append(products, product)
	}

	if len(products) > 0 {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
	}
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	query := "SELECT id,name,price,unit FROM product WHERE id = $1"

	fmt.Println(product)
	err := db.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Price, &product.Unit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}
	response := models.CustomResponse{
		Message: "successfully retreived product",
		Data:    product,
	}
	c.JSON(http.StatusOK, response)
}

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	err := c.ShouldBind(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO product (name,price,unit) VALUES ($1,$2,$3) RETURNING id"

	var productId string
	fmt.Println(newProduct)
	err = db.QueryRow(query, newProduct.Name, newProduct.Price, newProduct.Unit).Scan(&productId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newProduct.Id = productId

	response := models.CustomResponse{
		Message: "product successfully created",
		Data:    newProduct,
	}

	c.JSON(http.StatusCreated, response)
}

func UpdateProductById(c *gin.Context) {
	productId := c.Param("id")

	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Id = productId

	updateQuery := `UPDATE product SET name = $2, price = $3, unit = $4 WHERE id = $1`
	_, err := db.Exec(updateQuery, productId, product.Name, product.Price, product.Unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	response := models.CustomResponse{
		Message: "product successfully updated",
		Data:    product,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")

	deleteQuery := `DELETE FROM product  WHERE id = $1`
	_, err := db.Exec(deleteQuery, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.CustomResponse{
		Message: "product successfully deleted",
		Data:    "OK",
	}

	c.JSON(http.StatusOK, response)
}
