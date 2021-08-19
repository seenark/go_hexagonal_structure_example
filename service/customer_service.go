package service

import (
	"bank/repository"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(cusRepo repository.CustomerRepository) CustomerService {
	return customerService{
		customerRepository: cusRepo,
	}
}

func (cs customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := cs.customerRepository.GetAll()
	if err != nil {
		// this is technical error should not sent to user
		// should be keep in log file
		fmt.Printf("err: %v\n", err)
		return []CustomerResponse{}, err
	}

	// transfer data to response model
	cusResponses := []CustomerResponse{}
	for _, customer := range customers {
		cusRes := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		cusResponses = append(cusResponses, cusRes)
	}
	return cusResponses, nil
}

func (cs customerService) GetCustomer(id string) (*CustomerResponse, error) {
	customer, err := cs.customerRepository.GetById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("not found")
		}
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	cusResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &cusResponse, nil
}
