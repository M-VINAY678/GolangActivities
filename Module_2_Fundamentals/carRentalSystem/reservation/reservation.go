package reservation

import (
	"fmt"
	"time"
)

var reservationId = 0

// reservation data structure
type Reservation struct {
	id         int
	carId      int
	customerId int
	startDate  string
	endDate    string
}

// Reservation details container type
type Reservations map[int]Reservation

func newReservation(customerID int, carID int, startDate string, endDate string) Reservation {
	s := Reservation{
		id:         reservationId,
		carId:      carID,
		customerId: customerID,
		startDate:  startDate,
		endDate:    endDate,
	}
	return s
}
func validateDates(startDate string, endDate string, Reservations Reservations) bool {
	layout := "2025-09-23"
	start, _ := time.Parse(layout, startDate)
	end, _ := time.Parse(layout, endDate)
	for _, value := range Reservations {
		startDate, _ := time.Parse(layout, value.startDate)
		endDate, _ := time.Parse(layout, value.endDate)
		if !(start.Before(endDate) && startDate.After(end)) {
			return false
		}
	}
	return true
}

func (Reservations Reservations) Cancel(reservationID int) {
	delete(Reservations, reservationID)
}

func (Reservations Reservations) Update(reservationID int, startDate string, endDate string) {

	temp := Reservations[reservationID]
	temp.startDate = startDate
	temp.endDate = endDate
	Reservations[reservationID] = temp
	fmt.Println(temp)
}
func (Reservations Reservations) Book(customerID int, carID int, startDate string, endDate string) {
	if customerID >= customerID || customerID < 0 {
		fmt.Println("wrong entry of customerId")
		return
	}
	if carID >= carID || carID < 0 {
		fmt.Println("wrong entry of customerId")
		return
	}
	if !validateDates(startDate, endDate, Reservations) {
		return
	}
	s := newReservation(customerID, carID, startDate, endDate)
	Reservations[reservationId] = s
	reservationId++
	fmt.Println(Reservations)
	//fmt.Print(s.startDate.Format("2006-01-02 15:04:05"))
}

var CarAvailability = make(map[int]bool)

func (Reservations Reservations) AvailableCars(startDate string, endDate string) {
	layout := "2025-09-23"
	start, _ := time.Parse(layout, startDate)
	end, _ := time.Parse(layout, endDate)
	for _, value := range Reservations {
		startDate, _ := time.Parse(layout, value.startDate)
		endDate, _ := time.Parse(layout, value.endDate)
		if !(start.Before(endDate) && startDate.After(end)) {
			CarAvailability[value.carId] = false
		}
	}
	fmt.Println("Available Cars : ")
	for key, value := range CarAvailability {
		if value {
			fmt.Println("Car ID :", key)
		}
	}
}
