package car

import "fmt"

var carId = 0

// car data structure
type Car struct {
	id             int
	make           string
	model          string
	year           string
	pricePerDay    int
	licensePlateNo string
}

func NewCar(model string, licenseNo string) Car {
	s := Car{
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
type Cars map[int]Car //creating custom

func (Cars Cars) Filter(minPrice int, maxPrice int, model string) {
	var resultByPrice []int
	for key, value := range Cars {
		if maxPrice != 0 && value.pricePerDay >= minPrice && value.pricePerDay <= maxPrice {
			resultByPrice = append(resultByPrice, key)
		} else if maxPrice == 0 && value.pricePerDay >= minPrice {
			resultByPrice = append(resultByPrice, key)
		}
	}
	if (minPrice == 0 && maxPrice == 0) && model != "" {
		for _, value := range Cars {
			if value.model == model {
				fmt.Print("Car Id :", value)
			}
		}
	} else if model != "" {
		for _, value := range resultByPrice {
			if Cars[value].model == model {
				fmt.Print("Car Id :", value)
			}
		}
	} else {
		for _, value := range resultByPrice {
			fmt.Print("Car Id :", value)
		}
	}
}
func (Cars Cars) Add(model string, licenseNo string) {
	s := NewCar(model, licenseNo)
	Cars[carId] = s
	fmt.Print(Cars)
	carId++
}
