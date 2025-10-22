package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// fetches user details concurrently for each user ID
func FetchUsers(UsersId UsersId, wg *sync.WaitGroup, ch chan User) {
	semaphore := make(chan struct{}, 5) // limit to 5 concurrent requests
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
				return
			}
			defer response.Body.Close()
			byte, _ := io.ReadAll(response.Body)
			var user User
			err = json.Unmarshal(byte, &user)
			if err != nil {
				fmt.Print("Their is some error while Unmarshaling with the user id : ", id)
				return
			}
			ch <- user
		}(user.ID)
		wg.Add(1)
		// Separate goroutine to read from channel
		go func(ch chan User) {
			defer wg.Done()
			fmt.Println(<-ch)
		}(ch)
	}
}

// fetching a list of user IDs from the provided API endpoint
func FetchUserIDs(url string) (UsersId, error) {
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
