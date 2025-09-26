package main

import (
	"fmt"
	"time"
)

var reservationId = 0

// Reservation details container type
type reservationDetails map[int]reservation

// reservation data structure
type reservation struct {
	id         int
	carId      int
	customerId int
	startDate  string
	endDate    string
}

func validate(startDate string, endDate string, reservationDetails reservationDetails) bool {
	layout := "2025-09-23"
	start, _ := time.Parse(layout, startDate)
	end, _ := time.Parse(layout, endDate)
	for _, value := range reservationDetails {
		startDate, _ := time.Parse(layout, value.startDate)
		endDate, _ := time.Parse(layout, value.endDate)
		if !(start.Before(endDate) && startDate.After(end)) {
			return false
		}
	}
	return true
}

func (reservationDetails reservationDetails) cancelReservation() {
	fmt.Println("Enter Reservation Id")
	var reservationID int
	fmt.Scanln(&reservationID)
	delete(reservationDetails, reservationID)
}

func (reservationDetails reservationDetails) modifyReservation() {
	fmt.Println("Enter Reservation Id")
	var reservationID int
	fmt.Scanln(&reservationID)
	fmt.Println("Enter Start Date in this format '2025-09-23'")
	var startDate string
	fmt.Scanln(&startDate)
	fmt.Println("Enter End Date in this format '2025-09-23'")
	var endDate string
	fmt.Scanln(&endDate)
	temp := reservationDetails[reservationID]
	temp.startDate = startDate
	temp.endDate = endDate
	reservationDetails[reservationID] = temp
	fmt.Println(temp)
}

func (reservationDetails reservationDetails) carReservation() {
	fmt.Println("Enter customer Id")
	var customerID int
	fmt.Scanln(&customerID)
	if customerID >= customerId || customerID < 0 {
		fmt.Println("wrong entry of customerId")
		return
	}
	fmt.Println("Enter Car Id ")
	var carID int
	fmt.Scanln(&carID)
	if carID >= carId || carID < 0 {
		fmt.Println("wrong entry of customerId")
		return
	}
	fmt.Println("Enter Start Date in this format '2025-09-23'")
	var startDate string
	fmt.Scanln(&startDate)
	fmt.Println("Enter End Date in this format '2025-09-23'")
	var endDate string
	if !validate(startDate, endDate, reservationDetails) {
		return
	}

	fmt.Scanln(&endDate)
	s := reservation{
		id:         reservationId,
		carId:      carID,
		customerId: customerID,
		startDate:  startDate,
		endDate:    endDate,
	}
	reservationDetails[reservationId] = s
	reservationId++
	fmt.Println(reservationDetails)
	//fmt.Print(s.startDate.Format("2006-01-02 15:04:05"))
}

var availability = make(map[int]bool)

func (reservationDetails reservationDetails) getAvailabilityOfCar() {
	layout := "2025-09-23"
	fmt.Println("Enter Start Date in this format '2025-09-23'")
	var startDate string
	fmt.Scanln(&startDate)
	fmt.Println("Enter End Date in this format '2025-09-23'")
	var endDate string
	fmt.Scanln(&endDate)
	start, _ := time.Parse(layout, startDate)
	end, _ := time.Parse(layout, endDate)
	for _, value := range reservationDetails {
		startDate, _ := time.Parse(layout, value.startDate)
		endDate, _ := time.Parse(layout, value.endDate)
		if !(start.Before(endDate) && startDate.After(end)) {
			availability[value.carId] = false
		}
	}
	fmt.Println("Available Cars : ")
	for key, value := range availability {
		if value {
			fmt.Println("Car ID :", key)
		}
	}
}
