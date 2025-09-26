package main

import (
	"fmt"
)

func main() {
	// Reservation Details
	var reservationManager ReservationManagement = make(reservationDetails)
	var carManager CarManagement = make(carDetails)
	var customerAdder CustomerAdder = make(customerDetails)
	on := true
	for on {
		fmt.Println("Enter 1: To add customer Details \nEnter 2: To add Car Details")
		fmt.Println("Enter 3: Get Available Cars \nEnter 4: To Reserve the Car \nEnter 5: To modify Reservation \nEnter 6: To cancel Reservation \nEnter 7: To Filter car Based on Price and model \nEnter 8: To Exit")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			customerAdder.addCustomerDetails()
		case 2:
			carManager.addCarDetails()
		case 3:
			reservationManager.getAvailabilityOfCar()
		case 4:
			reservationManager.carReservation()
		case 5:
			reservationManager.modifyReservation()
		case 6:
			reservationManager.cancelReservation()
		case 7:
			carManager.filterCarByPriceBymodel()
		case 8:
			on = false
		}
	}
}
