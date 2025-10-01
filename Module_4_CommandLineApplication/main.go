package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func stringToInt(value string) (num int, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error : ", r)
			ok = false
		}
	}()
	num, err := strconv.Atoi(value)
	if err != nil {
		panic("It is not a numeric type")
	}
	ok = true
	return
}
func main() {
	var numbersString string
	sum := 0
	flag.StringVar(&numbersString, "numbersString", "", "comma separated list of numbers in string")
	flag.Parse()
	numberString := strings.Split(numbersString, ",")
	for _, value := range numberString {
		//Type conversion from string to int
		num, ok := stringToInt(value)
		if !ok {
			continue
		}
		fmt.Println("Adding : ", num)
		sum += num
	}
	fmt.Println(sum)

}
