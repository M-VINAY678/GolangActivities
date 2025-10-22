package main

import (
	"fmt"
	"module_10_Concurrency/fetch" // import your local fetch package
	"sync"
)

func main() {
	// Fetch User IDs
	usersID, err := fetch.FetchUserIDs("https://dummyjson.com/users?limit=10&select=id")
	if err != nil {
		fmt.Println("Error fetching user IDs:", err)
		return
	}
	fmt.Println("Fetched User IDs:", usersID)
	// Fetch user details concurrently
	wg := &sync.WaitGroup{}
	ch := make(chan fetch.User, 5)

	fetch.FetchUsers(usersID, wg, ch)

	wg.Wait()
	close(ch)
}
