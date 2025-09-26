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

func (customerDetails customerDetails) addCustomerDetails() {
	fmt.Println("Enter customer name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Enter customer license no")
	var licenseNo string
	fmt.Scanln(&licenseNo)
	s := customer{
		id:        customerId,
		name:      name,
		email:     "bc@123@gmail.com",
		phoneNo:   "812****957",
		licenseNo: licenseNo,
	}
	customerDetails[customerId] = s
	customerId++
	fmt.Println(customerDetails)
}
