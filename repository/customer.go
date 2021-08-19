package repository

type Customer struct {
	CustomerID  string `json:"customerId,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	City        string `json:"city,omitempty" bson:"city,omitempty"`
	Zipcode     string `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
	Status      int    `json:"status,omitempty" bson:"status,omitempty"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(string) (*Customer, error)
}
