package customer

import (
	"carRentalSystem/reservation"
	"fmt"
)

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
type Customers map[int]customer

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
func (Customers Customers) Add(name string, licenseNo string) {
	c := NewCustomer(name, licenseNo)
	Customers[c.id] = c
	reservation.CarAvailability[c.id] = true
	fmt.Println("Customer Details is Added")
}
