package service

type CustomerResponse struct {
	CustomerID string `json:"customerId,omitempty"`
	Name       string `json:"name,omitempty"`
	Status     int    `json:"status,omitempty"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(string) (*CustomerResponse, error)
}
