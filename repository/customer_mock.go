package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: "1", Name: "HadesGod", City: "Hell", Zipcode: "0000", Status: 0},
		{CustomerID: "2", Name: "TitonGod", City: "Heaven", Zipcode: "1111", Status: 1},
	}
	return customerRepositoryMock{customers: customers}
}

func (cm customerRepositoryMock) GetAll() ([]Customer, error) {
	return cm.customers, nil
}

func (cm customerRepositoryMock) GetById(id string) (*Customer, error) {
	for _, customer := range cm.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, mongo.ErrNoDocuments
}
