package controllers

import (
	"challenge-goapi/config"
	"challenge-goapi/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()
	request := c.Keys["data"].(models.Transaction)
	// request := c.MustGet("data").(models.Transaction)
	fmt.Println(request)

	// c.JSON(http.StatusOK, request)
	// return

	billDate, _ := time.Parse("02-01-2006", request.BillDate)
	entryDate, _ := time.Parse("02-01-2006", request.EntryDate)
	finishDate, _ := time.Parse("02-01-2006", request.FinishDate)
	FormattedBillDate := billDate.Format("2006-01-02")
	formattedEntryDate := entryDate.Format("2006-01-02")
	formattedFinishDate := finishDate.Format("2006-01-02")

	// Create a new transaction in the database using a transaction
	tx, err := db.Begin()
	if err != nil {
		// Failed to start transaction
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	// Insert the transaction record
	var transactionID string
	err = tx.QueryRow("INSERT INTO bill (bill_date, entry_date, finish_date, employee_id, customer_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		FormattedBillDate, formattedEntryDate, formattedFinishDate, request.EmployeeID, request.CustomerID).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		// failed to create transaction
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Insert the bill details
	for index, billDetail := range request.BillDetails {
		var billDetailId string
		var price int
		err = tx.QueryRow("SELECT price FROM product WHERE id = $1", billDetail.ProductID).Scan(&price)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get price"})
		}
		err = tx.QueryRow("INSERT INTO bill_detail (bill_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id",
			transactionID, billDetail.ProductID, billDetail.Qty).Scan(&billDetailId)
		if err != nil {
			tx.Rollback()
			// Failed to create bill detail
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bill detail"})
			return
		}
		request.BillDetails[index].ProductPrice = price
		request.BillDetails[index].BillID = transactionID
		request.BillDetails[index].ID = billDetailId
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		// Failed to commit
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit"})
	}
	request.ID = transactionID
	response := models.CustomResponse{
		Message: "Transaction created successfully",
		Data:    request,
	}

	c.JSON(http.StatusOK, response)
}

func GetTransactionByID(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()
	billID := c.Param("id")

	// Retrieve the transaction record from the database
	var transaction models.Transaction
	rows, err := db.Query(`
			SELECT t.id, to_char(t.bill_date, 'DD-MM-YYYY'), to_char(t.entry_date, 'DD-MM-YYYY'), to_char(t.finish_date, 'DD-MM-YYYY'),
				e.id, e.name, e.phone_number, e.address,
				c.id, c.name, c.phone_number, c.address,
				bd.id, bd.product_id, p.name, p.price, p.unit, bd.quantity
			FROM bill t
			JOIN employee e ON t.employee_id = e.id
			JOIN customer c ON t.customer_id = c.id
			JOIN bill_detail bd ON t.id = bd.bill_id
			JOIN product p ON bd.product_id = p.id
			WHERE t.id = $1
		`, billID)
	if err != nil {
		// Failed to retrieve transaction
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var billDetails []models.BillDetail
	for rows.Next() {
		var (
			transactionID   string
			billDate        string
			entryDate       string
			finishDate      string
			employeeID      string
			employeeName    string
			employeePhone   string
			employeeAddress string
			customerID      string
			customerName    string
			customerPhone   string
			customerAddress string
			billDetailID    string
			productID       string
			productName     string
			productPrice    int
			productUnit     string
			quantity        int
		)
		err := rows.Scan(
			&transactionID,
			&billDate,
			&entryDate,
			&finishDate,
			&employeeID,
			&employeeName,
			&employeePhone,
			&employeeAddress,
			&customerID,
			&customerName,
			&customerPhone,
			&customerAddress,
			&billDetailID,
			&productID,
			&productName,
			&productPrice,
			&productUnit,
			&quantity,
		)
		if err != nil {
			// Failed to retrieve transaction
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		transaction.ID = transactionID
		transaction.BillDate = billDate
		transaction.EntryDate = entryDate
		transaction.FinishDate = finishDate
		transaction.Employee = &models.Employee{
			ID:          employeeID,
			Name:        employeeName,
			PhoneNumber: employeePhone,
			Address:     employeeAddress,
		}
		transaction.Customer = &models.Customer{
			ID:          customerID,
			Name:        customerName,
			PhoneNumber: customerPhone,
			Address:     customerAddress,
		}

		billDetail := models.BillDetail{
			ID:     billDetailID,
			BillID: transactionID,
			Product: &models.Product{
				ID:    productID,
				Name:  productName,
				Price: productPrice,
				Unit:  productUnit,
			},
			ProductPrice: productPrice,
			Qty:          quantity,
		}
		billDetails = append(billDetails, billDetail)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to retreive transaction"})
		return
	}
	totalBill := 0
	for _, billDetail := range billDetails {
		totalBill += billDetail.ProductPrice * billDetail.Qty
	}

	transaction.BillDetails = billDetails
	transaction.TotalBill = totalBill

	response := models.CustomResponse{
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	}

	c.JSON(200, response)
}

func GetTransaction(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	transactionsMap := make(map[string]*models.Transaction)

	rows, err := db.Query(`
  SELECT t.id, to_char(t.bill_date, 'DD/MM/YYYY'), to_char(t.entry_date, 'DD/MM/YYYY'), to_char(t.finish_date, 'DD/MM/YYYY'),
         e.id, e.name, e.phone_number, e.address,
         c.id, c.name, c.phone_number, c.address,
         bd.id, bd.product_id, p.name, p.price, p.unit, bd.quantity
  FROM bill t
  JOIN employee e ON t.employee_id = e.id
  JOIN customer c ON t.customer_id = c.id
  JOIN bill_detail bd ON t.id = bd.bill_id
  JOIN product p ON bd.product_id = p.id
  ORDER BY t.id ASC
`)
	if err != nil {
		// Failed to retrieve transactions
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		var billDetail models.BillDetail
		var (
			employeeID      string
			employeeName    string
			employeePhone   string
			employeeAddress string
			customerID      string
			customerName    string
			customerPhone   string
			customerAddress string
			productID       string
			productName     string
			productPrice    int
			productUnit     string
		)
		err := rows.Scan(
			&transaction.ID,
			&transaction.BillDate,
			&transaction.EntryDate,
			&transaction.FinishDate,
			&employeeID,
			&employeeName,
			&employeePhone,
			&employeeAddress,
			&customerID,
			&customerName,
			&customerPhone,
			&customerAddress,
			&billDetail.ID,
			&productID,
			&productName,
			&productPrice,
			&productUnit,
			&billDetail.Qty,
		)
		billDetail.BillID = transaction.ID
		billDetail.ProductPrice = productPrice
		billDetail.Product = &models.Product{
			ID:    productID,
			Name:  productName,
			Price: productPrice,
			Unit:  productUnit,
		}
		transaction.Employee = &models.Employee{
			ID:          employeeID,
			Name:        employeeID,
			PhoneNumber: employeePhone,
			Address:     employeeAddress,
		}
		transaction.Customer = &models.Customer{
			ID:          customerID,
			Name:        customerName,
			PhoneNumber: customerPhone,
			Address:     customerAddress,
		}
		if err != nil {
			// Failed to retrieve transaction
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if _, ok := transactionsMap[transaction.ID]; !ok {
			transactionsMap[transaction.ID] = &transaction
		}
		transactionsMap[transaction.ID].BillDetails = append(transactionsMap[transaction.ID].BillDetails, billDetail)

	}
	transactions := make([]models.Transaction, 0, len(transactionsMap))

	// Calculate and set total bill for each transaction
	index := 0
	for _, transaction := range transactionsMap {
		transactions = append(transactions, *transaction)
		totalBill := 0
		for _, detail := range transactions[index].BillDetails {
			totalBill += detail.Product.Price * detail.Qty
		}
		transactions[index].TotalBill = totalBill
		index++
	}

	// Handle success and return transactions
	c.JSON(200, models.CustomResponse{
		Message: "success",
		Data:    transactions,
	})
}
