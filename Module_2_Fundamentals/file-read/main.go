package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Println("Enter the file path")
	fmt.Scanf("%s", &path)
	fileContent, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist or Given a wrong path")
			return
		}
		//log.Fatal(err)// we use log.Fatal so that it terminates at this point and gives non-zero exit status.
		fmt.Println("Error During File Read")
		return
	}
	fmt.Println(string(fileContent))

}
