package service

import (
	"fmt"

	"github.com/myrachanto/amicroservice/customermicroservice/httperors"
	"github.com/myrachanto/amicroservice/customermicroservice/model"
)

var (
	CustomerService customerService = customerService{}

	customers     = map[int]*model.Customer{}
	currentId int = 1
)

type customerService struct{}

func (service customerService) Create(customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	if err := customer.Validate(); err != nil {
		return nil, err
	}

	customer.ID = currentId
	currentId++
	customers[customer.ID] = customer
	return customer, nil
}

func (service customerService) GetOne(id int) (*model.Customer, *httperors.HttpError) {
	if customer := customers[id]; customer != nil {
		return customer, nil
	}
	return nil, httperors.NewNotFoundError(fmt.Sprintf("Customer with Id %d not found", id))
}

func (service customerService) GetAll(cust map[int]*model.Customer) (map[int]*model.Customer, *httperors.HttpSuccess) {
	cust = customers
	if cust == nil {
		return nil, httperors.NewNoResultsMessage("No results found")
	}

	return cust, nil
}

func (service customerService) Update(id int, customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	cust := customers[id]
	if cust == nil {
		return nil, httperors.NewNotFoundError("No results found")
	}
	if customer.Name == "" {
		customer.Name = cust.Name
	}
	if customer.Company == "" {
		customer.Company = cust.Company
	}
	if customer.Email == "" {
		customer.Email = cust.Email
	}
	if customer.Phone == "" {
		customer.Phone = cust.Phone
	}
	if customer.Address == "" {
		customer.Address = cust.Address
	}
	
	customers[id].Name = customer.Name
	customers[id].Company = customer.Company
	customers[id].Email = customer.Email
	customers[id].Phone = customer.Phone
	customers[id].Address = customer.Address
	

	return customer, nil
}
func (service customerService) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	if customer := customers[id]; customer == nil {
		return nil, httperors.NewNotFoundError("No results found")
	}
	delete(customers, id)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
