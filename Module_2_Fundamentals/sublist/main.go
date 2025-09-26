package main

import "fmt"

func sublist(list1 []int, list2 []int) bool {
	var j int = 0
	for i := 0; i < len(list1); i++ {
		if list2[j] != list1[i] {
			j = 0
			if list2[j] == list1[i] {
				j++
			}
		} else {
			j++
		}
		if j == len(list2) {
			return true
		}
	}
	return false
}
func main() {
	var list1 []int = []int{2, 3, 4}
	var list2 []int = []int{1, 2, 3, 2, 3, 4}
	if sublist(list1, list2) {
		if len(list1) == len(list2) {
			fmt.Println("A and B are equal")
		} else {
			fmt.Println("A is a superlist of B")
		}
	} else if sublist(list2, list1) {
		fmt.Println("B is a superlist of A")
	} else {
		fmt.Println("A and B are unequal")
	}
}
