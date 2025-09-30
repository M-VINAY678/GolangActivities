package main

import "fmt"

var carId = 0

// car data structure
type car struct {
	id             int
	make           string
	model          string
	year           string
	pricePerDay    int
	licensePlateNo string
}

func NewCar(model string, licenseNo string) car {
	s := car{
		id:             carId,
		make:           "india",
		model:          model,
		year:           "2025",
		pricePerDay:    2000,
		licensePlateNo: licenseNo,
	}
	return s
}

// Car details container type
type carDetails map[int]car //creating custom

func (carDetails carDetails) filterCarByPriceByModel(minPrice int, maxPrice int, model string) {
	var resultByPrice []int
	for key, value := range carDetails {
		if maxPrice != 0 && value.pricePerDay >= minPrice && value.pricePerDay <= maxPrice {
			resultByPrice = append(resultByPrice, key)
		} else if maxPrice == 0 && value.pricePerDay >= minPrice {
			resultByPrice = append(resultByPrice, key)
		}
	}
	if (minPrice == 0 && maxPrice == 0) && model != "" {
		for _, value := range carDetails {
			if value.model == model {
				fmt.Print("Car Id :", value)
			}
		}
	} else if model != "" {
		for _, value := range resultByPrice {
			if carDetails[value].model == model {
				fmt.Print("Car Id :", value)
			}
		}
	} else {
		for _, value := range resultByPrice {
			fmt.Print("Car Id :", value)
		}
	}
}
func (carDetails carDetails) addCarDetails(model string, licenseNo string) {
	s := NewCar(model, licenseNo)
	carDetails[carId] = s
	availability[carId] = true
	fmt.Print(carDetails)
	carId++
}
