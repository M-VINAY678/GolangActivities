package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
type UsersId struct {
	Users []struct {
		ID int `json:"id"`
	} `json:"users"`
}

func fetchUser(UsersId UsersId, wg *sync.WaitGroup, ch chan User) {
	semaphore := make(chan struct{}, 5)
	for _, user := range UsersId.Users {
		wg.Add(1)
		semaphore <- struct{}{} //occupy a slot
		go func(id int) {
			defer wg.Done()
			defer func() {
				<-semaphore
			}()
			url := fmt.Sprintf("https://dummyjson.com/users/%d", id)
			response, err := http.Get(url)
			if err != nil {
				fmt.Println("Error while fetching the user , whose id is : ", id)
			}
			defer response.Body.Close()
			byte, _ := io.ReadAll(response.Body)
			var user User
			err = json.Unmarshal(byte, &user)
			if err != nil {
				fmt.Print("Their is some error while Unmarshaling with the user id : ", id)
			}
			// fmt.Println(id, ":===>", user)
			ch <- user
		}(user.ID)
		wg.Add(1)
		go func(ch chan User) {
			defer wg.Done()
			fmt.Println(<-ch)
		}(ch)
	}
}
func fetchId(url string) (UsersId, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print("Their is some error while fetching users details")
		return UsersId{}, err
	}
	defer response.Body.Close()
	byte, _ := io.ReadAll(response.Body)
	var result UsersId
	err = json.Unmarshal(byte, &result)
	if err != nil {
		fmt.Print("Their is some error while Unmarshaling")
		return UsersId{}, err
	}
	return result, nil

}
func main() {
	//response, err := http.Get("https://dummyjson.com/users?limit=20")

	UsersId, err := fetchId("https://dummyjson.com/users?limit=10&select=id")
	if err != nil {
		fmt.Println("Error while fetching User's Id", err)
	}
	fmt.Println(UsersId)
	wg := &sync.WaitGroup{}
	ch := make(chan User, 5)
	fetchUser(UsersId, wg, ch)
	wg.Wait()
}
