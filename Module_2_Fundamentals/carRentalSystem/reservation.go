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

func NewReservation(customerID int, carID int, startDate string, endDate string) reservation {
	s := reservation{
		id:         reservationId,
		carId:      carID,
		customerId: customerID,
		startDate:  startDate,
		endDate:    endDate,
	}
	return s
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

func (reservationDetails reservationDetails) cancelReservation(reservationID int) {

	delete(reservationDetails, reservationID)
}

func (reservationDetails reservationDetails) modifyReservation(reservationID int, startDate string, endDate string) {

	temp := reservationDetails[reservationID]
	temp.startDate = startDate
	temp.endDate = endDate
	reservationDetails[reservationID] = temp
	fmt.Println(temp)
}
func (reservationDetails reservationDetails) carReservation(customerID int, carID int, startDate string, endDate string) {
	if customerID >= customerId || customerID < 0 {
		fmt.Println("wrong entry of customerId")
		return
	}
	if carID >= carId || carID < 0 {
		fmt.Println("wrong entry of customerId")
		return
	}
	if !validate(startDate, endDate, reservationDetails) {
		return
	}
	s := NewReservation(customerID, carID, startDate, endDate)
	reservationDetails[reservationId] = s
	reservationId++
	fmt.Println(reservationDetails)
	//fmt.Print(s.startDate.Format("2006-01-02 15:04:05"))
}

var availability = make(map[int]bool)

func (reservationDetails reservationDetails) getAvailabilityOfCar(startDate string, endDate string) {
	layout := "2025-09-23"
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
