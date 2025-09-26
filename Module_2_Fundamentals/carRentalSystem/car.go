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

// Car details container type
type carDetails map[int]car //creating custom

func (carDetails carDetails) filterCarByPriceBymodel() {
	var resultByPrice []int
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
func (carDetails carDetails) addCarDetails() {
	fmt.Println("Enter car model")
	var model string
	fmt.Scanln(&model)
	fmt.Println("Enter car's license plate no")
	var licenseNo string
	fmt.Scanln(&licenseNo)
	s := car{
		id:             carId,
		make:           "india",
		model:          model,
		year:           "2025",
		pricePerDay:    2000,
		licensePlateNo: licenseNo,
	}
	carDetails[carId] = s
	availability[carId] = true
	fmt.Print(carDetails)
	carId++
}
