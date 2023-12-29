package models

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
type CustomResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
