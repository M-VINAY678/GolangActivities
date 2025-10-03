package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Apk       string
}

func UnStructuredResponse(body []byte) {
	fields := []string{"firstName", "lastName", "age", "email"}

	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Their is error while unmarshalling the response")
	}
	users := data["users"].([]interface{})
	for i, u := range users {
		user := u.(map[string]interface{})
		fmt.Println("User : ", i)
		for _, value := range fields {
			fmt.Println(value, " ==> ", user[value])
		}
	}
}
func writingToLocalFile(Users []User) {
	fileData, err := json.Marshal(Users)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	// writing to user.json file
	err = os.WriteFile("users.json", fileData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Data saved to users.json successfully!")
}
func structuredResponse(body []byte) {
	//fmt.Print(string(body))
	var result struct {
		Users []User `json:"users"`
	}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Print("Their is some error while Unmarshaling")
		return
	}
	for k, v := range result.Users {
		fmt.Println(k, ",", v)
	}
	writingToLocalFile(result.Users)
}
func main() {
	response, err := http.Get("https://dummyjson.com/users?limit=100")
	if err != nil {
		fmt.Print("Their is some error while fetching users details")
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	//When we know the struture of Response data
	structuredResponse(body)
	//when you don't know the struture of Respone data
	UnStructuredResponse(body)
}
