package models

// type Request struct {
// 	ID          string       `json:"id"`
// 	BillDate    string       `json:"billDate"`
// 	EntryDate   string       `json:"entryDate"`
// 	FinishDate  string       `json:"finishDate"`
// 	EmployeeID  string       `json:"employeeId"`
// 	CustomerID  string       `json:"customerId"`
// 	BillDetails []BillDetail `json:"billDetails"`
// }

// Employee represents an employee object
type Employee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type BillDetail struct {
	ID           string   `json:"id"`
	BillID       string   `json:"billId"`
	ProductID    string   `json:"productId,omitempty"`
	Product      *Product `json:"product,omitempty"`
	ProductPrice int      `json:"productPrice"`
	Qty          int      `json:"qty"`
}

type Transaction struct {
	ID          string       `json:"id"`
	BillDate    string       `json:"billDate"`
	EntryDate   string       `json:"entryDate"`
	FinishDate  string       `json:"finishDate"`
	EmployeeID  string       `json:"employeeId,omitempty"`
	Employee    *Employee    `json:"employee,omitempty"`
	CustomerID  string       `json:"CustomerID,omitempty"`
	Customer    *Customer    `json:"customer,omitempty"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill   int          `json:"totalBill,omitempty"`
}
