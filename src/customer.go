package main

import (
	"fmt"
)

type customer struct {
	id int64
	companyName string
	firstName string
	lastName string
	address string
	postCode string
	contactNumber string
	emailAddress string
}

type customers []customer

func loadCustomers() customers {
	currentCustomers := []customer{
		{1, "Random Pty Ltd", "Bob" ,"Smith", "100 Test St", "1234", "12345678", "bob@smith.com"},
		{2, "Not So Random Pty Ltd", "Sue" ,"James", "200 Test St", "2345", "23456789", "sue@james.com"},
	}

	return currentCustomers
}

func (c customers) lookupCustomer(customerId int64) customer {
	var foundCustomer customer

	for _, customer := range c {
		if (customer.id == customerId) {
			fmt.Println("Found customer")
			foundCustomer = customer;
			break
		}
	}

	return foundCustomer
}