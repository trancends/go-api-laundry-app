package models

// Employee represents an employee object
type Employee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type BillDetail struct {
	ID           string  `json:"id"`
	BillID       string  `json:"billId"`
	Product      Product `json:"product"`
	ProductPrice int     `json:"productPrice"`
	Qty          int     `json:"qty"`
}

type Transaction struct {
	ID          string       `json:"id"`
	BillDate    string       `json:"billDate"`
	EntryDate   string       `json:"entryDate"`
	FinishDate  string       `json:"finishDate"`
	Employee    Employee     `json:"employee"`
	Customer    Customer     `json:"customer"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill   int          `json:"totalBill"`
}
