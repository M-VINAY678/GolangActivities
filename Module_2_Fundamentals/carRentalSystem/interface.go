package main

type ReservationManagement interface {
	cancelReservation()
	carReservation()
	modifyReservation()
	getAvailabilityOfCar()
}

type CarManagement interface {
	filterCarByPriceBymodel()
	addCarDetails()
}

type CustomerAdder interface {
	addCustomerDetails()
}
