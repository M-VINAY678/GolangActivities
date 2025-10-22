package fetch

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// UsersID holds a list of user IDs
type UsersId struct {
	Users []struct {
		ID int `json:"id"`
	} `json:"users"`
}
