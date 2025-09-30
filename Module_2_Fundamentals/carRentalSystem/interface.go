package main

type ReservationManagement interface {
	cancelReservation(int)
	carReservation(int, int, string, string)
	modifyReservation(int, string, string)
	getAvailabilityOfCar(string, string)
}

type CarManagement interface {
	filterCarByPriceByModel(int, int, string)
	addCarDetails(string, string)
}

type CustomerAdder interface {
	addCustomerDetails(string, string)
}
