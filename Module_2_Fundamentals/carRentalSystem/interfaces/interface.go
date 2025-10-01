package interfaces

type ReservationManagement interface {
	Cancel(int)
	Book(int, int, string, string)
	Update(int, string, string)
	AvailableCars(string, string)
}

type CarManagement interface {
	Filter(int, int, string)
	Add(string, string)
}

type CustomerAdder interface {
	Add(string, string)
}
