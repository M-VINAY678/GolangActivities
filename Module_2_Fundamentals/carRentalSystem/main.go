package main

import (
	"carRentalSystem/car"
	"carRentalSystem/customer"
	"carRentalSystem/interfaces"
	"carRentalSystem/reservation"
	"fmt"
)

func main() {
	// Reservation Details
	var reservationManager interfaces.ReservationManagement = make(reservation.Reservations)
	var carManager interfaces.CarManagement = make(car.Cars)
	var customerAdder interfaces.CustomerAdder = make(customer.Customers)
	on := true
	for on {
		fmt.Println("Enter 1: To add customer Details \nEnter 2: To add Car Details")
		fmt.Println("Enter 3: Get Available Cars \nEnter 4: To Reserve the Car \nEnter 5: To modify Reservation \nEnter 6: To cancel Reservation \nEnter 7: To Filter car Based on Price and model \nEnter 8: To Exit")
		var option int
		fmt.Scanln(&option)
		switch option {

		case 1:
			fmt.Println("Enter customer name")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Enter customer license no")
			var licenseNo string
			fmt.Scanln(&licenseNo)
			customerAdder.Add(name, licenseNo)

		case 2:
			fmt.Println("Enter car model")
			var model string
			fmt.Scanln(&model)
			fmt.Println("Enter car's license plate no")
			var licenseNo string
			fmt.Scanln(&licenseNo)
			carManager.Add(model, licenseNo)

		case 3:
			fmt.Println("Enter Start Date in this format '2025-09-23'")
			var startDate string
			fmt.Scanln(&startDate)
			fmt.Println("Enter End Date in this format '2025-09-23'")
			var endDate string
			fmt.Scanln(&endDate)
			reservationManager.AvailableCars(startDate, endDate)

		case 4:
			fmt.Println("Enter customer Id")
			var customerID int
			fmt.Scanln(&customerID)
			fmt.Println("Enter Car Id ")
			var carID int
			fmt.Scanln(&carID)
			fmt.Println("Enter Start Date in this format '2025-09-23'")
			var startDate string
			fmt.Scanln(&startDate)
			fmt.Println("Enter End Date in this format '2025-09-23'")
			var endDate string
			fmt.Scanln(&endDate)
			reservationManager.Book(customerID, carID, startDate, endDate)

		case 5:
			fmt.Println("Enter Reservation Id")
			var reservationID int
			fmt.Scanln(&reservationID)
			fmt.Println("Enter Start Date in this format '2025-09-23'")
			var startDate string
			fmt.Scanln(&startDate)
			fmt.Println("Enter End Date in this format '2025-09-23'")
			var endDate string
			fmt.Scanln(&endDate)
			reservationManager.Update(reservationID, startDate, endDate)

		case 6:
			fmt.Println("Enter Reservation Id")
			var reservationID int
			fmt.Scanln(&reservationID)
			reservationManager.Cancel(reservationID)

		case 7:
			fmt.Println("If you Want to skip filter just Press Enter ")
			fmt.Println("Enter Minimum price per day")
			var minPrice int
			fmt.Scanln(&minPrice)
			fmt.Println("Enter Maximum price per day")
			var maxPrice int
			fmt.Scanln(&maxPrice)
			fmt.Println("Enter car model")
			var model string
			fmt.Scanln(&model)
			carManager.Filter(minPrice, maxPrice, model)

		case 8:
			on = false
		}
	}
}
