package main

import (
	"fmt"
	"time"
)

// Availability of Car := key as carId and boolean as value
var availability = make(map[int]bool)

func getAvailabilityOfCar() {
	layout := "2025-09-23"
	fmt.Println("Enter Start Date in this format '2025-09-23'")
	var startDate string
	fmt.Scanln(&startDate)
	fmt.Println("Enter End Date in this format '2025-09-23'")
	var endDate string
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
