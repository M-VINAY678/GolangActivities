package main

import "fmt"

var customerId = 0

// customer data structure
type customer struct {
	id        int
	name      string
	email     string
	phoneNo   string
	licenseNo string
}

// customer Details
type customerDetails map[int]customer

func NewCustomer(name string, licenseNo string) customer {
	s := customer{
		id:        customerId,
		name:      name,
		email:     "bc@123@gmail.com",
		phoneNo:   "812****957",
		licenseNo: licenseNo,
	}
	customerId++
	return s
}
func (customerDetails customerDetails) addCustomerDetails(name string, licenseNo string) {
	c := NewCustomer(name, licenseNo)
	customerDetails[c.id] = c
	fmt.Println("Customer Details is Added")
}
