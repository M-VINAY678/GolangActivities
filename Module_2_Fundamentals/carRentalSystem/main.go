package main

import (
	"fmt"
)

func main() {
	allCars := make(carDetails)
	on := true
	for on {
		fmt.Println("Enter 1: To add customer Details \nEnter 2: To add Car Details")
		fmt.Println("Enter 3: Get Available Cars \nEnter 4: To Reserve the Car \nEnter 5: To modify Reservation \nEnter 6: To cancel Reservation \nEnter 7: To Filter car Based on Price and model \nEnter 8: To Exit")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addCustomerDetails()
		case 2:
			allCars.addCarDetails()
		case 3:
			getAvailabilityOfCar()
		case 4:
			carReservation()
		case 5:
			modifyReservation()
		case 6:
			cancelReservation()
		case 7:
			allCars.filterCarByPriceBymodel()
		case 8:
			on = false
		}
	}
}
